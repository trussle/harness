package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type int32AnyMatcher struct{}

func (int32AnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(int32)
	return ok
}

func (int32AnyMatcher) String() string {
	return "is int32"
}

// MatchAnyInt32 checks to see if the value it's trying to match is a int32
func MatchAnyInt32() gomock.Matcher { return int32AnyMatcher{} }

type int32Matcher struct {
	value int32
}

func (m int32Matcher) Matches(x interface{}) bool {
	if v, ok := x.(int32); ok {
		return v == m.value
	}
	return false
}

func (m int32Matcher) String() string {
	return fmt.Sprintf("is int32 %d", m.value)
}

// MatchInt32 checks to see if the value is the value int32 value
func MatchInt32(v int32) gomock.Matcher { return int32Matcher{v} }
