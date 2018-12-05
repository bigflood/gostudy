package pb_test

import (
	"github.com/bigflood/gostudy/todo/pb"
	fb "github.com/bigflood/gostudy/todo/pb/pb"
	"github.com/golang/protobuf/proto"
	fbs "github.com/google/flatbuffers/go"
	"testing"
)

/*
func BenchmarkEncodingFb2(t *testing.B) {
	tasks := []struct {
		desc string
		done bool
	}{
		{"aaa", false},
		{"aaa1", false},
		{"aaa2", true},
		{"aaa3", false},
	}

	builder := fbs.NewBuilder(0)

	for i := 0; i < t.N; i++ {
		builder.Reset()

		var taskOfsList []fbs.UOffsetT

		for _, task := range tasks {
			desc .. done
		}

		for i, task := range tasks {
			fb.TaskStart(builder)
			fb.TaskAddDesc(builder, descOfs[i])
			fb.TaskAddDone(builder, done)
			taskOfsList = append(taskOfsList, fb.TaskEnd(builder))
		}

		// -- tasks

		fb.ListReplyStartTasksVector(builder, len(tasks))
		for _,v := range taskOfsList {
			builder.PrependUOffsetT(v)
		}
		tasksOfs := builder.EndVector(len(tasks))

		//  --- list reply
		fb.ListReplyStart(builder)
		fb.ListReplyAddTasks(builder, tasksOfs)
		rootOfs := fb.ListReplyEnd(builder)

		builder.Finish(rootOfs)

		//_ := builder.Bytes[builder.Head():]
	}
}
*/

func BenchmarkEncodingFb(t *testing.B) {
	const (
		desc = "desc-abc"
	)

	builder := fbs.NewBuilder(0)

	for i := 0; i < t.N; i++ {
		builder.Reset()

		descOfs := builder.CreateString(desc)

		fb.AddRequestStart(builder)
		fb.AddRequestAddDesc(builder, descOfs)
		rootOfs := fb.AddRequestEnd(builder)

		builder.Finish(rootOfs)

		//_ := builder.Bytes[builder.Head():]
	}
}

func BenchmarkEncodingPb(t *testing.B) {
	const (
		desc = "desc-abc"
	)

	msg := &pb.AddRequest{
		Desc: desc,
	}

	for i := 0; i < t.N; i++ {
		_, err := proto.Marshal(msg)
		if err != nil {
			t.Fatal(err)
		}
	}
}
