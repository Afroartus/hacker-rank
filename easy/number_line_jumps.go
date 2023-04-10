package easy

func kangoroo(x1, v1, x2, v2 int32) (string, int32, int32) {
	bandera := true
	resp := "NO"
	if v2 > v1 && x2 > x1 {
		return "NO", x1, x2
	}
	for bandera == true || x2 <= 10000 {
		x2 += v2
		x1 += v1
		if x1 == x2 {
			bandera = false
			return "YES", x1, x2
		}
	}
	return resp, x1, x2
}
