/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
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

package firewall

import (
	"net"
)

// IncomingTrafficFirewall defines provider side firewall, to control which traffic is enabled to pass and which not.
type IncomingTrafficFirewall interface {
	Setup() error
	Teardown()
	BlockIncomingTraffic(network net.IPNet) (IncomingRuleRemove, error)
	AllowURLAccess(rawURLs ...string) (IncomingRuleRemove, error)
	AllowIPAccess(ip net.IP) (IncomingRuleRemove, error)
}

// IncomingRuleRemove type defines function for removal of created rule.
type IncomingRuleRemove func() error
