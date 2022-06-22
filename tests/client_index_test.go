package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhagnlu/milvus-sdk-go/v2/client"
	"github.com/zhagnlu/milvus-sdk-go/v2/entity"
)

func TestCreateIndex(t *testing.T) {
	c, err := client.NewGrpcClient(context.Background(), "localhost:19530")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	if c != nil {
		defer c.Close()
	}

	idx := entity.NewGenericIndex("", entity.Flat, map[string]string{
		"nlist":       "1024",
		"metric_type": "IP",
	})
	err = c.CreateIndex(context.Background(), "test_go_sdk", "vector", idx, false)
	assert.Nil(t, err)
	indexes, err := c.DescribeIndex(context.Background(), "test_go_sdk", "vector")
	if assert.Nil(t, err) {
		for _, idx := range indexes {
			t.Log(idx.IndexType())
			t.Log(idx.Params())
		}
	}
}
