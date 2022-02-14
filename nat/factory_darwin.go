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

package nat

import "os/exec"

// NewService returns fake nat service since there are no iptables on darwin
func NewService() NATService {
	return &servicePFCtl{
		ipForward: serviceIPForward{
			CommandFactory: func(name string, arg ...string) Command {
				return exec.Command(name, arg...)
			},
			CommandEnable:  []string{"/usr/sbin/sysctl", "-w", "net.inet.ip.forwarding=1"},
			CommandDisable: []string{"/usr/sbin/sysctl", "-w", "net.inet.ip.forwarding=0"},
			CommandRead:    []string{"/usr/sbin/sysctl", "-n", "net.inet.ip.forwarding"},
		},
		rules: []string{},
	}
}
