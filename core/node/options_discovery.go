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

package node

import "time"

// DiscoveryType identifies proposal discovery provider.
type DiscoveryType string

const (
	// DiscoveryTypeAPI defines type which discovers proposals through Mysterium API.
	DiscoveryTypeAPI = DiscoveryType("api")
	// DiscoveryTypeBroker defines type which discovers proposals through Broker (Mysterium Communication).
	DiscoveryTypeBroker = DiscoveryType("broker")
	// DiscoveryTypeDHT defines type which discovers proposals through DHT (Distributed Hash Table).
	DiscoveryTypeDHT = DiscoveryType("dht")
)

// OptionsDiscovery describes possible parameters of discovery configuration.
type OptionsDiscovery struct {
	Types         []DiscoveryType
	Address       string
	PingInterval  time.Duration
	FetchEnabled  bool
	FetchInterval time.Duration
	DHT           OptionsDHT
}

// OptionsDHT describes possible parameters of DHT configuration.
type OptionsDHT struct {
	Address        string
	Port           int
	Protocol       string
	BootstrapPeers []string
}
