package middlewares

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/url-shortener/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
	"log"
	"net/http"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:4001", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func runAuthenticateRoute(client auth.AuthClient, token string) *auth.AuthenticateResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	in := &auth.AuthenticateRequest{Token: token}
	authResponse, err := client.Authenticate(ctx, in)
	if err != nil {
		log.Fatalf("client.Authenticate failed: %v", err)
	}

	return authResponse
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-Auth-Token")
		if tokenString == "" {
			log.Print("empty token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			return
		}

		flag.Parse()
		var opts []grpc.DialOption
		if *tls {
			if *caFile == "" {
				*caFile = data.Path("x509/ca_cert.pem")
			}
			creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
			if err != nil {
				log.Fatalf("Failed to create TLS credentials: %v", err)
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.Dial(*serverAddr, opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := auth.NewAuthClient(conn)

		authResponse := runAuthenticateRoute(client, tokenString)

		if !authResponse.Authenticated {
			log.Print("invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			return
		}

		b := false
		for _, role := range authResponse.User.Roles {
			if role == "teacher" {
				b = true
				break
			}
		}

		if !b {
			log.Print("user has not permission")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Set("user", authResponse.User)
		c.Next()
	}
}
