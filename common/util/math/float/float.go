package float

import (
	"fmt"
	"strconv"
)

func Decimal(value float64, num ...int) float64 {
	dec := 5
	if len(num) != 0 {
		dec = num[0]
	}
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(dec)+"f", value), 64)
	return value
}
