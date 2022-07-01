package fuzztest

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    []byte
		b    []byte
		want bool
	}{
		// TODO: Add test cases.
		{"right case", []byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}, true},
		{"right case", []byte{'a', 'b', 'c'}, []byte{'b', 'c', 'd'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.a, tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
func FuzzEqual(f *testing.F) {
	//f.Add([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'})
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
