package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func isFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	return err == nil
}

// true for any number, can be float and negative
func isValidNumber(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		if !isFloat(str) {
			return false
		} else {
			err = nil
		}
	}
	return err == nil
}

// for values like window size
func isValidUint(str string) bool {
	_, err := strconv.ParseUint(str, 10, 64)
	return err == nil
}

var Effect1Label *widget.Label
var Effect2Label *widget.Label
var CustomTabletLabel *widget.Label
var EnableErrorsLabel *widget.Label
var TabletAreaLabel *widget.Label
var WidowSizeLabel *widget.Label
var CustomTabletSizeLabel *widget.Label
var CustomTabletActiveAreaLabel *widget.Label
var CustomTabletImageOffsetLabel *widget.Label
var CursorSizeLabel *widget.Label
var CursorSizeDataLabel *widget.Label
var CursorTrailSizeLabel *widget.Label
var CursorTrailSizeDataLabel *widget.Label
var FramerateLimitLabel *widget.Label
var FramerateLimitDataLabel *widget.Label
var TrailCirclesLifetimeLabel *widget.Label
var TrailCirclesLifetimeDataLabel *widget.Label
var CursorTrailDensityLabel *widget.Label
var CursorTrailDensityDataLabel *widget.Label
var TabletTransparencyLabel *widget.Label
var TabletTransparencyDataLabel *widget.Label

var Effect1Check *widget.Check
var Effect2Check *widget.Check
var CustomTabletCheck *widget.Check
var EnableErrorsCheck *widget.Check

var TabletAreaX *widget.Entry
var TabletAreaY *widget.Entry
var TabletAreaW *widget.Entry
var TabletAreaH *widget.Entry
var WindowSizeW *widget.Entry
var WindowSizeH *widget.Entry
var CustomTabletSizeW *widget.Entry
var CustomTabletSizeH *widget.Entry
var CustomTabletActiveAreaW *widget.Entry
var CustomTabletActiveAreaH *widget.Entry
var CustomTabletImageOffsetX *widget.Entry
var CustomTabletImageOffsetY *widget.Entry

var CursorSize *widget.Slider
var CursorTrailSize *widget.Slider
var FramerateLimit *widget.Slider
var TrailCirclesLifetime *widget.Slider
var CursorTrailDensity *widget.Slider
var TabletTransparency *widget.Slider

var CursorSizeVal binding.String
var CursorTrailSizeVal binding.String
var FramerateLimitVal binding.String
var TrailCirclesLifetimeVal binding.String
var CursorTrailDensityVal binding.String
var TabletTransparencyVal binding.String

var TabletAreaContainter *fyne.Container
var WindowSizeContainer *fyne.Container
var CustomTabletSizeContainer *fyne.Container
var CustomTabletActiveAreaContainer *fyne.Container
var CustomTabletImageOffsetContainer *fyne.Container
var CursorSizeContainer *fyne.Container
var CursorTrailSizeContainer *fyne.Container
var FramerateLimitContainer *fyne.Container
var TrailCirclesLifetimeContainer *fyne.Container
var CursorTrailDensityContainer *fyne.Container
var TabletTransparencyContainer *fyne.Container
var Labels *fyne.Container
var InputWidgetsContainer *fyne.Container
var LabelsAndInputsContainer *fyne.Container
var MainContainter *fyne.Container

var SaveConfig *widget.Button

func main() {
	App := app.New()
	App.Settings().SetTheme(&CustomTheme{})
	Window := App.NewWindow("Tablet-Visualizer-Configurator")
	Window.Resize(fyne.NewSize(616, 580))

	SetupWidgets()
	SetupAppLayout()
	LoadJSON() // loads settings from file and assigns loaded data to widgets

	Window.SetContent(MainContainter)

	go func() { // Update values of slider so user can see current value on labels
		for range time.Tick(time.Millisecond * 35) {
			CursorSizeVal.Set(strconv.FormatFloat(CursorSize.Value, 'f', 2, 64))
			CursorTrailSizeVal.Set(strconv.FormatFloat(CursorTrailSize.Value, 'f', 2, 64))
			TrailCirclesLifetimeVal.Set(strconv.FormatFloat(TrailCirclesLifetime.Value, 'f', 2, 64))

			FramerateLimitVal.Set(strconv.FormatUint(uint64(FramerateLimit.Value), 10))
			TabletTransparencyVal.Set(strconv.FormatUint(uint64(TabletTransparency.Value), 10))
			CursorTrailDensityVal.Set(strconv.FormatFloat(CursorTrailDensity.Value, 'f', 0, 64))
		}
	}()

	Window.ShowAndRun()
}
