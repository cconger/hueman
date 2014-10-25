package main

import (
  "github.com/savaki/go.hue"
  "github.com/lucasb-eyer/go-colorful"
  "fmt"
  "flag"
  "strconv"
)

func connectToDevice(connectionName string) {
  locators, err := hue.DiscoverBridges(false)
  if (err != nil) {
    fmt.Printf("Error discovering bridges %s", err)
    return
  }

  // Selecting the first one because I'm lazy
  locator := locators[0]
  deviceType := "hueman"

  bridge, err := locator.CreateUser(deviceType)
  if (err != nil) {
    fmt.Printf("Error creating User %s", err)
    return
  }

  fmt.Printf("registered new device => %+v\n", bridge)
  StoreBridge(connectionName, bridge)
}

func main() {
  connectionName := flag.String("connect", "", "After pressing the button on the bridge, pass a string to be the name to the new connection to that bridge")
  color := flag.String("color", "", "Pass a hex string for the color")
  hueVal := flag.Int("hue", -1, "Hue value to be set, [0,360]")
  satVal := flag.Int("sat", -1, "Saturation value to be set, [0,100]")
  brightness := flag.Int("brightness", -1, "Pass a number from 0-100 for brightness of the lights")

  flag.Parse()

  if (connectionName != nil && *connectionName != "") {
    connectToDevice(*connectionName)
    return
  }

  bridgeConfig, err := LoadBridge()
  if (err != nil) {
    fmt.Printf("Configuration not found, have you connected yet?\n%s\n", err)
    return
  }

  lights, err := bridgeConfig.Bridge.GetAllLights()
  if (err != nil) {
    fmt.Printf("Could not get lights %s\n", err)
    return
  }

  for _, light := range lights {
    var targetState hue.SetLightState

    if (*color != "") {
      c, err := colorful.Hex(*color)
      if (err != nil) {
        fmt.Printf("Could not parse color %s: %s", *color, err)
      }
      h,s,v := c.Hsv()
      fmt.Printf("H: %v, S: %v, V: %v\n", h,s,v)
      targetState.Bri = strconv.Itoa(int(v * 255))
      targetState.Sat = strconv.Itoa(int(s * 255))
      targetState.Hue = strconv.Itoa(int((h * 65535) / 360))
    }

    if (*brightness >= 0) {
      targetValue := strconv.Itoa((*brightness * 255) / 100)
      targetState.Bri = targetValue
    }

    if (*hueVal != -1) {
      targetState.Hue = strconv.Itoa(int((*hueVal * 65535) / 360))
    }

    if (*satVal != -1) {
      targetState.Sat = strconv.Itoa(int((*satVal * 255) / 100))
    }

    light.SetState(targetState)
  }
}
