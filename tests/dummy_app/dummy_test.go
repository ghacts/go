package main

import "testing"

func Test_dummy(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test",
			want: "dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dummy(); got != tt.want {
				t.Errorf("dummy() = %v, want %v", got, tt.want)
			}
		})
	}
}
