package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mclellac/ramify/data"
	"github.com/mclellac/ramify/services/auth"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authServer struct {
	users map[string]*auth.User
}

// VerifyToken returns a user from authentication token.
func (s *authServer) VerifyToken(ctx context.Context, req *auth.Request) (*auth.Result, error) {
	md, _ := metadata.FromContext(ctx)
	traceID := strings.Join(md["traceID"], ",")

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	user := s.users[req.AuthToken]
	if user == nil {
		return &auth.Result{}, errors.New("Invalid Token")
	}

	reply := new(auth.Result)
	reply.User = user
	return reply, nil
}

// loadUsers loads users from a JSON file.
func loadUserData(path string) map[string]*auth.User {
	file := data.MustAsset(path)
	users := []*auth.User{}

	// unmarshal JSON
	if err := json.Unmarshal(file, &users); err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	// create user lookup map
	cache := make(map[string]*auth.User)
	for _, c := range users {
		cache[c.AuthToken] = c
	}
	return cache
}

func main() {
	var port = flag.Int("port", 3000, "The server port")
	flag.Parse()

	// listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpc server
	srv := grpc.NewServer()
	auth.RegisterAuthServer(srv, &authServer{
		users: loadUserData("data/users.json"),
	})
	srv.Serve(lis)
}
