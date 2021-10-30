module github.com/yuriizinets/kyoto-starter

go 1.17

require (
	github.com/yuriizinets/kyoto v0.0.0-20211030071428-0c79c0d754b1
	github.com/yuriizinets/kyoto-uikit/twui v0.0.0
)

replace (
	github.com/yuriizinets/kyoto-uikit/twui => ./uikit/twui
)