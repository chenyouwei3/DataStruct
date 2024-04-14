package main

import "container/list"

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K, V int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if element, ok := c.Keys[key]; ok {
		c.List.MoveToFront(element)
		return element.Value.(pair).V
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if element, ok := c.Keys[key]; ok {
		element.Value = pair{K: key, V: value}
		c.List.MoveToFront(element)
	} else {
		element := c.List.PushFront(pair{K: key, V: value})
		c.Keys[key] = element
	}
	if c.List.Len() > c.Cap {
		element := c.List.Back()
		c.List.Remove(element)
		delete(c.Keys, element.Value.(pair).K)
	}
}

func (c *LRUCache) Add(key int, value int) {
	if _, ok := c.Keys[key]; !ok {
		element := c.List.PushFront(value)
		c.Keys[key] = element
	}
}

func (c *LRUCache) Remove(node *list.Element) {
	if node != nil {
		c.List.Remove(node)
		delete(c.Keys, node.Value.(pair).V)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
