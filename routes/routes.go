package routes

import (
	"github.com/go-chi/chi"
	"sugar/pkg/auth"
	"sugar/pkg/doors"
	"sugar/pkg/permissions"
	"sugar/pkg/users"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/", func(r chi.Router) {

		r.Post("/auth/login", auth.Login)
		r.Post("/auth/setpassword", auth.SetPassword)

		r.Post("/door/action", Authentication(doors.Action))
		r.Get("/door/list", Authentication(doors.MyDoorsList))

		//Admin APIs
		r.Route("/admin", func(r chi.Router) {
			r.Use(AdminOnly)
			r.Post("/door/add", doors.AddDoor)
			r.Post("/door/remove", doors.RemoveDoor)

			r.Post("/user/add", users.AddUser)
			r.Post("/user/remove", users.RemoveUser)

			r.Post("/permission/add", permissions.AddPermission)
			r.Post("/permission/remove", permissions.RemovePermission)

		})

	})
	return r
}
