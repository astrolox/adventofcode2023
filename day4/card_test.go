package main

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(line)

	if card.Number != 1 {
		t.Errorf("Expected card number to be 1, got %d", card.Number)
	}

	if len(card.WinningNumbers) != 5 {
		t.Errorf("Expected 5 winning numbers, got %d", len(card.WinningNumbers))
	}

	if len(card.SelectedNumbers) != 8 {
		t.Errorf("Expected 8 selected numbers, got %d", len(card.SelectedNumbers))
	}
}

func TestIsWinningNumber(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(line)

	if !card.IsWinningNumber(41) {
		t.Errorf("Expected 41 to be a winning number")
	}

	if card.IsWinningNumber(6) {
		t.Errorf("Expected 6 not to be a winning number")
	}
}

func TestPoints(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(line)

	if card.Points() != 8 {
		t.Errorf("Expected points to be 8, got %d", card.Points())
	}
}

func TestMatches(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(line)

	if card.Matches() != 4 {
		t.Errorf("Expected matches to be 4, got %d", card.Matches())
	}
}
