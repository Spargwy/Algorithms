package main

type List struct {
	root Element
	len  int
}

type Element struct {
	next, prev *Element

	list *List

	Value interface{}
}

type OutOfRangeError struct{}

func (e *OutOfRangeError) Error() string {
	return "Index out of range len"
}

func New() *List { return new(List).Init() }

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.list = l
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) PushHead(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushTail(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List) Head() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Tail() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (l *List) Search(value interface{}) *Element {
	x := l.Head()
	for x != nil && x.Value != value {
		x = x.next
	}
	return x
}

func (l *List) Delete(e *Element) interface{} {
	if e.list == l {
		l.delete(e)
	}
	return e.Value
}

func (l *List) delete(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) DeleteByIndex(i int) (interface{}, error) {
	if l.len <= i {
		return nil, &OutOfRangeError{}
	}
	e := l.Head()
	for j := 0; j <= i; j++ {
		if j == i {
			return (l.Delete(e)), nil
		}
		e = e.Next()
	}
	return nil, nil
}
