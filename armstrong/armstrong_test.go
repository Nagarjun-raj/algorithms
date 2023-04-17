package main

import "testing"

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"first":  {371, true},
		"second": {139, false},
		"third":  {222, false},
	}
	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			got := armstrongCheck(te.input)
			if te.want != got {
				t.Fatalf("we got %v but we wanted %v", got, te.want)
			} else {
				t.Logf("Successfull")
			}
		})
	}
}
