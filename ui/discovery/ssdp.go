/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package discovery

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"text/template"
	"time"

	"github.com/gofrs/uuid"
	"github.com/koron/go-ssdp"
	"github.com/mysteriumnetwork/node/core/ip"
	"github.com/mysteriumnetwork/node/metadata"
	"github.com/mysteriumnetwork/node/requests"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const deviceDescriptionTemplate = `<?xml version="1.0"?>
<root xmlns="urn:schemas-upnp-org:device-1-0" configId="1">
	<specVersion>
		<major>1</major>
		<minor>1</minor>
	</specVersion>
	<device>
		<deviceType>urn:schemas-upnp-org:device:node:1</deviceType>
		<friendlyName>Mysterium Node</friendlyName>

		<manufacturer>Mysterium Network</manufacturer>
		<manufacturerURL>https://mysterium.network/</manufacturerURL>

		<modelName>Raspberry Node</modelName>
		<modelNumber>{{.Version}}</modelNumber>
		<modelURL>https://mysterium.network/node/</modelURL>

		<UDN>uuid:{{.UUID}}</UDN>
		<presentationURL>{{.URL}}</presentationURL>
	</device>
</root>`

type ssdpServer struct {
	uiPort     int
	uuid       string
	ssdp       *ssdp.Advertiser
	quit       chan struct{}
	once       sync.Once
	httpClient *requests.HTTPClient
}

func newSSDPServer(uiPort int, httpClient *requests.HTTPClient) *ssdpServer {
	return &ssdpServer{
		uiPort:     uiPort,
		quit:       make(chan struct{}),
		httpClient: httpClient,
	}
}

func (ss *ssdpServer) Start() (err error) {
	ss.uuid, err = generateUUID()
	if err != nil {
		return errors.Wrap(err, "failed to generate new UUID")
	}

	url, err := ss.serveDeviceDescriptionDocument()
	if err != nil {
		return errors.Wrap(err, "failed to serve device description document")
	}

	ss.ssdp, err = ssdp.Advertise(
		"upnp:rootdevice",                   // ST: Type
		"uuid:"+ss.uuid+"::upnp:rootdevice", // USN: ID
		url.String(),                        // LOCATION: location header
		runtime.GOOS+" UPnP/1.1 node/"+metadata.VersionAsString(), // SERVER: server header
		1800) // cache control, max-age. A duration for which the advertisement is valid
	if err != nil {
		return errors.Wrap(err, "failed to start SSDP advertiser")
	}

	if ss.ssdpPubliclyAccessible() {
		log.Warn().Msg("SSDP publicly accessible. Stopping it to prevent abusing service")
		err := ss.Stop()
		if err != nil {
			log.Error().Err(err).Msg("Failed to stop SSDP")
		}
	}

	for {
		select {
		case <-time.After(30 * time.Second):
			if err := ss.ssdp.Alive(); err != nil {
				log.Warn().Err(err).Msg("Failed to sent SSDP Alive message")
			}
		case <-ss.quit:
			return nil
		}
	}
}

func (ss *ssdpServer) Stop() error {
	var closeError error
	ss.once.Do(func() {
		close(ss.quit)
		if ss.ssdp != nil {
			if err := ss.ssdp.Bye(); err != nil {
				log.Error().Err(err).Msg("Failed to send SSDP bye message")
			}
			closeError = errors.Wrap(ss.ssdp.Close(), "failed to send SSDP bye message")
		}
	})

	return closeError
}

func (ss *ssdpServer) serveDeviceDescriptionDocument() (url.URL, error) {
	resolver := ip.NewResolver(ss.httpClient, "0.0.0.0", "", ip.IPFallbackAddresses)

	outIP, err := resolver.GetOutboundIP()
	if err != nil {
		return url.URL{}, err
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return url.URL{}, err
	}

	deviceDoc := ss.deviceDescription(outIP)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, deviceDoc)
	})

	go func() {
		go http.Serve(listener, mux)
		<-ss.quit
		listener.Close()
	}()

	port := listener.Addr().(*net.TCPAddr).Port
	return url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", outIP, port),
	}, nil
}

func (ss *ssdpServer) deviceDescription(ip string) string {
	var buf bytes.Buffer
	deviceDescription := template.Must(template.New("SSDPDeviceDescription").Parse(deviceDescriptionTemplate))
	_ = deviceDescription.Execute(&buf,
		struct{ URL, Version, UUID string }{
			fmt.Sprintf("http://%s:%d/", ip, ss.uiPort),
			metadata.VersionAsString(),
			ss.uuid,
		})

	return buf.String()
}

func (ss *ssdpServer) ssdpPubliclyAccessible() bool {
	// https://github.com/cloudflare/badupnp
	// https://blog.cloudflare.com/ssdp-100gbps/
	// > Furthermore, we prepared on online checking website.
	// > Click if you want to know if your public IP address has a vulnerable SSDP service:
	// > https://badupnp.benjojo.co.uk/

	req, err := requests.NewGetRequest("https://badupnp.benjojo.co.uk", "test", nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create HTTP request to test a vulnerable SSDP service")
		return false
	}

	var result struct {
		Result bool `json:"result"`
	}

	err = ss.httpClient.DoRequestAndParseResponse(req, &result)
	if err != nil {
		log.Error().Err(err).Msg("Failed to detect IP has a vulnerable SSDP service")
		return false
	}

	return result.Result
}

func generateUUID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}
