package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type int64AnyMatcher struct{}

func (int64AnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(int64)
	return ok
}

func (int64AnyMatcher) String() string {
	return "is int64"
}

// MatchAnyInt64 checks to see if the value it's trying to match is a int64
func MatchAnyInt64() gomock.Matcher { return int64AnyMatcher{} }

type int64Matcher struct {
	value int64
}

func (m int64Matcher) Matches(x interface{}) bool {
	if v, ok := x.(int64); ok {
		return v == m.value
	}
	return false
}

func (m int64Matcher) String() string {
	return fmt.Sprintf("is int64 %d", m.value)
}

// MatchInt64 checks to see if the value is the value int64 value
func MatchInt64(v int64) gomock.Matcher { return int64Matcher{v} }
