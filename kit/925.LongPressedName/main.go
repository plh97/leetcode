package isLongPressedName

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	// vtkgn
	// vttkgnn
	i := 0
	// i, i+1
	// j, j+1

	// 只需要想,什么时候 i不变就好
	// len(typed) 不是最后一个 &&
	// typed[j] == typed[j+1]
	// 的同时, name[i] != name[i+1]

	for j := 0; j < len(typed); {
		if len(typed)-1 != j &&
			len(name)-1 != i &&
			typed[j] == typed[j+1] &&
			name[i] != name[i+1] {
		} else if len(typed)-1 != j &&
			len(name)-1 == i &&
			typed[j] == typed[j+1] {
		} else {
			i++
		}
		j++
	}

	return i == len(name)
}
