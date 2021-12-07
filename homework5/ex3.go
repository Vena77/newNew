func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//в третей домашке применила эту фуекцию для вычисления факториала
