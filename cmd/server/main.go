package main

import (
	"fmt"
	"tracerstudy-post-service/common/config"

	gormConn "tracerstudy-post-service/common/gorm"
	commonJwt "tracerstudy-post-service/common/jwt"
	"tracerstudy-post-service/common/mysql"
	"tracerstudy-post-service/server"

	postModule "tracerstudy-post-service/modules/post"
	commentModule "tracerstudy-post-service/modules/comment"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	splash(cfg)

	// http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/uploads"))))
	// log.Fatal(http.ListenAndServe(":8080", nil))

	dsn, derr := mysql.NewPool(&cfg.MySQL)
	checkError(derr)
	// errUtils.ConvertToRestError(derr)

	db, gerr := gormConn.NewMySQLGormDB(dsn)
	checkError(gerr)
	// errUtils.ConvertToRestError(gerr)

	jwtManager := commonJwt.NewJWT(cfg.JWT.JwtSecretKey, cfg.JWT.TokenDuration)

	grpcServer := server.NewGrpcServer(cfg.Port.GRPC, jwtManager)
	grpcConn := server.InitGRPCConn(fmt.Sprintf("127.0.0.1:%v", cfg.Port.GRPC), false, "")

	registerGrpcHandlers(grpcServer.Server, *cfg, db, grpcConn)

	_ = grpcServer.Run()
	_ = grpcServer.AwaitTermination()
}

func registerGrpcHandlers(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	postModule.InitGrpc(server, cfg, db, grpcConn)
	commentModule.InitGrpc(server, cfg, db, grpcConn)
}

// func createRestServer(port string) *server.Rest {
// 	return server.NewRest(port)
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func splash(cfg *config.Config) {
	version := "1.0.0"
	colorReset := "\033[0m"
	// colorBlue := "\033[34m"
	colorCyan := "\033[36m"

	fmt.Printf(`
    __                                                    __                    __
  _/  |_ _______ ______   ____  ______ _______     ______/  |__  __   __       |  \ __   __
  \   __\\_  __ \\___  \_/ ___\/  __  \\_  __ \   /  ___/\   __\|  | |  |   ___|  ||  | |  |
   |  |   |  | \/  / ___\  \___\   ___/ |  | \/   \___ \  |  |  |  |_|  | /  __   ||  |_|  | >>
   |__|   |__|    (____  /\__  >\___  > |__|     /____  > |__|  | ______|(  ____  )\__   __/
                       \/    \/     \/                \/        \/        \/    \/ _ /  /    v%s
                                                                                  / ___/
	`, version)

	// fmt.Println(colorBlue, fmt.Sprintf(`⇨ REST server started on port :%s`, cfg.Port.REST))
	fmt.Println(colorCyan, fmt.Sprintf(`⇨ GRPC post service server started on port :%s`, cfg.Port.GRPC))
	fmt.Println(colorReset, "")
}
