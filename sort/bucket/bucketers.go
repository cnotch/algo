// Code generated using genbucketers.go; DO NOT EDIT.

// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bucket

import (
	"reflect"
	"unsafe"

	"github.com/cnotch/algo/reflect/iface"
)

var bucketers = [256 + 1]func(typ reflect.Type, ptr unsafe.Pointer) bucket{

	0: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]struct{})(ptr)
		bslice := reflect.Zero(typ)
		subs := []struct{}(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	1: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]int8)(ptr)
		bslice := reflect.Zero(typ)
		subs := []int8(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	2: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]int16)(ptr)
		bslice := reflect.Zero(typ)
		subs := []int16(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	3: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][3]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][3]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	4: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]int32)(ptr)
		bslice := reflect.Zero(typ)
		subs := []int32(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	5: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][5]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][5]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	6: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][6]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][6]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	7: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][7]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][7]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	8: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]int64)(ptr)
		bslice := reflect.Zero(typ)
		subs := []int64(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	9: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][9]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][9]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	10: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][10]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][10]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	11: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][11]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][11]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	12: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][12]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][12]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	13: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][13]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][13]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	14: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][14]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][14]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	15: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][15]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][15]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	16: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][16]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][16]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	17: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][17]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][17]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	18: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][18]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][18]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	19: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][19]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][19]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	20: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][20]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][20]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	21: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][21]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][21]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	22: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][22]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][22]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	23: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][23]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][23]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	24: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][24]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][24]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	25: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][25]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][25]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	26: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][26]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][26]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	27: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][27]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][27]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	28: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][28]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][28]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	29: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][29]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][29]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	30: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][30]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][30]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	31: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][31]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][31]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	32: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][32]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][32]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	33: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][33]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][33]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	34: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][34]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][34]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	35: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][35]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][35]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	36: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][36]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][36]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	37: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][37]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][37]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	38: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][38]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][38]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	39: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][39]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][39]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	40: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][40]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][40]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	41: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][41]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][41]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	42: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][42]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][42]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	43: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][43]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][43]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	44: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][44]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][44]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	45: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][45]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][45]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	46: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][46]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][46]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	47: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][47]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][47]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	48: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][48]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][48]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	49: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][49]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][49]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	50: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][50]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][50]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	51: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][51]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][51]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	52: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][52]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][52]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	53: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][53]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][53]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	54: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][54]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][54]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	55: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][55]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][55]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	56: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][56]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][56]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	57: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][57]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][57]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	58: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][58]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][58]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	59: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][59]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][59]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	60: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][60]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][60]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	61: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][61]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][61]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	62: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][62]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][62]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	63: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][63]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][63]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	64: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][64]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][64]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	65: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][65]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][65]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	66: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][66]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][66]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	67: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][67]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][67]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	68: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][68]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][68]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	69: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][69]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][69]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	70: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][70]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][70]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	71: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][71]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][71]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	72: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][72]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][72]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	73: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][73]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][73]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	74: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][74]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][74]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	75: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][75]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][75]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	76: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][76]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][76]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	77: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][77]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][77]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	78: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][78]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][78]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	79: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][79]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][79]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	80: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][80]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][80]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	81: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][81]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][81]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	82: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][82]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][82]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	83: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][83]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][83]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	84: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][84]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][84]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	85: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][85]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][85]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	86: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][86]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][86]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	87: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][87]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][87]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	88: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][88]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][88]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	89: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][89]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][89]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	90: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][90]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][90]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	91: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][91]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][91]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	92: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][92]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][92]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	93: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][93]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][93]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	94: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][94]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][94]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	95: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][95]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][95]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	96: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][96]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][96]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	97: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][97]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][97]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	98: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][98]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][98]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	99: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][99]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][99]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	100: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][100]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][100]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	101: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][101]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][101]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	102: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][102]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][102]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	103: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][103]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][103]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	104: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][104]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][104]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	105: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][105]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][105]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	106: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][106]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][106]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	107: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][107]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][107]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	108: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][108]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][108]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	109: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][109]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][109]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	110: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][110]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][110]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	111: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][111]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][111]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	112: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][112]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][112]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	113: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][113]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][113]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	114: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][114]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][114]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	115: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][115]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][115]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	116: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][116]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][116]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	117: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][117]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][117]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	118: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][118]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][118]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	119: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][119]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][119]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	120: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][120]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][120]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	121: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][121]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][121]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	122: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][122]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][122]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	123: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][123]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][123]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	124: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][124]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][124]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	125: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][125]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][125]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	126: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][126]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][126]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	127: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][127]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][127]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	128: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][128]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][128]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	129: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][129]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][129]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	130: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][130]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][130]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	131: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][131]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][131]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	132: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][132]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][132]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	133: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][133]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][133]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	134: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][134]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][134]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	135: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][135]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][135]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	136: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][136]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][136]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	137: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][137]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][137]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	138: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][138]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][138]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	139: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][139]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][139]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	140: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][140]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][140]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	141: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][141]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][141]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	142: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][142]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][142]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	143: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][143]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][143]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	144: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][144]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][144]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	145: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][145]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][145]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	146: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][146]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][146]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	147: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][147]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][147]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	148: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][148]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][148]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	149: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][149]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][149]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	150: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][150]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][150]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	151: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][151]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][151]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	152: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][152]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][152]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	153: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][153]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][153]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	154: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][154]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][154]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	155: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][155]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][155]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	156: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][156]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][156]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	157: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][157]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][157]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	158: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][158]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][158]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	159: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][159]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][159]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	160: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][160]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][160]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	161: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][161]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][161]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	162: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][162]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][162]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	163: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][163]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][163]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	164: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][164]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][164]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	165: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][165]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][165]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	166: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][166]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][166]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	167: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][167]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][167]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	168: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][168]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][168]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	169: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][169]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][169]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	170: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][170]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][170]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	171: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][171]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][171]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	172: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][172]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][172]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	173: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][173]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][173]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	174: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][174]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][174]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	175: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][175]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][175]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	176: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][176]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][176]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	177: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][177]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][177]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	178: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][178]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][178]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	179: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][179]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][179]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	180: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][180]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][180]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	181: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][181]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][181]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	182: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][182]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][182]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	183: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][183]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][183]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	184: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][184]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][184]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	185: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][185]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][185]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	186: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][186]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][186]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	187: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][187]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][187]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	188: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][188]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][188]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	189: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][189]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][189]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	190: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][190]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][190]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	191: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][191]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][191]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	192: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][192]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][192]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	193: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][193]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][193]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	194: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][194]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][194]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	195: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][195]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][195]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	196: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][196]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][196]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	197: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][197]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][197]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	198: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][198]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][198]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	199: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][199]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][199]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	200: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][200]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][200]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	201: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][201]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][201]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	202: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][202]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][202]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	203: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][203]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][203]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	204: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][204]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][204]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	205: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][205]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][205]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	206: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][206]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][206]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	207: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][207]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][207]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	208: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][208]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][208]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	209: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][209]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][209]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	210: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][210]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][210]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	211: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][211]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][211]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	212: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][212]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][212]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	213: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][213]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][213]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	214: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][214]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][214]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	215: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][215]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][215]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	216: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][216]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][216]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	217: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][217]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][217]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	218: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][218]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][218]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	219: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][219]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][219]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	220: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][220]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][220]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	221: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][221]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][221]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	222: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][222]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][222]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	223: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][223]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][223]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	224: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][224]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][224]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	225: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][225]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][225]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	226: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][226]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][226]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	227: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][227]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][227]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	228: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][228]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][228]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	229: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][229]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][229]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	230: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][230]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][230]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	231: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][231]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][231]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	232: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][232]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][232]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	233: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][233]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][233]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	234: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][234]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][234]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	235: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][235]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][235]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	236: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][236]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][236]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	237: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][237]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][237]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	238: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][238]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][238]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	239: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][239]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][239]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	240: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][240]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][240]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	241: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][241]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][241]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	242: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][242]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][242]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	243: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][243]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][243]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	244: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][244]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][244]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	245: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][245]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][245]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	246: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][246]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][246]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	247: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][247]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][247]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	248: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][248]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][248]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	249: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][249]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][249]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	250: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][250]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][250]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	251: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][251]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][251]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	252: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][252]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][252]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	253: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][253]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][253]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	254: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][254]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][254]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	255: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][255]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][255]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},

	256: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[][256]byte)(ptr)
		bslice := reflect.Zero(typ)
		subs := [][256]byte(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},
}
