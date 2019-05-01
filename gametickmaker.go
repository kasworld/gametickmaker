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

// gametick을 생성 관리.
// 서버가 종료 되면 흐름이 멎고 재시작하면 계속 됨.
package gametickmaker

import (
	"fmt"
	"time"

	"github.com/kasworld/gametick"
)

func (gtm GameTickMaker) String() string {
	return fmt.Sprintf(
		"GameTickMaker[%v %v %f]",
		gtm.startTime,
		gtm.lastTick,
		gtm.nTimes)
}

type GameTickMaker struct {
	startTime time.Time         // maker create time
	lastTick  gametick.GameTick // initial tick == last server end tick
	nTimes    float64           // tick progress speed, 1 : normal
}

func New(lasttick gametick.GameTick, ntimes float64) GameTickMaker {
	if ntimes <= 0 {
		panic(fmt.Sprintf("GameTick speed cannot < 0 ntimes:%v %v", ntimes, lasttick))
	}
	gtm := GameTickMaker{
		startTime: time.Now().UTC(),
		lastTick:  lasttick,
		nTimes:    ntimes,
	}
	return gtm
}

func (gtm GameTickMaker) GetGameTick() gametick.GameTick {
	return gtm.FromTimeToTickType(time.Now().UTC())
}

func (gtm GameTickMaker) ToUTCTime(t gametick.GameTick) time.Time {
	tnow := (t - gtm.lastTick) / gametick.GameTick(gtm.nTimes)
	return gtm.startTime.Add(tnow.ToTimeDuration())
}

func (gtm GameTickMaker) FromTimeToTickType(s time.Time) gametick.GameTick {
	stdiff := gametick.FromTimeDurationToTickType(s.Sub(gtm.startTime))
	return stdiff*gametick.GameTick(gtm.nTimes) + gtm.lastTick
}

func (gtm GameTickMaker) GetAcceleratedBy(ntimes float64) GameTickMaker {
	if ntimes <= 0 {
		return gtm
	}
	return GameTickMaker{
		startTime: time.Now().UTC(),
		lastTick:  gtm.GetGameTick(),
		nTimes:    gtm.nTimes * ntimes,
	}
}
