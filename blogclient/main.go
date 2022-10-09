package main

import (
	"context"
	"fmt"
	"log"

	//  Cosmos blockchain client for RPC connection
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"blog/x/blog/types"
)

func main() {
	// Prefix to use for account addresses.
	// The address prefix was assigned to the blog blockchain
	// using the `--address-prefix` flag during scaffolding.
	addressPrefix := "blog"

	// Create a Cosmos client instance
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"
	//accountAddress := "blog1psw5nn4aja8ydr980d3vnag2ddv2apmsqaaaqh"

	// Get account from the keyring
	address, err := cosmos.Address(accountName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address)

	// addr := account.Address(addressPrefix)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Define a message to create a post
	msg := &types.MsgCreatePost{
		Creator: address.String(),
		Title:   "Hello!",
		Body:    "This is the first post",
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(cosmos.Context())

	// Query the blockchain using the client's `Posts` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
