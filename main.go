package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"
	"unsafe"

	"github.com/valyala/fasthttp"
)

// #include "filter.h"
// #include "libtopic.h"
import "C"

type TopicRequest struct {
	Text string `json:"text"`
}

type TopicResponse struct {
	Text   string `json:"text"`
	Topics string `json:"topics"`
}

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

var (
	contentType     = []byte("Content-Type")
	applicationJSON = []byte("application/json")
)

func startServer() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/topics":
		getTopicsHandler(ctx)
		return
	default:
		responseError(ctx, fasthttp.StatusNotFound, "Your request does not supported")
		return
	}
}

func responseError(ctx *fasthttp.RequestCtx, code int, message string) {
	ctx.Error(message, code)
}

func responseSuccess(ctx *fasthttp.RequestCtx, code int, data []byte) {
	ctx.Response.Header.SetCanonical(contentType, applicationJSON)
	ctx.Response.SetStatusCode(code)
	ctx.SetBody(data)
}

func getTopicsHandler(ctx *fasthttp.RequestCtx) {
	var payload TopicRequest
	if err := json.Unmarshal(ctx.PostBody(), &payload); err != nil {
		responseError(ctx, fasthttp.StatusBadRequest, "cannot unmarshal payload")
		return
	}

	println(payload.Text)

	// Call func from python
	result := callPythonFilterText(payload.Text)
	println(result)

	// Response
	response := TopicResponse{
		Text:   payload.Text,
		Topics: result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		responseError(ctx, fasthttp.StatusInternalServerError, "marshal response failed")
		return

	}

	responseSuccess(ctx, fasthttp.StatusOK, jsonResponse)
}

// Call python via C library
func callPythonFilterText(text string) string {
	p := C.CString(text)
	ret := C.Py_Filter_Text(p)
	result := C.GoString(ret)
	defer C.free(unsafe.Pointer(p))
	return result
}

// Every 100ms change one of the topics in the dictionary
func updateDictionary() {
	C.UpdateDictionary()
	time.Sleep(100 * time.Millisecond)
	updateDictionary()
}

// Load dictionary from json file
func loadDictionary(filePath string) {
	jsonPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(jsonPath))
	C.LoadDictionary(jsonPath)
}

/* ---- MAIN ---- */
func main() {

	// Load Python Env
	C.Py_Load()

	// Call load dictionary when start server
	loadDictionary("data/dictionary.json")

	// Change dictionary 100ms
	go updateDictionary()

	// Start server
	println("Start server on 8080")
	startServer()

	// Unload Python Env
	C.Py_Unload()
}
