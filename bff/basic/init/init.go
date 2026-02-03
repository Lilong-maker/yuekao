package init

import (
	"flag"
	"log"
	"yuekao/bff/basic/config"
	__ "yuekao/srv/basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	initDB()
	InitViper()
	InitRedis()
	InitMysql()
	InitEs()
	//middleware.CreateFile()
}
func initDB() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.UserClient = __.NewUserClient(conn)

}
