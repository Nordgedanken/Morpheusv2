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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewRoom(room matrix.Room, roomScroll *widgets.QScrollArea) (widgetR *widgets.QWidget, err error) {
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/ui/room.ui")

	file.Open(core.QIODevice__ReadOnly)
	wrapperWidget := loader.Load(file, widget)
	file.Close()

	roomAvatarQLabel := widgets.NewQLabelFromPointer(widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	roomName := widgets.NewQLabelFromPointer(widget.FindChild("roomName", core.Qt__FindChildrenRecursively).Pointer())

	var name string
	name, err = room.GetName()
	if err != nil {
		return
	}
	roomName.SetText(name)

	/*wrapperWidget.Resize2(roomScroll.Widget().Size().Width(), wrapperWidget.Size().Height())
	widget.Resize2(roomScroll.Widget().Size().Width(), wrapperWidget.Size().Height())*/

	var filterObject = core.NewQObject(nil)
	filterObject.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if event.Type() == core.QEvent__MouseButtonPress {
			var mouseEvent = gui.NewQMouseEventFromPointer(event.Pointer())

			if mouseEvent.Button() == core.Qt__LeftButton {
				util.E.Raise("changeRoom", room.GetRoomID())
				return true
			}

			return false
		}

		return false
	})

	wrapperWidget.InstallEventFilter(filterObject)

	var avatar []byte
	avatar, err = room.GetAvatar()
	if err != nil {
		return
	}
	roomAvatarQLabel.SetPixmap(matrix.ImageToPixmap(avatar))

	widgetR = wrapperWidget
	return
}
