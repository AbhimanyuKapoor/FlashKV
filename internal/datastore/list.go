package datastore

type List struct {
	items []string
}

func NewList() *List {
	return &List{
		items: make([]string, 0),
	}
}

func (l *List) LPush(value string) {

	l.items = append([]string{value}, l.items...)
}

func (l *List) RPush(value string) {

	l.items = append(l.items, value)
}

func (l *List) LPop() (string, bool) {

	if len(l.items) == 0 {
		return "", false
	}

	value := l.items[0]

	l.items = l.items[1:]

	return value, true
}

func (l *List) RPop() (string, bool) {

	if len(l.items) == 0 {
		return "", false
	}

	last := len(l.items) - 1

	value := l.items[last]

	l.items = l.items[:last]

	return value, true
}

func (l *List) Len() int {

	return len(l.items)
}

func (l *List) Range() []string {

	result := make([]string, len(l.items))

	copy(result, l.items)

	return result
}
