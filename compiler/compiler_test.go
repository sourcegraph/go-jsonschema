package compiler

import (
	"bytes"
	"encoding/json"
	"flag"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/sourcegraph/go-jsonschema/jsonschema"
)

var writeWant = flag.Bool("test.write-want", false, "(over)write want.go files in test cases with output")

func TestCompiler(t *testing.T) {
	entries, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.Mode().IsDir() {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			testCompiler(t, filepath.Join("testdata", entry.Name()))
		})
	}
}

func testCompiler(t *testing.T, dir string) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	var schemas []*jsonschema.Schema
	goFiles := map[string][]byte{}
	for _, entry := range entries {
		if entry.Mode().IsDir() {
			continue
		}
		data, err := ioutil.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			t.Fatalf("read %s: %s", entry.Name(), err)
		}
		switch filepath.Ext(entry.Name()) {
		case ".json":
			var schema jsonschema.Schema
			if err := json.Unmarshal(data, &schema); err != nil {
				t.Fatalf("unmarshal %s: %s", entry.Name(), err)
			}
			schemas = append(schemas, &schema)
		case ".go":
			goFiles[entry.Name()] = data
		}
	}

	decls, imports, err := Compile(schemas)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	file := &ast.File{Name: ast.NewIdent("p"), Imports: imports, Decls: decls}
	if err := format.Node(&buf, token.NewFileSet(), file); err != nil {
		t.Fatal(err)
	}
	out := buf.Bytes()
	if !bytes.HasSuffix(out, []byte("\n")) {
		out = append(out, '\n')
	}

	const goFile = "want.go"
	if *writeWant {
		if err := ioutil.WriteFile(filepath.Join(dir, goFile), out, 0600); err != nil {
			t.Fatal(err)
		}
	}
	if want := goFiles[goFile]; !bytes.Equal(out, want) {
		diff, err := diff(filepath.Join(dir, goFile), out)
		if err != nil {
			t.Fatal(err)
		}
		t.Errorf("got != want\n\n%s", diff)
	}
}

func diff(path string, data []byte) (string, error) {
	cmd := exec.Command("diff", "-N", "-u", path, "-")
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stderr = os.Stderr
	data, err := cmd.Output()
	if err != nil {
		waitStatus, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if ok && waitStatus.ExitStatus() == 1 /* files differ, no trouble */ && len(data) > 0 {
			err = nil
		}
	}
	return string(data), err
}
