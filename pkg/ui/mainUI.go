// Copyright © 2018 MTRNord <info@nordgedanken.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"log"
)

// MainUI defines the data for the main ui (that one with the chats)
type MainUI struct {
	widget       *widgets.QWidget
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

// GetWidget returns the QWidget of the MainUI
func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	return m.widget
}

// NewUI prepares the new UI
func (m *MainUI) NewUI() error {
	m.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, m.widget)
	file.Close()

	m.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	mainWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	m.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		mainWidget.Resize(event.Size())
		event.Accept()
	})

	// Setup functions and elements
	go m.setupLogout()

	m.window.SetWindowTitle("Morpheus")

	return nil
}

func (m *MainUI) setupLogout() {
	// Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(m.widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		go m.logout()
	})
	return
}

func (m *MainUI) logout() {
	_, err := util.User.GetCli().Logout()
	if err != nil {
		log.Panicln(err)
	}
	util.User.GetCli().ClearCredentials()

	err = util.DB.RemoveCurrentUser()
	if err != nil {
		log.Panicln(err)
	}

	loginUIs := NewLoginUI(m.windowWidth, m.windowHeight, m.window)
	err = SetNewWindow(loginUIs, m.window, m.windowWidth, m.windowHeight)
	if err != nil {
		log.Panicln(err)
	}
}
