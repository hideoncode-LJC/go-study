package variables

import "fmt"

func PrintVariables1() {
	var a int64 
	var b float64
	var c string
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables2() {
	var a int64 
	var b float64
	var c string
	a = 1
	b = 12.54646
	c = "string ???"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables3() {
	var a int64 = 1
	var b float64 = 123.45363645
	var c string = "string ??"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables4() {
	var a = 1
	var b = 123.4536435
	var c = "string ????"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables5() {
	var (
		a int
		b float64
		c string
	)
	a = 1
	b = 1.23543
	c = "string ????"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables6() {
	var (
		a int = 1
		b float64 = 23.4546
		c string = "string??"
	)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables7() {
	var (
		a = 1
		b = 213.546436
		c = "string >>>????"
	)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintVariables8() {
	a := 1
	b := 23.546547
	c := "string ??? 2314"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", &a, &a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", &b, &b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", &c, &c)
}

func PrintConstant() {
	const (
		pi = 3.14
		e = 2.7
	)
	const (
		n1 = 100
		n2
		n3
		n4 = 345
		n5
		n6
	)
	
	fmt.Printf("%v %T\n", pi, pi)
	fmt.Printf("%v %T\n", e, e)
	fmt.Printf("%v %T\n", n1, n1)
	fmt.Printf("%v %T\n", n2, n2)
	fmt.Printf("%v %T\n", n3, n3)
	fmt.Printf("%v %T\n", n4, n4)
	fmt.Printf("%v %T\n", n5, n5)
	fmt.Printf("%v %T\n", n6, n6)
}