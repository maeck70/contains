package contains

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		value    any
		expected bool
	}{
		{"String present", []string{"apple", "banana", "cherry"}, "banana", true},
		{"String absent", []string{"apple", "banana", "cherry"}, "grape", false},
		{"Int present", []string{"1", "2", "3"}, 2, true},
		{"Int absent", []string{"1", "2", "3"}, 4, false},
		{"Float64 present", []string{"1.1", "2.2", "3.3"}, 2.2, true},
		{"Float64 absent", []string{"1.1", "2.2", "3.3"}, 4.4, false},
		{"Unsupported type", []string{"apple", "banana", "cherry"}, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.name != "Unsupported type" {
						t.Errorf("Unexpected panic for test %s: %v", tt.name, r)
					}
				}
			}()
			result := Contains(tt.slice, tt.value)
			if result != tt.expected {
				t.Errorf("Contains(%v, %v) = %v; want %v", tt.slice, tt.value, result, tt.expected)
			}
		})
	}
}
