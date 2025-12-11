package math

// Pow raises an integer base to an integer exponent.
func Pow(base, exp int) (res int) {
	res = 1
	for ; exp > 0; exp -= 1 {
		res *= base
	}
	return
}

// Abs returns the absolute value of an integer.
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// GCD returns the greatest common factor of two integers.
func GCD(a, b int) int {
	a = Abs(a)
	b = Abs(b)
	r0 := max(a, b)
	r1 := min(a, b)
	r2 := r0 % r1
	for r2 != 0 {
		r0 = r1
		r1 = r2
		r2 = r0 % r1
	}
	return r1
}

// LCM returns the least common multiple of two integers.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Digits returns the number of digits in an integer in a given base.
func Digits(a, base int) (res int) {
	res = 1
	for ; a >= base; a /= base {
		res += 1
	}
	return
}

// IntSplit takes some numbre of digits off the end of an integer in a given base and returns both pieces.
func IntSplit(a, base, digits int) (start, end int) {
	splitter := Pow(base, digits)
	start = a / splitter
	end = a % splitter
	return
}
