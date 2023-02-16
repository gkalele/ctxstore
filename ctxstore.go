package ctxstore

import (
	"context"
	"sync"
)

const valKey = "__1234567890"
const lockKey = "__mutexlock"
const typeCollisionsAllowedKey = "__typecollisionsflag"

type Store map[string]any

type Options struct {
	TypeCollisionsAllowed bool
}

func GenerateRootContext(parent context.Context, options Options) (context.Context, context.CancelFunc) {
	c, cancelFn := context.WithCancel(context.Background())
	c = context.WithValue(c, valKey, make(Store))
	c = context.WithValue(c, lockKey, &sync.RWMutex{})
	c = context.WithValue(c, typeCollisionsAllowedKey, options.TypeCollisionsAllowed)
	return c, cancelFn
}

func Lock(ctx context.Context) *sync.RWMutex {
	return ctx.Value(lockKey).(*sync.RWMutex)
}

// Put
// Use Go generics to store all types into the map
// Detect collisions between types - panic to signal invalid key generation logic
func Put[T any](ctx context.Context, key string, val T) {
	l := Lock(ctx)
	l.Lock()
	defer l.Unlock()
	oldVal, ok := ctx.Value(valKey).(Store)[key]
	if !ok {
		// New entry created in the dictionary
		ctx.Value(valKey).(Store)[key] = val
		return
	}
	_, goodCast := oldVal.(T)
	if !goodCast {
		if !ctx.Value(typeCollisionsAllowedKey).(bool) {
			// Previous value in dictionary was NOT of the same type
			panic("Invalid types colliding for key " + key)
		}
	}
	ctx.Value(valKey).(Store)[key] = val
}

func Get[T any](ctx context.Context, key string) (T, bool) {
	var emptyValue T
	l := Lock(ctx)
	l.RLock()
	defer l.RUnlock()
	v, ok := ctx.Value(valKey).(Store)[key]
	if !ok {
		return emptyValue, false
	}
	return v.(T), ok
}
