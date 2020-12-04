package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
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
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	srcStats, _ := srcFile.Stat()
	regularFile := srcStats.Mode().IsRegular()

	if !regularFile && limit == 0 {
		return errors.New("limit should be >0 for irregular files")
	}

	var srcSize int64 = srcStats.Size()
	if offset > srcSize {
		return fmt.Errorf("offset %d exceeds file size %d", offset, srcSize)
	}

	dstFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	defer dstFile.Close()

	// some preparations
	_, err = srcFile.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("failed to seek within dest file: %w", err)
	}

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
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("error while reading source: %w", err)
		}

		ni64 := int64(n) // promote n as int to int64
		if ni64 > totalBytes-readCount {
			ni64 = totalBytes - readCount
		}

		_, err = dstFile.Write(buf[:ni64])
		if err != nil {
			return fmt.Errorf("error while writing to destination: %w", err)
		}

		readCount += ni64

		PrintProgress(float32(readCount), float32(totalBytes))
		time.Sleep(80 * time.Millisecond)
	}

	return nil
}
