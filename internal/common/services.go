package common

import (
	"path/filepath"
	"strings"
	"unicode"

	"github.com/adem02/epse/internal/utils/osutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetSrcPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src")
}

func GetFileOrDirectoryFromProjectRootPath(paths ...string) string {
	return filepath.Join(append([]string{osutils.GetCurrentDirPath()}, paths...)...)
}

func GetFileOrDirectoryPathFromSrcPath(paths ...string) string {
	return filepath.Join(append([]string{GetSrcPath()}, paths...)...)
}

func ToPascalCase(words []string) string {
	if len(words) == 0 {
		return ""
	}

	caser := cases.Title(language.Und, cases.NoLower)

	var b strings.Builder
	for _, word := range words {
		if word == "" {
			continue
		}
		b.WriteString(caser.String(strings.ToLower(word)))
	}

	return b.String()
}

func ToKebabCase(words []string) string {
	return strings.Join(words, "-")
}

func SplitCamelOrPascal(s string) string {
	if s == "" {
		return s
	}

	var b strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if i > 0 {
			prev := runes[i-1]

			if unicode.IsLower(prev) && unicode.IsUpper(r) {
				b.WriteRune(' ')
			}

			if i+1 < len(runes) &&
				unicode.IsUpper(prev) &&
				unicode.IsUpper(r) &&
				unicode.IsLower(runes[i+1]) {
				b.WriteRune(' ')
			}
		}

		b.WriteRune(r)
	}

	return b.String()
}
