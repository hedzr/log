package dir

import "strings"

// IsWildMatch provides a wild-matching ('*' and '?') test.
//
// For examples:
//
//     output := IsWildMatch("aa", "aa")
//     expectTrue(t, output)
//
//     output = IsWildMatch("aaaa", "*")
//     expectTrue(t, output)
//
//     output = IsWildMatch("ab", "a?")
//     expectTrue(t, output)
//
//     output = IsWildMatch("adceb", "*a*b")
//     expectTrue(t, output)
//
//     output = IsWildMatch("aa", "a")
//     expectFalse(t, output)
//
//     output = IsWildMatch("mississippi", "m??*ss*?i*pi")
//     expectFalse(t, output)
//
//     output = IsWildMatch("acdcb", "a*c?b")
//     expectFalse(t, output)
//
func IsWildMatch(s string, p string) bool {
	runeInputArray := []rune(s)
	runePatternArray := []rune(p)
	if len(runeInputArray) > 0 && len(runePatternArray) > 0 {
		if runePatternArray[len(runePatternArray)-1] != '*' && runePatternArray[len(runePatternArray)-1] != '?' && runeInputArray[len(runeInputArray)-1] != runePatternArray[len(runePatternArray)-1] {
			return false
		}
	}
	return isMatchUtil([]rune(s), []rune(p), 0, 0, len([]rune(s)), len([]rune(p)))
}

func isMatchUtil(input, pattern []rune, inputIndex, patternIndex int, inputLength, patternLength int) bool {

	if inputIndex == inputLength && patternIndex == patternLength {
		return true
	} else if patternIndex == patternLength {
		return false
	} else if inputIndex == inputLength {
		if pattern[patternIndex] == '*' && restPatternStar(pattern, patternIndex+1, patternLength) {
			return true
		} else {
			return false
		}
	}

	if pattern[patternIndex] == '*' {
		return isMatchUtil(input, pattern, inputIndex, patternIndex+1, inputLength, patternLength) ||
			isMatchUtil(input, pattern, inputIndex+1, patternIndex, inputLength, patternLength)
	}

	if pattern[patternIndex] == '?' {
		return isMatchUtil(input, pattern, inputIndex+1, patternIndex+1, inputLength, patternLength)
	}

	if inputIndex < inputLength {
		if input[inputIndex] == pattern[patternIndex] {
			return isMatchUtil(input, pattern, inputIndex+1, patternIndex+1, inputLength, patternLength)
		} else {
			return false
		}
	}

	return false
}

func restPatternStar(pattern []rune, patternIndex int, patternLength int) bool {
	for patternIndex < patternLength {
		if pattern[patternIndex] != '*' {
			return false
		}
		patternIndex++
	}
	return true
}

// ToBool translate a value to boolean
func ToBool(val interface{}, defaultVal ...bool) (ret bool) {
	if v, ok := val.(bool); ok {
		return v
	}
	if v, ok := val.(int); ok {
		return v != 0
	}
	if v, ok := val.(string); ok {
		return toBool(v, defaultVal...)
	}
	for _, vv := range defaultVal {
		ret = vv
	}
	return
}

func toBool(val string, defaultVal ...bool) (ret bool) {
	//ret = ToBool(val, defaultVal...)
	switch strings.ToLower(val) {
	case "1", "y", "t", "yes", "true", "ok", "on":
		ret = true
	case "":
		for _, vv := range defaultVal {
			ret = vv
		}
	}
	return
}
