package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type envList struct {
	timezone       string
	cciApiInterval string
	ghRepos        string
	ghBranches     string
	cciToken       string
}

func (e *envList) TokenInit(token string) {
	e.cciToken = stringOrDefault(token, "noMoreSecrets")
}

func (e *envList) SetEnvVars() {
	envVars := strings.Split(e.ToEnvFile(), "\n")

	for _, envVar := range envVars {
		kv := strings.Split(envVar, "=")
		if kv[0] != "" && kv[1] != "" {
			os.Setenv(kv[0], kv[1])
		}
	}
}

func (e *envList) WriteEnvFile(outFile string) error {
	_, err := os.Stat(outFile)
	if err == nil {
		return fmt.Errorf("file already exist")
	}

	err = ioutil.WriteFile(outFile, []byte(e.ToEnvFile()), 0666)

	if err != nil {
		return err
	}

	return nil
}

func (e *envList) ToEnvFile() string {
	var data strings.Builder
	if e.cciApiInterval != "" {
		data.WriteString(fmt.Sprintf("CIRCLECI_API_INTERVAL=%s", e.cciApiInterval))
	}
	if e.cciToken != "" {
		data.WriteString(fmt.Sprintf("\nCIRCLECI_TOKEN=%s", e.cciToken))
	}
	if e.ghBranches != "" {
		data.WriteString(fmt.Sprintf("\nGITHUB_BRANCH=%s", e.ghBranches))
	}
	if e.timezone != "" {
		data.WriteString(fmt.Sprintf("\nTZ=%s", e.timezone))
	}
	if e.ghRepos != "" {
		data.WriteString(fmt.Sprintf("\nGITHUB_REPOSITORY=%s", e.ghRepos))
	}
	return data.String()
}

func NewEnvList() envList {
	return envList{
		timezone:       "America/New_York",
		cciApiInterval: "300",
		ghRepos:        "aetna-digital-infrastructure/poplarj-hello-world", // "aetna-digital-infrastructure/DSSP-Developer-Portal", // ,aetna-digital-infrastructure/dfp-gitops-package-ccloud-definitions,aetna-digital-infrastructure/rally-service",
		ghBranches:     "fix/cci-image",                                    //"master",                                           // ,main",
		cciToken:       "noMoreSecrets",
	}
}
