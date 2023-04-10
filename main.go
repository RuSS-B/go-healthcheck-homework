package main

import (
	"dockerHomework/handler"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"os"
	"time"
)

var db *gorm.DB

func main() {
	LoadEnv()
	port := 8000

	db = ConnectDB(os.Getenv("DATABASE_URL"))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router(),
	}

	log.Printf("Starting server on port: %d", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func router() http.Handler {
	h := handler.NewHandlerRepository(db)
	mux := chi.NewRouter()

	mux.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	})

	mux.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)

			response := struct {
				Status string `json:"status"`
			}{
				Status: "OK",
			}

			data, _ := json.Marshal(response)

			_, _ = w.Write(data)
		})
	})

	mux.Route("/users", h.Users)

	return mux
}

func ConnectDB(dsn string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	if os.Getenv("APP_ENV") == "prod" {
		newLogger = nil
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Connection err:", err)
	}

	return db
}

func LoadEnv() {
	fileName := ".env"
	err := godotenv.Load(fileName)
	if err != nil {
		log.Println("Wasn't able to load the file", fileName)
	} else {
		log.Println("File found, envs are loaded")
	}
}
