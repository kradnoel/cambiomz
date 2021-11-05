package http

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func init() {
	godotenv.Load(".env")
}

type Server struct {
	Address   string
	RateLimit string
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
	rateLimit := os.Getenv("RATE_LIMIT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if port != "" {
		s.Address = ":" + port
	}

	if rateLimit == "" {
		log.Fatal("$RATELIMIT must be set")
	}

	if rateLimit != "" {
		s.RateLimit = rateLimit
	}
}

func (s *Server) Run() {

	rate, err := limiter.NewRateFromFormatted("5-S")
	if err != nil {
		log.Fatal(err)
		return
	}

	store := memory.NewStore()
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	r := gin.Default()
	r.Use(middleware)

	/*r.GET("/", index)*/
	r.GET("/health", ping2)
	r.GET("/api/v1.0.0/exchange", graphQuery2)

	r.Run(s.Address)
}
