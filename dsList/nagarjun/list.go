package nagarjun

import "errors"

type List struct {
	size int
	data []interface{}
}

func New(size int) *List {
	return &List{
		size: size,
		data: make([]interface{}, 0),
	}
}

func (l *List) SizeOfList() int {
	return len(l.data)
}

func (l *List) PushBack(ele interface{}) error {
	if l.SizeOfList() == l.size {
		return errors.New("list is full")
	}
	l.data = append(l.data, ele)
	return nil
}

func (l *List) PushFront(ele interface{}) error {
	if l.SizeOfList() == l.size {
		return errors.New("list is full")
	}
	l.data = append([]interface{}{ele}, l.data...)
	return nil
}

func (l *List) InsertAfter(ele interface{}, index int) (interface{}, error) {
	if index < 0 || index >= l.SizeOfList() {
		return nil, errors.New("invalid index")
	}
	if l.SizeOfList() == l.size {
		return nil, errors.New("list is full")
	}
	if index == l.size-1 {
		return nil, errors.New("index out of range")
	}
	desired := index + 1
	l.data = append(l.data[:desired], append([]interface{}{ele}, l.data[desired:]...)...)
	return ele, nil
}

func (l *List) IsEmpty() bool {
	return l.SizeOfList() == 0
}

func (l *List) Remove(index int) (interface{}, error) {
	if index > l.size-1 {
		return nil, errors.New("invalid Index")
	} else if index == 0 {
		k := l.data[0]
		l.data = l.data[1:]
		return k, nil
	} else if index == l.size-1 {
		m := l.data[l.size-1]
		l.data = l.data[:index]
		return m, nil
	} else {
		n := l.data[index]
		l.data = append(l.data[:index], l.data[index+1:]...)
		return n, nil

	}
}
