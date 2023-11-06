package strategy

import "sort"

type SortStrategy interface {
	Sort(x []int)
}

type OldSort struct{}

func (o *OldSort) Sort(x []int) {
	sort.Ints(x)
}

type InsertionSort struct{}

func (i *InsertionSort) Sort(x []int) {
	size := len(x)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = x[i]
		for j = i - 1; j >= 0; j-- {
			if x[j] < buff {
				break
			}
			x[j+1] = x[j]
		}
		x[j+1] = buff
	}
}

type Context struct {
	strategy SortStrategy
}

func (c *Context) Algorithm(a SortStrategy) {
	c.strategy = a
}

func (c *Context) Sort(x []int) {
	c.strategy.Sort(x)
}
