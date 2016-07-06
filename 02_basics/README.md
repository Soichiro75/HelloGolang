# A Tour of Go.

[Basics/Packages, variables, and functions.](https://go-tour-jp.appspot.com/basics/1)


## Note:

### What is `>>`?

`>>`, `<<` are `Bitwise Operators`

(11_basic-types.go, )

op1 >> op2 -- The SHIFT RIGHT operator moves the bits to the right, discards the far right bit, and assigns the leftmost bit a value of 0. Each move to the right effectively divides op1 in half.

For Example:
```
You can see bellowing Sample Code in 99_bitwiseOperatorTest.go:

package main

import (
	"fmt"
)

var (
	BitwiseOperator uint64 = 10 << 1
)

func main() {
	fmt.Println(BitwiseOperator) //=>20
}


"BitwiseOperator uint64 = 10<<1"
10(Decimal Number) is 1010(Binary Number).
"<<1" means that moves the bits to the left 1, and assigns the rightmost bit a value of 0.
That is means 1010(Binary Number) changes to 10100(Binary Number).
So x<<1 means that "<<1" makes "x" double.
10<<1 means that 20(Decimal Number)

```

For more details, see [https://golang.org/ > doc > References/Language Specification](https://golang.org/ref/spec#Operators)



### What is `%v`?

`%v`, `%q`, `%T`, `%g`, etc. are verbs.

(02_imports.go, 11_basic-types.go, 12_zero.go)

`fmt.Printf("%v %v %v %q\n", i, f, b, s)`

`%v` and `%q` are 'verbs' formatted by 'package fmt'.

`%v`: the value in a default format

`%q`: a single-quoted character literal safely escaped with Go syntax.

`%T`: a Go-syntax representation of the type of the value.

`%g`: %e for large exponents, %f otherwise.

`%e`: %e	scientific notation, e.g. -1.234456e+78

`%f`: decimal point but no exponent, e.g. 123.456

For more details, see [https://golang.org > Packages > fmt](https://golang.org/pkg/fmt/) .
