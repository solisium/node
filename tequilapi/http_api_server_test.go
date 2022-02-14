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

package tequilapi

import (
	"net"
	"strings"
	"testing"

	"github.com/mysteriumnetwork/node/core/node"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestLocalAPIServerPortIsAsExpected(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:31337")
	assert.Nil(t, err)

	server, err := NewServer(listener, *node.GetOptions(), []func(e *gin.Engine) error{})
	assert.NoError(t, err)

	server.StartServing()

	address, err := server.Address()
	assert.NoError(t, err)

	port := strings.Split(address, ":")[1]
	assert.Equal(t, "31337", port)

	server.Stop()
	server.Wait()
}

func TestStopBeforeStartingListeningDoesNotCausePanic(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:31337")
	assert.Nil(t, err)
	server, err := NewServer(listener, *node.GetOptions(), []func(e *gin.Engine) error{})
	assert.NoError(t, err)
	server.Stop()
}
