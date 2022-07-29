package api

import jwt "github.com/appleboy/gin-jwt/v2"

// Routes : ...
func (s *Server) Routes(authMw *jwt.GinJWTMiddleware) {
	s.g.GET("/", s.DefaultWelcome)

	api := s.g.Group("/api")
	{

		api.GET("/", s.Welcome)

		auth := api.Group("/auth")
		auth.POST("/login", s.RequestLogin)

		// user API group
		profile := api.Group("/user")
		profile.Use(authMw.MiddlewareFunc())
		{
			profile.GET("/profile", s.GetUserProfile)
		}
	}
}
