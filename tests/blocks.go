package tests

import (
	"github.com/keks/expath/naive"
)

var BlockTests = TestCases{
	TestCase{
		Name: "blockedPath",
		FollowPairs: []naive.StringNode{
			"alice", "bob",
			"bob", "claire",
			"bob", "carol",
			"claire", "debora",
			"debora", "eli",
		},
		BlockPairs: []naive.StringNode{
			"bob", "debora",
		},
		Src:          "alice",
		Dst:          "eli",
		MaxDist:      4,
		ExpReachable: false,
	},
}
