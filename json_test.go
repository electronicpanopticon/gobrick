package gobrick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	t.Run("ToJSONBytes()", func(t *testing.T) {
		var expected []byte = nil
		assert.Equal(t, expected, ToJSONBytes(nil))
	})

	t.Run("ToJSON", func(t *testing.T) {
		assert.Equal(t, "{}", ToJSON(nil))
		assert.Equal(t, "{}", ToJSON(""))
		//  TODO: get me to work: assert.Equal(t, "{}", ToJSON("  "))
	})
}
