// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gametickmaker

import (
	"fmt"
	"testing"
	"time"

	"github.com/kasworld/gametick"
)

func TestFrame3(t *testing.T) {
	gtm := New(943865745419, 1)

	tf := gametick.GameTick_None
	t.Logf("%v", tf)
	t.Logf("%v", gtm.ToUTCTime(tf))
	t.Logf("%v", gtm.ToUTCTime(tf).UTC())

	t.Logf("%v", time.Unix(0, 0).UTC().UnixNano())

	t1 := time.Now()
	t.Logf("%v %v", t1.UnixNano(), t1.UTC().UnixNano())

	t2 := time.Now().UTC()

	t.Logf("%v %v %v", t1, t2, t2.Sub(t1))
}

func TestSpeedChange(t *testing.T) {
	gtm := New(0, 1)
	fmt.Printf("%v\n", gtm)
	fmt.Printf("%v\n", gtm.GetGameTick())
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", gtm.GetGameTick())
	gtm = gtm.GetAcceleratedBy(2)
	fmt.Printf("%v\n", gtm.GetGameTick())
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", gtm.GetGameTick())
	gtm = gtm.GetAcceleratedBy(0.5)
	fmt.Printf("%v\n", gtm.GetGameTick())
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", gtm.GetGameTick())
}
