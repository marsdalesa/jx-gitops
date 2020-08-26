package helmhelpers

import (
	"github.com/jenkins-x/jx-helpers/pkg/cmdrunner"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/pkg/errors"
	"github.com/roboll/helmfile/pkg/state"
)

// AddHelmRepositories ensures the repositories in the helmfile are added to helm
// so that we can use helm templating etc
func AddHelmRepositories(helmState state.HelmState, runner cmdrunner.CommandRunner) error {
	repoMap := map[string]string{
		"jx": "http://chartmuseum.jenkins-x.io",
	}
	for _, repo := range helmState.Repositories {
		repoMap[repo.Name] = repo.URL
	}

	for repoName, repoURL := range repoMap {
		c := &cmdrunner.Command{
			Name: "helm",
			Args: []string{"repo", "add", repoName, repoURL},
		}
		_, err := runner(c)
		if err != nil {
			return errors.Wrap(err, "failed to add helm repo")
		}
		log.Logger().Infof("added helm repository %s %s", repoName, repoURL)
	}
	return nil
}
