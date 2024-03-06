package components

import (
	aes "aes_go/aes"
)

type Message struct {
	BlockNum uint32
	Block    aes.Block
}

type MessageHeap []Message

func (h MessageHeap) Len() int           { return len(h) }

func (h MessageHeap) Less(i, j int) bool { return h[i].BlockNum < h[j].BlockNum }

func (h MessageHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MessageHeap) down(i, n int) {
	for {
		l := 2*i + 1
		if l >= n || l < 0 { // avoid underflow
			break
		}
		j := l // left child
		if r := l + 1; r < n && !h.Less(l, r) {
			j = r // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *MessageHeap) Heapify() {
	for i := len(*h)/2 - 1; i >= 0; i-- {
		(*h).down(i, len(*h))
	}
}

func (h *MessageHeap) Push(x Message) {
	*h = append(*h, x)
	h.Heapify()
}

func (h *MessageHeap) Pop() Message {
	n := len(*h) - 1
	h.Swap(0, n)
	h.down(0, n)
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *MessageHeap) Peek() Message {
	return (*h)[0]
}