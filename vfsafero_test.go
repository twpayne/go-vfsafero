package vfsafero

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/twpayne/go-vfs/v2/vfst"
)

var _ afero.Fs = &AferoFS{}

func TestAferoFS(t *testing.T) {
	fs, cleanup, err := vfst.NewTestFS(map[string]interface{}{
		"/home/user/.bashrc": "# contents of .bashrc\n",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	aferoFS := NewAferoFS(fs)
	if err := afero.WriteFile(aferoFS, "/home/user/foo", []byte("bar"), 0666); err != nil {
		t.Fatal(err)
	}

	vfst.RunTests(t, fs, "",
		vfst.TestPath("/home/user/foo",
			vfst.TestContentsString("bar"),
		),
	)
}
