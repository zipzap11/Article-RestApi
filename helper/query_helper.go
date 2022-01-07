package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func JoinIds(tuples []uuid.UUID) string {
	var str = "("
	for i, val := range tuples {
		if i == len(tuples)-1 {
			str += fmt.Sprintf("%v)", val)
			break
		}
		str += fmt.Sprintf("%v,", val)
	}

	return str
}

func DynamicInject(n int) string {
	res := "("
	for i := 1; i <= n; i++ {
		if i == 1 {
			res += fmt.Sprintf("$%v", i)
			continue
		}
		if i == (n - 1) {
			res += fmt.Sprintf(", $%v)", i)
			break
		}
		res += fmt.Sprintf(", $%v", i)
	}

	return res
}
