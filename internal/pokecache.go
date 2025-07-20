package internal
import(
	"time"
	"sync"
)

type Cache struct{
mu sync.Mutex
cache map[string]CacheEntry
}

type CacheEntry struct{
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration)  *Cache{
	cache := Cache{ mu: sync.Mutex{},cache: make(map[string]CacheEntry)}
	go cache.ReapLoop(interval)
	return &cache

}
func (c *Cache) Add (key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key]= CacheEntry{time.Now(),val}
	

}

func (c *Cache) Get (key string) ([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	val,ok :=c.cache[key]
	if !ok{
		return []byte{},false
	}
	return val.val , ok
}

func (c *Cache) ReapLoop (interval time.Duration){
	for range time.Tick(interval){
		c.mu.Lock()
		
		var keysToDelete []string
		for i, value := range c.cache{
			if time.Since(value.createdAt) >= interval{
				keysToDelete = append(keysToDelete, i)
			} 
		}
	for _, key := range keysToDelete{
		delete(c.cache,key)
	}
	c.mu.Unlock()
	}

}