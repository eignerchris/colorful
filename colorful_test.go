package colorful

import (
  "testing"
  "fmt"
)

func TestBuildOutput(t *testing.T) {
  cases := []struct {
    code, output string
  }{
    {"black", "\033[38;5;0mhello world\033[0m"},
    {"light_black", "\033[38;5;240mhello world\033[0m"},
    {"red", "\033[38;5;9mhello world\033[0m"},
    {"light_red", "\033[38;5;204mhello world\033[0m"},
    {"green", "\033[38;5;2mhello world\033[0m"},
    {"light_green", "\033[38;5;120mhello world\033[0m"},
    {"yellow", "\033[38;5;11mhello world\033[0m"},
    {"light_yellow", "\033[38;5;228mhello world\033[0m"},
    {"blue", "\033[38;5;20mhello world\033[0m"},
    {"light_blue", "\033[38;5;33mhello world\033[0m"},
    {"magenta", "\033[38;5;200mhello world\033[0m"},
    {"light_magenta", "\033[38;5;213mhello world\033[0m"},
    {"cyan", "\033[38;5;51mhello world\033[0m"},
    {"light_cyan", "\033[38;5;159mhello world\033[0m"},
    {"white", "\033[38;5;15mhello world\033[0m"},
    {"light_white", "\033[38;5;251mhello world\033[0m"},

  }

  for _, c := range cases {
    got := Build("hello world", c.code)
    fmt.Println(got)
    if got != c.output {
      t.Errorf("Build(%q, %q) == %q, want %q", "hello world", c.code, got, c.output)
    }
  }
}
