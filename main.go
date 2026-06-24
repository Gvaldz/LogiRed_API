package main

import (
	"logired/src/cmd"
	_ "logired/docs" 
)

// @title           LogiRed API
// @version         1.0
// @description     API para la plataforma móvl LogiRed
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
// @host            api-logired.shop
// @BasePath        /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Escribe "Bearer <token>" para autenticarte

func main() {
	cmd.Init()
}