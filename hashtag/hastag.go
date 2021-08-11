package hashtag

import (
	"regexp"
	"strings"
)

// Ex.: "Hello World!" => "#helloworld"
func create(i string) string {
	i = strings.ToLower(i)
	i = strings.Replace(i, " ", "", -1)
	re := regexp.MustCompile(`[a-zA-Z]*`)
	matched := re.Find([]byte(i))
	return "#" + string(matched)
}
