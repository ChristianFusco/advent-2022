package main

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	day1()
}
