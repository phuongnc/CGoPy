G2P_LDFLAGS = "$(python-config --ldflags --embed) -L$$(pwd) -lfilter -ltopic"
G2P_CFLAGS = "-g -Wall"

build:
		go get github.com/valyala/fasthttp
		go build -o libtopic.so -buildmode=c-shared topic.go
		gcc -c -fPIC  `python-config --cflags --libs` filter.c
		gcc -shared filter.o -o libfilter.so `python-config --cflags --ldflags --embed`
		CGO_LDFLAGS=${G2P_LDFLAGS} CGO_CFLAGS=${G2P_CFLAGS} go build main.go

rebuild:	clean	build

clean:
		rm -f *.dylib *.so *.a *.o main
		rm -f libtopic.h
		rm -rf __pycache__
run:	build
		./main