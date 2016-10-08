Neighbor - 附近的人
===================

## 基于geohash以及mysql完成.

#### 依赖
 * mysql
 * golang >= 1.5

#### 启动

`go run server/app.go`

#### 示例接口
1. http://localhost:8080/neighborhood?lat=39.9046363143&lon=116.4071136987

```json
[
  {
    "user_id": 2,
    "username": "Jacky",
    "sex": "Male",
    "distance": 0.000008550965596164015,
    "distance_human": "10米以内",
    "latitude": 39.9046363143,
    "Longitude": 116.4071136988
  },
  {
    "user_id": 3,
    "username": "老顾大业",
    "sex": "Male",
    "distance": 1.110360933658212,
    "distance_human": "10米以内",
    "latitude": 39.9046463143,
    "Longitude": 116.4071135987
  },
  {
    "user_id": 6,
    "username": "老夏",
    "sex": "Male",
    "distance": 677.3004404492973,
    "distance_human": "700米以内",
    "latitude": 39.9107363143,
    "Longitude": 116.4071136987
  },
  {
    "user_id": 5,
    "username": "Tomy",
    "sex": "Male",
    "distance": 890.1167125329561,
    "distance_human": "900米以内",
    "latitude": 39.9097563143,
    "Longitude": 116.4151236988
  }
]
```

2. http://localhost:8080/neighbors

```json
[
  {
    "user_id": 7,
    "username": "刘力扬",
    "sex": "Female",
    "distance": 208.96470015523977,
    "distance_human": "300米以内",
    "latitude": 31.2519930642,
    "Longitude": 121.3578674217
  },
  {
    "user_id": 22,
    "username": "阿德",
    "sex": "Female",
    "distance": 210.12138203245195,
    "distance_human": "300米以内",
    "latitude": 31.2489378021,
    "Longitude": 121.3584549483
  },
  {
    "user_id": 19,
    "username": "高峰",
    "sex": "Female",
    "distance": 214.3164386771013,
    "distance_human": "300米以内",
    "latitude": 31.2520240642,
    "Longitude": 121.3579214217
  },
  {
    "user_id": 8,
    "username": "阿里",
    "sex": "Male",
    "distance": 341.74883506452574,
    "distance_human": "400米以内",
    "latitude": 31.2511440642,
    "Longitude": 121.3534664217
  },
  {
    "user_id": 21,
    "username": "德芙",
    "sex": "Male",
    "distance": 393.35329493189477,
    "distance_human": "400米以内",
    "latitude": 31.2483048021,
    "Longitude": 121.3534969483
  },
  {
    "user_id": 10,
    "username": "胡来的个",
    "sex": "Female",
    "distance": 406.3910888234207,
    "distance_human": "500米以内",
    "latitude": 31.2538910642,
    "Longitude": 121.3560894217
  },
  {
    "user_id": 24,
    "username": "玉溪",
    "sex": "Male",
    "distance": 438.6000300394465,
    "distance_human": "500米以内",
    "latitude": 31.2465758021,
    "Longitude": 121.3553469483
  },
  {
    "user_id": 20,
    "username": "哈里森",
    "sex": "Male",
    "distance": 533.8391984486166,
    "distance_human": "600米以内",
    "latitude": 31.2458658021,
    "Longitude": 121.3547179483
  },
  {
    "user_id": 16,
    "username": "乙二胺",
    "sex": "Female",
    "distance": 595.4151051533984,
    "distance_human": "600米以内",
    "latitude": 31.2480578021,
    "Longitude": 121.3512329483
  },
  {
    "user_id": 18,
    "username": "穷矮搓",
    "sex": "Male",
    "distance": 601.6754375924803,
    "distance_human": "700米以内",
    "latitude": 31.2452798021,
    "Longitude": 121.3545029483
  },
  {
    "user_id": 11,
    "username": "大丁哥",
    "sex": "Male",
    "distance": 645.5548520863088,
    "distance_human": "700米以内",
    "latitude": 31.2444768021,
    "Longitude": 121.3566589483
  },
  {
    "user_id": 17,
    "username": "二叉",
    "sex": "Male",
    "distance": 723.3458676492079,
    "distance_human": "800米以内",
    "latitude": 31.2438898021,
    "Longitude": 121.3554729483
  },
  {
    "user_id": 23,
    "username": "欧赔",
    "sex": "Male",
    "distance": 732.9363600998327,
    "distance_human": "800米以内",
    "latitude": 31.2441838021,
    "Longitude": 121.3539819483
  },
  {
    "user_id": 12,
    "username": "饭店阿出",
    "sex": "Male",
    "distance": 864.0096576454073,
    "distance_human": "900米以内",
    "latitude": 31.2471628021,
    "Longitude": 121.3486099483
  },
  {
    "user_id": 15,
    "username": "我哦哦",
    "sex": "Male",
    "distance": 907.762154435684,
    "distance_human": "1000米以内",
    "latitude": 31.2421608021,
    "Longitude": 121.3558319483
  },
  {
    "user_id": 13,
    "username": "呼呼",
    "sex": "Female",
    "distance": 956.6586330850344,
    "distance_human": "1000米以内",
    "latitude": 31.247013878,
    "Longitude": 121.3662039778
  },
  {
    "user_id": 14,
    "username": "华美",
    "sex": "Male",
    "distance": 995.6481638011498,
    "distance_human": "1000米以内",
    "latitude": 31.2431488021,
    "Longitude": 121.3505859483
  }
]
```

#### 性能

```
$ wrk -d10 -t100 -c400 http://localhost:8080/neighborhood\?lat\=39.9046363143\&lon\=116.4071136987
Running 10s test @ http://localhost:8080/neighborhood?lat=39.9046363143&lon=116.4071136987
  100 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   344.57ms   49.32ms 625.15ms   75.07%
    Req/Sec    12.29      6.15    30.00     59.98%
  11486 requests in 10.10s, 8.59MB read
  Socket errors: connect 0, read 1, write 0, timeout 0
Requests/sec:   1137.41
Transfer/sec:      0.85MB
```

```
wrk -d10 -t100 -c400 http://localhost:8080/neighbors
Running 10s test @ http://localhost:8080/neighbors
  100 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   628.25ms   58.07ms 887.54ms   74.10%
    Req/Sec     8.03      5.74    30.00     77.94%
  6161 requests in 10.11s, 17.17MB read
  Socket errors: connect 0, read 77, write 0, timeout 0
Requests/sec:    609.66
Transfer/sec:      1.70MB
```

#### Geohash 精确度误差

|geohash length	|lat bits	|lng bits	|lat error	|lng error	|km error |
|---------------|-----------|-----------|-----------|-----------|---------|
|1	            |2	        |3	        |± 23	    |± 23	    |± 2500   |
|2	            |5	        |5	        |± 2.8	    |± 5.6	    |±630     |
|3	            |7		    |8	        |± 0.70	    |± 0.7	    |±78      |
|4	            |10	       	|10	        |± 0.087	|± 0.18	    |±20      |
|5	            |12	       	|13	        |± 0.022	|± 0.022	|±2.4     |
|6	            |15	       	|15     	|± 0.0027	|± 0.0055	|±0.61    |
|7	            |17	       	|18	        |±0.00068	|±0.00068	|±0.076   |
|8	            |20	       	|20	        |±0.000085	|±0.00017	|±0.019   |

