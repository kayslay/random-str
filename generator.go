package str

import (
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// a map given the start and end range of characters for uppercase alphabets, lowercase alphabets, digits and all characters
	charRangeMap = map[rune][2]int{
		'A': [2]int{65, 90},
		'a': [2]int{97, 122},
		'd': [2]int{48, 57},
		'*': [2]int{32, 127},
	}
	// since symbol are not arranged in the same place on the ascii table, symRanges
	// hold the different start and end positions of symbols on the ascii table
	symRanges = [][2]int{[2]int{33, 47}, [2]int{58, 64}, [2]int{91, 96}, [2]int{123, 126}}
)

// WriteN generates a string with length n. if format is specified, the string will be of the given format
func WriteN(n int, format ...rune) string {
	bytes := make([]byte, n)
	key := '*'
	if len(format) > 0 {
		key = format[0]
	}
	for i := 0; i < n; i++ {
		bytes[i] = randomChar(key)
	}
	return string(bytes)
}

// WriteFromFormat generates a string that follow the given format
func WriteFromFormat(format string) string {
	bytes := make([]byte, len(format))
	for i, v := range format {
		bytes[i] = randomChar(v)
	}
	return string(bytes)

}

// randChar returns a random char based on the value of c.
// if c is not a valid format type, "*" will be used as the format
func randomChar(c rune) byte {
	// check if c is expecting a symbol
	if c == 's' {
		return randForSymbol()
	}
	v, ok := charRangeMap[c]
	if !ok {
		v = charRangeMap['*']
	}
	return randByte(v[0], v[1])
}

// return a random symbol char
func randForSymbol() byte {
	charRange := symRanges[r.Intn(len(symRanges))]
	return randByte(charRange[0], charRange[1])
}

// randByte return a byte which ascii value is within min and max
func randByte(min int, max int) byte {
	return byte(min + r.Intn(max-min+1))
}
