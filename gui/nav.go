package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makeNav(setContent func(page Page)) fyne.CanvasObject {
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return PageIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := PageIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := pages[uid]
			if !ok {
				fyne.LogError("Missing panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			if unsupportedPage(t) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			} else {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			}
		},
		OnSelected: func(uid string) {
			if page, ok := pages[uid]; ok {
				if unsupportedPage(page) {
					return
				}
				setContent(page)
			}
		},
	}

	return tree
}

func unsupportedPage(p Page) bool {
	return !p.SupportWeb && fyne.CurrentDevice().IsBrowser()
}
