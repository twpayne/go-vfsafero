package vfsafero_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/twpayne/go-vfs/v4/vfst"

	vfsafero "github.com/twpayne/go-vfsafero/v4"
)

func ExampleNewAferoFS() {
	Test := func(t *testing.T) {
		fs, cleanup, err := vfst.NewTestFS(map[string]interface{}{
			"/home/user/.bashrc": "# contents of .bashrc\n",
		})
		if err != nil {
			t.Fatal(err)
		}
		defer cleanup()

		aferoFS := vfsafero.NewAferoFS(fs)
		_ = afero.WriteFile(aferoFS, "/home/user/foo", []byte("bar"), 0666)

		vfst.RunTests(t, fs, "",
			vfst.TestPath("/home/user/foo",
				vfst.TestContentsString("bar"),
			),
		)
	}

	Test(&testing.T{})
}
