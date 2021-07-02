package iterate

// Count numbers for a range.
// One argument returns a slice of numbers from 0 up (and ony up) to that number.
// Two arguments return a slice ranging from the first to the second number (both inclusive, couunting up).
// Three arhguments allow for a positive or negative step size, which lets you count up and down and
// get non-consecutive ranges.
//
// Use in a for loop like this:
//	for x := range iterate.Count(10)
//	...
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
