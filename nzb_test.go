package main

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNZB(t *testing.T) {
	assert := assert.New(t)

	z := NewNZB("tests/test1.nzb")

	assert.Equal(len(z.Files), 3)
	assert.Equal(len(z.Files[0].Groups), 1)
	assert.Equal(len(z.Files[0].Segments), 2)
	assert.Equal(len(z.Files[1].Segments), 3)
}

func TestNZBSort(t *testing.T) {
	assert := assert.New(t)

	z := NewNZB("tests/test1.nzb")
	z.Sort()

	file0 := z.Files[0]
	assert.Equal(file0.Segments[0].Number, 1)
	assert.Equal(file0.Segments[0].ID, "file1-part1-seg1@test.example.com")
	assert.Equal(file0.Segments[1].Number, 2)
	assert.Equal(file0.Segments[1].ID, "file1-part1-seg2@test.example.com")

	file1 := z.Files[1]
	assert.Equal(file1.Segments[0].Number, 1)
	assert.Equal(file1.Segments[0].ID, "file2-part2-seg1@test.example.com")
	assert.Equal(file1.Segments[1].Number, 2)
	assert.Equal(file1.Segments[1].ID, "file2-part2-seg2@test.example.com")
	assert.Equal(file1.Segments[2].Number, 3)
	assert.Equal(file1.Segments[2].ID, "file2-part2-seg3@test.example.com")

}

func TestNZBAssignOffset(t *testing.T) {
	assert := assert.New(t)

	z := NewNZB("tests/test1.nzb")
	z.Sort()
	z.AssignOffset()

	file0 := z.Files[0]
	assert.Equal(file0.Segments[0].Offset, uint32(0))
	assert.Equal(file0.Segments[1].Offset, uint32(102400))

	file1 := z.Files[1]
	assert.Equal(file1.Segments[0].Offset, uint32(0))
	assert.Equal(file1.Segments[1].Offset, uint32(102400))
	assert.Equal(file1.Segments[2].Offset, uint32(102400+102400))
}
