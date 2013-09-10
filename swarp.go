package swrap

type SWrap []byte

func (sw *SWrap) Len() int {
	return len(*sw)
}

func (sw *SWrap) Add(a byte) {
	*sw = append(*sw, a)
}

func (sw *SWrap) Merge(a []byte) {
	*sw = append(*sw, a...)
}

func (sw *SWrap) Delete(i int) {
	s := *sw
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	*sw = s[:len(s)-1]
}

func (sw SWrap) Compare(b []byte) bool {
	if len(sw) != len(b) {
		return false
	}

	for i, v := range b {
		if sw[i] != v {
			return false
		}
	}
	return true
}

func (sw *SWrap) Push(b byte) {
	sw.Add(b)
}

func (sw *SWrap) Pop() byte {
	s := *sw
	last := s[len(s)-1]
	s[len(s)-1] = 0
	*sw = s[:len(s)-1]
	return last
}

func (sw *SWrap) Shift(b byte) {
	s := *sw
	s = append(s, 0)
	copy(s[1:], s[:])
	s[0] = b
	*sw = s
}

func (sw *SWrap) UnShift() byte {
	s := *sw
	top := s[0]
	s[len(s)-1] = 0
	*sw = s[1:]
	return top
}