package main

import "testing"

func Test_part2(t *testing.T) {
	t.Run("test2", func(t *testing.T) {
		part2("input-test2.txt")
	})
	t.Run("real", func(t *testing.T) {
		part2("input-real.txt")
	})
}

func Test_part2_line(t *testing.T) {
	type args struct {
		lineNum int
		line    string
	}
	tests := []struct {
		input string
		want  string
	}{
		{
			"two1nine",
			"2two1nine9",
		},
		{
			"eightwothree",
			"8eightwothree3",
		},
		{
			"abcone2threexyz",
			"1abcone2threexyz3",
		},
		{
			"xtwone3four",
			"2xtwone3four4",
		},
		{
			"4nineeightseven2",
			"4nineeightseven2",
		},
		{
			"zoneight234",
			"1zoneight234",
		},
		{
			"7pqrstsixteen",
			"7pqrstsixteen6",
		},
		{
			"eighthree",
			"8eighthree3",
		},
		{
			"sevenine",
			"7sevenine9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			lineNum := 0
			line := tt.input
			t.Logf("DEBUG: input: %-20s want: %s", tt.input, tt.want)
			if got := part2_line(lineNum, line); got != tt.want {
				t.Errorf("part2_line() = %v, want %v", got, tt.want)
			}
		})
	}
}
