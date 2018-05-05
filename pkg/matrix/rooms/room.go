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

package rooms

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

type Room struct {
}

func (r *Room) SetRoomID(id string) {

}

func (r *Room) SetRoomAliases([]string) {

}

func (r *Room) SetName(string) {

}
func (r *Room) SetAvatar(string) {

}
func (r *Room) SetTopic(string) {

}
func (r *Room) SetMessages([]matrix.Message) {

}

func (r *Room) SetMessageIDS([]string) {

}

func (r *Room) GetRoomID() string {
	return ""
}

func (r *Room) GetRoomAliases() []string {
	return nil
}

func (r *Room) GetName() (string, error) {
	return "", nil
}

func (r *Room) GetAvatar() (string, error) {
	return "", nil
}

func (r *Room) GetTopic() (string, error) {
	return "", nil
}

func (r *Room) GetMessages() []matrix.Message {
	return nil
}

func (r *Room) GetMessageIDS() ([]string, error) {
	return nil, nil
}
