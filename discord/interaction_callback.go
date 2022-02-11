package discord

// InteractionCallbackType indicates the type of slash command response, whether it's responding immediately or deferring to edit your response later
type InteractionCallbackType int

// Constants for the InteractionCallbackType(s)
const (
	InteractionCallbackTypePong InteractionCallbackType = iota + 1
	_
	_
	InteractionCallbackTypeChannelMessageWithSource
	InteractionCallbackTypeDeferredChannelMessageWithSource
	InteractionCallbackTypeDeferredUpdateMessage
	InteractionCallbackTypeUpdateMessage
	InteractionCallbackTypeAutocompleteResult
	InteractionCallbackTypeModal
)

// InteractionResponse is how you answer interactions. If an answer is not sent within 3 seconds of receiving it, the interaction is failed, and you will be unable to respond to it.
type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data InteractionCallbackData `json:"data,omitempty"`
}

// ToBody returns the InteractionResponse ready for body
func (r InteractionResponse) ToBody() (interface{}, error) {
	if v, ok := r.Data.(InteractionResponseCreator); ok {
		return v.ToResponseBody(r)
	}
	return r, nil
}

type InteractionCallbackData interface {
	interactionCallbackData()
}

type InteractionResponseCreator interface {
	ToResponseBody(response InteractionResponse) (interface{}, error)
}

type AutocompleteResult struct {
	Choices []AutocompleteChoice `json:"choices"`
}

func (AutocompleteResult) interactionCallbackData() {}

type AutocompleteChoice interface {
	autoCompleteChoice()
}

type AutocompleteChoiceString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (AutocompleteChoiceString) autoCompleteChoice() {}

type AutocompleteChoiceInt struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (AutocompleteChoiceInt) autoCompleteChoice() {}

type AutocompleteChoiceFloat struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (AutocompleteChoiceFloat) autoCompleteChoice() {}
