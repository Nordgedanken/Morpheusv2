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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/rooms"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type RoomLayout struct {
	widgets.QVBoxLayout

	Rooms     map[string]matrix.Room
	RoomCount int

	_ func(roomID string) `slot:"addRoom"`
}

func (r *RoomLayout) NewRoom(roomID string, roomScroll *widgets.QScrollArea) (err error) {
	room := r.Rooms[roomID]
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

	wrapperWidget.Resize2(roomScroll.Widget().Size().Width(), wrapperWidget.Size().Height())
	widget.Resize2(roomScroll.Widget().Size().Width(), wrapperWidget.Size().Height())

	var filterObject = core.NewQObject(nil)
	filterObject.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if event.Type() == core.QEvent__MouseButtonPress {
			var mouseEvent = gui.NewQMouseEventFromPointer(event.Pointer())

			if mouseEvent.Button() == core.Qt__LeftButton {
				util.E.Raise("changeRoom", room)
				return true
			}

			return false
		}

		return false
	})

	wrapperWidget.InstallEventFilter(filterObject)

	util.E.On("setRoomAvatar"+room.GetRoomID(), func(i interface{}) error {
		switch v := i.(type) {
		case *rooms.Room:
			var avatar []byte
			avatar, err := v.GetAvatar()
			if err != nil {
				return err
			}
			roomAvatarQLabel.SetPixmap(matrix.ImageToPixmap(avatar))
		}

		if (r.RoomCount % 5) == 0 {
			util.App.ProcessEvents(core.QEventLoop__AllEvents)
		}

		return nil
	})

	r.InsertWidget(r.Count()+1, wrapperWidget, 0, 0)
	return
}
