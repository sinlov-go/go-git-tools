package example_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func BenchmarkExampleBasic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
		if err != nil {
			b.Fatal(err)
		}
		assert.Truef(b, pathExistsFast(testDataFolderFullPath), "want testDataFolderFullPath exist: %s", testDataFolderFullPath)
	}
	b.StopTimer()
}

func BenchmarkExampleParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, 10, len(randomStr(10)))
		}
	})
	b.StopTimer()
}

func demoCunt() bool {
	return randomInt(10) > 5
}

func BenchmarkExampleTimer(b *testing.B) {
	// mock ExampleTimer

	// reset counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// do ExampleTimer
		flag := demoCunt()
		if flag {
			// need for timing
			b.StartTimer()
		} else {
			// no need for timing
			b.StopTimer()
		}
		// verify ExampleTimer
		assert.Truef(b, true, "please fix this")
	}
	b.StopTimer()
}

func BenchmarkStringsAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for _, v := range strData {
			s += v
		}
	}
	b.StopTimer()
}

func BenchmarkStringsFmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var _ string = fmt.Sprint(strData)
	}
	b.StopTimer()
}

func BenchmarkStringsJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(strData, "")
	}
	b.StopTimer()
}

func BenchmarkStringsBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len("") * (len(strData) - 1)
		for i := 0; i < len(strData); i++ {
			n += len(strData[i])
		}
		var s bytes.Buffer
		s.Grow(n)
		for _, v := range strData {
			s.WriteString(v)
		}
		_ = s.String()
	}
	b.StopTimer()
}

func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len("") * (len(strData) - 1)
		for i := 0; i < len(strData); i++ {
			n += len(strData[i])
		}
		var b strings.Builder
		b.Grow(n)
		b.WriteString(strData[0])
		for _, s := range strData[1:] {
			b.WriteString("")
			b.WriteString(s)
		}
		_ = b.String()
	}
	b.StopTimer()
}
