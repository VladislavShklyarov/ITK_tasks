package main

type Vehicle interface {
	StartEngine() error
	StopEngine() error
	GetInfo() string
}
