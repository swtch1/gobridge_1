package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestAdd(t *testing.T) {
// 	t.Parallel()
//
// 	tests := []struct {
// 		name     string
// 		x        int
// 		y        int
// 		expected int
// 	}{
// 		{
// 			name:     "basic",
// 			x:        1,
// 			y:        2,
// 			expected: 3,
// 		},
// 		{
// 			name:     "big numbers",
// 			x:        101,
// 			y:        202,
// 			expected: 303,
// 		},
// 		{
// 			name:     "big numbers",
// 			x:        101,
// 			y:        202,
// 			expected: 303,
// 		},
// 		{
// 			name:     "big numbers",
// 			x:        101,
// 			y:        202,
// 			expected: 303,
// 		},
// 		{
// 			name:     "big numbers",
// 			x:        101,
// 			y:        202,
// 			expected: 303,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := add(1, 2)
// 			require.Equal(t, 3, r)
// 		})
// 	}
// }

func TestNameCache_SetsAndGets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fName string
		lName string
	}{
		{
			fName: "john",
			lName: "smith",
		},
	}

	for _, tt := range tests {
		t.Run(tt.fName+" "+tt.lName, func(t *testing.T) {
			// arrange
			cache := newNameCache()
			cache.cache = make(map[string]string)

			// act
			cache.set(tt.fName, tt.lName)
			v := cache.get(tt.fName)

			// assert
			require.Equal(t, tt.lName, v, "the name was not expected")
		})
	}
}
