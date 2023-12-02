package main

type Turn struct {
	Red   int
	Green int
	Blue  int
}

func (turn *Turn) IsPossible() bool {
	if turn.Red > 12 {
		return false
	}
	if turn.Green > 13 {
		return false
	}
	if turn.Blue > 14 {
		return false
	}
	return true
}
