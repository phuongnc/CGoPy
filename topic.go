package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type Dictionary struct {
	dict map[string]interface{}
	sync.RWMutex
}

var dictionary *Dictionary

//export LoadDictionary
func LoadDictionary(cFilePath *C.char) *C.char {
	fmt.Println("Start loading data.")

	goFilePath := C.GoString(cFilePath)

	// Read json file
	file, err := ioutil.ReadFile(goFilePath)
	if err != nil {
		return C.CString(err.Error())
	}

	dictionary = &Dictionary{
		dict: make(map[string]interface{}),
	}

	// Unmarshal to dictionary struct
	err = json.Unmarshal([]byte(file), &dictionary.dict)
	if err != nil {
		fmt.Println(err.Error())
		return C.CString(err.Error())
	}

	fmt.Println("End loading data. Total Rows:", len(dictionary.dict))
	return C.CString("")
}

//export UpdateDictionaryTest
func UpdateDictionaryTest() int {
	// Change value between two random word
	rand.Seed(time.Now().UnixNano())
	word1 := strconv.Itoa(rand.Intn(39))
	word2 := strconv.Itoa(rand.Intn(39))

	// Lock and update value
	dictionary.Lock()
	dictionary.dict[word1] = dictionary.dict[word2]
	dictionary.Unlock()

	return 1
}

//export UpdateDictionary
func UpdateDictionary() {
	// Change value between two random word
	rand.Seed(time.Now().UnixNano())
	word1 := strconv.Itoa(rand.Intn(39))
	word2 := strconv.Itoa(rand.Intn(39))

	// Lock and update value
	dictionary.Lock()
	dictionary.dict[word1] = dictionary.dict[word2]
	dictionary.Unlock()
}

//export GetTopic
func GetTopic(cWord *C.char) *C.char {
	goWord := C.GoString(cWord)
	dictionary.RLock()
	topicsValues, isExit := dictionary.dict[goWord]
	dictionary.RUnlock()
	if !isExit {
		return C.CString("{}")
	}
	response, _ := json.Marshal(topicsValues)
	return C.CString(string(response))
}

//export Free
func Free(word *C.char) {
	C.free(unsafe.Pointer(word))
}

func main() {
}
