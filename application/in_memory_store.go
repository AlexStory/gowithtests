package main

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}

type InMemoryStore struct {
	store map[string]int
}

func (i *InMemoryStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.store[name]
}