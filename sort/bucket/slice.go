// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run genbucketers.go

package bucket

import (
	"reflect"
	"unsafe"

	"github.com/cnotch/algo/reflect/iface"
)

// Slice 对 slice 进行桶排序;如果对性能要求高，直接按照 Ints 方法实现专有类型的桶排序。
// 	slice 待排序的集合；
//  k 桶索引值最大值，它表示桶索引值的范围为 [0,k]；
//	key 计算集合中元素桶索引值函数，返回值必须在 [0,k]；
//	stable 对指定桶执行排序操作的函数
func Slice(slice interface{}, k int, key func(i int) int, stable func(slice interface{})) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("Slice isn't slice type")
	}

	n := rv.Len()
	if n < 2 {
		return
	}

	buckets := makeBuckets(rv, k+1)

	for i := 0; i < n; i++ {
		bi := key(i)
		buckets[bi].append(i)
	}

	offset := 0
	for i := 0; i <= k; i++ {
		bucket := buckets[i]
		if bucket.len() == 0 {
			continue
		}

		// 对单个桶排序
		stable(bucket.slice())
		// 复制到初始切片,并步进下次复制的位置
		offset += bucket.copy(offset)
	}
}

// 桶结构
type bucket struct {
	len    func() int           // 桶内元素个数
	append func(i int)          // 向桶添加元素
	slice  func() interface{}   // 返回桶的存储[]T
	copy   func(offset int) int // 拷贝桶内容到合适的位置
}

func makeBuckets(rv reflect.Value, count int) (b []bucket) {
	typ := rv.Type()         // slice Type
	etyp := typ.Elem()       // slice elem Type
	size := int(etyp.Size()) // slice elem size
	kind := etyp.Kind()      // slice elem kind

	ptr := iface.Ptr(rv.Interface()) // slice pointer
	b = make([]bucket, count)

	if kind < reflect.Chan || kind == reflect.Struct {
		if size < len(bucketers) {
			for i := 0; i < count; i++ {
				b[i] = bucketers[size](typ, ptr)
			}
			return
		}

		for i := 0; i < count; i++ {
			var alias, subs []byte
			aliasHeader := (*reflect.SliceHeader)(unsafe.Pointer(&alias))
			header := *(*reflect.SliceHeader)(ptr)

			aliasHeader.Data = header.Data
			aliasHeader.Len = size * header.Len
			aliasHeader.Cap = size * header.Cap

			bslice := reflect.Zero(typ)
			b[i] = bucket{
				func() int { return len(subs) / size },
				func(i int) {
					offset := i * size
					subs = append(subs, alias[offset:offset+size]...)
				},
				func() interface{} {
					bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
					sh := (*reflect.SliceHeader)(unsafe.Pointer(&subs))
					bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len/size, sh.Cap/size
					return bslice.Interface()
				},
				func(offset int) int { return copy(alias[offset:], subs) },
			}
		}
		return
	}

	switch kind {
	case reflect.String:
		for i := 0; i < count; i++ {
			alias := *(*[]string)(ptr)  // 初始切片别名
			bslice := reflect.Zero(typ) // 和源保持类型一致的桶切片
			subs := []string(nil)       // 桶切片的替身
			b[i] = bucket{
				func() int { return len(subs) },
				func(i int) { subs = append(subs, alias[i]) },
				func() interface{} {
					bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
					sh := (*reflect.SliceHeader)(iface.Ptr(subs))
					bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
					return bslice.Interface()
				},
				func(offset int) int { return copy(alias[offset:], subs) },
			}
		}
		return
	case reflect.Interface:
		for i := 0; i < count; i++ {
			alias := *(*[]interface{})(ptr) // 初始切片别名
			bslice := reflect.Zero(typ)     // 和源保持类型一致的桶切片
			subs := []interface{}(nil)      // 桶切片的替身
			b[i] = bucket{
				func() int { return len(subs) },
				func(i int) { subs = append(subs, alias[i]) },
				func() interface{} {
					bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
					sh := (*reflect.SliceHeader)(iface.Ptr(subs))
					bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
					return bslice.Interface()
				},
				func(offset int) int { return copy(alias[offset:], subs) },
			}
		}
		return
	case reflect.Slice:
		for i := 0; i < count; i++ {
			alias := *(*[][]reflect.SliceHeader)(ptr) // 初始切片别名
			bslice := reflect.Zero(typ)               // 和源保持类型一致的桶切片
			subs := [][]reflect.SliceHeader(nil)      // 桶切片的替身
			b[i] = bucket{
				func() int { return len(subs) },
				func(i int) { subs = append(subs, alias[i]) },
				func() interface{} {
					bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
					sh := (*reflect.SliceHeader)(iface.Ptr(subs))
					bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
					return bslice.Interface()
				},
				func(offset int) int { return copy(alias[offset:], subs) },
			}
		}
		return
	default: // 指针类: Chan、Func、Map、Ptr、UnsafePointer
		for i := 0; i < count; i++ {
			alias := *(*[]unsafe.Pointer)(ptr) // 初始切片别名
			bslice := reflect.Zero(typ)        // 和源保持类型一致的桶切片
			subs := []unsafe.Pointer(nil)      // 桶切片的替身
			b[i] = bucket{
				func() int { return len(subs) },
				func(i int) { subs = append(subs, alias[i]) },
				func() interface{} {
					bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
					sh := (*reflect.SliceHeader)(iface.Ptr(subs))
					bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
					return bslice.Interface()
				},
				func(offset int) int { return copy(alias[offset:], subs) },
			}
		}
		return
	}
}
