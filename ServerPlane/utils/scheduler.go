package utils
import "../core/ITimer"

var Scheduler *ITimer.TimerScheduler
var IntervalForever int

func init()  {
	IntervalForever = int(^uint(0) >> 1)
	Scheduler = ITimer.NewAutoExecTimerScheduler()
}

