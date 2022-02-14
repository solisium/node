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

package dns

import (
	"net"
	"strconv"

	"github.com/miekg/dns"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Proxy defines DNS server with all handler attached to it.
type Proxy struct {
	server *dns.Server
}

// NewProxy returns new instance of API server.
func NewProxy(lhost string, lport int, handler dns.Handler) *Proxy {
	return &Proxy{
		server: &dns.Server{
			Addr:         net.JoinHostPort(lhost, strconv.Itoa(lport)),
			Net:          "udp",
			ReadTimeout:  dnsTimeout,
			WriteTimeout: dnsTimeout,
			ReusePort:    true,
			Handler:      handler,
		},
	}
}

// Run starts DNS proxy server and waits for the startup to complete.
func (p *Proxy) Run() (err error) {
	dnsProxyCh := make(chan error)
	p.server.NotifyStartedFunc = func() { dnsProxyCh <- nil }
	go func() {
		log.Info().Msg("Starting DNS proxy on: " + p.server.Addr)
		if err := p.server.ListenAndServe(); err != nil {
			dnsProxyCh <- errors.Wrap(err, "failed to start DNS proxy")
		}
	}()

	return <-dnsProxyCh
}

// Stop shutdowns DNS proxy server.
func (p *Proxy) Stop() error {
	return p.server.Shutdown()
}
