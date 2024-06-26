package routes

import (
	"context"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func CreateRoutes(ctx context.Context, neo, mongo string, redis *redis.Client) error {
	router := mux.NewRouter()

	// TEST ENDPOINTS
	// =====================

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(handlers.TestCachePostHandler, redis, ctx)).Methods("POST")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(handlers.TestCacheGetHandler, redis, ctx)).Methods("GET")

	// APPLICATION ENDPOINTS
	// =====================

	// no name, no tags, time frame != mostRecent, query cache (Redis? Mongo?)
	router.HandleFunc("/api/search/cache", middleware.AddRedisDriver(handlers.SearchCacheHandler, redis, ctx)).Methods("GET")

	// timeFrame == mostRecent || name || tags
	router.HandleFunc("/api/search/standard", middleware.AddNeoDriver(handlers.SearchStandardHandler, neo)).Methods("GET")

	// time frame == mostRecent - query neo current year, order by date, not rank
	router.HandleFunc("/api/search/most-recent", middleware.AddNeoDriver(handlers.SearchMostRecentHandler, neo)).Methods("GET")

	// name || tags && timeFrame == allTime - query neo allTime db, order by absolute rank
	router.HandleFunc("/api/search/all-time", middleware.AddNeoDriver(handlers.SearchAllTimeHandler, neo)).Methods("GET")

	// timeFrame == previous year - query neo specific year, order by rank && rel_rank - DEPRECATED
	// router.HandleFunc("/api/search/year", middleware.AddNeoDriver(handlers.SearchYearHandler, neo)).Methods("GET")

	return http.ListenAndServe(":8080", router)
}
