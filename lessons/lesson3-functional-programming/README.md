# Functional Programming

### Is Go a functional programming language?

### What is functional programming?
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

