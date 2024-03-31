package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr       string
	limit       int64
	offset      int64
	expected    int64
	mustBeError bool
} {
	return []struct {
		descr       string
		limit       int64
		offset      int64
		expected    int64
		mustBeError bool
	}{
		{
			descr:       "limit 0, offset 0, filesize is 6742",
			limit:       0,
			offset:      0,
			expected:    6742,
			mustBeError: false,
		},
		{
			descr:       "limit 100, offset 0, filesize is 100",
			limit:       100,
			offset:      0,
			expected:    100,
			mustBeError: false,
		},
		{
			descr:       "limit 10000, offset 0, filesize is 6742",
			limit:       10000,
			offset:      0,
			expected:    6742,
			mustBeError: false,
		},
		{
			descr:       "limit 100, offset 1000, filesize is 100",
			limit:       100,
			offset:      1000,
			expected:    100,
			mustBeError: false,
		},
		{
			descr:       "limit 1000, offset 10000, must be error 'offset exceeds file size'",
			limit:       100,
			offset:      10000,
			expected:    0,
			mustBeError: true,
		},
		{
			descr:       "limit 0, offset 1000, filesize is 5742",
			limit:       0,
			offset:      1000,
			expected:    5742,
			mustBeError: false,
		},
	}
}

func TestCopy(t *testing.T) {
	srcFileName := "testdata/input.txt"
	dstFileName := "testdata/testoutput.txt"
	for _, tc := range tests() {
		tc := tc
		t.Run(tc.descr, func(t *testing.T) {
			err := Copy(srcFileName, dstFileName, tc.offset, tc.limit)
			if tc.mustBeError {
				require.Error(t, ErrOffsetExceedsFileSize, err)
				return
			}
			dstFile, _ := os.Open(dstFileName)
			defer dstFile.Close()
			stat, _ := dstFile.Stat()
			require.Equal(t, tc.expected, stat.Size())
		})
	}
	err := os.Remove(dstFileName)
	if err != nil {
		fmt.Println(err)
	}
}
