package tools

func SlicesContains(s []string, v string) bool {
	for _, vs := range s {
		if vs == v {
			return true
		}
	}
	return false
}
