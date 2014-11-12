// regen is a tool which generates all strings from the regular expression of a
// finite language.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp/syntax"
	"sort"
	"strings"
)

func init() {
	flag.Usage = usage
}

const use = `Usage: regen REGEX
Generate all strings from the regular expression of a finite language.

Examples:
  regen "r(8|9|1[0-5])(b|w|d)?"
  // Output:
  // r8
  // r8b
  // r8d
  // r8w
  // r9
  // r9b
  // r9d
  // r9w
  // ...
  // r15
  // r15b
  // r15d
  // r15w
`

func usage() {
	fmt.Fprintln(os.Stderr, use)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	ss, err := regen(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	sort.Strings(ss)
	for _, s := range ss {
		fmt.Println(s)
	}
}

// regen generates all strings from the regular expression of a finite language.
func regen(s string) ([]string, error) {
	re, err := syntax.Parse(s, syntax.POSIX)
	if err != nil {
		return nil, err
	}
	return gen(re), nil
}

// opName specifies the name of syntax operators.
var opName = map[syntax.Op]string{
	syntax.OpNoMatch:        "OpNoMatch",
	syntax.OpEmptyMatch:     "OpEmptyMatch",
	syntax.OpLiteral:        "OpLiteral",
	syntax.OpCharClass:      "OpCharClass",
	syntax.OpAnyCharNotNL:   "OpAnyCharNotNL",
	syntax.OpAnyChar:        "OpAnyChar",
	syntax.OpBeginLine:      "OpBeginLine",
	syntax.OpEndLine:        "OpEndLine",
	syntax.OpBeginText:      "OpBeginText",
	syntax.OpEndText:        "OpEndText",
	syntax.OpWordBoundary:   "OpWordBoundary",
	syntax.OpNoWordBoundary: "OpNoWordBoundary",
	syntax.OpCapture:        "OpCapture",
	syntax.OpStar:           "OpStar",
	syntax.OpPlus:           "OpPlus",
	syntax.OpQuest:          "OpQuest",
	syntax.OpRepeat:         "OpRepeat",
	syntax.OpConcat:         "OpConcat",
	syntax.OpAlternate:      "OpAlternate",
}

// gen generates all strings from the regular expression of a finite language.
func gen(re *syntax.Regexp) []string {
	switch re.Op {
	case syntax.OpAnyCharNotNL, syntax.OpAnyChar, syntax.OpBeginLine, syntax.OpEndLine, syntax.OpBeginText, syntax.OpEndText, syntax.OpWordBoundary, syntax.OpNoWordBoundary, syntax.OpStar, syntax.OpPlus:
		log.Fatalf("invalid regular expression %q; %v not supported.\n", re, opName[re.Op])
	case syntax.OpEmptyMatch:
		return []string{""}
	case syntax.OpLiteral:
		return []string{string(re.Rune)}
	case syntax.OpCharClass:
		var ss []string
		// The ranges are stored in start/end rune pairs, e.g.
		//    ac = a|b|c
		var start rune
		for i, end := range re.Rune {
			if i%2 == 0 {
				start = end
			} else {
				for c := start; c <= end; c++ {
					ss = append(ss, string(c))
				}
			}
		}
		return ss
	case syntax.OpCapture:
		return gen(re.Sub[0])
	case syntax.OpQuest:
		return append(gen(re.Sub[0]), "")
	case syntax.OpRepeat:
		if re.Max == -1 {
			log.Fatalf("invalid regular expression %q; %v with no upper limit not supported.\n", re, opName[re.Op])
		}
		var ss []string
		for _, s := range gen(re.Sub[0]) {
			for i := re.Min; i <= re.Max; i++ {
				ss = append(ss, strings.Repeat(s, i))
			}
		}
		return ss
	case syntax.OpConcat:
		ss := gen(re.Sub[0])
		for _, sub := range re.Sub[1:] {
			ss = merge(ss, gen(sub))
		}
		return ss
	case syntax.OpAlternate:
		var ss []string
		for _, sub := range re.Sub {
			ss = append(ss, gen(sub)...)
		}
		return ss
	}
	panic("unreachable")
}

// merge returns all combinations of the provided prefixes and suffixes.
func merge(prefixes, suffixes []string) []string {
	var ss []string
	for _, prefix := range prefixes {
		for _, suffix := range suffixes {
			ss = append(ss, prefix+suffix)
		}
	}
	return ss
}
