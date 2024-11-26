package main

import "testing"

func TestParseWailsConfig(t *testing.T) {
	wailsConfig, err := ParseWailsConfig()
	if err != nil {
		t.Fatal("error parsing wails.json.", err, wailsConfig)
	}

	want := "Marcus Nilszén"
	got := wailsConfig.Author.Name

	if want != got {
		t.Fatal(want, got)
	}
}
