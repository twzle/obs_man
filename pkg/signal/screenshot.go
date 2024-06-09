package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var ScreenshotSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[ScreenshotSaved](),
}

// Representation of screenshot saved signal
type ScreenshotSaved struct {
	SavedScreenshotPath string `hubman:"saved_screenshot_path"`
}

// Function returns string code of signal
func (s ScreenshotSaved) Code() string {
	return "ScreenshotSaved"
}

// Function returns string description of signal
func (s ScreenshotSaved) Description() string {
	return "Sent when a screenshot has been saved"
}
