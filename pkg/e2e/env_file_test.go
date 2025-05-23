/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package e2e

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/icmd"
)

func TestRawEnvFile(t *testing.T) {
	c := NewParallelCLI(t)
	defer c.cleanupWithDown(t, "dotenv")

	res := c.RunDockerComposeCmd(t, "-f", "./fixtures/dotenv/raw.yaml", "run", "test")
	assert.Equal(t, strings.TrimSpace(res.Stdout()), "'{\"key\": \"value\"}'")
}

func TestUnusedMissingEnvFile(t *testing.T) {
	c := NewParallelCLI(t)
	defer c.cleanupWithDown(t, "unused_dotenv")

	c.RunDockerComposeCmd(t, "-f", "./fixtures/env_file/compose.yaml", "up", "-d", "serviceA")
}

func TestRunEnvFile(t *testing.T) {
	c := NewParallelCLI(t)
	defer c.cleanupWithDown(t, "run_dotenv")

	res := c.RunDockerComposeCmd(t, "--project-directory", "./fixtures/env_file", "run", "serviceC", "env")
	res.Assert(t, icmd.Expected{Out: "FOO=BAR"})
}
