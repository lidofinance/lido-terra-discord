package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Participant struct {
	UserName string
	UserID   string
}

func (p Participant) Mention() string {
	return "<@" + p.UserID + ">"
}

type Bot interface {
	GetParticipants() ([]Participant, error)
	SendMessage(participantsToTag []Participant, msg string, pictures []Picture) error
}

func NewDefaultLidoBot(token string, guildID, channelID string) (Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot session instance: %w", err)
	}

	lb := NewLidoBot(guildID, channelID, session)
	return lb, nil
}

func NewLidoBot(guildID, channelID string, session Session) Bot {
	lb := &LidoBot{
		session:   session,
		channelID: channelID,
		guildID:   guildID,
	}
	return lb
}

type LidoBot struct {
	session   Session
	channelID string
	guildID   string
}

func (lb LidoBot) getAllMembers() ([]*discordgo.Member, error) {
	members := make([]*discordgo.Member, 0)
	lastFetchedMemberID := ""
	for {
		m, err := lb.session.GuildMembers(lb.guildID, lastFetchedMemberID, 100)
		if err != nil {
			return nil, fmt.Errorf("failed to get members of the channel(ID=%s): %w", lb.channelID, err)
		}
		if len(m) == 0 {
			break
		}
		members = append(members, m...)
		lastFetchedMemberID = m[len(m)-1].User.ID
	}
	return members, nil
}

func (lb LidoBot) GetParticipants() ([]Participant, error) {
	members, err := lb.getAllMembers()
	if err != nil {
		return nil, fmt.Errorf("failed to get all members: %w", err)
	}

	participants := make([]Participant, 0, len(members))
	for _, m := range members {
		if m.User.Bot {
			// ignore bots
			continue
		}
		participants = append(participants, Participant{
			UserName: m.User.Username,
			UserID:   m.User.ID,
		})
	}
	return participants, nil
}

func (lb LidoBot) SendMessage(participantsToTag []Participant, msgContent string, pictures []Picture) error {
	var messageContent string
	if len(participantsToTag) > 0 {
		mentions := make([]string, 0, len(participantsToTag))
		for _, p := range participantsToTag {
			mentions = append(mentions, p.Mention())
		}
		messageContent = fmt.Sprintf("%s: ", strings.Join(mentions, " "))
	}
	messageContent += msgContent
	message := &discordgo.MessageSend{
		Content: messageContent,
	}

	for _, p := range pictures {
		message.Files = append(message.Files, &discordgo.File{
			Name:   p.Name(),
			Reader: p.Body(),
		})
	}

	_, err := lb.session.ChannelMessageSendComplex(lb.channelID, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
