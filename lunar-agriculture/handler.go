package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {
	moonFuncURL := "http://127.0.0.1:8080/function/sriveros95/moon-phase"
	earthTime := time.Now()
	var err interface{}
	if len(string(req)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(req))
		if err != nil {
			// todo log error
			return "Failed lunar-agriculture time.Parse"
		}
	}
	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	if err != nil {
		fmt.Println(err)
		return "Failed http.Post"
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ",")
	fmt.Println(moonData)
	return ""
}
