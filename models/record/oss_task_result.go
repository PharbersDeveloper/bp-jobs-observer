// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     OssTask.avsc
 *     OssTaskResult.avsc
 */

package record

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type OssTaskResult struct {
	JobId    string
	TraceId  string
	Progress int64
	Error    string
}

func NewOssTaskResultWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &OssTaskResult{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func DeserializeOssTaskResult(r io.Reader) (*OssTaskResult, error) {
	t := NewOssTaskResult()

	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	return t, err
}

func NewOssTaskResult() *OssTaskResult {
	return &OssTaskResult{}
}

func (r *OssTaskResult) Schema() string {
	return "{\"fields\":[{\"name\":\"jobId\",\"type\":\"string\"},{\"name\":\"traceId\",\"type\":\"string\"},{\"name\":\"progress\",\"type\":\"long\"},{\"name\":\"error\",\"type\":\"string\"}],\"name\":\"OssTaskResult\",\"namespace\":\"com.pharbers.kafka.schema\",\"type\":\"record\"}"
}

func (r *OssTaskResult) SchemaName() string {
	return "com.pharbers.kafka.schema.OssTaskResult"
}

func (r *OssTaskResult) Serialize(w io.Writer) error {
	return writeOssTaskResult(r, w)
}

func (_ *OssTaskResult) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *OssTaskResult) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *OssTaskResult) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *OssTaskResult) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *OssTaskResult) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *OssTaskResult) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *OssTaskResult) SetString(v string)   { panic("Unsupported operation") }
func (_ *OssTaskResult) SetUnionElem(v int64) { panic("Unsupported operation") }
func (r *OssTaskResult) Get(i int) types.Field {
	switch i {
	case 0:
		return (*types.String)(&r.JobId)
	case 1:
		return (*types.String)(&r.TraceId)
	case 2:
		return (*types.Long)(&r.Progress)
	case 3:
		return (*types.String)(&r.Error)

	}
	panic("Unknown field index")
}
func (r *OssTaskResult) SetDefault(i int) {
	switch i {

	}
	panic("Unknown field index")
}
func (_ *OssTaskResult) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *OssTaskResult) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *OssTaskResult) Finalize()                        {}

type OssTaskResultReader struct {
	r io.Reader
	p *vm.Program
}

func NewOssTaskResultReader(r io.Reader) (*OssTaskResultReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewOssTaskResult()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &OssTaskResultReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r *OssTaskResultReader) Read() (*OssTaskResult, error) {
	t := NewOssTaskResult()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
