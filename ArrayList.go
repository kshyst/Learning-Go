package main

type ArrayList[T comparable] struct {
	prev  *ArrayList[T]
	next  *ArrayList[T]
	last  *ArrayList[T]
	first *ArrayList[T]
	val   T
}

func createArraylist[T comparable](initVal T) *ArrayList[T] {
	var list = &ArrayList[T]{val: initVal}
	list.first = list
	list.last = list

	return list
}

func (list *ArrayList[T]) add(val T) {
	if list == nil {
		list = &ArrayList[T]{val: val}
		list.first = list
		list.last = list
		return
	}

	list.last.next = &ArrayList[T]{val: val, prev: list.last}
	list.last = list.last.next
}

func (list *ArrayList[T]) remove(val T) {
	if list == nil {
		return
	}

	if list.val == val {
		if list.prev != nil {
			list.prev.next = list.next
		}

		if list.next != nil {
			list.next.prev = list.prev
		}
		return
	}

	list.next.remove(val)
}

func (list *ArrayList[T]) removeAll(val T) {
	if list == nil {
		return
	}

	if list.val == val {
		if list.prev != nil {
			list.prev.next = list.next
		}

		if list.next != nil {
			list.next.prev = list.prev
		}
	}

	list.next.removeAll(val)
}

func (list *ArrayList[T]) get(index int) T {
	if index == 0 {
		return list.val
	}

	return list.next.get(index - 1)

}

func (list *ArrayList[T]) size() int {
	if list == nil {
		return 0
	}

	return 1 + list.next.size()
}

func (list *ArrayList[T]) printList() []T {
	if list == nil {
		return []T{}
	}

	return append([]T{list.val}, list.next.printList()...)
}
