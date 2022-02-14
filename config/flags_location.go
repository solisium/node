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

package config

import (
	"fmt"

	"github.com/mysteriumnetwork/node/metadata"
	"github.com/urfave/cli/v2"
)

var (
	// FlagIPDetectorURL URL of IP detection service.
	FlagIPDetectorURL = cli.StringFlag{
		Name:  "ip-detector",
		Usage: "Address (URL form) of IP detection service",
		Value: metadata.DefaultNetwork.LocationAddress,
	}
	// FlagLocationType location detector type.
	FlagLocationType = cli.StringFlag{
		Name:  "location.type",
		Usage: "Location autodetect adapter. Options: { oracle, builtin, mmdb, manual }",
		Value: "oracle",
	}
	// FlagLocationAddress URL of location detector.
	FlagLocationAddress = cli.StringFlag{
		Name: "location.address",
		Usage: fmt.Sprintf(
			"Address of specific location adapter given in '--%s'",
			FlagLocationType.Name,
		),
		Value: metadata.DefaultNetwork.LocationAddress,
	}
	// FlagLocationCountry service location country.
	FlagLocationCountry = cli.StringFlag{
		Name:  "location.country",
		Usage: "Service location country",
	}
	// FlagLocationCity service location city.
	FlagLocationCity = cli.StringFlag{
		Name:  "location.city",
		Usage: "Service location city",
	}
	// FlagLocationIPType service location node type.
	FlagLocationIPType = cli.StringFlag{
		Name:  "location.ip-type",
		Usage: "Service location IP type (residential, datacenter, etc.)",
	}
)

// RegisterFlagsLocation function registers location flags to flag list.
func RegisterFlagsLocation(flags *[]cli.Flag) {
	*flags = append(*flags,
		&FlagIPDetectorURL,
		&FlagLocationType,
		&FlagLocationAddress,
		&FlagLocationCountry,
		&FlagLocationCity,
		&FlagLocationIPType,
	)
}

// ParseFlagsLocation function fills in location options from CLI context.
func ParseFlagsLocation(ctx *cli.Context) {
	Current.ParseStringFlag(ctx, FlagIPDetectorURL)
	Current.ParseStringFlag(ctx, FlagLocationType)
	Current.ParseStringFlag(ctx, FlagLocationAddress)
	Current.ParseStringFlag(ctx, FlagLocationCountry)
	Current.ParseStringFlag(ctx, FlagLocationCity)
	Current.ParseStringFlag(ctx, FlagLocationIPType)
}
