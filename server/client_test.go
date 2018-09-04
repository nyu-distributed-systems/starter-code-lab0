package main

import (
	"testing"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/apanda/fa18-lab0/pb"
)

func TestGet(t *testing.T) {
	// We assume service is running on localhost:3000
	const Endpoint = "localhost:3000"
	t.Logf("Connecting to %v", Endpoint)
	// Connect to the server. We use WithInsecure since we do not configure https in this class.
	conn, err := grpc.Dial(Endpoint, grpc.WithInsecure())
	//Ensure connection did not fail.
	if err != nil {
		t.Fatalf("Failed to dial GRPC server %v", err)
	}
	t.Logf("Connected")
	// Create a KvStore client
	kvc := pb.NewKvStoreClient(conn)
	// Create a request for the key hello
	req := &pb.Key{Key: "hello"}
	// Send request to server.
	res, err := kvc.Get(context.Background(), req)
	// Ensure request does not fail.
	if err != nil {
		t.Fatalf("Request error %v", err)
	}
	// Make sure we got back the right key.
	if res.Key != "hello" {
		t.Errorf("Expected key \"hello\" got \"%v\"", res.Key)
	}
	// Done
	t.Logf("Got response key:\"%v\" value:\"%v\"", res.Key, res.Value)
}
