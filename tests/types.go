package tests

import (
	"testing"

	"github.com/keks/expath"
	"github.com/keks/expath/naive"
)

type TestCase struct {
	Name     string
	FollowPairs    []naive.StringNode
	BlockPairs    []naive.StringNode
	Src, Dst naive.StringNode
	MaxDist  int

	ExpReachable bool
}

func (tc TestCase) RunTest(t *testing.T, nws expath.NewWalkState) {
	if len(tc.FollowPairs)%2 != 0 {
		t.Fatal("invalid test case: field `FollowPairs` needs to be of even length")
	}

	if len(tc.BlockPairs)%2 != 0 {
		t.Fatal("invalid test case: field `BlockPairs` needs to be of even length")
	}

	var (
		follows = naive.MakeGraph(tc.FollowPairs...)
		blocks  = naive.MakeGraph(tc.BlockPairs...)
	)

	actualReachable := expath.Do(follows, blocks, tc.Src, tc.Dst, tc.MaxDist, nws).Reachable()
	if tc.ExpReachable != actualReachable {
		t.Fatalf("expected %v, got %v", tc.ExpReachable, actualReachable)
	}
}

type TestCases []TestCase

func (tcs TestCases) RunTest(t *testing.T, nws expath.NewWalkState) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) { tc.RunTest(t, nws) })
	}
}

