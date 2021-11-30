package discord

import "github.com/bwmarrin/discordgo"

type Session interface {
	GuildMembers(guildID string, after string, limit int) (st []*discordgo.Member, err error)
	ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (st *discordgo.Message, err error)
}