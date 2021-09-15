package num

import (
	"fmt"
	"sort"
	"strconv"
)

type Int64Slice []int64

func (s Int64Slice) Len() int           { return len(s) }
func (s Int64Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Int64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Int64Slice) Sort() {
	sort.Sort(s)
}

func SearchInt64s(a []int64, v int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= v })
}

func Str2Float64(value string) (float64, error) {
	fee, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fee, fmt.Errorf("Invalid float value")
	}
	return fee, nil
}

func Int642Str(value int64) string {
	return strconv.FormatInt(value, 10)
}
