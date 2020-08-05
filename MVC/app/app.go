package app

import (
	"net/http"

	"github.com/GaryLouisStewart/LearningGo/MVC/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
