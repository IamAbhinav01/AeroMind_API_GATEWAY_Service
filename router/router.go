package router


type Router interface{
	Register()
}

func SetupRouter(UserRouter Router) *chi.Mux{
	
}