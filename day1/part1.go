package main

/*
--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take
a look. The Elves have even given you a map; on it, they've used stars to mark
the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you
need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day
in the Advent calendar; the second puzzle is unlocked when you complete the
first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough")
and where they're even sending you ("the sky") and why your map looks mostly
blank ("you sure ask a lot of questions") and hang on did you just say the sky
("of course, where do you think snow comes from") when you realize that the
Elves are already loading you into a trebuchet ("please hold still, we need to
strap you in").

As they're making the final adjustments, they discover that their calibration
document (your puzzle input) has been amended by a very young Elf who was
apparently just excited to show off her art skills. Consequently, the Elves are
having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line
originally contained a specific calibration value that the Elves now need to
recover. On each line, the calibration value can be found by combining the first
digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and
77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the
calibration values?

Your puzzle answer was 54081.

*/

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"

	re "regexp"
)

func part1(filename string) {

	log.Printf("Part 1: %s", filename)

	nums, err := re.Compile(`-?\d`)
	panicif(err)

	input, err := os.ReadFile(filename)
	panicif(err)

	scanner := bufio.NewScanner(bytes.NewReader(input))

	c := 0
	total := 0

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		found := nums.FindAllString(line, -1)
		c++

		number := ""

		if len(found) == 0 {
			log.Printf("No numbers found in line %d: %s", c, line)
			continue
		}
		if len(found) == 1 {
			number = found[0] + found[0]
		}
		if len(found) > 1 {
			number = found[0] + found[len(found)-1]
		}
		num, err := strconv.Atoi(number)
		panicif(err)

		log.Printf("Line %4d: %60s = %4d", c, line, num)

		total += num
	}

	// 17771531 is too high
	// 941791 is too high
	// 37071 is too low
	// 54081 = That's the right answer! You are one gold star closer to restoring snow operations.

	log.Printf("Total: %d", total)

}
