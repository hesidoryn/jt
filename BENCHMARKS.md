## TEST #1 - Servers performance test
### redis-benchmark -t ping_inline,set,get,incr,lpop,rpop,lpush,rpush,lrange -n 100000 -q

JT:
```shell
PING_INLINE: 121065.38 requests per second
SET: 57471.27 requests per second
GET: 65402.22 requests per second
INCR: 64557.78 requests per second
LPUSH: 59701.50 requests per second
RPUSH: 58719.91 requests per second
LPOP: 65402.22 requests per second
RPOP: 65274.15 requests per second
LPUSH (needed to benchmark LRANGE): 60096.15 requests per second
LRANGE_100 (first 100 elements): 53850.30 requests per second
LRANGE_300 (first 300 elements): 51572.98 requests per second
LRANGE_500 (first 450 elements): 51599.59 requests per second
LRANGE_600 (first 600 elements): 50251.26 requests per second
```

Redis:
```shell
PING_INLINE: 148367.95 requests per second
SET: 146627.56 requests per second
GET: 160513.64 requests per second
INCR: 161812.31 requests per second
LPUSH: 167504.19 requests per second
RPUSH: 168634.06 requests per second
LPOP: 166112.95 requests per second
RPOP: 163132.14 requests per second
LPUSH (needed to benchmark LRANGE): 167224.08 requests per second
LRANGE_100 (first 100 elements): 58823.53 requests per second
LRANGE_300 (first 300 elements): 25542.79 requests per second
LRANGE_500 (first 450 elements): 18248.18 requests per second
LRANGE_600 (first 600 elements): 14283.67 requests per second
```

## TEST #2 - jt client performance test

### [bench.go](https://github.com/hesidoryn/jt/blob/master/_examples/bench.go):
```shell
PING_INLINE: 53101.11 requests per second
SET: 45783.09 requests per second
GET: 41021.19 requests per second
RPUSH: 41570.65 requests per second
LPUSH: 41219.90 requests per second
LPOP: 40075.72 requests per second
RPOP: 40494.63 requests per second
```