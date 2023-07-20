package main

import (
	"path"

	"golang.org/x/tools/godoc/vfs"
)

// GatherRoots will collect all directories that can contain go files for documentation.
// If in a go module, it will collect the roots of all dependent modules.
// If not, it will collect the contents of all the directories in GOPATH.
func GatherRoots(goRoot string) vfs.NameSpace {
	fs := vfs.NameSpace{}
	fs.Bind("/", vfs.OS(goRoot), "/", vfs.BindReplace)

	dirs := codeRoots()
	for _, dir := range dirs {
		fs.Bind(path.Join("/src", "pkg", dir.importPath), vfs.OS(dir.dir), "/src", vfs.BindAfter)
	}

	return fs
}
