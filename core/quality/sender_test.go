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

package quality

import (
	"errors"
	"runtime"
	"testing"

	"github.com/mysteriumnetwork/node/identity"
	"github.com/stretchr/testify/assert"
)

type mockEventsTransport struct {
	sentEvent    Event
	mockResponse error
}

func buildMockEventsTransport(mockResponse error) *mockEventsTransport {
	return &mockEventsTransport{mockResponse: mockResponse}
}

func (transport *mockEventsTransport) SendEvent(event Event) error {
	transport.sentEvent = event
	return transport.mockResponse
}

func TestSender_SendStartupEvent_SendsToTransport(t *testing.T) {
	mockTransport := buildMockEventsTransport(nil)
	sender := &Sender{Transport: mockTransport, AppVersion: "test version"}

	sender.sendUnlockEvent(identity.AppEventIdentityUnlock{ID: identity.FromAddress("0x1234567890abcdef")})

	sentEvent := mockTransport.sentEvent
	assert.Equal(t, "unlock", sentEvent.EventName)
	assert.Equal(t, appInfo{Name: "myst", Version: "test version", OS: runtime.GOOS, Arch: runtime.GOARCH}, sentEvent.Application)
	assert.NotZero(t, sentEvent.CreatedAt)
}

func TestSender_SendNATMappingSuccessEvent_SendsToTransport(t *testing.T) {
	mockTransport := buildMockEventsTransport(nil)
	sender := &Sender{Transport: mockTransport, AppVersion: "test version"}

	sender.SendNATMappingSuccessEvent("id", "port_mapping", nil)

	sentEvent := mockTransport.sentEvent
	assert.Equal(t, "nat_mapping", sentEvent.EventName)
	assert.Equal(t, appInfo{Name: "myst", Version: "test version", OS: runtime.GOOS, Arch: runtime.GOARCH}, sentEvent.Application)
	assert.NotZero(t, sentEvent.CreatedAt)
	assert.Equal(t, natMappingContext{ID: "id", Successful: true, Stage: "port_mapping"}, sentEvent.Context)
}

func TestSender_SendNATMappingFailEvent_SendsToTransport(t *testing.T) {
	mockTransport := buildMockEventsTransport(nil)

	mockGateways := []map[string]string{
		{"test": "test"},
	}
	mockError := errors.New("mock nat mapping error")

	sender := &Sender{Transport: mockTransport, AppVersion: "test version"}
	sender.SendNATMappingFailEvent("id", "hole_punching", mockGateways, mockError)

	sentEvent := mockTransport.sentEvent
	assert.Equal(t, "nat_mapping", sentEvent.EventName)
	assert.Equal(t, appInfo{Name: "myst", Version: "test version", OS: runtime.GOOS, Arch: runtime.GOARCH}, sentEvent.Application)
	assert.NotZero(t, sentEvent.CreatedAt)
	c := sentEvent.Context.(natMappingContext)
	assert.False(t, c.Successful)
	assert.Equal(t, "mock nat mapping error", *c.ErrorMessage)
	assert.Equal(t, "hole_punching", c.Stage)
	assert.Equal(t, "id", c.ID)
	assert.Equal(t, mockGateways, c.Gateways)
}
