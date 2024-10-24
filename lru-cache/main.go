package main

func main() {

}

type LRUItem struct {
	key int
	value int
	next *LRUItem
	prev *LRUItem
}

type LRUCache struct {
    data map[int]*LRUItem
	head *LRUItem
	tail *LRUItem
	dataMaxCap int
}

func (l *LRUCache) AddToBegin(lruItem *LRUItem) {
	l.head = lruItem
	l.tail = lruItem
}

func (l *LRUCache) AddToBeginOfList(lruItem *LRUItem) {
	lruItem.next = l.head
	l.head.prev = lruItem
	l.head = lruItem
}

func (l *LRUCache) MoveToBegin(lruItem *LRUItem) {
	if lruItem.next == nil {
		lruItem.prev.next = nil
		l.tail = lruItem.prev
	}

	lruItem.next = l.head
	lruItem.prev = nil
	lruItem.next.prev = lruItem
	
	l.head = lruItem
}

func (l *LRUCache) DeleteItem(lruItem *LRUItem) {
	if lruItem.next == nil {
		l.tail = lruItem.prev
		lruItem.prev.next = nil
		lruItem.prev.prev = lruItem
	} else {
		lruItem.next.prev = lruItem.prev

		if lruItem.prev == nil {
			l.head = lruItem.next
		} else {
			lruItem.prev.next = lruItem.next
		}
	}
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
		data: make(map[int]*LRUItem, capacity),
		dataMaxCap: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
    if lruItem, found := l.data[key]; found {
		l.DeleteItem(lruItem)
		l.MoveToBegin(lruItem)

		return lruItem.value
	}

	return -1
}


func (l *LRUCache) Put(key int, value int) {
	if lruItem, found := l.data[key]; found {
		lruItem.value = value
		l.DeleteItem(lruItem)
		l.MoveToBegin(lruItem)
	} else {
		if len(l.data) == l.dataMaxCap {
			delete(l.data, l.tail.key)
			l.tail.prev.next = nil
			l.tail = l.tail.prev
		} else {
			newItem := &LRUItem{
				key: key,
				value: value,
			}
			if len(l.data) == 0 {
				l.AddToBegin(newItem)
			} else {
				l.AddToBeginOfList(newItem)
			}

			l.data[key] = newItem
		}
	}
}
