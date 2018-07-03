package random

import (
	"math"
	"math/rand"
	"time"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Creates a random string based on a variety of options, using
// supplied source of randomness.
//
// If start and end are both 0, start and end are set
// to ' ' and 'z', the ASCII printable
// characters, will be used, unless letters and numbers are both
// false, in which case, start and end are set to 0 and math.MaxInt32.
//
// If set is not nil, characters between start and end are chosen.
//
// This method accepts a user-supplied rand.Rand
// instance to use as a source of randomness. By seeding a single
// rand.Rand instance with a fixed seed and using it for each call,
// the same random sequence of strings can be generated repeatedly
// and predictably.
func Spec0(count uint, start, end int, letters, numbers bool,
	chars []rune, rand *rand.Rand) string {
	if count == 0 {
		return ""
	}
	if start == 0 && end == 0 {
		end = 'z' + 1
		start = ' '
		if !letters && !numbers {
			start = 0
			end = math.MaxInt32
		}
	}
	buffer := make([]rune, count)
	gap := end - start
	for count != 0 {
		count--
		var ch rune
		if len(chars) == 0 {
			ch = rune(rand.Intn(gap) + start)
		} else {
			ch = chars[rand.Intn(gap)+start]
		}
		if letters && ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) ||
			numbers && (ch >= '0' && ch <= '9') ||
			(!letters && !numbers) {
			if ch >= rune(56320) && ch <= rune(57343) {
				if count == 0 {
					count++
				} else {
					buffer[count] = ch
					count--
					buffer[count] = rune(55296 + rand.Intn(128))
				}
			} else if ch >= rune(55296) && ch <= rune(56191) {
				if count == 0 {
					count++
				} else {
					// high surrogate, insert low surrogate before putting it in
					buffer[count] = rune(56320 + rand.Intn(128))
					count--
					buffer[count] = ch
				}
			} else if ch >= rune(56192) && ch <= rune(56319) {
				// private high surrogate, no effing clue, so skip it
				count++
			} else {
				buffer[count] = ch
			}
		} else {
			count++
		}
	}
	return string(buffer)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param start - the position in set of chars to start at
// Param end   - the position in set of chars to end before
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func Spec1(count uint, start, end int, letters, numbers bool) string {
	return Spec0(count, start, end, letters, numbers, nil, defaultRand)
}

// Creates a random string whose length is the number of characters specified.
//
// Characters will be chosen from the set of alpha-numeric
// characters as indicated by the arguments.
//
// Param count - the length of random string to create
// Param letters - if true, generated string will include
//                 alphabetic characters
// Param numbers - if true, generated string will include
//                 numeric characters
func AlphaOrNumeric(count uint, letters, numbers bool) string {
	return Spec1(count, 0, 0, letters, numbers)
}

func String(count uint) string {
	return AlphaOrNumeric(count, false, false)
}

func StringSpec0(count uint, set []rune) string {
	return Spec0(count, 0, len(set)-1, false, false, set, defaultRand)
}

func StringSpec1(count uint, set string) string {
	return StringSpec0(count, []rune(set))
}

// Creates a random string whose length is the number of characters
// specified.
//
// Characters will be chosen from the set of characters whose
// ASCII value is between 32 and 126 (inclusive).
func Ascii(count uint) string {
	return Spec1(count, 32, 127, false, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alphabetic characters.
func Alphabetic(count uint) string {
	return AlphaOrNumeric(count, true, false)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of alpha-numeric characters.
func Alphanumeric(count uint) string {
	return AlphaOrNumeric(count, true, true)
}

// Creates a random string whose length is the number of characters specified.
// Characters will be chosen from the set of numeric characters.
func Numeric(count uint) string {
	return AlphaOrNumeric(count, false, true)
}

func Int(max int) int {
	return defaultRand.Intn(max)
}

// 从0到max中随机出n个不重复的数
func NInt(max, n int) []int {
	if max < 0 || n < 0 {
		return []int{}
	}

	arr := []int{}
	if max < n {
		for i := 0; i <= max; i++ {
			arr = append(arr, i)
		}
	} else {
		for i := 0; i < n; i++ {
		RAND:
			r := Int(max)
			for _, v := range arr {
				if v == r {
					goto RAND
				}
			}
			arr = append(arr, r)
		}
	}
	return arr
}

// [from, to)
func Range(from, to int) int {
	if from == to {
		return from
	}
	if from > to {
		panic("'to' mush be bigger than 'from'")
	}
	return Int(to-from) + from
}
