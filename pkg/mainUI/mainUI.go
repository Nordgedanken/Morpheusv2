package mainUI

import (
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// MainUI defines the data for the main ui (that one with the chats)
type MainUI struct {
	widget       *widgets.QWidget
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

// MainUI.SetCli sets the gomatrix Client for the MainUI
func (m *MainUI) SetCli(cli *gomatrix.Client) {
	m.cli = cli
}

// MainUI.GetWidget returns the QWidget of the MainUI
func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	return m.widget
}

// MainUI.NewUI prepares the new UI
func (m *MainUI) NewUI() error {
	m.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, m.widget)
	file.Close()

	var layout = widgets.NewQHBoxLayout()
	m.window.SetLayout(layout)
	layout.InsertWidget(0, mainWidget, 0, core.Qt__AlignTop|core.Qt__AlignLeft)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)

	m.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		mainWidget.Resize(event.Size())
		event.Accept()
	})

	return nil
}
