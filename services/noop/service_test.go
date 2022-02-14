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

package noop

import (
	"testing"
	"time"

	"github.com/mysteriumnetwork/node/core/service"
	"github.com/stretchr/testify/assert"
)

var _ service.Service = NewManager()

func Test_Manager_ProvideConfig(t *testing.T) {
	manager := NewManager()
	sessionConfig, err := manager.ProvideConfig("", nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, sessionConfig.SessionServiceConfig)
	assert.Nil(t, sessionConfig.SessionDestroyCallback)
}

func Test_Manager_Serve_Stop(t *testing.T) {
	manager := NewManager()
	go func() {
		err := manager.Serve(&service.Instance{})
		assert.NoError(t, err)
	}()

	time.Sleep(time.Millisecond * 10)
	err := manager.Stop()
	assert.NoError(t, err)
}
