package hw04lrucache

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (l *ListItem) isFirst() bool {
	return l.Prev == nil
}

func (l *ListItem) isLast() bool {
	return l.Next == nil
}

func (l *ListItem) isInside() bool {
	return !l.isLast() && !l.isFirst()
}

func NewListItem(i interface{}) *ListItem {
	return &ListItem{i, nil, nil}
}
