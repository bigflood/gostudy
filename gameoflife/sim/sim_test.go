package sim_test

import (
	. "github.com/onsi/ginkgo"

	. "github.com/bigflood/gostudy/gameoflife/sim"
)

var _ = Describe("Sim", func() {

	const (
		width  = 16
		height = 16
	)

	It("frame의 입력에 문제가 없는지 테스트", func() {
		s := New(width, height)

		for index := 0; index < 100; index++ {
			s.WaitForFrame()
			//log.Println("frame:", index)
		}
	})
})
