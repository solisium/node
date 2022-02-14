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

package servicestate

const (
	// AppTopicServiceStatus is used in event bus to announce the service status.
	AppTopicServiceStatus = "Service status"
)

// AppEventServiceStatus represents the service event related information
type AppEventServiceStatus struct {
	ID         string `json:"id"`
	ProviderID string `json:"provider_id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}

// State represents list of possible service states
type State string

const (
	// NotRunning means no service exists
	NotRunning = State("NotRunning")
	// Starting means that service is started but not yet fully established
	Starting = State("Starting")
	// Running means that fully established service exists
	Running = State("Running")
)
