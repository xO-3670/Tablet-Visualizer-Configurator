package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SetupWidgets() {
	SetupLaybels()

	SetupCheckboxes()

	SetupEntries()

	SetupSliders()

	SetupButtons()
}

func SetupLaybels() {
	Effect1Label = widget.NewLabel("Trail circles fading effect")
	Effect2Label = widget.NewLabel("Trail circles spacing out effect")
	CustomTabletLabel = widget.NewLabel("Custom tablet")
	EnableErrorsLabel = widget.NewLabel("Enable errors")
	TabletAreaLabel = widget.NewLabel("Tablet area")
	WidowSizeLabel = widget.NewLabel("Window size")
	CustomTabletSizeLabel = widget.NewLabel("Custom tablet size")
	CustomTabletActiveAreaLabel = widget.NewLabel("Custom active area size")
	CustomTabletImageOffsetLabel = widget.NewLabel("Custom image offset")
	CursorSizeLabel = widget.NewLabel("Cursor size")
	CursorTrailSizeLabel = widget.NewLabel("Cursor Trail size")
	FramerateLimitLabel = widget.NewLabel("Framerate limit")
	TrailCirclesLifetimeLabel = widget.NewLabel("Trail circles lifetime")
	CursorTrailDensityLabel = widget.NewLabel("Trail circles density")
	TabletTransparencyLabel = widget.NewLabel("Tablet image transparency")
}

func SetupCheckboxes() {
	Effect1Check = widget.NewCheck("", func(b bool) {})
	Effect2Check = widget.NewCheck("", func(b bool) {})
	CustomTabletCheck = widget.NewCheck("", func(b bool) {})
	EnableErrorsCheck = widget.NewCheck("", func(b bool) {})
}

func SetupEntries() {
	TabletAreaX = widget.NewEntry()
	TabletAreaY = widget.NewEntry()
	TabletAreaW = widget.NewEntry()
	TabletAreaH = widget.NewEntry()
	TabletAreaX.SetPlaceHolder("0.0")
	TabletAreaY.SetPlaceHolder("0.0")
	TabletAreaW.SetPlaceHolder("0.0")
	TabletAreaH.SetPlaceHolder("0.0")

	WindowSizeW = widget.NewEntry()
	WindowSizeH = widget.NewEntry()
	WindowSizeW.SetPlaceHolder("0")
	WindowSizeH.SetPlaceHolder("0")

	CustomTabletSizeW = widget.NewEntry()
	CustomTabletSizeH = widget.NewEntry()
	CustomTabletSizeW.SetPlaceHolder("0")
	CustomTabletSizeH.SetPlaceHolder("0")

	CustomTabletActiveAreaW = widget.NewEntry()
	CustomTabletActiveAreaH = widget.NewEntry()
	CustomTabletActiveAreaW.SetPlaceHolder("0")
	CustomTabletActiveAreaH.SetPlaceHolder("0")

	CustomTabletImageOffsetX = widget.NewEntry()
	CustomTabletImageOffsetY = widget.NewEntry()
	CustomTabletImageOffsetX.SetPlaceHolder("0")
	CustomTabletImageOffsetY.SetPlaceHolder("0")

	ValidateEntries()
}

func SetupSliders() {
	CursorSize = widget.NewSlider(0.01, 5.00)
	CursorSize.Step = 0.01
	CursorSizeVal = binding.NewString()
	CursorSizeVal.Set(strconv.FormatFloat(CursorSize.Value, 'f', 2, 64))
	CursorSizeDataLabel = widget.NewLabelWithData(CursorSizeVal)

	CursorTrailSize = widget.NewSlider(0.01, 5.00)
	CursorTrailSize.Step = 0.01
	CursorTrailSizeVal = binding.NewString()
	CursorTrailSizeVal.Set(strconv.FormatFloat(CursorTrailSize.Value, 'f', 2, 64))
	CursorTrailSizeDataLabel = widget.NewLabelWithData(CursorTrailSizeVal)

	FramerateLimit = widget.NewSlider(1, 400)
	FramerateLimit.Step = 1
	FramerateLimitVal = binding.NewString()
	FramerateLimitVal.Set(strconv.FormatFloat(FramerateLimit.Value, 'f', 0, 64))
	FramerateLimitDataLabel = widget.NewLabelWithData(FramerateLimitVal)

	TrailCirclesLifetime = widget.NewSlider(1, 480)
	TrailCirclesLifetime.Step = 0.01
	TrailCirclesLifetimeVal = binding.NewString()
	TrailCirclesLifetimeVal.Set(strconv.FormatFloat(TrailCirclesLifetime.Value, 'f', 2, 64))
	TrailCirclesLifetimeDataLabel = widget.NewLabelWithData(TrailCirclesLifetimeVal)

	CursorTrailDensity = widget.NewSlider(1, 480)
	CursorTrailDensity.Step = 0.01
	CursorTrailDensityVal = binding.NewString()
	CursorTrailDensityVal.Set(strconv.FormatFloat(TrailCirclesLifetime.Value, 'f', 0, 64))
	CursorTrailDensityDataLabel = widget.NewLabelWithData(CursorTrailDensityVal)

	TabletTransparency = widget.NewSlider(0, 255)
	TabletTransparency.Step = 1
	TabletTransparencyVal = binding.NewString()
	TabletTransparencyVal.Set(strconv.FormatUint(uint64(TabletTransparency.Value), 10))
	TabletTransparencyDataLabel = widget.NewLabelWithData(TabletTransparencyVal)

}

func SetupButtons() {
	SaveConfig = widget.NewButton("Save config", func() {
		SaveUserConfigToJSON()
	})
}

func ValidateEntries() {
	TabletAreaX.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("tablet area offset x is NaN")
		}
		return nil
	}
	TabletAreaY.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("tablet area offset y is NaN")
		}
		return nil
	}
	TabletAreaW.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("tablet area size x is NaN")
		}
		return nil
	}
	TabletAreaH.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("tablet area size y is NaN")
		}
		return nil
	}

	WindowSizeW.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("window size x is NaN")
		}
		return nil
	}
	WindowSizeH.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("window size y is NaN")
		}
		return nil
	}

	CustomTabletSizeW.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("custom tablet size x is NaN")
		}
		return nil
	}
	CustomTabletSizeH.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("custom tablet size y is NaN")
		}
		return nil
	}

	CustomTabletActiveAreaW.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("custom tablet area width is NaN")
		}
		return nil
	}
	CustomTabletActiveAreaH.Validator = func(s string) error {
		if !isValidUint(s) {
			return fmt.Errorf("custom tablet area height is NaN")
		}
		return nil
	}

	CustomTabletImageOffsetX.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("custom tablet image offset x is NaN")
		}
		return nil
	}
	CustomTabletImageOffsetY.Validator = func(s string) error {
		if !isValidNumber(s) {
			return fmt.Errorf("custom tablet image offset y is NaN")
		}
		return nil
	}

}

func SetupAppLayout() {
	TabletAreaContainter = container.New(layout.NewAdaptiveGridLayout(4),
		TabletAreaX,
		TabletAreaY,
		TabletAreaW,
		TabletAreaH)

	WindowSizeContainer = container.New(layout.NewAdaptiveGridLayout(2),
		WindowSizeW,
		WindowSizeH)

	CustomTabletSizeContainer = container.New(layout.NewAdaptiveGridLayout(2),
		CustomTabletSizeW,
		CustomTabletSizeH)

	CustomTabletActiveAreaContainer = container.New(layout.NewAdaptiveGridLayout(2),
		CustomTabletActiveAreaW,
		CustomTabletActiveAreaH)

	CustomTabletImageOffsetContainer = container.New(layout.NewAdaptiveGridLayout(2),
		CustomTabletImageOffsetX,
		CustomTabletImageOffsetY)

	CursorSizeContainer = container.New(layout.NewHBoxLayout(),
		CursorSizeLabel,
		CursorSizeDataLabel)

	CursorTrailSizeContainer = container.New(layout.NewHBoxLayout(),
		CursorTrailSizeLabel,
		CursorTrailSizeDataLabel)

	FramerateLimitContainer = container.New(layout.NewHBoxLayout(),
		FramerateLimitLabel,
		FramerateLimitDataLabel)

	TrailCirclesLifetimeContainer = container.New(layout.NewHBoxLayout(),
		TrailCirclesLifetimeLabel,
		TrailCirclesLifetimeDataLabel)

	CursorTrailDensityContainer = container.New(layout.NewHBoxLayout(),
		CursorTrailDensityLabel,
		CursorTrailDensityDataLabel)

	TabletTransparencyContainer = container.New(layout.NewHBoxLayout(),
		TabletTransparencyLabel,
		TabletTransparencyDataLabel)

	SetupMainLayoutContainers()
}

// This is final layout that user see when they open program
func SetupMainLayoutContainers() {
	Labels = container.New(layout.NewVBoxLayout(),
		Effect1Label,
		Effect2Label,
		CustomTabletLabel,
		EnableErrorsLabel,
		TabletAreaLabel,
		WidowSizeLabel,
		CustomTabletSizeLabel,
		CustomTabletActiveAreaLabel,
		CustomTabletImageOffsetLabel,
		layout.NewSpacer(),
		CursorSizeContainer,
		layout.NewSpacer(),
		CursorTrailSizeContainer,
		layout.NewSpacer(),
		FramerateLimitContainer,
		layout.NewSpacer(),
		TrailCirclesLifetimeContainer,
		layout.NewSpacer(),
		CursorTrailDensityContainer,
		layout.NewSpacer(),
		TabletTransparencyContainer)

	InputWidgetsContainer = container.NewVBox(
		Effect1Check,
		Effect2Check,
		CustomTabletCheck,
		EnableErrorsCheck,
		TabletAreaContainter,
		WindowSizeContainer,
		CustomTabletSizeContainer,
		CustomTabletActiveAreaContainer,
		CustomTabletImageOffsetContainer,
		layout.NewSpacer(),
		CursorSize,
		layout.NewSpacer(),
		CursorTrailSize,
		layout.NewSpacer(),
		FramerateLimit,
		layout.NewSpacer(),
		TrailCirclesLifetime,
		layout.NewSpacer(),
		CursorTrailDensity,
		layout.NewSpacer(),
		TabletTransparency)

	LabelsAndInputsContainer = container.New(layout.NewHBoxLayout(),
		Labels,
		layout.NewSpacer(),
		InputWidgetsContainer)

	MainContainter = container.New(layout.NewVBoxLayout(),
		LabelsAndInputsContainer,
		layout.NewSpacer(),
		SaveConfig)
}
