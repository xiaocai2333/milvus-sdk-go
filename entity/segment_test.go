package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhagnlu/milvus-sdk-go/v2/internal/proto/common"
)

func TestSegmentFlushed(t *testing.T) {
	segment := &Segment{}
	assert.False(t, segment.Flushed())
	segment.State = common.SegmentState_Growing
	assert.False(t, segment.Flushed())
	segment.State = common.SegmentState_Flushing
	assert.False(t, segment.Flushed())
	segment.State = common.SegmentState_Flushed
	assert.True(t, segment.Flushed())
}
