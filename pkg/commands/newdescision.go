package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func getTitleDesc() (string, string) {
	title := getTitle()
	desc := getDesc()
	return title, desc
}

func getTitle() string {
	var title string
	fmt.Print("Title Of Log:\n\n")
	fmt.Print("> ")
	scanner := bufio.NewReader(os.Stdin)
	if input, err := scanner.ReadString('\n'); err == nil {
		title = strings.TrimSpace(input)
	} else {
		log.Fatal("Failed to read user input for title:", err)
	}
	return title
}
func getDesc() string {
	fmt.Println("")
	m := initialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	return m.getTextAreaValue()
}

type errMsg error


type model struct {
	textarea textarea.Model
	lines    []string
	err      error
}

func initialModel() *model {
	ti := textarea.New()
	ti.Placeholder = "Your reasoning for your description....."
	ti.Focus()
	return &model{
		textarea: ti,
		err:      nil,
	}
}

func (m *model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *model) getTextAreaValue() string {
	return m.textarea.Value()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyEnter:
			cmd = m.textarea.Focus()
			m.lines = append(m.lines, m.getTextAreaValue())
			cmds = append(cmds, cmd)
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	return fmt.Sprintf(
		"Log Reasoning.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}
