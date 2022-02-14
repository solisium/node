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

package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifierIdentity_Verify(t *testing.T) {
	message := []byte("Boop!")
	signature := SignatureHex("1f89542f406b2d638fe09cd9912d0b8c0b5ebb4aef67d52ab046973e34fb430a1953576cd19d140eddb099aea34b2985fbd99e716d3b2f96a964141fdb84b32000")

	verifier := NewVerifierIdentity(FromAddress("0x53a835143c0ef3bbcbfa796d7eb738ca7dd28f68"))
	res, _ := verifier.Verify(message, signature)
	assert.True(t, res)
}

func TestVerifierIdentity_VerifyWithUppercaseIdentity(t *testing.T) {
	message := []byte("Boop!")
	signature := SignatureHex("1f89542f406b2d638fe09cd9912d0b8c0b5ebb4aef67d52ab046973e34fb430a1953576cd19d140eddb099aea34b2985fbd99e716d3b2f96a964141fdb84b32000")

	verifier := NewVerifierIdentity(FromAddress("0x53A835143C0EF3BBCBFA796d7eb738CA7DD28F68"))
	res, _ := verifier.Verify(message, signature)
	assert.True(t, res)
}

func TestVerifierIdentity_VerifyWhenWrongSender(t *testing.T) {
	message := []byte("boop!")
	signature := SignatureHex("1f89542f406b2d638fe09cd9912d0b8c0b5ebb4aef67d52ab046973e34fb430a1953576cd19d140eddb099aea34b2985fbd99e716d3b2f96a964141fdb84b32000")

	verifier := NewVerifierIdentity(FromAddress("0x53a835143c0ef3bbcbfa796d7eb738ca7dd28f68"))
	res, _ := verifier.Verify(message, signature)
	assert.False(t, res)
}

func TestVerifierSigned_Verify(t *testing.T) {
	message := []byte("Boop!")
	signature := SignatureHex("1f89542f406b2d638fe09cd9912d0b8c0b5ebb4aef67d52ab046973e34fb430a1953576cd19d140eddb099aea34b2985fbd99e716d3b2f96a964141fdb84b32000")

	verifier := NewVerifierSigned()
	res, _ := verifier.Verify(message, signature)
	assert.True(t, res)
}

func TestVerifierSigned_VerifyWhenMalformedSignature(t *testing.T) {
	message := []byte("Boop!")
	signature := SignatureHex("7369676e6564")

	verifier := NewVerifierSigned()
	res, _ := verifier.Verify(message, signature)
	assert.False(t, res)
}

func TestVerifierSigned_VerifyWhenMessageIsChanged(t *testing.T) {
	message := []byte("Boop changed!")
	signature := SignatureHex("1f89542f406b2d638fe09cd9912d0b8c0b5ebb4aef67d52ab046973e34fb430a1953576cd19d140eddb099aea34b2985fbd99e716d3b2f96a964141fdb84b32000")

	verifier := NewVerifierSigned()
	res, _ := verifier.Verify(message, signature)
	assert.True(t, res)
}
