package main

import "github.com/kyoto-framework/uikit/twui"

func UINavbar() *twui.AppUINavNavbar {
	return &twui.AppUINavNavbar{
		Logo: `<img src="/static/img/icons/kyoto.svg" class="h-8 w-8 scale-150" />`,
		Links: []twui.AppUINavNavbarLink{
			{Text: "Kyoto", Href: "https://github.com/kyoto-framework/kyoto"},
			{Text: "UIKit", Href: "https://github.com/kyoto-framework/uikit"},
			{Text: "Charts", Href: "https://github.com/kyoto-framework/kyoto-charts"},
			{Text: "Starter", Href: "https://github.com/kyoto-framework/starter"},
		},
		Profile: twui.AppUINavNavbarProfile{
			Enabled: true,
			Avatar: `
					<svg class="w-6 h-6 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path></svg>
				`,
			Links: []twui.AppUINavNavbarLink{
				{Text: "GitHub", Href: "https://github.com/kyoto-framework/kyoto/discussions/40"},
				{Text: "Telegram", Href: "https://t.me/yuriizinets"},
				{Text: "Email", Href: "mailto:yurii.zinets@icloud.com"},
			},
		},
	}
}
