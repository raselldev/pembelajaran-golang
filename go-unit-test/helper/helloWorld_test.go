package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Raka")
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Raka")
	if result != "Hello Raka" {
		//t.Fail() // Unit test gagal tapi tetap melanjutkan test
		t.FailNow() // Unit test gagal dan langsung tidak dilanjutkan
		//test failed
		//panic("Result is not Hello Raka")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Raka")
	assert.Equal(t, "Hello Raka", result)
	fmt.Printf("Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Raka")
	require.Equal(t, "Hello Raka", result)
	fmt.Printf("Assert Done")
}
