// Copyright (c) 2018, CAOHONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

func main() {
	const templateText = `// Code generated using genbucketers.go; DO NOT EDIT.

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
{{range $i,$typ := genETypes}}
	{{$i}}: func(typ reflect.Type, ptr unsafe.Pointer) bucket {
		alias := *(*[]{{$typ}})(ptr)
		bslice := reflect.Zero(typ)
		subs := []{{$typ}}(nil)
		return bucket{
			func() int { return len(subs) },
			func(i int) { subs = append(subs, alias[i]) },
			func() interface{} {
				bh := (*reflect.SliceHeader)(iface.Ptr(bslice.Interface()))
				sh := (*reflect.SliceHeader)(iface.Ptr(subs))
				bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Cap
				return bslice.Interface()
			},
			func(offset int) int { return copy(alias[offset:], subs) },
		}
	},
{{end}}}
`

	funcMap := template.FuncMap{
		"genETypes": genETypes,
	}

	tmpl, err := template.New("bucketers").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	var out bytes.Buffer

	// Run the template to verify the output.
	err = tmpl.Execute(&out, "")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	if err := ioutil.WriteFile("bucketers.go", out.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

func genETypes() (etypes [256 + 1]string) {
	for i := 0; i <= 256; i++ {
		etypes[i] = fmt.Sprintf("[%d]byte", i)
	}
	etypes[0] = "struct{}"
	etypes[1] = "int8"
	etypes[2] = "int16"
	etypes[4] = "int32"
	etypes[8] = "int64"
	return
}
