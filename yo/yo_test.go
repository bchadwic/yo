package yo

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProdYo(t *testing.T) {
	yo := ProdYo()
	fmt.Fprintf(yo.Out, "Hello world")
	yo.In.Read([]byte("Hello world, the sequel"))

	assert.Equal(t, os.Stdin, yo.In)
	assert.Equal(t, os.Stdout, yo.Out)
}

func Test_TestYo(t *testing.T) {
	yo, _, _, _ := TestYo()
	fmt.Fprintf(yo.Out, "Hello world")
	yo.In.Read([]byte("Hello world, the sequel"))

	assert.NotEqual(t, os.Stdin, yo.In)
	assert.NotEqual(t, os.Stdout, yo.Out)
}
