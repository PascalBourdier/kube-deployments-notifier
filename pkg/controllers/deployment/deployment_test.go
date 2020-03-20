package deployment

import (
	"sync"
	"testing"
	"time"

	apps_v1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bpineau/kube-deployments-notifier/config"
	"github.com/bpineau/kube-deployments-notifier/pkg/notifiers/count"
)

var (
	obj1 = &apps_v1.Deployment{ObjectMeta: meta_v1.ObjectMeta{
		Name:      "test1",
		Labels:    config.Labels,
		Namespace: v1.NamespaceDefault},
	}
)

func TestDeployment(t *testing.T) {
	notif := new(count.Notifier)
	cont := new(Controller)
	cont.Init(config.FakeConfig(obj1), notif)

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	go cont.Start(&wg)
	defer func(cont *Controller) {
		go cont.Stop()
	}(cont)

	time.Sleep(config.FakeResyncInterval + config.FakeResyncInterval/2)
	if notif.Count() != 2 { // 1 initial + 1 after resync
		t.Errorf("Expected 2 notification, got %d", notif.Count())
	}
}
