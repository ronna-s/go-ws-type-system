# Generics

Generics are a pretty late edition to the Go language, but few people know that they were already with us all this time, 
for instance: `append` takes a slice of a type T and append items of type T to it, resulting in another slice of type T.
The problem was that the language had generics and could use generic types, but we couldn't define our own.

## When to use them?
It was a long time consensus that "real gophers" don't need generics, so much so that around the time the generics draft of 2020 was released, many gophers still said that they are not likely to use them.

Let's understand first the point that they were trying to make.

Consider [this code](https://gist.github.com/Xaymar/7c82ed127c8f1def53075f414a7df153), made using C++.
We see here generic code (templates) that allows an event to add functions (listeners) to its subscribers.
Let's ignore for a second that this code adds functions, not objects and let's assume it did take in objects with the function `Handle(e Event)`.
We don't need generics in Go to make this work because interfaces are implicit. As we saw already in C++ an object has to be aware of it's implementations, this is why to allow plugging-in of functionality we have to use generics in C++ (and in Java).

In Go this code would look something like [this](https://go.dev/play/p/Tqm_Hb0vcZb):

```go
package main

import "fmt"

type Listener interface {
	Handle(Event)
}

type Event struct {
	Lis []Listener
}

func (e *Event) Add(l Listener) {
	e.Lis = append(e.Lis, l)
}

func main() {
	var l Listener
	var e Event
	e.Add(l)
	fmt.Println(e)
}
```

**We didn't need generics at all!**

However, there are cases in Go where we have to use generics and until recently we used code generation for.
Those cases are when the behavior is derived from the type or leaks to the type's behavior:

For example:
The linked list
```go
// https://go.dev/play/p/ZpAqvVFAIDZ
package main

import "fmt"

type Node[T any] struct { // any is builtin for interface{}
  Value T
  Next  *Node[T]
}

func main() {
  n1 := Node[int]{1, nil}
  n2 := Node[int]{3, &n1}
  fmt.Println(n2.Value, n2.Next.Value)
}
```
Example 2 - [Addition](https://go.dev/play/p/dmeQEVxpyAq)
```go
package main

import "fmt"

type A int

// Add takes any type with underlying type int 
func Add[T ~int](i T, j T) T { 
  return i + j
}

func main() {
  var i, j A
  fmt.Println(Add(i, j))
}
```
Of course, you might not be likely to use linked lists in your day to day, but you are likely to use:
1. Repositories, database models, data structures that are type specific, etc.
2. Event handlers and processors that act differently based on the type.
3. The [concurrent map in the sync package](https://pkg.go.dev/sync#Map) which uses the empty interface.
4. [The heap](https://pkg.go.dev/container/heap#example-package-IntHeap)

The common thread to these examples is that before generics we had to trade generalizing certain behavior for type safety (or generate code to do so), now we can have both.

