// This file was auto-generated by Fern from our API Definition.

package api

var Environments = struct {
	Production struct {
		Auth   string
		Plants string
	}
	Staging struct {
		Auth   string
		Plants string
	}
}{
	Production: struct {
		Auth   string
		Plants string
	}{
		Auth:   "https://auth.yoursite.com",
		Plants: "https://plants.yoursite.com",
	},
	Staging: struct {
		Auth   string
		Plants string
	}{
		Auth:   "https://auth.staging.yoursite.com",
		Plants: "https://plants.staging.yoursite.com",
	},
}