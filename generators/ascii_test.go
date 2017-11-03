package generators_test

import "testing"
import "github.com/trussle/harness/generators"
import "testing/quick"

func TestASCII(t *testing.T) {
	t.Parallel()

	t.Run("ascii", func(t *testing.T) {
		fn := func(a generators.ASCII) bool {
			return validASCII(a.String())
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Fatal(err)
		}
	})
}

func TestASCIISlice(t *testing.T) {
	t.Parallel()

	t.Run("ascii", func(t *testing.T) {
		fn := func(a generators.ASCIISlice) bool {
			return validASCIISlice(a.Slice())
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Fatal(err)
		}
	})
}

func validASCII(a string) bool {
	for _, v := range a {
		if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) {
			return false
		}
	}
	return true
}

func validASCIISlice(a []string) bool {
	for _, v := range a {
		if !validASCII(v) {
			return false
		}
	}
	return true
}
