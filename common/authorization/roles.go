package authorization

type AccessibleRoles map[string]map[string][]uint32

/*
	1. Super Admin
	2. Admin
	3. Manager
	4. Executive
	5. Prodi
	6. Alumni
	7. Pengguna Alumni
*/

const (
	BasePath = "tracer_study_grpc"
	PostSvc  = "PostService"
	CommentSvc = "CommentService"
)

var roles = AccessibleRoles{
	"/" + BasePath + "." + PostSvc + "/": {
		"CreatePost":  {1, 2},
		"UpdatePost":  {1, 2},
		"DeletePost":  {1, 2},
	},
	"/" + BasePath + "." + CommentSvc + "/": {
		"DeleteComment": {1, 2},
	},
}

func GetAccessibleRoles() map[string][]uint32 {
	routes := make(map[string][]uint32)

	for service, methods := range roles {
		for method, methodRoles := range methods {
			route := service + method
			routes[route] = methodRoles
		}
	}

	return routes
}
