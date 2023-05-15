package Day2

import (
	"fmt"
	"testing"
)

func TestLonelyInteger(t *testing.T) {
	a := []int32{1, 2, 3, 2, 1, 5, 5}
	fmt.Println(lonelyInteger(a))
}
