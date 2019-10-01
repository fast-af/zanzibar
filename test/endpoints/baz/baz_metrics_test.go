// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package baz

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/m3"
	bazClient "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	zanzibar "github.com/uber/zanzibar/runtime"
	testGateway "github.com/uber/zanzibar/test/lib/test_gateway"
	"github.com/uber/zanzibar/test/lib/util"
)

func TestCallMetrics(t *testing.T) {
	testcallCounter := 0

	gateway, err := testGateway.CreateGateway(t, map[string]interface{}{
		"clients.baz.serviceName": "bazService",
	}, &testGateway.Options{
		CountMetrics:          true,
		KnownTChannelBackends: []string{"baz"},
		TestBinary:            util.DefaultMainFile("example-gateway"),
		ConfigFiles:           util.DefaultConfigFiles("example-gateway"),
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	cg := gateway.(*testGateway.ChildProcessGateway)
	fakeCall := func(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Call_Args,
	) (map[string]string, error) {
		testcallCounter++

		var resHeaders map[string]string

		return resHeaders, nil
	}

	err = gateway.TChannelBackends()["baz"].Register(
		"baz", "call", "SimpleService::call",
		bazClient.NewSimpleServiceCallHandler(fakeCall),
	)
	assert.NoError(t, err)

	headers := map[string]string{}
	headers["x-token"] = "token"
	headers["x-uuid"] = "uuid"
	headers["regionname"] = "san_francisco"
	headers["device"] = "ios"
	headers["deviceversion"] = "carbon"

	numMetrics := 13
	cg.MetricsWaitGroup.Add(numMetrics)

	_, err = gateway.MakeRequest(
		"POST",
		"/baz/call",
		headers,
		bytes.NewReader([]byte(`{"arg":{"b1":true,"i3":42,"s2":"hello"}}`)),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	cg.MetricsWaitGroup.Wait()
	metrics := cg.M3Service.GetMetrics()
	cbKeys := make([]string, 0)
	for key := range metrics {
		if strings.Contains(key, "circuitbreaker") {
			cbKeys = append(cbKeys, key)
		}
	}
	assert.Equal(t, 6, len(cbKeys)) // number off because of the histogram
	// we don't care about jaeger emitted metrics
	for key := range metrics {
		if strings.HasPrefix(key, "jaeger") {
			delete(metrics, key)
		}
	}
	assert.Equal(t, numMetrics+4, len(metrics)) // magic number here because there are histogram entries

	endpointTags := map[string]string{
		"env":           "test",
		"service":       "test-gateway",
		"endpointid":    "baz",
		"handlerid":     "call",
		"regionname":    "san_francisco",
		"device":        "ios",
		"deviceversion": "carbon",
		"dc":            "unknown",
		"protocol":      "HTTP",
	}
	statusTags := map[string]string{
		"status": "204",
	}
	for k, v := range endpointTags {
		statusTags[k] = v
	}
	histogramTags := map[string]string{
		m3.DefaultHistogramBucketName:   "-infinity-10ms", // TODO(argouber): Remove the ugly hardcoding
		m3.DefaultHistogramBucketIDName: "0000",
	}
	for k, v := range statusTags {
		histogramTags[k] = v
	}
	key := tally.KeyForPrefixedStringMap("endpoint.request", endpointTags)
	assert.Contains(t, metrics, key, "expected metric: %s", key)
	recvdMetric := metrics[key]
	value := *recvdMetric.MetricValue.Count.I64Value
	assert.Equal(t, int64(1), value, "expected counter to be 1")

	key = tally.KeyForPrefixedStringMap("endpoint.latency", statusTags)
	assert.Contains(t, metrics, key, "expected metric: %s", key)
	latencyMetric := metrics[key]
	value = *latencyMetric.MetricValue.Timer.I64Value
	assert.True(t, value > 1000, "expected timer to be >1000 nano seconds")
	assert.True(t, value < 10*1000*1000, "expected timer to be < 10 milli seconds")

	key = tally.KeyForPrefixedStringMap("endpoint.latency-hist", histogramTags)
	assert.Contains(t, metrics, key, "expected metric: %s", key)
	latencyHistMetric := metrics[key]
	assert.Equal(t, int64(1), *latencyHistMetric.MetricValue.Count.I64Value)

	statusMetric := metrics[tally.KeyForPrefixedStringMap(
		"endpoint.status", statusTags,
	)]
	value = *statusMetric.MetricValue.Count.I64Value
	assert.Equal(t, int64(1), value, "expected counter to be 1")

	tchannelNames := []string{
		"outbound.calls.per-attempt.latency",
	}
	tchannelTags := map[string]string{
		"env":             "test",
		"app":             "test-gateway",
		"service":         "test-gateway",
		"target-service":  "bazService",
		"target-endpoint": "SimpleService__call",
		"host":            zanzibar.GetHostname(),
		"dc":              "unknown",
	}
	for _, name := range tchannelNames {
		key := tally.KeyForPrefixedStringMap(name, tchannelTags)
		assert.Contains(t, metrics, key, "expected metric: %s", key)
	}

	perAttemptOutboundLatency := metrics[tally.KeyForPrefixedStringMap(
		"outbound.calls.per-attempt.latency",
		tchannelTags,
	)]
	value = *perAttemptOutboundLatency.MetricValue.Timer.I64Value
	assert.True(t, value > 1000, "expected timer to be >1000 nano seconds")
	assert.True(t, value < 1000*1000*1000, "expected timer to be <1 second")

	clientNames := []string{
		"client.request",
		"client.success",
	}
	clientTags := map[string]string{
		"env":            "test",
		"service":        "test-gateway",
		"clientid":       "baz",
		"clientmethod":   "call",
		"targetservice":  "bazService",
		"targetendpoint": "SimpleService__call",
		"dc":             "unknown",
		"endpointid":     "baz",
		"handlerid":      "call",
		"regionname":     "san_francisco",
		"device":         "ios",
		"deviceversion":  "carbon",
		"protocol":       "HTTP",
	}
	for _, name := range clientNames {
		key := tally.KeyForPrefixedStringMap(name, clientTags)
		assert.Contains(t, metrics, key, "expected metric: %s", key)
		value := metrics[key]
		assert.Equal(t, int64(1), *value.MetricValue.Count.I64Value, "expected counter to be 1")
	}

	key = tally.KeyForPrefixedStringMap("client.latency", clientTags)
	assert.Contains(t, metrics, key, "expected metric: %s", key)
	value = *metrics[key].MetricValue.Timer.I64Value
	assert.True(t, value > 1000, "expected timer to be >1000 nano seconds")
	assert.True(t, value < 10*1000*1000, "expected timer to be < 10 milli seconds")

	cHistogramTags := map[string]string{
		m3.DefaultHistogramBucketName:   "-infinity-10ms",
		m3.DefaultHistogramBucketIDName: "0000",
	}
	for k, v := range clientTags {
		cHistogramTags[k] = v
	}
	key = tally.KeyForPrefixedStringMap("client.latency-hist", cHistogramTags)
	assert.Contains(t, metrics, key, "expected metric: %s", key)
	assert.Equal(t, int64(1), *metrics[key].MetricValue.Count.I64Value)
}
