package test

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// BenchmarkSpliceAddString10 測試使用+= 拼接N次長度為10的字串
var bN = 10

func BenchmarkSpliceAddString10(b *testing.B) {
	s := ""
	for i := 0; i < bN; i++ {
		s += GenRandString(10)
	}
}

// BenchmarkSpliceBuilderString10 測試使用strings.Builder拼接N次長度為10的字串
func BenchmarkSpliceBuilderString10(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < bN; i++ {
		builder.WriteString(GenRandString(10))
	}
}

// BenchmarkSpliceBufferString10 測試使用bytes.Buffer拼接N次長度為10的字串
func BenchmarkSpliceBufferString10(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < bN; i++ {
		buff.WriteString(GenRandString(10))
	}
}

// BenchmarkSpliceBufferByte10 測試使用bytes.Buffer拼接N次長度為10的[]byte
func BenchmarkSpliceBufferByte10(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < bN; i++ {
		buff.Write(GenRandBytes(10))
	}
}

// BenchmarkSpliceBuilderByte10 測試使用string.Builder拼接N次長度為10的[]byte
func BenchmarkSpliceBuilderByte10(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < bN; i++ {
		builder.Write(GenRandBytes(10))
	}
}

const (
	data = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890,.-=/ "
)

func init() {
	rand.Seed(time.Now().Unix()) //設定隨機種子
}

// GenRandString 產生n個隨機字元的string
func GenRandString(n int) string {
	max := len(data)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(data[rand.Intn(max)])
	}

	return buf.String()
}

// GenRandBytes 產生n個隨機字元的[]byte
func GenRandBytes(n int) []byte {
	max := len(data)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = data[rand.Intn(max)]
	}

	return buf
}
