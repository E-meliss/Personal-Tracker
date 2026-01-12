package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "personal-tracker/backend/internal/db"
    "personal-tracker/backend/internal/handlers"
    "personal-tracker/backend/internal/middleware"
)

func main() {
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "./tracker.db"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    database, err := db.Open(dbPath)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.CORS())

    h := handlers.TaskHandlers{DB: database}
    h.Register(r)

    addr := ":" + port
    log.Printf("API running on %s (DB: %s)", addr, dbPath)
    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
