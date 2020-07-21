// Code generated by 'goexports github.com/topxeq/imagetk'. DO NOT EDIT.

// +build go1.14,!go1.15

package main

import (
	"github.com/topxeq/imagetk"
	"reflect"
)

func init() {
	GotxSymbols["github.com/topxeq/imagetk"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Bicubic":           reflect.ValueOf(imagetk.Bicubic),
		"Bilinear":          reflect.ValueOf(imagetk.Bilinear),
		"Lanczos2":          reflect.ValueOf(imagetk.Lanczos2),
		"Lanczos3":          reflect.ValueOf(imagetk.Lanczos3),
		"MitchellNetravali": reflect.ValueOf(imagetk.MitchellNetravali),
		"NearestNeighbor":   reflect.ValueOf(imagetk.NearestNeighbor),
		"NewImageTK":        reflect.ValueOf(imagetk.NewImageTK),

		// type definitions
		"ImageTK":               reflect.ValueOf((*imagetk.ImageTK)(nil)),
		"InterpolationFunction": reflect.ValueOf((*imagetk.InterpolationFunction)(nil)),
	}
}