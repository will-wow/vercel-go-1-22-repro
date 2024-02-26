package one22

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var response []string

	for i := range 5 {
		response = append(response, strconv.Itoa(i))
	}

	fmt.Fprint(w, strings.Join(response, ","))
}
