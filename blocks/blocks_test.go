package blocks

import (
	"testing"

	"github.com/keks/expath/tests"
)

func TestBlocks(t *testing.T) {
	t.Run("General/Simple", func(t *testing.T) {
		tests.SimpleTests.RunTest(t, New)
	})

	t.Run("General/Block", func(t *testing.T) {
		tests.BlockTests.RunTest(t, New)
	})
}
