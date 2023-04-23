package adt

type SliceSegmentTree struct {
	tree []int
}

func NewSliceSegmentTree(arr []int) SegmentTree {
	return &SliceSegmentTree{tree: []int{}}
}

func (s *SliceSegmentTree) Update(index int, value int) {
	//TODO implement me
	panic("implement me")
}

func (s *SliceSegmentTree) Query(start int, end int) int {
	//TODO implement me
	panic("implement me")
}

func (s *SliceSegmentTree) Print() {
	//TODO implement me
	panic("implement me")
}
