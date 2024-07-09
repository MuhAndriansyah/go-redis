package book

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/MuhAndriansyah/go-redis-crud/internal/common/response"
	"github.com/MuhAndriansyah/go-redis-crud/pkg/redis"
	rds "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var books []Book

func ListBook(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	redisClient := redis.GetRedis()

	// Ambil data dari cache layer
	result, err := redisClient.Get(ctx, "books").Result()

	if err == rds.Nil {
		serveFromDataSource(ctx, w, redisClient)
		return
	} else if err != nil {
		logrus.Error("Error accessing Redis: ", err)
		response.JSON(w, http.StatusInternalServerError, response.ResponseBody{
			Message: "Internal server error",
		})
		return
	}

	if err := json.Unmarshal([]byte(result), &books); err != nil {
		logrus.Error("Error unmarshaling book data: ", err)
		response.JSON(w, http.StatusInternalServerError, response.ResponseBody{
			Message: "Internal server error",
		})
		return
	}

	response.JSON(w, http.StatusOK, response.ResponseBody{
		Message: "Data from Redis",
		Data:    books,
	})

}

func serveFromDataSource(ctx context.Context, w http.ResponseWriter, redisClient *redis.Client) {
	bookJSON, err := json.Marshal(books)
	if err != nil {
		logrus.Error("Error marshaling books: ", err)
		response.JSON(w, http.StatusInternalServerError, response.ResponseBody{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	cacheDuration := time.Minute * 1

	if err := redisClient.Set(ctx, "books", bookJSON, cacheDuration).Err(); err != nil {
		logrus.Warn("Error setting cache for book list", err)
	}

	response.JSON(w, http.StatusOK, response.ResponseBody{
		Message: "Data from main source",
		Data:    books,
	})

}

func InitializeBooks() {
	filepath := "data/book.json"

	loadedBooks, err := loadBookFromJson(filepath)
	if err != nil {
		logrus.Errorf("Error loading books : %v", err)
		books = make([]Book, 0)
		return
	}
	books = loadedBooks
}

func loadBookFromJson(filepath string) ([]Book, error) {
	var loadedBooks []Book

	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}
	// dari json ke type data struct
	err = json.Unmarshal(data, &loadedBooks)

	if err != nil {
		return nil, err
	}

	return loadedBooks, nil
}
