package v3

type LRUCache struct {
	lruCache map[int]*Node
	first    *Node
	last     *Node
	capacity int
}

type Node struct {
	previous *Node
	next     *Node
	value    int
	key      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lruCache: map[int]*Node{},
		first:    nil,
		last:     nil,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.lruCache[key]
	if !ok {
		return -1
	}
	if val == this.first {
		return val.value
	}
	if val == this.last {
		this.last = this.last.previous
		if this.last != nil {
			this.last.next = nil
		}
		this.first.previous = val
		val.next = this.first
		val.previous = nil
		this.first = val
		return val.value
	}

	if this.first != nil && this.last != nil && this.last != this.first {
		// check if last is val
		prevHead := this.first
		if val == this.last {
			this.last = this.last.previous
		}
		if val != this.first {
			this.first.previous = val
		}

		nodeBelow := val.next
		nodeAbove := val.previous
		if nodeBelow != nil && nodeAbove != nil {
			nodeBelow.previous = nodeAbove
			nodeAbove.next = nodeBelow
		}

		this.first = val
		val.previous = nil
		val.next = prevHead
	}
	return val.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 1 && this.first != nil {
		delete(this.lruCache, this.first.key)
		newNode := &Node{
			previous: nil,
			next:     nil,
			value:    value,
			key:      key,
		}
		this.lruCache[key] = newNode
		this.first = newNode
		this.last = newNode
		return
	}
	if val := this.Get(key); val != -1 {
		this.lruCache[key].value = value
		return
	}

	newNode := &Node{
		previous: nil,
		next:     this.first,
		value:    value,
		key:      key,
	}
	if this.last == nil && this.first == nil {
		this.last = newNode
		this.first = newNode
		this.lruCache[key] = newNode
		return
	}
	if len(this.lruCache) == this.capacity && this.last != nil {
		// evict the last
		prevTail := this.last
		prevTailKey := this.last.key
		this.last = prevTail.previous
		if this.last != nil {
			this.last.next = nil
		}
		delete(this.lruCache, prevTailKey)
	}
	this.first.previous = newNode
	this.first = newNode
	this.lruCache[key] = newNode
}
