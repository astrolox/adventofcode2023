package main

/*
--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola
lift will take you up to the water source, but this is as far as he can bring
you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem:
they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of
surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working
right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine,
but nobody can figure out which one. If you can add up all the part numbers in
the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of
the engine. There are lots of numbers and symbols you don't really understand,
but apparently any number adjacent to a symbol, even diagonally, is a "part
number" and should be included in your sum. (Periods (.) do not count as a
symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not
adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number
is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of
the part numbers in the engine schematic?

*/

import (
	"fmt"
	"log"
	"strconv"
)

func part1(filename string) {

	log.Printf("Part 1: %s", filename)

	grid := LoadFile(filename)

	total := 0

	numberStart := -1
	numberLast := -1

	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {

			switch grid[y][x] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if numberStart == -1 {
					numberStart = x
				}
				numberLast = x
				if x != grid.Width()-1 { // if not last column
					continue
				}
			}
			// this is not a number (or is a number in the last column)

			if numberStart == -1 {
				continue // we don't have a nearby number
			}

			// what is our number?
			numberBytes := grid[y][numberStart : numberLast+1]

			// is this number adjacent to a symbol?
			// we need to check the surrounding cells

			subGrid := grid.SubGrid(numberStart-1, y-1, len(numberBytes)+2, 3)
			if subGrid.ContainsSymbol() != ' ' {
				// this number is not adjacent to a symbol
				PrintRed(fmt.Sprintf("%3d,%-3d: Number %5s is NOT adjacent to a symbol", x, y, string(numberBytes)))
				PrintRed(subGrid.String())
				numberStart = -1 // reset
				continue
			}

			// this number is adjacent to a symbol!
			PrintBlack(fmt.Sprintf("%3d,%-3d: Number %5s is adjacent to a symbol", x, y, string(numberBytes)))
			PrintBlack(subGrid.String())
			numberStart = -1 // reset

			// convert the number to an integer
			number, err := strconv.Atoi(string(numberBytes))
			PanicOnError(err)
			total += number

		}
	}

	log.Printf("Total: %d", total)

	// 535299 = That's not the right answer; your answer is too low.
	// 536576 = That's the right answer! You are one gold star closer to restoring snow operations.

}
