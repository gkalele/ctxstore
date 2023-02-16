# ctxstore

Context Store

Golang use contexts as a generic store between goroutines

## An extension of the context.WithValue with Go Generics multiple type storage

The context.WithValue mechanism allows storing data inside a Context, that can be shared by all child contexts.

Using Go Generics, we can store and retrieve data from the context, in a typesafe concurrent manner.

This cleans up the code and message passing between sibling goroutines - especially when they are sharing common handles to databases etc.

Any goroutine encountering a need to reconnect can quickly store the new handle in the context store and it is immediately available to all goroutines using the same root context.
