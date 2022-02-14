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

package brokerdiscovery

import (
	"github.com/mysteriumnetwork/node/communication"
	"github.com/mysteriumnetwork/node/communication/nats"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/market"
)

// unregisterMessage structure represents message that the Provider sends about de-announced Proposal
type unregisterMessage struct {
	Proposal market.ServiceProposal `json:"proposal"`
}

const unregisterEndpoint = communication.MessageEndpoint("proposal-unregister.v3")

// unregisterProducer
type unregisterProducer struct {
	message *unregisterMessage
	signer  identity.Signer
}

// GetMessageEndpoint returns endpoint where to send messages
func (p *unregisterProducer) GetMessageEndpoint() (communication.MessageEndpoint, error) {
	subj, err := nats.SignedSubject(p.signer, string(unregisterEndpoint))
	return communication.MessageEndpoint(subj), err
}

// Produce creates message which will be serialized to endpoint
func (p *unregisterProducer) Produce() (requestPtr interface{}) {
	return p.message
}

// registerConsumer
type unregisterConsumer struct {
	Callback func(unregisterMessage) error
}

// GetMessageEndpoint returns endpoint where to receive messages
func (c *unregisterConsumer) GetMessageEndpoint() (communication.MessageEndpoint, error) {
	return unregisterEndpoint, nil
}

// NewMessage creates struct where message from endpoint will be serialized
func (c *unregisterConsumer) NewMessage() (messagePtr interface{}) {
	return &unregisterMessage{}
}

// Consume handles messages from endpoint
func (c *unregisterConsumer) Consume(messagePtr interface{}) error {
	return c.Callback(*messagePtr.(*unregisterMessage))
}
