package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type PostsCloud []Post

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var posts PostsCloud

	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(posts)
}
