package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "://github.com"
)

// Global DB pool reused across serverless invocations
var db *sql.DB

func init() {
	var err error
	// standard lib/pq connection string setup
	db, err = sql.Open("postgres", "YOUR_POSTGRES_CONNECTION_STRING")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse Image ID
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	// 2. Fetch MIME Type and BLOB from Postgres
	var mimeType string
	var imgBytes []byte
	query := `SELECT mime_type, img_blob FROM images WHERE id = $1 LIMIT 1`
	
	err = db.QueryRow(query, id).Scan(&mimeType, &imgBytes)
	if err == sql.ErrNoRows {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// 3. Apply Fast HTTP Caching Headers (Vercel Edge Network)
	// Cache the image at the Edge for 31 days (2678400 seconds)
	w.Header().Set("Cache-Control", "public, max-age=2678400, s-maxage=2678400, immutable")
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Length", strconv.Itoa(len(imgBytes)))

	// 4. Fast Stream Write to HTTP Network Buffer
	w.Write(imgBytes)
}
