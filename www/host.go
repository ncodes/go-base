package www

import (
	"log"
	"strings"

	"github.com/ellcrys/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ncodes/go-base/config"
	"github.com/ncodes/go-base/extend"
	"github.com/ncodes/go-base/lib"
	"gopkg.in/mgo.v2"
)

var (

// mongo params
// MongoDBHosts   = util.Env("MONGO_DB_HOST", "127.0.0.1")

)

// setup middleware, logger etc
func configRouter(router *echo.Echo, testMode bool) {
	if testMode {
		// ... do test setup here
		return
	}
}

// Defines and return an array of policies to pass
// to routes that require authentication to access
func UseAuthPolicy(policyCntrl *lib.PolicyController) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		extend.MiddlewareHandle(policyCntrl.Authenticate),
	}
}

// Fatally exits if an environment variable is unset
func requiresEnv(envName string) {
	if strings.TrimSpace(util.Env(envName, "")) == "" {
		log.Fatal(envName + " environment variable is unset")
	}
}

func App(testMode, runSeed bool) (*echo.Echo, *mgo.Session) {

	// create new server and router
	router := echo.New()
	var v1 = router.Group("/v1")
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	config.HandleError(router)

	// setup router
	configRouter(router, testMode)

	// redis connection
	// redisDB, _ := strconv.Atoi(RedisDatabase)
	// redisPool := GetRedisPool(RedisURL, RedisPassword, redisDB)
	// conn := redisPool.Get()
	// defer conn.Close()
	// if _, err := redisPool.Get().Do("PING"); err != nil {
	// 	util.Println("could not connect to redis database", err)
	// 	conn.Close()
	// 	os.Exit(1)
	// }

	// posgres connection
	// postgresDB, err := GetPostgresConn(PGDatabase, PGHost, PGUsername, PGPassword, PGSSLMODE)
	// if err != nil {
	// 	util.Println("could not connect to postgres database", err)
	// 	conn.Close()
	// 	os.Exit(1)
	// }

	// initialize controllers
	appCntrl := lib.NewAppController()

	// root route
	var rootRoute = v1.Group("/")
	rootRoute.GET("", extend.Handle(appCntrl.Index))

	return router, nil
}
