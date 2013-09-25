package config

import (
  "testing"
  "strings"
  "bufio"
  assert "github.com/pilu/miniassert"
)

func TestParse(t *testing.T) {
  content := `
  # comment 1
  ; comment 1

  foo 1
  bar 2

  [section_1]

  foo 3
  bar        4
  baz 5 6
  qux

  [section_2]

  `
  reader := bufio.NewReader(strings.NewReader(content))
  sections, _ := parse(reader, "main")

  assert.Equal(t, 3, len(sections))

  // Main section
  main_section := sections["main"]
  assert.Equal(t, 2, len(main_section))
  assert.Equal(t, "1", main_section["foo"])
  assert.Equal(t, "2", main_section["bar"])

  // Section 1
  section_1 := sections["section_1"]
  assert.Equal(t, 3, len(section_1))
  assert.Equal(t, "3", section_1["foo"])
  assert.Equal(t, "4", section_1["bar"])
  assert.Equal(t, "5 6", section_1["baz"])

  // Section 2
  section_2 := sections["section_2"]
  assert.Equal(t, 0, len(section_2))
}
