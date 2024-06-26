package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

type Response struct {
	Id     string `json:"id"`
	Result string `json:"result"`
}

func main() {
	// Define the server address
	serverAddr := "http://localhost:8080/process"

	var wg sync.WaitGroup
	requests := []Request{
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "4", Data: "Data 4"},
		{Id: "1", Data: "Data 1"},
		{Id: "2", Data: "Data 2"},
		{Id: "3", Data: "Data 3"},
		{Id: "FINISH", Data: "FINISH"},
	}

	for _, req := range requests {
		wg.Add(1)
		go func(r Request) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			// Marshal the request body
			reqBody, err := json.Marshal(r)
			if err != nil {
				log.Fatalf("Failed to marshal request body: %v", err)
			}

			// Create a new HTTP POST request
			httpReq, err := http.NewRequestWithContext(ctx, "POST", serverAddr, bytes.NewBuffer(reqBody))
			if err != nil {
				log.Fatalf("Failed to create HTTP request: %v", err)
			}
			httpReq.Header.Set("Content-Type", "application/json")

			// Send the request
			resp, err := http.DefaultClient.Do(httpReq)
			if err != nil {
				log.Fatalf("Failed to send HTTP request: %v", err)
			}
			defer resp.Body.Close()

			// Read the response body
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Failed to read response body: %v", err)
			}

			// Unmarshal the response body
			var response Response
			if err := json.Unmarshal(respBody, &response); err != nil {
				log.Fatalf("Failed to unmarshal response body: %v", err)
			}

			// Log the response
			log.Printf("Response for request ID %s: %s", response.Id, response.Result)

		}(req)
	}

	wg.Wait()
}
