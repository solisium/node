/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
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
	"github.com/rs/zerolog"
)

// Change this to slice as needed
var logLevelsByTopic = map[string]zerolog.Level{
	"*.proposal-register.v3":   zerolog.TraceLevel,
	"*.proposal-ping.v3":       zerolog.TraceLevel,
	"*.proposal-unregister.v3": zerolog.TraceLevel,
}

func levelFor(topic string) zerolog.Level {
	if level, exist := logLevelsByTopic[topic]; exist {
		return level
	}
	return zerolog.DebugLevel
}
