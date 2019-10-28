package gobrick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErr(t *testing.T) {
	t.Run("CheckErr()", func(t *testing.T) {
		CheckErr(nil)

		// Shouldn't fatal out if no error is passed in.
		assert.True(t, true)
	})
}
