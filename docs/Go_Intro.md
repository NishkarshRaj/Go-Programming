# Introduction to Go Programming

Developed at Google in 2009 for system programming with advanced features over C Programming.

## Basic improvements provided by Golang over C include the following:

* Garbage collection
* Type safety
* dynamic typing capability
* advanced built in data types like (key,value) pairs
* supports CSP-style programming for concurrent systems


## Features of Go programming

* Fast compilation speed
* Inbuilt support for concurrency
* Programs are consice and simple
* produces native statically linked libraries internally
* Supports adaptive environment features similar to dynamic languages to support portability.
* Support for interfaces
* Support for type embedding

## Features not supported in Go programming

####  Many features of common languages like C++ and Java are not supported in Golang to keep it simple and concise. Some of the features not supported are:

* Type inheritance
* Operator and Method Overloading
* support for Pointers
* Assertions
* Generic programming
* Circular dependencies among packages

## An overview of Golang

* Compile and linking of libraries
* Dependency Management in terms of packages
* Static Language: Type specification at time of declaration - Eg. int a = 0;
* Extension: **.go**

## Installing Go on Linux

```
$ apt install golang-go
$ go version
```

## Hello World Code

```go
package main

import "fmt"

func main() {
fmt.Println("Hello,World!")
}
```

* Packages are needed just like Java for namespaces
* fmt (format package) is the equivalent of stdio of C
* fmt.Println() == printf()
* Semicolons are not required
* The opening brace must be in the same line as of function definition.
* Execute program with `$ go run [filename].go`

## Structure of Go programs

* Package Declaration
* Import Packages
* Functions
* Variables
* Comments
* Expressions

**Note**

* The first line must always declare package in which current program lies.
* Import is a pre-processor directive to import necessary packages.
* Methods of imported packages always start with Capital Letters.
