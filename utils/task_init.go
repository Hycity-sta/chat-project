package utils

import "time"

type TimerFunc = func(any) bool

// 按一定间隔执行目标函数
// delay首次延迟, tick间隔, fun定时执行的方法, param方法的参数
func Timer(delay, tick time.Duration, method TimerFunc, param any) {
	action := func() {
		if method == nil {
			return
		}

		timer := time.NewTimer(delay)
		defer timer.Stop()

		for range timer.C {
			if !method(param) {
				return
			}

			timer.Reset(tick)
		}
	}

	go action()
}
