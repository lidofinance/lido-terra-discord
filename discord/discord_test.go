package discord

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
)

func TestParticipants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockSession(ctrl)
	m.EXPECT().GuildMembers("testGuildID", "", 100).Return(
		[]*discordgo.Member{
			{
				User: &discordgo.User{
					ID:  "1",
					Bot: false,
				},
			},
			{
				User: &discordgo.User{
					ID:  "2",
					Bot: false,
				},
			},
			{
				User: &discordgo.User{
					ID:  "3",
					Bot: true,
				},
			},
		}, nil)
	m.EXPECT().GuildMembers("testGuildID", "3", 100).Return(nil, nil)

	lb := NewLidoBot("testGuildID", "testChannelID", m)

	participants, err := lb.GetParticipants()
	if err != nil {
		t.Fatal(err)
	}

	expectedParticipants := 2
	if len(participants) != 2 {
		t.Fatalf("got participants %d, want %d\n", len(participants), expectedParticipants)
	}
}

func MessageMatcher(content string, pic Picture) eqMessageMatcher {
	return eqMessageMatcher{
		content: content,
		picture: pic,
	}
}

type eqMessageMatcher struct {
	content string
	picture Picture
}

func (mm eqMessageMatcher) Matches(x interface{}) bool {
	m, ok := x.(*discordgo.MessageSend)
	if !ok {
		return false
	}
	if mm.picture != nil {
		if len(m.Files) != 1 {
			return false
		}
	}
	return mm.content == m.Content
}

func (mm eqMessageMatcher) String() string {
	return fmt.Sprintln("message content - ", mm.content)
}

func TestSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPicture(ctrl)
	m := NewMockSession(ctrl)

	lb := NewLidoBot("testGuildID", "testChannelID", m)

	m.EXPECT().ChannelMessageSendComplex("testChannelID", MessageMatcher("test message", nil)).Times(1)
	m.EXPECT().ChannelMessageSendComplex("testChannelID", MessageMatcher("<@TestID>: test message", nil)).Times(1)
	m.EXPECT().ChannelMessageSendComplex("testChannelID", MessageMatcher("<@TestID>: test message with attachment", p)).Times(1)

	p.EXPECT().Name().Times(1)
	p.EXPECT().Body().Times(1)

	err := lb.SendMessage(nil, "test message", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = lb.SendMessage([]Participant{{UserName: "TestUser", UserID: "TestID"}}, "test message", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = lb.SendMessage([]Participant{{UserName: "TestUser", UserID: "TestID"}}, "test message with attachment", []Picture{p})
	if err != nil {
		t.Fatal(err)
	}
}
