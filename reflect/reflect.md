#### What’s Reflection?
Most of the time, variables, types, and functions in Go are pretty straightforward. When you need a type, you define a type:

``` go
type Foo struct {
  A int
  B string
}
```

When you need a variable, you define a variable:

``` go
var x Foo
```

And when you need a function, you define a function:

``` go
func DoSomething(f Foo) {
  fmt.Println(f.A, f.B)
}
```

But sometimes you want to work with variables at runtime using information that didn’t exist when the program was written. Maybe you’re trying to map data from a file or network request into a variable. Maybe you want to build a tool that works with different types. In those situations, you need to use reflection. Reflection gives you the ability to examine types at runtime. It also allows you to examine, modify, and create variables, functions, and structs at runtime.

Reflection in Go is built around three concepts: **Types**, **Kinds**, and **Values**. The reflect package in the standard library is the home for the types and functions that implement reflection in Go.

#### Finding Your Type
First let’s look at types. You can use reflection to get the type of a variable var with the function call `varType := reflect.TypeOf(var)`. This returns a variable of type reflect.Type, which has methods with all sorts of information about the type that defines the variable that was passed in.

The first method we’ll look at is Name(). This returns, not surprisingly, the name of the type. Some types, like a slice or a pointer, don’t have names and this method returns an empty string.
The next method, and in my opinion the first really useful one, is Kind(). The kind is what the type is made of — a slice, a map, a pointer, a struct, an interface, a string, an array, a function, an int or some other primitive type. The difference between the kind and the type can be tricky to understand, but think of it this way. If you define a struct named Foo, the kind is struct and the type is Foo.
One thing to be aware of when using reflection: everything in the reflect package assumes that you know what you are doing and many of the function and method calls will panic if used incorrectly. For example, if you call a method on reflect.Type that’s associated with a different kind of type than the current one, your code will panic. Always remember to use the kind of your reflected type to know which methods will work and which ones will panic.
If your variable is a pointer, map, slice, channel, or array, you can find out the contained type by using varType.Elem().
If your variable is a struct, you can use reflection to get the number of fields in the struct, and get back each field’s structure contained in a reflect.StructField struct. The reflect.StructField gives you the name, order, type, and struct tags on a fields.
Since a few lines of source code are worth a thousand words of prose, here’s a simple example for dumping out the type information for a variety of variables:

``` go
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
	simpleVar := []int{1, 2, 3, 4}
	simpleVarPtr := &simpleVar

	greeting := "Hello yo"
	greetingPtr := &greeting

	fooVar := Foo{A: 10, B: "twenty"}
	fooVarPtr := &fooVar

	typeOfSimpleVar := reflect.TypeOf(simpleVar)
	typeOfSimpleVarPtr := reflect.TypeOf(simpleVarPtr)
	typeOfGreeting := reflect.TypeOf(greeting)
	typeOfGreetingPtr := reflect.TypeOf(greetingPtr)
	typeOfFooVar := reflect.TypeOf(fooVar)
	typeOfFooVarPtr := reflect.TypeOf(fooVarPtr)

	examiner(typeOfSimpleVar, 0)
	examiner(typeOfSimpleVarPtr, 0)
	examiner(typeOfGreeting, 0)
	examiner(typeOfGreetingPtr, 0)
	examiner(typeOfFooVar, 0)
	examiner(typeOfFooVarPtr, 0)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "type is", t.Name(), "and kind is", t.Kind())

	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is:", f.Name, "type is:", f.Type.Name(), "kind is:", f.Type.Kind())

			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is:", f.Tag.Get("tag2"))
			}
		}
	}
}
```

And the output looks like this:


```bash
 type is  and kind is slice
         Contained type:
         type is int and kind is int
 type is  and kind is ptr
         Contained type:
         type is  and kind is slice
                 Contained type:
                 type is int and kind is int
 type is string and kind is string
 type is  and kind is ptr
         Contained type:
         type is string and kind is string
 type is Foo and kind is struct
         Field 1 name is: A type is: int kind is: int
                 tag is tag1:"First Tag" tag2:"Second Tag"
                 tag1 is First Tag tag2 is: Second Tag
         Field 2 name is: B type is: string kind is: string
 type is  and kind is ptr
         Contained type:
         type is Foo and kind is struct
                 Field 1 name is: A type is: int kind is: int
                         tag is tag1:"First Tag" tag2:"Second Tag"
                         tag1 is First Tag tag2 is: Second Tag
                 Field 2 name is: B type is: string kind is: string
```

#### Making a New Instance
In addition to examining the types of your variables, you can also use reflection to read, set, or create values. First you need to use refVal := reflect.ValueOf(var) to create a reflect.Value instance for your variable. If you want to be able to use reflection to modify the value, you have to get a pointer to the variable with `refPtrVal := reflect.ValueOf(&var);` if you don’t, you can read the value using reflection, but you can’t modify it.

Once you have a reflect.Value, you can get the reflect.Type of the variable with the Type() method.
If you want to modify a value, remember it has to be a pointer, and you have to dereference the pointer first. You use refPtrVal.Elem().Set(newRefVal) to make the change, and the value passed into Set() has to be a reflect.Value too.

If you want to create a new value, you can do so with the function call `newPtrVal := reflect.New(varType)`, passing in a reflect.Type. This returns a pointer value that you can then modify. using Elem().Set() as described above.

Finally, you can go back to a normal variable by calling the Interface() method. Because Go doesn’t have generics, the original type of the variable is lost; the method returns a value of type interface{}. If you created a pointer so that you could modify the value, you need to dereference the reflected pointer by using Elem().Interface(). In both cases, you will need to cast your empty interface to the actual type in order to use it.

Here’s some code to demonstrate these concepts:

``` go
package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	fooVar := Foo{A: 10, B: "twenty"}

	tempFooType := reflect.TypeOf(fooVar)
	tempFoo := reflect.New(tempFooType)
	tempFoo.Elem().Field(0).SetInt(1000)
	tempFoo.Elem().Field(1).SetString("fifty")

	newFoo := tempFoo.Elem().Interface().(Foo)

	fmt.Println(fooVar)
	fmt.Println(newFoo)
}
```

The output looks like:

``` bash 
{10 twenty}
{1000 fifty}
```