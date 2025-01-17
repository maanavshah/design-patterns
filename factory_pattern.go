package main

import "fmt"

type ModeFactory interface {
	SearchWithMode()
}

type AutosuggestMode struct{}

func (m AutosuggestMode) SearchWithMode() {
	fmt.Println("AutosuggestMode enabled")
}

type ThemeMode struct{}

func (m ThemeMode) SearchWithMode() {
	fmt.Println("ThemeMode enabled")
}

type DefaultSearchMode struct{}

func (m DefaultSearchMode) SearchWithMode() {
	fmt.Println("DefaultSearchMode enabled")
}

func GetFactory(mode string) ModeFactory {
	switch mode {
	case "autosuggest":
		return AutosuggestMode{}
	case "theme":
		return ThemeMode{}
	}
	return DefaultSearchMode{}
}

func main() {
	modeFactory := GetFactory("autosuggest")
	modeFactory.SearchWithMode()

	modeFactory = GetFactory("theme")
	modeFactory.SearchWithMode()

	modeFactory = GetFactory("default")
	modeFactory.SearchWithMode()
}
