package main

import (
    "os"
    "io"
    "fmt"
    "time"
    "errors"
)

const (
    bufSize = 256
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
    if offset < 0 {
        return fmt.Errorf("wrong offset (%d), should be >= 0", offset)
    }
    if limit < 0 {
        return fmt.Errorf("wrong limit (%d), should be >= 0", limit)
    }

    srcFile, err := os.Open(fromPath)
    if err != nil {
        return fmt.Errorf("failed to open source file: %v", err)
    }
    defer srcFile.Close()

    srcStats, _ := srcFile.Stat()
    regularFile := srcStats.Mode().IsRegular()

    if !regularFile && limit == 0 {
        return errors.New("limit should be >0 for irregular files")
    }

    var srcSize int64 = 0
    if regularFile {
        // get the file size
        srcSize, err = srcFile.Seek(0, 2) // at the ending
        if err != nil {
            return fmt.Errorf("failed to seek in source file: %v", err)
        }
        if offset > srcSize {
            return fmt.Errorf("offset %d exceeds file size %d", offset, srcSize)
        }
        // seek to the beginning of the source file
        _, err = srcFile.Seek(offset, 0)
        if err != nil {
            return fmt.Errorf("failed to seek back in source file: %v", err)
        }
    }

    dstFile, err := os.Create(toPath)
    if err != nil {
        return fmt.Errorf("failed to open destination file: %v", err)
    }
    defer dstFile.Close()

    // some preparations
    if limit == 0 {
        limit = srcSize - offset
    }
    totalBytes := limit
    if regularFile && totalBytes > srcSize {
        totalBytes = srcSize
    }
    var readCount int64 = 0
    buf := make([]byte, bufSize)

    // read by small chunks
    for readCount < totalBytes {
        n, err := srcFile.Read(buf)
        if err == io.EOF {
            break
        }
        if err != nil {
            return fmt.Errorf("error while reading source: %v", err)
        }

        n_ := int64(n) // promote n as int to int64
        if n_ > totalBytes - readCount {
            n_ = totalBytes - readCount
        }

        _, err = dstFile.Write(buf[:n_])
        if err != nil {
            return fmt.Errorf("error while writing to destination: %v", err)
        }

        readCount += n_

        PrintProgress(float32(readCount), float32(totalBytes))
        time.Sleep(80 * time.Millisecond)
    }

	return nil
}
