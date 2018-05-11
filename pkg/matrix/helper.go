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

package matrix

import (
	"github.com/therecipe/qt/gui"
)

func ImageToQImage(image []byte) *gui.QImage {
	imageS := string(image[:])
	qimage := gui.QImage_FromData(string(imageS[:]), len(imageS), "")

	return qimage
}

func ImageToPixmap(image []byte) *gui.QPixmap {
	pixmap := gui.NewQPixmap()

	imageS := string(image[:])

	pixmap.LoadFromData(string(imageS[:]), uint(len(imageS)), "", 0)

	return pixmap
}
