package main

import "testing"

func Test_part1(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		part1("input-test1.txt")
	})
	t.Run("real", func(t *testing.T) {
		part1("input-real.txt")
	})
}
