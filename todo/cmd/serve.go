// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/bigflood/gostudy/todo/pb"
	"github.com/bigflood/gostudy/todo/store"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := OpenStore(cmd)
		if err != nil {
			return err
		}

		s = NewLogging(s)

		addr := cmd.Flag("addr").Value.String()

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}

		svr := grpc.NewServer()

		pb.RegisterTodoServer(svr, NewTodoServer(s))

		log.Println("listen", addr)
		return svr.Serve(listener)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().String("addr", ":8082", "listen address")
}

func NewTodoServer(file Store) *TodoServer {
	return &TodoServer{file}
}

type TodoServer struct {
	file Store
}

func (svr *TodoServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{}, svr.file.Add(req.Desc)
}

func (svr *TodoServer) Done(ctx context.Context, req *pb.DoneRequest) (*pb.DoneReply, error) {
	return &pb.DoneReply{}, svr.file.Done(int(req.Index))
}

func (svr *TodoServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	reply := &pb.ListReply{}

	filter := store.Filter{}

	switch req.DoneFilter {
	case pb.DoneFilter_DONE:
		v := true
		filter = store.Filter{Done: &v}
	case pb.DoneFilter_NOT_DONE:
		v := false
		filter = store.Filter{Done: &v}
	}

	tasks, err := svr.file.List(filter)
	if err != nil {
		return reply, err
	}

	reply.Tasks = make([]*pb.Task, len(tasks))
	for i, task := range tasks {

		reply.Tasks[i] = &pb.Task{Done: task.Done, Desc: task.Desc}
	}

	return reply, err
}
