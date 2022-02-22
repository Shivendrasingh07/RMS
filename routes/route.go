package routes

import (
	"github.com/RMS/handler"
	"github.com/RMS/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	chi.Router
}

func Route() *Server {

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		//ADMINS
		r.Route("/admin", func(admin chi.Router) {
			admin.Post("/login", handler.Login)
			admin.Post("/signup", handler.AdminSignup)
			admin.Post("/reset", handler.ResetPassword)
			admin.Route("/", func(r chi.Router) {
				r.Use(middleware.AuthMiddleware)
				r.Post("/subadmin_signup", handler.SubadminSignup)
				r.Post("/user_signup", handler.UserSignup)
				r.Get("/allsubadmin", handler.Allsubadmins)
			})
		})
		//SUBADMINS
		r.Route("/subadmin", func(subadmin chi.Router) {
			subadmin.Post("/login", handler.Login)
			subadmin.Post("/reset", handler.ResetPassword)
			subadmin.Route("/", func(r chi.Router) {
				r.Use(middleware.AuthMiddleware)
				r.Post("/user_signup", handler.UserSignup)
				r.Get("/allusers", handler.Allusers)
			})

		})

		// USRES
		r.Route("/user", func(user chi.Router) {
			user.Post("/login", handler.Login)
			user.Post("/reset", handler.ResetPassword)
			user.Post("/distance", handler.Distance)

		})

		//RESTAURANTS
		r.Route("/resturants", func(r chi.Router) {
			r.Route("/", func(r chi.Router) {
				r.Use(middleware.AuthMiddleware)
				r.Post("/create_rest", handler.CreateResturants)
				r.Post("/create_dishes", handler.CreateDishes)
				r.Get("/allrestaurants", handler.Allrestaurants)
				r.Get("/alldishes", handler.Alldishes)
			})

		})

	})
	return &Server{router}
}

func (svc *Server) Run() error {
	err := http.ListenAndServe(":9000", svc)
	if err != nil {
		return err
	}
	return nil
}
