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

type list struct {
	length int
	first  *ListItem
	last   *ListItem
}

func (list *list) Len() int {
	return list.length
}

func (list *list) Front() *ListItem {
	return list.first
}

func (list *list) Back() *ListItem {
	return list.last
}

func (list *list) PushFront(v interface{}) *ListItem {
	listItem := NewListItem(v)
	if list.length == 0 {
		list.last = listItem
	} else {
		list.first.Prev = listItem
		listItem.Next = list.first
	}
	list.first = listItem
	list.length++
	return listItem
}

func (list *list) PushBack(v interface{}) *ListItem {
	newListItem := NewListItem(v)
	if list.length == 0 {
		newListItem.Prev = list.first
		list.first = newListItem
	} else {
		newListItem.Prev = list.last
		list.last.Next = newListItem
	}
	list.last = newListItem
	list.length++
	return newListItem
}

func (list *list) Remove(listItem *ListItem) {
	switch {
	case list.length == 1:
		list.first = nil
		list.last = nil
	case listItem.isFirst():
		list.first = listItem.Next
		list.first.Prev = nil
	case listItem.isLast():
		list.last = listItem.Prev
		list.last.Next = nil
	case listItem.isInside():
		listItem.Prev.Next = listItem.Next
		listItem.Next.Prev = listItem.Prev
	}
	list.length--
}

func (list *list) MoveToFront(listItem *ListItem) {
	if list.length == 1 || listItem.isFirst() {
		return
	}
	list.Remove(listItem)
	list.PushFront(listItem.Value)
}

func NewList() List {
	return new(list)
}
