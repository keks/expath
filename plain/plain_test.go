package plain

import (
	"testing"

	"github.com/keks/expath/tests"
)

func TestPlain(t *testing.T) {
	t.Run("General/Simple", func(t *testing.T) { tests.SimpleTests.RunTest(t, New) })
}
