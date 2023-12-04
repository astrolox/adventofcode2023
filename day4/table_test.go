package main

import (
	"testing"
)

func TestLoadString(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	table := LoadString(input)

	if len(table) != 2 {
		t.Errorf("Expected table length to be 2, got %d", len(table))
	}
}

func TestTotalPoints(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	table := LoadString(input)

	if table.TotalPoints() != 10 {
		t.Errorf("Expected total points to be 10, got %d", table.TotalPoints())
	}
}

func TestCreateNewInstancesFromWinners(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	table := LoadString(input)
	table.CreateNewInstancesFromWinners()

	if table[1].Instances != 2 {
		t.Errorf("Expected instances of card 2 to be 2, got %d", table[1].Instances)
	}
}

func TestTotalInstances(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	table := LoadString(input)
	table.CreateNewInstancesFromWinners()

	if table.TotalInstances() != 3 {
		t.Errorf("Expected total instances to be 3, got %d", table.TotalInstances())
	}
}
