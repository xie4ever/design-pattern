package main

import (
	"time"

	"design-pattern/observer/badexample/example1/article"
)

func main() {
	a := article.NewArticle("title", "content")
	a.Add()
	time.Sleep(time.Second)
	a.Modify()
	time.Sleep(time.Second)
	a.Delete()
}
