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

package license

import (
	"fmt"

	"github.com/mysteriumnetwork/node/metadata"
	"github.com/urfave/cli/v2"
)

var (
	// FlagShowWarranty allows to print license warranty.
	FlagShowWarranty = cli.BoolFlag{
		Name:  "warranty",
		Usage: "Show details of license warranty",
	}
	// FlagShowConditions allows to print license conditions.
	FlagShowConditions = cli.BoolFlag{
		Name:  "conditions",
		Usage: "Show details of license conditions",
	}
)

// NewCommand function creates license command.
func NewCommand(licenseCopyright string) *cli.Command {
	return &cli.Command{
		Name:      "license",
		Usage:     "Show license",
		ArgsUsage: " ",
		Flags:     []cli.Flag{&FlagShowWarranty, &FlagShowConditions},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool(FlagShowWarranty.Name) {
				_, err := fmt.Fprintln(ctx.App.Writer, metadata.LicenseWarranty)

				return err
			}

			if ctx.Bool(FlagShowConditions.Name) {
				_, err := fmt.Fprintln(ctx.App.Writer, metadata.LicenseConditions)

				return err
			}

			_, err := fmt.Fprintln(ctx.App.Writer, licenseCopyright)

			return err
		},
	}
}
