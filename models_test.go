package plex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlexibleBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		wantErr  bool
	}{
		// Boolean literals
		{"true literal", `true`, true, false},
		{"false literal", `false`, false, false},

		// String literals (Plex frequently returns "0" or "1")
		{"true string", `"true"`, true, false},
		{"false string", `"false"`, false, false},
		{"1 string", `"1"`, true, false},
		{"0 string", `"0"`, false, false},

		// Integer literals
		{"1 integer", `1`, true, false},
		{"0 integer", `0`, false, false},

		// Edge cases / Invalid
		{"empty string", `""`, false, true},
		{"null", `null`, false, false},
		{"random string", `"hello"`, false, true},
		{"random number", `42`, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b FlexibleBool
			err := json.Unmarshal([]byte(tt.input), &b)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, bool(b))
			}
		})
	}
}

func TestFlexibleFloat64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		wantErr  bool
	}{
		// Float literals
		{"float literal", `8.5`, 8.5, false},
		{"zero literal", `0`, 0.0, false},

		// String literals
		{"float string", `"8.5"`, 8.5, false},
		{"integer string", `"10"`, 10.0, false},
		{"zero string", `"0"`, 0.0, false},

		// Edge cases / Invalid
		{"empty string", `""`, 0, true},
		{"null", `null`, 0, false},
		{"random string", `"hello"`, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f FlexibleFloat64
			err := json.Unmarshal([]byte(tt.input), &f)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, float64(f))
			}
		})
	}
}

func TestUserPlexTV_Anonymous(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "anonymous boolean false",
			input:    `{"id": 123, "anonymous": false}`,
			expected: false,
		},
		{
			name:     "anonymous boolean true",
			input:    `{"id": 123, "anonymous": true}`,
			expected: true,
		},
		{
			name:     "anonymous string true",
			input:    `{"id": 123, "anonymous": "true"}`,
			expected: true,
		},
		{
			name:     "anonymous null",
			input:    `{"id": 123, "anonymous": null}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var user UserPlexTV
			err := json.Unmarshal([]byte(tt.input), &user)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, bool(user.Anonymous))
		})
	}
}

func TestSetting_Unmarshal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedDef   interface{}
		expectedValue interface{}
	}{
		{
			name:          "integer preference value and default",
			input:         `{"id": "LastAutomaticMappedPort", "default": 0, "value": 25018}`,
			expectedDef:   0.0, // json numbers unmarshal to float64 for interface{}
			expectedValue: 25018.0,
		},
		{
			name:          "string preference value and default",
			input:         `{"id": "some_pref", "default": "hello", "value": "world"}`,
			expectedDef:   "hello",
			expectedValue: "world",
		},
		{
			name:          "boolean preference value and default",
			input:         `{"id": "bool_pref", "default": false, "value": true}`,
			expectedDef:   false,
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Setting
			err := json.Unmarshal([]byte(tt.input), &s)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedDef, s.Default)
			assert.Equal(t, tt.expectedValue, s.Value)
		})
	}
}
