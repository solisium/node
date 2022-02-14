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

package communication

import (
	"testing"

	"github.com/mysteriumnetwork/node/testkit/assertkit"
	"github.com/stretchr/testify/assert"
)

var _ Codec = &codecBytes{}

func TestCodecBytesPack(t *testing.T) {
	table := []struct {
		payload          interface{}
		expectedData     []byte
		expectedErrorMsg string
	}{
		{`hello`, []byte(`hello`), ""},
		{`hello "name"`, []byte(`hello "name"`), ""},
		{nil, []byte{}, ""},
		{true, []byte{}, "can't pack payload: true"},
	}

	codec := codecBytes{}
	for _, tt := range table {
		data, err := codec.Pack(tt.payload)

		assert.Exactly(t, tt.expectedData, data)
		assertkit.EqualOptionalError(t, err, tt.expectedErrorMsg)
	}
}

func TestCodecBytesUnpackToString(t *testing.T) {
	codec := codecBytes{}

	var payload string
	err := codec.Unpack([]byte(`hello`), &payload)

	assert.Error(t, err)
	assert.EqualError(t, err, "can't unpack to payload: *string")
}

func TestCodecBytesUnpackToByte(t *testing.T) {
	codec := codecBytes{}

	var payload []byte
	err := codec.Unpack([]byte(`hello`), payload)

	assert.Error(t, err)
	assert.EqualError(t, err, "can't unpack to payload: []uint8")
}

func TestCodecBytesUnpack(t *testing.T) {
	table := []struct {
		data            []byte
		expectedPayload []byte
	}{
		{[]byte(`hello`), []byte(`hello`)},
		{[]byte(`hello "name"`), []byte(`hello "name"`)},
		{[]byte{}, []byte{}},
	}

	codec := codecBytes{}
	for _, tt := range table {
		var payload []byte
		err := codec.Unpack(tt.data, &payload)

		assert.NoError(t, err)
		assert.Exactly(t, tt.expectedPayload, payload)
	}
}
