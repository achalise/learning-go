package reverse_integer

func reverse(y int) int  {
	MAX := 1 <<31 - 1
	MIN := -MAX -1
	x := y
	if (y < 0) {
		x = -x
	}
	res := 0
	for x > 0 {
		tail := x % 10
		res = 10 * res + tail
		if res > MAX || -res < MIN {
			return 0
		}
		x = x / 10
	}
	if y < 0 {
		res = -res
	}
	return res
}
