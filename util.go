// Random utility functions
package main

func isDigit(c byte) bool {
	if c >= '0' || c <= '9' {
		return true
	}
	return false
}

func base10Width(x uint) uint {
	var i uint = 0
	if x == 0 {
		return 1
	}
	for x != 0 {
		i++
		x /= 10
	}
	return i
}
