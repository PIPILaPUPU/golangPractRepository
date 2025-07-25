package main

import "fmt"

// Queue описывает очередь.
type Queue struct {
	first *QueueItem // указатель на первый элемент очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	person *Character // указатель на персонажа
	next   *QueueItem // указатель на следующий элемент
}

// Character описывает персонажа игры.
type Character struct {
	Name  string
	Level int
}

// Pop удаляет первый элемент из очереди и возвращает указатель на персонажа.
func (queue *Queue) Pop() (*Character, bool) {
	// вставьте недостающий код
	if queue.first == nil {
		return nil, false
	}

	temp := queue.first.person
	queue.first = queue.first.next
	return temp, true
}

// Push добавляет в конец очереди элемент с указанным персонажем.
func (queue *Queue) Push(person *Character) {
	// вставьте недостающий код
	newItem := &QueueItem{
		person: person,
		next:   nil,
	}

	if queue.first == nil {
		queue.first = newItem
		return
	}

	ptr := queue.first
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = newItem
}

func main() {
	list := []*Character{
		{"царь", 90}, {"царевич", 50},
		{"король", 80}, {"королевич", 40}}

	queue := &Queue{}
	for _, v := range list {
		queue.Push(v)
	}

	for {
		v, ok := queue.Pop()
		if !ok {
			break
		}
		fmt.Println(*v)
	}
}
