package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
type testClientImpl struct {
}

func (c *testClientImpl) Get() int {
	return 2
}

func TestMainRequest(t *testing.T) {
	// testClientImpl の実体をセットする
	sample := &Sample{&testClientImpl{}}
	res := sample.doGet()
	assert.EqualValues(t, res, 2)
}
