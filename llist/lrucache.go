package llist

type dllnode struct {
	key []byte
    val []byte
    next *dllnode
    prev *dllnode
} 

type LRUCache struct {
	head 		*dllnode   // least recently used head
	tail 		*dllnode   // most recently used tail
	keyToNode	map[string]*dllnode // map of key to node in the dll
	capacity	int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache {
		head: nil,
		tail: nil,
		keyToNode: make(map[string]*dllnode),
		capacity: capacity,
	}
}

func (lru *LRUCache)addToTail(node *dllnode) {
	if lru.tail == nil {
		// If empty, set both head and tail
		lru.head, lru.tail = node, node
		return
	}
	node.prev = lru.tail // connect this nodes tx to last nodes
	node.next=nil
	lru.tail.next = node // last nodes tx points to this node
	lru.tail=node // finaly make this node as last node/ tail node
}

func (lru *LRUCache)removeNode(node *dllnode) {
	if node == nil {
		return
	}

	if node.prev == nil {
		// first node
		lru.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		// last node
		lru.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}

func (lru *LRUCache)get(key []byte) (val []byte) {
	node, ok := lru.keyToNode[string(key)]
	if !ok {
		return
	}
	// remove and add to tail
	lru.removeNode(node)
	lru.addToTail(node)
	return node.val
}

func (lru *LRUCache)put(key, val []byte) {
	node, ok := lru.keyToNode[string(key)]
	if ok {
		lru.removeNode(node)
	}
	node = &dllnode{
		key: key,
		val: val,
		next: nil,
		prev: nil,
	}
	lru.keyToNode[string(key)] = node
	if len(lru.keyToNode)>lru.capacity {
		delete(lru.keyToNode, string(lru.head.key)) // Remove key from map
		lru.removeNode(lru.head) // remove least recently used if overflows
	}
	lru.addToTail(node)
}