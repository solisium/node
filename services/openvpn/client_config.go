/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
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

package openvpn

import (
	"net"
	"strconv"

	"github.com/mysteriumnetwork/go-openvpn/openvpn/config"
	"github.com/mysteriumnetwork/node/core/connection"
)

// ClientConfig represents specific "openvpn as client" configuration
type ClientConfig struct {
	*config.GenericConfig
	VpnConfig *VPNConfig
}

// SetClientMode adds config arguments for openvpn behave as client
func (c *ClientConfig) SetClientMode(serverIP string, serverPort, localPort int) {
	c.SetFlag("client")
	c.SetParam("script-security", "2")
	c.SetFlag("auth-nocache")
	c.SetParam("remote", serverIP)
	c.SetPort(serverPort)
	c.SetParam("lport", strconv.Itoa(localPort))
	c.SetFlag("float")
	// more on this: https://www.v13.gr/blog/?p=386
	c.SetParam("remote-cert-ku", "84")
	c.SetFlag("auth-user-pass")
	c.SetFlag("management-query-passwords")
}

// SetProtocol specifies openvpn connection protocol type (tcp or udp)
func (c *ClientConfig) SetProtocol(protocol string) {
	if protocol == "tcp" {
		c.SetParam("proto", "tcp-client")
	} else if protocol == "udp" {
		c.SetFlag("explicit-exit-notify")
	}
}

func defaultClientConfig(runtimeDir string, scriptSearchPath string) *ClientConfig {
	clientConfig := ClientConfig{GenericConfig: config.NewConfig(runtimeDir, scriptSearchPath), VpnConfig: nil}

	clientConfig.SetDevice("tun")
	clientConfig.SetParam("cipher", "AES-256-GCM")
	clientConfig.SetParam("verb", "3")
	clientConfig.SetParam("tls-cipher", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384")
	clientConfig.SetKeepAlive(10, 60)
	clientConfig.SetPingTimerRemote()
	clientConfig.SetPersistKey()

	clientConfig.SetParam("auth", "none")

	clientConfig.SetParam("reneg-sec", "0")
	clientConfig.SetParam("resolv-retry", "infinite")
	clientConfig.SetParam("redirect-gateway", "def1", "bypass-dhcp")

	return &clientConfig
}

// NewClientConfigFromSession creates client configuration structure for given VPNConfig, configuration dir to store serialized file args, and
// configuration filename to store other args
// TODO this will become the part of openvpn service consumer separate package
func NewClientConfigFromSession(vpnConfig VPNConfig, scriptDir string, runtimeDir string, options connection.ConnectOptions) (*ClientConfig, error) {
	// TODO Rename `vpnConfig` to `sessionConfig`
	err := NewDefaultValidator().IsValid(vpnConfig)
	if err != nil {
		return nil, err
	}

	vpnConfig, err = FormatTLSPresharedKey(vpnConfig)
	if err != nil {
		return nil, err
	}

	clientFileConfig := newClientConfig(runtimeDir, scriptDir)
	dnsIPs, err := options.Params.DNS.ResolveIPs(vpnConfig.DNSIPs)
	if err != nil {
		return nil, err
	}
	for _, ip := range dnsIPs {
		clientFileConfig.SetParam("dhcp-option", "DNS", ip)
	}

	var remotePort, localPort int
	if options.ProviderNATConn != nil && vpnConfig.RemoteIP != "127.0.0.1" {
		options.ProviderNATConn.Close()
		remotePort = options.ProviderNATConn.RemoteAddr().(*net.UDPAddr).Port
		localPort = options.ProviderNATConn.LocalAddr().(*net.UDPAddr).Port
	} else {
		remotePort = vpnConfig.RemotePort
		localPort = vpnConfig.LocalPort
	}

	clientFileConfig.VpnConfig = &vpnConfig
	clientFileConfig.SetReconnectRetry(2)
	clientFileConfig.SetClientMode(vpnConfig.RemoteIP, remotePort, localPort)
	clientFileConfig.SetProtocol(vpnConfig.RemoteProtocol)
	clientFileConfig.SetTLSCACertificate(vpnConfig.CACertificate)
	clientFileConfig.SetTLSCrypt(vpnConfig.TLSPresharedKey)

	return clientFileConfig, nil
}
