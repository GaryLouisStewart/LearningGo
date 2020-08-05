package app

import (
	"net/http"

	"github.com/GaryLouisStewart/LearningGo/MVC/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
