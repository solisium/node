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

package nats

import (
	"testing"
	"time"

	"github.com/mysteriumnetwork/node/communication"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

type bytesRequestProducer struct {
	Request []byte
}

func (producer *bytesRequestProducer) GetRequestEndpoint() (communication.RequestEndpoint, error) {
	return communication.RequestEndpoint("bytes-request"), nil
}

func (producer *bytesRequestProducer) NewResponse() (responsePtr interface{}) {
	var response []byte
	return &response
}

func (producer *bytesRequestProducer) Produce() (requestPtr interface{}) {
	return producer.Request
}

func TestBytesRequest(t *testing.T) {
	connection := StartConnectionMock()
	connection.MockResponse("bytes-request", []byte("RESPONSE"))
	defer connection.Close()

	sender := &senderNATS{
		connection:     connection,
		codec:          communication.NewCodecBytes(),
		timeoutRequest: 100 * time.Millisecond,
	}

	response, err := sender.Request(&bytesRequestProducer{
		[]byte("REQUEST"),
	})
	assert.NoError(t, err)
	assert.Equal(t, []byte("REQUEST"), connection.GetLastRequest())
	assert.Equal(t, []byte("RESPONSE"), *response.(*[]byte))
}

type bytesRequestConsumer struct {
	requestReceived interface{}
}

func (consumer *bytesRequestConsumer) GetRequestEndpoint() (communication.RequestEndpoint, error) {
	return communication.RequestEndpoint("bytes-response"), nil
}

func (consumer *bytesRequestConsumer) NewRequest() (requestPtr interface{}) {
	var request []byte
	return &request
}

func (consumer *bytesRequestConsumer) Consume(requestPtr interface{}) (responsePtr interface{}, err error) {
	consumer.requestReceived = requestPtr
	return []byte("RESPONSE"), nil
}

func TestBytesRespond(t *testing.T) {
	connection := StartConnectionMock()
	defer connection.Close()

	receiver := &receiverNATS{
		connection: connection,
		codec:      communication.NewCodecBytes(),
		subs:       make(map[string]*nats.Subscription),
	}

	consumer := &bytesRequestConsumer{}
	err := receiver.Respond(consumer)
	assert.NoError(t, err)

	response, err := connection.Request("bytes-response", []byte("REQUEST"), 100*time.Millisecond)
	assert.NoError(t, err)
	assert.Equal(t, []byte("REQUEST"), *consumer.requestReceived.(*[]byte))
	assert.Equal(t, []byte("RESPONSE"), response.Data)
}
