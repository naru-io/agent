package test

import (
	"github.com/coreos/go-etcd/etcd"
	"testing"
)

var etcdUrl = []string{"http://etcd_1:4001"}

func TestEtcdSetGet(t *testing.T) {

	c := etcd.NewClient(etcdUrl)
	if c == nil {
		t.Error(c)
	}

	//Set kv
	resp, err := c.Set("test1", "{ \"test\" : true }", 0)
	if err != nil || resp == nil {
		t.Error(err)
	}
	t.Logf("Action:%v EtcdIndex:%v Node.Key:%v Node.Value:%v Node.Dir:%v Node.TTL:%v", resp.Action, resp.EtcdIndex, resp.Node.Key, resp.Node.Value, resp.Node.Dir, resp.Node.TTL)

	resp, err = c.Get("test1", true, false)
	if err != nil {
		t.Error(err)
	}
	t.Logf("EtcdIndex:%v Node.Value:%v", resp.EtcdIndex, resp.Node.Value)
}

func TestEtcdNodeCheck(t *testing.T) {
	c := etcd.NewClient(etcdUrl)
	if c == nil {
		t.Error(c)
	}

	resp, err := c.Get("test333", false, false)
	if err != nil {
		if etcdErr, ok := err.(*etcd.EtcdError); ok {
			t.Logf("%#v[%T]", etcdErr, etcdErr)
		} else {
			t.Error(resp, err)

		}
	}
}
