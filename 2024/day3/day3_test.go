package day3

import (
	"fmt"
	"os"
	"testing"
)

const smallInput = `xmul(2,4)%&`
const exampleInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const smallConditionalExampleInput = `xmul(2,4)&don't()*do()`
const conditionalExampleInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestLex(t *testing.T) {
	tokens := Lex(smallInput)

	if len(tokens) != 5 {
		t.Fatalf("unexpected result tokens = %#v", tokens)
	}

	fmt.Printf("tokens: %#v\n", tokens)
}

func TestLexConditional(t *testing.T) {
	tokens := Lex(smallConditionalExampleInput)

	if len(tokens) != 7 {
		t.Fatalf("unexpected lex result expected 7, got = %d\ntokens = %#v\n", len(tokens), tokens)
	}
}

func TestParse(t *testing.T) {
	program := Parse(exampleInput)

	if len(program) != 4 {
		t.Fatalf("Incorrect parsing got %#v", program)
	}
}

func TestParseConditional(t *testing.T) {
	program := Parse(conditionalExampleInput)

	if len(program) != 2 {
		t.Fatalf("Incorrect Conditional parse. Expected = 2, got = %v\nprogram = %#v\n", len(program), program)
	}
}

func TestRun(t *testing.T) {
	sum := Run(exampleInput)

	if sum != 161 {
		t.Fatalf("Expected = 161, got = %v", sum)
	}
}

func TestRunPuzzleInput(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}

	sum := Run(string(input))

	fmt.Printf("# puzzle\nsum = %d\n", sum)
}
