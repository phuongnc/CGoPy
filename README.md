### Introduction.
My challenge is have to speed up python server by combine with Golang. Golang will process the request/response to client as Restful Api level. When has a request comming, Go have to call to Python for some logic processing, then get back the result and response to client.
To ensure the best performance, I use C language is bridge to call native between Go<=>Python. Not via third-path libraries, or GRPC message protocol,...

### Technical using

1.  Golang, Go-CGO, FastHttp framework
2.  Python and C language
3.  Linux and GCC
4.  Docker, Git.
5. In memory store (no database)

### Build & Run

Open your Terminal

3.  `docker build -t cgopy_image .`

4.  `docker run --name cgopy -p 8080:8080 -v $(pwd):/app --rm cgopy_image`

### CURL Test
From your Terminal
run `curl --header "Content-Type: application/json" --request POST --data '{"text":"test 1"}' http://localhost:8080/topics`

### Benchmark Test
From your Terminal

1.  `cd benchmark`

2.  `go test -bench=BenchmarkServer -benchmem -benchtime=30s`

### Test Results

#### System Specs

- OS: Ubuntu 18.04

- CPU: 2.5 GHz Intel Core i5

- RAM: 8 GB

- Docker version 17.09

#### Results

Run a test in 30 seconds (execute ~1000 requests)

- Memory Usage (no data.json): 14M

- Memory Usage (data.json): 513M ~ 600M

- QPS: 0.0161 second/request