package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	b := &bytes.Buffer{}
	cfg := config{Arg: "🌏", out: b}
	require.NoError(t, cfg.AfterApply())
	require.NoError(t, cfg.Run())
	require.Equal(t, "Hello 🌏\n", b.String())
}
