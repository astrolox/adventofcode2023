package main

import (
	"log"
	"strconv"
	"strings"
)

/*
Test Data

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

type Game struct {
	Number int
	Turns  []*Turn
}

func NewGame(line string) (game *Game) {
	game = &Game{}
	game.Parse(line)
	return game
}

func (game *Game) Parse(line string) {

	parts := strings.Split(line, ":")

	// parse the game id number
	id := parts[0]
	{
		idFields := strings.Fields(id)
		if len(idFields) != 2 {
			log.Panicf("%s: invalid", id)
		}

		idNum, err := strconv.Atoi(idFields[1])
		panicif(err)

		game.Number = idNum
	}

	// parse the game turns
	turns := strings.Split(parts[1], ";")
	game.Turns = make([]*Turn, len(turns))
	for turnIndex, turn := range turns {
		game.Turns[turnIndex] = &Turn{}

		cubes := strings.Split(turn, ",")
		for _, cube := range cubes {

			cubeFields := strings.Fields(cube)
			if len(cubeFields) != 2 {
				log.Panicf("%s: turn %d: cube %s: invalid", id, turnIndex, cube)
			}

			count, err := strconv.Atoi(cubeFields[0])
			panicif(err)
			color := cubeFields[1]

			switch color {
			case "red":
				game.Turns[turnIndex].Red = count
			case "green":
				game.Turns[turnIndex].Green = count
			case "blue":
				game.Turns[turnIndex].Blue = count
			}
		}
	}

}

func (game *Game) IsPossible() bool {
	for _, turn := range game.Turns {
		if !turn.IsPossible() {
			return false
		}
	}
	return true
}

func (game *Game) MinCubes() (red, green, blue int) {
	for _, turn := range game.Turns {
		red = max(red, turn.Red)
		green = max(green, turn.Green)
		blue = max(blue, turn.Blue)
	}
	return red, green, blue
}

func (game *Game) MinCubesPower() int {
	red, green, blue := game.MinCubes()
	return red * green * blue
}
