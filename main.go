package ttlcache

import (
	"strconv"
	"time"
)

func main() {
	cache := NewCache()
	defer cache.Close()
	cache.SetTTL(time.Duration(10000000))
	count := 5
	for i := 0; i < count; i++ {
		val := NewCache()
		defer val.Close()

		cache.Set(strconv.FormatInt(int64(i), 10), val)
	}
}
