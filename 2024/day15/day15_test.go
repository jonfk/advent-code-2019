package day15

import (
	"fmt"
	"os"
	"testing"
)

const example = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const smallExample = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const smallExample2 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func TestExample(t *testing.T) {
	state := Parse(example)
	state.ConsumeAllMoves()
	// log.Printf("%#v", state)
	res := state.SumBoxesGPSCoordinates()

	if res != 10092 {
		t.Fatalf("expected = 10092, got = %v", res)
	}
}

func TestSmallExample(t *testing.T) {
	state := Parse(smallExample)
	state.ConsumeAllMoves()
	// log.Printf("%#v", state)
	res := state.SumBoxesGPSCoordinates()

	if res != 2028 {
		t.Fatalf("expected = 2028, got = %v", res)
	}
}

func TestRealInput(t *testing.T) {
	inputTxt, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("err reading file: %v", err)
	}

	state := Parse(string(inputTxt))
	state2 := state.Convert()
	state.ConsumeAllMoves()
	// log.Printf("%#v", state)
	res := state.SumBoxesGPSCoordinates()

	fmt.Printf("# Puzzle\nsum = %v\n", res)

	state2.ConsumeAllMoves()
	res2 := state2.SumBoxesGPSCoordinates()
	fmt.Printf("# Puzzle Part 2\nsum = %v\n", res2)
}

func TestSmallExamplePart2(t *testing.T) {
	state := Parse(smallExample2)

	state2 := state.Convert()
	state2.ConsumeAllMoves()
	state2.SumBoxesGPSCoordinates()
	state2.Print()

	// if res != 9021 {
	// 	t.Fatalf("expected = 9021, got = %v", res)
	// }
}

func TestExamplePart2(t *testing.T) {
	state := Parse(example)

	state2 := state.Convert()
	state2.ConsumeAllMoves()
	res := state2.SumBoxesGPSCoordinates()
	// state2.Print()

	if res != 9021 {
		t.Fatalf("expected = 9021, got = %v", res)
	}
}
