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

func TestParseRegex(t *testing.T) {
	program := ParseWithNamedRegex(conditionalExampleInput)

	if len(program) != 2 {
		t.Fatalf("Incorrect Conditional Parse with Regex. Expected = 2, got = %d\nprogram = %#v\n", len(program), program)
	}
}

func TestRunPuzzleInputWithRegex(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatalf("Could not read puzzle input: err = %s", err)
	}

	sum := Run(string(input))

	fmt.Printf("# puzzle\nsum = %d\n", sum)
}

// NOTE: The benchmark showed that regex parsing was much slower than the custom lexer and parser
// > go test -bench=. -benchtime=5s ./day3/...
// goos: darwin
// goarch: arm64
// pkg: jonfk.ca/advent-of-code/2024/day3
// cpu: Apple M4
// BenchmarkParse-10                          89307             66544 ns/op
// BenchmarkParseWithNamedRegex-10             6637            912917 ns/op
// BenchmarkParseWithMultiRegex-10            19225            314754 ns/op
// PASS
// ok      jonfk.ca/advent-of-code/2024/day3       22.133s

var parseResult, parseWithRegexResult []Mul

func BenchmarkParse(b *testing.B) {
	var r []Mul
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		b.Fatalf("Could not read puzzle input: err = %s", err)
	}

	for n := 0; n < b.N; n++ {
		r = Parse(string(input))
	}
	parseResult = r
}

func BenchmarkParseWithNamedRegex(b *testing.B) {
	var r []Mul
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		b.Fatalf("Could not read puzzle input: err = %s", err)
	}

	for n := 0; n < b.N; n++ {
		r = ParseWithNamedRegex(string(input))
	}
	parseWithRegexResult = r
}

func BenchmarkParseWithMultiRegex(b *testing.B) {
	var r []Mul
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		b.Fatalf("Could not read puzzle input: err = %s", err)
	}

	for n := 0; n < b.N; n++ {
		r = ParseWithMultipleRegex(string(input))
	}
	parseWithRegexResult = r
}
