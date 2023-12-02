package main

import "testing"

func TestTurn_IsPossible(t *testing.T) {
	tests := []struct {
		name string
		turn *Turn
		want bool
	}{
		{
			name: "All Maximum",
			turn: &Turn{Red: 12, Green: 13, Blue: 14},
			want: true,
		},
		{
			name: "Too Much Red",
			turn: &Turn{Red: 13, Green: 13, Blue: 14},
			want: false,
		},
		{
			name: "Too Much Green",
			turn: &Turn{Red: 12, Green: 14, Blue: 14},
			want: false,
		},
		{
			name: "Too Much Blue",
			turn: &Turn{Red: 12, Green: 13, Blue: 15},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.turn.IsPossible(); got != tt.want {
				t.Errorf("IsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
