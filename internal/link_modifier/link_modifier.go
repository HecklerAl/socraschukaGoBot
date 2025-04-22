package linkmodifier

import (
	md5 "crypto/md5"
	"fmt"
	"strings"

	nsql "gitlab.com/martyn.andrw/microlink/internal/nosql"
)

const (
	st   = 63
	host = "localhost:8082"
)

func ConvertURL(url string) *string {
	splited := strings.Split(url, "/")
	link := splited[len(splited)-1]
	link = "http://" + host + "/" + link
	return &link
}

func ShortenLink(actual string) (*string, error) {
	size, err := nsql.GetLen()
	if err != nil {
		return nil, err
	}

	array := []int{}

	if size == 0 {
		str := "0000000000"
		return &str, nil
	}

	for i := size; i != 0; i /= st {
		array = append(array, (i%st)+48)
	}

	shortLink := ""
	for i := 0; i < len(array); i++ {
		numb := array[i]
		if numb > 57 {
			numb += 7
		}
		if numb > 90 {
			numb += 4
		}
		if numb > 95 {
			numb += 1
		}

		shortLink = string(rune(numb)) + shortLink
	}

	for i := len(shortLink); i <= 10; i++ {
		shortLink = "0" + shortLink
	}

	return &shortLink, nil
}

func LengthenLink(actual string) (*string, error) {
	h := md5.Sum([]byte(actual))
	long := fmt.Sprintf("%x", h)
	return &long, nil
}
