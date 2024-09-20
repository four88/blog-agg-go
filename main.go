package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"github.com/four88/blog-agg-go/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbConn := os.Getenv("DB_CONN")
	if dbConn == "" {
		log.Fatal("DB_CONN environment variable is not set")
	}

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	dbQueries := database.New(db)

	apiCfg := &apiConfig{
		DB: dbQueries,
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mux := http.NewServeMux()

	// API endpoints
	// user route
	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGetInfo))
	mux.HandleFunc("GET /v1/healthz", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandle)

	// feed route
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handleFeedCreate))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handleGetAllFeeds)

	// feed follow route
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowCreate))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollows))
	mux.HandleFunc("DELETE /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowDelete))
	// Start the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on : %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
