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

package main

import (
	"os"
	"sync"

	"github.com/mysteriumnetwork/node/cmd/commands/account"
	command_cli "github.com/mysteriumnetwork/node/cmd/commands/cli"
	command_cfg "github.com/mysteriumnetwork/node/cmd/commands/config"
	"github.com/mysteriumnetwork/node/cmd/commands/connection"
	"github.com/mysteriumnetwork/node/cmd/commands/daemon"
	"github.com/mysteriumnetwork/node/cmd/commands/license"
	"github.com/mysteriumnetwork/node/cmd/commands/reset"
	"github.com/mysteriumnetwork/node/cmd/commands/service"
	"github.com/mysteriumnetwork/node/cmd/commands/version"
	"github.com/mysteriumnetwork/node/config"
	"github.com/mysteriumnetwork/node/logconfig"
	"github.com/mysteriumnetwork/node/metadata"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	licenseCopyright = metadata.LicenseCopyright(
		"run command 'license --warranty'",
		"run command 'license --conditions'",
	)
	versionSummary    = metadata.VersionAsSummary(licenseCopyright)
	daemonCommand     = daemon.NewCommand()
	versionCommand    = version.NewCommand(versionSummary)
	licenseCommand    = license.NewCommand(licenseCopyright)
	serviceCommand    = service.NewCommand(licenseCommand.Name)
	cliCommand        = command_cli.NewCommand()
	resetCommand      = reset.NewCommand()
	accountCommand    = account.NewCommand()
	connectionCommand = connection.NewCommand()
	configCommand     = command_cfg.NewCommand()
)

func main() {
	logconfig.Bootstrap()
	app, err := NewCommand()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create command: ")
		os.Exit(1)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to execute command: ")
		os.Exit(1)
	}
}

// NewCommand function creates application master command
func NewCommand() (*cli.App, error) {
	cli.VersionPrinter = func(ctx *cli.Context) {
		versionCommand.Run(ctx)
	}

	app, err := newApp()
	if err != nil {
		return nil, err
	}

	app.Usage = "VPN server and client for Mysterium Network https://mysterium.network/"
	app.Authors = []*cli.Author{
		{Name: `The "MysteriumNetwork/node" Authors`, Email: "mysterium-dev@mysterium.network"},
	}
	app.Version = metadata.VersionAsString()
	app.Copyright = licenseCopyright
	app.Before = configureLogging()

	app.Commands = []*cli.Command{
		versionCommand,
		licenseCommand,
		serviceCommand,
		daemonCommand,
		cliCommand,
		resetCommand,
		accountCommand,
		connectionCommand,
		configCommand,
	}

	return app, nil
}

func newApp() (*cli.App, error) {
	app := cli.NewApp()
	return app, config.RegisterFlagsNode(&app.Flags)
}

// uiCommands is a map which consists of all
// commands are used directly by a user.
var uiCommands = map[string]struct{}{
	command_cli.CommandName: {},
	account.CommandName:     {},
	connection.CommandName:  {},
	command_cfg.CommandName: {},
	reset.CommandName:       {},
}

// configureLogging returns a func which configures global
// logging settings depending on the command used.
// It only runs once.
func configureLogging() cli.BeforeFunc {
	var once sync.Once
	return func(ctx *cli.Context) error {
		once.Do(func() {
			// Keep default settings if verbose logging is enabled.
			if ctx.Bool(config.FlagVerbose.Name) {
				return
			}

			cmd := ctx.Args().First()
			if _, ok := uiCommands[cmd]; !ok {
				// If the command is not meant for user
				// interaction, skip.
				return
			}
			logconfig.SetLogLevel(zerolog.PanicLevel)
		})
		return nil
	}
}
