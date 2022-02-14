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

package service

// RegistryFactory initiates instance which is able to serve
type RegistryFactory func(options Options) (Service, error)

// Registry holds all pluggable services
type Registry struct {
	factories map[string]RegistryFactory
}

// NewRegistry creates a registry of pluggable services
func NewRegistry() *Registry {
	return &Registry{
		factories: make(map[string]RegistryFactory),
	}
}

// Register registers a new pluggable service
func (registry *Registry) Register(serviceType string, creator RegistryFactory) {
	registry.factories[serviceType] = creator
}

// Create creates pluggable service
func (registry *Registry) Create(serviceType string, options Options) (Service, error) {
	createService, exists := registry.factories[serviceType]
	if !exists {
		return nil, ErrUnsupportedServiceType
	}

	return createService(options)
}
