package function

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	moonFuncURL := "http://127.0.0.1:8080/function/sriveros95/moon-phase"
	earthTime := time.Now()
	var err interface{}
	if len(string(input)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(input))
		if err != nil {
			// todo log error
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed lunar-agriculture time.Parse"))
			return
		}
	}
	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed http.Post"))
		return
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ", ")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was: %s \n MoonString: %s \n %v", string(input), moonString, moonData)))
}
