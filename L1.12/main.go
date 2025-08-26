package main

import (
	"fmt"
	"sync"
)

type SafeSet[T comparable] struct {
	elements map[T]struct{}
	mu       sync.RWMutex
}

func NewSafeSet[T comparable]() *SafeSet[T] {
	return &SafeSet[T]{elements: make(map[T]struct{})}
}

func (s *SafeSet[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements[item] = struct{}{}
}

func (s *SafeSet[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.elements, item)
}

func (s *SafeSet[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.elements[item]
	return exists
}

func (s *SafeSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.elements)
}

func (s *SafeSet[T]) Print() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.elements {
		fmt.Println(k)
	}
}

func main() {
	set := NewSafeSet[string]()
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, val := range sl {
		set.Add(val)
	}

	set.Print()

}
