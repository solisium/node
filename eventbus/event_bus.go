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

package eventbus

import (
	asaskevichEventBus "github.com/asaskevich/EventBus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// EventBus allows subscribing and publishing data by topic
type EventBus interface {
	Publisher
	Subscriber
}

// Publisher publishes events
type Publisher interface {
	Publish(topic string, data interface{})
}

// Subscriber subscribes to events
type Subscriber interface {
	Subscribe(topic string, fn interface{}) error
	SubscribeAsync(topic string, fn interface{}) error
	Unsubscribe(topic string, fn interface{}) error
}

type simplifiedEventBus struct {
	bus asaskevichEventBus.Bus
}

func (simplifiedBus simplifiedEventBus) Unsubscribe(topic string, fn interface{}) error {
	return simplifiedBus.bus.Unsubscribe(topic, fn)
}

func (simplifiedBus simplifiedEventBus) Subscribe(topic string, fn interface{}) error {
	return simplifiedBus.bus.Subscribe(topic, fn)
}

func (simplifiedBus simplifiedEventBus) SubscribeAsync(topic string, fn interface{}) error {
	return simplifiedBus.bus.SubscribeAsync(topic, fn, false)
}

var logLevelsByTopic = map[string]zerolog.Level{
	"ProposalAdded":            zerolog.Disabled,
	"ProposalUpdated":          zerolog.Disabled,
	"ProposalRemoved":          zerolog.Disabled,
	"proposalEvent":            zerolog.Disabled,
	"Statistics":               zerolog.Disabled,
	"Throughput":               zerolog.Disabled,
	"State change":             zerolog.TraceLevel,
	"Session data transferred": zerolog.TraceLevel,
	"Session change":           zerolog.TraceLevel,
	"hermes_promise_received":  zerolog.TraceLevel,
}

func levelFor(topic string) zerolog.Level {
	if level, exist := logLevelsByTopic[topic]; exist {
		return level
	}
	return zerolog.DebugLevel
}

func (simplifiedBus simplifiedEventBus) Publish(topic string, data interface{}) {
	log.WithLevel(levelFor(topic)).Msgf("Published topic=%q event=%+v", topic, data)
	simplifiedBus.bus.Publish(topic, data)
}

// New returns implementation of EventBus
func New() EventBus {
	bus := asaskevichEventBus.New()
	return simplifiedEventBus{bus: bus}
}
