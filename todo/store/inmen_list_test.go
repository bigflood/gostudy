package store_test

import (
	"github.com/bigflood/gostudy/todo/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InmenList", func() {
	var inmem *store.InMem

	BeforeEach(func() {
		inmem = store.NewInMem()
		inmem.Add("task0")
		inmem.Add("task1")
		inmem.Add("task2")
		inmem.Add("task3")
		inmem.Done(0)
		inmem.Done(2)
	})

	It("can filter with done flag", func() {
		done := false
		tasks, err := inmem.List(store.Filter{Done:&done})
		Expect(err).NotTo(HaveOccurred())
		Expect(tasks).To(Equal([]store.Task{
			{Desc:"task1", Done:false},
			{Desc:"task3", Done:false},
		}))
	})

	It("범위를 벗어난 index로 Done함수를 호출하면 실패해야함", func() {
		err := inmem.Done(-1)
		Expect(err).To(HaveOccurred())
		err = inmem.Done(999)
		Expect(err).To(HaveOccurred())
	})
})
