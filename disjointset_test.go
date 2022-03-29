package main

import "testing"

func TestUnionFind(t *testing.T) {
	t.Run("QuickFind", func(t *testing.T) {
		qf := QuickFind{}
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