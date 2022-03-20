package main

import "fmt"

func main() {
	fmt.Println((9826296 - 1897752) / 100000)
	fmt.Println((5702008 - 1897752) / 50000)
	fmt.Println((3999864 - 1897752) / 30000)
	fmt.Println((2428792 - 1897752) / 10000)
}

/*
1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
10字节
SET: 90744.10 requests per second
GET: 90661.83 requests per second

20字节
SET: 91407.68 requests per second
GET: 90991.81 requests per second

50字节
SET: 92250.92 requests per second
GET: 91324.20 requests per second

100字节
SET: 90579.71 requests per second
GET: 92764.38 requests per second

200字节
SET: 90826.52 requests per second
GET: 91575.09 requests per second

1000字节
SET: 91324.20 requests per second
GET: 92336.11 requests per second

5000字节
SET: 88105.73 requests per second
GET: 91659.03 requests per second


2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
写入之前
used_memory:1897752
used_memory_human:1.81M
used_memory_rss:1860824
used_memory_rss_human:1.77M
used_memory_peak:5919848
used_memory_peak_human:5.65M
total_system_memory:0
total_system_memory_human:0B
used_memory_lua:37888
used_memory_lua_human:37.00K

写入数据脚本
for((i=1;i<=10000;i++));do
./redis-cli -c -h 127.0.0.1 -p 6379 set $i $i
done;

写入10000条数据后
used_memory:2428792
used_memory_human:2.32M
used_memory_rss:2391864
used_memory_rss_human:2.28M
used_memory_peak:5919848
used_memory_peak_human:5.65M
total_system_memory:0
total_system_memory_human:0B
used_memory_lua:37888
used_memory_lua_human:37.00K

写入30000条数据后
used_memory:3999864
used_memory_human:3.81M
used_memory_rss:3962936
used_memory_rss_human:3.78M
used_memory_peak:5919848
used_memory_peak_human:5.65M
total_system_memory:0
total_system_memory_human:0B
used_memory_lua:37888
used_memory_lua_human:37.00K

写入50000条数据后
used_memory:5702008
used_memory_human:5.44M
used_memory_rss:5665080
used_memory_rss_human:5.40M
used_memory_peak:5919848
used_memory_peak_human:5.65M
total_system_memory:0
total_system_memory_human:0B
used_memory_lua:37888
used_memory_lua_human:37.00K

写入100000条数据后
used_memory:9826296
used_memory_human:9.37M
used_memory_rss:9789368
used_memory_rss_human:9.34M
used_memory_peak:9930504
used_memory_peak_human:9.47M
total_system_memory:0
total_system_memory_human:0B
used_memory_lua:37888
used_memory_lua_human:37.00K

用used_memory计算可得
存入10000条数据时，平均每个key大小为53字节
存入30000条数据时，平均每个key大小为70字节
存入50000条数据时，平均每个key大小为76字节
存入100000条数据时，平均每个key大小为79字节
可能是随着key自身的增大，占用的内存也在变大
*/
