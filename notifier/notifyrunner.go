package notifier

import (
	"github.com/eternnoir/gmrn/apis"
)

type NotifyRunner interface {
	Trigger(mr *apis.MergeRequest) error
}
