package main

import (
    "io/ioutil"
    "testing"
    "os"
    "github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
    path := "/tmp/otus_hw07_test.txt"
    text := []byte("abcdef")
    err := ioutil.WriteFile(path, text, 0444)
    require.Empty(t, err)
    defer os.Remove(path)

    dstPath := path + ".1"

    t.Run("errors", func (t *testing.T) {
        err = Copy("wrong path", "", 0, 0)
        require.Error(t, err)

        err = Copy(path, "", 0, 0)
        require.Error(t, err)

        err = Copy("", dstPath, 0, 0)
        require.Error(t, err)

        err = Copy(path, dstPath, -1, 0)
        require.Error(t, err)

        err = Copy(path, dstPath, 0, -1)
        require.Error(t, err)
    })

    t.Run("positive cases", func (t *testing.T) {
        err := Copy(path, path + ".1", 0, 0)
        require.Empty(t, err)
        ds, _ := os.Stat(dstPath)
        require.Equal(t, len(text), int(ds.Size()))

        limit := int64(2)
        err = Copy(path, path + ".1", 0, limit)
        require.Empty(t, err)
        ds, _ = os.Stat(dstPath)
        require.Equal(t, limit, ds.Size())

        offset := int64(2)
        err = Copy(path, path + ".1", offset, 0)
        require.Empty(t, err)
        ds, _ = os.Stat(dstPath)
        require.Equal(t, int64(len(text)) - offset, ds.Size())
    })
}
