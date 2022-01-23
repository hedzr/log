package exec

import (
	"bufio"
	"fmt"
	"strings"
)

func toStringSimple(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// LeftPad inserts spaces at beginning of each line
func LeftPad(s string, pad int) string { return LeftPad(s, pad) }

func leftPad(s string, pad int) string {
	if pad <= 0 {
		return s
	}

	var sb strings.Builder
	padstr := strings.Repeat(" ", pad)
	scanner := bufio.NewScanner(bufio.NewReader(strings.NewReader(s)))
	for scanner.Scan() {
		sb.WriteString(padstr)
		sb.WriteString(scanner.Text())
		sb.WriteRune('\n')
	}
	return sb.String()
}

// TrimQuotes strips first and last quote char (double quote or single quote).
func TrimQuotes(s string) string { return trimQuotes(s) }

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// SplitCommandString allows split command-line by quote
// characters (default is double-quote).
//
// In: `bash -c 'echo hello world!'`
// Out: []string{ "bash", "-c", "echo hello world!"}
//
func SplitCommandString(s string, quoteChars ...rune) []string {
	var qc rune = '"'
	var m = map[rune]bool{qc: true}
	for _, q := range quoteChars {
		qc = q
		m[q] = true
	}

	quoted, ch := false, rune(0)

	a := strings.FieldsFunc(s, func(r rune) bool {
		if ch == 0 {
			if _, ok := m[r]; ok {
				quoted, ch = !quoted, r
			}
		} else if ch == r {
			quoted, ch = !quoted, r
		}
		return !quoted && r == ' '
	})

	var b []string
	for _, s := range a {
		b = append(b, trimQuotes(s))
	}

	return a
}
