package Day1

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	a := strings.Split(s, ":")
	b := a[2]
	if strings.ToLower(b[len(b)-2:]) == "am" {
		if a[0] == "12" {
			a[0] = "00"
		}
	} else {
		if a[0] != "12" {
			num, _ := strconv.Atoi(a[0])
			a[0] = strconv.Itoa(num + 12)
		}
	}
	return fmt.Sprintf("%s:%s:%s", a[0], a[1], b[:2])
}
