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

	authClientAddr     string
	iceCreamClientAddr string
	userClientAddr     string
)

func init() {
	if authClientAddr = os.Getenv("AUTH_CLIENT_ADDR"); authClientAddr == "" {
		authClientAddr = "0.0.0.0:50051"
	}

	if iceCreamClientAddr = os.Getenv("ICE_CREAM_CLIENT_ADDR"); iceCreamClientAddr == "" {
		iceCreamClientAddr = "0.0.0.0:50052"
	}

	if userClientAddr = os.Getenv("USER_CLIENT_ADDR"); userClientAddr == "" {
		userClientAddr = "0.0.0.0:50053"
	}
}

func RegisterGrpcServices() []*grpc.ClientConn {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	AuthConn, err := grpc.Dial(authClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to auth service err: %v", err)
	}

	IceCreamConn, err := grpc.Dial(iceCreamClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to ice cream service err: %v", err)
	}

	UserConn, err := grpc.Dial(userClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to user service err: %v", err)
	}

	Auth = auth.NewAuthClient(AuthConn)
	IceCream = ice_cream.NewIceCreamClient(IceCreamConn)
	User = user.NewUserClient(UserConn)

	return []*grpc.ClientConn{AuthConn, IceCreamConn, UserConn}
}
