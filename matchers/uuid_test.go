package matchers_test

import (
	"testing"
	"testing/quick"

	"github.com/trussle/harness/matchers"
	"github.com/trussle/uuid"
)

func TestUUIDAnyMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v uuid.UUID) bool {
			matcher := matchers.MatchAnyUUID()
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v string) bool {
			matcher := matchers.MatchAnyUUID()
			return !matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestUUIDMatcher(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		fn := func(v uuid.UUID) bool {
			matcher := matchers.MatchUUID(v)
			return matcher.Matches(v)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("no match", func(t *testing.T) {
		fn := func(v uuid.UUID, s string) bool {
			matcher := matchers.MatchUUID(v)
			return !matcher.Matches(s)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
