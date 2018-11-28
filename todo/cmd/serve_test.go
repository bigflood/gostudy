package cmd_test

import (
	"context"
	"github.com/bigflood/gostudy/todo/cmd"
	"github.com/bigflood/gostudy/todo/cmd/mocks"
	"github.com/bigflood/gostudy/todo/pb"
	"github.com/bigflood/gostudy/todo/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Serve", func() {

	var (
		fakeStore *mocks.Store
		svr       *cmd.TodoServer
		done      = true
		notDone   = false
	)

	BeforeEach(func() {
		fakeStore = &mocks.Store{}
		svr = cmd.NewTodoServer(fakeStore)
	})

	DescribeTable("done filter가 포함된 List 요청이 store로 잘 전달됨",
		func(input pb.DoneFilter, expected store.Filter) {
			_, err := svr.List(context.Background(), &pb.ListRequest{
				DoneFilter: input,
			})

			Expect(err).NotTo(HaveOccurred())
			filter := fakeStore.ListArgsForCall(0)
			Expect(filter).To(Equal(expected))
		},
		Entry("none", pb.DoneFilter_NONE, store.Filter{}),
		Entry("done", pb.DoneFilter_DONE, store.Filter{Done: &done}),
		Entry("not done", pb.DoneFilter_NOT_DONE, store.Filter{Done: &notDone}),
	)
})
