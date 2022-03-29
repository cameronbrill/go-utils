package main

import "fmt"

type UnionFinder interface {
	New(size int)
	Union(x, y int)  error
	Find(x int) (int, error)
	Connected(x, y int) (bool, error)
}


type QuickFind struct {
	root []int
	UnionFinder
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

type QuickUnion struct {
	root []int
	UnionFinder
}

func (q *QuickUnion) New(size int) {
	q.root = new(size)
}

func (q *QuickUnion) Find(x int) (int, error) {
	if x < 0 || x >= len(q.root) {
		return -1, fmt.Errorf("index out of bounds")
	}

	for x != q.root[x] {
		x = q.root[x]
	}
	return x, nil
}

func (q *QuickUnion) Union(x, y int) error {
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

	q.root[yr] = xr

	return nil
}

func (q *QuickUnion) Connected(x, y int) (bool, error) {
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

type UnionByRank struct {
	root []int
	rank []int
	UnionFinder
}

func (q *UnionByRank) New(size int) {
	q.root, q.rank = newRank(size)
}

func (q *UnionByRank) Find(x int) (int, error) {
	if x < 0 || x >= len(q.root) {
		return -1, fmt.Errorf("index out of bounds")
	}

	for x != q.root[x] {
		x = q.root[x]
	}
	return x, nil
}

func (q *UnionByRank) Union(x, y int) error {
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

	if q.rank[xr] > q.rank[yr] {
		q.root[yr] = xr
	} else if q.rank[xr] < q.rank[yr] {
		q.root[xr] = yr
	} else {
		q.root[yr] = xr
		q.rank[xr] += 1
	}

	return nil
}

func (q *UnionByRank) Connected(x, y int) (bool, error) {
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

type PathCompression struct {
	root []int
	UnionFinder
}

func (q *PathCompression) New(size int) {
	q.root = new(size)
}

func (q *PathCompression) Find(x int) (int, error) {
	if x < 0 || x >= len(q.root) {
		return -1, fmt.Errorf("index out of bounds")
	}

	if x == q.root[x] {
		return x, nil
	}

	var err error
	q.root[x], err = q.Find(q.root[x])
	if err != nil {
		return -1, err
	}

	return q.root[x], nil
}

func (q *PathCompression) Union(x, y int) error {
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

	q.root[yr] = xr

	return nil
}

func (q *PathCompression) Connected(x, y int) (bool, error) {
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

func new(size int) []int {
	t := make([]int, size)
	for i := range t {
		t[i] = i
	}
	return t
}

func newRank(size int) ([]int, []int) {
	rt := make([]int, size)
	rk := make([]int, size)
	for i := range rt {
		rt[i] = i
		rk[i] = 1
	}
	return rt, rk
}
