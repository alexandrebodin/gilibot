package listeners

import (
	"github.com/alexandrebodin/gilibot"
	"net/http"
	"regexp"
)

type JenkinsListener struct {
	BaseUri   string
	uriSuffix string
	username  string
	apiToken  string
}

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
			Regex: "jenkins build ([a-zA-Z.]*) .*",
			HandlerFunc: func(c *gilibot.Context) {

				//get parameters if any
				jobName := c.Matches[1]
				buildParameters := regexp.MustCompile("([a-zA-Z.]*=[a-zA-Z.]*)").FindAllString(c.Matches[0], -1)

				url := jenkins.BaseUri + "/job/" + jobName + "/buildWithParameters" + jenkins.uriSuffix
				if len(buildParameters) > 0 {
					url += "?"
				}

				for _, p := range buildParameters {
					url += p + "&"
				}

				req, err := http.NewRequest("POST", url, nil)
				if err != nil {
					c.Reply([]string{"Deploy error"})
					return
				}

				req.SetBasicAuth(jenkins.username, jenkins.apiToken)
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					c.Reply([]string{"Deploy error"})
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == 201 {
					c.Reply([]string{"Deploy launched"})
					return
				}
			},
		},
		{
			Regex: "merci (.*)",
			HandlerFunc: func(c *gilibot.Context) {
				c.Reply([]string{c.Matches[1]})
			},
		},
	}
}
