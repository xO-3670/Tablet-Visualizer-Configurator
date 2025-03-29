package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Settings struct {
	WindowDimensions []uint64  `json:"WindowDimensions"`
	TabletArea       []float64 `json:"TabletArea"`

	CustomTabletSize        []uint64 `json:"CustomTabletSize"`
	CustomTabletImageOffset []uint64 `json:"CustomTabletImageOffset"`
	CustomTabletActiveArea  []uint64 `json:"CustomTabletActiveArea"`

	TabletImageTransparency uint8  `json:"TabletImageTransparency"`
	CursorTrailDensity      uint16 `json:"CursorTrailDensity"`
	FramerateLimit          uint64 `json:"FramerateLimit"`

	CursorSize           float64 `json:"CursorSize"`
	CursorTrailSize      float64 `json:"CursorTrailSize"`
	TrailCirclesLifetime float64 `json:"TrailCirclesLifetime"`

	CustomTabletImage string `json:"CustomTabletImage"`
	CursorImage       string `json:"CursorImage"`
	CursorTrailImage  string `json:"CursorTrailImage"`

	CustomTablet              bool `json:"CustomTablet"`
	TrailCirclesFadingEffect  bool `json:"TrailCirclesFadingEffect"`
	TrailCirclesSpacingEffect bool `json:"TrailCirclesSpacingOutEffect"`
	EnableErrors              bool `json:"EnableErrors"`
}

var settings Settings

func LoadJSON() {
	byteValue, err := os.ReadFile("Settings.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(byteValue, &settings)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	AssignDataToWidgets()
}

func ApplyDataToSettingsStruct() {
	var x uint64
	var y uint64
	var w uint64
	var h uint64

	var xf float64
	var yf float64
	var wf float64
	var hf float64

	w, _ = strconv.ParseUint(CustomTabletActiveAreaW.Text, 10, 64)
	h, _ = strconv.ParseUint(CustomTabletActiveAreaH.Text, 10, 64)
	settings.CustomTabletActiveArea = []uint64{w, h}

	x, _ = strconv.ParseUint(CustomTabletImageOffsetX.Text, 10, 64)
	y, _ = strconv.ParseUint(CustomTabletImageOffsetY.Text, 10, 64)
	settings.CustomTabletImageOffset = []uint64{x, y}

	w, _ = strconv.ParseUint(CustomTabletSizeW.Text, 10, 64)
	h, _ = strconv.ParseUint(CustomTabletSizeH.Text, 10, 64)
	settings.CustomTabletSize = []uint64{w, h}

	xf, _ = strconv.ParseFloat(TabletAreaX.Text, 64)
	yf, _ = strconv.ParseFloat(TabletAreaY.Text, 64)
	wf, _ = strconv.ParseFloat(TabletAreaW.Text, 64)
	hf, _ = strconv.ParseFloat(TabletAreaH.Text, 64)
	settings.TabletArea = []float64{xf, yf, wf, hf}

	w, _ = strconv.ParseUint(WindowSizeW.Text, 10, 64)
	h, _ = strconv.ParseUint(WindowSizeH.Text, 10, 64)
	settings.WindowDimensions = []uint64{w, h}

	// float64
	settings.CursorSize = CursorSize.Value
	settings.CursorTrailSize = CursorTrailSize.Value
	settings.TrailCirclesLifetime = TrailCirclesLifetime.Value

	// uint8
	settings.TabletImageTransparency = uint8(TabletTransparency.Value)

	// uint16
	settings.CursorTrailDensity = uint16(CursorTrailDensity.Value)

	// uint64
	settings.FramerateLimit = uint64(FramerateLimit.Value)

	// bool
	settings.CustomTablet = CustomTabletCheck.Checked
	settings.EnableErrors = EnableErrorsCheck.Checked
	settings.TrailCirclesFadingEffect = Effect1Check.Checked
	settings.TrailCirclesSpacingEffect = Effect2Check.Checked
}

func SaveUserConfigToJSON() {
	ApplyDataToSettingsStruct()

	jsonData, err := json.Marshal(settings)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, jsonData, "", "    ")
	if err != nil {
		log.Fatalf("Error setting indent to JSON: %v", err)
	}

	err = os.WriteFile("Settings.json", out.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Error writing config to file: %v", err)
	}
}

func AssignDataToWidgets() {
	CursorSize.Value = settings.CursorSize
	CursorTrailSize.Value = settings.CursorTrailSize
	TrailCirclesLifetime.Value = settings.TrailCirclesLifetime

	FramerateLimit.Value = float64(settings.FramerateLimit)
	CursorTrailDensity.Value = float64(settings.CursorTrailDensity)

	TabletTransparency.Value = float64(settings.TabletImageTransparency)

	WindowSizeW.Text = strconv.FormatUint(settings.WindowDimensions[0], 10)
	WindowSizeH.Text = strconv.FormatUint(settings.WindowDimensions[1], 10)

	CustomTabletSizeW.Text = strconv.FormatUint(settings.CustomTabletSize[0], 10)
	CustomTabletSizeH.Text = strconv.FormatUint(settings.CustomTabletSize[1], 10)

	CustomTabletActiveAreaW.Text = strconv.FormatUint(settings.CustomTabletActiveArea[0], 10)
	CustomTabletActiveAreaH.Text = strconv.FormatUint(settings.CustomTabletActiveArea[1], 10)

	CustomTabletImageOffsetX.Text = strconv.FormatUint(settings.CustomTabletImageOffset[0], 10)
	CustomTabletImageOffsetY.Text = strconv.FormatUint(settings.CustomTabletImageOffset[1], 10)

	TabletAreaX.Text = strconv.FormatFloat(settings.TabletArea[0], 'f', 2, 64)
	TabletAreaY.Text = strconv.FormatFloat(settings.TabletArea[1], 'f', 2, 64)
	TabletAreaW.Text = strconv.FormatFloat(settings.TabletArea[2], 'f', 2, 64)
	TabletAreaH.Text = strconv.FormatFloat(settings.TabletArea[3], 'f', 2, 64)

	Effect1Check.Checked = settings.TrailCirclesFadingEffect
	Effect2Check.Checked = settings.TrailCirclesSpacingEffect
	CustomTabletCheck.Checked = settings.CustomTablet
	EnableErrorsCheck.Checked = settings.EnableErrors
}
