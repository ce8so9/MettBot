package plugins

import (
	"ircclient"
	"http"
	"fmt"
	"xml"
)

type root struct {
	Events Events
}

type Events struct {
	Event []Event
}

type Event struct {
	Id         int `xml:"attr"`
	Submission Submission
	Judging    Judging
}

type Submission struct {
	Id       int `xml:"attr"`
	Team     string
	Problem  string
	Language string
}

type Judging struct {
	Id       int    `xml:"attr"`
	Submitid int    `xml:"attr"`
	Result   string `xml:"chardata"`
}

type HalloWeltPlugin struct {
	ic *ircclient.IRCClient
}

func (q *HalloWeltPlugin) Register(cl *ircclient.IRCClient) {
	q.ic = cl
	var client http.Client
	//response, _ := client.Get("https://bot:hallowelt@icpc.informatik.uni-erlangen.de/domjudge/plugin/event.php")
	// Um den DOMJudge nicht uebermaessig in der Entwicklungsphase zu pollen 
	// TODO: Konfigurierbar
	response, _ := client.Get("http://d-paulus.de/tmp.xml")
	fmt.Println(response.StatusCode)
	var res root
	xml.Unmarshal(response.Body, &res)
	response.Body.Close()
	fmt.Println(len(res.Events.Event))
}

func (q *HalloWeltPlugin) String() string {
	return "hallowelt"
}

func (q *HalloWeltPlugin) Info() string {
	return "DomJudge live ticker"
}

func (q *HalloWeltPlugin) Usage(cmd string) string {
	return "This plugin provides no commands"
}

func (q *HalloWeltPlugin) ProcessLine(msg *ircclient.IRCMessage) {
}

func (q *HalloWeltPlugin) ProcessCommand(cmd *ircclient.IRCCommand) {
}

func (q *HalloWeltPlugin) Unregister() {
}
