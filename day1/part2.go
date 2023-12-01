package main

/*

https://adventofcode.com/2023

--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are
actually spelled out with letters: one, two, three, four, five, six, seven,
eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last
digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76.
Adding these together produces 281.

What is the sum of all of the calibration values?

Your puzzle answer was 54649.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

*/

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(filename string) {

	log.Printf("Part 2: %s", filename)

	// read
	input, err := os.ReadFile(filename)
	panicif(err)

	// process

	// read line by line (without a scanner so that we can modify the lines)
	lines := strings.Split(string(input), "\n")
	for lineNum, line := range lines {
		lines[lineNum] = part2_line(lineNum, line)
	}

	// write

	input = []byte(strings.Join(lines, "\n"))
	filename = WriteTempFile(filename+"-modified-by-part2-", input)
	defer os.Remove(filename)

	// run part1 on the modified file

	part1(filename)

	// 54649 = That's the right answer! You are one gold star closer to restoring snow operations.

}

func part2_line(lineNum int, line string) string {

	log.Printf("DEBUG: line %d: %s", lineNum, line)

	// eightwothree should equal 8wo3
	// but a simple search and replace (in numerical order) won't work because
	// eightwothree = eigh2three = eigh23 = 23
	// changing the search and replace order doesn't help either
	// so switched to a multiple pass scanning approach

	targetsForwards := []string{
		"one",   // index 0 = value 1 = length 3
		"two",   // index 1 = value 2 = length 3
		"three", // index 2 = value 3 = length 5
		"four",  // index 3 = value 4 = length 4
		"five",  // index 4 = value 5 = length 4
		"six",   // index 5 = value 6 = length 3
		"seven", // index 6 = value 7 = length 5
		"eight", // index 7 = value 8 = length 5
		"nine",  // index 8 = value 9 = length 4
	}
	targetsBackwards := make([]string, len(targetsForwards))
	for i, target := range targetsForwards {
		targetsBackwards[i] = Reverse(target)
	}

	findWrittenNumber := func(line string, targets []string) (string, int) {

		match := func(s string) string {
			for _, target := range targets {
				if s == target {
					return target
				}
			}
			return ""
		}

		getNumber := func(s string) int {
			for i, target := range targets {
				if s == target {
					return i + 1
				}
			}
			return 0
		}

		for pos := 0; pos < len(line); pos++ {
			for _, x := range []int{3, 4, 5} {
				if pos >= x {
					grab := line[pos-x : pos]
					matched := match(grab)
					if matched != "" {
						// do something
						number := getNumber(matched)
						log.Printf("DEBUG: %20s pos %2d = %5s = %d", line, pos, matched, number)
						return strconv.Itoa(number), pos
						//} else {
						//	log.Printf("line %d: %20s pos %d, x %d, grab %s = no match", lineNum, line, pos, x, grab)
					}
					//} else {
					//	log.Printf("line %d: %20s pos %d, x %d = x too big", lineNum, line, pos, x)
				}
			}
		}

		return "", 0
	}

	findNumericNumber := func(line string, targets []string) (bool, int) {
		for pos := 0; pos < len(line); pos++ {
			switch line[pos] {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				return true, pos
			}
		}
		return false, len(line)
	}

	log.Printf("DEBUG: finding first")
	firstNumeric := -1
	{
		// first numeric number
		ok, pos := findNumericNumber(line, targetsForwards)
		if ok {
			firstNumeric = pos
		}

		// first written number
		firstWritten := -1
		found, pos := findWrittenNumber(line, targetsForwards)
		if found != "" {
			firstWritten = pos
			log.Printf("DEBUG: firstNumeric %d, firstWritten %d", firstNumeric, firstWritten)

			if firstNumeric == -1 || firstNumeric >= firstWritten {
				line = found + line // prepend so this is the first number
			}
		}
	}

	log.Printf("DEBUG: finding last")
	lastNumeric := -1
	{
		lineReversed := Reverse(line)

		// last numeric number
		ok, pos := findNumericNumber(lineReversed, targetsBackwards)
		if ok {
			lastNumeric = pos
		}

		// last written number
		lastWritten := -1
		found, pos := findWrittenNumber(lineReversed, targetsBackwards)
		if found != "" {
			lastWritten = pos
			log.Printf("DEBUG: lastNumeric %d, lastWritten %d", lastNumeric, lastWritten)

			if lastNumeric == -1 || lastNumeric >= lastWritten {
				line = line + found // append so this is the last number
			}
		}

	}

	return line
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func WriteTempFile(pattern string, contents []byte) (filename string) {

	fileHandle, err := os.CreateTemp("", pattern)
	panicif(err)

	filename = fileHandle.Name()
	log.Printf("creating temp file: %s", filename)

	_, err = fileHandle.Write(contents)
	panicif(err)
	panicif(fileHandle.Close())

	return filename
}
