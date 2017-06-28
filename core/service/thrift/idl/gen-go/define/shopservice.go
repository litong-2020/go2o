// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package define

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ShopService interface {
	// Parameters:
	//  - VenderId
	GetStore(venderId int32) (r *Shop, err error)
	// Parameters:
	//  - ShopId
	GetStoreById(shopId int32) (r *Shop, err error)
}

type ShopServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewShopServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ShopServiceClient {
	return &ShopServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewShopServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ShopServiceClient {
	return &ShopServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - VenderId
func (p *ShopServiceClient) GetStore(venderId int32) (r *Shop, err error) {
	if err = p.sendGetStore(venderId); err != nil {
		return
	}
	return p.recvGetStore()
}

func (p *ShopServiceClient) sendGetStore(venderId int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetStore", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ShopServiceGetStoreArgs{
		VenderId: venderId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ShopServiceClient) recvGetStore() (value *Shop, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetStore" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetStore failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetStore failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error307 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error308 error
		error308, err = error307.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error308
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetStore failed: invalid message type")
		return
	}
	result := ShopServiceGetStoreResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - ShopId
func (p *ShopServiceClient) GetStoreById(shopId int32) (r *Shop, err error) {
	if err = p.sendGetStoreById(shopId); err != nil {
		return
	}
	return p.recvGetStoreById()
}

func (p *ShopServiceClient) sendGetStoreById(shopId int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetStoreById", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ShopServiceGetStoreByIdArgs{
		ShopId: shopId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ShopServiceClient) recvGetStoreById() (value *Shop, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetStoreById" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetStoreById failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetStoreById failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error309 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error310 error
		error310, err = error309.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error310
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetStoreById failed: invalid message type")
		return
	}
	result := ShopServiceGetStoreByIdResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ShopServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ShopService
}

func (p *ShopServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ShopServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ShopServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewShopServiceProcessor(handler ShopService) *ShopServiceProcessor {

	self311 := &ShopServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self311.processorMap["GetStore"] = &shopServiceProcessorGetStore{handler: handler}
	self311.processorMap["GetStoreById"] = &shopServiceProcessorGetStoreById{handler: handler}
	return self311
}

func (p *ShopServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x312 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x312.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x312

}

type shopServiceProcessorGetStore struct {
	handler ShopService
}

func (p *shopServiceProcessorGetStore) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ShopServiceGetStoreArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetStore", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ShopServiceGetStoreResult{}
	var retval *Shop
	var err2 error
	if retval, err2 = p.handler.GetStore(args.VenderId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetStore: "+err2.Error())
		oprot.WriteMessageBegin("GetStore", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetStore", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type shopServiceProcessorGetStoreById struct {
	handler ShopService
}

func (p *shopServiceProcessorGetStoreById) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ShopServiceGetStoreByIdArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetStoreById", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ShopServiceGetStoreByIdResult{}
	var retval *Shop
	var err2 error
	if retval, err2 = p.handler.GetStoreById(args.ShopId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetStoreById: "+err2.Error())
		oprot.WriteMessageBegin("GetStoreById", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetStoreById", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - VenderId
type ShopServiceGetStoreArgs struct {
	VenderId int32 `thrift:"venderId,1" json:"venderId"`
}

func NewShopServiceGetStoreArgs() *ShopServiceGetStoreArgs {
	return &ShopServiceGetStoreArgs{}
}

func (p *ShopServiceGetStoreArgs) GetVenderId() int32 {
	return p.VenderId
}
func (p *ShopServiceGetStoreArgs) Read(iprot thrift.TProtocol) error {
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

func (p *ShopServiceGetStoreArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.VenderId = v
	}
	return nil
}

func (p *ShopServiceGetStoreArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetStore_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *ShopServiceGetStoreArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("venderId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:venderId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.VenderId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.venderId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:venderId: ", p), err)
	}
	return err
}

func (p *ShopServiceGetStoreArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceGetStoreArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ShopServiceGetStoreResult struct {
	Success *Shop `thrift:"success,0" json:"success,omitempty"`
}

func NewShopServiceGetStoreResult() *ShopServiceGetStoreResult {
	return &ShopServiceGetStoreResult{}
}

var ShopServiceGetStoreResult_Success_DEFAULT *Shop

func (p *ShopServiceGetStoreResult) GetSuccess() *Shop {
	if !p.IsSetSuccess() {
		return ShopServiceGetStoreResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ShopServiceGetStoreResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ShopServiceGetStoreResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *ShopServiceGetStoreResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Shop{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *ShopServiceGetStoreResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetStore_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *ShopServiceGetStoreResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ShopServiceGetStoreResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceGetStoreResult(%+v)", *p)
}

// Attributes:
//  - ShopId
type ShopServiceGetStoreByIdArgs struct {
	ShopId int32 `thrift:"shopId,1" json:"shopId"`
}

func NewShopServiceGetStoreByIdArgs() *ShopServiceGetStoreByIdArgs {
	return &ShopServiceGetStoreByIdArgs{}
}

func (p *ShopServiceGetStoreByIdArgs) GetShopId() int32 {
	return p.ShopId
}
func (p *ShopServiceGetStoreByIdArgs) Read(iprot thrift.TProtocol) error {
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

func (p *ShopServiceGetStoreByIdArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ShopId = v
	}
	return nil
}

func (p *ShopServiceGetStoreByIdArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetStoreById_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *ShopServiceGetStoreByIdArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("shopId", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:shopId: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ShopId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.shopId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:shopId: ", p), err)
	}
	return err
}

func (p *ShopServiceGetStoreByIdArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceGetStoreByIdArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ShopServiceGetStoreByIdResult struct {
	Success *Shop `thrift:"success,0" json:"success,omitempty"`
}

func NewShopServiceGetStoreByIdResult() *ShopServiceGetStoreByIdResult {
	return &ShopServiceGetStoreByIdResult{}
}

var ShopServiceGetStoreByIdResult_Success_DEFAULT *Shop

func (p *ShopServiceGetStoreByIdResult) GetSuccess() *Shop {
	if !p.IsSetSuccess() {
		return ShopServiceGetStoreByIdResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ShopServiceGetStoreByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ShopServiceGetStoreByIdResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *ShopServiceGetStoreByIdResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Shop{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *ShopServiceGetStoreByIdResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetStoreById_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *ShopServiceGetStoreByIdResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ShopServiceGetStoreByIdResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShopServiceGetStoreByIdResult(%+v)", *p)
}
