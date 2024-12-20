package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (rt Runtime) MarshalJSON() ([]byte, error) {
	msg := fmt.Sprintf("%d mins", rt)

	quoted := strconv.Quote(msg)
	return []byte(quoted), nil
}
