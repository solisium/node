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

package daemon

import (
	"github.com/mysteriumnetwork/node/cmd"
	"github.com/mysteriumnetwork/node/config"
	"github.com/mysteriumnetwork/node/config/urfavecli/clicontext"
	"github.com/mysteriumnetwork/node/core/node"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// NewCommand function creates run command
func NewCommand() *cli.Command {
	var di cmd.Dependencies

	command := &cli.Command{
		Name:      "daemon",
		Usage:     "Starts Mysterium Tequilapi service",
		ArgsUsage: " ",
		Before:    clicontext.LoadUserConfigQuietly,
		Action: func(ctx *cli.Context) error {
			quit := make(chan error, 2)
			config.ParseFlagsServiceStart(ctx)
			config.ParseFlagsServiceOpenvpn(ctx)
			config.ParseFlagsServiceWireguard(ctx)
			config.ParseFlagsServiceNoop(ctx)
			config.ParseFlagsNode(ctx)

			nodeOptions := node.GetOptions()
			if err := di.Bootstrap(*nodeOptions); err != nil {
				return err
			}
			go func() { quit <- di.Node.Wait() }()

			cmd.RegisterSignalCallback(func() { quit <- nil })

			return describeQuit(<-quit)
		},
		After: func(ctx *cli.Context) error {
			return di.Shutdown()
		},
	}

	return command
}

func describeQuit(err error) error {
	if err == nil {
		log.Info().Msg("Stopping application")
	} else {
		log.Error().Err(err).Msgf("Terminating application due to error")
	}
	return err
}
