package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	moonFuncURL := "http://127.0.0.1:8080/function/sriveros95-moon-phase"
	resp, err := http.Get(moonFuncURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed http.Get"))
		return
	}
	defer resp.Body.Close()
	moonString, err := ioutil.ReadAll(resp.Body)
	fmt.Println(moonString)
	moonData := strings.Split(string(moonString), ", ")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was: %s \n MoonString: %s \n %v", string(input), moonString, moonData)))
}
