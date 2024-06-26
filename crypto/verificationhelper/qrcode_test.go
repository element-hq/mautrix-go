// Copyright (c) 2024 Sumner Evans
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package verificationhelper_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/element-hq/mautrix-go/crypto/verificationhelper"
)

func TestQRCode_Roundtrip(t *testing.T) {
	var key1, key2 [32]byte
	copy(key1[:], bytes.Repeat([]byte{0x01}, 32))
	copy(key2[:], bytes.Repeat([]byte{0x02}, 32))
	qrCode := verificationhelper.NewQRCode(verificationhelper.QRCodeModeCrossSigning, "test", key1, key2)

	encoded := qrCode.Bytes()
	decoded, err := verificationhelper.NewQRCodeFromBytes(encoded)
	require.NoError(t, err)

	assert.Equal(t, verificationhelper.QRCodeModeCrossSigning, decoded.Mode)
	assert.EqualValues(t, "test", decoded.TransactionID)
	assert.Equal(t, key1, decoded.Key1)
	assert.Equal(t, key2, decoded.Key2)
}

func TestQRCodeDecode(t *testing.T) {
	qrcodeData := []byte{
		0x4d, 0x41, 0x54, 0x52, 0x49, 0x58, 0x02, 0x01, 0x00, 0x20, 0x47, 0x6e, 0x41, 0x65, 0x43, 0x76,
		0x74, 0x57, 0x6a, 0x7a, 0x4d, 0x4f, 0x56, 0x57, 0x51, 0x54, 0x6b, 0x74, 0x33, 0x35, 0x59, 0x52,
		0x55, 0x72, 0x75, 0x6a, 0x6d, 0x52, 0x50, 0x63, 0x38, 0x61, 0x18, 0x32, 0x7c, 0xc3, 0x8c, 0xc2,
		0xa6, 0xc2, 0xb5, 0xc2, 0xa7, 0x50, 0x57, 0x67, 0x19, 0x5e, 0xc3, 0xaf, 0xc2, 0xa0, 0xc2, 0x98,
		0xc2, 0x9d, 0x36, 0xc3, 0xad, 0x7a, 0x10, 0x2e, 0x18, 0x3e, 0x4e, 0xc3, 0x84, 0xc3, 0x81, 0x45,
		0x0c, 0xc2, 0xae, 0x19, 0x78, 0xc2, 0x99, 0x06, 0xc2, 0x92, 0xc2, 0x94, 0xc2, 0x8e, 0xc2, 0xb7,
		0x59, 0xc2, 0x96, 0xc2, 0xad, 0xc3, 0xbd, 0x70, 0x6a, 0x11, 0xc2, 0xba, 0xc2, 0xa9, 0x29, 0xc3,
		0x8f, 0x0d, 0xc2, 0xb8, 0xc2, 0x88, 0x67, 0x5b, 0xc3, 0xb3, 0x01, 0xc2, 0xb0, 0x63, 0x2e, 0xc2,
		0xa5, 0xc3, 0xb3, 0x60, 0xc3, 0x82, 0x04, 0xc3, 0xa3, 0x72, 0x7d, 0x7c, 0x1d, 0xc2, 0xb6, 0xc2,
		0xba, 0xc2, 0x81, 0x1e, 0xc2, 0x99, 0xc2, 0xb8, 0x7f, 0x0a,
	}
	decoded, err := verificationhelper.NewQRCodeFromBytes(qrcodeData)
	require.NoError(t, err)
	assert.Equal(t, verificationhelper.QRCodeModeSelfVerifyingMasterKeyTrusted, decoded.Mode)
	assert.EqualValues(t, "GnAeCvtWjzMOVWQTkt35YRUrujmRPc8a", decoded.TransactionID)
	assert.Equal(t,
		[32]byte{0x18, 0x32, 0x7c, 0xc3, 0x8c, 0xc2, 0xa6, 0xc2, 0xb5, 0xc2, 0xa7, 0x50, 0x57, 0x67, 0x19, 0x5e, 0xc3, 0xaf, 0xc2, 0xa0, 0xc2, 0x98, 0xc2, 0x9d, 0x36, 0xc3, 0xad, 0x7a, 0x10, 0x2e, 0x18, 0x3e},
		decoded.Key1)
	assert.Equal(t,
		[32]byte{0x4e, 0xc3, 0x84, 0xc3, 0x81, 0x45, 0xc, 0xc2, 0xae, 0x19, 0x78, 0xc2, 0x99, 0x6, 0xc2, 0x92, 0xc2, 0x94, 0xc2, 0x8e, 0xc2, 0xb7, 0x59, 0xc2, 0x96, 0xc2, 0xad, 0xc3, 0xbd, 0x70, 0x6a, 0x11},
		decoded.Key2)
}
