package gobrick

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestPath(t *testing.T) {
	t.Run("FilePathInGoPath()", func(t *testing.T) {
		gopath := GetGOPATH()
		relPath := "src/github.com/electronicpanopticon/gobrick/data/foo.txt"
		expectedPath := fmt.Sprintf("%s/%s", gopath, relPath)

		path := FilePathInGoPath(relPath)

		assert.Equal(t, expectedPath, path)
	})

	t.Run("FileInGoPath() no such file", func(t *testing.T) {
		_, err := FileInGoPath("myxlplyx")

		assert.Error(t, err)
	})

	t.Run("FileInGoPath()", func(t *testing.T) {
		relPath := "src/github.com/electronicpanopticon/gobrick/data/foo.txt"

		f, err := FileInGoPath(relPath)
		b := make([]byte, 3)
		io.ReadAtLeast(f, b, 3)

		assert.NoError(t, err)
		assert.Equal(t, "bar", string(b))
	})
}
