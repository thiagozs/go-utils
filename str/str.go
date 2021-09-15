package str

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TrimBreakLines(str string) string {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(str, " ")
	return strings.TrimSpace(s)
}

func Float642Str(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64)
}

func Float642StrPrec(value float64, precision int) string {
	//return strconv.FormatFloat(value, 'f', 8, 64)
	return strconv.FormatFloat(value, 'f', precision, 64)
}

func SubStr(input, anchor string) string {
	asRunes := []rune(input)
	lastIndex := strings.LastIndex(input, anchor)
	length := len(anchor)

	if lastIndex < 0 {
		lastIndex = 0
	}
	if lastIndex >= len(asRunes) {
		return ""
	}
	if lastIndex+length > len(asRunes) {
		length = len(asRunes) - lastIndex
	}
	return string(asRunes[lastIndex : lastIndex+length])
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Int642Str(value int64) string {
	return strconv.FormatInt(value, 10)
}
