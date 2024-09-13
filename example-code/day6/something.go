package main

func did_something_happened() string {
	return "nothing happened"
}
func main() {
	if val := did_something_happened(); val != "Nothing Happened" {
		println("SOMETHING HAPPENED")
	} else {
		println("nothing")
	}
	// val 無法在它被宣告的if-else區塊外面使用
}
