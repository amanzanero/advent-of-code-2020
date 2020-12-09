package nqueue

type NQueue struct {
	size, cap int
	queue     [][]int
	lookup    map[int]int // key: sum value: count
}

func NewNQueue(n int) *NQueue {
	return &NQueue{
		size: 0, cap: n, queue: make([][]int, 0), lookup: make(map[int]int),
	}

}

func (nq *NQueue) push(value []int) {
	nq.size += 1
	nq.queue = append(nq.queue, value)

	for _, v := range value {
		_, exists := nq.lookup[v]
		if !exists {
			nq.lookup[v] = 0
		}
		nq.lookup[v] += 1
	}
}

func (nq *NQueue) pop() []int {
	front := nq.queue[0]
	nq.queue = nq.queue[1:]
	nq.size -= 1

	for _, num := range front {
		nq.lookup[num] -= 1
		if nq.lookup[num] == 0 {
			delete(nq.lookup, num)
		}
	}

	return front
}

func (nq *NQueue) PushPop(value []int) {
	if nq.size == nq.cap {
		nq.pop()
	}
	nq.push(value)
}

func (nq *NQueue) Has(value int) bool {
	_, exists := nq.lookup[value]
	return exists
}
