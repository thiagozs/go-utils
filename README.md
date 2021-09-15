# Golang utils

Helper for easily programing in golang

### Methods in package

* ***strs***
  * TrimBreakLines(str string) string
  * Float642Str(value float64) string
  * Float642StrPrec(value float64, precision int) string
  * Str2Float64(value string) (float64, error)
  * Int642Str(value int64) string
  * SubStr(input, anchor string) string

* ***nums***
  * Int64Slice([]int64)
  * Int64Slice([]int64).Sort()
  * SearchInt64s(a []int64, v int64)

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/mercadobitcoin/go-utils/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2021**, thiagozs.
