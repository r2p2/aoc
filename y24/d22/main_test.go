package d22

import (
	"testing"
)

func Test_nextPrn(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secret int
		want   int
	}{
		{secret: 123, want: 15887950},
		{secret: 15887950, want: 16495136},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextSecret(tt.secret)
			if got != tt.want {
				t.Errorf("nextPrn() = %v, want %v", got, tt.want)
			}
		})
	}
}
