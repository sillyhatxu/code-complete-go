package _46_LRU_Cache

import "fmt"

type LRUCache struct {
	first, last *Node
	cache       map[int]*Node
	capacity    int
}

type Node struct {
	key      int
	val      int
	previous *Node
	next     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*Node), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	current, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.refresh(current)
	return current.val
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.cache) == 0 {
		current := &Node{key: key, val: value}
		current.previous, current.next = current, current
		this.first = current
		this.last = current
		this.cache[key] = current
		return
	}
	if current, ok := this.cache[key]; ok {
		current.val = value
		this.refresh(current)
		return
	}
	current := &Node{key: key, val: value}
	if len(this.cache) == this.capacity {
		current.previous, current.next = this.last, this.first.next
	} else {
		current.previous, current.next = this.last, this.first
	}

	this.last.next = current
	this.first.previous = current
	this.last = current
	this.cache[key] = current
	this.refresh(current)
	//this.last.next, current.previous, this.last = current, this.last, current
}

func (this *LRUCache) refresh(current *Node) {
	if current == this.first {
		this.last.next, current.previous, current.next, this.last = current, this.last, this.first.next, current
		this.first = this.first.next
	} else if current == this.last {

	} else {
		current.previous.next, current.next.previous = current.next, current.previous
		current.next, current.previous = this.first, this.last
		this.first.previous, this.last.next = current, current
		this.last = current
	}

	//if current.previous == nil && current.next != nil {
	//	this.first = current.next
	//	this.last.next = current
	//	this.last = current
	//} else if current.previous != nil && current.next != nil {
	//	current.previous.next = current.next
	//	this.last.next = current
	//	this.last = current
	//}
	if len(this.cache) > this.capacity {
		removeKey := this.first.key
		this.first = this.first.next
		delete(this.cache, removeKey)
	}
	this.forloop(this.first)
}

func (this *LRUCache) forloop(current *Node) {
	if current == nil {
		return
	}
	test := ""
	for {
		test += fmt.Sprintf("%d -> ", current.key)
		if current.next == this.first {
			break
		}
		current = current.next
	}
	fmt.Println(test)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
