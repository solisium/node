/*
 * Copyright (C) 2021 The "MysteriumNetwork/node" Authors.
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

package contract

import (
	"net/http"
	"strconv"

	"github.com/mysteriumnetwork/node/tequilapi/validation"
)

// EntertainmentEstimateRequest request to estimate entertainment amounts.
type EntertainmentEstimateRequest struct {
	Amount float64
}

// Bind fills and validates EntertainmentEstimateRequest from API request.
func (req *EntertainmentEstimateRequest) Bind(request *http.Request) *validation.FieldErrorMap {
	errs := validation.NewErrorMap()

	amt := request.URL.Query().Get("amount")
	if amt == "" {
		errs.ForField("amount").Required()
		return errs
	}

	amtf, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		errs.ForField("amount").Invalid("Failed to parse amount")
		return errs
	}

	req.Amount = amtf
	return errs
}

// EntertainmentEstimateResponse represents estimated entertainment.
// swagger:model EntertainmentEstimateResponse
type EntertainmentEstimateResponse struct {
	VideoMinutes    uint64  `json:"video_minutes"`
	MusicMinutes    uint64  `json:"music_minutes"`
	BrowsingMinutes uint64  `json:"browsing_minutes"`
	TrafficMB       uint64  `json:"traffic_mb"`
	PriceGiB        float64 `json:"price_gib"`
	PriceMin        float64 `json:"price_min"`
}
