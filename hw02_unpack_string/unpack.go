package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
    "strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
    var result strings.Builder
    lastChar := '0'
    lastIsNumber := true

    // make sure it has enough capacity for 125% of original size
    result.Grow(int(float32(len(s)) * 1.25))

    for _, c := range s {
        //if unicode.IsNumber(c) {
        if c >= '0' && c <= '9' {
            if lastIsNumber {
                return "", ErrInvalidString
            }
            if c > '0' {
                //result.WriteString(strings.Repeat(string(lastChar), int(c - '0')));
                n := int(c - '0')
                for i := 0; i < n; i++ {
                    result.WriteRune(lastChar)
                }
            }
            lastIsNumber = true
        } else {
            if !lastIsNumber {
                result.WriteRune(lastChar)
            }
            lastIsNumber = false
        }

        lastChar = c
    }

    // write trailing character
    if !lastIsNumber {
        result.WriteRune(lastChar)
    }

	return result.String(), nil
}
