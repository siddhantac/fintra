package uid

import "github.com/segmentio/ksuid"

func NewID() string {
	return ksuid.New().String()
}
