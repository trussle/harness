package matchers_test

import (
	"testing"
	"testing/quick"

	"github.com/trussle/harness/matchers"
)

func TestFloat32AnyMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v float32) bool {
			matcher := matchers.MatchAnyFloat32()
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v string) bool {
			matcher := matchers.MatchAnyFloat32()
			return !matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestFloat32Matcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v float32) bool {
			matcher := matchers.MatchFloat32(v)
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v float32, s string) bool {
			matcher := matchers.MatchFloat32(v)
			return !matcher.Matches(s)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
