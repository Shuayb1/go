package main

import (
	"fmt"
	"math"
)

// func as(a, b int) int {
// 	return a + b
// }

// // func main() {
// // 	fmt.Println(as(2, 4))
// // }

// func multireturn(me, you string) (string, string) {
// 	return you, me
// }

// // func main() {
// // 	a, b := multireturn("hello", "world")
// // 	fmt.Println(a, b)
// // }

// func unnamed_return(num int) (x, y int) {
// 	x = num * 4 / 9
// 	y = num - x
// 	return //nake return returns the named return values i.e x y values
// 	// fmt.Println(unnamed_return(10))
// }

// var c, python int = 1, 2 //variables can be at level or package fn, if not initialised, will return their 'zero valu'
// i.e o for int and float, false for bool, empty string for string

// func main() {
// 	//short variable declarations used inside fn only i.e :=

// 	var a, b, d = true, false, "no" //same as
// 	k := 4
// 	fmt.Println(k, c, python, a, b, d)
// }

//varibales and basic types, string formatting'
// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z, z)

// }

// //type conversios. Go does explicit conversions unlike C.
// func main() {
// 	i := 3
// 	f := float64(i)
// 	g := uint(f)
// 	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(g))
// }

// //type inference
// func main() {
// 	v := 42                           //the variable's type is inferred from the value on the right hand side.
// 	fmt.Printf("v is of the %T\n", v) //same as
// 	fmt.Println(reflect.TypeOf(v))
// }

// // constants: Constants can be character, string, boolean, or numeric values.
// // Constants cannot be declared using the := syntax.
// const Pi = 3.14

// func main() {
// 	fmt.Println("give me ", Pi, "snacks")
// }

//check again and again
// const (
// 	// Create a huge number by shifting a 1 bit left 100 places.
// 	// In other words, the binary number that is 1 followed by 100 zeroes.
// 	Big = 1 << 100
// 	// Shift it right again 99 places, so we end up with 1<<1, or 2.
// 	Small = Big >> 99
// )

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

// func main() {
// 	fmt.Println(needInt(Big))
// 	fmt.Println(needFloat(Big))
// 	fmt.Println(needFloat(Big))
// }

// //forloop
// func main() {
// 	// sum := 0
// 	// for i := 0; i < 5; i++ {
// 	// 	sum += i
// 	// }
// 	//same thing, but init and post sttm are optional
// 	//If you omit the loop condition it loops forever
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

//if statement

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x)) //float64 type
// }

// func main() {
// 	fmt.Println(sqrt(-1), sqrt(0), sqrt(1))
// }

// //if statement can take short statement too, also else is here
// func pw(x, y, z float64) float64 {
// 	if m := math.Pow(x, y); m < z {
// 		return m
// 	} else {
// 		fmt.Printf("%g >= %g\n", m, z)
// 	}
// 	return z
// }
// func main() {
// 	fmt.Println(pw(1, 2, 3), pw(4, 5, 6))
// }

// //switch case in go
// func main() {
// 	fmt.Print("Go is on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X")
// 	case "linux":
// 		fmt.Println("Linux")
// 	default:
// 		fmt.Println("%s", os)
// 	}
// }

// func main() {
// 	fmt.Println("When is Firday?")
// 	today := time.Now().Weekday()
// 	switch time.Friday {
// 	case today:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("2 days time")
// 	default:
// 		fmt.Println("hmmm")

// 	}
// }
//switvh no condition is switch tru
// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("morning")
// 	case t.Hour() < 17:
// 		fmt.Println("afternoon")
// 	default:
// 		fmt.Println("evening")
// 	}
// }

//defer statement defers the execution of a function until the surrounding function returns.
// func main() {
// 	defer fmt.Print("world")

// 	fmt.Print("hello ")
// }

//defer functions are pushed to stack, last in first out
// i.e the last print comes first, print from bottom to top
// func main() {
// 	defer fmt.Println("count ")  //5th print
// 	defer fmt.Println("5th")     // 4th print
// 	defer fmt.Println("counted") //3rd print
// 	defer fmt.Println("4th")     //second print
// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i) //1st print
// 	}
// }

//Pointers: its initial value is nil, and * denotes pointer's underlying value, ds is known as dereferencing and indirecting
// func main() {
// 	i, j := 10, 20

// 	p := &i //p points to i memory
// 	*p = 30 //set new i value through p
// 	fmt.Println(i, j)

// 	p = &j      //p points to j memory
// 	*p = *p / 5 //set new j value through p
// 	fmt.Println(i, j)
// }

//structs: collection of fields
// type Vertex struct {
// 	X, Y int
// }

// var (
// 	v1 = Vertex{1, 2}  //has type vertex
// 	v2 = Vertex{X: 9}  //Y:0 is implicit
// 	v3 = Vertex{}      //X:0 and Y:0
// 	p  = &Vertex{1, 2} //has type *Vertex
// )

// func main() {
// 	// v := Vertex{1, 2}
// 	// v.Y = 3 //replaced Y value

// 	// p := &v   //point to v memory
// 	// p.X = 1e9 // change X value in struct without explicit dereferencing
// 	fmt.Println(v1, v2, v3, p)
// }

//Arrays and slices
// func main() {
// 	mylist := [4]string{"hello", "love"}
// 	mylist[2] = "Nihun"
// 	mylist[3] = "J"
// 	fmt.Println(mylist, mylist[0], mylist[1:])
// }

//slice literals, like array literal,  but without the length
//zero value of a slice is nil.
// func main() {
// 	q := []int{1, 2, 3, 2, 5}
// 	fmt.Println(q[2:5])
// 	fmt.Println(cap(q))

// 	r := []bool{true, false, true, true, true}
// 	fmt.Println(len(r))

// 	s := []struct {
// 		i int
// 		j string
// 		k bool
// 	}{
// 		{1, "true", true},
// 		{4, "do", false},
// 	}
// 	fmt.Println(s)
// }

//make function; used to create a slice.  create dynamically-sized arrays.
//append
// func main() {
// 	a := make([]int, 5) //len(a) = 4
// 	// a := make([]int, 4, 5) //len(a) = 4, cap(a) = 5
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c := append(a, 10, 20, 40, 23) //append
// 	printSlice("c", c)

// 	d := c[2:5]
// 	printSlice("d", d)
// }
// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
// }

//Range: When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.
// var pow = []int{1, 2, 3, 4, 5, 6}

// func main() { //tested a value, and returned the value

// 	// for i, v := range pow {
// 	// 	fmt.Printf("2**%d = %d\n", i, v)
// 	// }
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 		// fmt.Println(pow[i])
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}
// }

//Map maps keys to value
// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bells Labs"] = Vertex{
// 		40.3232, -23.2323,
// 	}
// 	fmt.Println(m["Bells Labs"])
// }

//Map literals are like struct literals, but the keys are required.
// type Mystruct struct {
// 	age, department string
// }

// var mydict = map[string]Mystruct{
// 	"shuayb": {"20", "integration"}, //you can remove the top-level name
// 	"ade":    {"300", "admin"},
// 	"sola":   {"748", "sales"},
// }

// func main() {
// 	fmt.Println(mydict["ade"])
// }

//mutating maps
// func main() {
// 	m := make(map[string]int)

// 	m["ade"] = 23 //assigne a key to a value
// 	fmt.Println(m["ade"])

// 	m["sola"] = 232
// 	fmt.Println(m["sola"])

// 	m["sola"] = 232
// 	fmt.Println(m["sola"])

// 	delete(m, "sola") //delete a key
// 	fmt.Println(m["sola"])

// 	// check if a key is present, with a 2-value argument
// 	value, ok := m["sola"]
// 	fmt.Println(m)
// 	fmt.Println(value, ok)
// }

// Methods:  a function with a special receiver argument.
// see a receiver as a parameter passed to a fn.
// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X * v.Y)
// }
// func main() {
// 	v := Vertex{3, 5}
// 	fmt.Println(v.Abs())
// }

// NON-STRUCT TYPE.
// NB: type declared must be same as receiver, and type should be same
// type Myfloat float64

// func (f Myfloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
// func main() {

// 	f := Myfloat(-math.Sqrt(2))
// 	fmt.Println(f.Abs())
// }

// POINTER receiver
// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.Y * v.Y)
// }

// // func (v *Vertex) Scale(f float64)
// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10)
// 	fmt.Println(Abs(v))
// }

// The equivalent thing happens in the reverse direction.

// Functions that take a value argument must take a value of that specific type:

// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!
// while methods with value receivers take either a value or a pointer as the receiver when they are called:

// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK


// INTERFACE
type Gh interface{
	Abs() float64
}

func main() {
	var a Gh
	f := 
}