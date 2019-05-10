package isLongPressedName

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	i := 0
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
