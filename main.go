package main

import (
	"encoding/json"
	"os"

	"github.com/satshub/opscript/cmd"
	"github.com/satshub/opscript/internal"
	"github.com/satshub/opscript/spec"
	"github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	specData, err := internal.Asset("spec.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	scriptSpec := make(spec.Script)
	if err := json.Unmarshal(specData, &scriptSpec); err != nil {
		logrus.Fatalln(err)
	}

	if err := cmd.New(scriptSpec).Execute(); err != nil {
		os.Exit(1)
	}
}
