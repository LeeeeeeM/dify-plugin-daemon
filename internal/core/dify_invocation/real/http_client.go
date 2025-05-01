package real

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/langgenius/dify-plugin-daemon/internal/core/dify_invocation"
)

func NewDifyInvocationDaemon(base string, calling_key string, write_timeout int, read_timeout int) (dify_invocation.BackwardsInvocation, error) {
	var err error
	invocation := &RealBackwardsInvocation{}
	baseurl, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 120 * time.Second,
			}).Dial,
			IdleConnTimeout: 120 * time.Second,
		},
	}

	invocation.difyInnerApiBaseurl = baseurl
	invocation.client = client
	invocation.difyInnerApiKey = calling_key
	invocation.writeTimeout = write_timeout
	invocation.readTimeout = read_timeout

	return invocation, nil
}
