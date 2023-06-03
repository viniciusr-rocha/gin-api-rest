package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	StudentRoutes(r)

	err := r.Run()
	if err != nil {
		log.Panic("Ocorreu um erro ao iniciar o servidor -", err)
	}
}
