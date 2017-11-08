package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type float64AnyMatcher struct{}

func (float64AnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(float64)
	return ok
}

func (float64AnyMatcher) String() string {
	return "is float64"
}

// MatchAnyFloat64 checks to see if the value it's trying to match is a float64
func MatchAnyFloat64() gomock.Matcher { return float64AnyMatcher{} }

type float64Matcher struct {
	value float64
}

func (m float64Matcher) Matches(x interface{}) bool {
	if v, ok := x.(float64); ok {
		return v == m.value
	}
	return false
}

func (m float64Matcher) String() string {
	return fmt.Sprintf("is float64 %f", m.value)
}

// MatchFloat64 checks to see if the value is the value float64 value
func MatchFloat64(v float64) gomock.Matcher { return float64Matcher{v} }
