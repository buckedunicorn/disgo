package handlers

import (
	"github.com/DisgoOrg/disgo/api"
	"github.com/DisgoOrg/disgo/api/events"
)

// MessageCreateHandler handles api.MessageCreateGatewayEvent
type MessageCreateHandler struct{}

// Event returns the raw gateway event Event
func (h MessageCreateHandler) Event() api.GatewayEventName {
	return api.GatewayEventMessageCreate
}

// New constructs a new payload receiver for the raw gateway event
func (h MessageCreateHandler) New() interface{} {
	return &api.Message{}
}

// Handle handles the specific raw gateway event
func (h MessageCreateHandler) Handle(disgo api.Disgo, eventManager api.EventManager, i interface{}) {
	message, ok := i.(*api.Message)
	if !ok {
		return
	}

	genericMessageEvent := events.GenericMessageEvent{
		GenericEvent:     api.NewEvent(disgo),
		MessageChannelID: message.ChannelID,
		MessageID:        message.ID,
	}
	eventManager.Dispatch(genericMessageEvent)

	genericGuildEvent := events.GenericGuildEvent{
		GenericEvent: api.NewEvent(disgo),
		GuildID:      *message.GuildID,
	}
	eventManager.Dispatch(genericGuildEvent)

	eventManager.Dispatch(events.MessageReceivedEvent{
		GenericMessageEvent: genericMessageEvent,
		Message:             *message,
	})

	if message.GuildID == nil {
		// dm channel
	} else {
		// text channel
		message.Disgo = disgo
		message.Author.Disgo = disgo
		eventManager.Dispatch(events.GuildMessageReceivedEvent{
			Message: *message,
			GenericGuildMessageEvent: events.GenericGuildMessageEvent{
				GenericMessageEvent: genericMessageEvent,
			},
		})
	}

}
