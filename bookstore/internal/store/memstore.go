// internal/store/memstore.go

package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"fmt"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (ms *MemStore) Create(b *mystore.Book) error {
	ms.books[b.Id] = b
	return nil
}

func (ms *MemStore) Update(b *mystore.Book) error {
	ms.books[b.Id] = b
	return nil
}

func (ms *MemStore) Get(id string) (mystore.Book, error) {
	return *ms.books[id], nil

}

func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	// arr := []mystore.Book{}
	fmt.Println("get All")
	arr := make([]mystore.Book, 0)
	for _, val := range ms.books {
		arr = append(arr, *val)
	}

	fmt.Println("arr", arr)
	return arr, nil
}

func (ms *MemStore) Delete(id string) error {

	delete(ms.books, id)
	return nil
}
