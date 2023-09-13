package repo

import (
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// SafeString converts the input string into a safe naming style in golang
func SafeString(in string) string {
	if in == "" {
		return in
	}

	data := strings.Map(func(r rune) rune {
		if isSafeRune(r) {
			return r
		}
		return '_'
	}, in)

	headRune := rune(data[0])
	if isNumber(headRune) {
		return "_" + data
	}
	return data
}
func isSafeRune(r rune) bool {
	return isLetter(r) || isNumber(r) || r == '_'
}
func isLetter(r rune) bool {
	return 'A' <= r && r <= 'z'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if dir == "" {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

func FillModelPkgPath(filePath string) string {
	pkg, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  filePath,
	})
	if err != nil {
		return ""
	}
	if len(pkg) == 0 {
		return ""
	}
	return pkg[0].PkgPath
}