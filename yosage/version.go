package yosage

import (
	"fmt"
	"time"

	latest "github.com/tcnksm/go-latest"
)

const TIMEOUT_SEC = 1

func verCheck(version string) <-chan *latest.CheckResponse {
	verCheckCh := make(chan *latest.CheckResponse)

	go func() {
		fixFunc := latest.DeleteFrontV()
		githubTag := &latest.GithubTag{
			Owner:             "hico-horiuchi",
			Repository:        "yosage",
			FixVersionStrFunc: fixFunc,
		}

		res, _ := latest.Check(githubTag, fixFunc(version))
		verCheckCh <- res
	}()

	return verCheckCh
}

func Version(version string) string {
	bytes := []byte(fmt.Sprintf("yosage version %s\n", version))
	verCheckCh := verCheck(version)

	for {
		select {
		case res := <-verCheckCh:
			if res != nil && res.Outdated {
				bytes = append(bytes, fmt.Sprintf("Latest version of yosage is %s, please update it\n", res.Current)...)
			}
			return string(bytes)
		case <-time.After(TIMEOUT_SEC * time.Second):
			return string(bytes)
		}
	}
}
