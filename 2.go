package main

func euler2() {
	//Project Euler Problem #2
	prevFib := 1
	currFib := 2
	sum := 0

	for currFib < 4000000 {
		if currFib%2 == 0 {
			sum += currFib
		}

		temp := currFib
		currFib = currFib + prevFib
		prevFib = temp

	}

	printAnswer(2, sum)
}
