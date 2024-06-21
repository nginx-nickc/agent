// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"testing"

	"go.uber.org/goleak"

	"github.com/nginx/agent/v3/test/helpers"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, helpers.GoLeakOptions...)
}