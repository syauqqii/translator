package helper

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

type UserData struct {
	Limiter   *rate.Limiter
	LastReset time.Time
	Requests  int
}

var (
	userDataMap = make(map[string]*UserData)
	mu          sync.Mutex
	limit       int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	limit, err = strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		fmt.Print(err.Error())
		limit = 0
	}
}

func RateLimitedHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("Username")

		userData := getUserData(username)

		mu.Lock()
		defer mu.Unlock()

		if limit != 0 && time.Since(userData.LastReset) >= time.Minute {
			userData.LastReset = time.Now()
			userData.Requests = 0
		}

		if limit == 0 || userData.Requests < limit {
			userData.Requests++
			next(w, r)
		} else {
			http.Error(w, " ! ERROR: Too many requests.", http.StatusTooManyRequests)
		}
	}
}

func getUserData(username string) *UserData {
	mu.Lock()
	defer mu.Unlock()

	userData, ok := userDataMap[username]
	if !ok {
		var limiter *rate.Limiter
		if limit != 0 {
			limiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(limit)), limit)
		} else {
			limiter = rate.NewLimiter(rate.Inf, 0)
		}
		userData = &UserData{
			Limiter:   limiter,
			LastReset: time.Now(),
			Requests:  0,
		}
		userDataMap[username] = userData
	}

	return userData
}

func SetRateLimit(newLimit int) {
	mu.Lock()
	defer mu.Unlock()
	limit = newLimit
}
