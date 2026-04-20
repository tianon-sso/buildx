package workers

import (
	"context"
	"os/exec"

	"github.com/pkg/errors"
)

// patchBuildkitd replaces /usr/bin/buildkitd in the named container with
// bkdBin and restarts it. dockerEnv is the environment for docker commands
// (must include DOCKER_CONTEXT=... when targeting a non-default daemon).
func patchBuildkitd(ctx context.Context, ctr, bkdBin string, dockerEnv []string) error {
	for _, args := range [][]string{
		{"docker", "cp", bkdBin, ctr + ":/usr/bin/buildkitd"},
		{"docker", "restart", ctr},
	} {
		c := exec.CommandContext(ctx, args[0], args[1:]...)
		c.Env = dockerEnv
		if out, err := c.CombinedOutput(); err != nil {
			return errors.Wrapf(err, "patching buildkitd in %s: %s", ctr, string(out))
		}
	}
	return nil
}
