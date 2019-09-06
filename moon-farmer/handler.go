package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		// read request payload
		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

			input = reqBody
		}
	}

	moonFuncURL := "http://127.0.0.1:8080/function/sriveros95/moon-phase"
	earthTime := time.Now()
	var err interface{}
	if len(string(input)) != 0 {
		earthTime, err = time.Parse(time.RFC3339, string(input))
		if err != nil {
			// todo log error
			fmt.Println(err)
			fmt.Println("Failed lunar-agriculture time.Parse")
			return
		}
	}
	resp, err := http.Post(moonFuncURL, "string", bytes.NewReader([]byte(earthTime.Format(time.RFC3339))))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ",")
	fmt.Println(moonData)

	response := struct {
		Payload     string              `json:"payload"`
		Headers     map[string][]string `json:"headers"`
		Environment []string            `json:"environment"`
	}{
		Payload:     string(moonString),
		Headers:     r.Header,
		Environment: os.Environ(),
	}

	resBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write result
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
