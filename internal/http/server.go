package http

import (
	"os"
 "log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/heroku/x/hmetrics/onload"
)

func init() {
	godotenv.Load(".env")
}

type Server struct {
	Address string
	//Handler *mux.Router
}

func New() Server {
	s := Server{}
	return s
}

/*func (s *Server) Default() {
	s.Address = ":3000"
	s.Handler = Handlers()
}

func (s *Server) Run() {
	fmt.Println("Server is running on port ", s.Address)

	server := http.Server{
		Addr:         s.Address,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      s.Handler,
	}
	server.ListenAndServe()
}*/

func (s *Server) Default() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if port != "" {
		s.Address = ":" + port
	}
}

func (s *Server) Run() {
	r := gin.Default()
	/*r.GET("/", index)*/
	r.GET("/health", ping2)
	r.GET("/api/v1.0.0/exchange", graphQuery2)

	r.Run(s.Address)
}
