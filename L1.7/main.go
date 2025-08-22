package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SafeMap) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *SafeMap) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func (s *SafeMap) Increment(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key]++
}

func main() {
	sm := NewSafeMap()

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key-%d", i), i)
		}(i)
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Increment(fmt.Sprintf("key-%d", i%25)) 
		}(i)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		if val, ok := sm.Get(fmt.Sprintf("key-%d", i)); ok {
			fmt.Printf("Before delete: key-%d = %d\n", i, val)
		}
		sm.Delete(fmt.Sprintf("key-%d", i))
	}

	fmt.Println("Keys after deletion:", sm.Keys())
	fmt.Println("Length after deletion:", sm.Len())

	sm.Increment("new-key")
	val, _ := sm.Get("new-key")
	fmt.Println("new-key after increment:", val)

	fmt.Println("Final keys and values:")
	for _, k := range sm.Keys() {
		v, _ := sm.Get(k)
		fmt.Printf("%s = %d\n", k, v)
	}
}
