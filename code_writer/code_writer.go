package codewriter

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iris-net/hackvm2asm/code_writer/arithmetic"
	"github.com/iris-net/hackvm2asm/code_writer/segment"
	"github.com/iris-net/hackvm2asm/parser"
	"github.com/iris-net/hackvm2asm/parser/command"
)

type CodeWriter struct {
	fileName     string
	assemblies   []string
	labelCounter int
}

func NewCodeWriter() CodeWriter {
	return CodeWriter{
		assemblies: make([]string, 0, 100),
	}
}

func (c *CodeWriter) addAssemblies(assemblies []string) {
	c.assemblies = append(c.assemblies, assemblies...)
	c.assemblies = append(c.assemblies, "")
}

func (c CodeWriter) GetAssemblies() []string {
	return c.assemblies
}

// Execute analyzes vmFile and translates to assemby codes
func (c *CodeWriter) Execute(vmFile, outFile string) error {
	c.assemblies = make([]string, 0, 100)
	c.fileName = strings.TrimSuffix(filepath.Base(vmFile), filepath.Ext(vmFile))

	p, err := parser.NewParser(vmFile)
	if err != nil {
		return err
	}

	for p.HasMoreCommands() {
		p.Advance()
		t, err := p.GetCommandType()
		if err != nil {
			return err
		}

		switch t {
		case command.Arithmetic:
			ari := arithmetic.NewCommand(p.GetCommand())
			err = c.TranslateArithmetic(ari)
			if err != nil {
				return err
			}
		case command.Push:
			arg1, err := p.GetArg1()
			if err != nil {
				return err
			}

			arg2, err := p.GetArg2()
			if err != nil {
				return err
			}

			seg := segment.NewSegment(arg1)
			index, err := strconv.Atoi(arg2)
			if err != nil {
				return err
			}

			err = c.TranslatePush(seg, index)
			if err != nil {
				return err
			}
		case command.Pop:
			arg1, err := p.GetArg1()
			if err != nil {
				return err
			}

			arg2, err := p.GetArg2()
			if err != nil {
				return err
			}

			seg := segment.NewSegment(arg1)
			index, err := strconv.Atoi(arg2)
			if err != nil {
				return err
			}

			err = c.TranslatePop(seg, index)
			if err != nil {
				return err
			}
		}
	}

	f, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, a := range c.GetAssemblies() {
		f.WriteString(a + "\n")
	}

	return nil
}

// translateArithmetic translates the given arithmetic command
func (c *CodeWriter) TranslateArithmetic(cmd arithmetic.Command) error {
	buf := make([]string, 0, 10)
	buf = append(buf, fmt.Sprintf("// ===== arithmetic command %s =====", cmd.String()))

	if cmd.DoNeedTwoArgs() {
		// decrement SP
		buf = c.decrementSP(buf)
		// get value from SP
		buf = append(buf, "// D=Memory[SP]")
		buf = append(buf, "@SP")
		buf = append(buf, "A=M")
		buf = append(buf, "D=M")
	}

	// decrement SP
	buf = c.decrementSP(buf)
	// set pointer from SP
	buf = append(buf, "// A=Memory[SP]")
	buf = append(buf, "A=M")

	buf = append(buf, fmt.Sprintf("// %s", cmd.String()))
	switch cmd {
	case arithmetic.Add:
		buf = append(buf, "M=M+D")
	case arithmetic.Sub:
		buf = append(buf, "M=M-D")
	case arithmetic.Neg:
		buf = append(buf, "D=0")
		buf = append(buf, "M=D-M")
	case arithmetic.EQ, arithmetic.GT, arithmetic.LT:
		buf = c.arithmeticLogic(cmd, buf)
	case arithmetic.And:
		buf = append(buf, "M=M&D")
	case arithmetic.Or:
		buf = append(buf, "M=M|D")
	case arithmetic.Not:
		buf = append(buf, "M=!M")
	}

	// increment SP
	buf = c.incrementSP(buf)

	c.addAssemblies(buf)

	return nil
}

// TranslatePush translates the given push command
func (c *CodeWriter) TranslatePush(seg segment.Segment, index int) error {
	buf := make([]string, 0, 10)
	buf = append(buf, fmt.Sprintf("// ===== push %s %d =====", seg.String(), index))
	switch seg {
	case segment.Constant:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
	case segment.Pointer:
		if index == 0 {
			buf = append(buf, "@THIS")
		} else if index == 1 {
			buf = append(buf, "@THAT")
		} else {
			return fmt.Errorf("invalid index of pointer. index=%d", index)
		}
		buf = append(buf, "D=M")
	case segment.Temp:
		buf = append(buf, fmt.Sprintf("@R%d", 5+index))
		buf = append(buf, "D=M")
	case segment.This:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@THIS")
		buf = append(buf, "A=M+D")
		buf = append(buf, "D=M")
	case segment.That:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@THAT")
		buf = append(buf, "A=M+D")
		buf = append(buf, "D=M")
	case segment.Local:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@LCL")
		buf = append(buf, "A=M+D")
		buf = append(buf, "D=M")
	case segment.Argument:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@ARG")
		buf = append(buf, "A=M+D")
		buf = append(buf, "D=M")
	case segment.Static:
		buf = append(buf, fmt.Sprintf("@%s.%d", c.fileName, index))
		buf = append(buf, "D=M")
	}

	// store D into memory indicated by SP
	buf = append(buf, "@SP")
	buf = append(buf, "A=M")
	buf = append(buf, "M=D")

	// increment SP
	buf = c.incrementSP(buf)

	c.addAssemblies(buf)

	return nil
}

// TranslatePop translates the given pop command
func (c *CodeWriter) TranslatePop(seg segment.Segment, index int) error {
	buf := make([]string, 0, 10)
	buf = append(buf, fmt.Sprintf("// ===== pop %s %d ======", seg.String(), index))

	switch seg {
	case segment.Pointer:
		if index == 0 {
			buf = append(buf, "@THIS")
		} else if index == 1 {
			buf = append(buf, "@THAT")
		} else {
			return fmt.Errorf("invalid index of pointer. index=%d", index)
		}
		buf = append(buf, "D=A")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.Temp:
		buf = append(buf, fmt.Sprintf("@R%d", 5+index))
		buf = append(buf, "D=A")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.This:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@THIS")
		buf = append(buf, "D=M+D")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.That:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@THAT")
		buf = append(buf, "D=M+D")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.Local:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@LCL")
		buf = append(buf, "D=M+D")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.Argument:
		buf = append(buf, fmt.Sprintf("@%d", index))
		buf = append(buf, "D=A")
		buf = append(buf, "@ARG")
		buf = append(buf, "D=M+D")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	case segment.Static:
		buf = append(buf, fmt.Sprintf("@%s.%d", c.fileName, index))
		buf = append(buf, "D=A")
		buf = append(buf, "@R13")
		buf = append(buf, "M=D")
	}

	// decrement SP
	buf = c.decrementSP(buf)

	// store memory indicated by SP into D
	buf = append(buf, "A=M")
	buf = append(buf, "D=M")

	buf = append(buf, "@R13")
	buf = append(buf, "A=M")
	buf = append(buf, "M=D")

	c.addAssemblies(buf)

	return nil
}

func (c CodeWriter) incrementSP(buf []string) []string {
	buf = append(buf, "// SP++")
	buf = append(buf, "@SP")
	buf = append(buf, "M=M+1")
	buf = append(buf, "")
	return buf
}

func (c CodeWriter) decrementSP(buf []string) []string {
	buf = append(buf, "// SP--")
	buf = append(buf, "@SP")
	buf = append(buf, "M=M-1")
	buf = append(buf, "")
	return buf
}

func (c *CodeWriter) arithmeticLogic(cmd arithmetic.Command, buf []string) []string {
	labelTrue := fmt.Sprintf("%sTRUE%d", cmd.String(), c.labelCounter)
	labelEnd := fmt.Sprintf("%sEND%d", cmd.String(), c.labelCounter)

	buf = append(buf, "D=M-D")
	buf = append(buf, fmt.Sprintf("@%s", labelTrue))

	switch cmd {
	case arithmetic.EQ:
		buf = append(buf, "D;JEQ")
	case arithmetic.GT:
		buf = append(buf, "D;JGT")
	case arithmetic.LT:
		buf = append(buf, "D;JLT")
	}

	buf = append(buf, "@SP")
	buf = append(buf, "A=M")
	buf = append(buf, "M=0")
	buf = append(buf, fmt.Sprintf("@%s", labelEnd))
	buf = append(buf, "0;JMP")

	buf = append(buf, fmt.Sprintf("(%s)", labelTrue))
	buf = append(buf, "@SP")
	buf = append(buf, "A=M")
	buf = append(buf, "M=-1")
	buf = append(buf, fmt.Sprintf("(%s)", labelEnd))
	c.labelCounter += 1

	return buf
}
