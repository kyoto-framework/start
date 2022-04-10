module github.com/kyoto-framework/start

go 1.18

require (
	github.com/kyoto-framework/i18n v0.0.0-20220407073653-0a4fc6b7f130 // indirect
	github.com/kyoto-framework/kyoto v0.2.1-0.20220410124033-0d16c397f65f // indirect
	github.com/kyoto-framework/scheduler v0.0.0-20220208110634-a2f0babca15f // indirect
	github.com/kyoto-framework/zen v0.0.0-20220408163909-add65cb11c5d // indirect
	github.com/kyoto-framework/uikit/twui v0.0.0
	golang.org/x/exp v0.0.0-20220328175248-053ad81199eb // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/kyoto-framework/uikit/twui => ./uikit/twui
)