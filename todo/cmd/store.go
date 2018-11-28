package cmd

import (
	"context"
	"github.com/bigflood/gostudy/todo/pb"
	"github.com/bigflood/gostudy/todo/store"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type Store interface {
	Add(desc string) error
	List(filter store.Filter) ([]store.Task, error)
	Done(index int) error
}

func OpenStore(cmd *cobra.Command) (Store, error) {
	grpcAddr := cmd.Flag("grpc").Value.String()
	if grpcAddr != "" {
		conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}

		return grpcAdaptor{pb.NewTodoClient(conn)}, nil
	}

	fileName := cmd.Flag("file").Value.String()
	return store.NewInFile(fileName), nil
}

type grpcAdaptor struct {
	client pb.TodoClient
}

func (adaptor grpcAdaptor) Add(desc string) error {
	_, err := adaptor.client.Add(context.Background(), &pb.AddRequest{Desc: desc})
	return err
}

func (adaptor grpcAdaptor) List(filter store.Filter) ([]store.Task, error) {

	req := &pb.ListRequest{}

	switch {
	case filter.Done != nil && *filter.Done:
		req.DoneFilter = pb.DoneFilter_DONE
	case filter.Done != nil && !*filter.Done:
		req.DoneFilter = pb.DoneFilter_NOT_DONE
	}

	reply, err := adaptor.client.List(context.Background(), req)
	if err != nil {
		return nil, err
	}

	tasks := make([]store.Task, len(reply.Tasks))
	for i, task := range reply.Tasks {
		tasks[i] = store.Task{Done: task.Done, Desc: task.Desc}
	}

	return tasks, err
}

func (adaptor grpcAdaptor) Done(index int) error {
	_, err := adaptor.client.Done(context.Background(), &pb.DoneRequest{Index: int32(index)})
	return err
}
