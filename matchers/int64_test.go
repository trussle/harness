package matchers_test

import (
	"testing"
	"testing/quick"

	"github.com/trussle/harness/matchers"
)

func TestInt64AnyMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v int64) bool {
			matcher := matchers.MatchAnyInt64()
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v string) bool {
			matcher := matchers.MatchAnyInt64()
			return !matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestInt64Matcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v int64) bool {
			matcher := matchers.MatchInt64(v)
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v int64, s string) bool {
			matcher := matchers.MatchInt64(v)
			return !matcher.Matches(s)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
