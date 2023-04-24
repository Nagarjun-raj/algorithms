package heap

type Heap struct {
	data []int
}

func New() *Heap {
	return &Heap{make([]int, 0)}
}

func (h *Heap) findParent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) findLeftChild(i int) int {
	return 2*i + 1
}

func (h *Heap) findRightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) swap(i, j int) {
	(*h).data[i], (*h).data[j] = (*h).data[j], (*h).data[i]
}

func (h *Heap) Push(ele int) {
	(*h).data = append((*h).data, ele)
	h.heapifyUp(len((*h).data) - 1)
}

func (h *Heap) Pop() int {
	n := len((*h).data)
	min := (*h).data[0]
	(*h).data[0] = (*h).data[n-1]
	(*h).data = (*h).data[:n-1]
	h.heapifyDown(0)
	return min
}

func (h *Heap) heapifyUp(i int) {
	for (*h).data[i] < (*h).data[(*h).findParent(i)] && i > 0 {
		(*h).swap(i, (*h).findParent(i))
		i = (*h).findParent(i)
	}
}

func (h *Heap) heapifyDown(i int) {
	for {
		left := (*h).findLeftChild(i)
		right := (*h).findRightChild(i)
		smallest := i
		if left < len((*h).data) && (*h).data[left] < (*h).data[smallest] {
			smallest = left
		}
		if right < len((*h).data) && (*h).data[right] < (*h).data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		(*h).swap(i, smallest)
		i = smallest
	}
}
