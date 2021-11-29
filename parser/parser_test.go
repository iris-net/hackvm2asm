package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/iris-net/hackvm2asm/parser"
	"github.com/iris-net/hackvm2asm/parser/command"
)

const TestBasicTestPath = "../test/BasicTest.vm"
const TestPointerTestPath = "../test/PointerTest.vm"
const TestSimpleAddPath = "../test/SimpleAdd.vm"
const TestStackTestPath = "../test/StackTest.vm"
const TestStaticTestPath = "../test/StaticTest.vm"

var _ = Describe("Parser", func() {
	Context("NewParser()", func() {
		When("load BasicTest.vm", func() {
			It("successfully load", func() {
				p, err := parser.NewParser(TestBasicTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"push constant 10", "pop local 0", "push constant 21", "push constant 22", "pop argument 2", "pop argument 1", "push constant 36", "pop this 6", "push constant 42", "push constant 45", "pop that 5", "pop that 2", "push constant 510", "pop temp 6", "push local 0", "push that 5", "add", "push argument 1", "sub", "push this 6", "push this 6", "add", "sub", "push temp 6", "add",
				}

				Expect(p.GetCommands()).To(Equal(exp))
			})
		})

		When("load PointerTest.vm", func() {
			It("successfully load", func() {
				p, err := parser.NewParser(TestPointerTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"push constant 3030", "pop pointer 0", "push constant 3040", "pop pointer 1", "push constant 32", "pop this 2", "push constant 46", "pop that 6", "push pointer 0", "push pointer 1", "add", "push this 2", "sub", "push that 6", "add",
				}

				Expect(p.GetCommands()).To(Equal(exp))
			})
		})

		When("load SimpleAdd.vm", func() {
			It("successfully load", func() {
				p, err := parser.NewParser(TestSimpleAddPath)
				Expect(err).To(BeNil())

				exp := []string{
					"push constant 7", "push constant 8", "add",
				}

				Expect(p.GetCommands()).To(Equal(exp))
			})
		})

		When("load StackTest.vm", func() {
			It("successfully load", func() {
				p, err := parser.NewParser(TestStackTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"push constant 17", "push constant 17", "eq", "push constant 17", "push constant 16", "eq", "push constant 16", "push constant 17", "eq", "push constant 892", "push constant 891", "lt", "push constant 891", "push constant 892", "lt", "push constant 891", "push constant 891", "lt", "push constant 32767", "push constant 32766", "gt", "push constant 32766", "push constant 32767", "gt", "push constant 32766", "push constant 32766", "gt", "push constant 57", "push constant 31", "push constant 53", "add", "push constant 112", "sub", "neg", "and", "push constant 82", "or", "not",
				}

				Expect(p.GetCommands()).To(Equal(exp))
			})
		})

		When("load StaticTest.vm", func() {
			It("successfully load", func() {
				p, err := parser.NewParser(TestStaticTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"push constant 111", "push constant 333", "push constant 888", "pop static 8", "pop static 3", "pop static 1", "push static 3", "push static 1", "sub", "push static 8", "add",
				}

				Expect(p.GetCommands()).To(Equal(exp))
			})
		})
	})

	Context("GetCommandType()", func() {
		When("load BasicTest.vm", func() {
			It("successfully get coomand type", func() {
				p, err := parser.NewParser(TestBasicTestPath)
				Expect(err).To(BeNil())

				exp := []command.Type{
					command.Push, command.Pop, command.Push, command.Push, command.Pop, command.Pop, command.Push, command.Pop, command.Push, command.Push, command.Pop, command.Pop, command.Push, command.Pop, command.Push, command.Push, command.Arithmetic, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Arithmetic, command.Push, command.Arithmetic,
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetCommandType()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load PointerTest.vm", func() {
			It("successfully get coomand type", func() {
				p, err := parser.NewParser(TestPointerTestPath)
				Expect(err).To(BeNil())

				exp := []command.Type{
					command.Push, command.Pop, command.Push, command.Pop, command.Push, command.Pop, command.Push, command.Pop, command.Push, command.Push, command.Arithmetic, command.Push, command.Arithmetic, command.Push, command.Arithmetic,
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetCommandType()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load SimpleAdd.vm", func() {
			It("successfully get coomand type", func() {
				p, err := parser.NewParser(TestSimpleAddPath)
				Expect(err).To(BeNil())

				exp := []command.Type{
					command.Push, command.Push, command.Arithmetic,
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetCommandType()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load StackTest.vm", func() {
			It("successfully get coomand type", func() {
				p, err := parser.NewParser(TestStackTestPath)
				Expect(err).To(BeNil())

				exp := []command.Type{
					command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Arithmetic, command.Push, command.Push, command.Push, command.Arithmetic, command.Push, command.Arithmetic, command.Arithmetic, command.Arithmetic, command.Push, command.Arithmetic, command.Arithmetic,
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetCommandType()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load StaticTest.vm", func() {
			It("successfully get coomand type", func() {
				p, err := parser.NewParser(TestStaticTestPath)
				Expect(err).To(BeNil())

				exp := []command.Type{
					command.Push, command.Push, command.Push, command.Pop, command.Pop, command.Pop, command.Push, command.Push, command.Arithmetic, command.Push, command.Arithmetic,
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetCommandType()).To(Equal(exp[i]))

					i++
				}
			})
		})
	})

	Context("GetArg1()", func() {
		When("load BasicTest.vm", func() {
			It("successfully get first argument", func() {
				p, err := parser.NewParser(TestBasicTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"constant", "local", "constant", "constant", "argument", "argument", "constant", "this", "constant", "constant", "that", "that", "constant", "temp", "local", "that", "add", "argument", "sub", "this", "this", "add", "sub", "temp", "add",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetArg1()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load PointerTest.vm", func() {
			It("successfully get first argument", func() {
				p, err := parser.NewParser(TestPointerTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"constant", "pointer", "constant", "pointer", "constant", "this", "constant", "that", "pointer", "pointer", "add", "this", "sub", "that", "add",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetArg1()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load SimpleAdd.vm", func() {
			It("successfully get first argument", func() {
				p, err := parser.NewParser(TestSimpleAddPath)
				Expect(err).To(BeNil())

				exp := []string{
					"constant", "constant", "add",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetArg1()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load StackTest.vm", func() {
			It("successfully get first argument", func() {
				p, err := parser.NewParser(TestStackTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"constant", "constant", "eq", "constant", "constant", "eq", "constant", "constant", "eq", "constant", "constant", "lt", "constant", "constant", "lt", "constant", "constant", "lt", "constant", "constant", "gt", "constant", "constant", "gt", "constant", "constant", "gt", "constant", "constant", "constant", "add", "constant", "sub", "neg", "and", "constant", "or", "not",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetArg1()).To(Equal(exp[i]))

					i++
				}
			})
		})

		When("load StaticTest.vm", func() {
			It("successfully get first argument", func() {
				p, err := parser.NewParser(TestStaticTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"constant", "constant", "constant", "static", "static", "static", "static", "static", "sub", "static", "add",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					Expect(p.GetArg1()).To(Equal(exp[i]))

					i++
				}
			})
		})
	})

	Context("GetArg2()", func() {
		When("load BasicTest.vm", func() {
			It("successfully get second argument", func() {
				p, err := parser.NewParser(TestBasicTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"10", "0", "21", "22", "2", "1", "36", "6", "42", "45", "5", "2", "510", "6", "0", "5", "", "1", "", "6", "6", "", "", "6", "",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					a, err := p.GetArg2()

					Expect(a).To(Equal(exp[i]))
					if len(a) == 0 {
						Expect(err == nil).To(BeFalse())
					} else {
						Expect(err == nil).To(BeTrue())
					}

					i++
				}
			})
		})

		When("load PointerTest.vm", func() {
			It("successfully get second argument", func() {
				p, err := parser.NewParser(TestPointerTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"3030", "0", "3040", "1", "32", "2", "46", "6", "0", "1", "", "2", "", "6", "",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					a, err := p.GetArg2()

					Expect(a).To(Equal(exp[i]))
					if len(a) == 0 {
						Expect(err == nil).To(BeFalse())
					} else {
						Expect(err == nil).To(BeTrue())
					}

					i++
				}
			})
		})

		When("load SimpleAdd.vm", func() {
			It("successfully get second argument", func() {
				p, err := parser.NewParser(TestSimpleAddPath)
				Expect(err).To(BeNil())

				exp := []string{
					"7", "8", "",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					a, err := p.GetArg2()

					Expect(a).To(Equal(exp[i]))
					if len(a) == 0 {
						Expect(err == nil).To(BeFalse())
					} else {
						Expect(err == nil).To(BeTrue())
					}

					i++
				}
			})
		})

		When("load StackTest.vm", func() {
			It("successfully get second argument", func() {
				p, err := parser.NewParser(TestStackTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"17", "17", "", "17", "16", "", "16", "17", "", "892", "891", "", "891", "892", "", "891", "891", "", "32767", "32766", "", "32766", "32767", "", "32766", "32766", "", "57", "31", "53", "", "112", "", "", "", "82", "", "",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					a, err := p.GetArg2()

					Expect(a).To(Equal(exp[i]))
					if len(a) == 0 {
						Expect(err == nil).To(BeFalse())
					} else {
						Expect(err == nil).To(BeTrue())
					}

					i++
				}
			})
		})

		When("load StaticTest.vm", func() {
			It("successfully get second argument", func() {
				p, err := parser.NewParser(TestStaticTestPath)
				Expect(err).To(BeNil())

				exp := []string{
					"111", "333", "888", "8", "3", "1", "3", "1", "", "8", "",
				}

				i := 0
				for p.HasMoreCommands() {
					p.Advance()

					a, err := p.GetArg2()

					Expect(a).To(Equal(exp[i]))
					if len(a) == 0 {
						Expect(err == nil).To(BeFalse())
					} else {
						Expect(err == nil).To(BeTrue())
					}

					i++
				}
			})
		})
	})
})
