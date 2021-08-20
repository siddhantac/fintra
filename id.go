package main

import (
	"strconv"
	"time"
)

type IDGeneratorImpl struct{}

func (IDGeneratorImpl) NewID() string {
	return strconv.Itoa(time.Now().Nanosecond())
}
