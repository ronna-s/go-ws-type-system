# Functional Programming

## Is Go a functional programming language?
No. Go is a multi-paradigm language that supports procedural, object-oriented, and concurrent programming. However, Go has some functional programming features, such as first-class functions, higher-order functions, and closures.
- Go doesn't provide lazy evaluation, which is a feature of functional programming languages that allows expressions to be evaluated only when needed (so we have to do it ourselves).
- Go doesn't provide immutable data structures, which are common in functional programming languages. 
- Go doesn't have tail call optimization, which is a feature of functional programming languages that allows recursive functions to be optimized to avoid stack overflow (and also generally good for performance).

### What is TCO (Tail Call Optimization)?
It's the practice of overriding the current stack frame with the next one (we just jump to the last function in the current function), instead of adding a new one.
Go doesn't support it for various reasons.

## What is functional programming?
Functional programming is the practice of building software by composing pure functions, avoiding shared state, mutable data, and side-effects. It is declarative rather than imperative, and application state flows through pure functions, e,g.:

```go
    fn1(fn2(fn3(someArgs)))
```
Instead of:
    
```go
    res := Proc1(someArgs)
    Proc2(otherArgs, res)
	sideEffects := ReadSideEffects()
	Proc3(sideEffects)
```

## Why use functional programming?
Functional programming is a powerful tool for building resilient and maintainable software. 

## See also:
- A very cohesive [Functional Programming package for Go by IBM](https://github.com/IBM/FP-GO) and the [video](https://www.youtube.com/watch?v=Jif3jL6DRdw) about it. 
- An Introduction to Functional Programming in Go - Eleanor McHugh [video](https://www.youtube.com/watch?v=OKlhUv8R1ag).
- The IO Monad in Haskell and how it's implemented [video](https://www.youtube.com/watch?v=fCoQb-zqYDI)
- Monadic operartions overview (in Haskell) [post](https://www.adit.io/posts/2013-04-17-functors,_applicatives,_and_monads_in_pictures.html)

