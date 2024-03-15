package helper

import (
	"net/http"
	"sync"
	"time"

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
	limit       = 30
)

func RateLimitedHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("Username")

		userData := getUserData(username)

		mu.Lock()
		defer mu.Unlock()

		if time.Since(userData.LastReset) >= time.Minute {
			userData.LastReset = time.Now()
			userData.Requests = 0
		}

		if userData.Requests < limit {
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
		limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(limit)), limit)
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
