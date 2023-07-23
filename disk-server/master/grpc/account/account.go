package account

import (
	"disk-server/master/global"
	account "disk-server/master/grpc/proto/account"
	"fmt"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var AccountClient account.UserServiceClient

func InitClient() error {
	grpcAddress := fmt.Sprintf("%s:%d", global.Config.AccountConfig.Host, global.Config.AccountConfig.Port)
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect of iot grpc: %v", err)
		return err
	}
	AccountClient = account.NewUserServiceClient(conn)
	return nil
}