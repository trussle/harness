package requests_test

import (
	"testing"

	"github.com/trussle/harness/requests"
)

func TestRange(t *testing.T) {
	t.Parallel()

	t.Run("match", func(t *testing.T) {
		r := requests.Range{
			Min: 200,
			Max: 299,
		}
		for i := 0; i < 100; i++ {
			if expected, actual := true, r.Match(i+200); expected != actual {
				t.Errorf("expected: %t, actual: %t", expected, actual)
			}
		}
	})

	t.Run("no match", func(t *testing.T) {
		r := requests.Range{
			Min: 200,
			Max: 299,
		}
		for i := 0; i < 100; i++ {
			if expected, actual := false, r.Match(i); expected != actual {
				t.Errorf("expected: %t, actual: %t", expected, actual)
			}
		}

		for i := 0; i < 100; i++ {
			if expected, actual := false, r.Match(i+300); expected != actual {
				t.Errorf("expected: %t, actual: %t", expected, actual)
			}
		}
	})

	t.Run("string", func(t *testing.T) {
		r := requests.Range{
			Min: 200,
			Max: 299,
		}
		want := "200-299"
		if expected, actual := want, r.String(); expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("string with min matching max", func(t *testing.T) {
		r := requests.Range{
			Min: 200,
			Max: 200,
		}
		want := "200"
		if expected, actual := want, r.String(); expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
}
