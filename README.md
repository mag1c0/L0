# WB Tech - L0 

Backend: Go + PostgreSQL + Nats + Rest api on `http://localhost:30016/api/v1/orders`

Frontend: NuxtJS on `http://localhost:3000`

## Run
```bash
# docker
docker-compose up
```

## Vegeta Testing
```bash
# vegeta attack -duration=5s -rate=100/s --targets=req.txt
Requests      [total, rate, throughput]         500, 100.21, 100.16
Duration      [total, attack, wait]             4.992s, 4.99s, 2.279ms
Latencies     [min, mean, 50, 90, 95, 99, max]  794.625µs, 1.858ms, 1.778ms, 2.327ms, 2.596ms, 3.239ms, 6.329ms
Bytes In      [total, mean]                     480000, 960.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:500
```

```bash
# vegeta attack -duration=5s -rate=1000/s --targets=req.txt
Requests      [total, rate, throughput]         5000, 1000.21, 1000.13
Duration      [total, attack, wait]             4.999s, 4.999s, 413.458µs
Latencies     [min, mean, 50, 90, 95, 99, max]  227.917µs, 483.455µs, 405.095µs, 514.738µs, 617.8µs, 2.987ms, 13.116ms
Bytes In      [total, mean]                     4800000, 960.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:5000
```

```bash
# vegeta attack -duration=5s -rate=10000/s --targets=req.txt
Requests      [total, rate, throughput]         50000, 10000.13, 9988.46
Duration      [total, attack, wait]             5.006s, 5s, 5.843ms
Latencies     [min, mean, 50, 90, 95, 99, max]  119.834µs, 15.172ms, 5.362ms, 41.34ms, 52.071ms, 93.865ms, 137.761ms
Bytes In      [total, mean]                     48000000, 960.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:50000
```