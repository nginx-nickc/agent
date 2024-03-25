// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNginxGatewayFabric_ParseConfig(t *testing.T) {
	result, err := NewNginxGatewayFabric().ParseConfig()
	// Not implemented yet so error is expected
	assert.Nil(t, result)
	require.Error(t, err)
}

func TestNginxGatewayFabric_Validate(t *testing.T) {
	err := NewNginxGatewayFabric().Validate()
	// Not implemented yet so error is expected )
	require.Error(t, err)
}

func TestNginxGatewayFabric_Apply(t *testing.T) {
	err := NewNginxGatewayFabric().Apply()
	// Not implemented yet so error is expected )
	require.Error(t, err)
}