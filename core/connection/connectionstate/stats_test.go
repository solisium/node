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

package connectionstate

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	exampleStats = Statistics{
		BytesReceived: 1,
		BytesSent:     2,
	}
)

func TestStatistics_Diff(t *testing.T) {
	tests := []struct {
		name string
		old  Statistics
		new  Statistics
		want Statistics
	}{
		{
			name: "calculates statistics correctly if they are continuous",
			old:  Statistics{},
			new:  exampleStats,
			want: exampleStats,
		},
		{
			name: "calculates statistics correctly if they are not continuous",
			old: Statistics{
				BytesReceived: 5,
				BytesSent:     6,
			},
			new:  exampleStats,
			want: exampleStats,
		},
		{
			name: "returns zeros on no change",
			old:  exampleStats,
			new:  exampleStats,
			want: Statistics{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.old.Diff(tt.new)
			assert.True(t, reflect.DeepEqual(result, tt.want))
		})
	}
}

func TestStatistics_Plus(t *testing.T) {
	tests := []struct {
		name  string
		stats Statistics
		diff  Statistics
		want  Statistics
	}{
		{
			name:  "adds up stats correctly",
			diff:  exampleStats,
			stats: exampleStats,
			want: Statistics{
				BytesReceived: 2,
				BytesSent:     4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.stats.Plus(tt.diff)
			assert.EqualValues(t, tt.want.BytesReceived, result.BytesReceived)
			assert.EqualValues(t, tt.want.BytesSent, result.BytesSent)
		})
	}
}
