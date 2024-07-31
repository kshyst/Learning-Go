package main

type ArrayList[T comparable] struct {
	prev  *ArrayList[T]
	next  *ArrayList[T]
	last  *ArrayList[T]
	first *ArrayList[T]
	val   T
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
