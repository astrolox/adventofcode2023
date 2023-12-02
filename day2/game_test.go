package main

import (
	"reflect"
	"testing"
)

func TestGame_Parse(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		wantGame *Game
	}{
		{
			name: "Game 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			wantGame: &Game{
				Number: 1,
				Turns: []*Turn{
					{Blue: 3, Red: 4},
					{Red: 1, Green: 2, Blue: 6},
					{Green: 2},
				},
			},
		},
		{
			name: "Game 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			wantGame: &Game{
				Number: 2,
				Turns: []*Turn{
					{Blue: 1, Green: 2},
					{Green: 3, Blue: 4, Red: 1},
					{Green: 1, Blue: 1},
				},
			},
		},
		{
			name: "Game 3",
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			wantGame: &Game{
				Number: 3,
				Turns: []*Turn{
					{Green: 8, Blue: 6, Red: 20},
					{Blue: 5, Red: 4, Green: 13},
					{Green: 5, Red: 1},
				},
			},
		},
		{
			name: "Game 4",
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			wantGame: &Game{
				Number: 4,
				Turns: []*Turn{
					{Green: 1, Red: 3, Blue: 6},
					{Green: 3, Red: 6},
					{Green: 3, Blue: 15, Red: 14},
				},
			},
		},
		{
			name: "Game 5",
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			wantGame: &Game{
				Number: 5,
				Turns: []*Turn{
					{Red: 6, Blue: 1, Green: 3},
					{Blue: 2, Red: 1, Green: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{}
			game.Parse(tt.line)
			if !reflect.DeepEqual(game, tt.wantGame) {
				t.Errorf("Parse() = %v, want %v", game, tt.wantGame)
			}
		})
	}
}

func TestGame_IsPossible(t *testing.T) {
	tests := []struct {
		name string
		line string
		want bool
	}{
		{
			name: "Test Case 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: true,
		},
		{
			name: "Test Case 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: true,
		},
		{
			name: "Test Case 3",
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: false,
		},
		{
			name: "Test Case 4",
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: false,
		},
		{
			name: "Test Case 5",
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.line)
			if got := game.IsPossible(); got != tt.want {
				t.Errorf("IsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_MinCubes(t *testing.T) {
	tests := []struct {
		name                         string
		line                         string
		wantRed, wantGreen, wantBlue int
	}{
		{
			name:      "Test Case 1",
			line:      "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			wantRed:   4,
			wantGreen: 2,
			wantBlue:  6,
		},
		{
			name:      "Test Case 2",
			line:      "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			wantRed:   1,
			wantGreen: 3,
			wantBlue:  4,
		},
		{
			name:      "Test Case 3",
			line:      "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			wantRed:   20,
			wantGreen: 13,
			wantBlue:  6,
		},
		{
			name:      "Test Case 4",
			line:      "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			wantRed:   14,
			wantGreen: 3,
			wantBlue:  15,
		},
		{
			name:      "Test Case 5",
			line:      "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			wantRed:   6,
			wantGreen: 3,
			wantBlue:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.line)
			gotRed, gotGreen, gotBlue := game.MinCubes()
			if gotRed != tt.wantRed || gotGreen != tt.wantGreen || gotBlue != tt.wantBlue {
				t.Errorf("MinCubes() = %v, %v, %v, want %v, %v, %v", gotRed, gotGreen, gotBlue, tt.wantRed, tt.wantGreen, tt.wantBlue)
			}
		})
	}
}

func TestGame_MinCubesPower(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "Test Case 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: 48,
		},
		{
			name: "Test Case 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: 12,
		},
		{
			name: "Test Case 3",
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: 1560,
		},
		{
			name: "Test Case 4",
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: 630,
		},
		{
			name: "Test Case 5",
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.line)
			if got := game.MinCubesPower(); got != tt.want {
				t.Errorf("MinCubesPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
