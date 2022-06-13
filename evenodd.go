package evenodd

func CekGanjilGenap(i int) string {

	result := ""

	if i%2 == 0 {
		result = "genap"
	} else {
		result = "ganjil"
	}

	return result
}
