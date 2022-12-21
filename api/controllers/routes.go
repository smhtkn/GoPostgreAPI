package controllers

import "github.com/smhtkn/testpostgre/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Clients routes
	s.Router.HandleFunc("/clients", middlewares.SetMiddlewareJSON(s.CreateClient)).Methods("POST")
	s.Router.HandleFunc("/clients", middlewares.SetMiddlewareJSON(s.GetClients)).Methods("GET")
	s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareJSON(s.GetClient)).Methods("GET")
	s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateClient))).Methods("PUT")
	s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteClient)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
