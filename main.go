package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")

}

//var res int
//var maxChan = 10
//var chanCount = 0
//var searchRequest = make(chan string)
//var chanDone = make(chan bool)
//var isFind = make(chan bool)
//var lock sync.Mutex
//
//func main() {
//	start := time.Now()
//	chanCount++
//	go search("C:\\Users\\mibbp\\", true)
//	waitForChan()
//	fmt.Println(res, "res")
//	fmt.Println(time.Since(start))
//}
//
//func waitForChan() {
//	for {
//		select {
//		case path := <-searchRequest:
//			lock.Lock()
//			chanCount++
//			lock.Unlock()
//			go search(path, true)
//		case <-chanDone:
//			lock.Lock()
//			chanCount--
//			lock.Unlock()
//			if chanCount == 0 {
//				return
//			}
//		case <-isFind:
//			res++
//		}
//	}
//}
//
//func search(path string, isChan bool) {
//	files, err := os.ReadDir(path)
//	if err == nil {
//		for _, file := range files {
//			name := file.Name()
//			if name == query {
//				res++
//				isFind <- true
//			}
//			if file.IsDir() {
//				if chanCount < maxChan {
//					searchRequest <- path + name + "\\"
//				}
//				search(path+name+"\\", false)
//			}
//		}
//	} else {
//		errors.New("read not success")
//	}
//	if isChan {
//		chanDone <- true
//	}
//
//}
