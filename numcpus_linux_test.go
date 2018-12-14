// Copyright 2018 Tobias Klauser
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package numcpus

import "testing"

func TestParseCPURange(t *testing.T) {
	testCases := []struct {
		cpus string
		n    int
	}{
		{"", 0},
		{"0", 1},
		{"0-1", 2},
		{"0-7", 8},
		{"1-7", 7},
		{"1-15", 15},
		{"0,2-4,7", 5},
		{"0,2-4,7-15", 13},
	}

	for _, tc := range testCases {
		n, err := parseCPURange(tc.cpus)
		if err != nil {
			t.Errorf("failed to parse CPU range: %v", err)
		}

		if n != tc.n {
			t.Errorf("parseCPURange(%s) = %d, expected %d", tc.cpus, n, tc.n)
		}
	}

	str := "invalid"
	_, err := parseCPURange(str)
	if err == nil {
		t.Errorf("parseCPURange(%s) unexpectedly succeeded", str)
	}
}
