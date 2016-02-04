## gcalc

`gcalc` is a simple calculator in Golang. 
Just for practicing lexical analysis and LL(1).

Differet from traditional implementation, this calculator using OOP and a small state machine for extracting a float number.

### Basic usage

```go
packge main

import "github.com/Focinfi/gcalc"

result := gcalc.Compute(" 1*(-1)* (5.00 + (-2)) * 3-(1.0+ 2)* 3 *(-1)")
// result will be 0
// any expression error will make gcalc panic :|
```

