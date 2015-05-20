package listeners

import (
	"github.com/alexandrebodin/gilibot"
	"net/http"
)

type JenkinsListener struct {
	BaseUri   string
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

func NewJenkinsListener(baseUri string, username string, apiToken string) *JenkinsListener {
	return &JenkinsListener{
		BaseUri:   baseUri,
		username:  username,
		apiToken:  apiToken,
		uriSuffix: "/api/json",
	}
}

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
				buildParameters := c.Matches[2]
				url := jenkins.BaseUri + "/job/" + jobName + "/buildWithParameters" + jenkins.uriSuffix + "?" + buildParameters

				req, err := http.NewRequest("POST", url, nil)
				if err != nil {
					c.Reply("Deploy error")
					return
				}

				req.SetBasicAuth(jenkins.username, jenkins.apiToken)
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					c.Reply("Deploy error")
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == 201 {
					c.Reply("Deploy launched")
					return
				}
			},
		},
	}
}
