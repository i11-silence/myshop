package service

func DeleteInSlice(a string, list []string) []string {
	var result []string
	for _, b := range list {
		if b != a {
			result = append(result, b)
		}
	}
	return result
}
