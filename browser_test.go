package gobrick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenBrowser(t *testing.T) {
	url := "http://electronicpanopticon.com"

	t.Run("OpenBrowser() error", func(t *testing.T) {
		osOffset = "boopOS"

		err := OpenBrowser(url)

		assert.Error(t, err)
		osOffset = ""
	})

	t.Run("openBrowser() linux", func(t *testing.T) {
		expectedExecCommand := ExecCommand{"xdg-open", []string{url}}

		execCommand, err := openBrowser("linux", url)

		assert.NoError(t, err)
		assert.Equal(t, expectedExecCommand, *execCommand)
	})

	t.Run("openBrowser() darwin", func(t *testing.T) {
		expectedExecCommand := ExecCommand{"open", []string{url}}

		execCommand, err := openBrowser("darwin", url)

		assert.NoError(t, err)
		assert.Equal(t, expectedExecCommand, *execCommand)
	})

	t.Run("openBrowser() windows", func(t *testing.T) {
		expectedExecCommand := ExecCommand{"rundll32", []string{"url.dll,FileProtocolHandler", url}}

		execCommand, err := openBrowser("windows", url)

		assert.NoError(t, err)
		assert.Equal(t, expectedExecCommand, *execCommand)
	})

	t.Run("openBrowser() unsupported", func(t *testing.T) {
		expectedExecCommand := ExecCommand{}

		execCommand, err := openBrowser("boopOS", url)

		assert.Error(t, err)
		assert.Equal(t, expectedExecCommand, *execCommand)
	})
}
