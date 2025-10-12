package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Greeting(prefix, name string) {
	fmt.Printf("Hello, %s %s!\n", prefix, name)
}

func Greeting2(out io.Writer, prefix, name string) {
	out.Write([]byte(fmt.Sprintf("Hello, %s %s!\n", prefix, name)))
}

func TestGreeting(t *testing.T) {
	originalStdout := os.Stdout

	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	Greeting("Mr.", "Joseph")
	w.Close()
	var buf strings.Builder
	_, err = io.Copy(&buf, r)
	assert.NoError(t, err)

	want := "Hello, Mr. Joseph!\n"
	got := buf.String()
	assert.Equal(t, want, got)
}

func TestGreeting2(t *testing.T) {

	var buf strings.Builder
	Greeting2(&buf, "Mr.", "Joseph")

	want := "Hello, Mr. Joseph!\n"
	got := buf.String()
	assert.Equal(t, want, got)
}
