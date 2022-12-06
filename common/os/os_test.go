package os

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	notExists := "../go.mod"
	exists := "../../go.mod"

	assert.False(t, FileExists(notExists))
	assert.True(t, FileExists(exists))
}

func TestCopyFile(t *testing.T) {
	var mode os.FileMode = 0644
	src := "src.txt"
	dst := "dst.txt"
	content := []byte("基于变色龙哈希函数和共识投票的可修改区块链.pdf")
	srcFile, err := os.OpenFile(src, os.O_CREATE|os.O_TRUNC|os.O_APPEND, mode)
	assert.Nil(t, err)
	defer func() {
		_ = os.Remove(src)
	}()
	writeN, err := srcFile.Write(content)
	assert.Nil(t, err)
	assert.Nil(t, srcFile.Close())

	err = CopyFile(src, dst)
	assert.Nil(t, err)

	dstFile, err := os.Open(dst)
	assert.Nil(t, err)
	defer func() {
		_ = dstFile.Close()
		_ = os.Remove(dst)
	}()
	buffer := new(bytes.Buffer)
	readN, err := buffer.ReadFrom(dstFile)
	assert.Nil(t, err)
	assert.Equal(t, writeN, int(readN))
	assert.Equal(t, content, buffer.Bytes())
}

func TestEnsureDir(t *testing.T) {
	dir := filepath.Join("root", "home", "cosmic")
	err := EnsureDir(dir, 0644)
	assert.Nil(t, err)
	assert.DirExists(t, dir)
	assert.Nil(t, os.RemoveAll("root"))
}
