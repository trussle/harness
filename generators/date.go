package generators

import (
	"math/rand"
	"reflect"
	"time"
)

// Date creates a random date useful for randomness
type Date int64

// Generate allows Date to be used within quickcheck scenarios.
func (Date) Generate(r *rand.Rand, size int) reflect.Value {
	var (
		min = time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		max = time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

		sec = r.Int63n(max-min) + min
	)
	return reflect.ValueOf(Date(sec))
}

// Date returns the raw underlying date in seconds
func (d Date) Date() time.Time {
	return time.Unix(int64(d), 0)
}

func (d Date) String() string {
	return d.Date().Format(time.RFC3339)
}
