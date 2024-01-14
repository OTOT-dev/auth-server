package common

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetExecPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}
