// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Keep the original license.

// Copyright 2019 Dmitry A. Mottl. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.

package ctimefmt

import (
	"testing"
	"time"
)

var format1 string = "%Y-%m-%d %H:%M:%S.%f"
var format2 string = "%Y-%m-%d %l:%M:%S.%L %P, %a"
var value1 string = "2019-01-02 15:04:05.666666"
var value2 string = "2019-01-02 3:04:05.666 pm, Wed"
var dt1 time.Time = time.Date(2019, 1, 2, 15, 4, 5, 666666000, time.UTC)
var dt2 time.Time = time.Date(2019, 1, 2, 15, 4, 5, 666000000, time.UTC)

func TestFormat(t *testing.T) {
	s, err := Format(format1, dt1)
	if err != nil {
		t.Fatal(err)
	}
	if s != value1 {
		t.Errorf("Given: %v, expected: %v", s, value1)
	}

	s, err = Format(format2, dt1)
	if err != nil {
		t.Fatal(err)
	}
	if s != value2 {
		t.Errorf("Given: %v, expected: %v", s, value2)
	}
}

func TestParse(t *testing.T) {
	dt_, err := Parse(format1, value1)
	if err != nil {
		t.Error(err)
	} else if dt_ != dt1 {
		t.Errorf("Given: %v, expected: %v", dt_, dt1)
	}

	dt_, err = Parse(format2, value2)
	if err != nil {
		t.Error(err)
	} else if dt_ != dt2 {
		t.Errorf("Given: %v, expected: %v", dt_, dt2)
	}
}
