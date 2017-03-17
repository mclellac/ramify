package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	ui "github.com/mclellac/amity/lib/ui"
	pb "github.com/mclellac/ramify/services/post"

	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	address    = "localhost:3001"
	defaultNum = 0
)

func add(client pb.PostClient, title string, article string) error {
	post := &pb.Content{
		Title:   title,
		Article: article,
	}
	res, err := client.Add(context.Background(), post)
	fmt.Println(res)

	return err
}

func delete(client pb.PostClient, id int64) error {
	post := &pb.Content{
		Id: id,
	}

	res, err := client.Delete(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	return nil
}

func list(client pb.PostClient) error {
	stream, err := client.List(context.Background(), new(pb.Request))
	if err != nil {
		return err
	}
	for {
		post, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		grpclog.Printf("%s %s %s", ui.Cyan, post, ui.Reset)
	}
	return nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPostClient(conn)

	app := cli.NewApp()
	app.Name = "ramify"
	app.Usage = "Ramify client."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "localhost:3001", Usage: "server host"},
	}

	app.Commands = []cli.Command{
		{
			Name:        "add",
			Usage:       "Create a new post.",
			Description: "Adds new article to the database.\n\nEXAMPLE:\n   $ ramify add \"Test Title\" \"Test article body...\"",
			ArgsUsage:   "[\"title\"] [\"article\"]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) != 2 {
					fmt.Println("You might want to double check your command there.")
					return nil
				}

				title := c.Args().Get(0)
				article := c.Args().Get(1)
				if err != nil {
					return err
				}
				return add(client, title, article)
			},
		},
		{
			Name:        "ls",
			Usage:       "List all posts.",
			Description: "Displays the IDs and titles of articles on the server.\n\nEXAMPLE:\n   $ ramify ls",
			Action: func(c *cli.Context) error {
				return list(client)
			},
		},
		{
			Name:        "rm",
			Usage:       "Delete an article.",
			Description: "Remove the post with the supplied ID from the server.\n\nEXAMPLE:\n   $ ramify rm 2",
			ArgsUsage:   "[ID]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					fmt.Println("What is it that you are trying to delete... air?")
					return nil
				}

				idStr := c.Args().Get(0)

				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return nil
				}

				return delete(client, int64(id))

				return nil
			},
		},
	}
	app.Run(os.Args)
}
