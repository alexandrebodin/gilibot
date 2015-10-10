package listeners

import (
	"net/http"
	"strings"

	"github.com/alexandrebodin/gilibot"
)

// JenkinsListener defines parameters and handlers to send commands to jenkins
type JenkinsListener struct {
	BaseURI   string
	uriSuffix string
	username  string
	apiToken  string
}

var jenkinsHelp = `Usage:

	jenkins command[,] [arguments]

The commands are:

    build       start a jenkins build
    b           alias for build

Use "jenkins help [command]" for more information about a command.
`

// NewJenkinsListener returns a new configured JenkinsListener
func NewJenkinsListener(baseURI string, username string, apiToken string) *JenkinsListener {
	return &JenkinsListener{
		BaseURI:   baseURI,
		username:  username,
		apiToken:  apiToken,
		uriSuffix: "/api/json",
	}
}

// GetHandlers Returns an array of ListenerHandler to add to the bot
func (jenkins *JenkinsListener) GetHandlers() []*gilibot.ListenerHandler {

	return []*gilibot.ListenerHandler{
		{
			Regex: `j(?:enkins)? help\s*(.*)?`,
			HandlerFunc: func(c *gilibot.Context) {
				c.Reply(jenkinsHelp)
			},
		},
		{
			Regex: `j(?:enkins)? build ([\w\.\-_ ]+),\s*(.+)?`,
			HandlerFunc: func(c *gilibot.Context) {

				//get parameters if any
				jobName := c.Matches[1]
				buildParameters := strings.TrimSpace(c.Matches[2])
				url := jenkins.BaseURI + "/job/" + jobName + "/buildWithParameters" + jenkins.uriSuffix + "?" + buildParameters

				req, err := http.NewRequest("POST", url, nil)
				if err != nil {
					c.Reply("Deployment error")
					return
				}

				req.SetBasicAuth(jenkins.username, jenkins.apiToken)
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					c.Reply("Deployment error")
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == 201 {
					c.Reply("Deploy launched")
					return
				}

				c.Reply("Deployment error")
			},
		},
	}
}
