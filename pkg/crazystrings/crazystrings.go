package crazystrings

import (
	"errors"
	"strings"
)

type CrazyString interface {
	Scramble() string
}

type crazystring string
type bigcrazystring string

func NewCrazyString(data string) (crazystring, error) {
	if len(data) == 0 {
		return "", errors.New("string empty")
	}
	return crazystring(data), nil
}

func NewBigCrazyString(data string) (bigcrazystring, error) {
	// call super constructor
	cs, err := NewCrazyString(data)

	// cast to own type
	return bigcrazystring(cs), err
}

func (c crazystring) Scramble() string {
	newStr1 := strings.ReplaceAll(string(c), "a", "@")
	newStr2 := strings.ReplaceAll(newStr1, "e", "3")
	newStr3 := strings.ReplaceAll(newStr2, "i", "1")
	newStr4 := strings.ReplaceAll(newStr3, "o", "0")
	return newStr4
}

func (c bigcrazystring) Scramble() string {
	crazystring := crazystring(c).Scramble()

	return strings.ToUpper(crazystring)
}
