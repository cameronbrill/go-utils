package main

import "testing"

type test struct {
	name string
	disjointSetImpl UnionFinder
}

func TestUnionFind(t *testing.T) {
	var tests = []test{
		{name: "DisjointSet", disjointSetImpl: &DisjointSet{}},
		{name: "QuickFind", disjointSetImpl: &QuickFind{}},
		{name: "QuickUnion", disjointSetImpl: &QuickUnion{}},
		{name: "UnionByRank", disjointSetImpl: &UnionByRank{}},
		{name: "PathCompression", disjointSetImpl: &PathCompression{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			qf := tc.disjointSetImpl
			qf.New(10)
		
			qf.Union(1, 2)
			qf.Union(2, 5)
			qf.Union(5, 6)
			qf.Union(6, 7)
			qf.Union(3, 8)
			qf.Union(8, 9)

			conn, err := qf.Connected(1, 5)
			if err != nil {
				t.Fatal(err)
			}
			if !conn {
				t.Fatal("1 is not connected to 5")
			}

			conn, err = qf.Connected(5, 7)
			if err != nil {
				t.Fatal(err)
			}
			if !conn {
				t.Fatal("5 is not connected to 7")
			}

			conn, err = qf.Connected(4, 9)
			if err != nil {
				t.Fatal(err)
			}
			if conn {
				t.Fatal("4 is connected to 9")
			}


			qf.Union(9, 4)

			conn, err = qf.Connected(4, 9)
			if err != nil {
				t.Fatal(err)
			}
			if !conn {
				t.Fatal("4 is not connected to 9")
			}
		})
	}
}