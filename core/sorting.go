package core

func SortNodes(nodes []FileNode, sortByTime bool, reverse bool) {
	if len(nodes) < 2 {
		return
	}

	if sortByTime {
		TimeSort(nodes)
	} else {
		AlphaSort(nodes)
	}

	if reverse {
		ReverseSlice(nodes)
	}
}

func AlphaSort(nodes []FileNode) {
	for i := 0; i < len(nodes)-1; i++ {
		for j := 0; j < len(nodes)-i-1; j++ {
			if compareCaseInsensitive(nodes[j].Name, nodes[j+1].Name) {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}
}

func compareCaseInsensitive(s1, s2 string) bool {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		c1 := s1[i]
		c2 := s2[i]

		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 32
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 32
		}

		if c1 != c2 {
			return c1 > c2
		}
	}
	return len(s1) > len(s2)
}

func TimeSort(nodes []FileNode) {
	for i := 0; i < len(nodes)-1; i++ {
		for j := 0; j < len(nodes)-i-1; j++ {

			shouldSwap := false

			if nodes[j].ModeTime.Before(nodes[j+1].ModeTime) {
				shouldSwap = true
			} else if nodes[j].ModeTime.Equal(nodes[j+1].ModeTime) {
				if nodes[j].Name > nodes[j+1].Name {
					shouldSwap = true
				}
			}

			if shouldSwap {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}
}

func ReverseSlice(nodes []FileNode) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}
