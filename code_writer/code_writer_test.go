package codewriter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	codewriter "github.com/iris-net/hackvm2asm/code_writer"
)

const TestBasicTestPath = "../test/BasicTest.vm"
const TestPointerTestPath = "../test/PointerTest.vm"
const TestSimpleAddPath = "../test/SimpleAdd.vm"
const TestStackTestPath = "../test/StackTest.vm"
const TestStaticTestPath = "../test/StaticTest.vm"

var _ = Describe("CodeWriter", func() {
	Context("Execute()", func() {
		When("load BasicTest.vm", func() {
			It("successfully translate", func() {
				cw := codewriter.NewCodeWriter()
				err := cw.Execute(TestBasicTestPath, TestBasicTestPath[0:len(TestBasicTestPath)-2]+"asm")
				Expect(err).To(BeNil())

			})
		})

		When("load PointerTest.vm", func() {
			It("successfully translate", func() {
				cw := codewriter.NewCodeWriter()
				err := cw.Execute(TestPointerTestPath, TestPointerTestPath[0:len(TestPointerTestPath)-2]+"asm")
				Expect(err).To(BeNil())

			})
		})

		When("load SimpleAdd.vm", func() {
			It("successfully translate", func() {
				cw := codewriter.NewCodeWriter()
				err := cw.Execute(TestSimpleAddPath, TestSimpleAddPath[0:len(TestSimpleAddPath)-2]+"asm")
				Expect(err).To(BeNil())

			})
		})

		When("load StackTest.vm", func() {
			It("successfully translate", func() {
				cw := codewriter.NewCodeWriter()
				err := cw.Execute(TestStackTestPath, TestStackTestPath[0:len(TestStackTestPath)-2]+"asm")
				Expect(err).To(BeNil())

			})
		})

		When("load StaticTest.vm", func() {
			It("successfully translate", func() {
				cw := codewriter.NewCodeWriter()
				err := cw.Execute(TestStaticTestPath, TestStaticTestPath[0:len(TestStaticTestPath)-2]+"asm")
				Expect(err).To(BeNil())

			})
		})
	})
})
