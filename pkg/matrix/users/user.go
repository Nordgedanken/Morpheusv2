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

package users

import "github.com/matrix-org/gomatrix"

type User struct {
}

func (u *User) SetCli(cli *gomatrix.Client) {

}

func (u *User) SetMXID(id string) {

}

func (u *User) SetDisplayName(roomID string, name string) {

}

func (u *User) SetAvatar(roomID string, avatar string) {

}

func (u *User) GetMXID() string {
	return ""
}

func (u *User) GetDisplayName(roomID string) (string, error) {
	return "", nil
}

func (u *User) GetAvatar(roomID string) (string, error) {
	return "", nil
}
