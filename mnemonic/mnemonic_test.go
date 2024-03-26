package mnemonic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOne(t *testing.T) {
	res1 := GenerateOne(Num12)
	assert.True(t, len(strings.Split(res1, " ")) == 12)

	res2 := GenerateOne(Num24)
	assert.True(t, len(strings.Split(res2, " ")) == 24)
}

func TestGenerate1(t *testing.T) {
	count := 5

	res1 := Generate(Num12, count)
	assert.True(t, len(res1) == count)

	for i := 0; i < count; i++ {
		mnemonic := res1[i]
		assert.True(t, len(strings.Split(mnemonic, " ")) == 12)
	}

	res2 := Generate(Num24, count)
	assert.True(t, len(res2) == count)

	for i := 0; i < count; i++ {
		mnemonic := res2[i]
		assert.True(t, len(strings.Split(mnemonic, " ")) == 24)
	}

	assert.Len(t, Generate(Num12, 1), 1)
	assert.Len(t, Generate(Num12, 0), 0)
	assert.Len(t, Generate(Num24, 0), 0)
	assert.Len(t, Generate(Num12, -1), 0)
	assert.Len(t, Generate(Num24, -1), 0)
}
