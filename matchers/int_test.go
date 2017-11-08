package matchers_test

import (
	"testing"
	"testing/quick"

	"github.com/trussle/harness/matchers"
)

func TestIntAnyMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v int) bool {
			matcher := matchers.MatchAnyInt()
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v string) bool {
			matcher := matchers.MatchAnyInt()
			return !matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestIntMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v int) bool {
			matcher := matchers.MatchInt(v)
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v int, s string) bool {
			matcher := matchers.MatchInt(v)
			return !matcher.Matches(s)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
