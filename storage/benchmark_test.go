package storage

import (
	"strconv"
	"testing"

	"github.com/hesidoryn/jt/config"
)

func BenchmarkSet(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), "foo")
	}
}

func BenchmarkGet(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.Get(strconv.Itoa(i))
	}
}

func BenchmarkIncr(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.IncrBy(strconv.Itoa(i), 1)
	}
}

func BenchmarkLPush(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.LPush(strconv.Itoa(i), strconv.Itoa(i))
	}
}

func BenchmarkRPush(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.RPush(strconv.Itoa(i), strconv.Itoa(i))
	}
}

func BenchmarkLPop(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.LPop(strconv.Itoa(i))
	}
}

func BenchmarkRPop(b *testing.B) {
	s := Init(config.Config{})
	for i := 0; i < b.N; i++ {
		s.RPop(strconv.Itoa(i))
	}
}
