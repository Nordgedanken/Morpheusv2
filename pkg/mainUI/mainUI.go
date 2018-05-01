package mainUI

import (
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/widgets"
)

type MainUI struct {
	ui           *widgets.QWidget
	cli          *gomatrix.Client
	window       *widgets.QMainWindow
	windowWidth  int
	windowHeight int
}

// NewMainUI gives you a MainUI struct with prefilled data
func NewMainUI(windowWidth, windowHeight int, window *widgets.QMainWindow) (mainUI *MainUI) {
	mainUI = &MainUI{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		window:       window,
	}
	return
}

func (m *MainUI) SetCli(cli *gomatrix.Client) {
	m.cli = cli
}

func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	return m.ui
}

func (m *MainUI) NewUI() error {

	return nil
}
