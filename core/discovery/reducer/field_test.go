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

package reducer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Field(t *testing.T) {
	match := Field(fieldProviderID, func(value interface{}) bool {
		return value == provider1
	})

	assert.False(t, match(proposalEmpty))
	assert.True(t, match(proposalProvider1Streaming))
	assert.True(t, match(proposalProvider1Noop))
	assert.False(t, match(proposalProvider2Streaming))
}

func Test_Empty(t *testing.T) {
	match := Empty(fieldProviderID)

	assert.True(t, match(proposalEmpty))
	assert.False(t, match(proposalProvider1Streaming))
	assert.False(t, match(proposalProvider1Noop))
	assert.False(t, match(proposalProvider2Streaming))
}
