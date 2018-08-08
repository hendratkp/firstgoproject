package main

import "fmt"
import "github.com/google/go-querystring/query"


type Product struct {
    name   string `url:"q"`
    active bool   `url:"all"`
    qty    int    `url:"page"`
}

func main() {
	fmt.Println("my first project in Golang")
    opt := Product{ "Sepatu", true, 2 }
    v, _ := query.Values(opt)
    fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"
}