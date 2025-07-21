package internal

import(
	"time"
	"sync"
)


// Interval of Time before cache reaping
const INTERVAL time.Duration = 5 * time.Second



// Cache Struct that holds mutex and the cache map
type Cache struct{
Mu sync.Mutex
Cache map[string]CacheEntry
}


// Cache value entry at cachemap
type CacheEntry struct{
	createdAt time.Time //time cacheentry was created, for internal use by Cache Reaping System
	val []byte    // actual cache value stored as bytes
}

//func to create cache based on the interval, concurrently executes ReapLoop
func NewCache(interval time.Duration)  *Cache{
	cache := Cache{ Mu: sync.Mutex{},Cache: make(map[string]CacheEntry)}
	go cache.ReapLoop(interval)
	return &cache

}

//Safe Way to add to cachemap
func (c *Cache) Add (key string, val []byte){
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cache[key]= CacheEntry{time.Now(),val}
	

}

//safe way to get cache value based on key string, returns data and bool ok value
func (c *Cache) Get (key string) ([]byte,bool){
	c.Mu.Lock()
	defer c.Mu.Unlock()
	val,ok :=c.Cache[key]
	if !ok{
		return []byte{},false
	}
	return val.val , ok
}


// reap loop to clear cache to make sure it doesn't get too big
func (c *Cache) ReapLoop (interval time.Duration){

	//for every time interval passes....
	for range time.Tick(interval){
		c.Mu.Lock()
		
		var keysToDelete []string
		for i, value := range c.Cache{
			if time.Since(value.createdAt) >= interval{
				keysToDelete = append(keysToDelete, i)
			} 
		}
	for _, key := range keysToDelete{
		delete(c.Cache,key)
	}
	c.Mu.Unlock()
	}

}