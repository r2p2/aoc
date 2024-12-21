package d21

import (
	"slices"
	"testing"
)

func Test_numPadPath(t *testing.T) {
	tests := []struct {
		name string
		src  byte
		dst  byte
		want paths
	}{
		{name: "A>0", src: 'A', dst: '0', want: []path{"<"}},
		{name: "A>5", src: 'A', dst: '5', want: []path{"<^^", "^<^", "^^<"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numPadPath(tt.src, tt.dst)

			slices.Sort(got)
			slices.Sort(tt.want)
			if slices.Compare(got, tt.want) != 0 {
				t.Errorf("numPadPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
