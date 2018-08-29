package colorful

import (
  "strings"
)

var colorCodes = map[string]string{
  "black":         "0m",
  "light_black":   "240m",
  "red":           "9m",
  "light_red":     "204m",
  "green":         "2m",
  "light_green":   "120m",
  "yellow":        "11m",
  "light_yellow":  "228m",
  "blue":          "20m",
  "light_blue":    "33m",
  "magenta":       "200m",
  "light_magenta": "213m",
  "cyan":          "51m",
  "light_cyan":    "159m",
  "white":         "15m",
  "light_white":   "251m",
}

func Build(s string, color string) string {
  colorKey := colorCodes[color]

  colorizedString := []string{"\033[38;5;", colorKey, s, "\033[0m"}
  return strings.Join(colorizedString, "")
}