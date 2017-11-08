package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type intAnyMatcher struct{}

func (intAnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(int)
	return ok
}

func (intAnyMatcher) String() string {
	return "is int"
}

// MatchAnyInt checks to see if the value it's trying to match is a int
func MatchAnyInt() gomock.Matcher { return intAnyMatcher{} }

type intMatcher struct {
	value int
}

func (m intMatcher) Matches(x interface{}) bool {
	if v, ok := x.(int); ok {
		return v == m.value
	}
	return false
}

func (m intMatcher) String() string {
	return fmt.Sprintf("is int %d", m.value)
}

// MatchInt checks to see if the value is the value int value
func MatchInt(v int) gomock.Matcher { return intMatcher{v} }
