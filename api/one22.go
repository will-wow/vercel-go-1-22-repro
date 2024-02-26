package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func One22(w http.ResponseWriter, r *http.Request) {

	var response []string

	for i := range 5 {
		response = append(response, strconv.Itoa(i))
	}

	fmt.Fprint(w, strings.Join(response, ","))
}
