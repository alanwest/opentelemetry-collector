// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "internal/data_generator/main.go". DO NOT EDIT.
// To regenerate this file run "go run internal/data_generator/main.go".

package data

import (
	"testing"

	otlpmetrics "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
	"github.com/stretchr/testify/assert"
)

func TestResourceMetricsSlice(t *testing.T) {
	es := NewResourceMetricsSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newResourceMetricsSlice(&[]*otlpmetrics.ResourceMetrics{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewResourceMetrics()
	emptyVal.InitEmpty()
	testVal := generateTestResourceMetrics()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestResourceMetrics(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.ResourceMetrics]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.ResourceMetrics]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.ResourceMetrics]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.ResourceMetrics]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewResourceMetricsSlice(), es)
}

func TestResourceMetrics(t *testing.T) {
	ms := NewResourceMetrics()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, true, ms.Resource().IsNil())
	ms.Resource().InitEmpty()
	assert.EqualValues(t, false, ms.Resource().IsNil())
	fillTestResource(ms.Resource())
	assert.EqualValues(t, generateTestResource(), ms.Resource())

	assert.EqualValues(t, NewInstrumentationLibraryMetricsSlice(), ms.InstrumentationLibraryMetrics())
	fillTestInstrumentationLibraryMetricsSlice(ms.InstrumentationLibraryMetrics())
	testValInstrumentationLibraryMetrics := generateTestInstrumentationLibraryMetricsSlice()
	assert.EqualValues(t, testValInstrumentationLibraryMetrics, ms.InstrumentationLibraryMetrics())

	assert.EqualValues(t, generateTestResourceMetrics(), ms)
}

func TestInstrumentationLibraryMetricsSlice(t *testing.T) {
	es := NewInstrumentationLibraryMetricsSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newInstrumentationLibraryMetricsSlice(&[]*otlpmetrics.InstrumentationLibraryMetrics{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewInstrumentationLibraryMetrics()
	emptyVal.InitEmpty()
	testVal := generateTestInstrumentationLibraryMetrics()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestInstrumentationLibraryMetrics(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.InstrumentationLibraryMetrics]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.InstrumentationLibraryMetrics]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.InstrumentationLibraryMetrics]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.InstrumentationLibraryMetrics]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewInstrumentationLibraryMetricsSlice(), es)
}

func TestInstrumentationLibraryMetrics(t *testing.T) {
	ms := NewInstrumentationLibraryMetrics()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, true, ms.InstrumentationLibrary().IsNil())
	ms.InstrumentationLibrary().InitEmpty()
	assert.EqualValues(t, false, ms.InstrumentationLibrary().IsNil())
	fillTestInstrumentationLibrary(ms.InstrumentationLibrary())
	assert.EqualValues(t, generateTestInstrumentationLibrary(), ms.InstrumentationLibrary())

	assert.EqualValues(t, NewMetricSlice(), ms.Metrics())
	fillTestMetricSlice(ms.Metrics())
	testValMetrics := generateTestMetricSlice()
	assert.EqualValues(t, testValMetrics, ms.Metrics())

	assert.EqualValues(t, generateTestInstrumentationLibraryMetrics(), ms)
}

func TestMetricSlice(t *testing.T) {
	es := NewMetricSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newMetricSlice(&[]*otlpmetrics.Metric{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewMetric()
	emptyVal.InitEmpty()
	testVal := generateTestMetric()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestMetric(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.Metric]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.Metric]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.Metric]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.Metric]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewMetricSlice(), es)
}

func TestMetric(t *testing.T) {
	ms := NewMetric()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, true, ms.MetricDescriptor().IsNil())
	ms.MetricDescriptor().InitEmpty()
	assert.EqualValues(t, false, ms.MetricDescriptor().IsNil())
	fillTestMetricDescriptor(ms.MetricDescriptor())
	assert.EqualValues(t, generateTestMetricDescriptor(), ms.MetricDescriptor())

	assert.EqualValues(t, NewInt64DataPointSlice(), ms.Int64DataPoints())
	fillTestInt64DataPointSlice(ms.Int64DataPoints())
	testValInt64DataPoints := generateTestInt64DataPointSlice()
	assert.EqualValues(t, testValInt64DataPoints, ms.Int64DataPoints())

	assert.EqualValues(t, NewDoubleDataPointSlice(), ms.DoubleDataPoints())
	fillTestDoubleDataPointSlice(ms.DoubleDataPoints())
	testValDoubleDataPoints := generateTestDoubleDataPointSlice()
	assert.EqualValues(t, testValDoubleDataPoints, ms.DoubleDataPoints())

	assert.EqualValues(t, NewHistogramDataPointSlice(), ms.HistogramDataPoints())
	fillTestHistogramDataPointSlice(ms.HistogramDataPoints())
	testValHistogramDataPoints := generateTestHistogramDataPointSlice()
	assert.EqualValues(t, testValHistogramDataPoints, ms.HistogramDataPoints())

	assert.EqualValues(t, NewSummaryDataPointSlice(), ms.SummaryDataPoints())
	fillTestSummaryDataPointSlice(ms.SummaryDataPoints())
	testValSummaryDataPoints := generateTestSummaryDataPointSlice()
	assert.EqualValues(t, testValSummaryDataPoints, ms.SummaryDataPoints())

	assert.EqualValues(t, generateTestMetric(), ms)
}

func TestMetricDescriptor(t *testing.T) {
	ms := NewMetricDescriptor()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())

	assert.EqualValues(t, "", ms.Description())
	testValDescription := "test_description"
	ms.SetDescription(testValDescription)
	assert.EqualValues(t, testValDescription, ms.Description())

	assert.EqualValues(t, "", ms.Unit())
	testValUnit := "1"
	ms.SetUnit(testValUnit)
	assert.EqualValues(t, testValUnit, ms.Unit())

	assert.EqualValues(t, MetricTypeUnspecified, ms.Type())
	testValType := MetricTypeGaugeInt64
	ms.SetType(testValType)
	assert.EqualValues(t, testValType, ms.Type())

	assert.EqualValues(t, NewStringMap(), ms.LabelsMap())
	fillTestStringMap(ms.LabelsMap())
	testValLabelsMap := generateTestStringMap()
	assert.EqualValues(t, testValLabelsMap, ms.LabelsMap())

	assert.EqualValues(t, generateTestMetricDescriptor(), ms)
}

func TestInt64DataPointSlice(t *testing.T) {
	es := NewInt64DataPointSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newInt64DataPointSlice(&[]*otlpmetrics.Int64DataPoint{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewInt64DataPoint()
	emptyVal.InitEmpty()
	testVal := generateTestInt64DataPoint()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestInt64DataPoint(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.Int64DataPoint]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.Int64DataPoint]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.Int64DataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.Int64DataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewInt64DataPointSlice(), es)
}

func TestInt64DataPoint(t *testing.T) {
	ms := NewInt64DataPoint()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, NewStringMap(), ms.LabelsMap())
	fillTestStringMap(ms.LabelsMap())
	testValLabelsMap := generateTestStringMap()
	assert.EqualValues(t, testValLabelsMap, ms.LabelsMap())

	assert.EqualValues(t, TimestampUnixNano(0), ms.StartTime())
	testValStartTime := TimestampUnixNano(1234567890)
	ms.SetStartTime(testValStartTime)
	assert.EqualValues(t, testValStartTime, ms.StartTime())

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, int64(0), ms.Value())
	testValValue := int64(-17)
	ms.SetValue(testValValue)
	assert.EqualValues(t, testValValue, ms.Value())

	assert.EqualValues(t, generateTestInt64DataPoint(), ms)
}

func TestDoubleDataPointSlice(t *testing.T) {
	es := NewDoubleDataPointSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newDoubleDataPointSlice(&[]*otlpmetrics.DoubleDataPoint{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewDoubleDataPoint()
	emptyVal.InitEmpty()
	testVal := generateTestDoubleDataPoint()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestDoubleDataPoint(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.DoubleDataPoint]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.DoubleDataPoint]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.DoubleDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.DoubleDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewDoubleDataPointSlice(), es)
}

func TestDoubleDataPoint(t *testing.T) {
	ms := NewDoubleDataPoint()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, NewStringMap(), ms.LabelsMap())
	fillTestStringMap(ms.LabelsMap())
	testValLabelsMap := generateTestStringMap()
	assert.EqualValues(t, testValLabelsMap, ms.LabelsMap())

	assert.EqualValues(t, TimestampUnixNano(0), ms.StartTime())
	testValStartTime := TimestampUnixNano(1234567890)
	ms.SetStartTime(testValStartTime)
	assert.EqualValues(t, testValStartTime, ms.StartTime())

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, float64(0.0), ms.Value())
	testValValue := float64(17.13)
	ms.SetValue(testValValue)
	assert.EqualValues(t, testValValue, ms.Value())

	assert.EqualValues(t, generateTestDoubleDataPoint(), ms)
}

func TestHistogramDataPointSlice(t *testing.T) {
	es := NewHistogramDataPointSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newHistogramDataPointSlice(&[]*otlpmetrics.HistogramDataPoint{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewHistogramDataPoint()
	emptyVal.InitEmpty()
	testVal := generateTestHistogramDataPoint()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestHistogramDataPoint(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.HistogramDataPoint]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.HistogramDataPoint]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.HistogramDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.HistogramDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewHistogramDataPointSlice(), es)
}

func TestHistogramDataPoint(t *testing.T) {
	ms := NewHistogramDataPoint()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, NewStringMap(), ms.LabelsMap())
	fillTestStringMap(ms.LabelsMap())
	testValLabelsMap := generateTestStringMap()
	assert.EqualValues(t, testValLabelsMap, ms.LabelsMap())

	assert.EqualValues(t, TimestampUnixNano(0), ms.StartTime())
	testValStartTime := TimestampUnixNano(1234567890)
	ms.SetStartTime(testValStartTime)
	assert.EqualValues(t, testValStartTime, ms.StartTime())

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, uint64(0), ms.Count())
	testValCount := uint64(17)
	ms.SetCount(testValCount)
	assert.EqualValues(t, testValCount, ms.Count())

	assert.EqualValues(t, float64(0.0), ms.Sum())
	testValSum := float64(17.13)
	ms.SetSum(testValSum)
	assert.EqualValues(t, testValSum, ms.Sum())

	assert.EqualValues(t, NewHistogramBucketSlice(), ms.Buckets())
	fillTestHistogramBucketSlice(ms.Buckets())
	testValBuckets := generateTestHistogramBucketSlice()
	assert.EqualValues(t, testValBuckets, ms.Buckets())

	assert.EqualValues(t, []float64(nil), ms.ExplicitBounds())
	testValExplicitBounds := []float64{1, 2, 3}
	ms.SetExplicitBounds(testValExplicitBounds)
	assert.EqualValues(t, testValExplicitBounds, ms.ExplicitBounds())

	assert.EqualValues(t, generateTestHistogramDataPoint(), ms)
}

func TestHistogramBucketSlice(t *testing.T) {
	es := NewHistogramBucketSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newHistogramBucketSlice(&[]*otlpmetrics.HistogramDataPoint_Bucket{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewHistogramBucket()
	emptyVal.InitEmpty()
	testVal := generateTestHistogramBucket()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestHistogramBucket(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.HistogramDataPoint_Bucket]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.HistogramDataPoint_Bucket]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.HistogramDataPoint_Bucket]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.HistogramDataPoint_Bucket]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewHistogramBucketSlice(), es)
}

func TestHistogramBucket(t *testing.T) {
	ms := NewHistogramBucket()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, uint64(0), ms.Count())
	testValCount := uint64(17)
	ms.SetCount(testValCount)
	assert.EqualValues(t, testValCount, ms.Count())

	assert.EqualValues(t, true, ms.Exemplar().IsNil())
	ms.Exemplar().InitEmpty()
	assert.EqualValues(t, false, ms.Exemplar().IsNil())
	fillTestHistogramBucketExemplar(ms.Exemplar())
	assert.EqualValues(t, generateTestHistogramBucketExemplar(), ms.Exemplar())

	assert.EqualValues(t, generateTestHistogramBucket(), ms)
}

func TestHistogramBucketExemplar(t *testing.T) {
	ms := NewHistogramBucketExemplar()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, float64(0.0), ms.Value())
	testValValue := float64(17.13)
	ms.SetValue(testValValue)
	assert.EqualValues(t, testValValue, ms.Value())

	assert.EqualValues(t, NewStringMap(), ms.Attachments())
	fillTestStringMap(ms.Attachments())
	testValAttachments := generateTestStringMap()
	assert.EqualValues(t, testValAttachments, ms.Attachments())

	assert.EqualValues(t, generateTestHistogramBucketExemplar(), ms)
}

func TestSummaryDataPointSlice(t *testing.T) {
	es := NewSummaryDataPointSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newSummaryDataPointSlice(&[]*otlpmetrics.SummaryDataPoint{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewSummaryDataPoint()
	emptyVal.InitEmpty()
	testVal := generateTestSummaryDataPoint()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestSummaryDataPoint(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.SummaryDataPoint]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.SummaryDataPoint]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.SummaryDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.SummaryDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewSummaryDataPointSlice(), es)
}

func TestSummaryDataPoint(t *testing.T) {
	ms := NewSummaryDataPoint()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, NewStringMap(), ms.LabelsMap())
	fillTestStringMap(ms.LabelsMap())
	testValLabelsMap := generateTestStringMap()
	assert.EqualValues(t, testValLabelsMap, ms.LabelsMap())

	assert.EqualValues(t, TimestampUnixNano(0), ms.StartTime())
	testValStartTime := TimestampUnixNano(1234567890)
	ms.SetStartTime(testValStartTime)
	assert.EqualValues(t, testValStartTime, ms.StartTime())

	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())

	assert.EqualValues(t, uint64(0), ms.Count())
	testValCount := uint64(17)
	ms.SetCount(testValCount)
	assert.EqualValues(t, testValCount, ms.Count())

	assert.EqualValues(t, float64(0.0), ms.Sum())
	testValSum := float64(17.13)
	ms.SetSum(testValSum)
	assert.EqualValues(t, testValSum, ms.Sum())

	assert.EqualValues(t, NewSummaryValueAtPercentileSlice(), ms.ValueAtPercentiles())
	fillTestSummaryValueAtPercentileSlice(ms.ValueAtPercentiles())
	testValValueAtPercentiles := generateTestSummaryValueAtPercentileSlice()
	assert.EqualValues(t, testValValueAtPercentiles, ms.ValueAtPercentiles())

	assert.EqualValues(t, generateTestSummaryDataPoint(), ms)
}

func TestSummaryValueAtPercentileSlice(t *testing.T) {
	es := NewSummaryValueAtPercentileSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newSummaryValueAtPercentileSlice(&[]*otlpmetrics.SummaryDataPoint_ValueAtPercentile{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(13)
	emptyVal := NewSummaryValueAtPercentile()
	emptyVal.InitEmpty()
	testVal := generateTestSummaryValueAtPercentile()
	assert.EqualValues(t, 13, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestSummaryValueAtPercentile(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}

	// Test Resize less elements.
	const resizeSmallLen = 10
	expectedEs := make(map[*otlpmetrics.SummaryDataPoint_ValueAtPercentile]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.SummaryDataPoint_ValueAtPercentile]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 13
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.SummaryDataPoint_ValueAtPercentile]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpmetrics.SummaryDataPoint_ValueAtPercentile]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewSummaryValueAtPercentileSlice(), es)
}

func TestSummaryValueAtPercentile(t *testing.T) {
	ms := NewSummaryValueAtPercentile()
	assert.EqualValues(t, true, ms.IsNil())
	ms.InitEmpty()
	assert.EqualValues(t, false, ms.IsNil())

	assert.EqualValues(t, float64(0.0), ms.Percentile())
	testValPercentile := float64(0.90)
	ms.SetPercentile(testValPercentile)
	assert.EqualValues(t, testValPercentile, ms.Percentile())

	assert.EqualValues(t, float64(0.0), ms.Value())
	testValValue := float64(17.13)
	ms.SetValue(testValValue)
	assert.EqualValues(t, testValValue, ms.Value())

	assert.EqualValues(t, generateTestSummaryValueAtPercentile(), ms)
}

func generateTestResourceMetricsSlice() ResourceMetricsSlice {
	tv := NewResourceMetricsSlice()
	fillTestResourceMetricsSlice(tv)
	return tv
}

func fillTestResourceMetricsSlice(tv ResourceMetricsSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestResourceMetrics(tv.At(i))
	}
}

func generateTestResourceMetrics() ResourceMetrics {
	tv := NewResourceMetrics()
	tv.InitEmpty()
	fillTestResourceMetrics(tv)
	return tv
}

func fillTestResourceMetrics(tv ResourceMetrics) {
	tv.Resource().InitEmpty()
	fillTestResource(tv.Resource())
	fillTestInstrumentationLibraryMetricsSlice(tv.InstrumentationLibraryMetrics())
}

func generateTestInstrumentationLibraryMetricsSlice() InstrumentationLibraryMetricsSlice {
	tv := NewInstrumentationLibraryMetricsSlice()
	fillTestInstrumentationLibraryMetricsSlice(tv)
	return tv
}

func fillTestInstrumentationLibraryMetricsSlice(tv InstrumentationLibraryMetricsSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestInstrumentationLibraryMetrics(tv.At(i))
	}
}

func generateTestInstrumentationLibraryMetrics() InstrumentationLibraryMetrics {
	tv := NewInstrumentationLibraryMetrics()
	tv.InitEmpty()
	fillTestInstrumentationLibraryMetrics(tv)
	return tv
}

func fillTestInstrumentationLibraryMetrics(tv InstrumentationLibraryMetrics) {
	tv.InstrumentationLibrary().InitEmpty()
	fillTestInstrumentationLibrary(tv.InstrumentationLibrary())
	fillTestMetricSlice(tv.Metrics())
}

func generateTestMetricSlice() MetricSlice {
	tv := NewMetricSlice()
	fillTestMetricSlice(tv)
	return tv
}

func fillTestMetricSlice(tv MetricSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestMetric(tv.At(i))
	}
}

func generateTestMetric() Metric {
	tv := NewMetric()
	tv.InitEmpty()
	fillTestMetric(tv)
	return tv
}

func fillTestMetric(tv Metric) {
	tv.MetricDescriptor().InitEmpty()
	fillTestMetricDescriptor(tv.MetricDescriptor())
	fillTestInt64DataPointSlice(tv.Int64DataPoints())
	fillTestDoubleDataPointSlice(tv.DoubleDataPoints())
	fillTestHistogramDataPointSlice(tv.HistogramDataPoints())
	fillTestSummaryDataPointSlice(tv.SummaryDataPoints())
}

func generateTestMetricDescriptor() MetricDescriptor {
	tv := NewMetricDescriptor()
	tv.InitEmpty()
	fillTestMetricDescriptor(tv)
	return tv
}

func fillTestMetricDescriptor(tv MetricDescriptor) {
	tv.SetName("test_name")
	tv.SetDescription("test_description")
	tv.SetUnit("1")
	tv.SetType(MetricTypeGaugeInt64)
	fillTestStringMap(tv.LabelsMap())
}

func generateTestInt64DataPointSlice() Int64DataPointSlice {
	tv := NewInt64DataPointSlice()
	fillTestInt64DataPointSlice(tv)
	return tv
}

func fillTestInt64DataPointSlice(tv Int64DataPointSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestInt64DataPoint(tv.At(i))
	}
}

func generateTestInt64DataPoint() Int64DataPoint {
	tv := NewInt64DataPoint()
	tv.InitEmpty()
	fillTestInt64DataPoint(tv)
	return tv
}

func fillTestInt64DataPoint(tv Int64DataPoint) {
	fillTestStringMap(tv.LabelsMap())
	tv.SetStartTime(TimestampUnixNano(1234567890))
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetValue(int64(-17))
}

func generateTestDoubleDataPointSlice() DoubleDataPointSlice {
	tv := NewDoubleDataPointSlice()
	fillTestDoubleDataPointSlice(tv)
	return tv
}

func fillTestDoubleDataPointSlice(tv DoubleDataPointSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestDoubleDataPoint(tv.At(i))
	}
}

func generateTestDoubleDataPoint() DoubleDataPoint {
	tv := NewDoubleDataPoint()
	tv.InitEmpty()
	fillTestDoubleDataPoint(tv)
	return tv
}

func fillTestDoubleDataPoint(tv DoubleDataPoint) {
	fillTestStringMap(tv.LabelsMap())
	tv.SetStartTime(TimestampUnixNano(1234567890))
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetValue(float64(17.13))
}

func generateTestHistogramDataPointSlice() HistogramDataPointSlice {
	tv := NewHistogramDataPointSlice()
	fillTestHistogramDataPointSlice(tv)
	return tv
}

func fillTestHistogramDataPointSlice(tv HistogramDataPointSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestHistogramDataPoint(tv.At(i))
	}
}

func generateTestHistogramDataPoint() HistogramDataPoint {
	tv := NewHistogramDataPoint()
	tv.InitEmpty()
	fillTestHistogramDataPoint(tv)
	return tv
}

func fillTestHistogramDataPoint(tv HistogramDataPoint) {
	fillTestStringMap(tv.LabelsMap())
	tv.SetStartTime(TimestampUnixNano(1234567890))
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetCount(uint64(17))
	tv.SetSum(float64(17.13))
	fillTestHistogramBucketSlice(tv.Buckets())
	tv.SetExplicitBounds([]float64{1, 2, 3})
}

func generateTestHistogramBucketSlice() HistogramBucketSlice {
	tv := NewHistogramBucketSlice()
	fillTestHistogramBucketSlice(tv)
	return tv
}

func fillTestHistogramBucketSlice(tv HistogramBucketSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestHistogramBucket(tv.At(i))
	}
}

func generateTestHistogramBucket() HistogramBucket {
	tv := NewHistogramBucket()
	tv.InitEmpty()
	fillTestHistogramBucket(tv)
	return tv
}

func fillTestHistogramBucket(tv HistogramBucket) {
	tv.SetCount(uint64(17))
	tv.Exemplar().InitEmpty()
	fillTestHistogramBucketExemplar(tv.Exemplar())
}

func generateTestHistogramBucketExemplar() HistogramBucketExemplar {
	tv := NewHistogramBucketExemplar()
	tv.InitEmpty()
	fillTestHistogramBucketExemplar(tv)
	return tv
}

func fillTestHistogramBucketExemplar(tv HistogramBucketExemplar) {
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetValue(float64(17.13))
	fillTestStringMap(tv.Attachments())
}

func generateTestSummaryDataPointSlice() SummaryDataPointSlice {
	tv := NewSummaryDataPointSlice()
	fillTestSummaryDataPointSlice(tv)
	return tv
}

func fillTestSummaryDataPointSlice(tv SummaryDataPointSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestSummaryDataPoint(tv.At(i))
	}
}

func generateTestSummaryDataPoint() SummaryDataPoint {
	tv := NewSummaryDataPoint()
	tv.InitEmpty()
	fillTestSummaryDataPoint(tv)
	return tv
}

func fillTestSummaryDataPoint(tv SummaryDataPoint) {
	fillTestStringMap(tv.LabelsMap())
	tv.SetStartTime(TimestampUnixNano(1234567890))
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetCount(uint64(17))
	tv.SetSum(float64(17.13))
	fillTestSummaryValueAtPercentileSlice(tv.ValueAtPercentiles())
}

func generateTestSummaryValueAtPercentileSlice() SummaryValueAtPercentileSlice {
	tv := NewSummaryValueAtPercentileSlice()
	fillTestSummaryValueAtPercentileSlice(tv)
	return tv
}

func fillTestSummaryValueAtPercentileSlice(tv SummaryValueAtPercentileSlice) {
	tv.Resize(13)
	for i := 0; i < tv.Len(); i++ {
		fillTestSummaryValueAtPercentile(tv.At(i))
	}
}

func generateTestSummaryValueAtPercentile() SummaryValueAtPercentile {
	tv := NewSummaryValueAtPercentile()
	tv.InitEmpty()
	fillTestSummaryValueAtPercentile(tv)
	return tv
}

func fillTestSummaryValueAtPercentile(tv SummaryValueAtPercentile) {
	tv.SetPercentile(float64(0.90))
	tv.SetValue(float64(17.13))
}