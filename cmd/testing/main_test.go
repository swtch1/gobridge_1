package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 2, 3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d", tt.x, tt.y), func(t *testing.T) {
			sum := add(tt.x, tt.y)
			require.Equal(t, tt.expected, sum)
		})
	}
}
