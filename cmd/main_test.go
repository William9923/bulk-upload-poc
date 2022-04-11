package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pingpong(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(pingpong(), "Ping - Pong")
}
