package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
        {
            input:    "l0o0l0",
            expected: "",
        },
		{
			input:    "aaa0b",
			expected: "aab",
		},
        {
            input:    "XYZ2",
            expected: "XYZZ",
        },
        {
            input:    "z9z",
            expected: "zzzzzzzzzz",
        },
        {
            input:    "Ф3Р",
            expected: "ФФФР",
        },
        {
            input:    `\3`,
            expected: `\\\`,
        },
        {
            input:    "\n5",
            expected: "\n\n\n\n\n",
        },
	} {
        for i := 0; i < 1000; i++ {
            result, err := Unpack(tst.input)
            require.Equal(t, tst.err, err)
            require.Equal(t, tst.expected, result)
        }
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
