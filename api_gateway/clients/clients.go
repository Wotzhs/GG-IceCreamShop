package clients

import (
	"log"
	"os"
	"proto/auth"
	"proto/ice_cream"
	"proto/user"

	"google.golang.org/grpc"
)

var (
	Auth     auth.AuthClient
	IceCream ice_cream.IceCreamClient
	User     user.UserClient

	AuthClientAddr     string
	IceCreamClientAddr string
	UserClientAddr     string
)

func init() {
	if AuthClientAddr = os.Getenv("AUTH_CLIENT_ADDR"); AuthClientAddr == "" {
		AuthClientAddr = "0.0.0.0:50051"
	}

	if IceCreamClientAddr = os.Getenv("ICE_CREAM_CLIENT_ADDR"); IceCreamClientAddr == "" {
		IceCreamClientAddr = "0.0.0.0:50052"
	}

	if UserClientAddr = os.Getenv("USER_CLIENT_ADDR"); UserClientAddr == "" {
		UserClientAddr = "0.0.0.0:50053"
	}
}

func RegisterGrpcServices() []*grpc.ClientConn {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	AuthConn, err := grpc.Dial(AuthClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to auth service err: %v", err)
	}

	IceCreamConn, err := grpc.Dial(IceCreamClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to ice cream service err: %v", err)
	}

	UserConn, err := grpc.Dial(UserClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to user service err: %v", err)
	}

	Auth = auth.NewAuthClient(AuthConn)
	IceCream = ice_cream.NewIceCreamClient(IceCreamConn)
	User = user.NewUserClient(UserConn)

	return []*grpc.ClientConn{AuthConn, IceCreamConn, UserConn}
}
