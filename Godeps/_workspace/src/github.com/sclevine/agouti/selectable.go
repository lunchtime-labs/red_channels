package agouti

import (
	"github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/sclevine/agouti/api"
	"github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/sclevine/agouti/internal/element"
	"github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/sclevine/agouti/internal/target"
)

type selectable struct {
	session   apiSession
	selectors target.Selectors
}

type apiSession interface {
	element.Client
	Delete() error
	GetActiveElement() (*api.Element, error)
	GetWindow() (*api.Window, error)
	GetWindows() ([]*api.Window, error)
	SetWindow(window *api.Window) error
	SetWindowByName(name string) error
	DeleteWindow() error
	GetScreenshot() ([]byte, error)
	GetCookies() ([]*api.Cookie, error)
	SetCookie(cookie *api.Cookie) error
	DeleteCookie(name string) error
	DeleteCookies() error
	GetURL() (string, error)
	SetURL(url string) error
	GetTitle() (string, error)
	GetSource() (string, error)
	MoveTo(element *api.Element, point api.Offset) error
	Frame(frame *api.Element) error
	FrameParent() error
	Execute(body string, arguments []interface{}, result interface{}) error
	Forward() error
	Back() error
	Refresh() error
	GetAlertText() (string, error)
	SetAlertText(text string) error
	AcceptAlert() error
	DismissAlert() error
	NewLogs(logType string) ([]api.Log, error)
	GetLogTypes() ([]string, error)
	DoubleClick() error
	Click(button api.Button) error
	ButtonDown(button api.Button) error
	ButtonUp(button api.Button) error
	TouchDown(x, y int) error
	TouchUp(x, y int) error
	TouchMove(x, y int) error
	TouchClick(element *api.Element) error
	TouchDoubleClick(element *api.Element) error
	TouchLongClick(element *api.Element) error
	TouchFlick(element *api.Element, offset api.Offset, speed api.Speed) error
	TouchScroll(element *api.Element, offset api.Offset) error
	DeleteLocalStorage() error
	DeleteSessionStorage() error
}

// Find finds exactly one element by CSS selector.
func (s *selectable) Find(selector string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.CSS, selector).Single())
}

// FindByXPath finds exactly one element by XPath selector.
func (s *selectable) FindByXPath(selector string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.XPath, selector).Single())
}

// FindByLink finds exactly one anchor element by its text content.
func (s *selectable) FindByLink(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Link, text).Single())
}

// FindByLabel finds exactly one element by associated label text.
func (s *selectable) FindByLabel(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Label, text).Single())
}

// FindByButton finds exactly one button element with the provided text.
// Supports <button>, <input type="button">, and <input type="submit">.
func (s *selectable) FindByButton(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Button, text).Single())
}

// FindByName finds exactly element with the provided name attribute.
func (s *selectable) FindByName(name string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Name, name).Single())
}

// First finds the first element by CSS selector.
func (s *selectable) First(selector string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.CSS, selector).At(0))
}

// FirstByXPath finds the first element by XPath selector.
func (s *selectable) FirstByXPath(selector string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.XPath, selector).At(0))
}

// FirstByLink finds the first anchor element by its text content.
func (s *selectable) FirstByLink(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Link, text).At(0))
}

// FirstByLabel finds the first element by associated label text.
func (s *selectable) FirstByLabel(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Label, text).At(0))
}

// FirstByButton finds the first button element with the provided text.
// Supports <button>, <input type="button">, and <input type="submit">.
func (s *selectable) FirstByButton(text string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Button, text).At(0))
}

// FirstByName finds the first element with the provided name attribute.
func (s *selectable) FirstByName(name string) *Selection {
	return newSelection(s.session, s.selectors.Append(target.Name, name).At(0))
}

// All finds zero or more elements by CSS selector.
func (s *selectable) All(selector string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.CSS, selector))
}

// AllByXPath finds zero or more elements by XPath selector.
func (s *selectable) AllByXPath(selector string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.XPath, selector))
}

// AllByLink finds zero or more anchor elements by their text content.
func (s *selectable) AllByLink(text string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.Link, text))
}

// AllByLabel finds zero or more elements by associated label text.
func (s *selectable) AllByLabel(text string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.Label, text))
}

// AllByButton finds zero or more button elements with the provided text.
// Supports <button>, <input type="button">, and <input type="submit">.
func (s *selectable) AllByButton(text string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.Button, text))
}

// AllByName finds zero or more elements with the provided name attribute.
func (s *selectable) AllByName(name string) *MultiSelection {
	return newMultiSelection(s.session, s.selectors.Append(target.Name, name))
}
