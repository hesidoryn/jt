# JT

## MASTER BRANCH IS DEPRECATED. PLEASE GO TO [REFACTOR](https://github.com/hesidoryn/jt/tree/refactor) BRANCH.

## Installation
[here](https://github.com/hesidoryn/jt/blob/master/DEPLOYMENT.md)

## Quickstart
```go
func SomeExamples() {
	client, err := jtclient.NewClient("client-config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = client.Set("k", "value")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resGet, err := client.Get("k")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(resGet)
}
```

## Examples
[here](https://github.com/hesidoryn/jt/tree/master/_examples)

## Some documentation
[here](https://godoc.org/github.com/hesidoryn/jt)

## Some tests
[here](https://github.com/hesidoryn/jt/tree/master/storage)

## Commands list
1. SET
2. GET
3. APPEND
4. GETSET
5. STRLEN
6. INCR
7. INCRBY
8. LPUSH
9. RPUSH
10. LPOP
11. RPOP
12. LREM
13. LINDEX
14. LRANGE
15. LLEN
16. DSET (Redis HMSET)
17. DGET (Redis HGET)
18. DDEL (Redis HDEL)
19. DEXISTS (Redis HEXISTS)
20. DLEN (Redis HLEN)
21. DINCRBY (Redis HINCRBY)
22. DINCRBYFLOAT (Redis HINCRBYFLOAT)
23. DELETE
24. RENAME
25. PERSIST
26. EXPIRE
27. TTL
28. TYPE
29. KEYS
30. EXISTS
31. PING
32. AUTH
33. SAVE

## Benchmarks
[here](https://github.com/hesidoryn/jt/blob/master/BENCHMARKS.md)
