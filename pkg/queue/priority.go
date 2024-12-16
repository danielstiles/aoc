package queue

type priorityItem[T any] struct {
	Priority int
	Value    T
}

type PriorityQueue[T any] struct {
	queue []priorityItem[T]
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.queue)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	return q.queue[i].Priority < q.queue[j].Priority
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.queue[i], q.queue[j] = q.queue[j], q.queue[i]
}

func (q *PriorityQueue[T]) Push(priority int, value T) {
	q.queue = append(q.queue, priorityItem[T]{
		Priority: priority,
		Value:    value,
	})
	q.up(q.Len() - 1)
}

func (q *PriorityQueue[T]) Pop() (ret T) {
	q.Swap(0, q.Len()-1)
	ret = q.queue[q.Len()-1].Value
	q.queue = q.queue[:q.Len()-1]
	q.down(0)
	return
}

func (q *PriorityQueue[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !q.Less(i, parent) {
			break
		}
		q.Swap(i, parent)
		i = parent
	}
}

func (q *PriorityQueue[T]) down(i int) {
	for {
		child := 2*i + 1
		if child >= q.Len() {
			break
		}
		if right := child + 1; right < q.Len() && q.Less(right, child) {
			child = right
		}
		if !q.Less(child, i) {
			break
		}
		q.Swap(i, child)
		i = child
	}
}
