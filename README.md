# Redis client for Golang

[![Build Status](https://travis-ci.org/hustfisher/redis.png?branch=master)](https://travis-ci.org/hustfisher/redis)
[![GoDoc](https://godoc.org/github.com/hustfisher/redis?status.svg)](https://godoc.org/github.com/hustfisher/redis)
[![Airbrake](https://img.shields.io/badge/kudos-airbrake.io-orange.svg)](https://airbrake.io)

Supports:

- Redis 3 commands except QUIT, MONITOR, SLOWLOG and SYNC.
- Automatic connection pooling with [circuit breaker](https://en.wikipedia.org/wiki/Circuit_breaker_design_pattern) support.
- [Pub/Sub](https://godoc.org/github.com/hustfisher/redis#PubSub).
- [Transactions](https://godoc.org/github.com/hustfisher/redis#example-Client-TxPipeline).
- [Pipeline](https://godoc.org/github.com/hustfisher/redis#example-Client-Pipeline) and [TxPipeline](https://godoc.org/github.com/hustfisher/redis#example-Client-TxPipeline).
- [Scripting](https://godoc.org/github.com/hustfisher/redis#Script).
- [Timeouts](https://godoc.org/github.com/hustfisher/redis#Options).
- [Redis Sentinel](https://godoc.org/github.com/hustfisher/redis#NewFailoverClient).
- [Redis Cluster](https://godoc.org/github.com/hustfisher/redis#NewClusterClient).
- [Cluster of Redis Servers](https://godoc.org/github.com/hustfisher/redis#example-NewClusterClient--ManualSetup) without using cluster mode and Redis Sentinel.
- [Ring](https://godoc.org/github.com/hustfisher/redis#NewRing).
- [Instrumentation](https://godoc.org/github.com/hustfisher/redis#ex-package--Instrumentation).
- [Cache friendly](https://github.com/hustfisher/cache).
- [Rate limiting](https://github.com/hustfisher/redis_rate).
- [Distributed Locks](https://github.com/bsm/redislock).

API docs: https://godoc.org/github.com/hustfisher/redis.
Examples: https://godoc.org/github.com/hustfisher/redis#pkg-examples.

## Installation

Install:

```shell
go get -u github.com/hustfisher/redis
```

Import:

```go
import "github.com/hustfisher/redis"
```

## Quickstart

```go
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
```

## Howto

Please go through [examples](https://godoc.org/github.com/hustfisher/redis#pkg-examples) to get an idea how to use this package.

## Look and feel

Some corner cases:

```go
// SET key value EX 10 NX
set, err := client.SetNX("key", "value", 10*time.Second).Result()

// SORT list LIMIT 0 2 ASC
vals, err := client.Sort("list", redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()

// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
vals, err := client.ZRangeByScoreWithScores("zset", redis.ZRangeBy{
	Min: "-inf",
	Max: "+inf",
	Offset: 0,
	Count: 2,
}).Result()

// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
vals, err := client.ZInterStore("out", redis.ZStore{Weights: []int64{2, 3}}, "zset1", "zset2").Result()

// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
vals, err := client.Eval("return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

// custom command
res, err := client.Do("set", "key", "value")
```

## See also

- [Golang PostgreSQL ORM](https://github.com/go-pg/pg)
- [Golang msgpack](https://github.com/vmihailenco/msgpack)
- [Golang message task queue](https://github.com/vmihailenco/taskq)
