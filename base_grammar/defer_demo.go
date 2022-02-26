package base_grammar

// 1
func deferReturn1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

// 2
func deferReturn2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

// 1
func deferReturn3() (b int) {
	a := 1
	defer func() {
		a++
	}()
	return 1
}

// 1?
func deferReturn4() (a int) {
	defer func(a int) {
		a++
	}(a)
	return 1
}
