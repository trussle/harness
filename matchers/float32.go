package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type float32AnyMatcher struct{}

func (float32AnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(float32)
	return ok
}

func (float32AnyMatcher) String() string {
	return "is float32"
}

// MatchAnyFloat32 checks to see if the value it's trying to match is a float32
func MatchAnyFloat32() gomock.Matcher { return float32AnyMatcher{} }

type float32Matcher struct {
	value float32
}

func (m float32Matcher) Matches(x interface{}) bool {
	if v, ok := x.(float32); ok {
		return v == m.value
	}
	return false
}

func (m float32Matcher) String() string {
	return fmt.Sprintf("is float32 %f", m.value)
}

// MatchFloat32 checks to see if the value is the value float32 value
func MatchFloat32(v float32) gomock.Matcher { return float32Matcher{v} }
