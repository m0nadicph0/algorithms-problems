package adt

type SegmentTree interface {
	Update(index int, value int)
	Query(start int, end int) int
	Print()
}
