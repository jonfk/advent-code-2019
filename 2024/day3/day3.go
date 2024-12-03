package day3

import (
	"strconv"
)

type itemType int

const (
	itemMulStart itemType = iota
	itemMulEnd
	itemComma
	itemNumber
	itemDo
	itemDoNot
)

type Item struct {
	typ itemType
	val string
}

type Mul struct {
	val1 int
	val2 int
}

type Lexer struct {
	input string
	start int
	pos   int
	items []Item
}

type stateFn func(l *Lexer) stateFn

func Run(input string) int {
	program := Parse(input)
	sum := ExecuteProgram(program)
	return sum
}

func ExecuteProgram(program []Mul) int {
	var sum int

	for _, instr := range program {
		sum += instr.val1 * instr.val2
	}
	return sum
}

func Parse(input string) []Mul {
	items := Lex(input)
	var program []Mul
	var isDontActive bool

	for i := 0; i < len(items); {
		if items[i].typ == itemDoNot {
			isDontActive = true
			i += 1
		} else if items[i].typ == itemDo {
			isDontActive = false
			i += 1
		} else if !isDontActive && items[i].typ == itemMulStart && i+4 < len(items) && items[i+1].typ == itemNumber && items[i+2].typ == itemComma && items[i+3].typ == itemNumber && items[i+4].typ == itemMulEnd {
			val1, _ := strconv.Atoi(items[i+1].val)
			val2, _ := strconv.Atoi(items[i+3].val)
			program = append(program, Mul{val1: val1, val2: val2})
			i += 5
		} else {
			i += 1
		}
	}
	return program
}

func Lex(input string) []Item {
	lexer := &Lexer{
		input: input,
	}

	for state := lexCorrupted(lexer); ; {
		if lexer.pos >= len(input) {
			break
		}
		state = state(lexer)
	}
	return lexer.items
}

func lexCorrupted(l *Lexer) stateFn {
	if l.input[l.pos] == 'm' {
		return lexMul
	} else if l.input[l.pos] == 'd' {
		return lexDoDont
	} else {
		l.pos += 1
		l.start = l.pos
		return lexCorrupted
	}
}

func lexMul(l *Lexer) stateFn {
	if l.pos+4 <= len(l.input) && l.input[l.pos:l.pos+4] == "mul(" {
		l.items = append(l.items, Item{typ: itemMulStart, val: string(l.input[l.pos : l.pos+4])})
		l.pos += 4
		l.start = l.pos
		return lexNumber
	} else {
		l.pos += 1
		l.start = l.pos
		return lexCorrupted
	}
}

func lexNumber(l *Lexer) stateFn {
	if IsNumber(l.input[l.pos]) {
		l.pos += 1
		return lexNumber
	} else {
		if l.pos > l.start {
			l.items = append(l.items, Item{typ: itemNumber, val: string(l.input[l.start:l.pos])})
		}
		l.start = l.pos
		return lexCommaOrEndMul
	}
}

func lexCommaOrEndMul(l *Lexer) stateFn {
	if l.input[l.pos] == ',' {
		l.items = append(l.items, Item{typ: itemComma, val: string(l.input[l.pos])})
		l.pos += 1
		l.start = l.pos
		return lexNumber
	} else if l.input[l.pos] == ')' {
		l.items = append(l.items, Item{typ: itemMulEnd, val: string(l.input[l.pos])})
		l.pos += 1
		l.start = l.pos
		return lexCorrupted
	} else {
		return lexCorrupted
	}
}

func lexDoDont(l *Lexer) stateFn {
	if l.pos+7 <= len(l.input) && l.input[l.pos:l.pos+7] == "don't()" {
		l.items = append(l.items, Item{typ: itemDoNot, val: string(l.input[l.pos : l.pos+7])})
		l.pos += 7
		l.start = l.pos
	} else if l.pos+4 <= len(l.input) && l.input[l.pos:l.pos+4] == "do()" {
		l.items = append(l.items, Item{typ: itemDo, val: string(l.input[l.pos : l.pos+4])})
		l.pos += 4
		l.start = l.pos
	} else {
		l.pos += 1
		l.start = l.pos
	}
	return lexCorrupted
}

func IsNumber(x byte) bool {
	switch x {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}
