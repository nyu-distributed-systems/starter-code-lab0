package main

import (
	"log"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/apanda/fa18-lab0/pb"
)

// The struct holding the kv-store data
type KvStore struct {
	store map[string]string
}

func main() {
	const Port = ":3000"
	// First initialize the store
	store := KvStore{store: make(map[string]string)}
	// Create socket that listens on port 3000
	c, err := net.Listen("tcp", Port)
	if err != nil {
		// Note the use of Fatalf which will exit the program after reporting the error.
		log.Fatalf("Could not create listening socket %v", err)
	}
	// Create a new GRPC server
	s := grpc.NewServer()
	// Tell GRPC that s will be serving requests for the KvStore service and should use store (defined on line 23)
	// as the struct whose methods should be called in response.
	pb.RegisterKvStoreServer(s, &store)
	log.Printf("Going to listen on port %v", Port)
	// Start serving, this will block this function and only return when done.
	if err := s.Serve(c); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
	log.Printf("Done listening")
}

// Handle the Get RPC from the KvStore servie
func (s *KvStore) Get(ctx context.Context, key *pb.Key) (*pb.KeyValue, error) {
	// The bit below works because Go maps return the 0 value for non existent keys, which is empty in this case.
	return &pb.KeyValue{Key: key.Key, Value: s.store[key.Key]}, nil
}
