package variablefinders

import (
	"os"

	jxc "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned"
	"github.com/jenkins-x/jx-api/pkg/config"
	"github.com/jenkins-x/jx-helpers/pkg/kube/jxenv"
	"github.com/pkg/errors"
)

// FindRepositoryURL finds the chart repository URL via environment variables or the dev Environment CRD
func FindRepositoryURL(jxClient jxc.Interface, ns string, requirements *config.RequirementsConfig) (string, error) {
	answer := os.Getenv("JX_CHART_REPOSITORY")
	if answer == "" {
		if requirements != nil {
			answer = requirements.Cluster.ChartRepository
		}
	}
	if answer == "" {
		// assume default chart museum
		answer = "http://jenkins-x-chartmuseum:8080"
	}
	return answer, nil
}

// FindRequirements finds the requirements from the dev Environment CRD
func FindRequirements(jxClient jxc.Interface, ns string) (*config.RequirementsConfig, error) {
	// try the dev environment
	devEnv, err := jxenv.GetDevEnvironment(jxClient, ns)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find the dev Environment in namespace %s", ns)
	}

	requirements, err := config.GetRequirementsConfigFromTeamSettings(&devEnv.Spec.TeamSettings)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load requirements from dev environment")
	}
	return requirements, nil
}