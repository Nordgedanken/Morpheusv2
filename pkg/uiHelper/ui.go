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
	"github.com/therecipe/qt/widgets"
)

type ui interface {
	GetWidget() (widget *widgets.QWidget)
	NewUI() error
}

// SetNewWindow loads the new UI into the QMainWindow
func SetNewWindow(ui ui, window *widgets.QMainWindow, windowWidth, windowHeight int) error {
	uiErr := ui.NewUI()
	if uiErr != nil {
		return uiErr
	}
	ui.GetWidget().Resize2(windowWidth, windowHeight)
	window.SetCentralWidget(ui.GetWidget())
	window.Show()
	return nil
}
