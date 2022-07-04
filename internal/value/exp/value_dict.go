package value

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/value/exp/allocator"
)

type (
	dictField struct {
		k V
		v V
	}
	dictValue struct {
		fields []dictField
	}
)

func (v *dictValue) toYDBType(a *allocator.Allocator) *Ydb.Type {
	var fields []dictField
	if v != nil {
		fields = v.fields
	}

	t := a.Type()

	typeDict := a.TypeDict()

	typeDict.DictType = a.Dict()

	typeDict.DictType.Key = fields[0].k.toYDBType(a)
	typeDict.DictType.Payload = fields[0].v.toYDBType(a)

	t.Type = typeDict

	return t
}

func (v *dictValue) toYDBValue(a *allocator.Allocator) *Ydb.Value {
	var fields []dictField
	if v != nil {
		fields = v.fields
	}
	vvv := a.Value()

	for _, vv := range fields {
		pair := a.Pair()

		pair.Key = vv.k.toYDBValue(a)
		pair.Payload = vv.v.toYDBValue(a)

		vvv.Pairs = append(vvv.Pairs, pair)
	}

	return vvv
}

func DictField(k, v V) dictField {
	return dictField{
		k: k,
		v: v,
	}
}

func DictValue(v ...dictField) *dictValue {
	return &dictValue{fields: v}
}