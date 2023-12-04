package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Number          int
	WinningNumbers  []int
	SelectedNumbers []int

	Instances int
}

func NewCard(line string) (card *Card) {
	card = &Card{}

	// populate the card id number
	line = strings.TrimPrefix(line, "Card ")
	parts := strings.Split(line, ":")
	number, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	PanicOnError(err)
	card.Number = number

	// split the winning numbers from the selected numbers
	parts = strings.Split(parts[1], "|")

	// populate the winning numbers
	winningNumbers := strings.Split(strings.TrimSpace(parts[0]), " ")
	card.WinningNumbers = make([]int, 0, len(winningNumbers))
	for _, numberStr := range winningNumbers {
		numberStr = strings.TrimSpace(numberStr)
		if len(numberStr) == 0 {
			continue
		}
		number, err = strconv.Atoi(numberStr)
		PanicOnError(err)
		card.WinningNumbers = append(card.WinningNumbers, number)
	}

	// populate the selected numbers
	selectedNumbers := strings.Split(strings.TrimSpace(parts[1]), " ")
	card.SelectedNumbers = make([]int, 0, len(selectedNumbers))
	for _, numberStr := range selectedNumbers {
		numberStr = strings.TrimSpace(numberStr)
		if len(numberStr) == 0 {
			continue
		}
		number, err = strconv.Atoi(numberStr)
		PanicOnError(err)
		card.SelectedNumbers = append(card.SelectedNumbers, number)
	}

	sort.Ints(card.SelectedNumbers)
	sort.Ints(card.WinningNumbers)

	card.Instances = 1

	return card
}

func (card *Card) String() string {
	sb := strings.Builder{}
	sb.WriteString("Card ")
	sb.WriteString(strconv.Itoa(card.Number))
	sb.WriteString(": ")
	sb.WriteString(strings.Join(strings.Fields(fmt.Sprintf("%d", card.WinningNumbers)), " "))
	sb.WriteString(" | ")
	sb.WriteString(strings.Join(strings.Fields(fmt.Sprintf("%d", card.SelectedNumbers)), " "))
	return sb.String()
}

func (card *Card) IsWinningNumber(number int) bool {
	for _, value := range card.WinningNumbers {
		if value == number {
			return true
		}
	}
	return false
}

func (card *Card) Points() (points int) {
	for _, number := range card.SelectedNumbers {
		if card.IsWinningNumber(number) {
			if points <= 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func (card *Card) Matches() (matches int) {
	for _, number := range card.SelectedNumbers {
		if card.IsWinningNumber(number) {
			matches++
		}
	}
	return matches
}
