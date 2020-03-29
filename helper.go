package logger

import "path/filepath"

func WindowsPath(path string) string {
	return "winfile:///" + filepath.ToSlash(path)
}
