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

	"github.com/stretchr/testify/assert"
)

type customPayload struct {
	Field int
}

var _ Codec = &codecJSON{}

func TestCodecJSONPack(t *testing.T) {
	table := []struct {
		payload      interface{}
		expectedData string
	}{
		{`hello`, `"hello"`},
		{`hello "name"`, `"hello \"name\""`},
		{true, `true`},
		{nil, `null`},
		{10, `10`},
		{10.20, `10.2`},
		{&customPayload{123}, `{"Field":123}`},
		{&customPayload{}, `{"Field":0}`},
	}

	codec := codecJSON{}
	for _, tt := range table {
		data, err := codec.Pack(tt.payload)

		assert.NoError(t, err)
		assert.Exactly(t, []byte(tt.expectedData), data)
	}
}

const (
	typeString = iota
	typeBool
	typeInt
	typeFloat
	typeCustom
)

func TestCodecJSONUnpack(t *testing.T) {
	table := []struct {
		data            string
		payloadType     int
		expectedPayload interface{}
	}{
		{`"hello"`, typeString, `hello`},
		{`"hello \"name\""`, typeString, `hello "name"`},
		{`true`, typeBool, true},
		{`null`, typeBool, nil},
		{`10`, typeInt, 10},
		{`10.2`, typeFloat, 10.20},
		{`{"Field":123}`, typeCustom, &customPayload{123}},
		{`{"Field":0}`, typeCustom, &customPayload{}},
	}

	codec := codecJSON{}
	for _, tt := range table {
		switch tt.payloadType {
		case typeString:
			var payload string
			err := codec.Unpack([]byte(tt.data), &payload)

			assert.NoError(t, err)
			assert.Exactly(t, tt.expectedPayload, payload)

		case typeBool:
			var payload interface{}
			err := codec.Unpack([]byte(tt.data), &payload)

			assert.NoError(t, err)
			assert.Exactly(t, tt.expectedPayload, payload)

		case typeInt:
			var payload int
			err := codec.Unpack([]byte(tt.data), &payload)

			assert.NoError(t, err)
			assert.Exactly(t, tt.expectedPayload, payload)

		case typeFloat:
			var payload float64
			err := codec.Unpack([]byte(tt.data), &payload)

			assert.NoError(t, err)
			assert.Exactly(t, tt.expectedPayload, payload)

		case typeCustom:
			var payload *customPayload
			err := codec.Unpack([]byte(tt.data), &payload)

			assert.NoError(t, err)
			assert.Exactly(t, tt.expectedPayload, payload)

		default:
			t.Fatalf("Unknown type: %d", tt.payloadType)
		}
	}
}
