package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {

	t.Run("simple case", func(t *testing.T) {
		// wrong path
		_, err := ReadDir("/some/wrong/path/")
		require.NotNil(t, err)

		// doesn't take dirs, links and files with =
		env, err := ReadDir("./testdata/notfiles")
		require.Equal(t, len(env), 0)
		require.Nil(t, err)
	})

	t.Run("values", func(t *testing.T) {
		// space/tabs are deleted at the end
		env, err := ReadDir("./testdata/env/")
		require.Nil(t, err)

		// multiline
		require.Equal(t, env["BAR"], "bar")

		// 0x0 replaced with \n
		require.Equal(t, env["FOO"], "   foo\nwith new line")

		// quotes
		require.Equal(t, env["HELLO"], "\"hello\"")
	})
}
