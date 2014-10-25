package main

import (
  "gopkg.in/yaml.v2"
  "github.com/savaki/go.hue"
  "os"
  "io/ioutil"
  "path"
)

var (
  config_path = path.Join(os.Getenv("HOME"), ".hueman")
)

type BridgeConfig struct {
  Name string
  Bridge *hue.Bridge
}

func StoreBridge(connectionName string, bridge *hue.Bridge) {
  bridgeConfig := &BridgeConfig{connectionName, bridge}
  bridgeString, _ := yaml.Marshal(bridgeConfig)

  ioutil.WriteFile(config_path, bridgeString, 0777)
}

func LoadBridge() (*BridgeConfig, error) {
  bytes, err := ioutil.ReadFile(config_path)

  if (err != nil) {
    return nil, err
  }

  var config BridgeConfig
  err = yaml.Unmarshal(bytes, &config)
  if (err != nil) {
    return nil, err
  }

  return &config, nil
}
