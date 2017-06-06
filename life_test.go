package main

import "testing"

var (
	board10   = New(10, 10)
	board20   = New(20, 20)
	board40   = New(40, 40)
	board100  = New(100, 100)
	board500  = New(500, 500)
	board1000 = New(1000, 1000)
)

func benchmarkNext(next func(*GameBoard) *GameBoard, board *GameBoard, b *testing.B) {
	for n := 0; n < b.N; n++ {
		next(board)
	}
}

func BenchmarkNext10(b *testing.B) {
	benchmarkNext(Next, board10, b)
}

func BenchmarkNext20(b *testing.B) {
	benchmarkNext(Next, board20, b)
}

func BenchmarkNext40(b *testing.B) {
	benchmarkNext(Next, board40, b)
}

func BenchmarkNext100(b *testing.B) {
	benchmarkNext(Next, board100, b)
}

func BenchmarkNext500(b *testing.B) {
	benchmarkNext(Next, board500, b)
}

func BenchmarkNext1000(b *testing.B) {
	benchmarkNext(Next, board1000, b)
}

func BenchmarkNextOld10(b *testing.B) {
	benchmarkNext(NextOld, board10, b)
}

func BenchmarkNextOld20(b *testing.B) {
	benchmarkNext(NextOld, board20, b)
}

func BenchmarkNextOld40(b *testing.B) {
	benchmarkNext(NextOld, board40, b)
}

func BenchmarkNextOld100(b *testing.B) {
	benchmarkNext(NextOld, board100, b)
}

func BenchmarkNextOld500(b *testing.B) {
	benchmarkNext(NextOld, board500, b)
}

func BenchmarkNextOld1000(b *testing.B) {
	benchmarkNext(NextOld, board1000, b)
}
