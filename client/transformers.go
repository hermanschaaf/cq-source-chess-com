package client

import (
	"reflect"

	"github.com/agoblet/chesscompubapi"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

// type TypeTransformer func(reflect.StructField) (arrow.DataType, error)
var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	timestamp := chesscompubapi.UnixSecondsTimestamp{}
	switch field.Type {
	case reflect.TypeOf(timestamp), reflect.TypeOf(&timestamp):
		return arrow.FixedWidthTypes.Timestamp_us, nil
	default:
		return nil, nil
	}
}
