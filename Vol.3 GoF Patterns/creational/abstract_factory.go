package main

import (
	"errors"
	"fmt"
)

const (
	Windows = iota
	Linux
)

// ////////Base interface/////////////
type StatusBar interface {
	CreateStatusBar()
}

type MainMenu interface {
	CreateMainMenu()
}

type MainWindow interface {
	CreateMainWindow()
}

// ///////Windows implementation/////////////
type WindowsMainMenu struct{}

func (w *WindowsMainMenu) CreateMainMenu() {
	fmt.Println("Created main menu for Windows")
}

type WindowsStatusBar struct{}

func (w *WindowsStatusBar) CreateStatusBar() {
	fmt.Println("Created status bar for Windows")
}

type WindowsMainWindow struct{}

func (w *WindowsMainWindow) CreateMainWindow() {
	fmt.Println("Created MainWindow for Windows")
}

// ///////Linux implementation/////////////
type LinuxMainMenu struct{}

func (w *LinuxMainMenu) CreateMainMenu() {
	fmt.Println("Created main menu for Linux")
}

type LinuxStatusBar struct{}

func (w *LinuxStatusBar) CreateStatusBar() {
	fmt.Println("Created status bar for Linux")
}

type LinuxMainWindow struct{}

func (w *LinuxMainWindow) CreateMainWindow() {
	fmt.Println("Created MainWindow for Linux")
}

// /////////Base interface abstract factory/////////////
type GuiAbstractFactory interface {
	GetStatusBar() StatusBar
	GetMainMenu() MainMenu
	GetMainWindow() MainWindow
}

// /////////Windows factory/////////////
type WindowsGuiFactory struct{}

func (w WindowsGuiFactory) GetStatusBar() StatusBar {
	return &WindowsStatusBar{}
}

func (w WindowsGuiFactory) GetMainMenu() MainMenu {
	return &WindowsMainMenu{}
}

func (w WindowsGuiFactory) GetMainWindow() MainWindow {
	return &WindowsMainWindow{}
}

// /////////Linux factory/////////////
type LinuxGuiFactory struct{}

func (l LinuxGuiFactory) GetStatusBar() StatusBar {
	return &LinuxStatusBar{}
}

func (l LinuxGuiFactory) GetMainMenu() MainMenu {
	return &LinuxMainMenu{}
}

func (l LinuxGuiFactory) GetMainWindow() MainWindow {
	return &LinuxMainWindow{}
}

// ////////Client code///////////
type Application struct {
	guiFactory GuiAbstractFactory
}

func (a *Application) CreateGui() {
	mainWindow := a.guiFactory.GetMainWindow()
	statusBar := a.guiFactory.GetStatusBar()
	mainMenu := a.guiFactory.GetMainMenu()
	mainWindow.CreateMainWindow()
	mainMenu.CreateMainMenu()
	statusBar.CreateStatusBar()
}

// ////////GUI Factory///////////
func GUIFactory(osType int) (GuiAbstractFactory, error) {
	switch osType {
	case Windows:
		return WindowsGuiFactory{}, nil
	case Linux:
		return LinuxGuiFactory{}, nil
	}
	return nil, errors.New("wrong os type")
}

func main() {
	osType := Windows
	ui, err := GUIFactory(osType)
	if err != nil {
		fmt.Println(err)
	} else {
		app := Application{ui}
		app.CreateGui()
	}
}
