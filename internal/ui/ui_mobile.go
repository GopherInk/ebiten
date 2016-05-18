// Copyright 2016 Hajime Hoshi
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

// +build android

package ui

import (
	"github.com/hajimehoshi/ebiten/internal/graphics/opengl"
)

func initialize() (*opengl.Context, error) {
	// TODO: Implement
	return nil, nil
}

func Main() error {
	return nil
}

type UserInterface struct {
}

func CurrentUI() *UserInterface {
	return nil
}

func (u *UserInterface) Start(width, height, scale int, title string) error {
	return nil
}

func (u *UserInterface) Terminate() error {
	return nil
}

func (u *UserInterface) Update() (interface{}, error) {
	return nil, nil
}

func (u *UserInterface) SwapBuffers() error {
	return nil
}

func (u *UserInterface) SetScreenSize(width, height int) bool {
	return false
}

func (u *UserInterface) SetScreenScale(scale int) bool {
	return false
}

func (u *UserInterface) ScreenScale() int {
	return 1
}
