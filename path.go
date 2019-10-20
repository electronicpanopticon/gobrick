package gobrick

import (
	"fmt"
	"go/build"
	"os"
)

func GetGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

// TODO: Support for multiple GOPATHs
func FilePathInGoPath(relativePath string) string {
	return fmt.Sprintf("%s/%s", GetGOPATH(), relativePath)
}

func FileInGoPath(relativePath string) (*os.File, error) {
	fp := FilePathInGoPath(relativePath)
	return os.Open(fp)
}