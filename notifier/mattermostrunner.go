package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"github.com/eternnoir/gmrn/utils"
	"text/template"
)

type MatterMostPayload struct {
	Channel  *string `json:"channel,omitempty"`
	Username *string `json:"username,omitempty"`
	Text     *string `json:"text"`
}

func (mmp *MatterMostPayload) Serialize() ([]byte, error) {
	return json.Marshal(mmp)
}

type MMNotifyRunner struct {
	WebhookUrl   string
	Channel      string
	Username     string
	TextTemplate string
}

func (mmnr *MMNotifyRunner) Trigger(mr *apis.MergeRequest) error {
	log.Debugf("MM NotifyRunner Get New MR %#v.Start Trigger runner %#v ", mr, mmnr)
	payload := &MatterMostPayload{}
	payloadText, err := mmnr.convertTemplateToText(mr)
	if err != nil {
		return fmt.Errorf("MMNotifyRunner Convert template to payload [Fail]. [Runner] %#v. [MR] %#v, [ERROR]%s", mmnr, mr, err)
	}
	if mmnr.Channel != "" {
		payload.Channel = &mmnr.Channel
	}
	if mmnr.Username != "" {
		payload.Username = &mmnr.Username
	}
	payload.Text = &payloadText
	err = mmnr.FirePayload(payload)
	return nil
}

func (mmnr *MMNotifyRunner) convertTemplateToText(mr *apis.MergeRequest) (string, error) {
	tmpl, err := template.New("name").Parse(mmnr.TextTemplate)
	if err != nil {
		return "", err
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, *mr)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}

func (mmnr *MMNotifyRunner) FirePayload(payload *MatterMostPayload) error {
	jsonbody, err := payload.Serialize()
	if err != nil {
		return nil
	}
	jsonStr := string(jsonbody)
	log.Infof("MMNotifyRunner Start to fire payload to %s. Payload text %s", mmnr.WebhookUrl, string(jsonStr))
	_, err = utils.PostRequest(mmnr.WebhookUrl, jsonStr)
	if err != nil {
		log.Errorf("Send to targer: %#v payload %#v fail. Error: %s.", mmnr, payload, err)
		return err
	}
	return nil
}
