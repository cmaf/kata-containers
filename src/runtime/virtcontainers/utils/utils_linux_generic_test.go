// Copyright (c) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIoctlVhostVsockGuestCid(t *testing.T){
	assert := assert.New(t)

	result := getIoctlVhostVsockGuestCid()
	assert.Equal(uintptr(ioctlVhostVsockSetGuestCid), result)
}
