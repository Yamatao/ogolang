package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("case", func(t *testing.T) {
		_, err := RunCmd([]string{}, Environment{})
		require.NotNil(t, err)

		_, err = RunCmd([]string{""}, Environment{})
		require.NotNil(t, err)

		_, err = RunCmd([]string{"/adsf5234_wr"}, Environment{})
		require.NotNil(t, err)

		_, err = RunCmd([]string{"/usr/bin/sh --help"}, Environment{})
		require.NotNil(t, err)

		// check exit code
		ret, err := RunCmd([]string{"true"}, Environment{})
		require.Nil(t, err)
		require.Equal(t, ret, 0)

		ret, err = RunCmd([]string{"false"}, Environment{})
		require.Nil(t, err)
		require.Equal(t, ret, 1)
	})
}
