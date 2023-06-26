package jsonmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathParse(t *testing.T) {
	tests := []struct {
		name        string
		testPath    string
		expectErr   bool
		expectValue []interface{}
	}{
		{
			name:        "Valid Simple Path",
			testPath:    "path1.path2",
			expectErr:   false,
			expectValue: []interface{}{"path1", "path2"},
		},
		{
			name:        "Valid Empty Path",
			testPath:    "",
			expectErr:   false,
			expectValue: []interface{}{},
		},
		{
			name:        "Valid Path With Extra Dots",
			testPath:    ".path1..path2.",
			expectErr:   false,
			expectValue: []interface{}{"path1", "path2"},
		},
		{
			name:        "Valid Path With Array Index",
			testPath:    "path1.path2[0]",
			expectErr:   false,
			expectValue: []interface{}{"path1", "path2", 0},
		},
		{
			name:        "Valid Path With Nested Array Index",
			testPath:    "path1.path2[0][1]",
			expectErr:   false,
			expectValue: []interface{}{"path1", "path2", 0, 1},
		},
		{
			name:        "Valid Path Beginning With Array Index",
			testPath:    "[0].path1.path2",
			expectErr:   false,
			expectValue: []interface{}{0, "path1", "path2"},
		},
		{
			name:        "Valid Path With Nested Array Index With Extra Dots",
			testPath:    ".path1..path2.[0].[1].",
			expectErr:   false,
			expectValue: []interface{}{"path1", "path2", 0, 1},
		},
		{
			name:        "Invalid Path With Non-Numeric Array Index",
			testPath:    "path1.path2[A]",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Invalid Path With Array Index Followed By Key",
			testPath:    "path1.[0]path2",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Invalid Path Beginning With Array Index Followed By Key",
			testPath:    "[0]path1.path2",
			expectErr:   true,
			expectValue: nil,
		},
		{
			name:        "Invalid Path With Nested Array Index Followed By Key",
			testPath:    ".path1[0][1]path2",
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := pathParse(tt.testPath)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
