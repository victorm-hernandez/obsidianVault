# Configuration tests

Command used to run test

```bash
# Terminal 1, Run web server
go run cmd/api/main.go
# Terminal 2, Run load test
hey -n 5000 -c 50 http://localhost:8888/
```

## Configuration 1

```go
config := &ServiceConfig{
		Address:          ":8888",
		LoggingLevel:     slog.LevelDebug,
		EnableProfiling:  true,
		ProfilingAddress: ":8889",

		// Timeouts
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 2,
		IddleTimeout: time.Second * 30,

		// Weather generator
		BufferSize:     1000,
		ClientCount:    20,
		MaxRetryCount:  3,
		MinRetryDelay:  time.Millisecond * 100,
		MaxRetryDelay:  time.Second * 6,
		MaxRecordAge:   time.Minute * 10,
		MaxFetchPerSec: 20, // This seems to be barely below the throttling threshold for the Weather APIs

		// Fetcher
		RandomLocAPIURL: "https://locations.patch3s.dev/api/random",
		PointsAPIURL:    "https://api.weather.gov/points/%v,%v",
	}
```

Result

```text
Summary:
  Total:        570.8299 secs
  Slowest:      9.5048 secs
  Fastest:      0.0001 secs
  Average:      5.6571 secs
  Requests/sec: 8.7592
  
  Total data:   873224 bytes
  Size/request: 174 bytes

Response time histogram:
  0.000 [1]     |
  0.951 [359]   |■■■■■■■■■
  1.901 [7]     |
  2.852 [12]    |
  3.802 [12]    |
  4.752 [221]   |■■■■■
  5.703 [1542]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  6.653 [1610]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  7.604 [860]   |■■■■■■■■■■■■■■■■■■■■■
  8.554 [299]   |■■■■■■■
  9.505 [77]    |■■


Latency distribution:
  10%% in 4.5555 secs
  25%% in 5.2613 secs
  50%% in 5.8603 secs
  75%% in 6.6403 secs
  90%% in 7.4225 secs
  95%% in 7.8621 secs
  99%% in 8.7908 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0071 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0026 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:    5.6568 secs, 0.0001 secs, 9.5046 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200] 5000 responses
```

## Configuration 2

```go
config := &ServiceConfig{
		Address:          ":8888",
		LoggingLevel:     slog.LevelDebug,
		EnableProfiling:  true,
		ProfilingAddress: ":8889",

		// Timeouts
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 2,
		IddleTimeout: time.Second * 30,

		// Weather generator
		BufferSize:     2000,
		ClientCount:    50,
		MaxRetryCount:  2,
		MinRetryDelay:  time.Second * 3,
		MaxRetryDelay:  time.Second * 6,
		MaxRecordAge:   time.Minute * 10,
		MaxFetchPerSec: 20, // This seems to be barely below the throttling threshold for the Weather APIs

		// Fetcher
		RandomLocAPIURL: "https://locations.patch3s.dev/api/random",
		PointsAPIURL:    "https://api.weather.gov/points/%v,%v",
	}
```

Results

```text
Summary:
  Total:        554.0853 secs
  Slowest:      7.9128 secs
  Fastest:      0.0003 secs
  Average:      5.4957 secs
  Requests/sec: 9.0239
  
  Total data:   854241 bytes
  Size/request: 170 bytes

Response time histogram:
  0.000 [1]     |
  0.792 [150]   |■■■
  1.583 [8]     |
  2.374 [19]    |
  3.165 [9]     |
  3.957 [22]    |
  4.748 [337]   |■■■■■■■
  5.539 [1680]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  6.330 [1934]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  7.122 [705]   |■■■■■■■■■■■■■■■
  7.913 [135]   |■■■


Latency distribution:
  10%% in 4.6983 secs
  25%% in 5.1484 secs
  50%% in 5.6523 secs
  75%% in 6.1158 secs
  90%% in 6.5688 secs
  95%% in 6.8811 secs
  99%% in 7.4546 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0034 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:    5.4955 secs, 0.0001 secs, 7.9126 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200] 5000 responses
```

## Configuration 3

```go
config := &ServiceConfig{
		Address:          ":8888",
		LoggingLevel:     slog.LevelDebug,
		EnableProfiling:  true,
		ProfilingAddress: ":8889",

		// Timeouts
		WriteTimeout: time.Second * 8,
		ReadTimeout:  time.Second * 2,
		IddleTimeout: time.Second * 30,

		// Weather generator
		BufferSize:     2000,
		ClientCount:    50,
		MaxRetryCount:  1,
		MinRetryDelay:  time.Second * 6,
		MaxRetryDelay:  time.Second * 6,
		MaxRecordAge:   time.Minute * 10,
		MaxFetchPerSec: 20, // This seems to be barely below the throttling threshold for the Weather APIs

		// Fetcher
		RandomLocAPIURL: "https://locations.patch3s.dev/api/random",
		PointsAPIURL:    "https://api.weather.gov/points/%v,%v",
	}

```

```text
Summary:
  Total:        593.3822 secs
  Slowest:      8.0022 secs
  Fastest:      0.0002 secs
  Average:      5.8915 secs
  Requests/sec: 8.4263
  
  Total data:   834360 bytes
  Size/request: 166 bytes

Response time histogram:
  0.000 [1]     |
  0.800 [115]   |■■
  1.601 [6]     |
  2.401 [8]     |
  3.201 [6]     |
  4.001 [11]    |
  4.801 [166]   |■■■
  5.602 [1154]  |■■■■■■■■■■■■■■■■■■■■■■
  6.402 [2109]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  7.202 [1056]  |■■■■■■■■■■■■■■■■■■■■
  8.002 [368]   |■■■■■■■


Latency distribution:
  10%% in 5.0214 secs
  25%% in 5.4962 secs
  50%% in 5.9754 secs
  75%% in 6.4922 secs
  90%% in 7.0378 secs
  95%% in 7.3977 secs
  99%% in 7.8748 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0058 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0021 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:    5.8913 secs, 0.0001 secs, 8.0017 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0055 secs

Status code distribution:
  [200] 5000 responses
```

## Configuration Test 4

```json
{
    "generatorStatus": {
        "bootDurationMs": 0,
        "itemCount": 59,
        "ready": false
    },
    "outgoingReqStatistics": {
        "api.weather.gov": {
            "domainName": "",
            "statsPerHttpCode": {
                "200": {
                    "requestCount": 2218,
                    "durationCeiling": 2620853125,
                    "durationFloor": 9793958,
                    "durationAverage": 221794742
                },
                "301": {
                    "requestCount": 1097,
                    "durationCeiling": 571166333,
                    "durationFloor": 9610917,
                    "durationAverage": 127903469
                }
            }
        },
        "locations.patch3s.dev": {
            "domainName": "",
            "statsPerHttpCode": {
                "200": {
                    "requestCount": 1109,
                    "durationCeiling": 7991946542,
                    "durationFloor":   73604500,
                    "durationAverage": 1228684684
                },
                "429": {
                    "requestCount": 7067,
                    "durationCeiling": 336654583,
                    "durationFloor": 73403708,
                    "durationAverage": 77572544
                },
                "500": {
                    "requestCount": 32,
                    "durationCeiling": 7311924500,
                    "durationFloor": 73852541,
                    "durationAverage": 1164099157
                },
                "502": {
                    "requestCount": 29,
                    "durationCeiling": 6491863000,
                    "durationFloor": 74079583,
                    "durationAverage": 1240638577
                },
                "503": {
                    "requestCount": 28,
                    "durationCeiling": 7589722083,
                    "durationFloor": 74213166,
                    "durationAverage": 1695601327
                },
                "504": {
                    "requestCount": 26,
                    "durationCeiling": 7141453042,
                    "durationFloor": 74243125,
                    "durationAverage": 1831866987
                }
            }
        }
    }
}
```


### Configuration 5 (refactor handler)

Before refactor

Latency distribution:
  10%% in 5.2269 secs
  25%% in 5.5585 secs
  50%% in 5.9682 secs
  75%% in 6.3543 secs
  90%% in 6.7127 secs
  95%% in 6.9156 secs
  99%% in 7.3161 secs


After refactor, using cache instead of buffer. 