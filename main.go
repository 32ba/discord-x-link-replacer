package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	linkReplaceRegex = regexp.MustCompile(`https:\/\/(?:twitter|x)\.com\/[A-Za-z0-9_]+\/status\/\d+`)
	ignoreSpoilerRegex = regexp.MustCompile(`\|\|([^|]+)\|\|`)
	config = Config{}
)

type Config struct {
	token string
	twitterProxyURL string
}


func main() {
	// load config
	config.token = os.Getenv("DISCORD_TOKEN")
	config.twitterProxyURL = os.Getenv("TWITTER_PROXY_URL")
	if config.token == "" {
		fmt.Println("DISCORD_TOKEN is not set")
		return
	}
	if config.twitterProxyURL == "" {
		config.twitterProxyURL = "https://vxtwitter.com"
	}

	dg, err := discordgo.New("Bot " + config.token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(replaceXLink)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
	dg.Close()
}

func replaceXLink(s *discordgo.Session, m *discordgo.MessageCreate) {
	withOutSpoilerContent := ignoreSpoilerRegex.ReplaceAllString(m.Content, "")
	matches := linkReplaceRegex.FindAllStringSubmatch(withOutSpoilerContent, -1)
	replacedLinks := []string{}
	for _, match := range matches {
		twUserName := strings.Split(match[0], "/")[3]
		twStatus := strings.Split(match[0], "/")[5]
		replacedLinks = append(replacedLinks, fmt.Sprintf("%s/%s/status/%s", config.twitterProxyURL, twUserName, twStatus))
	}
	if len(replacedLinks) == 0 {
		return
	}
	s.ChannelMessageSendReply(m.ChannelID, strings.Join(replacedLinks, "\n"), m.Reference())
}