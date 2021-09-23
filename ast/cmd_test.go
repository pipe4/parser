package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestParseFile(t *testing.T) {
	file, err := ParseFile("../examples/proto-build/proto-build.pipe4")
	require.Nil(t, err)

	assert.Less(t, 4, len(file.Import))

	log.Printf("%+v", file)
}
