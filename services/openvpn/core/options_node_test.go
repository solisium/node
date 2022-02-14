/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
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

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorIsReturnedOnBadBinaryPath(t *testing.T) {
	options := NodeOptions{"non-existent-binary"}
	assert.Error(t, options.Check())
}

func TestErrorIsReturnedOnExitCodeZero(t *testing.T) {
	options := NodeOptions{"testdata/exit-with-zero.sh"}
	assert.Error(t, options.Check())
}

func TestNoErrorIsReturnedOnExitCodeOne(t *testing.T) {
	options := NodeOptions{"testdata/openvpn-version.sh"}
	assert.NoError(t, options.Check())
}

func TestNoErrorIsReturnedOnOpenvpnWithCustomBuild(t *testing.T) {
	options := NodeOptions{"testdata/openvpn-version-custom-tag.sh"}
	assert.NoError(t, options.Check())
}
