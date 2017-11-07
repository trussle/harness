package generators

import (
	"math/rand"
	"reflect"
	"strings"
)

const asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateString creates a random string
func GenerateString(r *rand.Rand, size int) string {
	res := make([]byte, size)
	for k := range res {
		res[k] = byte(asciiChars[r.Intn(len(asciiChars)-1)])
	}
	return string(res)
}

// ASCII creates a value that is simple ascii characters from a-Z0-9.
type ASCII string

// Generate allows ASCII to be used within quickcheck scenarios.
func (ASCII) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(ASCII(GenerateString(r, size)))
}

func (a ASCII) String() string {
	return string(a)
}

// ASCIISlice creates a series of values that are simple ascii characters from
// a-Z0-9
type ASCIISlice []string

// Generate allows ASCIISlice to be used within quickcheck scenarios.
func (ASCIISlice) Generate(r *rand.Rand, size int) reflect.Value {
	res := make([]string, size)
	for k := range res {
		res[k] = GenerateString(r, size)
	}
	return reflect.ValueOf(res)
}

// Slice returns the underlying slice of the type
func (a ASCIISlice) Slice() []string {
	return []string(a)
}

func (a ASCIISlice) String() string {
	return strings.Join(a.Slice(), ",")
}
