package numlang

import (
	"fmt"
	"regexp"
	"strings"
)

type RomanNumeral = string

var mr = []string{"M", "D", "C", "L", "X", "V", "I"}

const min, max = 1, 1000

func Int(rn RomanNumeral) int {
	var total int = 0
	var r int = max
	// retrieve for each iteration
	for i, char := range mr {
		reg := regexp.MustCompile(fmt.Sprintf("^%s+", char))
		if sub := reg.FindStringSubmatch(rn); sub != nil {
			total += len(sub) * r
			rn = rn[len(sub)-1:]
		}
		r = next_range(i, r)
	}
	return total
}

// New verify is exist unsupported character and return
// RomanNumeral, or empty RomanNumeral in error case
func New(str string) RomanNumeral {
	reg := regexp.MustCompile(fmt.Sprintf("[^%s]+", strings.Join(mr, "")))
	if len(reg.FindStringSubmatch(str)) > 0 {
		return ""
	}
	return RomanNumeral(str)
}

// Parse uint to roman numbers
func Parse(i uint) RomanNumeral {
	seq := parse_seq(i)
	romseq := opt_seq(seq)
	return romseq
}

// parse_seq parse uint to un-optimized roman numerals string
func parse_seq(i uint) string {
	r := max
	render := ""
	for index, char := range mr {
		for i >= uint(r) {
			i = i - uint(r)
			render = fmt.Sprint(render, char)
		}
		r = next_range(index, r)
	}
	return render
}

// opt_seq optimize rendering of given roman numerals
func opt_seq(str string) string {
	for i, char := range mr[1:] {
		exp := fmt.Sprintf("%s{4}", char)
		reg := regexp.MustCompile(exp)
		match := reg.FindString(str)
		if match != "" {
			prefix := repeat(char, 5-len(match))
			str = strings.ReplaceAll(str, match, fmt.Sprint(prefix, mr[i]))
		}
	}
	return str
}

// repeat get string of given string that append 'num' times
func repeat(str string, num int) string {
	seq := ""
	for num > 0 {
		seq = fmt.Sprint(seq, str)
		num--
	}
	return seq
}

// next_range to perform
func next_range(index, r int) int {
	// modulo to 0 is letter 'M' corresponding to 1000, if is 1, only can be
	switch index % 2 {
	case 0: // must be 'M', 'C', 'X'
		r = r / 2

	case 1: // must be 'D', 'L', 'V'
		r = r / 5
	}
	return r
}
