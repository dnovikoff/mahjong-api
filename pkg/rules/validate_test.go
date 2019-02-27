package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	for _, v := range All() {
		t.Run(v.GetId(), func(t *testing.T) {
			assert.NoError(t, Validate(v))
		})
	}
}
