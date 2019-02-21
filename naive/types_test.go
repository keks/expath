package naive

import (
	"testing"

	"github.com/keks/expath"
)

func TestSliceNodes(t *testing.T) {
	type tcase struct {
		name string
		ns   *SliceNodes
		exp  []expath.Node
	}

	runner := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			for i, exp := range tc.exp {
				if !tc.ns.Next() {
					t.Fatalf("early end at i=%d", i)
				}

				n := tc.ns.Get()

				if n != exp {
					t.Fatalf("expected node %v, got %v", exp, n)
				}
			}

			if tc.ns.Next() {
				t.Fatal("Next returned true but should be finished")
			}
		}
	}

	tcs := []tcase{
		{
			name: "oneNode",
			ns:   NewSliceNodes(StringNode("a node")).(*SliceNodes),
			exp:  []expath.Node{StringNode("a node")},
		},
		{
			name: "empty",
			ns:   NewSliceNodes().(*SliceNodes),
			exp:  []expath.Node{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, runner(tc))
	}
}
