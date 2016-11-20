package point

// ByTop type for sorting points top results
type ByTop []*Point

// ByBottom type for sorting points by bottom results
type ByBottom []*Point

func (s ByTop) Len() int {
	return len(s)
}

func (s ByTop) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByTop) Less(i, j int) bool {
	return s[i].Distance <= s[j].Distance
}

func (s ByBottom) Len() int {
	return len(s)
}

func (s ByBottom) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByBottom) Less(i, j int) bool {
	return s[i].Distance >= s[j].Distance
}
