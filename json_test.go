package gobrick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	t.Run("ToJsonBytes()", func(t *testing.T) {
		var expected []byte = nil
		assert.Equal(t, expected, ToJsonBytes(nil))
	})

	t.Run("ToJson", func(t *testing.T) {
		assert.Equal(t, "{}", ToJson(nil))
		assert.Equal(t, "{}", ToJson(""))
		//  TODO: get me to work: assert.Equal(t, "{}", ToJson("  "))
	})
}
