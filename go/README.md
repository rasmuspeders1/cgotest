## Golang 

With automatic garbage collection, is `Go` really suitable as a systems programming language?


Definition of systems programming language:
"Systems programming aims to produce software and software platforms which provide services to other software, are performance constrained, or both"

Then answer is Yes.

While Go does not allow the same fine-tuned control of memory usage as given by C/C++, it gives much stronger control than managed languages such as Java, C#, Python, etc.

Best of Go:
Go has a peculiar, but highly efficient, concurrency model focused on goroutines. Using user mode scheduling along with miniature, dynamically growing stacks will get you a level of parallelism you can only dream about in other languages


### History

C++’s creators had their own goals: to build in higher layers of abstraction into C.
Google conceived of Go as a replacement for C++, the same motivations behind another major language: Rust.

Another iteration Google `Carbon`
1) C → C++, 2) C++ → Go, 3) Go → Carbon?

(there is no D language anymore!?)


Incrementally improving C++ is extremely difficult, both due to the technical debt itself and challenges with its evolution process.
Go has also some drawbacks.. 


### Go-C++ Interoperability

`cgo` is a tool that enables Go code to interoperate with C code. 
It allows Go programs to call C functions and use C libraries, and vice versa.

Go code that uses `cgo` can call C functions directly and access C variables, and C code can call Go functions as if they were regular C functions.

For most practical use cases, you will find yourself wrapping your C++ code with a C API (it is called CGO after all!) 
which can amount to a fair amount of boilerplate code.

(But `Carbon` have seamless, bidirectional interoperability with C++. Too young and too early to use!). 


### Simple examples

OK, lets see how good or complex is CGO implementation.

1. QR code parser:  C++ parser code -> C wrapper -> C library -> Go client: https://github.com/sigidagi/goffi 
2. Library in Go -> C++ client: https://github.com/sigidagi/goffi

