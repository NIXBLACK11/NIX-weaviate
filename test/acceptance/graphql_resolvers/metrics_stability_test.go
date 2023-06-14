//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package test

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/client/objects"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/test/helper"
	graphqlhelper "github.com/weaviate/weaviate/test/helper/graphql"
)

func metricsCount(t *testing.T) {
	defer cleanupMetricsClasses(t, 0, 20)
	createImportQueryMetricsClasses(t, 0, 10)
	metricsLinesBefore := countMetricsLines(t)
	createImportQueryMetricsClasses(t, 10, 20)
	metricsLinesAfter := countMetricsLines(t)
	assert.Equal(t, metricsLinesBefore, metricsLinesAfter, "number of metrics should not have changed")
}

func createImportQueryMetricsClasses(t *testing.T, start, end int) {
	for i := start; i < end; i++ {
		createMetricsClass(t, i)
		importMetricsClass(t, i)
		queryMetricsClass(t, i)
	}
}

func createMetricsClass(t *testing.T, classIndex int) {
	createObjectClass(t, &models.Class{
		Class:      fmt.Sprintf("MetricsClass_%d", classIndex),
		Vectorizer: "none",
		Properties: []*models.Property{
			{
				Name:     "some_text",
				DataType: schema.DataTypeText.PropString(),
			},
		},
		VectorIndexConfig: map[string]any{
			"efConstruction": 10,
			"maxConnextions": 2,
			"ef":             10,
		},
	})
}

func queryMetricsClass(t *testing.T, classIndex int) {
	// object by ID which exists
	resp, err := helper.Client(t).Objects.
		ObjectsClassGet(
			objects.NewObjectsClassGetParams().
				WithID(intToUUID(1)).
				WithClassName(metricsClassName(classIndex)),
			nil)

	require.Nil(t, err)
	assert.NotNil(t, resp.Payload)

	// object by ID which doesn't exist
	// ignore any return values
	helper.Client(t).Objects.
		ObjectsClassGet(
			objects.NewObjectsClassGetParams().
				WithID(intToUUID(math.MaxUint64)).
				WithClassName(metricsClassName(classIndex)),
			nil)

	// vector search
	result := graphqlhelper.AssertGraphQL(t, helper.RootAuth,
		fmt.Sprintf(
			"{  Get { %s(nearVector:{vector: [0.3,0.3,0.7,0.7]}, limit:5) { some_text } } }",
			metricsClassName(classIndex),
		),
	)
	objs := result.Get("Get", metricsClassName(classIndex)).AsSlice()
	assert.Len(t, objs, 5)

	// filtered vector search (which has specific metrics)
	// vector search
	result = graphqlhelper.AssertGraphQL(t, helper.RootAuth,
		fmt.Sprintf(
			"{  Get { %s(nearVector:{vector:[0.3,0.3,0.7,0.7]}, limit:5, where: %s) { some_text } } }",
			metricsClassName(classIndex),
			`{operator:Equal, valueText: "individually", path:["some_text"]}`,
		),
	)
	objs = result.Get("Get", metricsClassName(classIndex)).AsSlice()
	assert.Len(t, objs, 1)
}

func intToUUID(in uint64) strfmt.UUID {
	id := [16]byte{}
	binary.BigEndian.PutUint64(id[8:16], in)
	return strfmt.UUID(uuid.UUID(id).String())
}

// make sure that we use both individual as well as batch imports, as they
// might produce different metrics
func importMetricsClass(t *testing.T, classIndex int) {
	// individual
	createObject(t, &models.Object{
		Class: metricsClassName(classIndex),
		Properties: map[string]interface{}{
			"some_text": "this object was created individually",
		},
		ID:     intToUUID(1),
		Vector: randomVector(4),
	})

	// with batches
	const (
		batchSize  = 100
		numBatches = 50
	)

	for i := 0; i < numBatches; i++ {
		batch := make([]*models.Object, batchSize)
		for j := 0; j < batchSize; j++ {
			batch[j] = &models.Object{
				Class: metricsClassName(classIndex),
				Properties: map[string]interface{}{
					"some_text": fmt.Sprintf("this is object %d of batch %d", j, i),
				},
				Vector: randomVector(4),
			}
		}

		createObjectsBatch(t, batch)
	}
}

func cleanupMetricsClasses(t *testing.T, start, end int) {
	for i := start; i < end; i++ {
		deleteObjectClass(t, fmt.Sprintf("MetricsClass_%d", i))
	}
}

func randomVector(dims int) []float32 {
	out := make([]float32, dims)
	for i := range out {
		out[i] = rand.Float32()
	}
	return out
}

func countMetricsLines(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://localhost:2112/metrics", nil)
	require.Nil(t, err)

	c := &http.Client{}
	res, err := c.Do(req)
	require.Nil(t, err)

	defer res.Body.Close()
	require.Equal(t, http.StatusOK, res.StatusCode)

	scanner := bufio.NewScanner(res.Body)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	require.Nil(t, scanner.Err())

	return lineCount
}

func metricsClassName(classIndex int) string {
	return fmt.Sprintf("MetricsClass_%d", classIndex)
}
