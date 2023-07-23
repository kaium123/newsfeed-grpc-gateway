package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/route"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func Start() {
	router := route.Setup()
	port := viper.GetString("SERVER_PORT")

	logger.LogInfo("Running server on port: " + port)

	/*g := goGen.NewGenerator(goGen.Config{
		OutPath: "./dal/query",
	})
	db := dic.Container.Get(dic.EntDbService).(*gorm.DB)
	g.UseDB(db)
	g.ApplyBasic(entity.User{})

	//g.ApplyInterface(func(method models.Method) {}, models.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.Execute()*/

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	//router.Run(":" + viper.GetString("SERVER_PORT"))
	go func() {

		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// send signal to worker to shutdown
	//cleanup
	//wait for worker to finish stopping
	logger.LogInfo("Shutting down server...")

	//shutdown gin
	serverCloseContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("getting ready for pettycash server shutdown")
	if err := server.Shutdown(serverCloseContext); err != nil {
		log.Fatal("Server Shutdown err:", err)
	}

	<-serverCloseContext.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("bye bye")
}
