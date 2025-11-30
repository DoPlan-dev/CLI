package cli

import (
	"fmt"
	"io"
	"strings"
)

// ChallengeIntegration provides easy integration of challenges into commands
type ChallengeIntegration struct {
	system *ChallengeSystem
}

// NewChallengeIntegration creates a new challenge integration helper
func NewChallengeIntegration() (*ChallengeIntegration, error) {
	system, err := NewChallengeSystem()
	if err != nil {
		return nil, err
	}

	return &ChallengeIntegration{
		system: system,
	}, nil
}

// CheckAndCelebrateChallenges checks for completed challenges and celebrates them
func (ci *ChallengeIntegration) CheckAndCelebrateChallenges(context map[string]interface{}, out io.Writer) error {
	if ci.system == nil {
		return nil // Graceful degradation
	}

	// Check for challenges
	completed, err := ci.system.CheckAndAwardChallenges(context)
	if err != nil {
		return fmt.Errorf("failed to check challenges: %w", err)
	}

	// Celebrate if any completed
	if len(completed) > 0 {
		if len(completed) == 1 {
			CelebrateChallenge(completed[0], out)
		} else {
			CelebrateMultipleChallenges(completed, out)
		}
	}

	return nil
}

// CheckOnIntegration checks challenges when integration is completed
func (ci *ChallengeIntegration) CheckOnIntegration(integrationType string, tested bool, context map[string]interface{}, out io.Writer) error {
	context["integration"] = integrationType
	context["api_created"] = true
	if tested {
		context["api_tested"] = true
		context["tests_passed"] = true
	}

	return ci.CheckAndCelebrateChallenges(context, out)
}

// CheckOnDatabaseAction checks challenges when database action is completed
func (ci *ChallengeIntegration) CheckOnDatabaseAction(action string, context map[string]interface{}, out io.Writer) error {
	context["database_action"] = action
	if action == "linked" {
		context["database_linked"] = true
	} else if action == "merged" {
		context["database_merged"] = true
	}

	return ci.CheckAndCelebrateChallenges(context, out)
}

// CheckOnDeployment checks challenges when deployment happens
func (ci *ChallengeIntegration) CheckOnDeployment(environment string, context map[string]interface{}, out io.Writer) error {
	context["deployed"] = true
	context["environment"] = environment
	if environment == "production" {
		context["production_deployed"] = true
	}

	return ci.CheckAndCelebrateChallenges(context, out)
}

// CheckOnPublicRelease checks challenges when public release happens
func (ci *ChallengeIntegration) CheckOnPublicRelease(version string, context map[string]interface{}, out io.Writer) error {
	context["public_release"] = true
	context["release_type"] = "public"
	if version != "" {
		context["version"] = version
	}

	return ci.CheckAndCelebrateChallenges(context, out)
}

// CheckOnTestResults checks challenges based on test results
func (ci *ChallengeIntegration) CheckOnTestResults(coverage float64, allPassing bool, context map[string]interface{}, out io.Writer) error {
	context["test_coverage"] = coverage
	context["all_tests_passing"] = allPassing

	return ci.CheckAndCelebrateChallenges(context, out)
}

// CheckOnWorkflowQuality checks challenges based on workflow quality
func (ci *ChallengeIntegration) CheckOnWorkflowQuality(quality string, context map[string]interface{}, out io.Writer) error {
	context["workflow_quality"] = quality
	if quality == "best" {
		context["best_github_workflow"] = true
	}

	return ci.CheckAndCelebrateChallenges(context, out)
}

// GetPendingChallenges returns challenges user can work towards
func (ci *ChallengeIntegration) GetPendingChallenges() []ChallengeDefinition {
	if ci.system == nil {
		return []ChallengeDefinition{}
	}
	return ci.system.GetPendingChallenges()
}

// DisplayPendingChallenges displays challenges user can complete
func (ci *ChallengeIntegration) DisplayPendingChallenges(out io.Writer) {
	if ci.system == nil {
		return
	}

	pending := ci.system.GetPendingChallenges()
	if len(pending) == 0 {
		fmt.Fprintln(out, "\n🎉 Congratulations! You've completed all challenges!")
		return
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🎯 Available Challenges")
	fmt.Fprintln(out, strings.Repeat("-", 60))

	// Group by category
	categories := make(map[string][]ChallengeDefinition)
	for _, challenge := range pending {
		categories[challenge.Category] = append(categories[challenge.Category], challenge)
	}

	for category, challenges := range categories {
		fmt.Fprintf(out, "\n%s:\n", strings.Title(category))
		for _, challenge := range challenges {
			fmt.Fprintf(out, "  %s %s - %d points (%s)\n",
				challenge.Icon,
				challenge.Title,
				challenge.Points,
				challenge.Rarity)
		}
	}

	fmt.Fprintln(out, "")
}
