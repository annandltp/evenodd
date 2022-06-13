package evenodd

func CekGanjilGenap(num ...int) string {

	result := ""
	coma := ","
	last := num[len(num)-1]

	for _, i := range num {
		if last == i {
			coma = ""
		}
		if i%2 == 0 {
			result = "genap" + coma + " "
		} else {
			result = "ganjil" + coma + " "
		}
	}
	return result
}
