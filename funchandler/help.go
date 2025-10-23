package funchandler

func CopySlice(group [][][]string) [][][]string {
	copyGroup := make([][][]string, len(group))
	for i := range group {
		copyGroup[i] = make([][]string, len(group[i]))[:0] // allocate zero-length slices with capacity
		for _, p := range group[i] {
			cp := make([]string, len(p))
			copy(cp, p)
			copyGroup[i] = append(copyGroup[i], cp)
		}
	}
	return copyGroup
}