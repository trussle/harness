package generators_test

import (
	"testing"
	"testing/quick"
	"time"

	"github.com/trussle/harness/generators"
)

func TestDate(t *testing.T) {
	t.Parallel()

	t.Run("date", func(t *testing.T) {
		fn := func(a generators.Date) bool {
			return validDate(a.Date())
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Fatal(err)
		}
	})
}

func validDate(a time.Time) bool {
	return !a.IsZero()
}
