package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	InsertBefore(i *ListItem, other *ListItem)
	PopBack() interface{}
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Clear()
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	size  int
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{v, nil, l.front}
	// update the current front
	if l.front != nil {
		l.front.Prev = item
	}
	// replace the list's front and back (if needed), update the size
	l.front = item
	if l.size == 0 {
		l.back = item
	}
	l.size++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{v, l.back, nil}
	// update the current back
	if l.back != nil {
		l.back.Next = item
	}
	// replace the list's back and front (if needed), update the size
	l.back = item
	if l.size == 0 {
		l.front = item
	}
	l.size++
	return item
}

func (l *list) InsertBefore(i *ListItem, other *ListItem) {
	if other == nil {
		// when i - becomes the first element
		l.front = i
		l.back = i
		i.Prev, i.Next = nil, nil
	} else {
		// put the i before the other
		i.Next = other
		i.Prev = other.Prev
		if l.front == other {
			l.front = i
		}
		other.Prev = i
	}
	l.size++
}

func (l *list) PopBack() interface{} {
	item := l.Back()
	l.Remove(item)
	return item.Value
}

func (l *list) Remove(i *ListItem) {
	// update the element's siblings
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	// update the list's front/back if needed (also works if the element was the last one)
	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}

	l.size--
	// update the element itself
	i.Prev, i.Next = nil, nil
}

func (l *list) MoveToFront(i *ListItem) {
	if l.size == 1 || i == l.front {
		// nothing to do
		return
	}
	l.Remove(i)
	l.InsertBefore(i, l.Front())
}

func (l *list) Clear() {
	// unlink each and every item
	for item := l.front; item != nil; {
		next := item.Next
		item.Prev, item.Next = nil, nil
		item = next
	}
	l.back, l.front = nil, nil
	l.size = 0
}

func NewList() List {
	return &list{nil, nil, 0}
}
