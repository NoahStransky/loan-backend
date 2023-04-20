package src

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"xorm.io/xorm"

	"bernie.top/demyst/loan-backend/src/controller"
	"bernie.top/demyst/loan-backend/src/util"
)

func Serve() {

	engine := SetupDatabase()

	accountingProvider := controller.NewAccountingProvider(engine)

	r := gin.Default()

	r.GET("/v1/accounting-providers", func(c *gin.Context) {
		providers, err := accountingProvider.GetProviders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, providers)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupDatabase() *xorm.Engine {
	appEnv := os.Getenv("ENV")
	driveName := viper.GetString("PG.driveName")
	host := viper.GetString("PG.host")
	port := viper.GetString("PG.port")
	user := viper.GetString("PG.user")
	password := viper.GetString("PG.password")
	database := viper.GetString("PG.database")
	if appEnv == "production" {
		host = os.Getenv("PG_HOST")
		port = os.Getenv("PG_PORT")
		user = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		database = os.Getenv("PG_DATABASE")
	}
	return util.CreateConnection(driveName, host, port, user, password, database)
}
