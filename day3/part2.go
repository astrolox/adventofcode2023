package main

/*

--- Part Two ---

The engineer finds the missing part and installs it in the engine! As the engine
springs to life, you jump in the closest gondola, finally ready to ascend to the
water source.

You don't seem to be going very fast, though. Maybe something is still wrong?
Fortunately, the gondola has a phone labeled "help", so you pick it up and the
engineer answers.

Before you can explain the situation, she suggests that you look out the window.
There stands the engineer, holding a phone in one hand and waving with the
other. You're going so slowly that you haven't even left the station. You exit
the gondola.

The missing part wasn't the only issue - one of the gears in the engine is
wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its
gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so
that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

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

In this schematic, there are two gears. The first is in the top left; it has
part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the
lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear
because it is only adjacent to one part number.) Adding up all of the gear
ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?

*/

import (
	"fmt"
	"log"
	"strconv"
)

func part2(filename string) {

	log.Printf("Part 2: %s", filename)

	grid := LoadFile(filename)

	total := 0

	found := make(map[Point][]int) // asterisk -> numbers

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

			subX := numberStart - 1
			subY := y - 1
			subW := len(numberBytes) + 2
			subH := 3

			subGrid := grid.SubGrid(subX, subY, subW, subH)

			if subGrid.ContainsSymbol() != '*' {
				PrintRed(fmt.Sprintf("%3d,%-3d: Number %5s is NOT adjacent to an asterisk", x, y, string(numberBytes)))
				//PrintRed(subGrid.String())
				numberStart = -1 // reset
				continue
			}

			PrintBlack(fmt.Sprintf("%3d,%-3d: Number %5s is adjacent to an asterisk", x, y, string(numberBytes)))
			//PrintBlack(subGrid.String())
			numberStart = -1 // reset

			// Work out the exact location of the asterisk

			symbolLocations := subGrid.Find('*') // find within the subgrid
			if symbolLocations == nil {
				panic("symbol not found")
			}
			if len(symbolLocations) != 1 {
				panic("too many symbols found")
			}

			// translate the coordinates of the asterisk back to the main grid
			asterisk := symbolLocations[0]
			subX, subY = grid.Clamp(subX, subY)
			asterisk.X += subX
			asterisk.Y += subY

			if _, ok := found[asterisk]; !ok {
				PrintBlue(fmt.Sprintf("Asterisk found at %d,%d (first time)", asterisk.X, asterisk.Y))
			} else {
				PrintBlue(fmt.Sprintf("Asterisk found at %d,%d", asterisk.X, asterisk.Y))
			}

			// convert the number to an integer
			number, err := strconv.Atoi(string(numberBytes))
			PanicOnError(err)

			// add to the found list
			found[asterisk] = append(found[asterisk], number)
		}
	}

	log.Println("FOUND")
	for asterisk, numbers := range found {
		if len(numbers) == 2 {
			ratio := numbers[0] * numbers[1]
			total += ratio
			log.Printf("Asterisk %3d,%-3d: %#v: ratio %d", asterisk.X, asterisk.Y, numbers, ratio)
		} else {
			PrintRed(fmt.Sprintf("Asterisk %3d,%-3d: %#v", asterisk.X, asterisk.Y, numbers))
		}
	}

	log.Printf("Total: %d", total)

	// 75741499 = That's the right answer! You are one gold star closer to restoring snow operations.

}
