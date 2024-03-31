package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	stat, err := srcFile.Stat()
	if err != nil {
		return err
	}
	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}
	available := stat.Size()
	if offset > 0 {
		available -= offset
	}
	if limit > 0 {
		if limit < available {
			available = limit
		}
	}

	dstFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if offset > 0 {
		_, err := srcFile.Seek(offset, 0)
		if err != nil {
			return err
		}
	}

	totalRead := 0
	bufSize := 1024
	buf := make([]byte, bufSize)

	for {
		nRead := 0
		if limit > 0 && int64(totalRead+bufSize) > limit {
			nRead, err = srcFile.Read(buf[:limit-int64(totalRead)])
		} else {
			nRead, err = srcFile.Read(buf)
		}
		totalRead += nRead
		if nRead > 0 {
			_, err := dstFile.Write(buf[:nRead])
			if err != nil {
				return err
			}
			fmt.Printf("Copied %d from %d\n", totalRead, available)
		}
		if limit > 0 && int64(totalRead) == limit {
			break
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}

	return nil
}
