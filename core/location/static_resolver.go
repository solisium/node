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

package location

import (
	"github.com/mysteriumnetwork/node/core/ip"
	"github.com/mysteriumnetwork/node/core/location/locationstate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// StaticResolver struct represents country by ip ExternalDBResolver which always returns specified country
type StaticResolver struct {
	country string
	city    string
	ipType  string
	err     error

	ipResolver ip.Resolver
}

// NewStaticResolver creates new StaticResolver with specified country
func NewStaticResolver(country, city, ipType string, ipResolver ip.Resolver) *StaticResolver {
	return &StaticResolver{
		country: country,
		city:    city,
		ipType:  ipType,
		err:     nil,

		ipResolver: ipResolver,
	}
}

// NewFailingResolver returns StaticResolver with entered error
func NewFailingResolver(err error) *StaticResolver {
	return &StaticResolver{
		err:        err,
		ipResolver: ip.NewResolverMock(""),
	}
}

// DetectLocation detects current IP-address provides location information for the IP.
func (d *StaticResolver) DetectLocation() (locationstate.Location, error) {
	log.Debug().Msg("Detecting with static resolver")
	pubIP, err := d.ipResolver.GetPublicIP()
	if err != nil {
		return locationstate.Location{}, errors.Wrap(err, "failed to get public IP")
	}
	return locationstate.Location{
		Country: d.country,
		City:    d.city,
		IPType:  d.ipType,
		IP:      pubIP,
	}, d.err
}
