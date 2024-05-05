package main

import "fmt"

const FIXED  = "hello"
var (
	 a int
	 b  = "how are you"
	 c string = "hello world"
	 d int = 6
	 i = 10
	 txt = "hello"
	 t = true
	 f = false
	 k = 3.141
)
func main() {
	// fmt.Print(a,"\n", d)
	// fmt.Println(a)

	// fmt.Println(c)
	// fmt.Printf("b has value: %v and type: %T\n", b,b )
	// fmt.Printf("c has value: %v and type: %T\n", d ,d)
	// fmt.Printf("FIXED has value: %v and type: %T\n", FIXED, FIXED)

	// fmt.Printf("%v\n", d)
	// fmt.Printf("%#v\n",d)
	// fmt.Printf("%v%%\n", d)
	// fmt.Printf("%T\n", d)

	// fmt.Printf("%v\n", b)
	// fmt.Printf("%#v\n",b)
	// fmt.Printf("%v%%\n", b)
	// fmt.Printf("%T\n", b)

//  fmt.Printf("%b\n", i) // output : 1010
//   fmt.Printf("%d\n", i) // output : 10
//   fmt.Printf("%+d\n", i) // output : +10
//   fmt.Printf("%o\n", i) // output : 12
//   fmt.Printf("%O\n", i) // output : 0o12
//   fmt.Printf("%x\n", i) // output : a
//   fmt.Printf("%X\n", i) // output : A
//   fmt.Printf("%#x\n", i) // output : 0xa
//   fmt.Printf("%4d\n", i) // output :    10
//   fmt.Printf("%-4d\n", i) // output : 10
//   fmt.Printf("%04d\n", i) // output : 0010


//   fmt.Printf("%s\n", txt) // output : hello
//   fmt.Printf("%q\n", txt) // output : "hello"
//   fmt.Printf("%8s\n", txt) // output :     hello
//   fmt.Printf("%-8s\n", txt) // output : hello
//   fmt.Printf("%x\n", txt)  // output : 68656c6c6f
//   fmt.Printf("% x\n", txt) // 68 65 6c 6c 6f


// fmt.Printf("%t\n", t) // output : true
// fmt.Printf("%t\n", f) // output : false

  fmt.Printf("%e\n", k) // output : 3.141000e+00
  fmt.Printf("%f\n", k) // output : 3.141000
  fmt.Printf("%.2f\n", k) // output : 3.14
  fmt.Printf("%6.2f\n", k) // output :   3.14
  fmt.Printf("%g\n", k) // output : 3.141



}