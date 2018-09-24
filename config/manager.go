package config

import "sync"

type Request struct {
	Pattern string
	Method  string
	Headers map[string]string
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

type Stub struct {
	Request  Request
	Response Response
}

type Manager struct {
	mu    sync.RWMutex
	Stubs *[]Stub
}
