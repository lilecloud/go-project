package cache

type Cache struct {
	m map[string]*interface{}
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]*interface{}),
	}
}

// add element to cache
func (c *Cache) Put(key string, val interface{}) bool {
	if _, ok := c.m[key]; ok {
		return true
	}

	c.m[key] = &val

	return true
}

// get element from cache
func (c *Cache) Get(key string) (interface{}, bool) {
	if val, ok := c.m[key]; ok {
		return val, true
	}
	return nil, false
}

func (c *Cache) Del(key string) (interface{}, bool) {
	if val, ok := c.m[key]; ok {
		delete(c.m, key)
		return val, true
	}
	return nil, false
}
