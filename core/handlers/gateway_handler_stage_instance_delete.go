package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

// gatewayHandlerStageInstanceDelete handles discord.GatewayEventTypeStageInstanceDelete
type gatewayHandlerStageInstanceDelete struct{}

// EventType returns the discord.GatewayEventType
func (h *gatewayHandlerStageInstanceDelete) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeStageInstanceDelete
}

// New constructs a new payload receiver for the raw gateway event
func (h *gatewayHandlerStageInstanceDelete) New() any {
	return &discord.StageInstance{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *gatewayHandlerStageInstanceDelete) HandleGatewayEvent(bot core.Bot, sequenceNumber discord.GatewaySequence, v any) {
	stageInstance := *v.(*discord.StageInstance)

	bot.Caches().StageInstances().Remove(stageInstance.GuildID, stageInstance.ID)

	bot.EventManager().Dispatch(&events.StageInstanceDeleteEvent{
		GenericStageInstanceEvent: &events.GenericStageInstanceEvent{
			GenericEvent:    events.NewGenericEvent(bot, sequenceNumber),
			StageInstanceID: stageInstance.ID,
			StageInstance:   stageInstance,
		},
	})
}
