package iterate

// Count numbers for a range.
func Count(arg ...int) []int {
	a := []int{}
	switch len(arg) {
	case 1:
		for n := 0; n <= arg[0]; n++ {
			a = append(a, n)
		}

	case 2:
		for n := arg[0]; n <= arg[1]; n++ {
			a = append(a, n)
		}

	case 3:
		if arg[2] < 0 {
			for n := arg[0]; n >= arg[1]; n-- {
				a = append(a, n)
			}
		} else {
			n := arg[0]
			for n <= arg[1] {
				a = append(a, n)
				n += arg[2]
			}
		}
	}

	return a
}
