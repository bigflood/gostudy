package store_test

import (
	"github.com/bigflood/gostudy/todo/store"
	"github.com/bigflood/gostudy/todo/store/mocks"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Infile", func() {
	var infile *store.InFile

	BeforeEach(func() {
		file := &mocks.File{}
		infile = store.NewFromDataSource(file)
	})

	It("done 필터 InMem로 잘 전달한다", func() {

	})
})
