# Qoala Cache Library
Qoala's Cache

# How to Use LRUCache
```
import "github.com/qoala-engineering/qoala-common/lrucache"

func main() {
    // some code

    ....
    // create log entry
    key := "sample_key"
    data := "sample data"

    l := New(10)
    l.Set(key, data)

    val, ok := l.Get(key)
    ....

    // some code
}
```

# Author
[Windy Hendra Supardi](https://github.com/windyhendra)


# How to Use Redis
```
import "github.com/qoala-engineering/qoala-common/redis"

func main() {
    // some code

    ....
    // create log entry
    key := "sample_key"
    data := "sample data"

    redisCache := redis.NewRedisCache(ac.RedisSession)

	respCache := redisCache.GetRedisValue(aws.String(key))
	if respCache != nil {
        ..
	}

    redisCache.SetRedisValue(aws.String(key), aws.String(data))

    ....

    // some code
}
```