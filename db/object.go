package db

import "strings"

type Object interface {
	i()
	SetAndValues(val interface{})
	SetAndParams(param string)
	SetOrValues(val interface{})
	SetOrParams(param string)
	GetAndValues() []interface{}
	GetAndParams() []string
	GetOrValues() []interface{}
	GetOrParams() []string
	GetOrParamsToBuf(buf *strings.Builder, flag string)
	GetAndParamsToBuf(buf *strings.Builder, flag string)
}

type obj struct {
	andValues []interface{}
	andParams []string
	orValues  []interface{}
	orParams  []string
}

func (*obj) i() {}

func NewDbObject() Object {
	return new(obj)
}

func (o *obj) SetAndValues(val interface{}) {
	o.andValues = append(o.andValues, val)
}

func (o *obj) SetAndParams(param string) {
	o.andParams = append(o.andParams, param)
}

func (o *obj) SetOrValues(val interface{}) {
	o.orValues = append(o.orValues, val)
}
func (o *obj) SetOrParams(param string) {
	o.orParams = append(o.orParams, param)
}

func (o *obj) GetAndValues() []interface{} {
	return o.andValues
}

func (o *obj) GetAndParams() []string {
	return o.andParams
}

func (o *obj) GetOrValues() []interface{} {
	return o.orValues
}

func (o *obj) GetOrParams() []string {
	return o.orParams
}

func (o *obj) GetAndParamsToBuf(buf *strings.Builder, flag string) {
	if len(o.andParams) != 0 {
		buf.WriteString(o.andParams[0])
		for _, param := range o.andParams[1:] {
			buf.WriteString(flag)
			buf.WriteString(param)
		}
	}
}

func (o *obj) GetOrParamsToBuf(buf *strings.Builder, flag string) {
	if len(o.orParams) != 0 {
		buf.WriteString(o.orParams[0])
		for _, param := range o.orParams[1:] {
			buf.WriteString(flag)
			buf.WriteString(param)
		}
	}
}
