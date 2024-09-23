// Auto-generated structs for Workspace Add-Ons based on the JSON schema files
// In addition there are a helper method library to avoid unneccessary templating code when using the structs
// Since this is basically just putting values into structs, there are no error checking/handling. Client programs should instead do error checking on their input.
//
// Unfortunately the JSON schemas are not 100% suited for Go code, so there are also included some "fixes" for this in the library. If this is causing trouble, please check the source code to see what has been done
//
// Terminology
// This helper library uses Create for main components like Actions and Cards
// while Add is used for adding additional components to the main components,
// like adding a Widget to a Card, or adding a Button to a Widget
package WorkspaceAddOns

import (
	"encoding/json"
)

// Create a new NavigationAction in this RenderAction
func (ra *RenderAction) CreateAction() *NavigationAction {
	var action NavigationAction
	//	action.Link = OpenLink{}
	ra.Action = &action
	return ra.Action
}

// Create a new Card
func CreateCard(title string) *Card {
	var card Card
	var header CardHeader
	header.Title = title
	card.Header = &header
	return &card
}

// Add a new Navigation to this NavigationAction
func (na *NavigationAction) AddNavigation() *Navigation {
	navigation := new(Navigation)
	na.Navigations = append(na.Navigations, *navigation)
	return &na.Navigations[len(na.Navigations)-1]
}

// Add a Section to a Card
func (c *Card) AddSection(title string) *Section {
	var section Section
	section.Header = &title
	c.Sections = append(c.Sections, section)
	return &c.Sections[len(c.Sections)-1]
}

// Add a Widget to a Section
func (s *Section) AddWidget() *Widget {
	var widget Widget
	s.Widgets = append(s.Widgets, widget)
	return &s.Widgets[len(s.Widgets)-1]
}

// Add TextInput to a Widget
func (w *Widget) AddTextInput(name, label, value string) *TextInput {
	var textInput TextInput
	textInput.Name = name
	textInput.Label = &label
	textInput.Value = &value
	w.TextInput = &textInput
	return w.TextInput
}

// Add TextParagraph to a Widget
func (w *Widget) AddTextParagraph(text string) *TextParagraph {
	var textParagraph TextParagraph
	textParagraph.Text = text
	w.TextParagraph = &textParagraph
	return w.TextParagraph
}

// Add an Image to a Widget
func (w *Widget) AddImage(altText, url string) *Image {
	var image Image
	image.AltText = &altText
	image.ImageUrl = url
	w.Image = &image
	return w.Image
}

// Add a ButtonList to a Widget
func (w *Widget) AddButtonList() *ButtonList {
	var buttonList ButtonList
	w.ButtonList = &buttonList
	return &buttonList
}

// Add a Button to a ButtonList
func (b *ButtonList) AddButton() *Button {
	var button Button
	b.Buttons = append(b.Buttons, button)
	return &b.Buttons[len(b.Buttons)-1]
}

// Add OnClick to a Button
func (b *Button) AddOnClick(label string) *OnClick {
	var onClick OnClick
	b.Text = label
	b.OnClick = &onClick
	return &onClick
}

// Add OpenLink to an OnClick
func (o *OnClick) AddOpenLink(url string) *OpenLink {
	var openLink OpenLink
	openLink.Url = url
	o.OpenLink = &openLink
	return &openLink
}

// Not generated automatically since Action is used in several places with different meaning
type FormAction struct {
	Function string `json:"function,omitempty"`
}

// Wrapper since the payload must start with renderAction when it is a response to a form submit
type RenderActionWrapper struct {
	RenderAction *RenderAction `json:"renderActions,omitempty"`
}

// Add a RenderAction to the RenderActionWrapper (for response to a form submit)
func (ra *RenderActionWrapper) AddRenderAction() *RenderAction {
	var renderAction RenderAction
	ra.RenderAction = &renderAction
	return &renderAction
}

// Add Notification to a NavigationAction
func (a *NavigationAction) AddNotification(text string) {
	var notification Notification
	notification.Text = &text
	a.Notification = &notification
}

// Add a new Card to a Navigation
func (n *Navigation) AddCard() *Card {
	var card Card
	n.PushCard = &card
	return &card
}

// Add a Header to a Card
func (c *Card) AddHeader(imageType string) {
	var h CardHeader
	h.ImageType = (*CardHeaderImageType)(&imageType)
	c.Header = &h
}

// Add Fixed Footer to a Card
func (c *Card) AddFixedFooter() *CardFixedFooterWrapper {
	var f CardFixedFooterWrapper
	c.FixedFooter = &f
	return &f
}

// Add a button as Primary button in the Fixed Footer
func (f *CardFixedFooterWrapper) AddPrimaryButton(text, url string) {
	var button Button
	button.Text = text
	var openLink OpenLink
	openLink.Url = url
	var onClick OnClick
	onClick.OpenLink = &openLink
	button.OnClick = &onClick
	f.PrimaryButton = button
}

// Add a button as Secondary button in the Fixed Footer
func (f *CardFixedFooterWrapper) AddSecondaryButton(text, url string) {
	var button Button
	button.Text = text
	var openLink OpenLink
	openLink.Url = url
	var onClick OnClick
	onClick.OpenLink = &openLink
	button.OnClick = &onClick
	f.SecondaryButton = button
}

// Add a Submit Button to a ButtonList
func (b *ButtonList) AddSubmitButton(text, url string) {
	var button Button
	button.Text = text
	formAction := new(Action)
	formAction.Function = url
	var onClick OnClick
	onClick.Action = formAction
	button.OnClick = &onClick
	b.Buttons = append(b.Buttons, button)
}

// A custom JSON marshaller is neccessary to make sure that an empty url is not included (url = ”). If url had been tagged with omitempty in OpenLink or Link in NavigationAction had been a pointer, this wouldn't have been neccessary
func (na NavigationAction) MarshalJSON() ([]byte, error) {
	type Alias NavigationAction // Alias for å unngå uendelig rekursiv kall
	if na.Link.Url == "" {
		return json.Marshal(struct {
			Navigations  []Navigation  `json:"navigations,omitempty"`
			Notification *Notification `json:"notification,omitempty"`
		}{
			Navigations:  na.Navigations,
			Notification: na.Notification,
		})
	}
	return json.Marshal(Alias(na))
}

type CardFixedFooterWrapper struct {
	PrimaryButton   Button `json:"primaryButton,omitempty"`
	SecondaryButton Button `json:"secondaryButton,omitempty"`
}
