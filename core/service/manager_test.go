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

package service

import (
	"errors"
	"testing"
	"time"

	"github.com/mysteriumnetwork/node/core/location/locationstate"
	"github.com/mysteriumnetwork/node/core/policy"
	"github.com/mysteriumnetwork/node/core/service/servicestate"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/market"
	"github.com/mysteriumnetwork/node/mocks"
	"github.com/mysteriumnetwork/node/p2p"
	"github.com/mysteriumnetwork/node/requests"
	"github.com/mysteriumnetwork/node/utils/netutil"
	"github.com/stretchr/testify/assert"
)

var (
	serviceType      = "the-very-awesome-test-service-type"
	mockPolicyOracle = policy.NewOracle(requests.NewHTTPClient("0.0.0.0", requests.DefaultTimeout), "http://policy.localhost/", 1*time.Minute)
)

func init() {
	netutil.LogNetworkStats = func() {}
}

func TestManager_StartRemovesServiceFromPoolIfServiceCrashes(t *testing.T) {
	registry := NewRegistry()
	mockCopy := *serviceMock
	mockCopy.onStartReturnError = errors.New("some error")
	registry.Register(serviceType, func(options Options) (Service, error) {
		return &mockCopy, nil
	})

	discovery := mockDiscovery{}
	discoveryFactory := MockDiscoveryFactoryFunc(&discovery)
	manager := NewManager(
		registry,
		discoveryFactory,
		mocks.NewEventBus(),
		mockPolicyOracle,
		&mockP2PListener{}, nil, nil, mockLocationResolver{},
	)
	_, err := manager.Start(identity.FromAddress(proposalMock.ProviderID), serviceType, nil, struct{}{})
	assert.Nil(t, err)

	discovery.Wait()
	assert.Len(t, manager.servicePool.List(), 0)
}

func TestManager_StartDoesNotCrashIfStoppedByUser(t *testing.T) {
	registry := NewRegistry()
	mockCopy := *serviceMock
	mockCopy.mockProcess = make(chan struct{})
	registry.Register(serviceType, func(options Options) (Service, error) {
		return &mockCopy, nil
	})

	discovery := mockDiscovery{}
	discoveryFactory := MockDiscoveryFactoryFunc(&discovery)
	manager := NewManager(
		registry,
		discoveryFactory,
		mocks.NewEventBus(),
		mockPolicyOracle,
		&mockP2PListener{}, nil, nil,
		mockLocationResolver{},
	)
	id, err := manager.Start(identity.FromAddress(proposalMock.ProviderID), serviceType, nil, struct{}{})
	assert.Nil(t, err)
	err = manager.Stop(id)
	assert.Nil(t, err)
	discovery.Wait()
	assert.Len(t, manager.servicePool.List(), 0)
}

func TestManager_StopSendsEvent_SucceedsAndPublishesEvent(t *testing.T) {
	registry := NewRegistry()
	mockCopy := *serviceMock
	mockCopy.mockProcess = make(chan struct{})
	registry.Register(serviceType, func(options Options) (Service, error) {
		return &mockCopy, nil
	})

	discovery := mockDiscovery{}
	discoveryFactory := MockDiscoveryFactoryFunc(&discovery)
	eventBus := &mockPublisher{}
	manager := NewManager(
		registry,
		discoveryFactory,
		eventBus,
		mockPolicyOracle,
		&mockP2PListener{}, nil, nil,
		mockLocationResolver{},
	)

	id, err := manager.Start(identity.FromAddress(proposalMock.ProviderID), serviceType, nil, struct{}{})
	assert.NoError(t, err)

	services := manager.servicePool.List()

	var serviceID ID
	for k := range services {
		serviceID = services[k].ID
	}

	err = manager.Stop(id)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 30)

	eventBus.lock.Lock()
	defer eventBus.lock.Unlock()

	assert.Equal(t, servicestate.AppTopicServiceStatus, eventBus.publishedTopic)

	var matchFound bool
	expectedPayload := servicestate.AppEventServiceStatus{ID: string(serviceID), ProviderID: "", Type: "", Status: "NotRunning"}
	for i := range eventBus.publishedData {
		e, ok := eventBus.publishedData[i].(servicestate.AppEventServiceStatus)
		if !ok {
			continue
		}
		if e.Status == expectedPayload.Status && e.ID == expectedPayload.ID {
			matchFound = true
			break
		}
	}
	assert.True(t, matchFound)
}

type mockP2PListener struct {
}

func (m mockP2PListener) GetContact() market.Contact {
	return market.Contact{}
}

func (m mockP2PListener) Listen(_ identity.Identity, serviceType string, channelHandler func(ch p2p.Channel)) (func(), error) {
	return func() {}, nil
}

type mockLocationResolver struct {
}

func (m mockLocationResolver) DetectLocation() (locationstate.Location, error) {
	return locationstate.Location{}, nil
}
