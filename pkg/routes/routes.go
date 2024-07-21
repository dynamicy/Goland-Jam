package routes

import (
	"Goland-Jam/pkg/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SetupRoutes(client *mongo.Client) {
	memberController := controllers.NewMemberController(client)

	// 健康檢查路由
	http.HandleFunc("/health", controllers.HealthCheckHandler)

	// 會員相關路由
	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			memberController.CreateMember(w, r)
		} else {
			memberController.ListMembers(w, r)
		}
	})
	http.HandleFunc("/member", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			memberController.GetMember(w, r)
		case http.MethodPut:
			memberController.UpdateMember(w, r)
		case http.MethodDelete:
			memberController.DeleteMember(w, r)
		}
	})
}
