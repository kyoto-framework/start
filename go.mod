module github.com/kyoto-framework/starter

go 1.17

require (
	github.com/kyoto-framework/kyoto v0.0.0-20211030071428-0c79c0d754b1
	github.com/kyoto-framework/uikit/twui v0.0.0
)

replace (
	github.com/kyoto-framework/uikit/twui => ./uikit/twui
)