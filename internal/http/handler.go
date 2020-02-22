package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kradnoel/cambiomz/internal/graphql"
)

/*var graphQuery = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := r.URL.Query().Get("query")
	if query == "" {
		fmt.Println("error of search...")
		return
	}

	result := graphql.ExecuteQuery(query)
	json.NewEncoder(w).Encode(result)
}

var ping = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Server is alive...")
}*/

/*var index = func(c *gin.Context) {
	var tpl = template.Must(template.ParseFiles("../../README.txt"))
	tpl.Execute(c.Writer, nil)
}*/

var ping2 = func(c *gin.Context) {
	c.JSON(200, "Server is alive")
}

var graphQuery2 = func(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		fmt.Println("error of search...")
		return
	}

	result := graphql.ExecuteQuery(query)
	c.JSON(200, result)
}

/*var Handlers = func() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", ping)
	r.HandleFunc("/api/v1.0.0/exchange", graphQuery)
	return r
}*/
