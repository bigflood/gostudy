package store_test

import (
	"github.com/bigflood/gostudy/todo/store"
	"github.com/bigflood/gostudy/todo/store/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InFileList", func() {
	var obj *store.InFile

	BeforeEach(func() {
		data := []byte(`{}`)

		file := &mocks.File{}
		file.ReadAllCalls(func() ([]byte, error) {
			return data, nil
		})
		file.WriteAllCalls(func(bytes []byte) error {
			data = bytes
			return nil
		})

		obj = store.NewFromDataSource(file)
		obj.Add("task0")
		obj.Add("task1")
		obj.Add("task2")
		obj.Add("task3")
		obj.Done(0)
		obj.Done(2)
	})

	It("can filter with done flag", func() {
		done := false
		tasks, err := obj.List(store.Filter{Done: &done})
		Expect(err).NotTo(HaveOccurred())
		Expect(tasks).To(Equal([]store.Task{
			{Desc: "task1", Done: false},
			{Desc: "task3", Done: false},
		}))
	})
})
