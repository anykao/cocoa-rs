package main

type Foo struct {
	Arch        string `json:"arch"`
	Definitions struct {
		Aliases struct {
			NSDockWindowLevel                     string `json:"NSDockWindowLevel"`
			NSDragOperationAll                    string `json:"NSDragOperationAll"`
			NSDraggingItemEnumerationConcurrent   string `json:"NSDraggingItemEnumerationConcurrent"`
			NSEventDurationForever                string `json:"NSEventDurationForever"`
			NSFileHandlingPanelCancelButton       string `json:"NSFileHandlingPanelCancelButton"`
			NSFileHandlingPanelOKButton           string `json:"NSFileHandlingPanelOKButton"`
			NSFloatingWindowLevel                 string `json:"NSFloatingWindowLevel"`
			NSGestureRecognizerStateRecognized    string `json:"NSGestureRecognizerStateRecognized"`
			NSImageRepRegistryChangedNotification string `json:"NSImageRepRegistryChangedNotification"`
			NSMainMenuWindowLevel                 string `json:"NSMainMenuWindowLevel"`
			NSModalPanelWindowLevel               string `json:"NSModalPanelWindowLevel"`
			NSNormalWindowLevel                   string `json:"NSNormalWindowLevel"`
			NSPopUpMenuWindowLevel                string `json:"NSPopUpMenuWindowLevel"`
			NSScreenSaverWindowLevel              string `json:"NSScreenSaverWindowLevel"`
			NSStatusWindowLevel                   string `json:"NSStatusWindowLevel"`
			NSSubmenuWindowLevel                  string `json:"NSSubmenuWindowLevel"`
			NSTickMarkLeft                        string `json:"NSTickMarkLeft"`
			NSTickMarkRight                       string `json:"NSTickMarkRight"`
			NSTornOffMenuWindowLevel              string `json:"NSTornOffMenuWindowLevel"`
		} `json:"aliases"`
		CalledDefinitions struct{} `json:"called_definitions"`
		Cftypes           struct{} `json:"cftypes"`
		Classes           struct {
			CIColor struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"CIColor"`
			CIImage struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"CIImage"`
			NSATSTypesetter struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSATSTypesetter"`
			NSAccessibilityElement struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSAccessibilityElement"`
			NSActionCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSActionCell"`
			NSAffineTransform struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"NSAffineTransform"`
			NSAlert struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Variadic   bool   `json:"variadic"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSAlert"`
			NSAnimation struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSAnimation"`
			NSAnimationContext struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSAnimationContext"`
			NSAppearance struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSAppearance"`
			NSAppleScript struct {
				Methods []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
			} `json:"NSAppleScript"`
			NSApplication struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSApplication"`
			NSArrayController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSArrayController"`
			NSAttributedString struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
			} `json:"NSAttributedString"`
			NSBezierPath struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSBezierPath"`
			NSBitmapImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSBitmapImageRep"`
			NSBox struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSBox"`
			NSBrowser struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSBrowser"`
			NSBrowserCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSBrowserCell"`
			NSBundle struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"NSBundle"`
			NSButton struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSButton"`
			NSButtonCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSButtonCell"`
			NSCIImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSCIImageRep"`
			NSCachedImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSCachedImageRep"`
			NSCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSCell"`
			NSClickGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSClickGestureRecognizer"`
			NSClipView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSClipView"`
			NSCoder struct {
				Methods []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"NSCoder"`
			NSCollectionView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSCollectionView"`
			NSCollectionViewItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSCollectionViewItem"`
			NSColor struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSColor"`
			NSColorList struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSColorList"`
			NSColorPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSColorPanel"`
			NSColorPicker struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSColorPicker"`
			NSColorSpace struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSColorSpace"`
			NSColorWell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSColorWell"`
			NSComboBox struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSComboBox"`
			NSComboBoxCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSComboBoxCell"`
			NSControl struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSControl"`
			NSController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSController"`
			NSCursor struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSCursor"`
			NSCustomImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSCustomImageRep"`
			NSDatePicker struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDatePicker"`
			NSDatePickerCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDatePickerCell"`
			NSDictionaryController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDictionaryController"`
			NSDockTile struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDockTile"`
			NSDocument struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSDocument"`
			NSDocumentController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSDocumentController"`
			NSDraggingImageComponent struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDraggingImageComponent"`
			NSDraggingItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDraggingItem"`
			NSDraggingSession struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSDraggingSession"`
			NSDrawer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSDrawer"`
			NSEPSImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSEPSImageRep"`
			NSEvent struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSEvent"`
			NSFileWrapper struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
			} `json:"NSFileWrapper"`
			NSFont struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSFont"`
			NSFontCollection struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSFontCollection"`
			NSFontDescriptor struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSFontDescriptor"`
			NSFontManager struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSFontManager"`
			NSFontPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSFontPanel"`
			NSForm struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSForm"`
			NSFormCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSFormCell"`
			NSGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSGestureRecognizer"`
			NSGlyphGenerator struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSGlyphGenerator"`
			NSGlyphInfo struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSGlyphInfo"`
			NSGradient struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Variadic   bool   `json:"variadic"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSGradient"`
			NSGraphicsContext struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSGraphicsContext"`
			NSHelpManager struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSHelpManager"`
			NSImage struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSImage"`
			NSImageCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSImageCell"`
			NSImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSImageRep"`
			NSImageView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSImageView"`
			NSInputManager struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSInputManager"`
			NSInputServer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSInputServer"`
			NSItemProvider struct {
				Methods []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
			} `json:"NSItemProvider"`
			NSLayoutConstraint struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSLayoutConstraint"`
			NSLayoutManager struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSLayoutManager"`
			NSLevelIndicator struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSLevelIndicator"`
			NSLevelIndicatorCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSLevelIndicatorCell"`
			NSMagnificationGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSMagnificationGestureRecognizer"`
			NSMatrix struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSMatrix"`
			NSMediaLibraryBrowserController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSMediaLibraryBrowserController"`
			NSMenu struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSMenu"`
			NSMenuItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSMenuItem"`
			NSMenuItemCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSMenuItemCell"`
			NSMovie struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSMovie"`
			NSMutableAttributedString struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
			} `json:"NSMutableAttributedString"`
			NSMutableFontCollection struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSMutableFontCollection"`
			NSMutableParagraphStyle struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSMutableParagraphStyle"`
			NSNib struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSNib"`
			NSObject struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
			} `json:"NSObject"`
			NSObjectController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSObjectController"`
			NSOpenGLContext struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSOpenGLContext"`
			NSOpenGLLayer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSOpenGLLayer"`
			NSOpenGLPixelBuffer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSOpenGLPixelBuffer"`
			NSOpenGLPixelFormat struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSOpenGLPixelFormat"`
			NSOpenGLView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSOpenGLView"`
			NSOpenPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSOpenPanel"`
			NSOutlineView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSOutlineView"`
			NSPDFImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPDFImageRep"`
			NSPDFInfo struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPDFInfo"`
			NSPDFPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPDFPanel"`
			NSPICTImageRep struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPICTImageRep"`
			NSPageController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPageController"`
			NSPageLayout struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPageLayout"`
			NSPanGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPanGestureRecognizer"`
			NSPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPanel"`
			NSParagraphStyle struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSParagraphStyle"`
			NSPasteboard struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPasteboard"`
			NSPasteboardItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPasteboardItem"`
			NSPathCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPathCell"`
			NSPathComponentCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPathComponentCell"`
			NSPathControl struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPathControl"`
			NSPathControlItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPathControlItem"`
			NSPersistentDocument struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPersistentDocument"`
			NSPopUpButton struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPopUpButton"`
			NSPopUpButtonCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPopUpButtonCell"`
			NSPopover struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPopover"`
			NSPredicateEditor struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPredicateEditor"`
			NSPredicateEditorRowTemplate struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPredicateEditorRowTemplate"`
			NSPressGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPressGestureRecognizer"`
			NSPrintInfo struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPrintInfo"`
			NSPrintOperation struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPrintOperation"`
			NSPrintPanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSPrintPanel"`
			NSPrinter struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSPrinter"`
			NSProgressIndicator struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSProgressIndicator"`
			NSResponder struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSResponder"`
			NSRotationGestureRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSRotationGestureRecognizer"`
			NSRuleEditor struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSRuleEditor"`
			NSRulerMarker struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSRulerMarker"`
			NSRulerView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSRulerView"`
			NSRunningApplication struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSRunningApplication"`
			NSSavePanel struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSavePanel"`
			NSScreen struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSScreen"`
			NSScrollView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSScrollView"`
			NSScroller struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSScroller"`
			NSSearchField struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSearchField"`
			NSSearchFieldCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSearchFieldCell"`
			NSSecureTextField struct {
				Categories []interface{} `json:"categories"`
				Methods    []interface{} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSSecureTextField"`
			NSSecureTextFieldCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSecureTextFieldCell"`
			NSSegmentedCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSegmentedCell"`
			NSSegmentedControl struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSegmentedControl"`
			NSShadow struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSShadow"`
			NSSharingService struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSharingService"`
			NSSharingServicePicker struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSharingServicePicker"`
			NSSlider struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSSlider"`
			NSSliderCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSliderCell"`
			NSSound struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSSound"`
			NSSpeechRecognizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSpeechRecognizer"`
			NSSpeechSynthesizer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSpeechSynthesizer"`
			NSSpellChecker struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSpellChecker"`
			NSSplitView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSSplitView"`
			NSSplitViewController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSSplitViewController"`
			NSSplitViewItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSSplitViewItem"`
			NSStackView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStackView"`
			NSStatusBar struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStatusBar"`
			NSStatusBarButton struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStatusBarButton"`
			NSStatusItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStatusItem"`
			NSStepper struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSStepper"`
			NSStepperCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStepperCell"`
			NSStoryboard struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSStoryboard"`
			NSStoryboardSegue struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSStoryboardSegue"`
			NSString struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
			} `json:"NSString"`
			NSTabView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTabView"`
			NSTabViewController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTabViewController"`
			NSTabViewItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTabViewItem"`
			NSTableCellView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTableCellView"`
			NSTableColumn struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTableColumn"`
			NSTableHeaderCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []interface{} `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSTableHeaderCell"`
			NSTableHeaderView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTableHeaderView"`
			NSTableRowView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTableRowView"`
			NSTableView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTableView"`
			NSText struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSText"`
			NSTextAlternatives struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextAlternatives"`
			NSTextAttachment struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextAttachment"`
			NSTextAttachmentCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []interface{} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
				Super      string        `json:"super"`
			} `json:"NSTextAttachmentCell"`
			NSTextBlock struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextBlock"`
			NSTextContainer struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextContainer"`
			NSTextField struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextField"`
			NSTextFieldCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextFieldCell"`
			NSTextFinder struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextFinder"`
			NSTextInputContext struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextInputContext"`
			NSTextList struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextList"`
			NSTextStorage struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextStorage"`
			NSTextTab struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextTab"`
			NSTextTable struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextTable"`
			NSTextTableBlock struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTextTableBlock"`
			NSTextView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTextView"`
			NSTitlebarAccessoryViewController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTitlebarAccessoryViewController"`
			NSTokenField struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTokenField"`
			NSTokenFieldCell struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTokenFieldCell"`
			NSToolbar struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSToolbar"`
			NSToolbarItem struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSToolbarItem"`
			NSToolbarItemGroup struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSToolbarItemGroup"`
			NSTouch struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTouch"`
			NSTrackingArea struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSTrackingArea"`
			NSTreeController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTreeController"`
			NSTreeNode struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTreeNode"`
			NSTypesetter struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSTypesetter"`
			Nsurl struct {
				Methods []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string        `json:"name"`
				Properties []interface{} `json:"properties"`
				Protocols  []string      `json:"protocols"`
			} `json:"NSURL"`
			NSUserDefaultsController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSUserDefaultsController"`
			NSView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSView"`
			NSViewAnimation struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSViewAnimation"`
			NSViewController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSViewController"`
			NSVisualEffectView struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSVisualEffectView"`
			NSWindow struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSWindow"`
			NSWindowController struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []string `json:"protocols"`
				Super     string   `json:"super"`
			} `json:"NSWindowController"`
			NSWorkspace struct {
				Categories []interface{} `json:"categories"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Name       string `json:"name"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
				Protocols []interface{} `json:"protocols"`
				Super     string        `json:"super"`
			} `json:"NSWorkspace"`
		} `json:"classes"`
		Enum struct {
			NS16BitBigEndianBitmapFormat                          int64 `json:"NS16BitBigEndianBitmapFormat"`
			NS16BitLittleEndianBitmapFormat                       int64 `json:"NS16BitLittleEndianBitmapFormat"`
			NS32BitBigEndianBitmapFormat                          int64 `json:"NS32BitBigEndianBitmapFormat"`
			NS32BitLittleEndianBitmapFormat                       int64 `json:"NS32BitLittleEndianBitmapFormat"`
			NSAWTEventType                                        int64 `json:"NSAWTEventType"`
			NSAboveBottom                                         int64 `json:"NSAboveBottom"`
			NSAboveTop                                            int64 `json:"NSAboveTop"`
			NSAcceleratorButton                                   int64 `json:"NSAcceleratorButton"`
			NSAccessibilityOrientationHorizontal                  int64 `json:"NSAccessibilityOrientationHorizontal"`
			NSAccessibilityOrientationUnknown                     int64 `json:"NSAccessibilityOrientationUnknown"`
			NSAccessibilityOrientationVertical                    int64 `json:"NSAccessibilityOrientationVertical"`
			NSAccessibilityPriorityHigh                           int64 `json:"NSAccessibilityPriorityHigh"`
			NSAccessibilityPriorityLow                            int64 `json:"NSAccessibilityPriorityLow"`
			NSAccessibilityPriorityMedium                         int64 `json:"NSAccessibilityPriorityMedium"`
			NSAccessibilityRulerMarkerTypeIndentFirstLine         int64 `json:"NSAccessibilityRulerMarkerTypeIndentFirstLine"`
			NSAccessibilityRulerMarkerTypeIndentHead              int64 `json:"NSAccessibilityRulerMarkerTypeIndentHead"`
			NSAccessibilityRulerMarkerTypeIndentTail              int64 `json:"NSAccessibilityRulerMarkerTypeIndentTail"`
			NSAccessibilityRulerMarkerTypeTabStopCenter           int64 `json:"NSAccessibilityRulerMarkerTypeTabStopCenter"`
			NSAccessibilityRulerMarkerTypeTabStopDecimal          int64 `json:"NSAccessibilityRulerMarkerTypeTabStopDecimal"`
			NSAccessibilityRulerMarkerTypeTabStopLeft             int64 `json:"NSAccessibilityRulerMarkerTypeTabStopLeft"`
			NSAccessibilityRulerMarkerTypeTabStopRight            int64 `json:"NSAccessibilityRulerMarkerTypeTabStopRight"`
			NSAccessibilityRulerMarkerTypeUnknown                 int64 `json:"NSAccessibilityRulerMarkerTypeUnknown"`
			NSAccessibilitySortDirectionAscending                 int64 `json:"NSAccessibilitySortDirectionAscending"`
			NSAccessibilitySortDirectionDescending                int64 `json:"NSAccessibilitySortDirectionDescending"`
			NSAccessibilitySortDirectionUnknown                   int64 `json:"NSAccessibilitySortDirectionUnknown"`
			NSAccessibilityUnitsCentimeters                       int64 `json:"NSAccessibilityUnitsCentimeters"`
			NSAccessibilityUnitsInches                            int64 `json:"NSAccessibilityUnitsInches"`
			NSAccessibilityUnitsPicas                             int64 `json:"NSAccessibilityUnitsPicas"`
			NSAccessibilityUnitsPoints                            int64 `json:"NSAccessibilityUnitsPoints"`
			NSAccessibilityUnitsUnknown                           int64 `json:"NSAccessibilityUnitsUnknown"`
			NSAddTraitFontAction                                  int64 `json:"NSAddTraitFontAction"`
			NSAdobeCNS1CharacterCollection                        int64 `json:"NSAdobeCNS1CharacterCollection"`
			NSAdobeGB1CharacterCollection                         int64 `json:"NSAdobeGB1CharacterCollection"`
			NSAdobeJapan1CharacterCollection                      int64 `json:"NSAdobeJapan1CharacterCollection"`
			NSAdobeJapan2CharacterCollection                      int64 `json:"NSAdobeJapan2CharacterCollection"`
			NSAdobeKorea1CharacterCollection                      int64 `json:"NSAdobeKorea1CharacterCollection"`
			NSAlertAlternateReturn                                int64 `json:"NSAlertAlternateReturn"`
			NSAlertDefaultReturn                                  int64 `json:"NSAlertDefaultReturn"`
			NSAlertErrorReturn                                    int64 `json:"NSAlertErrorReturn"`
			NSAlertFirstButtonReturn                              int64 `json:"NSAlertFirstButtonReturn"`
			NSAlertOtherReturn                                    int64 `json:"NSAlertOtherReturn"`
			NSAlertSecondButtonReturn                             int64 `json:"NSAlertSecondButtonReturn"`
			NSAlertThirdButtonReturn                              int64 `json:"NSAlertThirdButtonReturn"`
			NSAllScrollerParts                                    int64 `json:"NSAllScrollerParts"`
			NSAlphaFirstBitmapFormat                              int64 `json:"NSAlphaFirstBitmapFormat"`
			NSAlphaNonpremultipliedBitmapFormat                   int64 `json:"NSAlphaNonpremultipliedBitmapFormat"`
			NSAlphaShiftKeyMask                                   int64 `json:"NSAlphaShiftKeyMask"`
			NSAlternateKeyMask                                    int64 `json:"NSAlternateKeyMask"`
			NSAnimationBlocking                                   int64 `json:"NSAnimationBlocking"`
			NSAnimationEaseIn                                     int64 `json:"NSAnimationEaseIn"`
			NSAnimationEaseInOut                                  int64 `json:"NSAnimationEaseInOut"`
			NSAnimationEaseOut                                    int64 `json:"NSAnimationEaseOut"`
			NSAnimationEffectDisappearingItemDefault              int64 `json:"NSAnimationEffectDisappearingItemDefault"`
			NSAnimationEffectPoof                                 int64 `json:"NSAnimationEffectPoof"`
			NSAnimationLinear                                     int64 `json:"NSAnimationLinear"`
			NSAnimationNonblocking                                int64 `json:"NSAnimationNonblocking"`
			NSAnimationNonblockingThreaded                        int64 `json:"NSAnimationNonblockingThreaded"`
			NSAnyEventMask                                        int64 `json:"NSAnyEventMask"`
			NSAnyType                                             int64 `json:"NSAnyType"`
			NSAppKitDefined                                       int64 `json:"NSAppKitDefined"`
			NSAppKitDefinedMask                                   int64 `json:"NSAppKitDefinedMask"`
			NSAppKitVersionNumber10_0                             int64 `json:"NSAppKitVersionNumber10_0"`
			NSAppKitVersionNumber10_1                             int64 `json:"NSAppKitVersionNumber10_1"`
			NSAppKitVersionNumber10_10                            int64 `json:"NSAppKitVersionNumber10_10"`
			NSAppKitVersionNumber10_2                             int64 `json:"NSAppKitVersionNumber10_2"`
			NSAppKitVersionNumber10_3                             int64 `json:"NSAppKitVersionNumber10_3"`
			NSAppKitVersionNumber10_4                             int64 `json:"NSAppKitVersionNumber10_4"`
			NSAppKitVersionNumber10_5                             int64 `json:"NSAppKitVersionNumber10_5"`
			NSAppKitVersionNumber10_6                             int64 `json:"NSAppKitVersionNumber10_6"`
			NSAppKitVersionNumber10_7                             int64 `json:"NSAppKitVersionNumber10_7"`
			NSAppKitVersionNumber10_8                             int64 `json:"NSAppKitVersionNumber10_8"`
			NSAppKitVersionNumber10_9                             int64 `json:"NSAppKitVersionNumber10_9"`
			NSApplicationActivateAllWindows                       int64 `json:"NSApplicationActivateAllWindows"`
			NSApplicationActivateIgnoringOtherApps                int64 `json:"NSApplicationActivateIgnoringOtherApps"`
			NSApplicationActivatedEventType                       int64 `json:"NSApplicationActivatedEventType"`
			NSApplicationActivationPolicyAccessory                int64 `json:"NSApplicationActivationPolicyAccessory"`
			NSApplicationActivationPolicyProhibited               int64 `json:"NSApplicationActivationPolicyProhibited"`
			NSApplicationActivationPolicyRegular                  int64 `json:"NSApplicationActivationPolicyRegular"`
			NSApplicationDeactivatedEventType                     int64 `json:"NSApplicationDeactivatedEventType"`
			NSApplicationDefined                                  int64 `json:"NSApplicationDefined"`
			NSApplicationDefinedMask                              int64 `json:"NSApplicationDefinedMask"`
			NSApplicationDelegateReplyCancel                      int64 `json:"NSApplicationDelegateReplyCancel"`
			NSApplicationDelegateReplyFailure                     int64 `json:"NSApplicationDelegateReplyFailure"`
			NSApplicationDelegateReplySuccess                     int64 `json:"NSApplicationDelegateReplySuccess"`
			NSApplicationOcclusionStateVisible                    int64 `json:"NSApplicationOcclusionStateVisible"`
			NSApplicationPresentationAutoHideDock                 int64 `json:"NSApplicationPresentationAutoHideDock"`
			NSApplicationPresentationAutoHideMenuBar              int64 `json:"NSApplicationPresentationAutoHideMenuBar"`
			NSApplicationPresentationAutoHideToolbar              int64 `json:"NSApplicationPresentationAutoHideToolbar"`
			NSApplicationPresentationDefault                      int64 `json:"NSApplicationPresentationDefault"`
			NSApplicationPresentationDisableAppleMenu             int64 `json:"NSApplicationPresentationDisableAppleMenu"`
			NSApplicationPresentationDisableForceQuit             int64 `json:"NSApplicationPresentationDisableForceQuit"`
			NSApplicationPresentationDisableHideApplication       int64 `json:"NSApplicationPresentationDisableHideApplication"`
			NSApplicationPresentationDisableMenuBarTransparency   int64 `json:"NSApplicationPresentationDisableMenuBarTransparency"`
			NSApplicationPresentationDisableProcessSwitching      int64 `json:"NSApplicationPresentationDisableProcessSwitching"`
			NSApplicationPresentationDisableSessionTermination    int64 `json:"NSApplicationPresentationDisableSessionTermination"`
			NSApplicationPresentationFullScreen                   int64 `json:"NSApplicationPresentationFullScreen"`
			NSApplicationPresentationHideDock                     int64 `json:"NSApplicationPresentationHideDock"`
			NSApplicationPresentationHideMenuBar                  int64 `json:"NSApplicationPresentationHideMenuBar"`
			NSAscendingPageOrder                                  int64 `json:"NSAscendingPageOrder"`
			NSAtBottom                                            int64 `json:"NSAtBottom"`
			NSAtTop                                               int64 `json:"NSAtTop"`
			NSAttachmentCharacter                                 int64 `json:"NSAttachmentCharacter"`
			NSAutoPagination                                      int64 `json:"NSAutoPagination"`
			NSAutosaveAsOperation                                 int64 `json:"NSAutosaveAsOperation"`
			NSAutosaveElsewhereOperation                          int64 `json:"NSAutosaveElsewhereOperation"`
			NSAutosaveInPlaceOperation                            int64 `json:"NSAutosaveInPlaceOperation"`
			NSAutosaveOperation                                   int64 `json:"NSAutosaveOperation"`
			NSBMPFileType                                         int64 `json:"NSBMPFileType"`
			NSBackTabCharacter                                    int64 `json:"NSBackTabCharacter"`
			NSBackgroundStyleDark                                 int64 `json:"NSBackgroundStyleDark"`
			NSBackgroundStyleLight                                int64 `json:"NSBackgroundStyleLight"`
			NSBackgroundStyleLowered                              int64 `json:"NSBackgroundStyleLowered"`
			NSBackgroundStyleRaised                               int64 `json:"NSBackgroundStyleRaised"`
			NSBackgroundTab                                       int64 `json:"NSBackgroundTab"`
			NSBackingStoreBuffered                                int64 `json:"NSBackingStoreBuffered"`
			NSBackingStoreNonretained                             int64 `json:"NSBackingStoreNonretained"`
			NSBackingStoreRetained                                int64 `json:"NSBackingStoreRetained"`
			NSBackspaceCharacter                                  int64 `json:"NSBackspaceCharacter"`
			NSBacktabTextMovement                                 int64 `json:"NSBacktabTextMovement"`
			NSBeginFunctionKey                                    int64 `json:"NSBeginFunctionKey"`
			NSBelowBottom                                         int64 `json:"NSBelowBottom"`
			NSBelowTop                                            int64 `json:"NSBelowTop"`
			NSBevelLineJoinStyle                                  int64 `json:"NSBevelLineJoinStyle"`
			NSBezelBorder                                         int64 `json:"NSBezelBorder"`
			NSBlueControlTint                                     int64 `json:"NSBlueControlTint"`
			NSBoldFontMask                                        int64 `json:"NSBoldFontMask"`
			NSBorderlessWindowMask                                int64 `json:"NSBorderlessWindowMask"`
			NSBottomTabsBezelBorder                               int64 `json:"NSBottomTabsBezelBorder"`
			NSBoxCustom                                           int64 `json:"NSBoxCustom"`
			NSBoxOldStyle                                         int64 `json:"NSBoxOldStyle"`
			NSBoxPrimary                                          int64 `json:"NSBoxPrimary"`
			NSBoxSecondary                                        int64 `json:"NSBoxSecondary"`
			NSBoxSeparator                                        int64 `json:"NSBoxSeparator"`
			NSBreakFunctionKey                                    int64 `json:"NSBreakFunctionKey"`
			NSBrowserAutoColumnResizing                           int64 `json:"NSBrowserAutoColumnResizing"`
			NSBrowserDropAbove                                    int64 `json:"NSBrowserDropAbove"`
			NSBrowserDropOn                                       int64 `json:"NSBrowserDropOn"`
			NSBrowserNoColumnResizing                             int64 `json:"NSBrowserNoColumnResizing"`
			NSBrowserUserColumnResizing                           int64 `json:"NSBrowserUserColumnResizing"`
			NSButtLineCapStyle                                    int64 `json:"NSButtLineCapStyle"`
			NSCMYKColorSpaceModel                                 int64 `json:"NSCMYKColorSpaceModel"`
			NSCMYKModeColorPanel                                  int64 `json:"NSCMYKModeColorPanel"`
			NSCancelButton                                        int64 `json:"NSCancelButton"`
			NSCancelTextMovement                                  int64 `json:"NSCancelTextMovement"`
			NSCarriageReturnCharacter                             int64 `json:"NSCarriageReturnCharacter"`
			NSCellAllowsMixedState                                int64 `json:"NSCellAllowsMixedState"`
			NSCellChangesContents                                 int64 `json:"NSCellChangesContents"`
			NSCellDisabled                                        int64 `json:"NSCellDisabled"`
			NSCellEditable                                        int64 `json:"NSCellEditable"`
			NSCellHasImageHorizontal                              int64 `json:"NSCellHasImageHorizontal"`
			NSCellHasImageOnLeftOrBottom                          int64 `json:"NSCellHasImageOnLeftOrBottom"`
			NSCellHasOverlappingImage                             int64 `json:"NSCellHasOverlappingImage"`
			NSCellHighlighted                                     int64 `json:"NSCellHighlighted"`
			NSCellHitContentArea                                  int64 `json:"NSCellHitContentArea"`
			NSCellHitEditableTextArea                             int64 `json:"NSCellHitEditableTextArea"`
			NSCellHitNone                                         int64 `json:"NSCellHitNone"`
			NSCellHitTrackableArea                                int64 `json:"NSCellHitTrackableArea"`
			NSCellIsBordered                                      int64 `json:"NSCellIsBordered"`
			NSCellIsInsetButton                                   int64 `json:"NSCellIsInsetButton"`
			NSCellLightsByBackground                              int64 `json:"NSCellLightsByBackground"`
			NSCellLightsByContents                                int64 `json:"NSCellLightsByContents"`
			NSCellLightsByGray                                    int64 `json:"NSCellLightsByGray"`
			NSCellState                                           int64 `json:"NSCellState"`
			NSCenterTabStopType                                   int64 `json:"NSCenterTabStopType"`
			NSCenterTextAlignment                                 int64 `json:"NSCenterTextAlignment"`
			NSChangeAutosaved                                     int64 `json:"NSChangeAutosaved"`
			NSChangeBackgroundCell                                int64 `json:"NSChangeBackgroundCell"`
			NSChangeBackgroundCellMask                            int64 `json:"NSChangeBackgroundCellMask"`
			NSChangeCleared                                       int64 `json:"NSChangeCleared"`
			NSChangeDiscardable                                   int64 `json:"NSChangeDiscardable"`
			NSChangeDone                                          int64 `json:"NSChangeDone"`
			NSChangeGrayCell                                      int64 `json:"NSChangeGrayCell"`
			NSChangeGrayCellMask                                  int64 `json:"NSChangeGrayCellMask"`
			NSChangeReadOtherContents                             int64 `json:"NSChangeReadOtherContents"`
			NSChangeRedone                                        int64 `json:"NSChangeRedone"`
			NSChangeUndone                                        int64 `json:"NSChangeUndone"`
			NSCircularBezelStyle                                  int64 `json:"NSCircularBezelStyle"`
			NSCircularSlider                                      int64 `json:"NSCircularSlider"`
			NSClearControlTint                                    int64 `json:"NSClearControlTint"`
			NSClearDisplayFunctionKey                             int64 `json:"NSClearDisplayFunctionKey"`
			NSClearLineFunctionKey                                int64 `json:"NSClearLineFunctionKey"`
			NSClipPagination                                      int64 `json:"NSClipPagination"`
			NSClockAndCalendarDatePickerStyle                     int64 `json:"NSClockAndCalendarDatePickerStyle"`
			NSClosableWindowMask                                  int64 `json:"NSClosableWindowMask"`
			NSClosePathBezierPathElement                          int64 `json:"NSClosePathBezierPathElement"`
			NSCollectionViewDropBefore                            int64 `json:"NSCollectionViewDropBefore"`
			NSCollectionViewDropOn                                int64 `json:"NSCollectionViewDropOn"`
			NSColorListModeColorPanel                             int64 `json:"NSColorListModeColorPanel"`
			NSColorPanelAllModesMask                              int64 `json:"NSColorPanelAllModesMask"`
			NSColorPanelCMYKModeMask                              int64 `json:"NSColorPanelCMYKModeMask"`
			NSColorPanelColorListModeMask                         int64 `json:"NSColorPanelColorListModeMask"`
			NSColorPanelCrayonModeMask                            int64 `json:"NSColorPanelCrayonModeMask"`
			NSColorPanelCustomPaletteModeMask                     int64 `json:"NSColorPanelCustomPaletteModeMask"`
			NSColorPanelGrayModeMask                              int64 `json:"NSColorPanelGrayModeMask"`
			NSColorPanelHSBModeMask                               int64 `json:"NSColorPanelHSBModeMask"`
			NSColorPanelRGBModeMask                               int64 `json:"NSColorPanelRGBModeMask"`
			NSColorPanelWheelModeMask                             int64 `json:"NSColorPanelWheelModeMask"`
			NSColorRenderingIntentAbsoluteColorimetric            int64 `json:"NSColorRenderingIntentAbsoluteColorimetric"`
			NSColorRenderingIntentDefault                         int64 `json:"NSColorRenderingIntentDefault"`
			NSColorRenderingIntentPerceptual                      int64 `json:"NSColorRenderingIntentPerceptual"`
			NSColorRenderingIntentRelativeColorimetric            int64 `json:"NSColorRenderingIntentRelativeColorimetric"`
			NSColorRenderingIntentSaturation                      int64 `json:"NSColorRenderingIntentSaturation"`
			NSCommandKeyMask                                      int64 `json:"NSCommandKeyMask"`
			NSCompositeClear                                      int64 `json:"NSCompositeClear"`
			NSCompositeColor                                      int64 `json:"NSCompositeColor"`
			NSCompositeColorBurn                                  int64 `json:"NSCompositeColorBurn"`
			NSCompositeColorDodge                                 int64 `json:"NSCompositeColorDodge"`
			NSCompositeCopy                                       int64 `json:"NSCompositeCopy"`
			NSCompositeDarken                                     int64 `json:"NSCompositeDarken"`
			NSCompositeDestinationAtop                            int64 `json:"NSCompositeDestinationAtop"`
			NSCompositeDestinationIn                              int64 `json:"NSCompositeDestinationIn"`
			NSCompositeDestinationOut                             int64 `json:"NSCompositeDestinationOut"`
			NSCompositeDestinationOver                            int64 `json:"NSCompositeDestinationOver"`
			NSCompositeDifference                                 int64 `json:"NSCompositeDifference"`
			NSCompositeExclusion                                  int64 `json:"NSCompositeExclusion"`
			NSCompositeHardLight                                  int64 `json:"NSCompositeHardLight"`
			NSCompositeHighlight                                  int64 `json:"NSCompositeHighlight"`
			NSCompositeHue                                        int64 `json:"NSCompositeHue"`
			NSCompositeLighten                                    int64 `json:"NSCompositeLighten"`
			NSCompositeLuminosity                                 int64 `json:"NSCompositeLuminosity"`
			NSCompositeMultiply                                   int64 `json:"NSCompositeMultiply"`
			NSCompositeOverlay                                    int64 `json:"NSCompositeOverlay"`
			NSCompositePlusDarker                                 int64 `json:"NSCompositePlusDarker"`
			NSCompositePlusLighter                                int64 `json:"NSCompositePlusLighter"`
			NSCompositeSaturation                                 int64 `json:"NSCompositeSaturation"`
			NSCompositeScreen                                     int64 `json:"NSCompositeScreen"`
			NSCompositeSoftLight                                  int64 `json:"NSCompositeSoftLight"`
			NSCompositeSourceAtop                                 int64 `json:"NSCompositeSourceAtop"`
			NSCompositeSourceIn                                   int64 `json:"NSCompositeSourceIn"`
			NSCompositeSourceOut                                  int64 `json:"NSCompositeSourceOut"`
			NSCompositeSourceOver                                 int64 `json:"NSCompositeSourceOver"`
			NSCompositeXOR                                        int64 `json:"NSCompositeXOR"`
			NSCompressedFontMask                                  int64 `json:"NSCompressedFontMask"`
			NSCondensedFontMask                                   int64 `json:"NSCondensedFontMask"`
			NSContentsCellMask                                    int64 `json:"NSContentsCellMask"`
			NSContinuousCapacityLevelIndicatorStyle               int64 `json:"NSContinuousCapacityLevelIndicatorStyle"`
			NSControlGlyph                                        int64 `json:"NSControlGlyph"`
			NSControlKeyMask                                      int64 `json:"NSControlKeyMask"`
			NSCorrectionIndicatorTypeDefault                      int64 `json:"NSCorrectionIndicatorTypeDefault"`
			NSCorrectionIndicatorTypeGuesses                      int64 `json:"NSCorrectionIndicatorTypeGuesses"`
			NSCorrectionIndicatorTypeReversion                    int64 `json:"NSCorrectionIndicatorTypeReversion"`
			NSCorrectionResponseAccepted                          int64 `json:"NSCorrectionResponseAccepted"`
			NSCorrectionResponseEdited                            int64 `json:"NSCorrectionResponseEdited"`
			NSCorrectionResponseIgnored                           int64 `json:"NSCorrectionResponseIgnored"`
			NSCorrectionResponseNone                              int64 `json:"NSCorrectionResponseNone"`
			NSCorrectionResponseRejected                          int64 `json:"NSCorrectionResponseRejected"`
			NSCorrectionResponseReverted                          int64 `json:"NSCorrectionResponseReverted"`
			NSCrayonModeColorPanel                                int64 `json:"NSCrayonModeColorPanel"`
			NSCriticalAlertStyle                                  int64 `json:"NSCriticalAlertStyle"`
			NSCriticalRequest                                     int64 `json:"NSCriticalRequest"`
			NSCursorPointingDevice                                int64 `json:"NSCursorPointingDevice"`
			NSCursorUpdate                                        int64 `json:"NSCursorUpdate"`
			NSCursorUpdateMask                                    int64 `json:"NSCursorUpdateMask"`
			NSCurveToBezierPathElement                            int64 `json:"NSCurveToBezierPathElement"`
			NSCustomPaletteModeColorPanel                         int64 `json:"NSCustomPaletteModeColorPanel"`
			NSDecimalTabStopType                                  int64 `json:"NSDecimalTabStopType"`
			NSDefaultControlTint                                  int64 `json:"NSDefaultControlTint"`
			NSDeleteCharFunctionKey                               int64 `json:"NSDeleteCharFunctionKey"`
			NSDeleteCharacter                                     int64 `json:"NSDeleteCharacter"`
			NSDeleteFunctionKey                                   int64 `json:"NSDeleteFunctionKey"`
			NSDeleteLineFunctionKey                               int64 `json:"NSDeleteLineFunctionKey"`
			NSDescendingPageOrder                                 int64 `json:"NSDescendingPageOrder"`
			NSDeviceIndependentModifierFlagsMask                  int64 `json:"NSDeviceIndependentModifierFlagsMask"`
			NSDeviceNColorSpaceModel                              int64 `json:"NSDeviceNColorSpaceModel"`
			NSDirectSelection                                     int64 `json:"NSDirectSelection"`
			NSDisclosureBezelStyle                                int64 `json:"NSDisclosureBezelStyle"`
			NSDiscreteCapacityLevelIndicatorStyle                 int64 `json:"NSDiscreteCapacityLevelIndicatorStyle"`
			NSDisplayWindowRunLoopOrdering                        int64 `json:"NSDisplayWindowRunLoopOrdering"`
			NSDocModalWindowMask                                  int64 `json:"NSDocModalWindowMask"`
			NSDoubleType                                          int64 `json:"NSDoubleType"`
			NSDownArrowFunctionKey                                int64 `json:"NSDownArrowFunctionKey"`
			NSDownTextMovement                                    int64 `json:"NSDownTextMovement"`
			NSDragOperationAll                                    int64 `json:"NSDragOperationAll"`
			NSDragOperationAllObsolete                            int64 `json:"NSDragOperationAll_Obsolete"`
			NSDragOperationCopy                                   int64 `json:"NSDragOperationCopy"`
			NSDragOperationDelete                                 int64 `json:"NSDragOperationDelete"`
			NSDragOperationEvery                                  int64 `json:"NSDragOperationEvery"`
			NSDragOperationGeneric                                int64 `json:"NSDragOperationGeneric"`
			NSDragOperationLink                                   int64 `json:"NSDragOperationLink"`
			NSDragOperationMove                                   int64 `json:"NSDragOperationMove"`
			NSDragOperationNone                                   int64 `json:"NSDragOperationNone"`
			NSDragOperationPrivate                                int64 `json:"NSDragOperationPrivate"`
			NSDraggingContextOutsideApplication                   int64 `json:"NSDraggingContextOutsideApplication"`
			NSDraggingContextWithinApplication                    int64 `json:"NSDraggingContextWithinApplication"`
			NSDraggingFormationDefault                            int64 `json:"NSDraggingFormationDefault"`
			NSDraggingFormationList                               int64 `json:"NSDraggingFormationList"`
			NSDraggingFormationNone                               int64 `json:"NSDraggingFormationNone"`
			NSDraggingFormationPile                               int64 `json:"NSDraggingFormationPile"`
			NSDraggingFormationStack                              int64 `json:"NSDraggingFormationStack"`
			NSDraggingItemEnumerationClearNonenumeratedImages     int64 `json:"NSDraggingItemEnumerationClearNonenumeratedImages"`
			NSDraggingItemEnumerationConcurrent                   int64 `json:"NSDraggingItemEnumerationConcurrent"`
			NSDrawerClosedState                                   int64 `json:"NSDrawerClosedState"`
			NSDrawerClosingState                                  int64 `json:"NSDrawerClosingState"`
			NSDrawerOpenState                                     int64 `json:"NSDrawerOpenState"`
			NSDrawerOpeningState                                  int64 `json:"NSDrawerOpeningState"`
			NSEndFunctionKey                                      int64 `json:"NSEndFunctionKey"`
			NSEnterCharacter                                      int64 `json:"NSEnterCharacter"`
			NSEraDatePickerElementFlag                            int64 `json:"NSEraDatePickerElementFlag"`
			NSEraserPointingDevice                                int64 `json:"NSEraserPointingDevice"`
			NSEvenOddWindingRule                                  int64 `json:"NSEvenOddWindingRule"`
			NSEventGestureAxisHorizontal                          int64 `json:"NSEventGestureAxisHorizontal"`
			NSEventGestureAxisNone                                int64 `json:"NSEventGestureAxisNone"`
			NSEventGestureAxisVertical                            int64 `json:"NSEventGestureAxisVertical"`
			NSEventMaskBeginGesture                               int64 `json:"NSEventMaskBeginGesture"`
			NSEventMaskEndGesture                                 int64 `json:"NSEventMaskEndGesture"`
			NSEventMaskGesture                                    int64 `json:"NSEventMaskGesture"`
			NSEventMaskMagnify                                    int64 `json:"NSEventMaskMagnify"`
			NSEventMaskPressure                                   int64 `json:"NSEventMaskPressure"`
			NSEventMaskRotate                                     int64 `json:"NSEventMaskRotate"`
			NSEventMaskSmartMagnify                               int64 `json:"NSEventMaskSmartMagnify"`
			NSEventMaskSwipe                                      int64 `json:"NSEventMaskSwipe"`
			NSEventPhaseBegan                                     int64 `json:"NSEventPhaseBegan"`
			NSEventPhaseCancelled                                 int64 `json:"NSEventPhaseCancelled"`
			NSEventPhaseChanged                                   int64 `json:"NSEventPhaseChanged"`
			NSEventPhaseEnded                                     int64 `json:"NSEventPhaseEnded"`
			NSEventPhaseMayBegin                                  int64 `json:"NSEventPhaseMayBegin"`
			NSEventPhaseNone                                      int64 `json:"NSEventPhaseNone"`
			NSEventPhaseStationary                                int64 `json:"NSEventPhaseStationary"`
			NSEventSwipeTrackingClampGestureAmount                int64 `json:"NSEventSwipeTrackingClampGestureAmount"`
			NSEventSwipeTrackingLockDirection                     int64 `json:"NSEventSwipeTrackingLockDirection"`
			NSEventTypeBeginGesture                               int64 `json:"NSEventTypeBeginGesture"`
			NSEventTypeEndGesture                                 int64 `json:"NSEventTypeEndGesture"`
			NSEventTypeGesture                                    int64 `json:"NSEventTypeGesture"`
			NSEventTypeMagnify                                    int64 `json:"NSEventTypeMagnify"`
			NSEventTypePressure                                   int64 `json:"NSEventTypePressure"`
			NSEventTypeQuickLook                                  int64 `json:"NSEventTypeQuickLook"`
			NSEventTypeRotate                                     int64 `json:"NSEventTypeRotate"`
			NSEventTypeSmartMagnify                               int64 `json:"NSEventTypeSmartMagnify"`
			NSEventTypeSwipe                                      int64 `json:"NSEventTypeSwipe"`
			NSExclude10_4ElementsIconCreationOption               int64 `json:"NSExclude10_4ElementsIconCreationOption"`
			NSExcludeQuickDrawElementsIconCreationOption          int64 `json:"NSExcludeQuickDrawElementsIconCreationOption"`
			NSExecuteFunctionKey                                  int64 `json:"NSExecuteFunctionKey"`
			NSExpandedFontMask                                    int64 `json:"NSExpandedFontMask"`
			NSF10FunctionKey                                      int64 `json:"NSF10FunctionKey"`
			NSF11FunctionKey                                      int64 `json:"NSF11FunctionKey"`
			NSF12FunctionKey                                      int64 `json:"NSF12FunctionKey"`
			NSF13FunctionKey                                      int64 `json:"NSF13FunctionKey"`
			NSF14FunctionKey                                      int64 `json:"NSF14FunctionKey"`
			NSF15FunctionKey                                      int64 `json:"NSF15FunctionKey"`
			NSF16FunctionKey                                      int64 `json:"NSF16FunctionKey"`
			NSF17FunctionKey                                      int64 `json:"NSF17FunctionKey"`
			NSF18FunctionKey                                      int64 `json:"NSF18FunctionKey"`
			NSF19FunctionKey                                      int64 `json:"NSF19FunctionKey"`
			NSF1FunctionKey                                       int64 `json:"NSF1FunctionKey"`
			NSF20FunctionKey                                      int64 `json:"NSF20FunctionKey"`
			NSF21FunctionKey                                      int64 `json:"NSF21FunctionKey"`
			NSF22FunctionKey                                      int64 `json:"NSF22FunctionKey"`
			NSF23FunctionKey                                      int64 `json:"NSF23FunctionKey"`
			NSF24FunctionKey                                      int64 `json:"NSF24FunctionKey"`
			NSF25FunctionKey                                      int64 `json:"NSF25FunctionKey"`
			NSF26FunctionKey                                      int64 `json:"NSF26FunctionKey"`
			NSF27FunctionKey                                      int64 `json:"NSF27FunctionKey"`
			NSF28FunctionKey                                      int64 `json:"NSF28FunctionKey"`
			NSF29FunctionKey                                      int64 `json:"NSF29FunctionKey"`
			NSF2FunctionKey                                       int64 `json:"NSF2FunctionKey"`
			NSF30FunctionKey                                      int64 `json:"NSF30FunctionKey"`
			NSF31FunctionKey                                      int64 `json:"NSF31FunctionKey"`
			NSF32FunctionKey                                      int64 `json:"NSF32FunctionKey"`
			NSF33FunctionKey                                      int64 `json:"NSF33FunctionKey"`
			NSF34FunctionKey                                      int64 `json:"NSF34FunctionKey"`
			NSF35FunctionKey                                      int64 `json:"NSF35FunctionKey"`
			NSF3FunctionKey                                       int64 `json:"NSF3FunctionKey"`
			NSF4FunctionKey                                       int64 `json:"NSF4FunctionKey"`
			NSF5FunctionKey                                       int64 `json:"NSF5FunctionKey"`
			NSF6FunctionKey                                       int64 `json:"NSF6FunctionKey"`
			NSF7FunctionKey                                       int64 `json:"NSF7FunctionKey"`
			NSF8FunctionKey                                       int64 `json:"NSF8FunctionKey"`
			NSF9FunctionKey                                       int64 `json:"NSF9FunctionKey"`
			NSFPCurrentField                                      int64 `json:"NSFPCurrentField"`
			NSFPPreviewButton                                     int64 `json:"NSFPPreviewButton"`
			NSFPPreviewField                                      int64 `json:"NSFPPreviewField"`
			NSFPRevertButton                                      int64 `json:"NSFPRevertButton"`
			NSFPSetButton                                         int64 `json:"NSFPSetButton"`
			NSFPSizeField                                         int64 `json:"NSFPSizeField"`
			NSFPSizeTitle                                         int64 `json:"NSFPSizeTitle"`
			NSFileHandlingPanelCancelButton                       int64 `json:"NSFileHandlingPanelCancelButton"`
			NSFileHandlingPanelOKButton                           int64 `json:"NSFileHandlingPanelOKButton"`
			NSFindFunctionKey                                     int64 `json:"NSFindFunctionKey"`
			NSFindPanelActionNext                                 int64 `json:"NSFindPanelActionNext"`
			NSFindPanelActionPrevious                             int64 `json:"NSFindPanelActionPrevious"`
			NSFindPanelActionReplace                              int64 `json:"NSFindPanelActionReplace"`
			NSFindPanelActionReplaceAll                           int64 `json:"NSFindPanelActionReplaceAll"`
			NSFindPanelActionReplaceAllInSelection                int64 `json:"NSFindPanelActionReplaceAllInSelection"`
			NSFindPanelActionReplaceAndFind                       int64 `json:"NSFindPanelActionReplaceAndFind"`
			NSFindPanelActionSelectAll                            int64 `json:"NSFindPanelActionSelectAll"`
			NSFindPanelActionSelectAllInSelection                 int64 `json:"NSFindPanelActionSelectAllInSelection"`
			NSFindPanelActionSetFindString                        int64 `json:"NSFindPanelActionSetFindString"`
			NSFindPanelActionShowFindPanel                        int64 `json:"NSFindPanelActionShowFindPanel"`
			NSFindPanelSubstringMatchTypeContains                 int64 `json:"NSFindPanelSubstringMatchTypeContains"`
			NSFindPanelSubstringMatchTypeEndsWith                 int64 `json:"NSFindPanelSubstringMatchTypeEndsWith"`
			NSFindPanelSubstringMatchTypeFullWord                 int64 `json:"NSFindPanelSubstringMatchTypeFullWord"`
			NSFindPanelSubstringMatchTypeStartsWith               int64 `json:"NSFindPanelSubstringMatchTypeStartsWith"`
			NSFitPagination                                       int64 `json:"NSFitPagination"`
			NSFixedPitchFontMask                                  int64 `json:"NSFixedPitchFontMask"`
			NSFlagsChanged                                        int64 `json:"NSFlagsChanged"`
			NSFlagsChangedMask                                    int64 `json:"NSFlagsChangedMask"`
			NSFloatType                                           int64 `json:"NSFloatType"`
			NSFloatingPointSamplesBitmapFormat                    int64 `json:"NSFloatingPointSamplesBitmapFormat"`
			NSFocusRingAbove                                      int64 `json:"NSFocusRingAbove"`
			NSFocusRingBelow                                      int64 `json:"NSFocusRingBelow"`
			NSFocusRingOnly                                       int64 `json:"NSFocusRingOnly"`
			NSFocusRingTypeDefault                                int64 `json:"NSFocusRingTypeDefault"`
			NSFocusRingTypeExterior                               int64 `json:"NSFocusRingTypeExterior"`
			NSFocusRingTypeNone                                   int64 `json:"NSFocusRingTypeNone"`
			NSFontAntialiasedIntegerAdvancementsRenderingMode     int64 `json:"NSFontAntialiasedIntegerAdvancementsRenderingMode"`
			NSFontAntialiasedRenderingMode                        int64 `json:"NSFontAntialiasedRenderingMode"`
			NSFontBoldTrait                                       int64 `json:"NSFontBoldTrait"`
			NSFontClarendonSerifsClass                            int64 `json:"NSFontClarendonSerifsClass"`
			NSFontCollectionApplicationOnlyMask                   int64 `json:"NSFontCollectionApplicationOnlyMask"`
			NSFontCollectionVisibilityComputer                    int64 `json:"NSFontCollectionVisibilityComputer"`
			NSFontCollectionVisibilityProcess                     int64 `json:"NSFontCollectionVisibilityProcess"`
			NSFontCollectionVisibilityUser                        int64 `json:"NSFontCollectionVisibilityUser"`
			NSFontCondensedTrait                                  int64 `json:"NSFontCondensedTrait"`
			NSFontDefaultRenderingMode                            int64 `json:"NSFontDefaultRenderingMode"`
			NSFontExpandedTrait                                   int64 `json:"NSFontExpandedTrait"`
			NSFontFamilyClassMask                                 int64 `json:"NSFontFamilyClassMask"`
			NSFontFreeformSerifsClass                             int64 `json:"NSFontFreeformSerifsClass"`
			NSFontIntegerAdvancementsRenderingMode                int64 `json:"NSFontIntegerAdvancementsRenderingMode"`
			NSFontItalicTrait                                     int64 `json:"NSFontItalicTrait"`
			NSFontModernSerifsClass                               int64 `json:"NSFontModernSerifsClass"`
			NSFontMonoSpaceTrait                                  int64 `json:"NSFontMonoSpaceTrait"`
			NSFontOldStyleSerifsClass                             int64 `json:"NSFontOldStyleSerifsClass"`
			NSFontOrnamentalsClass                                int64 `json:"NSFontOrnamentalsClass"`
			NSFontPanelAllEffectsModeMask                         int64 `json:"NSFontPanelAllEffectsModeMask"`
			NSFontPanelAllModesMask                               int64 `json:"NSFontPanelAllModesMask"`
			NSFontPanelCollectionModeMask                         int64 `json:"NSFontPanelCollectionModeMask"`
			NSFontPanelDocumentColorEffectModeMask                int64 `json:"NSFontPanelDocumentColorEffectModeMask"`
			NSFontPanelFaceModeMask                               int64 `json:"NSFontPanelFaceModeMask"`
			NSFontPanelShadowEffectModeMask                       int64 `json:"NSFontPanelShadowEffectModeMask"`
			NSFontPanelSizeModeMask                               int64 `json:"NSFontPanelSizeModeMask"`
			NSFontPanelStandardModesMask                          int64 `json:"NSFontPanelStandardModesMask"`
			NSFontPanelStrikethroughEffectModeMask                int64 `json:"NSFontPanelStrikethroughEffectModeMask"`
			NSFontPanelTextColorEffectModeMask                    int64 `json:"NSFontPanelTextColorEffectModeMask"`
			NSFontPanelUnderlineEffectModeMask                    int64 `json:"NSFontPanelUnderlineEffectModeMask"`
			NSFontSansSerifClass                                  int64 `json:"NSFontSansSerifClass"`
			NSFontScriptsClass                                    int64 `json:"NSFontScriptsClass"`
			NSFontSlabSerifsClass                                 int64 `json:"NSFontSlabSerifsClass"`
			NSFontSymbolicClass                                   int64 `json:"NSFontSymbolicClass"`
			NSFontTransitionalSerifsClass                         int64 `json:"NSFontTransitionalSerifsClass"`
			NSFontUIOptimizedTrait                                int64 `json:"NSFontUIOptimizedTrait"`
			NSFontUnknownClass                                    int64 `json:"NSFontUnknownClass"`
			NSFontVerticalTrait                                   int64 `json:"NSFontVerticalTrait"`
			NSFormFeedCharacter                                   int64 `json:"NSFormFeedCharacter"`
			NSFullScreenWindowMask                                int64 `json:"NSFullScreenWindowMask"`
			NSFullSizeContentViewWindowMask                       int64 `json:"NSFullSizeContentViewWindowMask"`
			NSFunctionKeyMask                                     int64 `json:"NSFunctionKeyMask"`
			NSGIFFileType                                         int64 `json:"NSGIFFileType"`
			NSGestureRecognizerStateBegan                         int64 `json:"NSGestureRecognizerStateBegan"`
			NSGestureRecognizerStateCancelled                     int64 `json:"NSGestureRecognizerStateCancelled"`
			NSGestureRecognizerStateChanged                       int64 `json:"NSGestureRecognizerStateChanged"`
			NSGestureRecognizerStateEnded                         int64 `json:"NSGestureRecognizerStateEnded"`
			NSGestureRecognizerStateFailed                        int64 `json:"NSGestureRecognizerStateFailed"`
			NSGestureRecognizerStatePossible                      int64 `json:"NSGestureRecognizerStatePossible"`
			NSGestureRecognizerStateRecognized                    int64 `json:"NSGestureRecognizerStateRecognized"`
			NSGlyphAttributeBidiLevel                             int64 `json:"NSGlyphAttributeBidiLevel"`
			NSGlyphAttributeElastic                               int64 `json:"NSGlyphAttributeElastic"`
			NSGlyphAttributeInscribe                              int64 `json:"NSGlyphAttributeInscribe"`
			NSGlyphAttributeSoft                                  int64 `json:"NSGlyphAttributeSoft"`
			NSGlyphInscribeAbove                                  int64 `json:"NSGlyphInscribeAbove"`
			NSGlyphInscribeBase                                   int64 `json:"NSGlyphInscribeBase"`
			NSGlyphInscribeBelow                                  int64 `json:"NSGlyphInscribeBelow"`
			NSGlyphInscribeOverBelow                              int64 `json:"NSGlyphInscribeOverBelow"`
			NSGlyphInscribeOverstrike                             int64 `json:"NSGlyphInscribeOverstrike"`
			NSGradientConcaveStrong                               int64 `json:"NSGradientConcaveStrong"`
			NSGradientConcaveWeak                                 int64 `json:"NSGradientConcaveWeak"`
			NSGradientConvexStrong                                int64 `json:"NSGradientConvexStrong"`
			NSGradientConvexWeak                                  int64 `json:"NSGradientConvexWeak"`
			NSGradientDrawsAfterEndingLocation                    int64 `json:"NSGradientDrawsAfterEndingLocation"`
			NSGradientDrawsBeforeStartingLocation                 int64 `json:"NSGradientDrawsBeforeStartingLocation"`
			NSGradientNone                                        int64 `json:"NSGradientNone"`
			NSGraphiteControlTint                                 int64 `json:"NSGraphiteControlTint"`
			NSGrayColorSpaceModel                                 int64 `json:"NSGrayColorSpaceModel"`
			NSGrayModeColorPanel                                  int64 `json:"NSGrayModeColorPanel"`
			NSGrooveBorder                                        int64 `json:"NSGrooveBorder"`
			NSHSBModeColorPanel                                   int64 `json:"NSHSBModeColorPanel"`
			NSHUDWindowMask                                       int64 `json:"NSHUDWindowMask"`
			NSHeavierFontAction                                   int64 `json:"NSHeavierFontAction"`
			NSHelpButtonBezelStyle                                int64 `json:"NSHelpButtonBezelStyle"`
			NSHelpFunctionKey                                     int64 `json:"NSHelpFunctionKey"`
			NSHelpKeyMask                                         int64 `json:"NSHelpKeyMask"`
			NSHighlightModeMatrix                                 int64 `json:"NSHighlightModeMatrix"`
			NSHomeFunctionKey                                     int64 `json:"NSHomeFunctionKey"`
			NSHorizontalRuler                                     int64 `json:"NSHorizontalRuler"`
			NSHourMinuteDatePickerElementFlag                     int64 `json:"NSHourMinuteDatePickerElementFlag"`
			NSHourMinuteSecondDatePickerElementFlag               int64 `json:"NSHourMinuteSecondDatePickerElementFlag"`
			NSIdentityMappingCharacterCollection                  int64 `json:"NSIdentityMappingCharacterCollection"`
			NSIllegalTextMovement                                 int64 `json:"NSIllegalTextMovement"`
			NSImageAbove                                          int64 `json:"NSImageAbove"`
			NSImageAlignBottom                                    int64 `json:"NSImageAlignBottom"`
			NSImageAlignBottomLeft                                int64 `json:"NSImageAlignBottomLeft"`
			NSImageAlignBottomRight                               int64 `json:"NSImageAlignBottomRight"`
			NSImageAlignCenter                                    int64 `json:"NSImageAlignCenter"`
			NSImageAlignLeft                                      int64 `json:"NSImageAlignLeft"`
			NSImageAlignRight                                     int64 `json:"NSImageAlignRight"`
			NSImageAlignTop                                       int64 `json:"NSImageAlignTop"`
			NSImageAlignTopLeft                                   int64 `json:"NSImageAlignTopLeft"`
			NSImageAlignTopRight                                  int64 `json:"NSImageAlignTopRight"`
			NSImageBelow                                          int64 `json:"NSImageBelow"`
			NSImageCacheAlways                                    int64 `json:"NSImageCacheAlways"`
			NSImageCacheBySize                                    int64 `json:"NSImageCacheBySize"`
			NSImageCacheDefault                                   int64 `json:"NSImageCacheDefault"`
			NSImageCacheNever                                     int64 `json:"NSImageCacheNever"`
			NSImageCellType                                       int64 `json:"NSImageCellType"`
			NSImageFrameButton                                    int64 `json:"NSImageFrameButton"`
			NSImageFrameGrayBezel                                 int64 `json:"NSImageFrameGrayBezel"`
			NSImageFrameGroove                                    int64 `json:"NSImageFrameGroove"`
			NSImageFrameNone                                      int64 `json:"NSImageFrameNone"`
			NSImageFramePhoto                                     int64 `json:"NSImageFramePhoto"`
			NSImageInterpolationDefault                           int64 `json:"NSImageInterpolationDefault"`
			NSImageInterpolationHigh                              int64 `json:"NSImageInterpolationHigh"`
			NSImageInterpolationLow                               int64 `json:"NSImageInterpolationLow"`
			NSImageInterpolationMedium                            int64 `json:"NSImageInterpolationMedium"`
			NSImageInterpolationNone                              int64 `json:"NSImageInterpolationNone"`
			NSImageLeft                                           int64 `json:"NSImageLeft"`
			NSImageLoadStatusCancelled                            int64 `json:"NSImageLoadStatusCancelled"`
			NSImageLoadStatusCompleted                            int64 `json:"NSImageLoadStatusCompleted"`
			NSImageLoadStatusInvalidData                          int64 `json:"NSImageLoadStatusInvalidData"`
			NSImageLoadStatusReadError                            int64 `json:"NSImageLoadStatusReadError"`
			NSImageLoadStatusUnexpectedEOF                        int64 `json:"NSImageLoadStatusUnexpectedEOF"`
			NSImageOnly                                           int64 `json:"NSImageOnly"`
			NSImageOverlaps                                       int64 `json:"NSImageOverlaps"`
			NSImageRepLoadStatusCompleted                         int64 `json:"NSImageRepLoadStatusCompleted"`
			NSImageRepLoadStatusInvalidData                       int64 `json:"NSImageRepLoadStatusInvalidData"`
			NSImageRepLoadStatusReadingHeader                     int64 `json:"NSImageRepLoadStatusReadingHeader"`
			NSImageRepLoadStatusUnexpectedEOF                     int64 `json:"NSImageRepLoadStatusUnexpectedEOF"`
			NSImageRepLoadStatusUnknownType                       int64 `json:"NSImageRepLoadStatusUnknownType"`
			NSImageRepLoadStatusWillNeedAllData                   int64 `json:"NSImageRepLoadStatusWillNeedAllData"`
			NSImageRepMatchesDevice                               int64 `json:"NSImageRepMatchesDevice"`
			NSImageResizingModeStretch                            int64 `json:"NSImageResizingModeStretch"`
			NSImageResizingModeTile                               int64 `json:"NSImageResizingModeTile"`
			NSImageRight                                          int64 `json:"NSImageRight"`
			NSImageScaleAxesIndependently                         int64 `json:"NSImageScaleAxesIndependently"`
			NSImageScaleNone                                      int64 `json:"NSImageScaleNone"`
			NSImageScaleProportionallyDown                        int64 `json:"NSImageScaleProportionallyDown"`
			NSImageScaleProportionallyUpOrDown                    int64 `json:"NSImageScaleProportionallyUpOrDown"`
			NSIndexedColorSpaceModel                              int64 `json:"NSIndexedColorSpaceModel"`
			NSInformationalAlertStyle                             int64 `json:"NSInformationalAlertStyle"`
			NSInformationalRequest                                int64 `json:"NSInformationalRequest"`
			NSInlineBezelStyle                                    int64 `json:"NSInlineBezelStyle"`
			NSInsertCharFunctionKey                               int64 `json:"NSInsertCharFunctionKey"`
			NSInsertFunctionKey                                   int64 `json:"NSInsertFunctionKey"`
			NSInsertLineFunctionKey                               int64 `json:"NSInsertLineFunctionKey"`
			NSIntType                                             int64 `json:"NSIntType"`
			NSItalicFontMask                                      int64 `json:"NSItalicFontMask"`
			NSJPEG2000FileType                                    int64 `json:"NSJPEG2000FileType"`
			NSJPEGFileType                                        int64 `json:"NSJPEGFileType"`
			NSJustifiedTextAlignment                              int64 `json:"NSJustifiedTextAlignment"`
			NSKeyDown                                             int64 `json:"NSKeyDown"`
			NSKeyDownMask                                         int64 `json:"NSKeyDownMask"`
			NSKeyUp                                               int64 `json:"NSKeyUp"`
			NSKeyUpMask                                           int64 `json:"NSKeyUpMask"`
			NSLABColorSpaceModel                                  int64 `json:"NSLABColorSpaceModel"`
			NSLandscapeOrientation                                int64 `json:"NSLandscapeOrientation"`
			NSLayoutAttributeBaseline                             int64 `json:"NSLayoutAttributeBaseline"`
			NSLayoutAttributeBottom                               int64 `json:"NSLayoutAttributeBottom"`
			NSLayoutAttributeCenterX                              int64 `json:"NSLayoutAttributeCenterX"`
			NSLayoutAttributeCenterY                              int64 `json:"NSLayoutAttributeCenterY"`
			NSLayoutAttributeHeight                               int64 `json:"NSLayoutAttributeHeight"`
			NSLayoutAttributeLeading                              int64 `json:"NSLayoutAttributeLeading"`
			NSLayoutAttributeLeft                                 int64 `json:"NSLayoutAttributeLeft"`
			NSLayoutAttributeNotAnAttribute                       int64 `json:"NSLayoutAttributeNotAnAttribute"`
			NSLayoutAttributeRight                                int64 `json:"NSLayoutAttributeRight"`
			NSLayoutAttributeTop                                  int64 `json:"NSLayoutAttributeTop"`
			NSLayoutAttributeTrailing                             int64 `json:"NSLayoutAttributeTrailing"`
			NSLayoutAttributeWidth                                int64 `json:"NSLayoutAttributeWidth"`
			NSLayoutConstraintOrientationHorizontal               int64 `json:"NSLayoutConstraintOrientationHorizontal"`
			NSLayoutConstraintOrientationVertical                 int64 `json:"NSLayoutConstraintOrientationVertical"`
			NSLayoutFormatAlignAllBaseline                        int64 `json:"NSLayoutFormatAlignAllBaseline"`
			NSLayoutFormatAlignAllBottom                          int64 `json:"NSLayoutFormatAlignAllBottom"`
			NSLayoutFormatAlignAllCenterX                         int64 `json:"NSLayoutFormatAlignAllCenterX"`
			NSLayoutFormatAlignAllCenterY                         int64 `json:"NSLayoutFormatAlignAllCenterY"`
			NSLayoutFormatAlignAllLeading                         int64 `json:"NSLayoutFormatAlignAllLeading"`
			NSLayoutFormatAlignAllLeft                            int64 `json:"NSLayoutFormatAlignAllLeft"`
			NSLayoutFormatAlignAllRight                           int64 `json:"NSLayoutFormatAlignAllRight"`
			NSLayoutFormatAlignAllTop                             int64 `json:"NSLayoutFormatAlignAllTop"`
			NSLayoutFormatAlignAllTrailing                        int64 `json:"NSLayoutFormatAlignAllTrailing"`
			NSLayoutFormatAlignmentMask                           int64 `json:"NSLayoutFormatAlignmentMask"`
			NSLayoutFormatDirectionLeadingToTrailing              int64 `json:"NSLayoutFormatDirectionLeadingToTrailing"`
			NSLayoutFormatDirectionLeftToRight                    int64 `json:"NSLayoutFormatDirectionLeftToRight"`
			NSLayoutFormatDirectionMask                           int64 `json:"NSLayoutFormatDirectionMask"`
			NSLayoutFormatDirectionRightToLeft                    int64 `json:"NSLayoutFormatDirectionRightToLeft"`
			NSLayoutRelationEqual                                 int64 `json:"NSLayoutRelationEqual"`
			NSLayoutRelationGreaterThanOrEqual                    int64 `json:"NSLayoutRelationGreaterThanOrEqual"`
			NSLayoutRelationLessThanOrEqual                       int64 `json:"NSLayoutRelationLessThanOrEqual"`
			NSLeftArrowFunctionKey                                int64 `json:"NSLeftArrowFunctionKey"`
			NSLeftMouseDown                                       int64 `json:"NSLeftMouseDown"`
			NSLeftMouseDownMask                                   int64 `json:"NSLeftMouseDownMask"`
			NSLeftMouseDragged                                    int64 `json:"NSLeftMouseDragged"`
			NSLeftMouseDraggedMask                                int64 `json:"NSLeftMouseDraggedMask"`
			NSLeftMouseUp                                         int64 `json:"NSLeftMouseUp"`
			NSLeftMouseUpMask                                     int64 `json:"NSLeftMouseUpMask"`
			NSLeftTabStopType                                     int64 `json:"NSLeftTabStopType"`
			NSLeftTabsBezelBorder                                 int64 `json:"NSLeftTabsBezelBorder"`
			NSLeftTextAlignment                                   int64 `json:"NSLeftTextAlignment"`
			NSLeftTextMovement                                    int64 `json:"NSLeftTextMovement"`
			NSLighterFontAction                                   int64 `json:"NSLighterFontAction"`
			NSLineBorder                                          int64 `json:"NSLineBorder"`
			NSLineBreakByCharWrapping                             int64 `json:"NSLineBreakByCharWrapping"`
			NSLineBreakByClipping                                 int64 `json:"NSLineBreakByClipping"`
			NSLineBreakByTruncatingHead                           int64 `json:"NSLineBreakByTruncatingHead"`
			NSLineBreakByTruncatingMiddle                         int64 `json:"NSLineBreakByTruncatingMiddle"`
			NSLineBreakByTruncatingTail                           int64 `json:"NSLineBreakByTruncatingTail"`
			NSLineBreakByWordWrapping                             int64 `json:"NSLineBreakByWordWrapping"`
			NSLineDoesntMove                                      int64 `json:"NSLineDoesntMove"`
			NSLineMovesDown                                       int64 `json:"NSLineMovesDown"`
			NSLineMovesLeft                                       int64 `json:"NSLineMovesLeft"`
			NSLineMovesRight                                      int64 `json:"NSLineMovesRight"`
			NSLineMovesUp                                         int64 `json:"NSLineMovesUp"`
			NSLineSeparatorCharacter                              int64 `json:"NSLineSeparatorCharacter"`
			NSLineSweepDown                                       int64 `json:"NSLineSweepDown"`
			NSLineSweepLeft                                       int64 `json:"NSLineSweepLeft"`
			NSLineSweepRight                                      int64 `json:"NSLineSweepRight"`
			NSLineSweepUp                                         int64 `json:"NSLineSweepUp"`
			NSLineToBezierPathElement                             int64 `json:"NSLineToBezierPathElement"`
			NSLinearSlider                                        int64 `json:"NSLinearSlider"`
			NSListModeMatrix                                      int64 `json:"NSListModeMatrix"`
			NSMacintoshInterfaceStyle                             int64 `json:"NSMacintoshInterfaceStyle"`
			NSMediaLibraryAudio                                   int64 `json:"NSMediaLibraryAudio"`
			NSMediaLibraryImage                                   int64 `json:"NSMediaLibraryImage"`
			NSMediaLibraryMovie                                   int64 `json:"NSMediaLibraryMovie"`
			NSMenuFunctionKey                                     int64 `json:"NSMenuFunctionKey"`
			NSMenuPropertyItemAccessibilityDescription            int64 `json:"NSMenuPropertyItemAccessibilityDescription"`
			NSMenuPropertyItemAttributedTitle                     int64 `json:"NSMenuPropertyItemAttributedTitle"`
			NSMenuPropertyItemEnabled                             int64 `json:"NSMenuPropertyItemEnabled"`
			NSMenuPropertyItemImage                               int64 `json:"NSMenuPropertyItemImage"`
			NSMenuPropertyItemKeyEquivalent                       int64 `json:"NSMenuPropertyItemKeyEquivalent"`
			NSMenuPropertyItemTitle                               int64 `json:"NSMenuPropertyItemTitle"`
			NSMiniControlSize                                     int64 `json:"NSMiniControlSize"`
			NSMiniaturizableWindowMask                            int64 `json:"NSMiniaturizableWindowMask"`
			NSMiterLineJoinStyle                                  int64 `json:"NSMiterLineJoinStyle"`
			NSMixedState                                          int64 `json:"NSMixedState"`
			NSModalResponseAbort                                  int64 `json:"NSModalResponseAbort"`
			NSModalResponseCancel                                 int64 `json:"NSModalResponseCancel"`
			NSModalResponseContinue                               int64 `json:"NSModalResponseContinue"`
			NSModalResponseOK                                     int64 `json:"NSModalResponseOK"`
			NSModalResponseStop                                   int64 `json:"NSModalResponseStop"`
			NSModeSwitchFunctionKey                               int64 `json:"NSModeSwitchFunctionKey"`
			NSMomentaryChangeButton                               int64 `json:"NSMomentaryChangeButton"`
			NSMomentaryLight                                      int64 `json:"NSMomentaryLight"`
			NSMomentaryLightButton                                int64 `json:"NSMomentaryLightButton"`
			NSMomentaryPushButton                                 int64 `json:"NSMomentaryPushButton"`
			NSMomentaryPushInButton                               int64 `json:"NSMomentaryPushInButton"`
			NSMouseEntered                                        int64 `json:"NSMouseEntered"`
			NSMouseEnteredMask                                    int64 `json:"NSMouseEnteredMask"`
			NSMouseEventSubtype                                   int64 `json:"NSMouseEventSubtype"`
			NSMouseExited                                         int64 `json:"NSMouseExited"`
			NSMouseExitedMask                                     int64 `json:"NSMouseExitedMask"`
			NSMouseMoved                                          int64 `json:"NSMouseMoved"`
			NSMouseMovedMask                                      int64 `json:"NSMouseMovedMask"`
			NSMoveToBezierPathElement                             int64 `json:"NSMoveToBezierPathElement"`
			NSMultiLevelAcceleratorButton                         int64 `json:"NSMultiLevelAcceleratorButton"`
			NSNarrowFontMask                                      int64 `json:"NSNarrowFontMask"`
			NSNativeShortGlyphPacking                             int64 `json:"NSNativeShortGlyphPacking"`
			NSNaturalTextAlignment                                int64 `json:"NSNaturalTextAlignment"`
			NSNewlineCharacter                                    int64 `json:"NSNewlineCharacter"`
			NSNextFunctionKey                                     int64 `json:"NSNextFunctionKey"`
			NSNextStepInterfaceStyle                              int64 `json:"NSNextStepInterfaceStyle"`
			NSNoBorder                                            int64 `json:"NSNoBorder"`
			NSNoCellMask                                          int64 `json:"NSNoCellMask"`
			NSNoFontChangeAction                                  int64 `json:"NSNoFontChangeAction"`
			NSNoImage                                             int64 `json:"NSNoImage"`
			NSNoInterfaceStyle                                    int64 `json:"NSNoInterfaceStyle"`
			NSNoModeColorPanel                                    int64 `json:"NSNoModeColorPanel"`
			NSNoScrollerParts                                     int64 `json:"NSNoScrollerParts"`
			NSNoTabsBezelBorder                                   int64 `json:"NSNoTabsBezelBorder"`
			NSNoTabsLineBorder                                    int64 `json:"NSNoTabsLineBorder"`
			NSNoTabsNoBorder                                      int64 `json:"NSNoTabsNoBorder"`
			NSNoTitle                                             int64 `json:"NSNoTitle"`
			NSNoUnderlineStyle                                    int64 `json:"NSNoUnderlineStyle"`
			NSNonStandardCharacterSetFontMask                     int64 `json:"NSNonStandardCharacterSetFontMask"`
			NSNonZeroWindingRule                                  int64 `json:"NSNonZeroWindingRule"`
			NSNonactivatingPanelMask                              int64 `json:"NSNonactivatingPanelMask"`
			NSNullCellType                                        int64 `json:"NSNullCellType"`
			NSNullGlyph                                           int64 `json:"NSNullGlyph"`
			NSNumericPadKeyMask                                   int64 `json:"NSNumericPadKeyMask"`
			NSOKButton                                            int64 `json:"NSOKButton"`
			NsopenglCurrentVersion                                int64 `json:"NSOPENGL_CURRENT_VERSION"`
			NSOffState                                            int64 `json:"NSOffState"`
			NSOnOffButton                                         int64 `json:"NSOnOffButton"`
			NSOnState                                             int64 `json:"NSOnState"`
			NSOnlyScrollerArrows                                  int64 `json:"NSOnlyScrollerArrows"`
			NSOpenGLCPCurrentRendererID                           int64 `json:"NSOpenGLCPCurrentRendererID"`
			NSOpenGLCPGPUFragmentProcessing                       int64 `json:"NSOpenGLCPGPUFragmentProcessing"`
			NSOpenGLCPGPUVertexProcessing                         int64 `json:"NSOpenGLCPGPUVertexProcessing"`
			NSOpenGLCPHasDrawable                                 int64 `json:"NSOpenGLCPHasDrawable"`
			NSOpenGLCPMPSwapsInFlight                             int64 `json:"NSOpenGLCPMPSwapsInFlight"`
			NSOpenGLCPRasterizationEnable                         int64 `json:"NSOpenGLCPRasterizationEnable"`
			NSOpenGLCPReclaimResources                            int64 `json:"NSOpenGLCPReclaimResources"`
			NSOpenGLCPStateValidation                             int64 `json:"NSOpenGLCPStateValidation"`
			NSOpenGLCPSurfaceBackingSize                          int64 `json:"NSOpenGLCPSurfaceBackingSize"`
			NSOpenGLCPSurfaceOpacity                              int64 `json:"NSOpenGLCPSurfaceOpacity"`
			NSOpenGLCPSurfaceOrder                                int64 `json:"NSOpenGLCPSurfaceOrder"`
			NSOpenGLCPSurfaceSurfaceVolatile                      int64 `json:"NSOpenGLCPSurfaceSurfaceVolatile"`
			NSOpenGLCPSwapInterval                                int64 `json:"NSOpenGLCPSwapInterval"`
			NSOpenGLCPSwapRectangle                               int64 `json:"NSOpenGLCPSwapRectangle"`
			NSOpenGLCPSwapRectangleEnable                         int64 `json:"NSOpenGLCPSwapRectangleEnable"`
			NSOpenGLGOClearFormatCache                            int64 `json:"NSOpenGLGOClearFormatCache"`
			NSOpenGLGOFormatCacheSize                             int64 `json:"NSOpenGLGOFormatCacheSize"`
			NSOpenGLGOResetLibrary                                int64 `json:"NSOpenGLGOResetLibrary"`
			NSOpenGLGORetainRenderers                             int64 `json:"NSOpenGLGORetainRenderers"`
			NSOpenGLGOUseBuildCache                               int64 `json:"NSOpenGLGOUseBuildCache"`
			NSOpenGLPFAAccelerated                                int64 `json:"NSOpenGLPFAAccelerated"`
			NSOpenGLPFAAcceleratedCompute                         int64 `json:"NSOpenGLPFAAcceleratedCompute"`
			NSOpenGLPFAAccumSize                                  int64 `json:"NSOpenGLPFAAccumSize"`
			NSOpenGLPFAAllRenderers                               int64 `json:"NSOpenGLPFAAllRenderers"`
			NSOpenGLPFAAllowOfflineRenderers                      int64 `json:"NSOpenGLPFAAllowOfflineRenderers"`
			NSOpenGLPFAAlphaSize                                  int64 `json:"NSOpenGLPFAAlphaSize"`
			NSOpenGLPFAAuxBuffers                                 int64 `json:"NSOpenGLPFAAuxBuffers"`
			NSOpenGLPFAAuxDepthStencil                            int64 `json:"NSOpenGLPFAAuxDepthStencil"`
			NSOpenGLPFABackingStore                               int64 `json:"NSOpenGLPFABackingStore"`
			NSOpenGLPFAClosestPolicy                              int64 `json:"NSOpenGLPFAClosestPolicy"`
			NSOpenGLPFAColorFloat                                 int64 `json:"NSOpenGLPFAColorFloat"`
			NSOpenGLPFAColorSize                                  int64 `json:"NSOpenGLPFAColorSize"`
			NSOpenGLPFACompliant                                  int64 `json:"NSOpenGLPFACompliant"`
			NSOpenGLPFADepthSize                                  int64 `json:"NSOpenGLPFADepthSize"`
			NSOpenGLPFADoubleBuffer                               int64 `json:"NSOpenGLPFADoubleBuffer"`
			NSOpenGLPFAFullScreen                                 int64 `json:"NSOpenGLPFAFullScreen"`
			NSOpenGLPFAMPSafe                                     int64 `json:"NSOpenGLPFAMPSafe"`
			NSOpenGLPFAMaximumPolicy                              int64 `json:"NSOpenGLPFAMaximumPolicy"`
			NSOpenGLPFAMinimumPolicy                              int64 `json:"NSOpenGLPFAMinimumPolicy"`
			NSOpenGLPFAMultiScreen                                int64 `json:"NSOpenGLPFAMultiScreen"`
			NSOpenGLPFAMultisample                                int64 `json:"NSOpenGLPFAMultisample"`
			NSOpenGLPFANoRecovery                                 int64 `json:"NSOpenGLPFANoRecovery"`
			NSOpenGLPFAOffScreen                                  int64 `json:"NSOpenGLPFAOffScreen"`
			NSOpenGLPFAOpenGLProfile                              int64 `json:"NSOpenGLPFAOpenGLProfile"`
			NSOpenGLPFAPixelBuffer                                int64 `json:"NSOpenGLPFAPixelBuffer"`
			NSOpenGLPFARemotePixelBuffer                          int64 `json:"NSOpenGLPFARemotePixelBuffer"`
			NSOpenGLPFARendererID                                 int64 `json:"NSOpenGLPFARendererID"`
			NSOpenGLPFARobust                                     int64 `json:"NSOpenGLPFARobust"`
			NSOpenGLPFASampleAlpha                                int64 `json:"NSOpenGLPFASampleAlpha"`
			NSOpenGLPFASampleBuffers                              int64 `json:"NSOpenGLPFASampleBuffers"`
			NSOpenGLPFASamples                                    int64 `json:"NSOpenGLPFASamples"`
			NSOpenGLPFAScreenMask                                 int64 `json:"NSOpenGLPFAScreenMask"`
			NSOpenGLPFASingleRenderer                             int64 `json:"NSOpenGLPFASingleRenderer"`
			NSOpenGLPFAStencilSize                                int64 `json:"NSOpenGLPFAStencilSize"`
			NSOpenGLPFAStereo                                     int64 `json:"NSOpenGLPFAStereo"`
			NSOpenGLPFASupersample                                int64 `json:"NSOpenGLPFASupersample"`
			NSOpenGLPFATripleBuffer                               int64 `json:"NSOpenGLPFATripleBuffer"`
			NSOpenGLPFAVirtualScreenCount                         int64 `json:"NSOpenGLPFAVirtualScreenCount"`
			NSOpenGLPFAWindow                                     int64 `json:"NSOpenGLPFAWindow"`
			NSOpenGLProfileVersion3_2Core                         int64 `json:"NSOpenGLProfileVersion3_2Core"`
			NSOpenGLProfileVersion4_1Core                         int64 `json:"NSOpenGLProfileVersion4_1Core"`
			NSOpenGLProfileVersionLegacy                          int64 `json:"NSOpenGLProfileVersionLegacy"`
			NSOtherMouseDown                                      int64 `json:"NSOtherMouseDown"`
			NSOtherMouseDownMask                                  int64 `json:"NSOtherMouseDownMask"`
			NSOtherMouseDragged                                   int64 `json:"NSOtherMouseDragged"`
			NSOtherMouseDraggedMask                               int64 `json:"NSOtherMouseDraggedMask"`
			NSOtherMouseUp                                        int64 `json:"NSOtherMouseUp"`
			NSOtherMouseUpMask                                    int64 `json:"NSOtherMouseUpMask"`
			NSOtherTextMovement                                   int64 `json:"NSOtherTextMovement"`
			NSOutlineViewDropOnItemIndex                          int64 `json:"NSOutlineViewDropOnItemIndex"`
			NSPDFPanelRequestsParentDirectory                     int64 `json:"NSPDFPanelRequestsParentDirectory"`
			NSPDFPanelShowsOrientation                            int64 `json:"NSPDFPanelShowsOrientation"`
			NSPDFPanelShowsPaperSize                              int64 `json:"NSPDFPanelShowsPaperSize"`
			NSPNGFileType                                         int64 `json:"NSPNGFileType"`
			NSPageControllerTransitionStyleHorizontalStrip        int64 `json:"NSPageControllerTransitionStyleHorizontalStrip"`
			NSPageControllerTransitionStyleStackBook              int64 `json:"NSPageControllerTransitionStyleStackBook"`
			NSPageControllerTransitionStyleStackHistory           int64 `json:"NSPageControllerTransitionStyleStackHistory"`
			NSPageDownFunctionKey                                 int64 `json:"NSPageDownFunctionKey"`
			NSPageUpFunctionKey                                   int64 `json:"NSPageUpFunctionKey"`
			NSPaperOrientationLandscape                           int64 `json:"NSPaperOrientationLandscape"`
			NSPaperOrientationPortrait                            int64 `json:"NSPaperOrientationPortrait"`
			NSParagraphSeparatorCharacter                         int64 `json:"NSParagraphSeparatorCharacter"`
			NSPasteboardReadingAsData                             int64 `json:"NSPasteboardReadingAsData"`
			NSPasteboardReadingAsKeyedArchive                     int64 `json:"NSPasteboardReadingAsKeyedArchive"`
			NSPasteboardReadingAsPropertyList                     int64 `json:"NSPasteboardReadingAsPropertyList"`
			NSPasteboardReadingAsString                           int64 `json:"NSPasteboardReadingAsString"`
			NSPasteboardWritingPromised                           int64 `json:"NSPasteboardWritingPromised"`
			NSPathStyleNavigationBar                              int64 `json:"NSPathStyleNavigationBar"`
			NSPathStylePopUp                                      int64 `json:"NSPathStylePopUp"`
			NSPathStyleStandard                                   int64 `json:"NSPathStyleStandard"`
			NSPatternColorSpaceModel                              int64 `json:"NSPatternColorSpaceModel"`
			NSPauseFunctionKey                                    int64 `json:"NSPauseFunctionKey"`
			NSPenLowerSideMask                                    int64 `json:"NSPenLowerSideMask"`
			NSPenPointingDevice                                   int64 `json:"NSPenPointingDevice"`
			NSPenTipMask                                          int64 `json:"NSPenTipMask"`
			NSPenUpperSideMask                                    int64 `json:"NSPenUpperSideMask"`
			NSPeriodic                                            int64 `json:"NSPeriodic"`
			NSPeriodicMask                                        int64 `json:"NSPeriodicMask"`
			NSPopUpArrowAtBottom                                  int64 `json:"NSPopUpArrowAtBottom"`
			NSPopUpArrowAtCenter                                  int64 `json:"NSPopUpArrowAtCenter"`
			NSPopUpNoArrow                                        int64 `json:"NSPopUpNoArrow"`
			NSPopoverAppearanceHUD                                int64 `json:"NSPopoverAppearanceHUD"`
			NSPopoverAppearanceMinimal                            int64 `json:"NSPopoverAppearanceMinimal"`
			NSPopoverBehaviorApplicationDefined                   int64 `json:"NSPopoverBehaviorApplicationDefined"`
			NSPopoverBehaviorSemitransient                        int64 `json:"NSPopoverBehaviorSemitransient"`
			NSPopoverBehaviorTransient                            int64 `json:"NSPopoverBehaviorTransient"`
			NSPortraitOrientation                                 int64 `json:"NSPortraitOrientation"`
			NSPositiveDoubleType                                  int64 `json:"NSPositiveDoubleType"`
			NSPositiveFloatType                                   int64 `json:"NSPositiveFloatType"`
			NSPositiveIntType                                     int64 `json:"NSPositiveIntType"`
			NSPosterFontMask                                      int64 `json:"NSPosterFontMask"`
			NSPowerOffEventType                                   int64 `json:"NSPowerOffEventType"`
			NSPressedTab                                          int64 `json:"NSPressedTab"`
			NSPrevFunctionKey                                     int64 `json:"NSPrevFunctionKey"`
			NSPrintFunctionKey                                    int64 `json:"NSPrintFunctionKey"`
			NSPrintPanelShowsCopies                               int64 `json:"NSPrintPanelShowsCopies"`
			NSPrintPanelShowsOrientation                          int64 `json:"NSPrintPanelShowsOrientation"`
			NSPrintPanelShowsPageRange                            int64 `json:"NSPrintPanelShowsPageRange"`
			NSPrintPanelShowsPageSetupAccessory                   int64 `json:"NSPrintPanelShowsPageSetupAccessory"`
			NSPrintPanelShowsPaperSize                            int64 `json:"NSPrintPanelShowsPaperSize"`
			NSPrintPanelShowsPreview                              int64 `json:"NSPrintPanelShowsPreview"`
			NSPrintPanelShowsPrintSelection                       int64 `json:"NSPrintPanelShowsPrintSelection"`
			NSPrintPanelShowsScaling                              int64 `json:"NSPrintPanelShowsScaling"`
			NSPrintRenderingQualityBest                           int64 `json:"NSPrintRenderingQualityBest"`
			NSPrintRenderingQualityResponsive                     int64 `json:"NSPrintRenderingQualityResponsive"`
			NSPrintScreenFunctionKey                              int64 `json:"NSPrintScreenFunctionKey"`
			NSPrinterTableError                                   int64 `json:"NSPrinterTableError"`
			NSPrinterTableNotFound                                int64 `json:"NSPrinterTableNotFound"`
			NSPrinterTableOK                                      int64 `json:"NSPrinterTableOK"`
			NSPrintingCancelled                                   int64 `json:"NSPrintingCancelled"`
			NSPrintingFailure                                     int64 `json:"NSPrintingFailure"`
			NSPrintingReplyLater                                  int64 `json:"NSPrintingReplyLater"`
			NSPrintingSuccess                                     int64 `json:"NSPrintingSuccess"`
			NSProgressIndicatorBarStyle                           int64 `json:"NSProgressIndicatorBarStyle"`
			NSProgressIndicatorPreferredAquaThickness             int64 `json:"NSProgressIndicatorPreferredAquaThickness"`
			NSProgressIndicatorPreferredLargeThickness            int64 `json:"NSProgressIndicatorPreferredLargeThickness"`
			NSProgressIndicatorPreferredSmallThickness            int64 `json:"NSProgressIndicatorPreferredSmallThickness"`
			NSProgressIndicatorPreferredThickness                 int64 `json:"NSProgressIndicatorPreferredThickness"`
			NSProgressIndicatorSpinningStyle                      int64 `json:"NSProgressIndicatorSpinningStyle"`
			NSPushInCell                                          int64 `json:"NSPushInCell"`
			NSPushInCellMask                                      int64 `json:"NSPushInCellMask"`
			NSPushOnPushOffButton                                 int64 `json:"NSPushOnPushOffButton"`
			NSRGBColorSpaceModel                                  int64 `json:"NSRGBColorSpaceModel"`
			NSRGBModeColorPanel                                   int64 `json:"NSRGBModeColorPanel"`
			NSRadioButton                                         int64 `json:"NSRadioButton"`
			NSRadioModeMatrix                                     int64 `json:"NSRadioModeMatrix"`
			NSRangeDateMode                                       int64 `json:"NSRangeDateMode"`
			NSRatingLevelIndicatorStyle                           int64 `json:"NSRatingLevelIndicatorStyle"`
			NSRecessedBezelStyle                                  int64 `json:"NSRecessedBezelStyle"`
			NSRedoFunctionKey                                     int64 `json:"NSRedoFunctionKey"`
			NSRegularControlSize                                  int64 `json:"NSRegularControlSize"`
			NSRegularSquareBezelStyle                             int64 `json:"NSRegularSquareBezelStyle"`
			NSRelevancyLevelIndicatorStyle                        int64 `json:"NSRelevancyLevelIndicatorStyle"`
			NSRemoteNotificationTypeAlert                         int64 `json:"NSRemoteNotificationTypeAlert"`
			NSRemoteNotificationTypeBadge                         int64 `json:"NSRemoteNotificationTypeBadge"`
			NSRemoteNotificationTypeNone                          int64 `json:"NSRemoteNotificationTypeNone"`
			NSRemoteNotificationTypeSound                         int64 `json:"NSRemoteNotificationTypeSound"`
			NSRemoveTraitFontAction                               int64 `json:"NSRemoveTraitFontAction"`
			NSResetCursorRectsRunLoopOrdering                     int64 `json:"NSResetCursorRectsRunLoopOrdering"`
			NSResetFunctionKey                                    int64 `json:"NSResetFunctionKey"`
			NSResizableWindowMask                                 int64 `json:"NSResizableWindowMask"`
			NSReturnTextMovement                                  int64 `json:"NSReturnTextMovement"`
			NSRightArrowFunctionKey                               int64 `json:"NSRightArrowFunctionKey"`
			NSRightMouseDown                                      int64 `json:"NSRightMouseDown"`
			NSRightMouseDownMask                                  int64 `json:"NSRightMouseDownMask"`
			NSRightMouseDragged                                   int64 `json:"NSRightMouseDragged"`
			NSRightMouseDraggedMask                               int64 `json:"NSRightMouseDraggedMask"`
			NSRightMouseUp                                        int64 `json:"NSRightMouseUp"`
			NSRightMouseUpMask                                    int64 `json:"NSRightMouseUpMask"`
			NSRightTabStopType                                    int64 `json:"NSRightTabStopType"`
			NSRightTabsBezelBorder                                int64 `json:"NSRightTabsBezelBorder"`
			NSRightTextAlignment                                  int64 `json:"NSRightTextAlignment"`
			NSRightTextMovement                                   int64 `json:"NSRightTextMovement"`
			NSRoundLineCapStyle                                   int64 `json:"NSRoundLineCapStyle"`
			NSRoundLineJoinStyle                                  int64 `json:"NSRoundLineJoinStyle"`
			NSRoundRectBezelStyle                                 int64 `json:"NSRoundRectBezelStyle"`
			NSRoundedBezelStyle                                   int64 `json:"NSRoundedBezelStyle"`
			NSRoundedDisclosureBezelStyle                         int64 `json:"NSRoundedDisclosureBezelStyle"`
			NSRuleEditorNestingModeCompound                       int64 `json:"NSRuleEditorNestingModeCompound"`
			NSRuleEditorNestingModeList                           int64 `json:"NSRuleEditorNestingModeList"`
			NSRuleEditorNestingModeSimple                         int64 `json:"NSRuleEditorNestingModeSimple"`
			NSRuleEditorNestingModeSingle                         int64 `json:"NSRuleEditorNestingModeSingle"`
			NSRuleEditorRowTypeCompound                           int64 `json:"NSRuleEditorRowTypeCompound"`
			NSRuleEditorRowTypeSimple                             int64 `json:"NSRuleEditorRowTypeSimple"`
			NSRunAbortedResponse                                  int64 `json:"NSRunAbortedResponse"`
			NSRunContinuesResponse                                int64 `json:"NSRunContinuesResponse"`
			NSRunStoppedResponse                                  int64 `json:"NSRunStoppedResponse"`
			NSSaveAsOperation                                     int64 `json:"NSSaveAsOperation"`
			NSSaveOperation                                       int64 `json:"NSSaveOperation"`
			NSSaveToOperation                                     int64 `json:"NSSaveToOperation"`
			NSScaleNone                                           int64 `json:"NSScaleNone"`
			NSScaleProportionally                                 int64 `json:"NSScaleProportionally"`
			NSScaleToFit                                          int64 `json:"NSScaleToFit"`
			NSScreenChangedEventType                              int64 `json:"NSScreenChangedEventType"`
			NSScrollElasticityAllowed                             int64 `json:"NSScrollElasticityAllowed"`
			NSScrollElasticityAutomatic                           int64 `json:"NSScrollElasticityAutomatic"`
			NSScrollElasticityNone                                int64 `json:"NSScrollElasticityNone"`
			NSScrollLockFunctionKey                               int64 `json:"NSScrollLockFunctionKey"`
			NSScrollViewFindBarPositionAboveContent               int64 `json:"NSScrollViewFindBarPositionAboveContent"`
			NSScrollViewFindBarPositionAboveHorizontalRuler       int64 `json:"NSScrollViewFindBarPositionAboveHorizontalRuler"`
			NSScrollViewFindBarPositionBelowContent               int64 `json:"NSScrollViewFindBarPositionBelowContent"`
			NSScrollWheel                                         int64 `json:"NSScrollWheel"`
			NSScrollWheelMask                                     int64 `json:"NSScrollWheelMask"`
			NSScrollerArrowsDefaultSetting                        int64 `json:"NSScrollerArrowsDefaultSetting"`
			NSScrollerArrowsMaxEnd                                int64 `json:"NSScrollerArrowsMaxEnd"`
			NSScrollerArrowsMinEnd                                int64 `json:"NSScrollerArrowsMinEnd"`
			NSScrollerArrowsNone                                  int64 `json:"NSScrollerArrowsNone"`
			NSScrollerDecrementArrow                              int64 `json:"NSScrollerDecrementArrow"`
			NSScrollerDecrementLine                               int64 `json:"NSScrollerDecrementLine"`
			NSScrollerDecrementPage                               int64 `json:"NSScrollerDecrementPage"`
			NSScrollerIncrementArrow                              int64 `json:"NSScrollerIncrementArrow"`
			NSScrollerIncrementLine                               int64 `json:"NSScrollerIncrementLine"`
			NSScrollerIncrementPage                               int64 `json:"NSScrollerIncrementPage"`
			NSScrollerKnob                                        int64 `json:"NSScrollerKnob"`
			NSScrollerKnobSlot                                    int64 `json:"NSScrollerKnobSlot"`
			NSScrollerKnobStyleDark                               int64 `json:"NSScrollerKnobStyleDark"`
			NSScrollerKnobStyleDefault                            int64 `json:"NSScrollerKnobStyleDefault"`
			NSScrollerKnobStyleLight                              int64 `json:"NSScrollerKnobStyleLight"`
			NSScrollerNoPart                                      int64 `json:"NSScrollerNoPart"`
			NSScrollerStyleLegacy                                 int64 `json:"NSScrollerStyleLegacy"`
			NSScrollerStyleOverlay                                int64 `json:"NSScrollerStyleOverlay"`
			NSSearchFieldClearRecentsMenuItemTag                  int64 `json:"NSSearchFieldClearRecentsMenuItemTag"`
			NSSearchFieldNoRecentsMenuItemTag                     int64 `json:"NSSearchFieldNoRecentsMenuItemTag"`
			NSSearchFieldRecentsMenuItemTag                       int64 `json:"NSSearchFieldRecentsMenuItemTag"`
			NSSearchFieldRecentsTitleMenuItemTag                  int64 `json:"NSSearchFieldRecentsTitleMenuItemTag"`
			NSSegmentStyleAutomatic                               int64 `json:"NSSegmentStyleAutomatic"`
			NSSegmentStyleCapsule                                 int64 `json:"NSSegmentStyleCapsule"`
			NSSegmentStyleRoundRect                               int64 `json:"NSSegmentStyleRoundRect"`
			NSSegmentStyleRounded                                 int64 `json:"NSSegmentStyleRounded"`
			NSSegmentStyleSeparated                               int64 `json:"NSSegmentStyleSeparated"`
			NSSegmentStyleSmallSquare                             int64 `json:"NSSegmentStyleSmallSquare"`
			NSSegmentStyleTexturedRounded                         int64 `json:"NSSegmentStyleTexturedRounded"`
			NSSegmentStyleTexturedSquare                          int64 `json:"NSSegmentStyleTexturedSquare"`
			NSSegmentSwitchTrackingMomentary                      int64 `json:"NSSegmentSwitchTrackingMomentary"`
			NSSegmentSwitchTrackingMomentaryAccelerator           int64 `json:"NSSegmentSwitchTrackingMomentaryAccelerator"`
			NSSegmentSwitchTrackingSelectAny                      int64 `json:"NSSegmentSwitchTrackingSelectAny"`
			NSSegmentSwitchTrackingSelectOne                      int64 `json:"NSSegmentSwitchTrackingSelectOne"`
			NSSelectByCharacter                                   int64 `json:"NSSelectByCharacter"`
			NSSelectByParagraph                                   int64 `json:"NSSelectByParagraph"`
			NSSelectByWord                                        int64 `json:"NSSelectByWord"`
			NSSelectFunctionKey                                   int64 `json:"NSSelectFunctionKey"`
			NSSelectedTab                                         int64 `json:"NSSelectedTab"`
			NSSelectingNext                                       int64 `json:"NSSelectingNext"`
			NSSelectingPrevious                                   int64 `json:"NSSelectingPrevious"`
			NSSelectionAffinityDownstream                         int64 `json:"NSSelectionAffinityDownstream"`
			NSSelectionAffinityUpstream                           int64 `json:"NSSelectionAffinityUpstream"`
			NSServiceApplicationLaunchFailedError                 int64 `json:"NSServiceApplicationLaunchFailedError"`
			NSServiceApplicationNotFoundError                     int64 `json:"NSServiceApplicationNotFoundError"`
			NSServiceErrorMaximum                                 int64 `json:"NSServiceErrorMaximum"`
			NSServiceErrorMinimum                                 int64 `json:"NSServiceErrorMinimum"`
			NSServiceInvalidPasteboardDataError                   int64 `json:"NSServiceInvalidPasteboardDataError"`
			NSServiceMalformedServiceDictionaryError              int64 `json:"NSServiceMalformedServiceDictionaryError"`
			NSServiceMiscellaneousError                           int64 `json:"NSServiceMiscellaneousError"`
			NSServiceRequestTimedOutError                         int64 `json:"NSServiceRequestTimedOutError"`
			NSShadowlessSquareBezelStyle                          int64 `json:"NSShadowlessSquareBezelStyle"`
			NSSharingContentScopeFull                             int64 `json:"NSSharingContentScopeFull"`
			NSSharingContentScopeItem                             int64 `json:"NSSharingContentScopeItem"`
			NSSharingContentScopePartial                          int64 `json:"NSSharingContentScopePartial"`
			NSSharingServiceErrorMaximum                          int64 `json:"NSSharingServiceErrorMaximum"`
			NSSharingServiceErrorMinimum                          int64 `json:"NSSharingServiceErrorMinimum"`
			NSSharingServiceNotConfiguredError                    int64 `json:"NSSharingServiceNotConfiguredError"`
			NSShiftKeyMask                                        int64 `json:"NSShiftKeyMask"`
			NSShowControlGlyphs                                   int64 `json:"NSShowControlGlyphs"`
			NSShowInvisibleGlyphs                                 int64 `json:"NSShowInvisibleGlyphs"`
			NSSingleDateMode                                      int64 `json:"NSSingleDateMode"`
			NSSingleUnderlineStyle                                int64 `json:"NSSingleUnderlineStyle"`
			NSSizeDownFontAction                                  int64 `json:"NSSizeDownFontAction"`
			NSSizeUpFontAction                                    int64 `json:"NSSizeUpFontAction"`
			NSSmallCapsFontMask                                   int64 `json:"NSSmallCapsFontMask"`
			NSSmallControlSize                                    int64 `json:"NSSmallControlSize"`
			NSSmallIconButtonBezelStyle                           int64 `json:"NSSmallIconButtonBezelStyle"`
			NSSmallSquareBezelStyle                               int64 `json:"NSSmallSquareBezelStyle"`
			NSSpecialPageOrder                                    int64 `json:"NSSpecialPageOrder"`
			NSSpeechImmediateBoundary                             int64 `json:"NSSpeechImmediateBoundary"`
			NSSpeechSentenceBoundary                              int64 `json:"NSSpeechSentenceBoundary"`
			NSSpeechWordBoundary                                  int64 `json:"NSSpeechWordBoundary"`
			NSSpellingStateGrammarFlag                            int64 `json:"NSSpellingStateGrammarFlag"`
			NSSpellingStateSpellingFlag                           int64 `json:"NSSpellingStateSpellingFlag"`
			NSSplitViewDividerStylePaneSplitter                   int64 `json:"NSSplitViewDividerStylePaneSplitter"`
			NSSplitViewDividerStyleThick                          int64 `json:"NSSplitViewDividerStyleThick"`
			NSSplitViewDividerStyleThin                           int64 `json:"NSSplitViewDividerStyleThin"`
			NSSquareLineCapStyle                                  int64 `json:"NSSquareLineCapStyle"`
			NSStackViewGravityBottom                              int64 `json:"NSStackViewGravityBottom"`
			NSStackViewGravityCenter                              int64 `json:"NSStackViewGravityCenter"`
			NSStackViewGravityLeading                             int64 `json:"NSStackViewGravityLeading"`
			NSStackViewGravityTop                                 int64 `json:"NSStackViewGravityTop"`
			NSStackViewGravityTrailing                            int64 `json:"NSStackViewGravityTrailing"`
			NSStopFunctionKey                                     int64 `json:"NSStopFunctionKey"`
			NSStringDrawingDisableScreenFontSubstitution          int64 `json:"NSStringDrawingDisableScreenFontSubstitution"`
			NSStringDrawingOneShot                                int64 `json:"NSStringDrawingOneShot"`
			NSStringDrawingTruncatesLastVisibleLine               int64 `json:"NSStringDrawingTruncatesLastVisibleLine"`
			NSStringDrawingUsesDeviceMetrics                      int64 `json:"NSStringDrawingUsesDeviceMetrics"`
			NSStringDrawingUsesFontLeading                        int64 `json:"NSStringDrawingUsesFontLeading"`
			NSStringDrawingUsesLineFragmentOrigin                 int64 `json:"NSStringDrawingUsesLineFragmentOrigin"`
			NSSwitchButton                                        int64 `json:"NSSwitchButton"`
			NSSysReqFunctionKey                                   int64 `json:"NSSysReqFunctionKey"`
			NSSystemDefined                                       int64 `json:"NSSystemDefined"`
			NSSystemDefinedMask                                   int64 `json:"NSSystemDefinedMask"`
			NSSystemFunctionKey                                   int64 `json:"NSSystemFunctionKey"`
			NSTIFFCompressionCCITTFAX3                            int64 `json:"NSTIFFCompressionCCITTFAX3"`
			NSTIFFCompressionCCITTFAX4                            int64 `json:"NSTIFFCompressionCCITTFAX4"`
			NSTIFFCompressionJPEG                                 int64 `json:"NSTIFFCompressionJPEG"`
			NSTIFFCompressionLZW                                  int64 `json:"NSTIFFCompressionLZW"`
			NSTIFFCompressionNEXT                                 int64 `json:"NSTIFFCompressionNEXT"`
			NSTIFFCompressionNone                                 int64 `json:"NSTIFFCompressionNone"`
			NSTIFFCompressionOldJPEG                              int64 `json:"NSTIFFCompressionOldJPEG"`
			NSTIFFCompressionPackBits                             int64 `json:"NSTIFFCompressionPackBits"`
			NSTIFFFileType                                        int64 `json:"NSTIFFFileType"`
			NSTabCharacter                                        int64 `json:"NSTabCharacter"`
			NSTabTextMovement                                     int64 `json:"NSTabTextMovement"`
			NSTabViewControllerTabStyleSegmentedControlOnBottom   int64 `json:"NSTabViewControllerTabStyleSegmentedControlOnBottom"`
			NSTabViewControllerTabStyleSegmentedControlOnTop      int64 `json:"NSTabViewControllerTabStyleSegmentedControlOnTop"`
			NSTabViewControllerTabStyleToolbar                    int64 `json:"NSTabViewControllerTabStyleToolbar"`
			NSTabViewControllerTabStyleUnspecified                int64 `json:"NSTabViewControllerTabStyleUnspecified"`
			NSTableColumnAutoresizingMask                         int64 `json:"NSTableColumnAutoresizingMask"`
			NSTableColumnNoResizing                               int64 `json:"NSTableColumnNoResizing"`
			NSTableColumnUserResizingMask                         int64 `json:"NSTableColumnUserResizingMask"`
			NSTableViewAnimationEffectFade                        int64 `json:"NSTableViewAnimationEffectFade"`
			NSTableViewAnimationEffectGap                         int64 `json:"NSTableViewAnimationEffectGap"`
			NSTableViewAnimationEffectNone                        int64 `json:"NSTableViewAnimationEffectNone"`
			NSTableViewAnimationSlideDown                         int64 `json:"NSTableViewAnimationSlideDown"`
			NSTableViewAnimationSlideLeft                         int64 `json:"NSTableViewAnimationSlideLeft"`
			NSTableViewAnimationSlideRight                        int64 `json:"NSTableViewAnimationSlideRight"`
			NSTableViewAnimationSlideUp                           int64 `json:"NSTableViewAnimationSlideUp"`
			NSTableViewDashedHorizontalGridLineMask               int64 `json:"NSTableViewDashedHorizontalGridLineMask"`
			NSTableViewDraggingDestinationFeedbackStyleGap        int64 `json:"NSTableViewDraggingDestinationFeedbackStyleGap"`
			NSTableViewDraggingDestinationFeedbackStyleNone       int64 `json:"NSTableViewDraggingDestinationFeedbackStyleNone"`
			NSTableViewDraggingDestinationFeedbackStyleRegular    int64 `json:"NSTableViewDraggingDestinationFeedbackStyleRegular"`
			NSTableViewDraggingDestinationFeedbackStyleSourceList int64 `json:"NSTableViewDraggingDestinationFeedbackStyleSourceList"`
			NSTableViewDropAbove                                  int64 `json:"NSTableViewDropAbove"`
			NSTableViewDropOn                                     int64 `json:"NSTableViewDropOn"`
			NSTableViewFirstColumnOnlyAutoresizingStyle           int64 `json:"NSTableViewFirstColumnOnlyAutoresizingStyle"`
			NSTableViewGridNone                                   int64 `json:"NSTableViewGridNone"`
			NSTableViewLastColumnOnlyAutoresizingStyle            int64 `json:"NSTableViewLastColumnOnlyAutoresizingStyle"`
			NSTableViewNoColumnAutoresizing                       int64 `json:"NSTableViewNoColumnAutoresizing"`
			NSTableViewReverseSequentialColumnAutoresizingStyle   int64 `json:"NSTableViewReverseSequentialColumnAutoresizingStyle"`
			NSTableViewRowSizeStyleCustom                         int64 `json:"NSTableViewRowSizeStyleCustom"`
			NSTableViewRowSizeStyleDefault                        int64 `json:"NSTableViewRowSizeStyleDefault"`
			NSTableViewRowSizeStyleLarge                          int64 `json:"NSTableViewRowSizeStyleLarge"`
			NSTableViewRowSizeStyleMedium                         int64 `json:"NSTableViewRowSizeStyleMedium"`
			NSTableViewRowSizeStyleSmall                          int64 `json:"NSTableViewRowSizeStyleSmall"`
			NSTableViewSelectionHighlightStyleNone                int64 `json:"NSTableViewSelectionHighlightStyleNone"`
			NSTableViewSelectionHighlightStyleRegular             int64 `json:"NSTableViewSelectionHighlightStyleRegular"`
			NSTableViewSelectionHighlightStyleSourceList          int64 `json:"NSTableViewSelectionHighlightStyleSourceList"`
			NSTableViewSequentialColumnAutoresizingStyle          int64 `json:"NSTableViewSequentialColumnAutoresizingStyle"`
			NSTableViewSolidHorizontalGridLineMask                int64 `json:"NSTableViewSolidHorizontalGridLineMask"`
			NSTableViewSolidVerticalGridLineMask                  int64 `json:"NSTableViewSolidVerticalGridLineMask"`
			NSTableViewUniformColumnAutoresizingStyle             int64 `json:"NSTableViewUniformColumnAutoresizingStyle"`
			NSTabletPoint                                         int64 `json:"NSTabletPoint"`
			NSTabletPointEventSubtype                             int64 `json:"NSTabletPointEventSubtype"`
			NSTabletPointMask                                     int64 `json:"NSTabletPointMask"`
			NSTabletProximity                                     int64 `json:"NSTabletProximity"`
			NSTabletProximityEventSubtype                         int64 `json:"NSTabletProximityEventSubtype"`
			NSTabletProximityMask                                 int64 `json:"NSTabletProximityMask"`
			NSTerminateCancel                                     int64 `json:"NSTerminateCancel"`
			NSTerminateLater                                      int64 `json:"NSTerminateLater"`
			NSTerminateNow                                        int64 `json:"NSTerminateNow"`
			NSTextBlockAbsoluteValueType                          int64 `json:"NSTextBlockAbsoluteValueType"`
			NSTextBlockBaselineAlignment                          int64 `json:"NSTextBlockBaselineAlignment"`
			NSTextBlockBorder                                     int64 `json:"NSTextBlockBorder"`
			NSTextBlockBottomAlignment                            int64 `json:"NSTextBlockBottomAlignment"`
			NSTextBlockHeight                                     int64 `json:"NSTextBlockHeight"`
			NSTextBlockMargin                                     int64 `json:"NSTextBlockMargin"`
			NSTextBlockMaximumHeight                              int64 `json:"NSTextBlockMaximumHeight"`
			NSTextBlockMaximumWidth                               int64 `json:"NSTextBlockMaximumWidth"`
			NSTextBlockMiddleAlignment                            int64 `json:"NSTextBlockMiddleAlignment"`
			NSTextBlockMinimumHeight                              int64 `json:"NSTextBlockMinimumHeight"`
			NSTextBlockMinimumWidth                               int64 `json:"NSTextBlockMinimumWidth"`
			NSTextBlockPadding                                    int64 `json:"NSTextBlockPadding"`
			NSTextBlockPercentageValueType                        int64 `json:"NSTextBlockPercentageValueType"`
			NSTextBlockTopAlignment                               int64 `json:"NSTextBlockTopAlignment"`
			NSTextBlockWidth                                      int64 `json:"NSTextBlockWidth"`
			NSTextCellType                                        int64 `json:"NSTextCellType"`
			NSTextFieldAndStepperDatePickerStyle                  int64 `json:"NSTextFieldAndStepperDatePickerStyle"`
			NSTextFieldDatePickerStyle                            int64 `json:"NSTextFieldDatePickerStyle"`
			NSTextFieldRoundedBezel                               int64 `json:"NSTextFieldRoundedBezel"`
			NSTextFieldSquareBezel                                int64 `json:"NSTextFieldSquareBezel"`
			NSTextFinderActionHideFindInterface                   int64 `json:"NSTextFinderActionHideFindInterface"`
			NSTextFinderActionHideReplaceInterface                int64 `json:"NSTextFinderActionHideReplaceInterface"`
			NSTextFinderActionNextMatch                           int64 `json:"NSTextFinderActionNextMatch"`
			NSTextFinderActionPreviousMatch                       int64 `json:"NSTextFinderActionPreviousMatch"`
			NSTextFinderActionReplace                             int64 `json:"NSTextFinderActionReplace"`
			NSTextFinderActionReplaceAll                          int64 `json:"NSTextFinderActionReplaceAll"`
			NSTextFinderActionReplaceAllInSelection               int64 `json:"NSTextFinderActionReplaceAllInSelection"`
			NSTextFinderActionReplaceAndFind                      int64 `json:"NSTextFinderActionReplaceAndFind"`
			NSTextFinderActionSelectAll                           int64 `json:"NSTextFinderActionSelectAll"`
			NSTextFinderActionSelectAllInSelection                int64 `json:"NSTextFinderActionSelectAllInSelection"`
			NSTextFinderActionSetSearchString                     int64 `json:"NSTextFinderActionSetSearchString"`
			NSTextFinderActionShowFindInterface                   int64 `json:"NSTextFinderActionShowFindInterface"`
			NSTextFinderActionShowReplaceInterface                int64 `json:"NSTextFinderActionShowReplaceInterface"`
			NSTextFinderMatchingTypeContains                      int64 `json:"NSTextFinderMatchingTypeContains"`
			NSTextFinderMatchingTypeEndsWith                      int64 `json:"NSTextFinderMatchingTypeEndsWith"`
			NSTextFinderMatchingTypeFullWord                      int64 `json:"NSTextFinderMatchingTypeFullWord"`
			NSTextFinderMatchingTypeStartsWith                    int64 `json:"NSTextFinderMatchingTypeStartsWith"`
			NSTextLayoutOrientationHorizontal                     int64 `json:"NSTextLayoutOrientationHorizontal"`
			NSTextLayoutOrientationVertical                       int64 `json:"NSTextLayoutOrientationVertical"`
			NSTextListPrependEnclosingMarker                      int64 `json:"NSTextListPrependEnclosingMarker"`
			NSTextReadInapplicableDocumentTypeError               int64 `json:"NSTextReadInapplicableDocumentTypeError"`
			NSTextReadWriteErrorMaximum                           int64 `json:"NSTextReadWriteErrorMaximum"`
			NSTextReadWriteErrorMinimum                           int64 `json:"NSTextReadWriteErrorMinimum"`
			NSTextStorageEditedAttributes                         int64 `json:"NSTextStorageEditedAttributes"`
			NSTextStorageEditedCharacters                         int64 `json:"NSTextStorageEditedCharacters"`
			NSTextTableAutomaticLayoutAlgorithm                   int64 `json:"NSTextTableAutomaticLayoutAlgorithm"`
			NSTextTableFixedLayoutAlgorithm                       int64 `json:"NSTextTableFixedLayoutAlgorithm"`
			NSTextWriteInapplicableDocumentTypeError              int64 `json:"NSTextWriteInapplicableDocumentTypeError"`
			NSTextWritingDirectionEmbedding                       int64 `json:"NSTextWritingDirectionEmbedding"`
			NSTextWritingDirectionOverride                        int64 `json:"NSTextWritingDirectionOverride"`
			NSTexturedBackgroundWindowMask                        int64 `json:"NSTexturedBackgroundWindowMask"`
			NSTexturedRoundedBezelStyle                           int64 `json:"NSTexturedRoundedBezelStyle"`
			NSTexturedSquareBezelStyle                            int64 `json:"NSTexturedSquareBezelStyle"`
			NSThickSquareBezelStyle                               int64 `json:"NSThickSquareBezelStyle"`
			NSThickerSquareBezelStyle                             int64 `json:"NSThickerSquareBezelStyle"`
			NSTickMarkAbove                                       int64 `json:"NSTickMarkAbove"`
			NSTickMarkBelow                                       int64 `json:"NSTickMarkBelow"`
			NSTickMarkLeft                                        int64 `json:"NSTickMarkLeft"`
			NSTickMarkRight                                       int64 `json:"NSTickMarkRight"`
			NSTimeZoneDatePickerElementFlag                       int64 `json:"NSTimeZoneDatePickerElementFlag"`
			NSTitledWindowMask                                    int64 `json:"NSTitledWindowMask"`
			NSToggleButton                                        int64 `json:"NSToggleButton"`
			NSTokenStyleDefault                                   int64 `json:"NSTokenStyleDefault"`
			NSTokenStyleNone                                      int64 `json:"NSTokenStyleNone"`
			NSTokenStylePlainSquared                              int64 `json:"NSTokenStylePlainSquared"`
			NSTokenStyleRounded                                   int64 `json:"NSTokenStyleRounded"`
			NSTokenStyleSquared                                   int64 `json:"NSTokenStyleSquared"`
			NSToolbarDisplayModeDefault                           int64 `json:"NSToolbarDisplayModeDefault"`
			NSToolbarDisplayModeIconAndLabel                      int64 `json:"NSToolbarDisplayModeIconAndLabel"`
			NSToolbarDisplayModeIconOnly                          int64 `json:"NSToolbarDisplayModeIconOnly"`
			NSToolbarDisplayModeLabelOnly                         int64 `json:"NSToolbarDisplayModeLabelOnly"`
			NSToolbarItemVisibilityPriorityHigh                   int64 `json:"NSToolbarItemVisibilityPriorityHigh"`
			NSToolbarItemVisibilityPriorityLow                    int64 `json:"NSToolbarItemVisibilityPriorityLow"`
			NSToolbarItemVisibilityPriorityStandard               int64 `json:"NSToolbarItemVisibilityPriorityStandard"`
			NSToolbarItemVisibilityPriorityUser                   int64 `json:"NSToolbarItemVisibilityPriorityUser"`
			NSToolbarSizeModeDefault                              int64 `json:"NSToolbarSizeModeDefault"`
			NSToolbarSizeModeRegular                              int64 `json:"NSToolbarSizeModeRegular"`
			NSToolbarSizeModeSmall                                int64 `json:"NSToolbarSizeModeSmall"`
			NSTopTabsBezelBorder                                  int64 `json:"NSTopTabsBezelBorder"`
			NSTouchEventSubtype                                   int64 `json:"NSTouchEventSubtype"`
			NSTouchPhaseAny                                       int64 `json:"NSTouchPhaseAny"`
			NSTouchPhaseBegan                                     int64 `json:"NSTouchPhaseBegan"`
			NSTouchPhaseCancelled                                 int64 `json:"NSTouchPhaseCancelled"`
			NSTouchPhaseEnded                                     int64 `json:"NSTouchPhaseEnded"`
			NSTouchPhaseMoved                                     int64 `json:"NSTouchPhaseMoved"`
			NSTouchPhaseStationary                                int64 `json:"NSTouchPhaseStationary"`
			NSTouchPhaseTouching                                  int64 `json:"NSTouchPhaseTouching"`
			NSTrackModeMatrix                                     int64 `json:"NSTrackModeMatrix"`
			NSTrackingActiveAlways                                int64 `json:"NSTrackingActiveAlways"`
			NSTrackingActiveInActiveApp                           int64 `json:"NSTrackingActiveInActiveApp"`
			NSTrackingActiveInKeyWindow                           int64 `json:"NSTrackingActiveInKeyWindow"`
			NSTrackingActiveWhenFirstResponder                    int64 `json:"NSTrackingActiveWhenFirstResponder"`
			NSTrackingAssumeInside                                int64 `json:"NSTrackingAssumeInside"`
			NSTrackingCursorUpdate                                int64 `json:"NSTrackingCursorUpdate"`
			NSTrackingEnabledDuringMouseDrag                      int64 `json:"NSTrackingEnabledDuringMouseDrag"`
			NSTrackingInVisibleRect                               int64 `json:"NSTrackingInVisibleRect"`
			NSTrackingMouseEnteredAndExited                       int64 `json:"NSTrackingMouseEnteredAndExited"`
			NSTrackingMouseMoved                                  int64 `json:"NSTrackingMouseMoved"`
			NSTypesetterBehavior10_2                              int64 `json:"NSTypesetterBehavior_10_2"`
			NSTypesetterBehavior10_2WithCompatibility             int64 `json:"NSTypesetterBehavior_10_2_WithCompatibility"`
			NSTypesetterBehavior10_3                              int64 `json:"NSTypesetterBehavior_10_3"`
			NSTypesetterBehavior10_4                              int64 `json:"NSTypesetterBehavior_10_4"`
			NSTypesetterContainerBreakAction                      int64 `json:"NSTypesetterContainerBreakAction"`
			NSTypesetterHorizontalTabAction                       int64 `json:"NSTypesetterHorizontalTabAction"`
			NSTypesetterLatestBehavior                            int64 `json:"NSTypesetterLatestBehavior"`
			NSTypesetterLineBreakAction                           int64 `json:"NSTypesetterLineBreakAction"`
			NSTypesetterOriginalBehavior                          int64 `json:"NSTypesetterOriginalBehavior"`
			NSTypesetterParagraphBreakAction                      int64 `json:"NSTypesetterParagraphBreakAction"`
			NSTypesetterWhitespaceAction                          int64 `json:"NSTypesetterWhitespaceAction"`
			NSTypesetterZeroAdvancementAction                     int64 `json:"NSTypesetterZeroAdvancementAction"`
			NSUnboldFontMask                                      int64 `json:"NSUnboldFontMask"`
			NSUnderlinePatternDash                                int64 `json:"NSUnderlinePatternDash"`
			NSUnderlinePatternDashDot                             int64 `json:"NSUnderlinePatternDashDot"`
			NSUnderlinePatternDashDotDot                          int64 `json:"NSUnderlinePatternDashDotDot"`
			NSUnderlinePatternDot                                 int64 `json:"NSUnderlinePatternDot"`
			NSUnderlinePatternSolid                               int64 `json:"NSUnderlinePatternSolid"`
			NSUnderlineStyleDouble                                int64 `json:"NSUnderlineStyleDouble"`
			NSUnderlineStyleNone                                  int64 `json:"NSUnderlineStyleNone"`
			NSUnderlineStyleSingle                                int64 `json:"NSUnderlineStyleSingle"`
			NSUnderlineStyleThick                                 int64 `json:"NSUnderlineStyleThick"`
			NSUndoFunctionKey                                     int64 `json:"NSUndoFunctionKey"`
			NSUnifiedTitleAndToolbarWindowMask                    int64 `json:"NSUnifiedTitleAndToolbarWindowMask"`
			NSUnitalicFontMask                                    int64 `json:"NSUnitalicFontMask"`
			NSUnknownColorSpaceModel                              int64 `json:"NSUnknownColorSpaceModel"`
			NSUnknownPageOrder                                    int64 `json:"NSUnknownPageOrder"`
			NSUnknownPointingDevice                               int64 `json:"NSUnknownPointingDevice"`
			NSUnscaledWindowMask                                  int64 `json:"NSUnscaledWindowMask"`
			NSUpArrowFunctionKey                                  int64 `json:"NSUpArrowFunctionKey"`
			NSUpTextMovement                                      int64 `json:"NSUpTextMovement"`
			NSUpdateWindowsRunLoopOrdering                        int64 `json:"NSUpdateWindowsRunLoopOrdering"`
			NSUserFunctionKey                                     int64 `json:"NSUserFunctionKey"`
			NSUserInterfaceLayoutDirectionLeftToRight             int64 `json:"NSUserInterfaceLayoutDirectionLeftToRight"`
			NSUserInterfaceLayoutDirectionRightToLeft             int64 `json:"NSUserInterfaceLayoutDirectionRightToLeft"`
			NSUserInterfaceLayoutOrientationHorizontal            int64 `json:"NSUserInterfaceLayoutOrientationHorizontal"`
			NSUserInterfaceLayoutOrientationVertical              int64 `json:"NSUserInterfaceLayoutOrientationVertical"`
			NSUtilityWindowMask                                   int64 `json:"NSUtilityWindowMask"`
			NSVerticalRuler                                       int64 `json:"NSVerticalRuler"`
			NSViaPanelFontAction                                  int64 `json:"NSViaPanelFontAction"`
			NSViewControllerTransitionAllowUserInteraction        int64 `json:"NSViewControllerTransitionAllowUserInteraction"`
			NSViewControllerTransitionCrossfade                   int64 `json:"NSViewControllerTransitionCrossfade"`
			NSViewControllerTransitionNone                        int64 `json:"NSViewControllerTransitionNone"`
			NSViewControllerTransitionSlideBackward               int64 `json:"NSViewControllerTransitionSlideBackward"`
			NSViewControllerTransitionSlideDown                   int64 `json:"NSViewControllerTransitionSlideDown"`
			NSViewControllerTransitionSlideForward                int64 `json:"NSViewControllerTransitionSlideForward"`
			NSViewControllerTransitionSlideLeft                   int64 `json:"NSViewControllerTransitionSlideLeft"`
			NSViewControllerTransitionSlideRight                  int64 `json:"NSViewControllerTransitionSlideRight"`
			NSViewControllerTransitionSlideUp                     int64 `json:"NSViewControllerTransitionSlideUp"`
			NSViewHeightSizable                                   int64 `json:"NSViewHeightSizable"`
			NSViewLayerContentsPlacementBottom                    int64 `json:"NSViewLayerContentsPlacementBottom"`
			NSViewLayerContentsPlacementBottomLeft                int64 `json:"NSViewLayerContentsPlacementBottomLeft"`
			NSViewLayerContentsPlacementBottomRight               int64 `json:"NSViewLayerContentsPlacementBottomRight"`
			NSViewLayerContentsPlacementCenter                    int64 `json:"NSViewLayerContentsPlacementCenter"`
			NSViewLayerContentsPlacementLeft                      int64 `json:"NSViewLayerContentsPlacementLeft"`
			NSViewLayerContentsPlacementRight                     int64 `json:"NSViewLayerContentsPlacementRight"`
			NSViewLayerContentsPlacementScaleAxesIndependently    int64 `json:"NSViewLayerContentsPlacementScaleAxesIndependently"`
			NSViewLayerContentsPlacementScaleProportionallyToFill int64 `json:"NSViewLayerContentsPlacementScaleProportionallyToFill"`
			NSViewLayerContentsPlacementScaleProportionallyToFit  int64 `json:"NSViewLayerContentsPlacementScaleProportionallyToFit"`
			NSViewLayerContentsPlacementTop                       int64 `json:"NSViewLayerContentsPlacementTop"`
			NSViewLayerContentsPlacementTopLeft                   int64 `json:"NSViewLayerContentsPlacementTopLeft"`
			NSViewLayerContentsPlacementTopRight                  int64 `json:"NSViewLayerContentsPlacementTopRight"`
			NSViewLayerContentsRedrawBeforeViewResize             int64 `json:"NSViewLayerContentsRedrawBeforeViewResize"`
			NSViewLayerContentsRedrawCrossfade                    int64 `json:"NSViewLayerContentsRedrawCrossfade"`
			NSViewLayerContentsRedrawDuringViewResize             int64 `json:"NSViewLayerContentsRedrawDuringViewResize"`
			NSViewLayerContentsRedrawNever                        int64 `json:"NSViewLayerContentsRedrawNever"`
			NSViewLayerContentsRedrawOnSetNeedsDisplay            int64 `json:"NSViewLayerContentsRedrawOnSetNeedsDisplay"`
			NSViewMaxXMargin                                      int64 `json:"NSViewMaxXMargin"`
			NSViewMaxYMargin                                      int64 `json:"NSViewMaxYMargin"`
			NSViewMinXMargin                                      int64 `json:"NSViewMinXMargin"`
			NSViewMinYMargin                                      int64 `json:"NSViewMinYMargin"`
			NSViewNotSizable                                      int64 `json:"NSViewNotSizable"`
			NSViewWidthSizable                                    int64 `json:"NSViewWidthSizable"`
			NSVisualEffectBlendingModeBehindWindow                int64 `json:"NSVisualEffectBlendingModeBehindWindow"`
			NSVisualEffectBlendingModeWithinWindow                int64 `json:"NSVisualEffectBlendingModeWithinWindow"`
			NSVisualEffectMaterialAppearanceBased                 int64 `json:"NSVisualEffectMaterialAppearanceBased"`
			NSVisualEffectMaterialDark                            int64 `json:"NSVisualEffectMaterialDark"`
			NSVisualEffectMaterialLight                           int64 `json:"NSVisualEffectMaterialLight"`
			NSVisualEffectMaterialTitlebar                        int64 `json:"NSVisualEffectMaterialTitlebar"`
			NSVisualEffectStateActive                             int64 `json:"NSVisualEffectStateActive"`
			NSVisualEffectStateFollowsWindowActiveState           int64 `json:"NSVisualEffectStateFollowsWindowActiveState"`
			NSVisualEffectStateInactive                           int64 `json:"NSVisualEffectStateInactive"`
			NSWantsBidiLevels                                     int64 `json:"NSWantsBidiLevels"`
			NSWarningAlertStyle                                   int64 `json:"NSWarningAlertStyle"`
			NSWheelModeColorPanel                                 int64 `json:"NSWheelModeColorPanel"`
			NSWindowAbove                                         int64 `json:"NSWindowAbove"`
			NSWindowAnimationBehaviorAlertPanel                   int64 `json:"NSWindowAnimationBehaviorAlertPanel"`
			NSWindowAnimationBehaviorDefault                      int64 `json:"NSWindowAnimationBehaviorDefault"`
			NSWindowAnimationBehaviorDocumentWindow               int64 `json:"NSWindowAnimationBehaviorDocumentWindow"`
			NSWindowAnimationBehaviorNone                         int64 `json:"NSWindowAnimationBehaviorNone"`
			NSWindowAnimationBehaviorUtilityWindow                int64 `json:"NSWindowAnimationBehaviorUtilityWindow"`
			NSWindowBackingLocationDefault                        int64 `json:"NSWindowBackingLocationDefault"`
			NSWindowBackingLocationMainMemory                     int64 `json:"NSWindowBackingLocationMainMemory"`
			NSWindowBackingLocationVideoMemory                    int64 `json:"NSWindowBackingLocationVideoMemory"`
			NSWindowBelow                                         int64 `json:"NSWindowBelow"`
			NSWindowCloseButton                                   int64 `json:"NSWindowCloseButton"`
			NSWindowCollectionBehaviorCanJoinAllSpaces            int64 `json:"NSWindowCollectionBehaviorCanJoinAllSpaces"`
			NSWindowCollectionBehaviorDefault                     int64 `json:"NSWindowCollectionBehaviorDefault"`
			NSWindowCollectionBehaviorFullScreenAuxiliary         int64 `json:"NSWindowCollectionBehaviorFullScreenAuxiliary"`
			NSWindowCollectionBehaviorFullScreenPrimary           int64 `json:"NSWindowCollectionBehaviorFullScreenPrimary"`
			NSWindowCollectionBehaviorIgnoresCycle                int64 `json:"NSWindowCollectionBehaviorIgnoresCycle"`
			NSWindowCollectionBehaviorManaged                     int64 `json:"NSWindowCollectionBehaviorManaged"`
			NSWindowCollectionBehaviorMoveToActiveSpace           int64 `json:"NSWindowCollectionBehaviorMoveToActiveSpace"`
			NSWindowCollectionBehaviorParticipatesInCycle         int64 `json:"NSWindowCollectionBehaviorParticipatesInCycle"`
			NSWindowCollectionBehaviorStationary                  int64 `json:"NSWindowCollectionBehaviorStationary"`
			NSWindowCollectionBehaviorTransient                   int64 `json:"NSWindowCollectionBehaviorTransient"`
			NSWindowDepthOnehundredtwentyeightBitRGB              int64 `json:"NSWindowDepthOnehundredtwentyeightBitRGB"`
			NSWindowDepthSixtyfourBitRGB                          int64 `json:"NSWindowDepthSixtyfourBitRGB"`
			NSWindowDepthTwentyfourBitRGB                         int64 `json:"NSWindowDepthTwentyfourBitRGB"`
			NSWindowDocumentIconButton                            int64 `json:"NSWindowDocumentIconButton"`
			NSWindowDocumentVersionsButton                        int64 `json:"NSWindowDocumentVersionsButton"`
			NSWindowExposedEventType                              int64 `json:"NSWindowExposedEventType"`
			NSWindowFullScreenButton                              int64 `json:"NSWindowFullScreenButton"`
			NSWindowMiniaturizeButton                             int64 `json:"NSWindowMiniaturizeButton"`
			NSWindowMovedEventType                                int64 `json:"NSWindowMovedEventType"`
			NSWindowNumberListAllApplications                     int64 `json:"NSWindowNumberListAllApplications"`
			NSWindowNumberListAllSpaces                           int64 `json:"NSWindowNumberListAllSpaces"`
			NSWindowOcclusionStateVisible                         int64 `json:"NSWindowOcclusionStateVisible"`
			NSWindowOut                                           int64 `json:"NSWindowOut"`
			NSWindowSharingNone                                   int64 `json:"NSWindowSharingNone"`
			NSWindowSharingReadOnly                               int64 `json:"NSWindowSharingReadOnly"`
			NSWindowSharingReadWrite                              int64 `json:"NSWindowSharingReadWrite"`
			NSWindowTitleHidden                                   int64 `json:"NSWindowTitleHidden"`
			NSWindowTitleVisible                                  int64 `json:"NSWindowTitleVisible"`
			NSWindowToolbarButton                                 int64 `json:"NSWindowToolbarButton"`
			NSWindowZoomButton                                    int64 `json:"NSWindowZoomButton"`
			NSWindows95InterfaceStyle                             int64 `json:"NSWindows95InterfaceStyle"`
			NSWorkspaceLaunchAllowingClassicStartup               int64 `json:"NSWorkspaceLaunchAllowingClassicStartup"`
			NSWorkspaceLaunchAndHide                              int64 `json:"NSWorkspaceLaunchAndHide"`
			NSWorkspaceLaunchAndHideOthers                        int64 `json:"NSWorkspaceLaunchAndHideOthers"`
			NSWorkspaceLaunchAndPrint                             int64 `json:"NSWorkspaceLaunchAndPrint"`
			NSWorkspaceLaunchAsync                                int64 `json:"NSWorkspaceLaunchAsync"`
			NSWorkspaceLaunchDefault                              int64 `json:"NSWorkspaceLaunchDefault"`
			NSWorkspaceLaunchInhibitingBackgroundOnly             int64 `json:"NSWorkspaceLaunchInhibitingBackgroundOnly"`
			NSWorkspaceLaunchNewInstance                          int64 `json:"NSWorkspaceLaunchNewInstance"`
			NSWorkspaceLaunchPreferringClassic                    int64 `json:"NSWorkspaceLaunchPreferringClassic"`
			NSWorkspaceLaunchWithErrorPresentation                int64 `json:"NSWorkspaceLaunchWithErrorPresentation"`
			NSWorkspaceLaunchWithoutActivation                    int64 `json:"NSWorkspaceLaunchWithoutActivation"`
			NSWorkspaceLaunchWithoutAddingToRecents               int64 `json:"NSWorkspaceLaunchWithoutAddingToRecents"`
			NSWritingDirectionLeftToRight                         int64 `json:"NSWritingDirectionLeftToRight"`
			NSWritingDirectionNatural                             int64 `json:"NSWritingDirectionNatural"`
			NSWritingDirectionRightToLeft                         int64 `json:"NSWritingDirectionRightToLeft"`
			NSYearMonthDatePickerElementFlag                      int64 `json:"NSYearMonthDatePickerElementFlag"`
			NSYearMonthDayDatePickerElementFlag                   int64 `json:"NSYearMonthDayDatePickerElementFlag"`
			NsUserActivitySupported                               int64 `json:"NS_USER_ACTIVITY_SUPPORTED"`
		} `json:"enum"`
		Expressions struct{} `json:"expressions"`
		Externs     struct {
			NSAbortModalException struct {
				Typestr string `json:"typestr"`
			} `json:"NSAbortModalException"`
			NSAbortPrintingException struct {
				Typestr string `json:"typestr"`
			} `json:"NSAbortPrintingException"`
			NSAccessibilityActivationPointAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityActivationPointAttribute"`
			NSAccessibilityAllowedValuesAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAllowedValuesAttribute"`
			NSAccessibilityAlternateUIVisibleAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAlternateUIVisibleAttribute"`
			NSAccessibilityAnnouncementKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAnnouncementKey"`
			NSAccessibilityAnnouncementRequestedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAnnouncementRequestedNotification"`
			NSAccessibilityApplicationActivatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityApplicationActivatedNotification"`
			NSAccessibilityApplicationDeactivatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityApplicationDeactivatedNotification"`
			NSAccessibilityApplicationHiddenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityApplicationHiddenNotification"`
			NSAccessibilityApplicationRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityApplicationRole"`
			NSAccessibilityApplicationShownNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityApplicationShownNotification"`
			NSAccessibilityAscendingSortDirectionValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAscendingSortDirectionValue"`
			NSAccessibilityAttachmentTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAttachmentTextAttribute"`
			NSAccessibilityAttributedStringForRangeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAttributedStringForRangeParameterizedAttribute"`
			NSAccessibilityAutocorrectedTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityAutocorrectedTextAttribute"`
			NSAccessibilityBackgroundColorTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityBackgroundColorTextAttribute"`
			NSAccessibilityBoundsForRangeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityBoundsForRangeParameterizedAttribute"`
			NSAccessibilityBrowserRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityBrowserRole"`
			NSAccessibilityBusyIndicatorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityBusyIndicatorRole"`
			NSAccessibilityButtonRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityButtonRole"`
			NSAccessibilityCancelAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCancelAction"`
			NSAccessibilityCancelButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCancelButtonAttribute"`
			NSAccessibilityCellForColumnAndRowParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCellForColumnAndRowParameterizedAttribute"`
			NSAccessibilityCellRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCellRole"`
			NSAccessibilityCenterTabStopMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCenterTabStopMarkerTypeValue"`
			NSAccessibilityCentimetersUnitValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCentimetersUnitValue"`
			NSAccessibilityCheckBoxRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCheckBoxRole"`
			NSAccessibilityChildrenAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityChildrenAttribute"`
			NSAccessibilityClearButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityClearButtonAttribute"`
			NSAccessibilityCloseButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCloseButtonAttribute"`
			NSAccessibilityCloseButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCloseButtonSubrole"`
			NSAccessibilityColorWellRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColorWellRole"`
			NSAccessibilityColumnCountAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnCountAttribute"`
			NSAccessibilityColumnHeaderUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnHeaderUIElementsAttribute"`
			NSAccessibilityColumnIndexRangeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnIndexRangeAttribute"`
			NSAccessibilityColumnRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnRole"`
			NSAccessibilityColumnTitlesAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnTitlesAttribute"`
			NSAccessibilityColumnsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityColumnsAttribute"`
			NSAccessibilityComboBoxRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityComboBoxRole"`
			NSAccessibilityConfirmAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityConfirmAction"`
			NSAccessibilityContainsProtectedContentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityContainsProtectedContentAttribute"`
			NSAccessibilityContentListSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityContentListSubrole"`
			NSAccessibilityContentsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityContentsAttribute"`
			NSAccessibilityCreatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCreatedNotification"`
			NSAccessibilityCriticalValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityCriticalValueAttribute"`
			NSAccessibilityDecimalTabStopMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDecimalTabStopMarkerTypeValue"`
			NSAccessibilityDecrementAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDecrementAction"`
			NSAccessibilityDecrementArrowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDecrementArrowSubrole"`
			NSAccessibilityDecrementButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDecrementButtonAttribute"`
			NSAccessibilityDecrementPageSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDecrementPageSubrole"`
			NSAccessibilityDefaultButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDefaultButtonAttribute"`
			NSAccessibilityDefinitionListSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDefinitionListSubrole"`
			NSAccessibilityDeleteAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDeleteAction"`
			NSAccessibilityDescendingSortDirectionValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDescendingSortDirectionValue"`
			NSAccessibilityDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDescriptionAttribute"`
			NSAccessibilityDescriptionListSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDescriptionListSubrole"`
			NSAccessibilityDialogSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDialogSubrole"`
			NSAccessibilityDisclosedByRowAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDisclosedByRowAttribute"`
			NSAccessibilityDisclosedRowsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDisclosedRowsAttribute"`
			NSAccessibilityDisclosingAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDisclosingAttribute"`
			NSAccessibilityDisclosureLevelAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDisclosureLevelAttribute"`
			NSAccessibilityDisclosureTriangleRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDisclosureTriangleRole"`
			NSAccessibilityDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDocumentAttribute"`
			NSAccessibilityDrawerCreatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDrawerCreatedNotification"`
			NSAccessibilityDrawerRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityDrawerRole"`
			NSAccessibilityEditedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityEditedAttribute"`
			NSAccessibilityEnabledAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityEnabledAttribute"`
			NSAccessibilityErrorCodeExceptionInfo struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityErrorCodeExceptionInfo"`
			NSAccessibilityException struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityException"`
			NSAccessibilityExpandedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityExpandedAttribute"`
			NSAccessibilityExtrasMenuBarAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityExtrasMenuBarAttribute"`
			NSAccessibilityFilenameAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFilenameAttribute"`
			NSAccessibilityFirstLineIndentMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFirstLineIndentMarkerTypeValue"`
			NSAccessibilityFloatingWindowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFloatingWindowSubrole"`
			NSAccessibilityFocusedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFocusedAttribute"`
			NSAccessibilityFocusedUIElementAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFocusedUIElementAttribute"`
			NSAccessibilityFocusedUIElementChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFocusedUIElementChangedNotification"`
			NSAccessibilityFocusedWindowAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFocusedWindowAttribute"`
			NSAccessibilityFocusedWindowChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFocusedWindowChangedNotification"`
			NSAccessibilityFontFamilyKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFontFamilyKey"`
			NSAccessibilityFontNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFontNameKey"`
			NSAccessibilityFontSizeKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFontSizeKey"`
			NSAccessibilityFontTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFontTextAttribute"`
			NSAccessibilityForegroundColorTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityForegroundColorTextAttribute"`
			NSAccessibilityFrontmostAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFrontmostAttribute"`
			NSAccessibilityFullScreenButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFullScreenButtonAttribute"`
			NSAccessibilityFullScreenButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityFullScreenButtonSubrole"`
			NSAccessibilityGridRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityGridRole"`
			NSAccessibilityGroupRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityGroupRole"`
			NSAccessibilityGrowAreaAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityGrowAreaAttribute"`
			NSAccessibilityGrowAreaRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityGrowAreaRole"`
			NSAccessibilityHandleRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHandleRole"`
			NSAccessibilityHandlesAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHandlesAttribute"`
			NSAccessibilityHeadIndentMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHeadIndentMarkerTypeValue"`
			NSAccessibilityHeaderAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHeaderAttribute"`
			NSAccessibilityHelpAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHelpAttribute"`
			NSAccessibilityHelpTagCreatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHelpTagCreatedNotification"`
			NSAccessibilityHelpTagRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHelpTagRole"`
			NSAccessibilityHiddenAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHiddenAttribute"`
			NSAccessibilityHorizontalOrientationValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHorizontalOrientationValue"`
			NSAccessibilityHorizontalScrollBarAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHorizontalScrollBarAttribute"`
			NSAccessibilityHorizontalUnitDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHorizontalUnitDescriptionAttribute"`
			NSAccessibilityHorizontalUnitsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityHorizontalUnitsAttribute"`
			NSAccessibilityIdentifierAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIdentifierAttribute"`
			NSAccessibilityImageRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityImageRole"`
			NSAccessibilityInchesUnitValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityInchesUnitValue"`
			NSAccessibilityIncrementAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIncrementAction"`
			NSAccessibilityIncrementArrowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIncrementArrowSubrole"`
			NSAccessibilityIncrementButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIncrementButtonAttribute"`
			NSAccessibilityIncrementPageSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIncrementPageSubrole"`
			NSAccessibilityIncrementorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIncrementorRole"`
			NSAccessibilityIndexAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityIndexAttribute"`
			NSAccessibilityInsertionPointLineNumberAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityInsertionPointLineNumberAttribute"`
			NSAccessibilityLabelUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLabelUIElementsAttribute"`
			NSAccessibilityLabelValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLabelValueAttribute"`
			NSAccessibilityLayoutAreaRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLayoutAreaRole"`
			NSAccessibilityLayoutChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLayoutChangedNotification"`
			NSAccessibilityLayoutItemRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLayoutItemRole"`
			NSAccessibilityLayoutPointForScreenPointParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLayoutPointForScreenPointParameterizedAttribute"`
			NSAccessibilityLayoutSizeForScreenSizeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLayoutSizeForScreenSizeParameterizedAttribute"`
			NSAccessibilityLeftTabStopMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLeftTabStopMarkerTypeValue"`
			NSAccessibilityLevelIndicatorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLevelIndicatorRole"`
			NSAccessibilityLineForIndexParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLineForIndexParameterizedAttribute"`
			NSAccessibilityLinkRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLinkRole"`
			NSAccessibilityLinkTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLinkTextAttribute"`
			NSAccessibilityLinkedUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityLinkedUIElementsAttribute"`
			NSAccessibilityListRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityListRole"`
			NSAccessibilityMainAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMainAttribute"`
			NSAccessibilityMainWindowAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMainWindowAttribute"`
			NSAccessibilityMainWindowChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMainWindowChangedNotification"`
			NSAccessibilityMarkedMisspelledTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkedMisspelledTextAttribute"`
			NSAccessibilityMarkerGroupUIElementAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkerGroupUIElementAttribute"`
			NSAccessibilityMarkerTypeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkerTypeAttribute"`
			NSAccessibilityMarkerTypeDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkerTypeDescriptionAttribute"`
			NSAccessibilityMarkerUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkerUIElementsAttribute"`
			NSAccessibilityMarkerValuesAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMarkerValuesAttribute"`
			NSAccessibilityMatteContentUIElementAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMatteContentUIElementAttribute"`
			NSAccessibilityMatteHoleAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMatteHoleAttribute"`
			NSAccessibilityMatteRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMatteRole"`
			NSAccessibilityMaxValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMaxValueAttribute"`
			NSAccessibilityMenuBarAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMenuBarAttribute"`
			NSAccessibilityMenuBarRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMenuBarRole"`
			NSAccessibilityMenuButtonRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMenuButtonRole"`
			NSAccessibilityMenuItemRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMenuItemRole"`
			NSAccessibilityMenuRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMenuRole"`
			NSAccessibilityMinValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMinValueAttribute"`
			NSAccessibilityMinimizeButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMinimizeButtonAttribute"`
			NSAccessibilityMinimizeButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMinimizeButtonSubrole"`
			NSAccessibilityMinimizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMinimizedAttribute"`
			NSAccessibilityMisspelledTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMisspelledTextAttribute"`
			NSAccessibilityModalAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityModalAttribute"`
			NSAccessibilityMovedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityMovedNotification"`
			NSAccessibilityNextContentsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityNextContentsAttribute"`
			NSAccessibilityNumberOfCharactersAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityNumberOfCharactersAttribute"`
			NSAccessibilityOrderedByRowAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityOrderedByRowAttribute"`
			NSAccessibilityOrientationAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityOrientationAttribute"`
			NSAccessibilityOutlineRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityOutlineRole"`
			NSAccessibilityOutlineRowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityOutlineRowSubrole"`
			NSAccessibilityOverflowButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityOverflowButtonAttribute"`
			NSAccessibilityParentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityParentAttribute"`
			NSAccessibilityPicasUnitValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPicasUnitValue"`
			NSAccessibilityPickAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPickAction"`
			NSAccessibilityPlaceholderValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPlaceholderValueAttribute"`
			NSAccessibilityPointsUnitValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPointsUnitValue"`
			NSAccessibilityPopUpButtonRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPopUpButtonRole"`
			NSAccessibilityPopoverRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPopoverRole"`
			NSAccessibilityPositionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPositionAttribute"`
			NSAccessibilityPressAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPressAction"`
			NSAccessibilityPreviousContentsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPreviousContentsAttribute"`
			NSAccessibilityPriorityKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityPriorityKey"`
			NSAccessibilityProgressIndicatorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityProgressIndicatorRole"`
			NSAccessibilityProxyAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityProxyAttribute"`
			NSAccessibilityRTFForRangeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRTFForRangeParameterizedAttribute"`
			NSAccessibilityRadioButtonRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRadioButtonRole"`
			NSAccessibilityRadioGroupRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRadioGroupRole"`
			NSAccessibilityRaiseAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRaiseAction"`
			NSAccessibilityRangeForIndexParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRangeForIndexParameterizedAttribute"`
			NSAccessibilityRangeForLineParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRangeForLineParameterizedAttribute"`
			NSAccessibilityRangeForPositionParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRangeForPositionParameterizedAttribute"`
			NSAccessibilityRatingIndicatorSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRatingIndicatorSubrole"`
			NSAccessibilityRelevanceIndicatorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRelevanceIndicatorRole"`
			NSAccessibilityResizedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityResizedNotification"`
			NSAccessibilityRightTabStopMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRightTabStopMarkerTypeValue"`
			NSAccessibilityRoleAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRoleAttribute"`
			NSAccessibilityRoleDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRoleDescriptionAttribute"`
			NSAccessibilityRowCollapsedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowCollapsedNotification"`
			NSAccessibilityRowCountAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowCountAttribute"`
			NSAccessibilityRowCountChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowCountChangedNotification"`
			NSAccessibilityRowExpandedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowExpandedNotification"`
			NSAccessibilityRowHeaderUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowHeaderUIElementsAttribute"`
			NSAccessibilityRowIndexRangeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowIndexRangeAttribute"`
			NSAccessibilityRowRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowRole"`
			NSAccessibilityRowsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRowsAttribute"`
			NSAccessibilityRulerMarkerRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRulerMarkerRole"`
			NSAccessibilityRulerRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityRulerRole"`
			NSAccessibilityScreenPointForLayoutPointParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityScreenPointForLayoutPointParameterizedAttribute"`
			NSAccessibilityScreenSizeForLayoutSizeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityScreenSizeForLayoutSizeParameterizedAttribute"`
			NSAccessibilityScrollAreaRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityScrollAreaRole"`
			NSAccessibilityScrollBarRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityScrollBarRole"`
			NSAccessibilitySearchButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySearchButtonAttribute"`
			NSAccessibilitySearchFieldSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySearchFieldSubrole"`
			NSAccessibilitySearchMenuAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySearchMenuAttribute"`
			NSAccessibilitySecureTextFieldSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySecureTextFieldSubrole"`
			NSAccessibilitySelectedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedAttribute"`
			NSAccessibilitySelectedCellsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedCellsAttribute"`
			NSAccessibilitySelectedCellsChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedCellsChangedNotification"`
			NSAccessibilitySelectedChildrenAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedChildrenAttribute"`
			NSAccessibilitySelectedChildrenChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedChildrenChangedNotification"`
			NSAccessibilitySelectedChildrenMovedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedChildrenMovedNotification"`
			NSAccessibilitySelectedColumnsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedColumnsAttribute"`
			NSAccessibilitySelectedColumnsChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedColumnsChangedNotification"`
			NSAccessibilitySelectedRowsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedRowsAttribute"`
			NSAccessibilitySelectedRowsChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedRowsChangedNotification"`
			NSAccessibilitySelectedTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedTextAttribute"`
			NSAccessibilitySelectedTextChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedTextChangedNotification"`
			NSAccessibilitySelectedTextRangeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedTextRangeAttribute"`
			NSAccessibilitySelectedTextRangesAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySelectedTextRangesAttribute"`
			NSAccessibilityServesAsTitleForUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityServesAsTitleForUIElementsAttribute"`
			NSAccessibilityShadowTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityShadowTextAttribute"`
			NSAccessibilitySharedCharacterRangeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySharedCharacterRangeAttribute"`
			NSAccessibilitySharedFocusElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySharedFocusElementsAttribute"`
			NSAccessibilitySharedTextUIElementsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySharedTextUIElementsAttribute"`
			NSAccessibilitySheetCreatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySheetCreatedNotification"`
			NSAccessibilitySheetRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySheetRole"`
			NSAccessibilityShowAlternateUIAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityShowAlternateUIAction"`
			NSAccessibilityShowDefaultUIAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityShowDefaultUIAction"`
			NSAccessibilityShowMenuAction struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityShowMenuAction"`
			NSAccessibilityShownMenuAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityShownMenuAttribute"`
			NSAccessibilitySizeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySizeAttribute"`
			NSAccessibilitySliderRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySliderRole"`
			NSAccessibilitySortButtonRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySortButtonRole"`
			NSAccessibilitySortButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySortButtonSubrole"`
			NSAccessibilitySortDirectionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySortDirectionAttribute"`
			NSAccessibilitySplitGroupRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySplitGroupRole"`
			NSAccessibilitySplitterRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySplitterRole"`
			NSAccessibilitySplittersAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySplittersAttribute"`
			NSAccessibilityStandardWindowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStandardWindowSubrole"`
			NSAccessibilityStaticTextRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStaticTextRole"`
			NSAccessibilityStrikethroughColorTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStrikethroughColorTextAttribute"`
			NSAccessibilityStrikethroughTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStrikethroughTextAttribute"`
			NSAccessibilityStringForRangeParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStringForRangeParameterizedAttribute"`
			NSAccessibilityStyleRangeForIndexParameterizedAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityStyleRangeForIndexParameterizedAttribute"`
			NSAccessibilitySubroleAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySubroleAttribute"`
			NSAccessibilitySuperscriptTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySuperscriptTextAttribute"`
			NSAccessibilitySwitchSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySwitchSubrole"`
			NSAccessibilitySystemDialogSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySystemDialogSubrole"`
			NSAccessibilitySystemFloatingWindowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySystemFloatingWindowSubrole"`
			NSAccessibilitySystemWideRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilitySystemWideRole"`
			NSAccessibilityTabGroupRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTabGroupRole"`
			NSAccessibilityTableRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTableRole"`
			NSAccessibilityTableRowSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTableRowSubrole"`
			NSAccessibilityTabsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTabsAttribute"`
			NSAccessibilityTailIndentMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTailIndentMarkerTypeValue"`
			NSAccessibilityTextAreaRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTextAreaRole"`
			NSAccessibilityTextAttachmentSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTextAttachmentSubrole"`
			NSAccessibilityTextFieldRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTextFieldRole"`
			NSAccessibilityTextLinkSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTextLinkSubrole"`
			NSAccessibilityTimelineSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTimelineSubrole"`
			NSAccessibilityTitleAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTitleAttribute"`
			NSAccessibilityTitleChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTitleChangedNotification"`
			NSAccessibilityTitleUIElementAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTitleUIElementAttribute"`
			NSAccessibilityToggleSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityToggleSubrole"`
			NSAccessibilityToolbarButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityToolbarButtonAttribute"`
			NSAccessibilityToolbarButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityToolbarButtonSubrole"`
			NSAccessibilityToolbarRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityToolbarRole"`
			NSAccessibilityTopLevelUIElementAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityTopLevelUIElementAttribute"`
			NSAccessibilityUIElementDestroyedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUIElementDestroyedNotification"`
			NSAccessibilityUIElementsKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUIElementsKey"`
			NSAccessibilityURLAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityURLAttribute"`
			NSAccessibilityUnderlineColorTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnderlineColorTextAttribute"`
			NSAccessibilityUnderlineTextAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnderlineTextAttribute"`
			NSAccessibilityUnitDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnitDescriptionAttribute"`
			NSAccessibilityUnitsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnitsAttribute"`
			NSAccessibilityUnitsChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnitsChangedNotification"`
			NSAccessibilityUnknownMarkerTypeValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownMarkerTypeValue"`
			NSAccessibilityUnknownOrientationValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownOrientationValue"`
			NSAccessibilityUnknownRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownRole"`
			NSAccessibilityUnknownSortDirectionValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownSortDirectionValue"`
			NSAccessibilityUnknownSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownSubrole"`
			NSAccessibilityUnknownUnitValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityUnknownUnitValue"`
			NSAccessibilityValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityValueAttribute"`
			NSAccessibilityValueChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityValueChangedNotification"`
			NSAccessibilityValueDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityValueDescriptionAttribute"`
			NSAccessibilityValueIndicatorRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityValueIndicatorRole"`
			NSAccessibilityVerticalOrientationValue struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVerticalOrientationValue"`
			NSAccessibilityVerticalScrollBarAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVerticalScrollBarAttribute"`
			NSAccessibilityVerticalUnitDescriptionAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVerticalUnitDescriptionAttribute"`
			NSAccessibilityVerticalUnitsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVerticalUnitsAttribute"`
			NSAccessibilityVisibleCellsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleCellsAttribute"`
			NSAccessibilityVisibleCharacterRangeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleCharacterRangeAttribute"`
			NSAccessibilityVisibleChildrenAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleChildrenAttribute"`
			NSAccessibilityVisibleColumnsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleColumnsAttribute"`
			NSAccessibilityVisibleNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleNameKey"`
			NSAccessibilityVisibleRowsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityVisibleRowsAttribute"`
			NSAccessibilityWarningValueAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWarningValueAttribute"`
			NSAccessibilityWindowAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowAttribute"`
			NSAccessibilityWindowCreatedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowCreatedNotification"`
			NSAccessibilityWindowDeminiaturizedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowDeminiaturizedNotification"`
			NSAccessibilityWindowMiniaturizedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowMiniaturizedNotification"`
			NSAccessibilityWindowMovedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowMovedNotification"`
			NSAccessibilityWindowResizedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowResizedNotification"`
			NSAccessibilityWindowRole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowRole"`
			NSAccessibilityWindowsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityWindowsAttribute"`
			NSAccessibilityZoomButtonAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityZoomButtonAttribute"`
			NSAccessibilityZoomButtonSubrole struct {
				Typestr string `json:"typestr"`
			} `json:"NSAccessibilityZoomButtonSubrole"`
			NSAlignmentBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAlignmentBinding"`
			NSAllRomanInputSourcesLocaleIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSAllRomanInputSourcesLocaleIdentifier"`
			NSAllowsEditingMultipleValuesSelectionBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSAllowsEditingMultipleValuesSelectionBindingOption"`
			NSAllowsNullArgumentBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSAllowsNullArgumentBindingOption"`
			NSAlternateImageBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAlternateImageBinding"`
			NSAlternateTitleBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAlternateTitleBinding"`
			NSAlwaysPresentsApplicationModalAlertsBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSAlwaysPresentsApplicationModalAlertsBindingOption"`
			NSAnimateBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimateBinding"`
			NSAnimationDelayBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimationDelayBinding"`
			NSAnimationProgressMark struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimationProgressMark"`
			NSAnimationProgressMarkNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimationProgressMarkNotification"`
			NSAnimationTriggerOrderIn struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimationTriggerOrderIn"`
			NSAnimationTriggerOrderOut struct {
				Typestr string `json:"typestr"`
			} `json:"NSAnimationTriggerOrderOut"`
			NSAntialiasThresholdChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSAntialiasThresholdChangedNotification"`
			NSApp struct {
				Typestr string `json:"typestr"`
			} `json:"NSApp"`
			NSAppKitIgnoredException struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppKitIgnoredException"`
			NSAppKitVersionNumber struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppKitVersionNumber"`
			NSAppKitVirtualMemoryException struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppKitVirtualMemoryException"`
			NSAppearanceNameAqua struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppearanceNameAqua"`
			NSAppearanceNameLightContent struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppearanceNameLightContent"`
			NSAppearanceNameVibrantDark struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppearanceNameVibrantDark"`
			NSAppearanceNameVibrantLight struct {
				Typestr string `json:"typestr"`
			} `json:"NSAppearanceNameVibrantLight"`
			NSApplicationDidBecomeActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidBecomeActiveNotification"`
			NSApplicationDidChangeOcclusionStateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidChangeOcclusionStateNotification"`
			NSApplicationDidChangeScreenParametersNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidChangeScreenParametersNotification"`
			NSApplicationDidFinishLaunchingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidFinishLaunchingNotification"`
			NSApplicationDidFinishRestoringWindowsNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidFinishRestoringWindowsNotification"`
			NSApplicationDidHideNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidHideNotification"`
			NSApplicationDidResignActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidResignActiveNotification"`
			NSApplicationDidUnhideNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidUnhideNotification"`
			NSApplicationDidUpdateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationDidUpdateNotification"`
			NSApplicationFileType struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationFileType"`
			NSApplicationLaunchIsDefaultLaunchKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationLaunchIsDefaultLaunchKey"`
			NSApplicationLaunchRemoteNotificationKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationLaunchRemoteNotificationKey"`
			NSApplicationLaunchUserNotificationKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationLaunchUserNotificationKey"`
			NSApplicationWillBecomeActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillBecomeActiveNotification"`
			NSApplicationWillFinishLaunchingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillFinishLaunchingNotification"`
			NSApplicationWillHideNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillHideNotification"`
			NSApplicationWillResignActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillResignActiveNotification"`
			NSApplicationWillTerminateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillTerminateNotification"`
			NSApplicationWillUnhideNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillUnhideNotification"`
			NSApplicationWillUpdateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSApplicationWillUpdateNotification"`
			NSArgumentBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSArgumentBinding"`
			NSAttachmentAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSAttachmentAttributeName"`
			NSAttributedStringBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSAttributedStringBinding"`
			NSAuthorDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSAuthorDocumentAttribute"`
			NSBackgroundColorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSBackgroundColorAttributeName"`
			NSBackgroundColorDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSBackgroundColorDocumentAttribute"`
			NSBackingPropertyOldColorSpaceKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSBackingPropertyOldColorSpaceKey"`
			NSBackingPropertyOldScaleFactorKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSBackingPropertyOldScaleFactorKey"`
			NSBadBitmapParametersException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadBitmapParametersException"`
			NSBadComparisonException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadComparisonException"`
			NSBadRTFColorTableException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadRTFColorTableException"`
			NSBadRTFDirectiveException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadRTFDirectiveException"`
			NSBadRTFFontTableException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadRTFFontTableException"`
			NSBadRTFStyleSheetException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBadRTFStyleSheetException"`
			NSBaseURLDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSBaseURLDocumentOption"`
			NSBaselineOffsetAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSBaselineOffsetAttributeName"`
			NSBlack struct {
				Typestr string `json:"typestr"`
			} `json:"NSBlack"`
			NSBottomMarginDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSBottomMarginDocumentAttribute"`
			NSBrowserColumnConfigurationDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSBrowserColumnConfigurationDidChangeNotification"`
			NSBrowserIllegalDelegateException struct {
				Typestr string `json:"typestr"`
			} `json:"NSBrowserIllegalDelegateException"`
			NSCalibratedBlackColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSCalibratedBlackColorSpace"`
			NSCalibratedRGBColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSCalibratedRGBColorSpace"`
			NSCalibratedWhiteColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSCalibratedWhiteColorSpace"`
			NSCategoryDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCategoryDocumentAttribute"`
			NSCharacterEncodingDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCharacterEncodingDocumentAttribute"`
			NSCharacterEncodingDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSCharacterEncodingDocumentOption"`
			NSCharacterShapeAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSCharacterShapeAttributeName"`
			NSCocoaVersionDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCocoaVersionDocumentAttribute"`
			NSColorListDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSColorListDidChangeNotification"`
			NSColorListIOException struct {
				Typestr string `json:"typestr"`
			} `json:"NSColorListIOException"`
			NSColorListNotEditableException struct {
				Typestr string `json:"typestr"`
			} `json:"NSColorListNotEditableException"`
			NSColorPanelColorDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSColorPanelColorDidChangeNotification"`
			NSColorPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSColorPboardType"`
			NSComboBoxSelectionDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSComboBoxSelectionDidChangeNotification"`
			NSComboBoxSelectionIsChangingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSComboBoxSelectionIsChangingNotification"`
			NSComboBoxWillDismissNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSComboBoxWillDismissNotification"`
			NSComboBoxWillPopUpNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSComboBoxWillPopUpNotification"`
			NSCommentDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCommentDocumentAttribute"`
			NSCompanyDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCompanyDocumentAttribute"`
			NSConditionallySetsEditableBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSConditionallySetsEditableBindingOption"`
			NSConditionallySetsEnabledBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSConditionallySetsEnabledBindingOption"`
			NSConditionallySetsHiddenBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSConditionallySetsHiddenBindingOption"`
			NSContentArrayBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentArrayBinding"`
			NSContentArrayForMultipleSelectionBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentArrayForMultipleSelectionBinding"`
			NSContentBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentBinding"`
			NSContentDictionaryBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentDictionaryBinding"`
			NSContentHeightBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentHeightBinding"`
			NSContentObjectBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentObjectBinding"`
			NSContentObjectsBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentObjectsBinding"`
			NSContentPlacementTagBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentPlacementTagBindingOption"`
			NSContentSetBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentSetBinding"`
			NSContentValuesBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentValuesBinding"`
			NSContentWidthBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSContentWidthBinding"`
			NSContextHelpModeDidActivateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSContextHelpModeDidActivateNotification"`
			NSContextHelpModeDidDeactivateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSContextHelpModeDidDeactivateNotification"`
			NSContinuouslyUpdatesValueBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSContinuouslyUpdatesValueBindingOption"`
			NSControlTextDidBeginEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSControlTextDidBeginEditingNotification"`
			NSControlTextDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSControlTextDidChangeNotification"`
			NSControlTextDidEndEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSControlTextDidEndEditingNotification"`
			NSControlTintDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSControlTintDidChangeNotification"`
			NSConvertedDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSConvertedDocumentAttribute"`
			NSCopyrightDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCopyrightDocumentAttribute"`
			NSCreatesSortDescriptorBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSCreatesSortDescriptorBindingOption"`
			NSCreationTimeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSCreationTimeDocumentAttribute"`
			NSCriticalValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSCriticalValueBinding"`
			NSCursorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSCursorAttributeName"`
			NSCustomColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSCustomColorSpace"`
			NSDarkGray struct {
				Typestr string `json:"typestr"`
			} `json:"NSDarkGray"`
			NSDataBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDataBinding"`
			NSDefaultAttributesDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSDefaultAttributesDocumentOption"`
			NSDefaultTabIntervalDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSDefaultTabIntervalDocumentAttribute"`
			NSDefinitionPresentationTypeDictionaryApplication struct {
				Typestr string `json:"typestr"`
			} `json:"NSDefinitionPresentationTypeDictionaryApplication"`
			NSDefinitionPresentationTypeKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSDefinitionPresentationTypeKey"`
			NSDefinitionPresentationTypeOverlay struct {
				Typestr string `json:"typestr"`
			} `json:"NSDefinitionPresentationTypeOverlay"`
			NSDeletesObjectsOnRemoveBindingsOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeletesObjectsOnRemoveBindingsOption"`
			NSDeviceBitsPerSample struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceBitsPerSample"`
			NSDeviceBlackColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceBlackColorSpace"`
			NSDeviceCMYKColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceCMYKColorSpace"`
			NSDeviceColorSpaceName struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceColorSpaceName"`
			NSDeviceIsPrinter struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceIsPrinter"`
			NSDeviceIsScreen struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceIsScreen"`
			NSDeviceRGBColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceRGBColorSpace"`
			NSDeviceResolution struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceResolution"`
			NSDeviceSize struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceSize"`
			NSDeviceWhiteColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSDeviceWhiteColorSpace"`
			NSDirectoryFileType struct {
				Typestr string `json:"typestr"`
			} `json:"NSDirectoryFileType"`
			NSDisplayNameBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSDisplayNameBindingOption"`
			NSDisplayPatternBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSDisplayPatternBindingOption"`
			NSDisplayPatternTitleBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDisplayPatternTitleBinding"`
			NSDisplayPatternValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDisplayPatternValueBinding"`
			NSDocFormatTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSDocFormatTextDocumentType"`
			NSDocumentEditedBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDocumentEditedBinding"`
			NSDocumentTypeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSDocumentTypeDocumentAttribute"`
			NSDocumentTypeDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSDocumentTypeDocumentOption"`
			NSDoubleClickArgumentBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDoubleClickArgumentBinding"`
			NSDoubleClickTargetBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSDoubleClickTargetBinding"`
			NSDragPboard struct {
				Typestr string `json:"typestr"`
			} `json:"NSDragPboard"`
			NSDraggingException struct {
				Typestr string `json:"typestr"`
			} `json:"NSDraggingException"`
			NSDraggingImageComponentIconKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSDraggingImageComponentIconKey"`
			NSDraggingImageComponentLabelKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSDraggingImageComponentLabelKey"`
			NSDrawerDidCloseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSDrawerDidCloseNotification"`
			NSDrawerDidOpenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSDrawerDidOpenNotification"`
			NSDrawerWillCloseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSDrawerWillCloseNotification"`
			NSDrawerWillOpenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSDrawerWillOpenNotification"`
			NSEditableBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSEditableBinding"`
			NSEditorDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSEditorDocumentAttribute"`
			NSEnabledBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSEnabledBinding"`
			NSEventTrackingRunLoopMode struct {
				Typestr string `json:"typestr"`
			} `json:"NSEventTrackingRunLoopMode"`
			NSExcludedElementsDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSExcludedElementsDocumentAttribute"`
			NSExcludedKeysBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSExcludedKeysBinding"`
			NSExpansionAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSExpansionAttributeName"`
			NSFileContentsPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFileContentsPboardType"`
			NSFileTypeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFileTypeDocumentAttribute"`
			NSFileTypeDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSFileTypeDocumentOption"`
			NSFilenamesPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFilenamesPboardType"`
			NSFilesPromisePboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFilesPromisePboardType"`
			NSFilesystemFileType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFilesystemFileType"`
			NSFilterPredicateBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFilterPredicateBinding"`
			NSFindPanelCaseInsensitiveSearch struct {
				Typestr string `json:"typestr"`
			} `json:"NSFindPanelCaseInsensitiveSearch"`
			NSFindPanelSearchOptionsPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFindPanelSearchOptionsPboardType"`
			NSFindPanelSubstringMatch struct {
				Typestr string `json:"typestr"`
			} `json:"NSFindPanelSubstringMatch"`
			NSFindPboard struct {
				Typestr string `json:"typestr"`
			} `json:"NSFindPboard"`
			NSFontAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontAttributeName"`
			NSFontBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontBinding"`
			NSFontBoldBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontBoldBinding"`
			NSFontCascadeListAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCascadeListAttribute"`
			NSFontCharacterSetAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCharacterSetAttribute"`
			NSFontCollectionActionKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionActionKey"`
			NSFontCollectionAllFonts struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionAllFonts"`
			NSFontCollectionDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionDidChangeNotification"`
			NSFontCollectionDisallowAutoActivationOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionDisallowAutoActivationOption"`
			NSFontCollectionFavorites struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionFavorites"`
			NSFontCollectionIncludeDisabledFontsOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionIncludeDisabledFontsOption"`
			NSFontCollectionNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionNameKey"`
			NSFontCollectionOldNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionOldNameKey"`
			NSFontCollectionRecentlyUsed struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionRecentlyUsed"`
			NSFontCollectionRemoveDuplicatesOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionRemoveDuplicatesOption"`
			NSFontCollectionUser struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionUser"`
			NSFontCollectionVisibilityKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionVisibilityKey"`
			NSFontCollectionWasHidden struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionWasHidden"`
			NSFontCollectionWasRenamed struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionWasRenamed"`
			NSFontCollectionWasShown struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontCollectionWasShown"`
			NSFontColorAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontColorAttribute"`
			NSFontFaceAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFaceAttribute"`
			NSFontFamilyAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFamilyAttribute"`
			NSFontFamilyNameBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFamilyNameBinding"`
			NSFontFeatureSelectorIdentifierKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFeatureSelectorIdentifierKey"`
			NSFontFeatureSettingsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFeatureSettingsAttribute"`
			NSFontFeatureTypeIdentifierKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFeatureTypeIdentifierKey"`
			NSFontFixedAdvanceAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontFixedAdvanceAttribute"`
			NSFontIdentityMatrix struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontIdentityMatrix"`
			NSFontItalicBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontItalicBinding"`
			NSFontMatrixAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontMatrixAttribute"`
			NSFontNameAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontNameAttribute"`
			NSFontNameBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontNameBinding"`
			NSFontPboard struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontPboard"`
			NSFontPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontPboardType"`
			NSFontSetChangedNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontSetChangedNotification"`
			NSFontSizeAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontSizeAttribute"`
			NSFontSizeBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontSizeBinding"`
			NSFontSlantTrait struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontSlantTrait"`
			NSFontSymbolicTrait struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontSymbolicTrait"`
			NSFontTraitsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontTraitsAttribute"`
			NSFontUnavailableException struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontUnavailableException"`
			NSFontVariationAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAttribute"`
			NSFontVariationAxisDefaultValueKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAxisDefaultValueKey"`
			NSFontVariationAxisIdentifierKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAxisIdentifierKey"`
			NSFontVariationAxisMaximumValueKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAxisMaximumValueKey"`
			NSFontVariationAxisMinimumValueKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAxisMinimumValueKey"`
			NSFontVariationAxisNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVariationAxisNameKey"`
			NSFontVisibleNameAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontVisibleNameAttribute"`
			NSFontWeightTrait struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontWeightTrait"`
			NSFontWidthTrait struct {
				Typestr string `json:"typestr"`
			} `json:"NSFontWidthTrait"`
			NSForegroundColorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSForegroundColorAttributeName"`
			NSFullScreenModeAllScreens struct {
				Typestr string `json:"typestr"`
			} `json:"NSFullScreenModeAllScreens"`
			NSFullScreenModeApplicationPresentationOptions struct {
				Typestr string `json:"typestr"`
			} `json:"NSFullScreenModeApplicationPresentationOptions"`
			NSFullScreenModeSetting struct {
				Typestr string `json:"typestr"`
			} `json:"NSFullScreenModeSetting"`
			NSFullScreenModeWindowLevel struct {
				Typestr string `json:"typestr"`
			} `json:"NSFullScreenModeWindowLevel"`
			NSGeneralPboard struct {
				Typestr string `json:"typestr"`
			} `json:"NSGeneralPboard"`
			NSGlyphInfoAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSGlyphInfoAttributeName"`
			NSGraphicsContextDestinationAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSGraphicsContextDestinationAttributeName"`
			NSGraphicsContextPDFFormat struct {
				Typestr string `json:"typestr"`
			} `json:"NSGraphicsContextPDFFormat"`
			NSGraphicsContextPSFormat struct {
				Typestr string `json:"typestr"`
			} `json:"NSGraphicsContextPSFormat"`
			NSGraphicsContextRepresentationFormatAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSGraphicsContextRepresentationFormatAttributeName"`
			NSHTMLPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSHTMLPboardType"`
			NSHTMLTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSHTMLTextDocumentType"`
			NSHandlesContentAsCompoundValueBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSHandlesContentAsCompoundValueBindingOption"`
			NSHeaderTitleBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSHeaderTitleBinding"`
			NSHiddenBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSHiddenBinding"`
			NSHyphenationFactorDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSHyphenationFactorDocumentAttribute"`
			NSIllegalSelectorException struct {
				Typestr string `json:"typestr"`
			} `json:"NSIllegalSelectorException"`
			NSImageBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageBinding"`
			NSImageCacheException struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageCacheException"`
			NSImageColorSyncProfileData struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageColorSyncProfileData"`
			NSImageCompressionFactor struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageCompressionFactor"`
			NSImageCompressionMethod struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageCompressionMethod"`
			NSImageCurrentFrame struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageCurrentFrame"`
			NSImageCurrentFrameDuration struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageCurrentFrameDuration"`
			NSImageDitherTransparency struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageDitherTransparency"`
			NSImageEXIFData struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageEXIFData"`
			NSImageFallbackBackgroundColor struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageFallbackBackgroundColor"`
			NSImageFrameCount struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageFrameCount"`
			NSImageGamma struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageGamma"`
			NSImageHintCTM struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageHintCTM"`
			NSImageHintInterpolation struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageHintInterpolation"`
			NSImageInterlaced struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageInterlaced"`
			NSImageLoopCount struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageLoopCount"`
			NSImageNameActionTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameActionTemplate"`
			NSImageNameAddTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameAddTemplate"`
			NSImageNameAdvanced struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameAdvanced"`
			NSImageNameApplicationIcon struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameApplicationIcon"`
			NSImageNameBluetoothTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameBluetoothTemplate"`
			NSImageNameBonjour struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameBonjour"`
			NSImageNameBookmarksTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameBookmarksTemplate"`
			NSImageNameCaution struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameCaution"`
			NSImageNameColorPanel struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameColorPanel"`
			NSImageNameColumnViewTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameColumnViewTemplate"`
			NSImageNameComputer struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameComputer"`
			NSImageNameDotMac struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameDotMac"`
			NSImageNameEnterFullScreenTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameEnterFullScreenTemplate"`
			NSImageNameEveryone struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameEveryone"`
			NSImageNameExitFullScreenTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameExitFullScreenTemplate"`
			NSImageNameFlowViewTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFlowViewTemplate"`
			NSImageNameFolder struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFolder"`
			NSImageNameFolderBurnable struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFolderBurnable"`
			NSImageNameFolderSmart struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFolderSmart"`
			NSImageNameFollowLinkFreestandingTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFollowLinkFreestandingTemplate"`
			NSImageNameFontPanel struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameFontPanel"`
			NSImageNameGoLeftTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameGoLeftTemplate"`
			NSImageNameGoRightTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameGoRightTemplate"`
			NSImageNameHomeTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameHomeTemplate"`
			NSImageNameIChatTheaterTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameIChatTheaterTemplate"`
			NSImageNameIconViewTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameIconViewTemplate"`
			NSImageNameInfo struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameInfo"`
			NSImageNameInvalidDataFreestandingTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameInvalidDataFreestandingTemplate"`
			NSImageNameLeftFacingTriangleTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameLeftFacingTriangleTemplate"`
			NSImageNameListViewTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameListViewTemplate"`
			NSImageNameLockLockedTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameLockLockedTemplate"`
			NSImageNameLockUnlockedTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameLockUnlockedTemplate"`
			NSImageNameMenuMixedStateTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameMenuMixedStateTemplate"`
			NSImageNameMenuOnStateTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameMenuOnStateTemplate"`
			NSImageNameMobileMe struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameMobileMe"`
			NSImageNameMultipleDocuments struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameMultipleDocuments"`
			NSImageNameNetwork struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameNetwork"`
			NSImageNamePathTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNamePathTemplate"`
			NSImageNamePreferencesGeneral struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNamePreferencesGeneral"`
			NSImageNameQuickLookTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameQuickLookTemplate"`
			NSImageNameRefreshFreestandingTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameRefreshFreestandingTemplate"`
			NSImageNameRefreshTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameRefreshTemplate"`
			NSImageNameRemoveTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameRemoveTemplate"`
			NSImageNameRevealFreestandingTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameRevealFreestandingTemplate"`
			NSImageNameRightFacingTriangleTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameRightFacingTriangleTemplate"`
			NSImageNameShareTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameShareTemplate"`
			NSImageNameSlideshowTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameSlideshowTemplate"`
			NSImageNameSmartBadgeTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameSmartBadgeTemplate"`
			NSImageNameStatusAvailable struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStatusAvailable"`
			NSImageNameStatusNone struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStatusNone"`
			NSImageNameStatusPartiallyAvailable struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStatusPartiallyAvailable"`
			NSImageNameStatusUnavailable struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStatusUnavailable"`
			NSImageNameStopProgressFreestandingTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStopProgressFreestandingTemplate"`
			NSImageNameStopProgressTemplate struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameStopProgressTemplate"`
			NSImageNameTrashEmpty struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameTrashEmpty"`
			NSImageNameTrashFull struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameTrashFull"`
			NSImageNameUser struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameUser"`
			NSImageNameUserAccounts struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameUserAccounts"`
			NSImageNameUserGroup struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameUserGroup"`
			NSImageNameUserGuest struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageNameUserGuest"`
			NSImageProgressive struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageProgressive"`
			NSImageRGBColorTable struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageRGBColorTable"`
			NSImageRepRegistryDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSImageRepRegistryDidChangeNotification"`
			NSIncludedKeysBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSIncludedKeysBinding"`
			NSInitialKeyBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSInitialKeyBinding"`
			NSInitialValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSInitialValueBinding"`
			NSInkTextPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSInkTextPboardType"`
			NSInsertsNullPlaceholderBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSInsertsNullPlaceholderBindingOption"`
			NSInterfaceStyleDefault struct {
				Typestr string `json:"typestr"`
			} `json:"NSInterfaceStyleDefault"`
			NSInvokesSeparatelyWithArrayObjectsBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSInvokesSeparatelyWithArrayObjectsBindingOption"`
			NSIsIndeterminateBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSIsIndeterminateBinding"`
			NSKernAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSKernAttributeName"`
			NSKeywordsDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSKeywordsDocumentAttribute"`
			NSLabelBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSLabelBinding"`
			NSLeftMarginDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSLeftMarginDocumentAttribute"`
			NSLigatureAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSLigatureAttributeName"`
			NSLightGray struct {
				Typestr string `json:"typestr"`
			} `json:"NSLightGray"`
			NSLinkAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSLinkAttributeName"`
			NSLocalizedKeyDictionaryBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSLocalizedKeyDictionaryBinding"`
			NSMacSimpleTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSMacSimpleTextDocumentType"`
			NSManagedObjectContextBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSManagedObjectContextBinding"`
			NSManagerDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSManagerDocumentAttribute"`
			NSMarkedClauseSegmentAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSMarkedClauseSegmentAttributeName"`
			NSMaxValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMaxValueBinding"`
			NSMaxWidthBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMaxWidthBinding"`
			NSMaximumRecentsBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMaximumRecentsBinding"`
			NSMenuDidAddItemNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidAddItemNotification"`
			NSMenuDidBeginTrackingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidBeginTrackingNotification"`
			NSMenuDidChangeItemNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidChangeItemNotification"`
			NSMenuDidEndTrackingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidEndTrackingNotification"`
			NSMenuDidRemoveItemNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidRemoveItemNotification"`
			NSMenuDidSendActionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuDidSendActionNotification"`
			NSMenuWillSendActionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSMenuWillSendActionNotification"`
			NSMinValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMinValueBinding"`
			NSMinWidthBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMinWidthBinding"`
			NSMixedStateImageBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSMixedStateImageBinding"`
			NSModalPanelRunLoopMode struct {
				Typestr string `json:"typestr"`
			} `json:"NSModalPanelRunLoopMode"`
			NSModificationTimeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSModificationTimeDocumentAttribute"`
			NSMultipleTextSelectionPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSMultipleTextSelectionPboardType"`
			NSMultipleValuesMarker struct {
				Typestr string `json:"typestr"`
			} `json:"NSMultipleValuesMarker"`
			NSMultipleValuesPlaceholderBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSMultipleValuesPlaceholderBindingOption"`
			NSNamedColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSNamedColorSpace"`
			NSNibLoadingException struct {
				Typestr string `json:"typestr"`
			} `json:"NSNibLoadingException"`
			NSNibOwner struct {
				Typestr string `json:"typestr"`
			} `json:"NSNibOwner"`
			NSNibTopLevelObjects struct {
				Typestr string `json:"typestr"`
			} `json:"NSNibTopLevelObjects"`
			NSNoSelectionMarker struct {
				Typestr string `json:"typestr"`
			} `json:"NSNoSelectionMarker"`
			NSNoSelectionPlaceholderBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSNoSelectionPlaceholderBindingOption"`
			NSNotApplicableMarker struct {
				Typestr string `json:"typestr"`
			} `json:"NSNotApplicableMarker"`
			NSNotApplicablePlaceholderBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSNotApplicablePlaceholderBindingOption"`
			NSNullPlaceholderBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSNullPlaceholderBindingOption"`
			NSObliquenessAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSObliquenessAttributeName"`
			NSObservedKeyPathKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSObservedKeyPathKey"`
			NSObservedObjectKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSObservedObjectKey"`
			NSOffStateImageBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSOffStateImageBinding"`
			NSOfficeOpenXMLTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSOfficeOpenXMLTextDocumentType"`
			NSOnStateImageBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSOnStateImageBinding"`
			NSOpenDocumentTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSOpenDocumentTextDocumentType"`
			NSOptionsKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSOptionsKey"`
			NSOutlineViewColumnDidMoveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewColumnDidMoveNotification"`
			NSOutlineViewColumnDidResizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewColumnDidResizeNotification"`
			NSOutlineViewDisclosureButtonKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewDisclosureButtonKey"`
			NSOutlineViewItemDidCollapseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewItemDidCollapseNotification"`
			NSOutlineViewItemDidExpandNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewItemDidExpandNotification"`
			NSOutlineViewItemWillCollapseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewItemWillCollapseNotification"`
			NSOutlineViewItemWillExpandNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewItemWillExpandNotification"`
			NSOutlineViewSelectionDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewSelectionDidChangeNotification"`
			NSOutlineViewSelectionIsChangingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewSelectionIsChangingNotification"`
			NSOutlineViewShowHideButtonKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSOutlineViewShowHideButtonKey"`
			NSPDFPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSPDFPboardType"`
			NSPICTPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSPICTPboardType"`
			NSPPDIncludeNotFoundException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPPDIncludeNotFoundException"`
			NSPPDIncludeStackOverflowException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPPDIncludeStackOverflowException"`
			NSPPDIncludeStackUnderflowException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPPDIncludeStackUnderflowException"`
			NSPPDParseException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPPDParseException"`
			NSPaperSizeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSPaperSizeDocumentAttribute"`
			NSParagraphStyleAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSParagraphStyleAttributeName"`
			NSPasteboardCommunicationException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardCommunicationException"`
			NSPasteboardTypeColor struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeColor"`
			NSPasteboardTypeFindPanelSearchOptions struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeFindPanelSearchOptions"`
			NSPasteboardTypeFont struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeFont"`
			NSPasteboardTypeHTML struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeHTML"`
			NSPasteboardTypeMultipleTextSelection struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeMultipleTextSelection"`
			NSPasteboardTypePDF struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypePDF"`
			NSPasteboardTypePNG struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypePNG"`
			NSPasteboardTypeRTF struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeRTF"`
			NSPasteboardTypeRTFD struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeRTFD"`
			NSPasteboardTypeRuler struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeRuler"`
			NSPasteboardTypeSound struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeSound"`
			NSPasteboardTypeString struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeString"`
			NSPasteboardTypeTIFF struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeTIFF"`
			NSPasteboardTypeTabularText struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeTabularText"`
			NSPasteboardTypeTextFinderOptions struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardTypeTextFinderOptions"`
			NSPasteboardURLReadingContentsConformToTypesKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardURLReadingContentsConformToTypesKey"`
			NSPasteboardURLReadingFileURLsOnlyKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSPasteboardURLReadingFileURLsOnlyKey"`
			NSPatternColorSpace struct {
				Typestr string `json:"typestr"`
			} `json:"NSPatternColorSpace"`
			NSPlainFileType struct {
				Typestr string `json:"typestr"`
			} `json:"NSPlainFileType"`
			NSPlainTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSPlainTextDocumentType"`
			NSPopUpButtonCellWillPopUpNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopUpButtonCellWillPopUpNotification"`
			NSPopUpButtonWillPopUpNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopUpButtonWillPopUpNotification"`
			NSPopoverCloseReasonDetachToWindow struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverCloseReasonDetachToWindow"`
			NSPopoverCloseReasonKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverCloseReasonKey"`
			NSPopoverCloseReasonStandard struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverCloseReasonStandard"`
			NSPopoverDidCloseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverDidCloseNotification"`
			NSPopoverDidShowNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverDidShowNotification"`
			NSPopoverWillCloseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverWillCloseNotification"`
			NSPopoverWillShowNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPopoverWillShowNotification"`
			NSPositioningRectBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSPositioningRectBinding"`
			NSPostScriptPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSPostScriptPboardType"`
			NSPredicateBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSPredicateBinding"`
			NSPredicateFormatBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSPredicateFormatBindingOption"`
			NSPreferredScrollerStyleDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSPreferredScrollerStyleDidChangeNotification"`
			NSPrefixSpacesDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrefixSpacesDocumentAttribute"`
			NSPrintAllPages struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintAllPages"`
			NSPrintAllPresetsJobStyleHint struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintAllPresetsJobStyleHint"`
			NSPrintBottomMargin struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintBottomMargin"`
			NSPrintCancelJob struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintCancelJob"`
			NSPrintCopies struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintCopies"`
			NSPrintDetailedErrorReporting struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintDetailedErrorReporting"`
			NSPrintFaxNumber struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintFaxNumber"`
			NSPrintFirstPage struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintFirstPage"`
			NSPrintFormName struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintFormName"`
			NSPrintHeaderAndFooter struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintHeaderAndFooter"`
			NSPrintHorizontalPagination struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintHorizontalPagination"`
			NSPrintHorizontallyCentered struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintHorizontallyCentered"`
			NSPrintJobDisposition struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintJobDisposition"`
			NSPrintJobFeatures struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintJobFeatures"`
			NSPrintJobSavingFileNameExtensionHidden struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintJobSavingFileNameExtensionHidden"`
			NSPrintJobSavingURL struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintJobSavingURL"`
			NSPrintLastPage struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintLastPage"`
			NSPrintLeftMargin struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintLeftMargin"`
			NSPrintManualFeed struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintManualFeed"`
			NSPrintMustCollate struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintMustCollate"`
			NSPrintNoPresetsJobStyleHint struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintNoPresetsJobStyleHint"`
			NSPrintOperationExistsException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintOperationExistsException"`
			NSPrintOrientation struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintOrientation"`
			NSPrintPackageException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPackageException"`
			NSPrintPagesAcross struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPagesAcross"`
			NSPrintPagesDown struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPagesDown"`
			NSPrintPagesPerSheet struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPagesPerSheet"`
			NSPrintPanelAccessorySummaryItemDescriptionKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPanelAccessorySummaryItemDescriptionKey"`
			NSPrintPanelAccessorySummaryItemNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPanelAccessorySummaryItemNameKey"`
			NSPrintPaperFeed struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPaperFeed"`
			NSPrintPaperName struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPaperName"`
			NSPrintPaperSize struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPaperSize"`
			NSPrintPhotoJobStyleHint struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPhotoJobStyleHint"`
			NSPrintPreviewJob struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPreviewJob"`
			NSPrintPrinter struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPrinter"`
			NSPrintPrinterName struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintPrinterName"`
			NSPrintReversePageOrder struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintReversePageOrder"`
			NSPrintRightMargin struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintRightMargin"`
			NSPrintSaveJob struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintSaveJob"`
			NSPrintSavePath struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintSavePath"`
			NSPrintScalingFactor struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintScalingFactor"`
			NSPrintSelectionOnly struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintSelectionOnly"`
			NSPrintSpoolJob struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintSpoolJob"`
			NSPrintTime struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintTime"`
			NSPrintTopMargin struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintTopMargin"`
			NSPrintVerticalPagination struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintVerticalPagination"`
			NSPrintVerticallyCentered struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintVerticallyCentered"`
			NSPrintingCommunicationException struct {
				Typestr string `json:"typestr"`
			} `json:"NSPrintingCommunicationException"`
			NSRTFDPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRTFDPboardType"`
			NSRTFDTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRTFDTextDocumentType"`
			NSRTFPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRTFPboardType"`
			NSRTFPropertyStackOverflowException struct {
				Typestr string `json:"typestr"`
			} `json:"NSRTFPropertyStackOverflowException"`
			NSRTFTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRTFTextDocumentType"`
			NSRaisesForNotApplicableKeysBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSRaisesForNotApplicableKeysBindingOption"`
			NSReadOnlyDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSReadOnlyDocumentAttribute"`
			NSRecentSearchesBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSRecentSearchesBinding"`
			NSRepresentedFilenameBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSRepresentedFilenameBinding"`
			NSRightMarginDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSRightMarginDocumentAttribute"`
			NSRowHeightBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSRowHeightBinding"`
			NSRuleEditorPredicateComparisonModifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateComparisonModifier"`
			NSRuleEditorPredicateCompoundType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateCompoundType"`
			NSRuleEditorPredicateCustomSelector struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateCustomSelector"`
			NSRuleEditorPredicateLeftExpression struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateLeftExpression"`
			NSRuleEditorPredicateOperatorType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateOperatorType"`
			NSRuleEditorPredicateOptions struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateOptions"`
			NSRuleEditorPredicateRightExpression struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorPredicateRightExpression"`
			NSRuleEditorRowsDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSRuleEditorRowsDidChangeNotification"`
			NSRulerPboard struct {
				Typestr string `json:"typestr"`
			} `json:"NSRulerPboard"`
			NSRulerPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSRulerPboardType"`
			NSScreenColorSpaceDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScreenColorSpaceDidChangeNotification"`
			NSScrollViewDidEndLiveMagnifyNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScrollViewDidEndLiveMagnifyNotification"`
			NSScrollViewDidEndLiveScrollNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScrollViewDidEndLiveScrollNotification"`
			NSScrollViewDidLiveScrollNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScrollViewDidLiveScrollNotification"`
			NSScrollViewWillStartLiveMagnifyNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScrollViewWillStartLiveMagnifyNotification"`
			NSScrollViewWillStartLiveScrollNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSScrollViewWillStartLiveScrollNotification"`
			NSSelectedIdentifierBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedIdentifierBinding"`
			NSSelectedIndexBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedIndexBinding"`
			NSSelectedLabelBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedLabelBinding"`
			NSSelectedObjectBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedObjectBinding"`
			NSSelectedObjectsBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedObjectsBinding"`
			NSSelectedTagBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedTagBinding"`
			NSSelectedValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedValueBinding"`
			NSSelectedValuesBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectedValuesBinding"`
			NSSelectionIndexPathsBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectionIndexPathsBinding"`
			NSSelectionIndexesBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectionIndexesBinding"`
			NSSelectorNameBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectorNameBindingOption"`
			NSSelectsAllWhenSettingContentBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSSelectsAllWhenSettingContentBindingOption"`
			NSShadowAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSShadowAttributeName"`
			NSSharingServiceNameAddToAperture struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameAddToAperture"`
			NSSharingServiceNameAddToIPhoto struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameAddToIPhoto"`
			NSSharingServiceNameAddToSafariReadingList struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameAddToSafariReadingList"`
			NSSharingServiceNameComposeEmail struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameComposeEmail"`
			NSSharingServiceNameComposeMessage struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameComposeMessage"`
			NSSharingServiceNamePostImageOnFlickr struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostImageOnFlickr"`
			NSSharingServiceNamePostOnFacebook struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostOnFacebook"`
			NSSharingServiceNamePostOnLinkedIn struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostOnLinkedIn"`
			NSSharingServiceNamePostOnSinaWeibo struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostOnSinaWeibo"`
			NSSharingServiceNamePostOnTencentWeibo struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostOnTencentWeibo"`
			NSSharingServiceNamePostOnTwitter struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostOnTwitter"`
			NSSharingServiceNamePostVideoOnTudou struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostVideoOnTudou"`
			NSSharingServiceNamePostVideoOnVimeo struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostVideoOnVimeo"`
			NSSharingServiceNamePostVideoOnYouku struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNamePostVideoOnYouku"`
			NSSharingServiceNameSendViaAirDrop struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameSendViaAirDrop"`
			NSSharingServiceNameUseAsDesktopPicture struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameUseAsDesktopPicture"`
			NSSharingServiceNameUseAsFacebookProfileImage struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameUseAsFacebookProfileImage"`
			NSSharingServiceNameUseAsLinkedInProfileImage struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameUseAsLinkedInProfileImage"`
			NSSharingServiceNameUseAsTwitterProfileImage struct {
				Typestr string `json:"typestr"`
			} `json:"NSSharingServiceNameUseAsTwitterProfileImage"`
			NSShellCommandFileType struct {
				Typestr string `json:"typestr"`
			} `json:"NSShellCommandFileType"`
			NSSortDescriptorsBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSSortDescriptorsBinding"`
			NSSoundPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSSoundPboardType"`
			NSSpeechCharacterModeProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechCharacterModeProperty"`
			NSSpeechCommandDelimiterProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechCommandDelimiterProperty"`
			NSSpeechCommandPrefix struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechCommandPrefix"`
			NSSpeechCommandSuffix struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechCommandSuffix"`
			NSSpeechCurrentVoiceProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechCurrentVoiceProperty"`
			NSSpeechDictionaryAbbreviations struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryAbbreviations"`
			NSSpeechDictionaryEntryPhonemes struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryEntryPhonemes"`
			NSSpeechDictionaryEntrySpelling struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryEntrySpelling"`
			NSSpeechDictionaryLocaleIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryLocaleIdentifier"`
			NSSpeechDictionaryModificationDate struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryModificationDate"`
			NSSpeechDictionaryPronunciations struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechDictionaryPronunciations"`
			NSSpeechErrorCount struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorCount"`
			NSSpeechErrorNewestCharacterOffset struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorNewestCharacterOffset"`
			NSSpeechErrorNewestCode struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorNewestCode"`
			NSSpeechErrorOldestCharacterOffset struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorOldestCharacterOffset"`
			NSSpeechErrorOldestCode struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorOldestCode"`
			NSSpeechErrorsProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechErrorsProperty"`
			NSSpeechInputModeProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechInputModeProperty"`
			NSSpeechModeLiteral struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechModeLiteral"`
			NSSpeechModeNormal struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechModeNormal"`
			NSSpeechModePhoneme struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechModePhoneme"`
			NSSpeechModeText struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechModeText"`
			NSSpeechNumberModeProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechNumberModeProperty"`
			NSSpeechOutputToFileURLProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechOutputToFileURLProperty"`
			NSSpeechPhonemeInfoExample struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeInfoExample"`
			NSSpeechPhonemeInfoHiliteEnd struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeInfoHiliteEnd"`
			NSSpeechPhonemeInfoHiliteStart struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeInfoHiliteStart"`
			NSSpeechPhonemeInfoOpcode struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeInfoOpcode"`
			NSSpeechPhonemeInfoSymbol struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeInfoSymbol"`
			NSSpeechPhonemeSymbolsProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPhonemeSymbolsProperty"`
			NSSpeechPitchBaseProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPitchBaseProperty"`
			NSSpeechPitchModProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechPitchModProperty"`
			NSSpeechRateProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechRateProperty"`
			NSSpeechRecentSyncProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechRecentSyncProperty"`
			NSSpeechResetProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechResetProperty"`
			NSSpeechStatusNumberOfCharactersLeft struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechStatusNumberOfCharactersLeft"`
			NSSpeechStatusOutputBusy struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechStatusOutputBusy"`
			NSSpeechStatusOutputPaused struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechStatusOutputPaused"`
			NSSpeechStatusPhonemeCode struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechStatusPhonemeCode"`
			NSSpeechStatusProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechStatusProperty"`
			NSSpeechSynthesizerInfoIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechSynthesizerInfoIdentifier"`
			NSSpeechSynthesizerInfoProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechSynthesizerInfoProperty"`
			NSSpeechSynthesizerInfoVersion struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechSynthesizerInfoVersion"`
			NSSpeechVolumeProperty struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpeechVolumeProperty"`
			NSSpellCheckerDidChangeAutomaticDashSubstitutionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpellCheckerDidChangeAutomaticDashSubstitutionNotification"`
			NSSpellCheckerDidChangeAutomaticQuoteSubstitutionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpellCheckerDidChangeAutomaticQuoteSubstitutionNotification"`
			NSSpellCheckerDidChangeAutomaticSpellingCorrectionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpellCheckerDidChangeAutomaticSpellingCorrectionNotification"`
			NSSpellCheckerDidChangeAutomaticTextReplacementNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpellCheckerDidChangeAutomaticTextReplacementNotification"`
			NSSpellingStateAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSSpellingStateAttributeName"`
			NSSplitViewDidResizeSubviewsNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSplitViewDidResizeSubviewsNotification"`
			NSSplitViewWillResizeSubviewsNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSplitViewWillResizeSubviewsNotification"`
			NSStrikethroughColorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSStrikethroughColorAttributeName"`
			NSStrikethroughStyleAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSStrikethroughStyleAttributeName"`
			NSStringPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSStringPboardType"`
			NSStrokeColorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSStrokeColorAttributeName"`
			NSStrokeWidthAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSStrokeWidthAttributeName"`
			NSSubjectDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSSubjectDocumentAttribute"`
			NSSuperscriptAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSSuperscriptAttributeName"`
			NSSystemColorsDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSSystemColorsDidChangeNotification"`
			NSTIFFException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTIFFException"`
			NSTIFFPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSTIFFPboardType"`
			NSTabColumnTerminatorsAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSTabColumnTerminatorsAttributeName"`
			NSTableViewColumnDidMoveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTableViewColumnDidMoveNotification"`
			NSTableViewColumnDidResizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTableViewColumnDidResizeNotification"`
			NSTableViewRowViewKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTableViewRowViewKey"`
			NSTableViewSelectionDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTableViewSelectionDidChangeNotification"`
			NSTableViewSelectionIsChangingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTableViewSelectionIsChangingNotification"`
			NSTabularTextPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSTabularTextPboardType"`
			NSTargetBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSTargetBinding"`
			NSTextAlternativesAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextAlternativesAttributeName"`
			NSTextAlternativesSelectedAlternativeStringNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextAlternativesSelectedAlternativeStringNotification"`
			NSTextCheckingDocumentAuthorKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingDocumentAuthorKey"`
			NSTextCheckingDocumentTitleKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingDocumentTitleKey"`
			NSTextCheckingDocumentURLKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingDocumentURLKey"`
			NSTextCheckingOrthographyKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingOrthographyKey"`
			NSTextCheckingQuotesKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingQuotesKey"`
			NSTextCheckingReferenceDateKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingReferenceDateKey"`
			NSTextCheckingReferenceTimeZoneKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingReferenceTimeZoneKey"`
			NSTextCheckingRegularExpressionsKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingRegularExpressionsKey"`
			NSTextCheckingReplacementsKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextCheckingReplacementsKey"`
			NSTextColorBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextColorBinding"`
			NSTextDidBeginEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextDidBeginEditingNotification"`
			NSTextDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextDidChangeNotification"`
			NSTextDidEndEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextDidEndEditingNotification"`
			NSTextEffectAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextEffectAttributeName"`
			NSTextEffectLetterpressStyle struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextEffectLetterpressStyle"`
			NSTextEncodingNameDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextEncodingNameDocumentAttribute"`
			NSTextEncodingNameDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextEncodingNameDocumentOption"`
			NSTextFinderCaseInsensitiveKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextFinderCaseInsensitiveKey"`
			NSTextFinderMatchingTypeKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextFinderMatchingTypeKey"`
			NSTextInputContextKeyboardSelectionDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextInputContextKeyboardSelectionDidChangeNotification"`
			NSTextLayoutSectionOrientation struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextLayoutSectionOrientation"`
			NSTextLayoutSectionRange struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextLayoutSectionRange"`
			NSTextLayoutSectionsAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextLayoutSectionsAttribute"`
			NSTextLineTooLongException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextLineTooLongException"`
			NSTextNoSelectionException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextNoSelectionException"`
			NSTextReadException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextReadException"`
			NSTextSizeMultiplierDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextSizeMultiplierDocumentOption"`
			NSTextStorageDidProcessEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextStorageDidProcessEditingNotification"`
			NSTextStorageWillProcessEditingNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextStorageWillProcessEditingNotification"`
			NSTextViewDidChangeSelectionNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextViewDidChangeSelectionNotification"`
			NSTextViewDidChangeTypingAttributesNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextViewDidChangeTypingAttributesNotification"`
			NSTextViewWillChangeNotifyingTextViewNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextViewWillChangeNotifyingTextViewNotification"`
			NSTextWriteException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTextWriteException"`
			NSTimeoutDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSTimeoutDocumentOption"`
			NSTitleBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSTitleBinding"`
			NSTitleDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSTitleDocumentAttribute"`
			NSToolTipAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolTipAttributeName"`
			NSToolTipBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolTipBinding"`
			NSToolbarCustomizeToolbarItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarCustomizeToolbarItemIdentifier"`
			NSToolbarDidRemoveItemNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarDidRemoveItemNotification"`
			NSToolbarFlexibleSpaceItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarFlexibleSpaceItemIdentifier"`
			NSToolbarPrintItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarPrintItemIdentifier"`
			NSToolbarSeparatorItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarSeparatorItemIdentifier"`
			NSToolbarShowColorsItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarShowColorsItemIdentifier"`
			NSToolbarShowFontsItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarShowFontsItemIdentifier"`
			NSToolbarSpaceItemIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarSpaceItemIdentifier"`
			NSToolbarWillAddItemNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSToolbarWillAddItemNotification"`
			NSTopMarginDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSTopMarginDocumentAttribute"`
			NSTransparentBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSTransparentBinding"`
			NSTypeIdentifierAddressText struct {
				Typestr string `json:"typestr"`
			} `json:"NSTypeIdentifierAddressText"`
			NSTypeIdentifierDateText struct {
				Typestr string `json:"typestr"`
			} `json:"NSTypeIdentifierDateText"`
			NSTypeIdentifierPhoneNumberText struct {
				Typestr string `json:"typestr"`
			} `json:"NSTypeIdentifierPhoneNumberText"`
			NSTypeIdentifierTransitInformationText struct {
				Typestr string `json:"typestr"`
			} `json:"NSTypeIdentifierTransitInformationText"`
			NSTypedStreamVersionException struct {
				Typestr string `json:"typestr"`
			} `json:"NSTypedStreamVersionException"`
			NSURLPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSURLPboardType"`
			NSUnderlineByWordMask struct {
				Typestr string `json:"typestr"`
			} `json:"NSUnderlineByWordMask"`
			NSUnderlineColorAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSUnderlineColorAttributeName"`
			NSUnderlineStrikethroughMask struct {
				Typestr string `json:"typestr"`
			} `json:"NSUnderlineStrikethroughMask"`
			NSUnderlineStyleAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSUnderlineStyleAttributeName"`
			NSUserActivityDocumentURLKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSUserActivityDocumentURLKey"`
			NSUsesScreenFontsDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSUsesScreenFontsDocumentAttribute"`
			NSVCardPboardType struct {
				Typestr string `json:"typestr"`
			} `json:"NSVCardPboardType"`
			NSValidatesImmediatelyBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSValidatesImmediatelyBindingOption"`
			NSValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSValueBinding"`
			NSValuePathBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSValuePathBinding"`
			NSValueTransformerBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSValueTransformerBindingOption"`
			NSValueTransformerNameBindingOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSValueTransformerNameBindingOption"`
			NSValueURLBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSValueURLBinding"`
			NSVerticalGlyphFormAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSVerticalGlyphFormAttributeName"`
			NSViewAnimationEffectKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationEffectKey"`
			NSViewAnimationEndFrameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationEndFrameKey"`
			NSViewAnimationFadeInEffect struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationFadeInEffect"`
			NSViewAnimationFadeOutEffect struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationFadeOutEffect"`
			NSViewAnimationStartFrameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationStartFrameKey"`
			NSViewAnimationTargetKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewAnimationTargetKey"`
			NSViewBoundsDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewBoundsDidChangeNotification"`
			NSViewDidUpdateTrackingAreasNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewDidUpdateTrackingAreasNotification"`
			NSViewFocusDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewFocusDidChangeNotification"`
			NSViewFrameDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewFrameDidChangeNotification"`
			NSViewGlobalFrameDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewGlobalFrameDidChangeNotification"`
			NSViewModeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewModeDocumentAttribute"`
			NSViewNoInstrinsicMetric struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewNoInstrinsicMetric"`
			NSViewSizeDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewSizeDocumentAttribute"`
			NSViewZoomDocumentAttribute struct {
				Typestr string `json:"typestr"`
			} `json:"NSViewZoomDocumentAttribute"`
			NSVisibleBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSVisibleBinding"`
			NSVoiceAge struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceAge"`
			NSVoiceDemoText struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceDemoText"`
			NSVoiceGender struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceGender"`
			NSVoiceGenderFemale struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceGenderFemale"`
			NSVoiceGenderMale struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceGenderMale"`
			NSVoiceGenderNeuter struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceGenderNeuter"`
			NSVoiceIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceIdentifier"`
			NSVoiceIndividuallySpokenCharacters struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceIndividuallySpokenCharacters"`
			NSVoiceLanguage struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceLanguage"`
			NSVoiceLocaleIdentifier struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceLocaleIdentifier"`
			NSVoiceName struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceName"`
			NSVoiceSupportedCharacters struct {
				Typestr string `json:"typestr"`
			} `json:"NSVoiceSupportedCharacters"`
			NSWarningValueBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSWarningValueBinding"`
			NSWebArchiveTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSWebArchiveTextDocumentType"`
			NSWebPreferencesDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSWebPreferencesDocumentOption"`
			NSWebResourceLoadDelegateDocumentOption struct {
				Typestr string `json:"typestr"`
			} `json:"NSWebResourceLoadDelegateDocumentOption"`
			NSWhite struct {
				Typestr string `json:"typestr"`
			} `json:"NSWhite"`
			NSWidthBinding struct {
				Typestr string `json:"typestr"`
			} `json:"NSWidthBinding"`
			NSWindowDidBecomeKeyNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidBecomeKeyNotification"`
			NSWindowDidBecomeMainNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidBecomeMainNotification"`
			NSWindowDidChangeBackingPropertiesNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidChangeBackingPropertiesNotification"`
			NSWindowDidChangeOcclusionStateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidChangeOcclusionStateNotification"`
			NSWindowDidChangeScreenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidChangeScreenNotification"`
			NSWindowDidChangeScreenProfileNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidChangeScreenProfileNotification"`
			NSWindowDidDeminiaturizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidDeminiaturizeNotification"`
			NSWindowDidEndLiveResizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidEndLiveResizeNotification"`
			NSWindowDidEndSheetNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidEndSheetNotification"`
			NSWindowDidEnterFullScreenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidEnterFullScreenNotification"`
			NSWindowDidEnterVersionBrowserNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidEnterVersionBrowserNotification"`
			NSWindowDidExitFullScreenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidExitFullScreenNotification"`
			NSWindowDidExitVersionBrowserNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidExitVersionBrowserNotification"`
			NSWindowDidExposeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidExposeNotification"`
			NSWindowDidMiniaturizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidMiniaturizeNotification"`
			NSWindowDidMoveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidMoveNotification"`
			NSWindowDidResignKeyNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidResignKeyNotification"`
			NSWindowDidResignMainNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidResignMainNotification"`
			NSWindowDidResizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidResizeNotification"`
			NSWindowDidUpdateNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowDidUpdateNotification"`
			NSWindowServerCommunicationException struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowServerCommunicationException"`
			NSWindowWillBeginSheetNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillBeginSheetNotification"`
			NSWindowWillCloseNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillCloseNotification"`
			NSWindowWillEnterFullScreenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillEnterFullScreenNotification"`
			NSWindowWillEnterVersionBrowserNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillEnterVersionBrowserNotification"`
			NSWindowWillExitFullScreenNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillExitFullScreenNotification"`
			NSWindowWillExitVersionBrowserNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillExitVersionBrowserNotification"`
			NSWindowWillMiniaturizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillMiniaturizeNotification"`
			NSWindowWillMoveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillMoveNotification"`
			NSWindowWillStartLiveResizeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWindowWillStartLiveResizeNotification"`
			NSWordMLTextDocumentType struct {
				Typestr string `json:"typestr"`
			} `json:"NSWordMLTextDocumentType"`
			NSWordTablesReadException struct {
				Typestr string `json:"typestr"`
			} `json:"NSWordTablesReadException"`
			NSWordTablesWriteException struct {
				Typestr string `json:"typestr"`
			} `json:"NSWordTablesWriteException"`
			NSWorkspaceAccessibilityDisplayOptionsDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceAccessibilityDisplayOptionsDidChangeNotification"`
			NSWorkspaceActiveSpaceDidChangeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceActiveSpaceDidChangeNotification"`
			NSWorkspaceApplicationKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceApplicationKey"`
			NSWorkspaceCompressOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceCompressOperation"`
			NSWorkspaceCopyOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceCopyOperation"`
			NSWorkspaceDecompressOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDecompressOperation"`
			NSWorkspaceDecryptOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDecryptOperation"`
			NSWorkspaceDesktopImageAllowClippingKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDesktopImageAllowClippingKey"`
			NSWorkspaceDesktopImageFillColorKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDesktopImageFillColorKey"`
			NSWorkspaceDesktopImageScalingKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDesktopImageScalingKey"`
			NSWorkspaceDestroyOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDestroyOperation"`
			NSWorkspaceDidActivateApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidActivateApplicationNotification"`
			NSWorkspaceDidChangeFileLabelsNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidChangeFileLabelsNotification"`
			NSWorkspaceDidDeactivateApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidDeactivateApplicationNotification"`
			NSWorkspaceDidHideApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidHideApplicationNotification"`
			NSWorkspaceDidLaunchApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidLaunchApplicationNotification"`
			NSWorkspaceDidMountNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidMountNotification"`
			NSWorkspaceDidPerformFileOperationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidPerformFileOperationNotification"`
			NSWorkspaceDidRenameVolumeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidRenameVolumeNotification"`
			NSWorkspaceDidTerminateApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidTerminateApplicationNotification"`
			NSWorkspaceDidUnhideApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidUnhideApplicationNotification"`
			NSWorkspaceDidUnmountNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidUnmountNotification"`
			NSWorkspaceDidWakeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDidWakeNotification"`
			NSWorkspaceDuplicateOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceDuplicateOperation"`
			NSWorkspaceEncryptOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceEncryptOperation"`
			NSWorkspaceLaunchConfigurationAppleEvent struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceLaunchConfigurationAppleEvent"`
			NSWorkspaceLaunchConfigurationArchitecture struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceLaunchConfigurationArchitecture"`
			NSWorkspaceLaunchConfigurationArguments struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceLaunchConfigurationArguments"`
			NSWorkspaceLaunchConfigurationEnvironment struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceLaunchConfigurationEnvironment"`
			NSWorkspaceLinkOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceLinkOperation"`
			NSWorkspaceMoveOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceMoveOperation"`
			NSWorkspaceRecycleOperation struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceRecycleOperation"`
			NSWorkspaceScreensDidSleepNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceScreensDidSleepNotification"`
			NSWorkspaceScreensDidWakeNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceScreensDidWakeNotification"`
			NSWorkspaceSessionDidBecomeActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceSessionDidBecomeActiveNotification"`
			NSWorkspaceSessionDidResignActiveNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceSessionDidResignActiveNotification"`
			NSWorkspaceVolumeLocalizedNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceVolumeLocalizedNameKey"`
			NSWorkspaceVolumeOldLocalizedNameKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceVolumeOldLocalizedNameKey"`
			NSWorkspaceVolumeOldURLKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceVolumeOldURLKey"`
			NSWorkspaceVolumeURLKey struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceVolumeURLKey"`
			NSWorkspaceWillLaunchApplicationNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceWillLaunchApplicationNotification"`
			NSWorkspaceWillPowerOffNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceWillPowerOffNotification"`
			NSWorkspaceWillSleepNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceWillSleepNotification"`
			NSWorkspaceWillUnmountNotification struct {
				Typestr string `json:"typestr"`
			} `json:"NSWorkspaceWillUnmountNotification"`
			NSWritingDirectionAttributeName struct {
				Typestr string `json:"typestr"`
			} `json:"NSWritingDirectionAttributeName"`
		} `json:"externs"`
		FormalProtocols struct {
			NSAccessibility struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []interface{} `json:"attributes"`
					Name           string        `json:"name"`
					Typestr        string        `json:"typestr"`
					TypestrSpecial bool          `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSAccessibility"`
			NSAccessibilityButton struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityButton"`
			NSAccessibilityCheckBox struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityCheckBox"`
			NSAccessibilityContainsTransientUI struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityContainsTransientUI"`
			NSAccessibilityElement struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityElement"`
			NSAccessibilityGroup struct {
				Implements []string      `json:"implements"`
				Methods    []interface{} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityGroup"`
			NSAccessibilityImage struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityImage"`
			NSAccessibilityLayoutArea struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityLayoutArea"`
			NSAccessibilityLayoutItem struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityLayoutItem"`
			NSAccessibilityList struct {
				Implements []string      `json:"implements"`
				Methods    []interface{} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityList"`
			NSAccessibilityNavigableStaticText struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityNavigableStaticText"`
			NSAccessibilityOutline struct {
				Implements []string      `json:"implements"`
				Methods    []interface{} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityOutline"`
			NSAccessibilityProgressIndicator struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityProgressIndicator"`
			NSAccessibilityRadioButton struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityRadioButton"`
			NSAccessibilityRow struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityRow"`
			NSAccessibilitySlider struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilitySlider"`
			NSAccessibilityStaticText struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityStaticText"`
			NSAccessibilityStepper struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityStepper"`
			NSAccessibilitySwitch struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilitySwitch"`
			NSAccessibilityTable struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityTable"`
			NSAlertDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAlertDelegate"`
			NSAnimatablePropertyContainer struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAnimatablePropertyContainer"`
			NSAnimationDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAnimationDelegate"`
			NSAppearanceCustomization struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSAppearanceCustomization"`
			NSApplicationDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSApplicationDelegate"`
			NSBrowserDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSBrowserDelegate"`
			NSChangeSpelling struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSChangeSpelling"`
			NSCollectionViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSCollectionViewDelegate"`
			NSColorPickingCustom struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSColorPickingCustom"`
			NSColorPickingDefault struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSColorPickingDefault"`
			NSComboBoxCellDataSource struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSComboBoxCellDataSource"`
			NSComboBoxDataSource struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSComboBoxDataSource"`
			NSComboBoxDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSComboBoxDelegate"`
			NSControlTextEditingDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSControlTextEditingDelegate"`
			NSDatePickerCellDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDatePickerCellDelegate"`
			NSDockTilePlugIn struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDockTilePlugIn"`
			NSDraggingDestination struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDraggingDestination"`
			NSDraggingInfo struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Name           string `json:"name"`
					Typestr        string `json:"typestr"`
					TypestrSpecial bool   `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSDraggingInfo"`
			NSDraggingSource struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDraggingSource"`
			NSDrawerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDrawerDelegate"`
			NSGestureRecognizerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSGestureRecognizerDelegate"`
			NSGlyphStorage struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSGlyphStorage"`
			NSIgnoreMisspelledWords struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSIgnoreMisspelledWords"`
			NSImageDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSImageDelegate"`
			NSInputServerMouseTracker struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSInputServerMouseTracker"`
			NSInputServiceProvider struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSInputServiceProvider"`
			NSLayoutManagerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSLayoutManagerDelegate"`
			NSMatrixDelegate struct {
				Implements []string      `json:"implements"`
				Methods    []interface{} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSMatrixDelegate"`
			NSMenuDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSMenuDelegate"`
			NSOpenSavePanelDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSOpenSavePanelDelegate"`
			NSOutlineViewDataSource struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSOutlineViewDataSource"`
			NSOutlineViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSOutlineViewDelegate"`
			NSPageControllerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPageControllerDelegate"`
			NSPasteboardItemDataProvider struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPasteboardItemDataProvider"`
			NSPasteboardReading struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPasteboardReading"`
			NSPasteboardWriting struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPasteboardWriting"`
			NSPathCellDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPathCellDelegate"`
			NSPathControlDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPathControlDelegate"`
			NSPopoverDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPopoverDelegate"`
			NSPrintPanelAccessorizing struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPrintPanelAccessorizing"`
			NSRuleEditorDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSRuleEditorDelegate"`
			NSSeguePerforming struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSeguePerforming"`
			NSServicesMenuRequestor struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSServicesMenuRequestor"`
			NSSharingServiceDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSharingServiceDelegate"`
			NSSharingServicePickerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSharingServicePickerDelegate"`
			NSSoundDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSoundDelegate"`
			NSSpeechRecognizerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSpeechRecognizerDelegate"`
			NSSpeechSynthesizerDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSpeechSynthesizerDelegate"`
			NSSplitViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSplitViewDelegate"`
			NSStackViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSStackViewDelegate"`
			NSTabViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTabViewDelegate"`
			NSTableViewDataSource struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTableViewDataSource"`
			NSTableViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTableViewDelegate"`
			NSTextAttachmentCell struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextAttachmentCell"`
			NSTextDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextDelegate"`
			NSTextFieldDelegate struct {
				Implements []string      `json:"implements"`
				Methods    []interface{} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextFieldDelegate"`
			NSTextFinderBarContainer struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     [][]string `json:"attributes"`
					Name           string     `json:"name"`
					Typestr        string     `json:"typestr"`
					TypestrSpecial bool       `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSTextFinderBarContainer"`
			NSTextFinderClient struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSTextFinderClient"`
			NSTextInput struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextInput"`
			NSTextInputClient struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextInputClient"`
			NSTextLayoutOrientationProvider struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextLayoutOrientationProvider"`
			NSTextStorageDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextStorageDelegate"`
			NSTextViewDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTextViewDelegate"`
			NSTokenFieldCellDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTokenFieldCellDelegate"`
			NSTokenFieldDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTokenFieldDelegate"`
			NSToolbarDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSToolbarDelegate"`
			NSUserInterfaceItemIdentification struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSUserInterfaceItemIdentification"`
			NSUserInterfaceItemSearching struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSUserInterfaceItemSearching"`
			NSUserInterfaceValidations struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSUserInterfaceValidations"`
			NSValidatedUserInterfaceItem struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Required    bool          `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSValidatedUserInterfaceItem"`
			NSViewControllerPresentationAnimator struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSViewControllerPresentationAnimator"`
			NSWindowDelegate struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSWindowDelegate"`
			NSWindowRestoration struct {
				Implements []string `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Required    bool `json:"required"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSWindowRestoration"`
		} `json:"formal_protocols"`
		FuncMacros struct{} `json:"func_macros"`
		Functions  struct {
			NSAccessibilityActionDescription struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityActionDescription"`
			NSAccessibilityFrameInView struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityFrameInView"`
			NSAccessibilityPointInView struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityPointInView"`
			NSAccessibilityPostNotification struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityPostNotification"`
			NSAccessibilityPostNotificationWithUserInfo struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityPostNotificationWithUserInfo"`
			NSAccessibilityRaiseBadArgumentException struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityRaiseBadArgumentException"`
			NSAccessibilityRoleDescription struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityRoleDescription"`
			NSAccessibilityRoleDescriptionForUIElement struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityRoleDescriptionForUIElement"`
			NSAccessibilitySetMayContainProtectedContent struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilitySetMayContainProtectedContent"`
			NSAccessibilityUnignoredAncestor struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityUnignoredAncestor"`
			NSAccessibilityUnignoredChildren struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityUnignoredChildren"`
			NSAccessibilityUnignoredChildrenForOnlyChild struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityUnignoredChildrenForOnlyChild"`
			NSAccessibilityUnignoredDescendant struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAccessibilityUnignoredDescendant"`
			NSApplicationLoad struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSApplicationLoad"`
			NSApplicationMain struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSApplicationMain"`
			NSAvailableWindowDepths struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSAvailableWindowDepths"`
			NSBeep struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSBeep"`
			NSBeginAlertSheet struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSBeginAlertSheet"`
			NSBeginCriticalAlertSheet struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSBeginCriticalAlertSheet"`
			NSBeginInformationalAlertSheet struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSBeginInformationalAlertSheet"`
			NSBestDepth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSBestDepth"`
			NSBitsPerPixelFromDepth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSBitsPerPixelFromDepth"`
			NSBitsPerSampleFromDepth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSBitsPerSampleFromDepth"`
			NSColorSpaceFromDepth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSColorSpaceFromDepth"`
			NSConvertGlyphsToPackedGlyphs struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSConvertGlyphsToPackedGlyphs"`
			NSCopyBits struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					AlreadyCfretained bool   `json:"already_cfretained"`
					Typestr           string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSCopyBits"`
			NSCountWindows struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSCountWindows"`
			NSCountWindowsForContext struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSCountWindowsForContext"`
			NSCreateFileContentsPboardType struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					AlreadyCfretained bool   `json:"already_cfretained"`
					Typestr           string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSCreateFileContentsPboardType"`
			NSCreateFilenamePboardType struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					AlreadyCfretained bool   `json:"already_cfretained"`
					Typestr           string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSCreateFilenamePboardType"`
			NSDisableScreenUpdates struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDisableScreenUpdates"`
			NSDottedFrameRect struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDottedFrameRect"`
			NSDrawBitmap struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawBitmap"`
			NSDrawButton struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawButton"`
			NSDrawColorTiledRects struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawColorTiledRects"`
			NSDrawDarkBezel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawDarkBezel"`
			NSDrawGrayBezel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawGrayBezel"`
			NSDrawGroove struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawGroove"`
			NSDrawLightBezel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawLightBezel"`
			NSDrawNinePartImage struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawNinePartImage"`
			NSDrawThreePartImage struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawThreePartImage"`
			NSDrawTiledRects struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawTiledRects"`
			NSDrawWhiteBezel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawWhiteBezel"`
			NSDrawWindowBackground struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSDrawWindowBackground"`
			NSEnableScreenUpdates struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSEnableScreenUpdates"`
			NSEraseRect struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSEraseRect"`
			NSEventMaskFromType struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Inline bool `json:"inline"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSEventMaskFromType"`
			NSFrameRect struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSFrameRect"`
			NSFrameRectWithWidth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSFrameRectWithWidth"`
			NSFrameRectWithWidthUsingOperation struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSFrameRectWithWidthUsingOperation"`
			NSGetAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSGetAlertPanel"`
			NSGetCriticalAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSGetCriticalAlertPanel"`
			NSGetFileType struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSGetFileType"`
			NSGetFileTypes struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSGetFileTypes"`
			NSGetInformationalAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSGetInformationalAlertPanel"`
			NSGetWindowServerMemory struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSGetWindowServerMemory"`
			NSHighlightRect struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSHighlightRect"`
			NSInterfaceStyleForKey struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSInterfaceStyleForKey"`
			NSIsControllerMarker struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSIsControllerMarker"`
			NSNumberOfColorComponents struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSNumberOfColorComponents"`
			NSOpenGLGetOption struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSOpenGLGetOption"`
			NSOpenGLGetVersion struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSOpenGLGetVersion"`
			NSOpenGLSetOption struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSOpenGLSetOption"`
			NSPerformService struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSPerformService"`
			NSPlanarFromDepth struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSPlanarFromDepth"`
			NSReadPixel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSReadPixel"`
			NSRectClip struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectClip"`
			NSRectClipList struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectClipList"`
			NSRectFill struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFill"`
			NSRectFillList struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillList"`
			NSRectFillListUsingOperation struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillListUsingOperation"`
			NSRectFillListWithColors struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillListWithColors"`
			NSRectFillListWithColorsUsingOperation struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillListWithColorsUsingOperation"`
			NSRectFillListWithGrays struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillListWithGrays"`
			NSRectFillUsingOperation struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRectFillUsingOperation"`
			NSRegisterServicesProvider struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSRegisterServicesProvider"`
			NSReleaseAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSReleaseAlertPanel"`
			NSRunAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunAlertPanel"`
			NSRunAlertPanelRelativeToWindow struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunAlertPanelRelativeToWindow"`
			NSRunCriticalAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunCriticalAlertPanel"`
			NSRunCriticalAlertPanelRelativeToWindow struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunCriticalAlertPanelRelativeToWindow"`
			NSRunInformationalAlertPanel struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunInformationalAlertPanel"`
			NSRunInformationalAlertPanelRelativeToWindow struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"NSRunInformationalAlertPanelRelativeToWindow"`
			NSSetFocusRingStyle struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSSetFocusRingStyle"`
			NSSetShowsServicesMenuItem struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSSetShowsServicesMenuItem"`
			NSShowAnimationEffect struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSShowAnimationEffect"`
			NSShowsServicesMenuItem struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSShowsServicesMenuItem"`
			NSUnregisterServicesProvider struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSUnregisterServicesProvider"`
			NSUpdateDynamicServices struct {
				Args []struct {
					Name    interface{} `json:"name"`
					Typestr string      `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSUpdateDynamicServices"`
			NSWindowList struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSWindowList"`
			NSWindowListForContext struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
			} `json:"NSWindowListForContext"`
			NSDictionaryOfVariableBindings struct {
				Args []struct {
					Name    string `json:"name"`
					Typestr string `json:"typestr"`
				} `json:"args"`
				Retval struct {
					Typestr string `json:"typestr"`
				} `json:"retval"`
				Variadic bool `json:"variadic"`
			} `json:"_NSDictionaryOfVariableBindings"`
		} `json:"functions"`
		InformalProtocols struct {
			NSAccessibility struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSAccessibility"`
			NSAccessibilityAdditions struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSAccessibilityAdditions"`
			NSApplicationScriptingDelegation struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSApplicationScriptingDelegation"`
			NSColorPanelResponderMethod struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSColorPanelResponderMethod"`
			NSControlSubclassNotifications struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSControlSubclassNotifications"`
			NSDictionaryControllerKeyValuePair struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDictionaryControllerKeyValuePair"`
			NSDraggingSourceDeprecated struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSDraggingSourceDeprecated"`
			NSEditor struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSEditor"`
			NSEditorRegistration struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSEditorRegistration"`
			NSFontManagerDelegate struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSFontManagerDelegate"`
			NSFontManagerResponderMethod struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSFontManagerResponderMethod"`
			NSFontPanelValidationAdditions struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSFontPanelValidationAdditions"`
			NSKeyValueBindingCreation struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []struct {
					Attributes     []string `json:"attributes"`
					Name           string   `json:"name"`
					Typestr        string   `json:"typestr"`
					TypestrSpecial bool     `json:"typestr_special"`
				} `json:"properties"`
			} `json:"NSKeyValueBindingCreation"`
			NSLayerDelegateContentsScaleUpdating struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSLayerDelegateContentsScaleUpdating"`
			NSMenuValidation struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSMenuValidation"`
			NSNibAwaking struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args        []interface{} `json:"args"`
					ClassMethod bool          `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSNibAwaking"`
			NSPasteboardOwner struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPasteboardOwner"`
			NSPlaceholders struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSPlaceholders"`
			NSSavePanelDelegateDeprecated struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSSavePanelDelegateDeprecated"`
			NSTableViewDataSourceDeprecated struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSTableViewDataSourceDeprecated"`
			NSToolTipOwner struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSToolTipOwner"`
			NSToolbarItemValidation struct {
				Implements []interface{} `json:"implements"`
				Methods    []struct {
					Args []struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"args"`
					ClassMethod bool `json:"class_method"`
					Retval      struct {
						Typestr        string `json:"typestr"`
						TypestrSpecial bool   `json:"typestr_special"`
					} `json:"retval"`
					Selector   string `json:"selector"`
					Visibility string `json:"visibility"`
				} `json:"methods"`
				Properties []interface{} `json:"properties"`
			} `json:"NSToolbarItemValidation"`
		} `json:"informal_protocols"`
		Literals struct {
			NSAppKitVersionNumber10_2_3                           float64 `json:"NSAppKitVersionNumber10_2_3"`
			NSAppKitVersionNumber10_3_2                           float64 `json:"NSAppKitVersionNumber10_3_2"`
			NSAppKitVersionNumber10_3_3                           float64 `json:"NSAppKitVersionNumber10_3_3"`
			NSAppKitVersionNumber10_3_5                           float64 `json:"NSAppKitVersionNumber10_3_5"`
			NSAppKitVersionNumber10_3_7                           float64 `json:"NSAppKitVersionNumber10_3_7"`
			NSAppKitVersionNumber10_3_9                           float64 `json:"NSAppKitVersionNumber10_3_9"`
			NSAppKitVersionNumber10_4_1                           float64 `json:"NSAppKitVersionNumber10_4_1"`
			NSAppKitVersionNumber10_4_3                           float64 `json:"NSAppKitVersionNumber10_4_3"`
			NSAppKitVersionNumber10_4_4                           float64 `json:"NSAppKitVersionNumber10_4_4"`
			NSAppKitVersionNumber10_4_7                           float64 `json:"NSAppKitVersionNumber10_4_7"`
			NSAppKitVersionNumber10_5_2                           float64 `json:"NSAppKitVersionNumber10_5_2"`
			NSAppKitVersionNumber10_5_3                           float64 `json:"NSAppKitVersionNumber10_5_3"`
			NSAppKitVersionNumber10_7_2                           float64 `json:"NSAppKitVersionNumber10_7_2"`
			NSAppKitVersionNumber10_7_3                           float64 `json:"NSAppKitVersionNumber10_7_3"`
			NSAppKitVersionNumber10_7_4                           float64 `json:"NSAppKitVersionNumber10_7_4"`
			NSAppKitVersionNumberWithColumnResizingBrowser        int64   `json:"NSAppKitVersionNumberWithColumnResizingBrowser"`
			NSAppKitVersionNumberWithContinuousScrollingBrowser   int64   `json:"NSAppKitVersionNumberWithContinuousScrollingBrowser"`
			NSAppKitVersionNumberWithCursorSizeSupport            int64   `json:"NSAppKitVersionNumberWithCursorSizeSupport"`
			NSAppKitVersionNumberWithCustomSheetPosition          int64   `json:"NSAppKitVersionNumberWithCustomSheetPosition"`
			NSAppKitVersionNumberWithDeferredWindowDisplaySupport int64   `json:"NSAppKitVersionNumberWithDeferredWindowDisplaySupport"`
			NSAppKitVersionNumberWithDirectionalTabs              int64   `json:"NSAppKitVersionNumberWithDirectionalTabs"`
			NSAppKitVersionNumberWithDockTilePlugInSupport        int64   `json:"NSAppKitVersionNumberWithDockTilePlugInSupport"`
			NSAppKitVersionNumberWithPatternColorLeakFix          int64   `json:"NSAppKitVersionNumberWithPatternColorLeakFix"`
		} `json:"literals"`
		Structs struct {
			NSOpenGLContextAuxiliary struct {
				Fieldnames []interface{} `json:"fieldnames"`
				Special    bool          `json:"special"`
				Typestr    string        `json:"typestr"`
			} `json:"NSOpenGLContextAuxiliary"`
			NSScreenAuxiliaryOpaque struct {
				Fieldnames []interface{} `json:"fieldnames"`
				Special    bool          `json:"special"`
				Typestr    string        `json:"typestr"`
			} `json:"NSScreenAuxiliaryOpaque"`
			BCFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_BCFlags"`
			BCFlags2 struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_BCFlags2"`
			Brflags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_Brflags"`
			CFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_CFlags"`
			MFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_MFlags"`
			NSProgressIndicatorThreadInfo struct {
				Fieldnames []interface{} `json:"fieldnames"`
				Special    bool          `json:"special"`
				Typestr    string        `json:"typestr"`
			} `json:"_NSProgressIndicatorThreadInfo"`
			NSThreadPrivate struct {
				Fieldnames []interface{} `json:"fieldnames"`
				Special    bool          `json:"special"`
				Typestr    string        `json:"typestr"`
			} `json:"_NSThreadPrivate"`
			OVFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_OVFlags"`
			SFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_SFlags"`
			SPFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_SPFlags"`
			TvFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_TvFlags"`
			VFlags struct {
				Fieldnames []string `json:"fieldnames"`
				Special    bool     `json:"special"`
				Typestr    string   `json:"typestr"`
			} `json:"_VFlags"`
		} `json:"structs"`
	} `json:"definitions"`
	Framework string   `json:"framework"`
	Headers   []string `json:"headers"`
	Release   string   `json:"release"`
	Sdk       string   `json:"sdk"`
}
