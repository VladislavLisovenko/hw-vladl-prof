package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
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

	reader := io.LimitReader(srcFile, available)
	bar := pb.Full.Start64(available)
	barReader := bar.NewProxyReader(reader)
	_, err = io.CopyN(dstFile, barReader, available)
	if err != nil {
		return err
	}
	bar.Finish()

	return nil
}
