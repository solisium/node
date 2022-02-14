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

package openvpn

import (
	"bufio"
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

// ValidateConfig is function which takes VPNConfig as argument, checks it and returns error if validation fails
type ValidateConfig func(config VPNConfig) error

// ConfigValidator represents structure which contains list of validating functions
type ConfigValidator struct {
	validators []ValidateConfig
}

// NewDefaultValidator returns ConfigValidator with predefined list of validating functions
func NewDefaultValidator() *ConfigValidator {
	return &ConfigValidator{
		validators: []ValidateConfig{
			validProtocol,
			validIPFormat,
			validTLSPresharedKey,
			validCACertificate,
		},
	}
}

// IsValid function checks if provided config is valid against given config validator and returns first encountered error
func (v *ConfigValidator) IsValid(config VPNConfig) error {
	for _, validator := range v.validators {
		if err := validator(config); err != nil {
			return err
		}
	}
	return nil
}

func validProtocol(config VPNConfig) error {
	switch config.RemoteProtocol {
	case
		"udp",
		"tcp":
		return nil
	}
	return errors.New("invalid protocol: " + config.RemoteProtocol)
}

func validIPFormat(config VPNConfig) error {
	parsed := net.ParseIP(config.RemoteIP)
	if parsed == nil {
		return errors.New("unable to parse ip address " + config.RemoteIP)
	}
	if parsed.To4() == nil {
		return errors.New("IPv4 address is expected")
	}
	return nil
}

func validTLSPresharedKey(config VPNConfig) error {
	_, err := FormatTLSPresharedKey(config)
	return err
}

// FormatTLSPresharedKey formats preshared key (PEM blocks with data encoded to hex) are taken from
// openvpn --genkey --secret static.key, which is openvpn specific.
// it reformats key from single line to multiline fixed length strings.
func FormatTLSPresharedKey(config VPNConfig) (VPNConfig, error) {
	contentScanner := bufio.NewScanner(bytes.NewBufferString(config.TLSPresharedKey))
	for contentScanner.Scan() {
		line := contentScanner.Text()
		//skip empty lines or comments
		if len(line) > 0 || strings.HasPrefix(line, "#") {
			break
		}
	}
	if err := contentScanner.Err(); err != nil {
		return VPNConfig{}, contentScanner.Err()
	}
	header := contentScanner.Text()
	if header != "-----BEGIN OpenVPN Static key V1-----" {
		return VPNConfig{}, errors.New("Invalid key header: " + header)
	}

	var key string
	for contentScanner.Scan() {
		line := contentScanner.Text()
		if line == "-----END OpenVPN Static key V1-----" {
			break
		} else {
			key = key + line
		}
	}
	if err := contentScanner.Err(); err != nil {
		return VPNConfig{}, err
	}
	// 256 bytes key is 512 bytes if encoded to hex
	if len(key) != 512 {
		return VPNConfig{}, errors.New("invalid key length")
	}

	var buff = &bytes.Buffer{}
	fmt.Fprintln(buff, "-----BEGIN OpenVPN Static key V1-----")
	left := key
	for ; len(left) > 64; left = left[64:] {
		fmt.Fprintln(buff, left[0:64])
	}
	fmt.Fprintln(buff, left)
	fmt.Fprintln(buff, "-----END OpenVPN Static key V1-----")
	config.TLSPresharedKey = buff.String()

	return config, nil
}

func validCACertificate(config VPNConfig) error {
	pemBlock, _ := pem.Decode([]byte(config.CACertificate))
	if pemBlock.Type != "CERTIFICATE" {
		return errors.New("invalid CA certificate. Certificate block expected")
	}
	//if we parse it correctly - at least structure is right
	_, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}
	return nil
}
