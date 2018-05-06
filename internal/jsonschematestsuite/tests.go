package jsonschematestsuite

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sourcegraph/go-jsonschema/jsonschema"
)

// File is a file of test groups from the JSON Schema test suite.
type File struct {
	Name string `json:"-"`
	path string

	Groups []*Group // call Read to populate this field
}

// Read reads and unmarshals the test file from disk.
func (f *File) Read() error {
	f2, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer f2.Close()
	if err := json.NewDecoder(f2).Decode(&f.Groups); err != nil {
		return err
	}
	for _, g := range f.Groups {
		if err := json.Unmarshal(g.RawSchema, &g.Schema); err != nil {
			return err
		}
	}
	return nil
}

// ReadT wraps the Read method for easier use in tests. It calls t.Fatal(err) if Read returns an
// error.
func (f *File) ReadT(t testing.TB) {
	t.Helper()
	if err := f.Read(); err != nil {
		t.Fatal(err)
	}
}

// Group is a group of test cases from the JSON Schema test suite.
type Group struct {
	Description string
	RawSchema   json.RawMessage    `json:"schema"`
	Schema      *jsonschema.Schema `json:"-"`
	Tests       []TestCase
}

// TestCase is a test case from the JSON Schema test suite.
type TestCase struct {
	Description string
	Data        json.RawMessage
	Valid       bool
}

// Files returns all test files from the JSON Schema official test suite and this library's own test
// suite.
func Files(internalDir string) (files []File, err error) {
	officialTestSuiteDir := filepath.Join(internalDir, "jsonschematestsuite", "testdata", "official")
	for _, root := range []string{
		filepath.Join(internalDir, "jsonschematestsuite", "testdata"),
		filepath.Join(officialTestSuiteDir, "tests", "draft7"),
	} {
		err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if path == officialTestSuiteDir {
				return filepath.SkipDir
			}
			if !info.Mode().IsRegular() || filepath.Ext(info.Name()) != ".json" {
				return nil
			}
			files = append(files, File{
				Name: strings.TrimSuffix(strings.TrimPrefix(path, root+string(os.PathSeparator)), ".json"),
				path: path,
			})
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return files, err
}
