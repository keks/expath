package tests

import (
	"github.com/keks/expath/naive"
)

var SimpleTests = TestCases{
	TestCase{
		Name: "longPath",
		FollowPairs: []naive.StringNode{
			"alice", "bob",
			"bob", "claire",
			"bob", "carol",
			"claire", "debora",
			"debora", "eli",
		},
		Src:          "alice",
		Dst:          "eli",
		MaxDist:      4,
		ExpReachable: true,
	},
	TestCase{
		Name: "tooShortPath",
		FollowPairs: []naive.StringNode{
			"alice", "bob",
			"bob", "claire",
			"bob", "carol",
			"claire", "debora",
			"debora", "eli",
		},
		Src:          "alice",
		Dst:          "eli",
		MaxDist:      3,
		ExpReachable: false,
	},
	TestCase{
		Name: "noPath",
		FollowPairs: []naive.StringNode{
			"alice", "bob",
			"claire", "debora",
			"debora", "eli",
		},
		Src:          "alice",
		Dst:          "eli",
		MaxDist:      10,
		ExpReachable: false,
	},
	TestCase{
		Name: "twoPath",
		FollowPairs: []naive.StringNode{
			"alice", "bob",
			"alice", "barbara",
			"bob", "claire",
			"barbara", "claire",
		},
		Src:          "alice",
		Dst:          "claire",
		MaxDist:      2,
		ExpReachable: true,
	},
}
