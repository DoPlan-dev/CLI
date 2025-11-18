package dpr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// DesignTokens represents the design tokens structure
type DesignTokens struct {
	Colors       ColorTokens        `json:"colors"`
	Typography   TypographyTokens   `json:"typography"`
	Spacing      SpacingTokens      `json:"spacing"`
	BorderRadius BorderRadiusTokens `json:"borderRadius"`
	Shadows      ShadowTokens       `json:"shadows"`
	Breakpoints  BreakpointTokens   `json:"breakpoints"`
}

type ColorTokens struct {
	Primary   string            `json:"primary"`
	Secondary string            `json:"secondary"`
	Accent    string            `json:"accent,omitempty"`
	Neutral   map[string]string `json:"neutral"`
	Semantic  SemanticColors    `json:"semantic"`
}

type SemanticColors struct {
	Success string `json:"success"`
	Warning string `json:"warning"`
	Error   string `json:"error"`
	Info    string `json:"info"`
}

type TypographyTokens struct {
	FontFamily map[string]string `json:"fontFamily"`
	FontSize   map[string]string `json:"fontSize"`
	FontWeight map[string]string `json:"fontWeight"`
	LineHeight map[string]string `json:"lineHeight"`
}

type SpacingTokens struct {
	Scale  string            `json:"scale"`
	Values map[string]string `json:"values"`
}

type BorderRadiusTokens struct {
	None   string `json:"none"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
	Full   string `json:"full"`
}

type ShadowTokens struct {
	None   string `json:"none"`
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type BreakpointTokens struct {
	Mobile  string `json:"mobile"`
	Tablet  string `json:"tablet"`
	Desktop string `json:"desktop"`
	Large   string `json:"large"`
}

// TokenGenerator generates design-tokens.json
type TokenGenerator struct {
	projectRoot string
	data        *DPRData
}

// NewTokenGenerator creates a new token generator
func NewTokenGenerator(projectRoot string, data *DPRData) *TokenGenerator {
	return &TokenGenerator{
		projectRoot: projectRoot,
		data:        data,
	}
}

// Generate creates the design-tokens.json file
func (tg *TokenGenerator) Generate() error {
	tokensPath := filepath.Join(tg.projectRoot, "doplan", "design", "design-tokens.json")

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(tokensPath), 0755); err != nil {
		return fmt.Errorf("failed to create design directory: %w", err)
	}

	tokens := tg.generateTokens()

	jsonData, err := json.MarshalIndent(tokens, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tokens: %w", err)
	}

	return os.WriteFile(tokensPath, jsonData, 0644)
}

func (tg *TokenGenerator) generateTokens() *DesignTokens {
	primary := tg.getAnswerString("color_primary", "#667eea")
	secondary := tg.getAnswerString("color_secondary", "#764ba2")
	scheme := tg.getAnswerString("color_scheme", "Both")

	// Generate neutral colors based on scheme
	neutral := tg.generateNeutralColors(scheme)

	return &DesignTokens{
		Colors: ColorTokens{
			Primary:   primary,
			Secondary: secondary,
			Accent:    tg.generateAccentColor(primary, secondary),
			Neutral:   neutral,
			Semantic: SemanticColors{
				Success: "#10b981",
				Warning: "#f59e0b",
				Error:   "#ef4444",
				Info:    "#3b82f6",
			},
		},
		Typography: tg.generateTypographyTokens(),
		Spacing:    tg.generateSpacingTokens(),
		BorderRadius: BorderRadiusTokens{
			None:   "0",
			Small:  "4px",
			Medium: "8px",
			Large:  "12px",
			Full:   "9999px",
		},
		Shadows: ShadowTokens{
			None:   "none",
			Small:  "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
			Medium: "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
			Large:  "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)",
		},
		Breakpoints: BreakpointTokens{
			Mobile:  "320px",
			Tablet:  "768px",
			Desktop: "1024px",
			Large:   "1440px",
		},
	}
}

func (tg *TokenGenerator) generateNeutralColors(scheme string) map[string]string {
	neutral := make(map[string]string)

	if scheme == "Dark" || scheme == "Both" {
		neutral["50"] = "#18181b"
		neutral["100"] = "#27272a"
		neutral["200"] = "#3f3f46"
		neutral["300"] = "#52525b"
		neutral["400"] = "#71717a"
		neutral["500"] = "#a1a1aa"
		neutral["600"] = "#d4d4d8"
		neutral["700"] = "#e4e4e7"
		neutral["800"] = "#f4f4f5"
		neutral["900"] = "#fafafa"
	} else {
		neutral["50"] = "#fafafa"
		neutral["100"] = "#f4f4f5"
		neutral["200"] = "#e4e4e7"
		neutral["300"] = "#d4d4d8"
		neutral["400"] = "#a1a1aa"
		neutral["500"] = "#71717a"
		neutral["600"] = "#52525b"
		neutral["700"] = "#3f3f46"
		neutral["800"] = "#27272a"
		neutral["900"] = "#18181b"
	}

	return neutral
}

func (tg *TokenGenerator) generateAccentColor(primary, secondary string) string {
	// Generate accent color by blending primary and secondary
	// For now, return a complementary color
	return "#f59e0b"
}

func (tg *TokenGenerator) generateTypographyTokens() TypographyTokens {
	style := tg.getAnswerString("typography_style", "Sans-serif")

	fontFamily := make(map[string]string)
	switch style {
	case "Sans-serif (modern)":
		fontFamily["sans"] = "system-ui, -apple-system, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif"
		fontFamily["mono"] = "'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, monospace"
	case "Serif (classic)":
		fontFamily["serif"] = "Georgia, 'Times New Roman', Times, serif"
		fontFamily["sans"] = "system-ui, sans-serif"
	case "Monospace (technical)":
		fontFamily["mono"] = "'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, monospace"
		fontFamily["sans"] = "system-ui, sans-serif"
	default:
		fontFamily["sans"] = "system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif"
		fontFamily["mono"] = "'SF Mono', Monaco, monospace"
	}

	return TypographyTokens{
		FontFamily: fontFamily,
		FontSize: map[string]string{
			"xs":   "0.75rem",
			"sm":   "0.875rem",
			"base": "1rem",
			"lg":   "1.125rem",
			"xl":   "1.25rem",
			"2xl":  "1.5rem",
			"3xl":  "1.875rem",
			"4xl":  "2.25rem",
			"5xl":  "3rem",
		},
		FontWeight: map[string]string{
			"thin":      "100",
			"light":     "300",
			"normal":    "400",
			"medium":    "500",
			"semibold":  "600",
			"bold":      "700",
			"extrabold": "800",
			"black":     "900",
		},
		LineHeight: map[string]string{
			"none":    "1",
			"tight":   "1.25",
			"snug":    "1.375",
			"normal":  "1.5",
			"relaxed": "1.625",
			"loose":   "2",
		},
	}
}

func (tg *TokenGenerator) generateSpacingTokens() SpacingTokens {
	spacing := tg.getAnswerString("layout_spacing", "Moderate")

	scale := "1rem" // Base unit
	if spacing == "Tight" {
		scale = "0.75rem"
	} else if spacing == "Generous" {
		scale = "1.25rem"
	} else if spacing == "Very spacious" {
		scale = "1.5rem"
	}

	return SpacingTokens{
		Scale: scale,
		Values: map[string]string{
			"0":  "0",
			"1":  fmt.Sprintf("calc(%s * 0.25)", scale),
			"2":  fmt.Sprintf("calc(%s * 0.5)", scale),
			"3":  fmt.Sprintf("calc(%s * 0.75)", scale),
			"4":  scale,
			"5":  fmt.Sprintf("calc(%s * 1.25)", scale),
			"6":  fmt.Sprintf("calc(%s * 1.5)", scale),
			"8":  fmt.Sprintf("calc(%s * 2)", scale),
			"10": fmt.Sprintf("calc(%s * 2.5)", scale),
			"12": fmt.Sprintf("calc(%s * 3)", scale),
			"16": fmt.Sprintf("calc(%s * 4)", scale),
			"20": fmt.Sprintf("calc(%s * 5)", scale),
			"24": fmt.Sprintf("calc(%s * 6)", scale),
			"32": fmt.Sprintf("calc(%s * 8)", scale),
			"40": fmt.Sprintf("calc(%s * 10)", scale),
			"48": fmt.Sprintf("calc(%s * 12)", scale),
			"56": fmt.Sprintf("calc(%s * 14)", scale),
			"64": fmt.Sprintf("calc(%s * 16)", scale),
		},
	}
}

func (tg *TokenGenerator) getAnswerString(key string, defaultValue string) string {
	if val, ok := tg.data.Answers[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", val)
	}
	return defaultValue
}
