package golanglearn

import "fmt"

func mergesort(arr []int, l, r int, ch chan int) {
	if r-l <= 0 {
		return
	}

	if r-l == 1 {
		ch <- arr[l]
		return
	}

	lch := make(chan int)
	defer close(lch)
	rch := make(chan int)
	defer close(rch)

	mid := (r + l) / 2
	go mergesort(arr, l, mid, lch)
	go mergesort(arr, mid, r, rch)

	pl := l
	pr := mid
	p := l

	tmp := make([]int, r-l)

	lvalue := <-lch
	rvalue := <-rch

	for p < r {
		if pl == mid {
			tmp[p-l] = rvalue
			pr++
			if pr != r {
				rvalue = <-rch
			}
			p++
			continue
		}
		if pr == r {
			tmp[p-l] = lvalue
			pl++
			if pl != mid {
				lvalue = <-lch
			}
			p++
			continue
		}
		if lvalue > rvalue {
			tmp[p-l] = rvalue
			pr++
			if pr != r {
				rvalue = <-rch
			}
			p++
		} else {
			tmp[p-l] = lvalue
			pl++
			if pl != mid {
				lvalue = <-lch
			}
			p++
		}
	}

	for i := l; i < r; i++ {
		arr[i] = tmp[i-l]
		ch <- arr[i]
	}
}
