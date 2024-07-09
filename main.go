package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
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
