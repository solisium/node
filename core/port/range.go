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

package port

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Range represents a networking port range
type Range struct {
	Start, End int
}

// ParseRange parses port range expression, e.g. "1000:1200" into Range struct
func ParseRange(rangeExpr string) (Range, error) {
	bounds := strings.Split(rangeExpr, ":")
	if len(bounds) != 2 {
		return Range{}, errors.New("invalid port range expression: " + rangeExpr)
	}

	start, err := strconv.Atoi(bounds[0])
	if err != nil {
		return Range{}, errors.New("invalid start port number: " + bounds[0])
	}

	end, err := strconv.Atoi(bounds[1])
	if err != nil {
		return Range{}, errors.New("invalid end port number: " + bounds[0])
	}

	if start > end {
		return Range{}, errors.New("start port cannot be greater than end port: " + rangeExpr)
	}

	return Range{start, end}, nil
}

// Capacity returns range capacity
func (r *Range) Capacity() int {
	return r.End - r.Start
}

// String returns port range expression, e.g. "1000:1200"
func (r *Range) String() string {
	return fmt.Sprintf("%d:%d", r.Start, r.End)
}
