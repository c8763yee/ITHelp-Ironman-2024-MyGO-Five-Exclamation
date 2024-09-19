package main

type Counter func()
type Setter func(int)
type Getter func() int

func closureCounter() (counter Counter, adder, setter Setter, getter Getter) {
	count := 0

	// assign function as value to variable
	counter = func() {
		count += 1
	}
	adder = func(val int) {
		count += val
	}
	setter = func(num int) {
		count = num
	}
	getter = func() int {
		return count
	}

	return counter, adder, setter, getter
}

func main() {
	cCounter, cAdder, cSetter, cGetter := closureCounter()
	// counter process
	cCounter()
	cCounter()
	cCounter()
	cCounter()
	cAdder(100)
	println(cGetter()) // 104
	cSetter(10)
	println(cGetter()) // 10
}
