package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/trussle/uuid"
)

type uuidAnyMatcher struct{}

func (uuidAnyMatcher) Matches(x interface{}) bool {
	_, ok := x.(uuid.UUID)
	return ok
}

func (uuidAnyMatcher) String() string {
	return "is uuid"
}

// MatchAnyUUID checks to see if the value it's trying to match is a uuid
func MatchAnyUUID() gomock.Matcher { return uuidAnyMatcher{} }

type uuidMatcher struct {
	value uuid.UUID
}

func (m uuidMatcher) Matches(x interface{}) bool {
	if v, ok := x.(uuid.UUID); ok {
		return v.Equals(m.value)
	}
	return false
}

func (m uuidMatcher) String() string {
	return fmt.Sprintf("is uuid %s", m.value.String())
}

// MatchUUID checks to see if the value is the value uuid value
func MatchUUID(v uuid.UUID) gomock.Matcher { return uuidMatcher{v} }
