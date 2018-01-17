package main

import (
	"fmt"
)

/*
	f[i][j] = k:
*/

type voucher struct {
	l, r, c int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	data := []*voucher{}
	for i := 0; i < n; i++ {
		temp := &voucher{}
		fmt.Scan(&temp.l, &temp.r, &temp.c)
		data = append(data)
	}

	lengthMap := make(map[int][]*voucher)

	for i := 0; i < n; i++ {
		length := data[i].r - data[i].l + 1

		if lengthMap[length] == nil {
			lengthMap[length] = make([]*voucher, 0)
		}

		lengthMap[length] = append(lengthMap[length], data[i])
	}

	for i := 2; i <= m/2+1; i++ {
		if lengthMap[i] == nil || lengthMap[m-i] == nil {
			continue
		}

	}

}
