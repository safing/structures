package dsd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeFirst16(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		"47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|",
		safeFirst16Bytes([]byte("Go is an open source programming language.")),
	)
	assert.Equal(t,
		"47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|",
		safeFirst16Chars("Go is an open source programming language."),
	)

	assert.Equal(t,
		"<empty>",
		safeFirst16Bytes(nil),
	)
	assert.Equal(t,
		"<empty>",
		safeFirst16Chars(""),
	)
}
