package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/iris-net/hackvm2asm/parser/command"
)

type Parser struct {
	commands []string
	cursor   int
}

// load input and get ready to parse it
func NewParser(path string) (Parser, error) {
	p := Parser{cursor: -1}
	err := p.load(path)
	if err != nil {
		return Parser{}, err
	}
	return p, nil
}

func (p Parser) GetCursor() int {
	return p.cursor
}

func (p Parser) GetCommands() []string {
	return p.commands
}

func (p Parser) getCurrentCommand() string {
	return p.commands[p.cursor]
}

// load input file
func (p *Parser) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := make([]string, 0, 100)
	for scanner.Scan() {
		cmd := p.removeNotCommand(scanner.Text())
		if len(cmd) > 0 {
			commands = append(commands, cmd)
		}
	}

	p.commands = commands

	return nil
}

// removeNotCommand removes not command descriptions
func (p Parser) removeNotCommand(text string) (ret string) {
	r := regexp.MustCompile(`//(\w|\W)*`)
	ret = r.ReplaceAllString(text, "")

	r = regexp.MustCompile(`\t`)
	ret = r.ReplaceAllString(ret, "")

	r = regexp.MustCompile(`\s+$`)
	ret = r.ReplaceAllString(ret, "")

	return ret
}

// HasMoreCommands returns whether there are more commands in the input
func (p Parser) HasMoreCommands() bool {
	return p.cursor+1 < len(p.commands)
}

// Advance reads the next command from the input and makes it the command.
func (p *Parser) Advance() {
	p.cursor += 1
}

// GetCommandType gets the command type of current command.
func (p Parser) GetCommandType() (command.Type, error) {
	cmdLine := p.getCurrentCommand()
	cmd := strings.Split(cmdLine, " ")[0]
	cmdType := command.NewType(cmd)

	if cmdType == command.Unknown {
		return cmdType, fmt.Errorf("unknown command type. %s", cmdLine)
	}

	return cmdType, nil
}

// GetCommand gets the command of current command.
func (p Parser) GetCommand() string {
	cmdLine := p.getCurrentCommand()
	return strings.Split(cmdLine, " ")[0]
}

// GetArg1 gets the first argument of current code.
func (p Parser) GetArg1() (string, error) {
	cmdLine := p.getCurrentCommand()
	cmdType, err := p.GetCommandType()
	if err != nil {
		return "", err
	}

	if cmdType == command.Return {
		return "", fmt.Errorf("the return command can't have any argments. %s", cmdLine)
	}

	spt := strings.Split(cmdLine, " ")

	if cmdType == command.Arithmetic {
		return spt[0], nil
	}

	if len(spt) < 2 {
		return "", fmt.Errorf("failed to load the first argument. %s", cmdLine)
	}

	return spt[1], nil
}

// GetArg2 gets the second argument of current code.
func (p Parser) GetArg2() (string, error) {
	cmdLine := p.getCurrentCommand()
	cmdType, err := p.GetCommandType()
	if err != nil {
		return "", err
	}

	switch cmdType {
	case command.Push, command.Pop, command.Function, command.Call:
	default:
		return "", fmt.Errorf("the command can't have second argment. %s", cmdLine)
	}

	spt := strings.Split(cmdLine, " ")

	if len(spt) < 3 {
		return "", fmt.Errorf("failed to load the second argument. %s", cmdLine)
	}

	return spt[2], nil
}
