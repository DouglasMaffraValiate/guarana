package requestCTX

import (
	"sync"
	"sync/atomic"
)

var nextKeyID atomic.Uint64

// Key identifica um valor tipado armazenado no Context.
type Key[T any] struct {
	id uint64
}

// NewKey cria uma chave tipada.
func NewKey[T any]() Key[T] {
	return Key[T]{
		id: nextKeyID.Add(1),
	}
}

type valueStore struct {
	mu     sync.RWMutex
	values map[uint64]any
}

func newValueStore() valueStore {
	return valueStore{
		values: make(map[uint64]any),
	}
}

func (s *valueStore) set(id uint64, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.values[id] = value
}

func (s *valueStore) get(id uint64) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.values[id]
	return value, ok
}

func (s *valueStore) delete(id uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.values, id)
}

// Set armazena um valor tipado.
func Set[T any](ctx *Context, key Key[T], value T) {
	if ctx == nil || key.id == 0 {
		return
	}

	ctx.values.set(key.id, value)
}

// Get recupera um valor tipado.
func Get[T any](ctx *Context, key Key[T]) (T, bool) {
	var zero T

	if ctx == nil || key.id == 0 {
		return zero, false
	}

	value, ok := ctx.values.get(key.id)
	if !ok {
		return zero, false
	}

	result, ok := value.(T)
	if !ok {
		return zero, false
	}

	return result, true
}

// Delete remove um valor armazenado.
func Delete[T any](ctx *Context, key Key[T]) {
	if ctx == nil || key.id == 0 {
		return
	}

	ctx.values.delete(key.id)
}
