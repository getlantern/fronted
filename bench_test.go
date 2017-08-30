package fronted

import (
	"testing"
	"time"
)

func BenchmarkCopyPointer(b *testing.B) {
	src := make([]*masquerade, 1000)
	dst := make([]*masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src = append(src, &masquerade{Masquerade{"Doamin", "IP"}, time.Now()})
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		copy(dst, src)
	}
}

func BenchmarkCopyValue(b *testing.B) {
	src := make([]masquerade, 1000)
	dst := make([]masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src = append(src, masquerade{Masquerade{"Doamin", "IP"}, time.Now()})
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		copy(dst, src)
	}
}

func BenchmarkAssignPointer(b *testing.B) {
	src := make([]*masquerade, 1000)
	dst := make([]*masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src = append(src, &masquerade{Masquerade{"Doamin", "IP"}, time.Now()})
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			dst[i] = src[i]
		}
	}
}

func BenchmarkAssignValue(b *testing.B) {
	src := make([]masquerade, 1000)
	dst := make([]masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src = append(src, masquerade{Masquerade{"Doamin", "IP"}, time.Now()})
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			dst[i] = src[i]
		}
	}
}

func BenchmarkChannelPointer(b *testing.B) {
	src := make(chan *masquerade, 1000)
	dst := make(chan *masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src <- &masquerade{Masquerade{"Doamin", "IP"}, time.Now()}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			dst <- <-src
			src <- <-dst
		}
	}
}

func BenchmarkChannelValue(b *testing.B) {
	src := make(chan masquerade, 1000)
	dst := make(chan masquerade, 1000)
	for i := 0; i < 1000; i++ {
		src <- masquerade{Masquerade{"Doamin", "IP"}, time.Now()}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			dst <- <-src
			src <- <-dst
		}
	}
}
