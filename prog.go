package main

func main() {
	println("hello")

	print(eq(map[string]int{"zb": 111}, map[string]int{"zb": 111}))
}

func eq(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}

	return true
}
