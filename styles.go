package main

import "github.com/charmbracelet/lipgloss"

// Color constants
var (
	accent    = lipgloss.Color("#DC2626")
	accentDim = lipgloss.Color("#F87171")
	secondary = lipgloss.Color("#06B6D4")
	green     = lipgloss.Color("#10B981")
	pink      = lipgloss.Color("#EC4899")
	orange    = lipgloss.Color("#F97316")
	yellow    = lipgloss.Color("#FBBF24")
	textColor = lipgloss.Color("#E2E8F0")
	muted     = lipgloss.Color("#94A3B8")
	dim       = lipgloss.Color("#64748B")
	subtle    = lipgloss.Color("#334155")
)

// styles holds all lipgloss styles, created from a session-aware renderer.
type styles struct {
	title         lipgloss.Style
	subtitle      lipgloss.Style
	sectionHeader lipgloss.Style
	accentText    lipgloss.Style
	secondaryText lipgloss.Style
	mutedText     lipgloss.Style
	dimText       lipgloss.Style
	highlightText lipgloss.Style
	greenText     lipgloss.Style
	pinkText      lipgloss.Style
	orangeText    lipgloss.Style
	tag           lipgloss.Style
	activeTab     lipgloss.Style
	inactiveTab   lipgloss.Style
	footerKey     lipgloss.Style
	footerDesc    lipgloss.Style
	bullet        lipgloss.Style
	contentBox    lipgloss.Style
	divider       lipgloss.Style
	base          lipgloss.Style // unstyled, for building ad-hoc styles
	r             *lipgloss.Renderer
}

func newStyles(r *lipgloss.Renderer) styles {
	return styles{
		r: r,
		base: r.NewStyle(),
		title: r.NewStyle().
			Bold(true).
			Foreground(accent).
			MarginBottom(1),
		subtitle: r.NewStyle().
			Foreground(muted).
			Italic(true),
		sectionHeader: r.NewStyle().
			Bold(true).
			Foreground(accent).
			Underline(true).
			MarginBottom(1),
		accentText: r.NewStyle().
			Foreground(accent).
			Bold(true),
		secondaryText: r.NewStyle().
			Foreground(secondary).
			Bold(true),
		mutedText: r.NewStyle().
			Foreground(muted),
		dimText: r.NewStyle().
			Foreground(dim),
		highlightText: r.NewStyle().
			Foreground(yellow).
			Bold(true),
		greenText: r.NewStyle().
			Foreground(green).
			Bold(true),
		pinkText: r.NewStyle().
			Foreground(pink).
			Bold(true),
		orangeText: r.NewStyle().
			Foreground(orange).
			Bold(true),
		tag: r.NewStyle().
			Foreground(secondary).
			Background(lipgloss.Color("#0E3A4A")).
			Bold(true).
			Padding(0, 1),
		activeTab: r.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(accent).
			Padding(0, 2),
		inactiveTab: r.NewStyle().
			Foreground(muted).
			Padding(0, 2),
		footerKey: r.NewStyle().
			Foreground(accentDim).
			Bold(true),
		footerDesc: r.NewStyle().
			Foreground(dim),
		bullet: r.NewStyle().
			Foreground(green).
			Bold(true),
		contentBox: r.NewStyle().
			Padding(1, 2),
		divider: r.NewStyle().
			Foreground(subtle),
	}
}

func (s styles) renderTabBar(tabs []string, active int, width int) string {
	var rendered []string
	for i, t := range tabs {
		if i == active {
			rendered = append(rendered, s.activeTab.Render(t))
		} else {
			rendered = append(rendered, s.inactiveTab.Render(t))
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, rendered...)
	bar := s.r.NewStyle().
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(accent).
		Width(width).
		Render(row)
	return bar
}

func (s styles) renderFooter(width int) string {
	parts := []string{
		s.footerKey.Render("←/→") + s.footerDesc.Render(" navigate"),
		s.footerKey.Render("↑/↓") + s.footerDesc.Render(" scroll"),
		s.footerKey.Render("q") + s.footerDesc.Render(" quit"),
	}
	sep := s.dimText.Render("  •  ")
	line := parts[0] + sep + parts[1] + sep + parts[2]
	return s.r.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		MarginTop(1).
		Render(line)
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
