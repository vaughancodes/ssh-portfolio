package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var asciiLines = []string{
	` ____              _      _  __     __                 _`,
	`|  _ \  __ _ _ __ (_) ___| | \ \   / /_ _ _   _  __ _| |__   __ _ _ __`,
	"| | | |/ _` | '_ \\| |/ _ \\ |  \\ \\ / / _` | | | |/ _` | '_ \\ / _` | '_ \\",
	`| |_| | (_| | | | | |  __/ |   \ V / (_| | |_| | (_| | | | | (_| | | | |`,
	`|____/ \__,_|_| |_|_|\___|_|    \_/ \__,_|\__,_|\__, |_| |_|\__,_|_| |_|`,
	`                                                 |___/`,
}

var bannerColors = []lipgloss.Color{
	"#F87171", "#DC2626", "#B91C1C", "#DC2626", "#F87171", "#FCA5A5",
}

// hyperlink wraps text in an OSC 8 clickable hyperlink escape sequence.
// Terminals that don't support it will just show the display text.
func hyperlink(url, text string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

var nonDigitRe = regexp.MustCompile(`[^\d+]`)

// urlFor returns a full URL for a contact value.
func urlFor(label, value string) string {
	switch label {
	case "Email":
		return "mailto:" + value
	case "Phone", "Office":
		return "tel:" + nonDigitRe.ReplaceAllString(value, "")
	case "GitHub", "LinkedIn":
		return "https://" + value
	default:
		return ""
	}
}

func renderAbout(s styles, width int) string {
	var b strings.Builder

	for i, line := range asciiLines {
		color := bannerColors[i%len(bannerColors)]
		b.WriteString(s.r.NewStyle().Foreground(color).Bold(true).Render(line))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	role := s.greenText.Render(profile.Role)
	sep := s.dimText.Render("  ◆  ")
	loc := s.secondaryText.Render(profile.Location)
	b.WriteString(role + sep + loc)
	b.WriteString("\n\n")

	bioWidth := min(width-8, 70)
	bio := s.r.NewStyle().
		Width(bioWidth).
		Foreground(textColor).
		Render(profile.Bio)
	box := s.r.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(subtle).
		Padding(1, 2).
		Render(bio)
	b.WriteString(box)
	b.WriteString("\n")

	return b.String()
}

func renderExperience(s styles, width int) string {
	var b strings.Builder
	contentWidth := min(width-4, 72)

	b.WriteString(s.sectionHeader.Render("Work Experience"))
	b.WriteString("\n\n")

	for i, exp := range experiences {
		marker := s.greenText.Render("●")
		line := s.dimText.Render("│")

		period := s.r.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#334155")).
			Bold(true).
			Padding(0, 1).
			Render(exp.Period)

		b.WriteString(marker + "  " + s.accentText.Render(exp.Title) + "\n")
		b.WriteString(line + "  " + s.secondaryText.Render(exp.Company) + "  " + period + "\n")
		b.WriteString(line + "\n")
		b.WriteString(line + "  " + s.r.NewStyle().
			Width(contentWidth-6).
			Foreground(textColor).
			Italic(true).
			Render(exp.Description))
		b.WriteString("\n")

		for _, h := range exp.Highlights {
			b.WriteString(line + "  " + s.bullet.Render("▸ "))
			b.WriteString(s.r.NewStyle().
				Width(contentWidth-8).
				Foreground(muted).
				Render(h))
			b.WriteString("\n")
		}

		if i < len(experiences)-1 {
			b.WriteString(line + "\n")
		} else {
			b.WriteString(s.dimText.Render("╵") + "\n")
		}
	}

	return b.String()
}

func renderProjects(s styles, width int) string {
	var b strings.Builder
	cardWidth := min(width-8, 68)

	b.WriteString(s.sectionHeader.Render("Projects"))
	b.WriteString("\n\n")

	for _, proj := range projects {
		name := s.accentText.Render("◈  " + proj.Name)

		desc := s.r.NewStyle().
			Width(cardWidth - 4).
			Foreground(textColor).
			Render(proj.Description)

		var tags []string
		for _, t := range proj.Tech {
			tags = append(tags, s.tag.Render(t))
		}
		tagLine := lipgloss.JoinHorizontal(lipgloss.Top, tags...)

		url := s.dimText.Render("→ ") + s.secondaryText.Render(hyperlink("https://"+proj.URL, proj.URL))

		inner := lipgloss.JoinVertical(lipgloss.Left, name, "", desc, "", tagLine, url)

		card := s.r.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Width(cardWidth).
			Padding(1, 2).
			Render(inner)

		b.WriteString(card)
		b.WriteString("\n")
	}

	return b.String()
}

func renderSkills(s styles, width int) string {
	var b strings.Builder
	contentWidth := min(width-4, 72)

	b.WriteString(s.sectionHeader.Render("Skills & Technologies"))
	b.WriteString("\n\n")

	categoryColors := []lipgloss.Color{yellow, green, pink, orange, secondary}

	for i, group := range skillGroups {
		color := categoryColors[i%len(categoryColors)]
		header := s.r.NewStyle().Foreground(color).Bold(true).Render("■ " + group.Category)
		b.WriteString(header)
		b.WriteString("\n")

		var tags []string
		for _, sk := range group.Skills {
			tags = append(tags, s.tag.Render(sk))
		}
		b.WriteString("  " + lipgloss.JoinHorizontal(lipgloss.Top, tags...))
		b.WriteString("\n")

		if i < len(skillGroups)-1 {
			b.WriteString(s.dimText.Render("  "+repeat("·", contentWidth-4)) + "\n")
		}
	}

	return b.String()
}

func renderEducation(s styles, width int) string {
	var b strings.Builder
	contentWidth := min(width-4, 72)

	b.WriteString(s.sectionHeader.Render("Education & Certifications"))
	b.WriteString("\n\n")

	for i, edu := range education {
		b.WriteString(s.highlightText.Render("🎓 ") + s.accentText.Render(edu.Degree))
		b.WriteString("\n")
		b.WriteString("   " + s.secondaryText.Render(edu.Institution))
		b.WriteString("  " + s.dimText.Render("("+edu.Period+")"))
		b.WriteString("\n")
		b.WriteString("   " + s.r.NewStyle().
			Width(contentWidth-6).
			Foreground(muted).
			Render(edu.Details))
		b.WriteString("\n")

		if i < len(education)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func renderContact(s styles, width int) string {
	var b strings.Builder

	b.WriteString(s.sectionHeader.Render("Get In Touch"))
	b.WriteString("\n\n")

	intro := s.r.NewStyle().
		Width(min(width-4, 72)).
		Foreground(textColor).
		Render("I'm always interested in hearing about new opportunities, collaborations, or just connecting with fellow engineers.")
	b.WriteString(intro)
	b.WriteString("\n\n")

	iconMap := map[string]string{
		"Email": "✉", "Phone": "☎", "Office": "☏",
		"Portfolio": "⌂", "GitHub": "◆", "LinkedIn": "∞", "Location": "◉",
	}

	labelColors := []lipgloss.Color{pink, green, green, secondary, muted, accent, orange}

	for i, c := range contacts {
		icon := iconMap[c.Label]
		if icon == "" {
			icon = "→"
		}
		color := labelColors[i%len(labelColors)]
		styledIcon := s.r.NewStyle().Foreground(color).Bold(true).Render(icon)
		styledLabel := s.r.NewStyle().Foreground(color).Bold(true).Width(12).Render(c.Label)

		if c.Label == "Portfolio" {
			valuePart := s.r.NewStyle().Foreground(textColor).Bold(true).Render(
				hyperlink("https://"+c.Value, c.Value),
			)
			httpsLink := s.r.NewStyle().Foreground(textColor).Bold(true).Render(
				hyperlink("https://"+c.Value, "HTTPS"),
			)
			sshLink := s.r.NewStyle().Foreground(textColor).Bold(true).Render(
				hyperlink("ssh://"+c.Value, "SSH"),
			)
			suffix := s.dimText.Render(" (") + httpsLink + s.dimText.Render(" or ") + sshLink + s.dimText.Render(")") +
				s.dimText.Render("\t// you're already here!")
			b.WriteString(fmt.Sprintf("  %s  %s  %s%s\n", styledIcon, styledLabel, valuePart, suffix))
		} else {
			valueText := c.Value
			if u := urlFor(c.Label, c.Value); u != "" {
				valueText = hyperlink(u, c.Value)
			}
			styledValue := s.r.NewStyle().Foreground(textColor).Bold(true).Render(valueText)
			b.WriteString(fmt.Sprintf("  %s  %s  %s\n", styledIcon, styledLabel, styledValue))
		}
	}

	b.WriteString("\n")
	b.WriteString(s.dimText.Render("  Thanks for stopping by! ") + s.highlightText.Render("👋"))
	b.WriteString("\n")

	return b.String()
}
