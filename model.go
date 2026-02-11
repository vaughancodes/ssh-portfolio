package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = model{}

var tabNames = []string{"About", "Experience", "Projects", "Skills", "Education", "Contact"}

type model struct {
	activeTab int
	hoverTab  int
	tabs      []string
	viewport  viewport.Model
	width     int
	height    int
	ready     bool
	styles    styles
}

func newModel(width, height int, r *lipgloss.Renderer) model {
	return model{
		tabs:     tabNames,
		hoverTab: -1,
		width:    width,
		height:   height,
		styles:   newStyles(r),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 4
		footerHeight := 2
		contentHeight := m.height - headerHeight - footerHeight
		if contentHeight < 1 {
			contentHeight = 1
		}

		if !m.ready {
			m.viewport = viewport.New(m.width, contentHeight)
			m.viewport.SetContent(m.currentTabContent())
			m.ready = true
		} else {
			m.viewport.Width = m.width
			m.viewport.Height = contentHeight
			m.viewport.SetContent(m.currentTabContent())
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab", "right", "l":
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
			m.viewport.SetContent(m.currentTabContent())
			m.viewport.GotoTop()
			return m, nil

		case "shift+tab", "left", "h":
			m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
			m.viewport.SetContent(m.currentTabContent())
			m.viewport.GotoTop()
			return m, nil
		}

	case tea.MouseMsg:
		if msg.Y <= 1 {
			m.hoverTab = m.tabHitTest(msg.X)
		} else {
			m.hoverTab = -1
		}

		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			if m.hoverTab >= 0 && m.hoverTab != m.activeTab {
				m.activeTab = m.hoverTab
				m.viewport.SetContent(m.currentTabContent())
				m.viewport.GotoTop()
				return m, nil
			}
		}
	}

	if m.ready {
		m.viewport, cmd = m.viewport.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	tabBar := m.styles.renderTabBar(m.tabs, m.activeTab, m.hoverTab, m.width)
	content := m.styles.contentBox.Render(m.viewport.View())
	footer := m.styles.renderFooter(m.width)

	return lipgloss.JoinVertical(lipgloss.Left, tabBar, content, footer)
}

func (m model) tabHitTest(x int) int {
	for i, name := range m.tabs {
		tabWidth := len(name) + 4 // padding(0,2) adds 4
		if x < tabWidth {
			return i
		}
		x -= tabWidth
	}
	return -1
}

func (m model) currentTabContent() string {
	s := m.styles
	w := m.width
	switch m.activeTab {
	case 0:
		return renderAbout(s, w)
	case 1:
		return renderExperience(s, w)
	case 2:
		return renderProjects(s, w)
	case 3:
		return renderSkills(s, w)
	case 4:
		return renderEducation(s, w)
	case 5:
		return renderContact(s, w)
	default:
		return ""
	}
}
