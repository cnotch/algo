// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package partition

import (
	"fmt"
	"sort"
)

// Mode 切分模式
type Mode uint8

// 切分模式常量
const (
	Basic          = Mode(iota) // 基本模式
	Randomized                  // 随机化模式
	ThreeWay                    // 三向切分模式
	StrictThreeWay              // 严格的三向切分模式
)

// Interface 标准库 sort.Interface 别名
type Interface = sort.Interface

// LessSwap 用于自动转换的函数优化版本。
// 它用于替换 Interface 接口。
type LessSwap struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}

// Partition 面向 sort.Interface 的切分函数类型
type Partition func(data Interface, lo, hi int) (midlo int, midhi int)

// ZfuncPartition 面向 LessSwap 函数型的切分函数类型
type ZfuncPartition func(data LessSwap, lo, hi int) (midlo int, midhi int)

// Get 获取指定模式的切分函数
func Get(mode Mode) Partition {
	switch mode {
	case Basic:
		return basic
	case Randomized:
		return randomized
	case ThreeWay:
		return threeWay
	case StrictThreeWay:
		return strictThreeWay
	default:
		panic(fmt.Sprintf("Not supported mode: %d", mode))
	}
}

// GetZfunc 获取指定模式的 zfunc 版的切分函数
func GetZfunc(mode Mode) ZfuncPartition {
	switch mode {
	case Basic:
		return basicZfunc
	case Randomized:
		return randomizedZfunc
	case ThreeWay:
		return threeWayZfunc
	case StrictThreeWay:
		return strictThreeWayZfunc
	default:
		panic(fmt.Sprintf("Not supported mode: %d", mode))
	}
}
