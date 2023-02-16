# ctxstore

Context Store

A dependency-free library to use contexts as a generic store between goroutines

## An extension of the context.WithValue with Go Generics multiple type storage

The context.WithValue mechanism allows storing data inside a Context, that can be shared by all child contexts.

Using Go Generics, we can store and retrieve data from the context, in a typesafe concurrent manner.

This cleans up the code and message passing between sibling goroutines - especially when they are sharing common handles to databases etc.

Any goroutine encountering a need to reconnect can quickly store the new handle in the context store and it is immediately available to all goroutines using the same root context.

## Concurrent Puts and Gets

A RWMutex lock embedded in the context ensures that Puts work under
Locks and Gets under RLocks.

## Type collisions

Attempting to overwrite an existing key with a value of
a different type than previously inserted can be configured to panic
or allow overwriting, just like maps in Python can.

## Examples

	import (
		"ctxstore"

		"fmt"
	)

	func main() {
			ctx, _ := ctxstore.GenerateRootContext(context.Background())
			ctxstore.Put[int](ctx, "number", 0)
			ctxstore.Put[string](ctx, "somestring", "string now")

			val, ok := ctxstore.Get[int](ctx, "number")
			if ok {
				fmt.Printf("Value retrieved is %d\n", val)
			}
			val1, ok1 := ctxstore.Get[string](ctx, "somestring")
			if ok1 {
				fmt.Printf("Value retrieved is %s\n", val1)
			}
	}
