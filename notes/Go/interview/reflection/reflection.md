# Reflection

*Clear is better than clever. Reflection is never clear.*
*Rob Pike*

## What is the reflection

 Reflection gives you the ability to examine(invistigate) types at runtime. It also allows you to examine, modify, and create variables, functions, and structs at runtime.

 Reflection in Go is built around three concepts: `Types`, `Kinds`, and `Values`. The reflect package in the standard library is the home for the types and functions that implement reflection in Go.

 ## Types
  You can use reflection to get the type of a variable 

`varType := reflects.TypeOf(var)`

`Name()` 
 This returns, not surprisingly, the name of the type. Some types, like a **slice** or a **pointer**, don’t have names and this method returns an **empty** string.


`Kind()` 
The kind is what the type is made of — a slice, a map, a pointer, a struct, an interface, a string, an array, a function, an int or some other primitive type. The difference between the kind and the type can be tricky to understand, but think of it this way. If you define a struct named Foo, the kind is struct and the type is Foo.

`varType.Elem()`
If your variable is a pointer, map, slice, channel, or array, you can find out the contained type by using varType.Elem().

`reflect.StructField`
If your variable is a struct, you can use reflection to get the number of fields in the struct, and get back each field’s structure contained in a reflect.StructField struct. The reflect.StructField gives you the name, order, type, and struct tags on a fields.

`NumField()` and `Field()`
The NumField() method returns the number of fields in a struct and the Field(i int) method returns the reflect.Value of the ith field.

```
func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
```

## Making a New Instance

We can  use reflection to read, set, or create values

`refVal := reflect.ValueOf(var)`
Create a reflect.Value instance for variable.  If you want to be able to use reflection to modify the value, you have to get a pointer to the variable with `refPtrVal := reflect.ValueOf(&var)`; if you don’t, you can read the value using reflection, but you can’t modify it.
If you want to modify a value, remember it has to be a pointer, and you have to dereference the pointer first. You use `refPtrVal.Elem().Set(newRefVal)` to make the change, and the value passed into Set() has to be a `reflect.Value` too.


## Making Without Make
In addition to making instances of built-in and user-defined types, you can also use reflection to make instances that normally require the make function. You can make a slice, map, or channel using the reflect.MakeSlice, reflect.MakeMap, and reflect.MakeChan functions. In all cases, you supply a reflect.Type and get back a reflect.Value that you can manipulate with reflection, or that you can assign back to a standard variable.

**Example**

```
func main() {
	// declaring these vars, so I can make a reflect.Type
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)

	// here are the reflect.Types
	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)

	// and here are the new values that we are making
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	mapReflect := reflect.MakeMap(mapType)

	// and here we are using them
	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
 
```

## Making Functions
Reflection doesn’t just let you make new places to store data. You can use reflection to make new functions using the reflect.MakeFunc function. This function expects the reflect.Type for the function that we want to make and a closure whose input parameters are of type []reflect.Value and whose output parameters are also of type []reflect.Value. Here’s a quick example, which creates a timing wrapper for any function that’s passed into it:

```
func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	fmt.Println("starting")
	time.Sleep(1 * time.Second)
	fmt.Println("ending")
}

func timeMeToo(a int) int {
	fmt.Println("starting")
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	fmt.Println("ending")
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func())
	timed()
	timedToo := MakeTimedFunction(timeMeToo).(func(int) int)
	fmt.Println(timedToo(2))
}
```

## New structs

 You can make brand-new structs at runtime by passing a slice of `reflect.StructField` instances to the `reflect.StructOf` function. This one is a bit weird; we are making a new type, but we don’t have a name for it, so you can’t really turn it back into a “normal” variable. You can create a new instance and use `Interface()` to put the value into a variable of type `interface{}`, but if you want to set any values on it, you need to use reflection.

```
func MakeStruct(vals ...interface{}) interface{} {
	var sfs []reflect.StructField
	for k, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", (k + 1)),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so.Interface()
}

func main() {
	s := MakeStruct(0, "", []int{})
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can’t actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(s)

	// getting and setting the int field
	fmt.Println(sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface())

	// getting and setting the string field
	fmt.Println(sr.Elem().Field(1).Interface())
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println(sr.Elem().Field(1).Interface())

	// getting and setting the []int field
	fmt.Println(sr.Elem().Field(2).Interface())
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println(sr.Elem().Field(2).Interface())
}
```

## What Can’t You Do?
There’s one big limitation on reflection. While you can use reflection to create new functions, there’s no way to create new methods at runtime. This means you cannot use reflection to implement an interface at runtime. 