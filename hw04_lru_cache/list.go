package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	tail *ListItem
	head *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

// добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	newNode := &ListItem{
		Value: v,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		curHead := l.head
		newNode.Next = curHead
		curHead.Prev = newNode
		l.head = newNode
	}
	l.len++

	return newNode
}

// добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	newNode := &ListItem{
		Value: v,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		curTail := l.tail
		newNode.Prev = curTail
		curTail.Next = newNode
		l.tail = newNode
	}
	l.len++

	return newNode
}

// удалить элемент
func (l *list) Remove(i *ListItem) {
	if i == l.head && i == l.tail {
		l.head = nil
		l.tail = nil
	} else if i == l.head {
		l.head = i.Next
		l.head.Prev = nil
	} else if i == l.tail {
		l.tail = i.Prev
		l.tail.Next = nil
	} else {
		next := i.Next
		prev := i.Prev
		next.Prev = prev
		prev.Next = next
	}
	l.len--
}

// переместить элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}
	if l.tail == i {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else {
		next := i.Next
		prev := i.Prev
		next.Prev = prev
		prev.Next = next
	}
	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}

func NewList() List {
	return &list{}
}
