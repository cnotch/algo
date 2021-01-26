// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package radix

import "math"

var decDivs [19]int

func init() {
	for i := 0; i < len(decDivs); i++ {
		decDivs[i] = int(math.Pow10(i))
	}
}

// MaxDigits 集合中元素最大位数
func MaxDigits(n int, digits func(i int) int) int {
	max := 0
	for i := 0; i < n; i++ {
		l := digits(i)
		for l > max {
			max = l
		}
	}
	return max
}

// MaxDecDigits 十进制整数集合中元素最大位数
func MaxDecDigits(n int, index func(i int) int) int {
	max := 1
	divisor := 10
	for i := 0; i < n; i++ {
		for index(i) >= divisor {
			divisor *= 10
			max++
		}
	}
	return max
}

// DecDigit 十进制整数指定位的值
func DecDigit(num int, p int) int {
	return (num / decDivs[p]) % 10
}

// CharDigit 字串指定位置的字符值，字串最大长度为 d
func CharDigit(s string, d, p int) int {
	i := d - p - 1
	if i < len(s) {
		return int(s[i])
	}
	return 0
}
