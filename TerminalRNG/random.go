package main

type Random interface {
	Mode(Source int64) int64
	Print(Source int64)
}
