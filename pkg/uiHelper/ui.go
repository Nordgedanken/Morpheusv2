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

package uiHelper

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/loginUI"
	"github.com/Nordgedanken/Morpheusv2/pkg/mainUI"
	"github.com/Nordgedanken/Morpheusv2/pkg/registerUI"
	"github.com/therecipe/qt/widgets"
	"log"
)

type ui interface {
	GetWidget() (widget *widgets.QWidget)
	NewUI() error
}

// setNewWindow loads the new UI into the QMainWindow
func setNewWindow(ui ui, window *widgets.QMainWindow, windowWidth, windowHeight int) error {
	log.Println("Start changing UI")
	uiErr := ui.NewUI()
	if uiErr != nil {
		return uiErr
	}
	ui.GetWidget().Resize2(windowWidth, windowHeight)
	window.SetCentralWidget(ui.GetWidget())
	log.Println("Finished changing UI")
	return nil
}

func NewLoginUI(windowWidth, windowHeight int, window *widgets.QMainWindow) {
	loginUIs := loginUI.NewLoginUI(windowWidth, windowHeight, window)
	setNewWindow(loginUIs, window, windowWidth, windowHeight)
}

func NewRegisterUI(windowWidth, windowHeight int, window *widgets.QMainWindow) {
	registerUIs := registerUI.NewRegisterUI(windowWidth, windowHeight, window)
	setNewWindow(registerUIs, window, windowWidth, windowHeight)
}

func NewMainUI(windowWidth, windowHeight int, window *widgets.QMainWindow) {
	mainUIs := mainUI.NewMainUI(windowWidth, windowHeight, window)
	setNewWindow(mainUIs, window, windowWidth, windowHeight)
}
