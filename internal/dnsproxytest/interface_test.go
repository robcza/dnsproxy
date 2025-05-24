package dnsproxytest_test

import (
	"github.com/robcza/dnsproxy/internal/dnsmsg"
	"github.com/robcza/dnsproxy/internal/dnsproxytest"
	"github.com/robcza/dnsproxy/upstream"
)

// type checks
var (
	_ upstream.Upstream         = (*dnsproxytest.FakeUpstream)(nil)
	_ dnsmsg.MessageConstructor = (*dnsproxytest.TestMessageConstructor)(nil)
)
