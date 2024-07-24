package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	/* 	sl := []int{1, 2, 3}
	   	greeting := "hello"
	   	greetingPtr := &greeting
	   	f := Foo{A: 10, B: "Salutations"}
	   	fp := &f

	   	slType := reflect.TypeOf(sl)
	   	gType := reflect.TypeOf(greeting)
	   	grpType := reflect.TypeOf(greetingPtr)
	   	fType := reflect.TypeOf(f)
	   	fpType := reflect.TypeOf(fp)

	   	examiner(slType, 0)
	   	examiner(gType, 0)
	   	examiner(grpType, 0)
	   	examiner(fType, 0)
	   	examiner(fpType, 0) */

	greeting := "hello"
	f := Foo{A: 10, B: "Salutations"}

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)

	fType := reflect.TypeOf(f)
	fVal := reflect.New(fType)
	fVal.Elem().Field(0).SetInt(20)
	fVal.Elem().Field(1).SetString("Greetings")
	f2 := fVal.Elem().Interface().(Foo)
	fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)

}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
