package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	//"gopkg.in/yaml.v2"
	yaml "github.com/ghodss/yaml"
)

const (
	yamlPath = "./"
	yamlFile = "testmc.yaml"
)

func parseMachineConfigYAML(yamlData []byte) (*mcfgv1.MachineConfig, error) {
	//machineConfig := &mcfgv1.MachineConfig{}
	machineConfig := new(mcfgv1.MachineConfig)
	err := yaml.Unmarshal(yamlData, machineConfig)
	if err != nil {
		return nil, err
	}
	return machineConfig, nil
}

func readMachineConfigYAML(mcFileName string) ([]byte, error) {
	machineConfigFilePath := filepath.Join(yamlPath, yamlFile)
	yamlData, err := ioutil.ReadFile(machineConfigFilePath)
	if err != nil {
		return nil, err
	}
	return yamlData, nil
}

func main() {
	yamlData, err := readMachineConfigYAML(yamlFile)
	if err != nil {
		fmt.Printf("Unable to read MachineConfigYaml (%s) : err (%v)\n", yamlFile, err)
		return
	}
	fmt.Printf("yamlData (%s)\n", yamlData)

	machineConfig, err := parseMachineConfigYAML(yamlData)
	if err != nil {
		fmt.Printf("Unable to parse MachineConfigYaml (%s) : err (%v)\n", yamlFile, err)
		return
	}

	fmt.Printf("machineConfig (%s)\n", *machineConfig)
}
