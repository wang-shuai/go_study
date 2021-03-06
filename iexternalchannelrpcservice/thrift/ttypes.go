// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thrift

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// 类:渠道分组,描述:渠道分组
//
// Attributes:
//  - TypeId: 属性:渠道类型id,描述:主键id
//  - TypeName: 属性:渠道类型名,描述:渠道类型名
//  - Items: 属性:组内条目,描述:组内条目
type ExternalChannelGroup struct {
	TypeId   int32                  `thrift:"TypeId,1" json:"TypeId"`
	TypeName string                 `thrift:"TypeName,2" json:"TypeName"`
	Items    []*ExternalChannelItem `thrift:"Items,3" json:"Items"`
}

func NewExternalChannelGroup() *ExternalChannelGroup {
	return &ExternalChannelGroup{}
}

func (p *ExternalChannelGroup) GetTypeId() int32 {
	return p.TypeId
}

func (p *ExternalChannelGroup) GetTypeName() string {
	return p.TypeName
}

func (p *ExternalChannelGroup) GetItems() []*ExternalChannelItem {
	return p.Items
}
func (p *ExternalChannelGroup) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ExternalChannelGroup) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.TypeId = v
	}
	return nil
}

func (p *ExternalChannelGroup) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.TypeName = v
	}
	return nil
}

func (p *ExternalChannelGroup) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*ExternalChannelItem, 0, size)
	p.Items = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &ExternalChannelItem{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Items = append(p.Items, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ExternalChannelGroup) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ExternalChannelGroup"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ExternalChannelGroup) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TypeId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:TypeId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.TypeId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TypeId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:TypeId: ", p), err)
	}
	return err
}

func (p *ExternalChannelGroup) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TypeName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:TypeName: ", p), err)
	}
	if err := oprot.WriteString(string(p.TypeName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TypeName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:TypeName: ", p), err)
	}
	return err
}

func (p *ExternalChannelGroup) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Items", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Items: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Items)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Items {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Items: ", p), err)
	}
	return err
}

func (p *ExternalChannelGroup) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExternalChannelGroup(%+v)", *p)
}

// 类:渠道条目,描述:渠道条目
//
// Attributes:
//  - Code: 属性:渠道码,描述:渠道码
//  - Name: 属性:渠道名,描述:渠道名
//  - ParentId: 属性:所属渠道类型id,描述:所属渠道类型id
type ExternalChannelItem struct {
	Code     string `thrift:"Code,1" json:"Code"`
	Name     string `thrift:"Name,2" json:"Name"`
	ParentId int32  `thrift:"ParentId,3" json:"ParentId"`
}

func NewExternalChannelItem() *ExternalChannelItem {
	return &ExternalChannelItem{}
}

func (p *ExternalChannelItem) GetCode() string {
	return p.Code
}

func (p *ExternalChannelItem) GetName() string {
	return p.Name
}

func (p *ExternalChannelItem) GetParentId() int32 {
	return p.ParentId
}
func (p *ExternalChannelItem) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ExternalChannelItem) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *ExternalChannelItem) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *ExternalChannelItem) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ParentId = v
	}
	return nil
}

func (p *ExternalChannelItem) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ExternalChannelItem"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ExternalChannelItem) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Code", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Code: ", p), err)
	}
	if err := oprot.WriteString(string(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Code: ", p), err)
	}
	return err
}

func (p *ExternalChannelItem) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Name", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Name (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Name: ", p), err)
	}
	return err
}

func (p *ExternalChannelItem) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ParentId", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:ParentId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ParentId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ParentId (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:ParentId: ", p), err)
	}
	return err
}

func (p *ExternalChannelItem) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExternalChannelItem(%+v)", *p)
}
