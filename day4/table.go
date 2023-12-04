package main

import (
	"log"
	"os"
	"strings"
)

type Table []*Card

func LoadFile(filename string) (table Table) {
	input, err := os.ReadFile(filename)
	PanicOnError(err)
	return LoadString(string(input))
}

func LoadString(input string) (table Table) {
	table = Table{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		table = append(table, NewCard(line))
	}
	return table
}

func (table Table) TotalPoints() (total int) {
	for _, card := range table {
		total += card.Points()
	}
	return total
}

func (table Table) CreateNewInstancesFromWinners() {
	for x, card := range table {
		instances := card.Instances
		matches := card.Matches()
		log.Printf("%d: Card %d: %d matches with %d instances", x, card.Number, matches, instances)
		if matches == 0 {
			continue
		}

		for y := x + 1; y <= x+matches; y++ {
			if y >= len(table) {
				break
			}
			log.Printf(
				"   >>> %d: Card %d: + %d copies = %d",
				y,
				table[y].Number,
				instances,
				table[y].Instances+instances,
			)
			table[y].Instances += instances
		}

	}
}

func (table Table) TotalInstances() (total int) {
	for _, card := range table {
		total += card.Instances
	}
	return total
}
