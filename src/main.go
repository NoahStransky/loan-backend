package src

import (
	"encoding/json"
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
	handlerRoutes(engine)
}

func handlerRoutes(engine *xorm.Engine) {
	accountingProviderController := controller.NewAccountingProvider(engine)
	balanceController := controller.NewBalance()
	decisionController := controller.NewDecision()

	r := gin.Default()

	r.GET("/v1/accounting-providers", func(c *gin.Context) {
		providers, err := accountingProviderController.GetProviders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"data":    providers,
			"message": "ok",
		})
	})
	r.GET("/v1/get-balance-by-provider", func(c *gin.Context) {
		provider := c.Query("provider")
		balance := balanceController.GetBalanceByProvider(provider)
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"data":    balance,
			"message": "ok",
		})
		return
	})
	r.POST("/v1/make-decision", func(c *gin.Context) {
		var requestBody *controller.MakeDecisionRequestBody
		err := json.NewDecoder(c.Request.Body).Decode(&requestBody)
		defer c.Request.Body.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		decision := decisionController.MakeDecision(requestBody)
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"data":    decision,
			"message": "ok",
		})
		return
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
