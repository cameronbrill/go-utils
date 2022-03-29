package main

import "fmt"

type DisjointSet interface {
	Union(x, y int)  error
	Find(x int) (int, error)
	Connected(x, y int) (bool, error)
}

func new(size int) []int {
	t := make([]int, size)
	for i := range t {
		t[i] = i
	}
	return t
}

type QuickFind struct {
	root []int
	DisjointSet
}

func (q *QuickFind) New(size int) {
	q.root = new(size)
}

func (q *QuickFind) Find(x int) (int, error) {
	if x < 0 || x >= len(q.root) {
		return -1, fmt.Errorf("index out of bounds")
	}

	return q.root[x], nil
}

func (q *QuickFind) Union(x, y int)  error {
	xr, err := q.Find(x)
	if err != nil {
		return err
	}

	yr, err := q.Find(y)
	if err != nil {
		return err
	}

	if xr == yr {
		return fmt.Errorf("cannot union vertices that share a root")
	}

	for i := range q.root {
		if q.root[i] == yr {
			q.root[i] = xr
		}
	}

	return nil
}

func (q *QuickFind) Connected(x, y int) (bool, error) {
	xp, err := q.Find(x)
	if err != nil {
		return false, err
	}
	yp, err := q.Find(y)
	if err != nil {
		return false, err
	}

	return xp == yp, nil
}