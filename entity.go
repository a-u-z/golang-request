package main

type Req struct {
	baseUrl string
	header  []keyValue
	data    []keyValue
}

type keyValue struct {
	key   string
	value interface{}
}
