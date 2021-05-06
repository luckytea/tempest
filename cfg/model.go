package cfg

import "errors"

type Config struct {
	Type  string
	Name  string
	Desc  string
	Label string
	From  int64
	To    int64
}

var ErrMalformedTime = errors.New("malformed time: to can't be less or equal to from")
