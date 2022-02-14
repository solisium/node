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

package firewall

import (
	"fmt"

	"github.com/mysteriumnetwork/node/utils/cmdutil"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// AddInboundRule adds new inbound rule to the platform specific firewall.
func AddInboundRule(proto string, port int) error {
	name := fmt.Sprintf("myst-%d:%s", port, proto)
	cmd := fmt.Sprintf(`netsh advfirewall firewall add rule name="%s" dir=in action=allow protocol=%s localport=%d`, name, proto, port)

	if inboundRuleExists(name) {
		return nil
	}

	_, err := cmdutil.PowerShell(cmd)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to add firewall rule")
		return err
	}

	return nil
}

// RemoveInboundRule removes inbound rule from the platform specific firewall.
func RemoveInboundRule(proto string, port int) error {
	name := fmt.Sprintf("myst-%d:%s", port, proto)
	cmd := fmt.Sprintf(`netsh advfirewall firewall delete rule name="%s" dir=in`, name)

	if !inboundRuleExists(name) {
		return errors.New("firewall rule not found")
	}

	_, err := cmdutil.PowerShell(cmd)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to remove firewall rule")
		return err
	}

	return nil
}

func inboundRuleExists(name string) bool {
	cmd := fmt.Sprintf(`netsh advfirewall firewall show rule name="%s" dir=in`, name)

	if _, err := cmdutil.PowerShell(cmd); err != nil {
		log.Warn().Err(err).Msg("Failed to get firewall rule")
		return false
	}

	return true
}
