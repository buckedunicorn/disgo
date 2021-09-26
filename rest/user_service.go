package rest

import (
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/rest/route"
)

var _ UserService = (*userServiceImpl)(nil)

func NewUserService(restClient Client) UserService {
	return &userServiceImpl{restClient: restClient}
}

type UserService interface {
	Service
	GetUser(userID discord.Snowflake, opts ...RequestOpt) (*discord.User, Error)
	GetSelfUser(opts ...RequestOpt) (*discord.OAuth2User, Error)
	UpdateSelfUser(selfUserUpdate discord.SelfUserUpdate, opts ...RequestOpt) (*discord.OAuth2User, Error)
	GetGuilds(before int, after int, limit int, opts ...RequestOpt) ([]discord.OAuth2Guild, Error)
	LeaveGuild(guildID discord.Snowflake, opts ...RequestOpt) Error
	GetDMChannels(opts ...RequestOpt) ([]discord.Channel, Error)
	CreateDMChannel(userID discord.Snowflake, opts ...RequestOpt) (*discord.Channel, Error)
}

type userServiceImpl struct {
	restClient Client
}

func (s *userServiceImpl) RestClient() Client {
	return s.restClient
}

func (s *userServiceImpl) GetUser(userID discord.Snowflake, opts ...RequestOpt) (user *discord.User, rErr Error) {
	compiledRoute, err := route.GetUser.Compile(nil, userID)
	if err != nil {
		return nil, NewError(nil, err)
	}
	rErr = s.restClient.Do(compiledRoute, nil, &user, opts...)
	return
}

func (s *userServiceImpl) GetSelfUser(opts ...RequestOpt) (selfUser *discord.OAuth2User, rErr Error) {
	compiledRoute, err := route.GetCurrentUser.Compile(nil)
	if err != nil {
		return nil, NewError(nil, err)
	}
	rErr = s.restClient.Do(compiledRoute, nil, &selfUser, opts...)
	return
}

func (s *userServiceImpl) UpdateSelfUser(updateSelfUser discord.SelfUserUpdate, opts ...RequestOpt) (selfUser *discord.OAuth2User, rErr Error) {
	compiledRoute, err := route.UpdateSelfUser.Compile(nil)
	if err != nil {
		return nil, NewError(nil, err)
	}
	var user *discord.User
	rErr = s.restClient.Do(compiledRoute, updateSelfUser, &user, opts...)
	return
}

func (s *userServiceImpl) GetGuilds(before int, after int, limit int, opts ...RequestOpt) (guilds []discord.OAuth2Guild, rErr Error) {
	queryParams := route.QueryValues{}
	if before > 0 {
		queryParams["before"] = before
	}
	if after > 0 {
		queryParams["after"] = after
	}
	if limit > 0 {
		queryParams["limit"] = limit
	}
	compiledRoute, err := route.GetCurrentUserGuilds.Compile(queryParams)
	if err != nil {
		return nil, NewError(nil, NewError(nil, err))
	}

	rErr = s.restClient.Do(compiledRoute, nil, &guilds, opts...)
	return
}

func (s *userServiceImpl) LeaveGuild(guildID discord.Snowflake, opts ...RequestOpt) Error {
	compiledRoute, err := route.LeaveGuild.Compile(nil, guildID)
	if err != nil {
		return NewError(nil, err)
	}
	return s.restClient.Do(compiledRoute, nil, nil, opts...)
}

func (s *userServiceImpl) GetDMChannels(opts ...RequestOpt) (channels []discord.Channel, rErr Error) {
	compiledRoute, err := route.GetDMChannels.Compile(nil)
	if err != nil {
		return nil, NewError(nil, err)
	}

	rErr = s.restClient.Do(compiledRoute, nil, &channels, opts...)
	return
}

func (s *userServiceImpl) CreateDMChannel(userID discord.Snowflake, opts ...RequestOpt) (channel *discord.Channel, rErr Error) {
	compiledRoute, err := route.CreateDMChannel.Compile(nil)
	if err != nil {
		return nil, NewError(nil, err)
	}

	rErr = s.restClient.Do(compiledRoute, discord.DMChannelCreate{RecipientID: userID}, &channel, opts...)
	return
}