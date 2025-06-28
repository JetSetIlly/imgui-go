package imgui_test

import (
	"testing"

	"github.com/jetsetilly/imgui-go/v5"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	version := imgui.Version()
	assert.Equal(t, "1.91.9b", version)
}
