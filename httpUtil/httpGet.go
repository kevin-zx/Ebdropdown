package httpUtil

import (
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var c []byte
	for {
		buf := make([]byte, 1024)
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		c = append(c, buf...)
	}
	return string(c)
}
