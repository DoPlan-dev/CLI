package dpr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// QuestionnaireModel represents the DPR questionnaire TUI
type QuestionnaireModel struct {
	width         int
	height        int
	currentStep   int
	totalSteps    int
	answers       map[string]interface{}
	projectRoot   string
	err           error
	
	// Input components
	textInput    textinput.Model
	selectList   list.Model
	scaleValue   int
	multiSelect  map[int]bool
	currentInput string // "text", "select", "scale", "multi-select"
}

// Question represents a single question in the questionnaire
type Question struct {
	ID          string
	Category    string
	Text        string
	Type        string // "text", "select", "multi-select", "scale"
	Options     []string
	Required    bool
	Placeholder string
}

// RunQuestionnaire launches the interactive DPR questionnaire
func RunQuestionnaire(projectRoot string) (*DPRData, error) {
	ti := textinput.New()
	ti.Placeholder = "Type your answer..."
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 50

	m := &QuestionnaireModel{
		currentStep: 0,
		totalSteps:  getTotalQuestions(),
		answers:     make(map[string]interface{}),
		projectRoot: projectRoot,
		textInput:   ti,
		scaleValue:  3, // Default to middle
		multiSelect: make(map[int]bool),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	model := finalModel.(*QuestionnaireModel)
	return model.toDPRData(), nil
}

func (m *QuestionnaireModel) Init() tea.Cmd {
	return m.setupCurrentQuestion()
}

func (m *QuestionnaireModel) setupCurrentQuestion() tea.Cmd {
	question := getQuestion(m.currentStep)
	if question == nil {
		return nil
	}

	// Check if already answered
	if val, ok := m.answers[question.ID]; ok {
		switch question.Type {
		case "text":
			m.textInput.SetValue(fmt.Sprintf("%v", val))
		case "select":
			// Find index in options
			for i, opt := range question.Options {
				if opt == fmt.Sprintf("%v", val) {
					m.selectList.Select(i)
					break
				}
			}
		case "scale":
			if num, ok := val.(int); ok {
				m.scaleValue = num
			}
		}
	}

	switch question.Type {
	case "text":
		m.textInput.Focus()
		m.textInput.Placeholder = question.Placeholder
		m.currentInput = "text"
		return textinput.Blink
	case "select":
		items := make([]list.Item, len(question.Options))
		for i, opt := range question.Options {
			items[i] = selectItem{text: opt, index: i}
		}
		m.selectList = list.New(items, list.NewDefaultDelegate(), 0, 0)
		m.selectList.Title = ""
		m.selectList.SetShowStatusBar(false)
		m.selectList.SetFilteringEnabled(false)
		m.currentInput = "select"
		return nil
	case "scale":
		m.currentInput = "scale"
		return nil
	case "multi-select":
		items := make([]list.Item, len(question.Options))
		for i, opt := range question.Options {
			selected := m.multiSelect[i]
			items[i] = multiSelectItem{text: opt, index: i, selected: selected}
		}
		m.selectList = list.New(items, list.NewDefaultDelegate(), 0, 0)
		m.selectList.Title = ""
		m.selectList.SetShowStatusBar(false)
		m.selectList.SetFilteringEnabled(false)
		m.currentInput = "multi-select"
		return nil
	}

	return nil
}

func (m *QuestionnaireModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.currentInput == "select" || m.currentInput == "multi-select" {
			m.selectList.SetWidth(msg.Width - 4)
			m.selectList.SetHeight(msg.Height - 15)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m.handleEnter()
		case "esc", "b":
			return m.handleBack()
		case "1", "2", "3", "4", "5":
			if m.currentInput == "scale" {
				val, _ := strconv.Atoi(msg.String())
				m.scaleValue = val
				return m, nil
			}
		case " ":
			if m.currentInput == "multi-select" {
				return m.handleMultiSelectToggle()
			}
		}

		// Handle input based on current type
		switch m.currentInput {
		case "text":
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		case "select", "multi-select":
			var cmd tea.Cmd
			m.selectList, cmd = m.selectList.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m *QuestionnaireModel) handleEnter() (tea.Model, tea.Cmd) {
	question := getQuestion(m.currentStep)
	if question == nil {
		// Questionnaire complete
		return m, tea.Quit
	}

	// Save answer based on input type
	switch m.currentInput {
	case "text":
		answer := strings.TrimSpace(m.textInput.Value())
		if answer != "" || !question.Required {
			m.answers[question.ID] = answer
			m.textInput.SetValue("")
		} else {
			return m, nil // Don't advance if required and empty
		}
	case "select":
		selected := m.selectList.SelectedItem()
		if selected != nil {
			item := selected.(selectItem)
			m.answers[question.ID] = item.text
		}
	case "scale":
		m.answers[question.ID] = m.scaleValue
	case "multi-select":
		// Save multi-select answers on enter
		selectedOptions := []string{}
		for i, opt := range question.Options {
			if m.multiSelect[i] {
				selectedOptions = append(selectedOptions, opt)
			}
		}
		if len(selectedOptions) > 0 || !question.Required {
			m.answers[question.ID] = strings.Join(selectedOptions, ", ")
		} else {
			return m, nil // Don't advance if required and none selected
		}
	}

	// Move to next question
	m.currentStep++
	if m.currentStep >= m.totalSteps {
		return m, tea.Quit
	}

	return m, m.setupCurrentQuestion()
}

func (m *QuestionnaireModel) handleBack() (tea.Model, tea.Cmd) {
	if m.currentStep > 0 {
		m.currentStep--
		return m, m.setupCurrentQuestion()
	}
	return m, tea.Quit
}

func (m *QuestionnaireModel) handleMultiSelectToggle() (tea.Model, tea.Cmd) {
	selected := m.selectList.SelectedItem()
	if selected != nil {
		item := selected.(multiSelectItem)
		m.multiSelect[item.index] = !m.multiSelect[item.index]
		// Update list
		question := getQuestion(m.currentStep)
		items := make([]list.Item, len(question.Options))
		for i, opt := range question.Options {
			selected := m.multiSelect[i]
			items[i] = multiSelectItem{text: opt, index: i, selected: selected}
		}
		m.selectList.SetItems(items)
	}
	return m, nil
}

func (m *QuestionnaireModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	header := renderQuestionnaireHeader(m.width, m.currentStep, m.totalSteps)

	question := getQuestion(m.currentStep)
	content := m.renderQuestion(question)

	body := lipgloss.NewStyle().
		Width(m.width - 4).
		Height(m.height - lipgloss.Height(header) - 5).
		Padding(1, 2).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
}

func renderQuestionnaireHeader(width, current, total int) string {
	progress := float64(current) / float64(total)
	progressBar := renderProgressBar(width-4, progress)

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Bold(true).
		Width(width - 4).
		Align(lipgloss.Center).
		Render("🎨 Design Preferences & Requirements (DPR)")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")).
		Width(width - 4).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("Question %d of %d", current+1, total))

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		title,
		subtitle,
		"",
		progressBar,
		"",
	)
}

func renderProgressBar(width int, progress float64) string {
	filled := int(float64(width-4) * progress)
	empty := (width - 4) - filled

	filledBar := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#667eea")).
		Render(strings.Repeat("█", filled))

	emptyBar := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#333333")).
		Render(strings.Repeat("░", empty))

	return filledBar + emptyBar
}

func (m *QuestionnaireModel) renderQuestion(q *Question) string {
	if q == nil {
		return "✅ Questionnaire complete!\n\nPress Enter to generate DPR..."
	}

	var content strings.Builder
	content.WriteString(fmt.Sprintf("%s\n\n", q.Text))

	// Show current answer if exists
	if val, ok := m.answers[q.ID]; ok {
		answerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10b981")).
			Italic(true)
		content.WriteString(answerStyle.Render(fmt.Sprintf("Current answer: %v\n\n", val)))
	}

	switch m.currentInput {
	case "text":
		content.WriteString(m.textInput.View())
		content.WriteString("\n\n")
		if q.Placeholder != "" {
			content.WriteString(fmt.Sprintf("💡 Example: %s\n", q.Placeholder))
		}
	case "select":
		content.WriteString(m.selectList.View())
	case "scale":
		content.WriteString("Rate from 1-5:\n\n")
		for i := 1; i <= 5; i++ {
			marker := "  "
			if i == m.scaleValue {
				marker = "→ "
			}
			desc := getScaleDescription(i)
			content.WriteString(fmt.Sprintf("%s%d - %s\n", marker, i, desc))
		}
		content.WriteString("\n")
	case "multi-select":
		content.WriteString("Select multiple options (Space to toggle, Enter when done):\n\n")
		content.WriteString(m.selectList.View())
	}

	help := "\n"
	if m.currentInput == "multi-select" {
		help += "Space: Toggle | Enter: Continue | Esc: Back | Q: Quit"
	} else {
		help += "Enter: Continue | Esc: Back | Q: Quit"
	}
	content.WriteString(help)

	return content.String()
}

func getScaleDescription(value int) string {
	descriptions := map[int]string{
		1: "Not important",
		2: "Somewhat important",
		3: "Moderately important",
		4: "Very important",
		5: "Critical",
	}
	return descriptions[value]
}

// List item types
type selectItem struct {
	text  string
	index int
}

func (i selectItem) FilterValue() string { return i.text }
func (i selectItem) Title() string       { return i.text }
func (i selectItem) Description() string { return "" }

type multiSelectItem struct {
	text     string
	index    int
	selected bool
}

func (i multiSelectItem) FilterValue() string { return i.text }
func (i multiSelectItem) Title() string {
	if i.selected {
		return fmt.Sprintf("✓ %s", i.text)
	}
	return fmt.Sprintf("  %s", i.text)
}
func (i multiSelectItem) Description() string { return "" }

func (m *QuestionnaireModel) toDPRData() *DPRData {
	return &DPRData{
		Answers: m.answers,
	}
}

// DPRData holds the collected questionnaire data
type DPRData struct {
	Answers map[string]interface{}
}

func getTotalQuestions() int {
	return len(getAllQuestions())
}

func getQuestion(index int) *Question {
	questions := getAllQuestions()
	if index < 0 || index >= len(questions) {
		return nil
	}
	return &questions[index]
}

func getAllQuestions() []Question {
	return []Question{
		// Audience Analysis
		{ID: "audience_primary", Category: "audience", Text: "Who is your primary target audience?", Type: "text", Required: true, Placeholder: "e.g., Developers, Designers, End users"},
		{ID: "audience_age", Category: "audience", Text: "What is the age range of your target audience?", Type: "select", Options: []string{"18-25", "26-35", "36-45", "46-55", "55+"}, Required: true},
		{ID: "audience_tech_level", Category: "audience", Text: "What is the technical proficiency level?", Type: "select", Options: []string{"Beginner", "Intermediate", "Advanced", "Expert"}, Required: true},

		// Emotional Design
		{ID: "emotion_primary", Category: "emotion", Text: "What primary emotion should your design evoke?", Type: "select", Options: []string{"Trust", "Excitement", "Calm", "Professional", "Playful", "Innovative"}, Required: true},
		{ID: "emotion_secondary", Category: "emotion", Text: "What secondary emotion?", Type: "select", Options: []string{"Confidence", "Joy", "Serenity", "Authority", "Creativity", "Modern"}, Required: false},

		// Style Preferences
		{ID: "style_overall", Category: "style", Text: "What overall design style do you prefer?", Type: "select", Options: []string{"Minimalist", "Modern", "Classic", "Bold", "Playful", "Professional"}, Required: true},
		{ID: "style_inspiration", Category: "style", Text: "Do you have design inspiration or references?", Type: "text", Required: false, Placeholder: "e.g., Apple, Google Material Design, Stripe"},

		// Colors
		{ID: "color_primary", Category: "colors", Text: "What is your primary brand color?", Type: "text", Required: true, Placeholder: "e.g., #667eea or blue"},
		{ID: "color_secondary", Category: "colors", Text: "What is your secondary color?", Type: "text", Required: false, Placeholder: "e.g., #764ba2 or purple"},
		{ID: "color_scheme", Category: "colors", Text: "What color scheme do you prefer?", Type: "select", Options: []string{"Light", "Dark", "Auto (system preference)", "Both"}, Required: true},

		// Typography
		{ID: "typography_style", Category: "typography", Text: "What typography style do you prefer?", Type: "select", Options: []string{"Sans-serif (modern)", "Serif (classic)", "Monospace (technical)", "Mixed"}, Required: true},
		{ID: "typography_importance", Category: "typography", Text: "How important is typography to your design?", Type: "scale", Required: true},

		// Layout
		{ID: "layout_style", Category: "layout", Text: "What layout style do you prefer?", Type: "select", Options: []string{"Centered", "Full-width", "Sidebar", "Grid", "Asymmetric"}, Required: true},
		{ID: "layout_spacing", Category: "layout", Text: "How much spacing do you prefer?", Type: "select", Options: []string{"Tight", "Moderate", "Generous", "Very spacious"}, Required: true},

		// Components
		{ID: "components_style", Category: "components", Text: "What component style do you prefer?", Type: "select", Options: []string{"Flat", "Elevated (shadows)", "Outlined", "Filled", "Mixed"}, Required: true},
		{ID: "components_interactivity", Category: "components", Text: "How interactive should components be?", Type: "scale", Required: true},

		// Animation
		{ID: "animation_level", Category: "animation", Text: "What level of animation do you prefer?", Type: "select", Options: []string{"None", "Subtle", "Moderate", "Prominent"}, Required: true},
		{ID: "animation_style", Category: "animation", Text: "What animation style?", Type: "select", Options: []string{"Smooth", "Bouncy", "Sharp", "Elegant"}, Required: false},

		// Accessibility
		{ID: "accessibility_importance", Category: "accessibility", Text: "How important is accessibility?", Type: "scale", Required: true},
		{ID: "accessibility_requirements", Category: "accessibility", Text: "Any specific accessibility requirements?", Type: "text", Required: false, Placeholder: "e.g., WCAG AAA, screen reader support"},

		// Responsive Design
		{ID: "responsive_priority", Category: "responsive", Text: "What devices are most important?", Type: "multi-select", Options: []string{"Mobile", "Tablet", "Desktop", "Large screens"}, Required: true},
		{ID: "responsive_approach", Category: "responsive", Text: "What responsive approach?", Type: "select", Options: []string{"Mobile-first", "Desktop-first", "Adaptive", "Fluid"}, Required: true},
	}
}

