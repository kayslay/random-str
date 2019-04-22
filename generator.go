package str

import (
	"math/rand"
	"time"
)

var (
	r            = rand.New(rand.NewSource(time.Now().UnixNano()))
	charRangeMap = map[rune][2]int{
		'A': [2]int{65, 90},
		'a': [2]int{97, 122},
		'd': [2]int{48, 57},
		'*': [2]int{32, 127},
	}
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

func randomChar(c rune) byte {
	if c == 's' {
		return randForSymbol()
	}
	v, ok := charRangeMap[c]
	if !ok {
		v = charRangeMap['*']
	}
	return randInt(v[0], v[1])
}

func randForSymbol() byte {
	charRange := symRanges[r.Intn(len(symRanges))]
	return randInt(charRange[0], charRange[1])
}

// randByte return a byte which ascii value is within min and max
func randByte(min int, max int) byte {
	return byte(min + r.Intn(max-min+1))
}
