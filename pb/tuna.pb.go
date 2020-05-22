// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/tuna.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceMetadata struct {
	Ip              string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	TcpPort         uint32   `protobuf:"varint,2,opt,name=tcp_port,json=tcpPort,proto3" json:"tcp_port,omitempty"`
	UdpPort         uint32   `protobuf:"varint,3,opt,name=udp_port,json=udpPort,proto3" json:"udp_port,omitempty"`
	ServiceId       uint32   `protobuf:"varint,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceTcp      []uint32 `protobuf:"varint,5,rep,packed,name=service_tcp,json=serviceTcp" json:"service_tcp,omitempty"`
	ServiceUdp      []uint32 `protobuf:"varint,6,rep,packed,name=service_udp,json=serviceUdp" json:"service_udp,omitempty"`
	Price           string   `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
	BeneficiaryAddr string   `protobuf:"bytes,8,opt,name=beneficiary_addr,json=beneficiaryAddr,proto3" json:"beneficiary_addr,omitempty"`
}

func (m *ServiceMetadata) Reset()      { *m = ServiceMetadata{} }
func (*ServiceMetadata) ProtoMessage() {}
func (*ServiceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tuna_607c72e3a0605a47, []int{0}
}
func (m *ServiceMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ServiceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMetadata.Merge(dst, src)
}
func (m *ServiceMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ServiceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMetadata proto.InternalMessageInfo

func (m *ServiceMetadata) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ServiceMetadata) GetTcpPort() uint32 {
	if m != nil {
		return m.TcpPort
	}
	return 0
}

func (m *ServiceMetadata) GetUdpPort() uint32 {
	if m != nil {
		return m.UdpPort
	}
	return 0
}

func (m *ServiceMetadata) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ServiceMetadata) GetServiceTcp() []uint32 {
	if m != nil {
		return m.ServiceTcp
	}
	return nil
}

func (m *ServiceMetadata) GetServiceUdp() []uint32 {
	if m != nil {
		return m.ServiceUdp
	}
	return nil
}

func (m *ServiceMetadata) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *ServiceMetadata) GetBeneficiaryAddr() string {
	if m != nil {
		return m.BeneficiaryAddr
	}
	return ""
}

type StreamMetadata struct {
	ServiceId uint32 `protobuf:"varint,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	PortId    uint32 `protobuf:"varint,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	IsPayment bool   `protobuf:"varint,3,opt,name=is_payment,json=isPayment,proto3" json:"is_payment,omitempty"`
}

func (m *StreamMetadata) Reset()      { *m = StreamMetadata{} }
func (*StreamMetadata) ProtoMessage() {}
func (*StreamMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tuna_607c72e3a0605a47, []int{1}
}
func (m *StreamMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StreamMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMetadata.Merge(dst, src)
}
func (m *StreamMetadata) XXX_Size() int {
	return m.Size()
}
func (m *StreamMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMetadata proto.InternalMessageInfo

func (m *StreamMetadata) GetServiceId() uint32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *StreamMetadata) GetPortId() uint32 {
	if m != nil {
		return m.PortId
	}
	return 0
}

func (m *StreamMetadata) GetIsPayment() bool {
	if m != nil {
		return m.IsPayment
	}
	return false
}

func init() {
	proto.RegisterType((*ServiceMetadata)(nil), "pb.ServiceMetadata")
	proto.RegisterType((*StreamMetadata)(nil), "pb.StreamMetadata")
}
func (this *ServiceMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceMetadata)
	if !ok {
		that2, ok := that.(ServiceMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Ip != that1.Ip {
		return false
	}
	if this.TcpPort != that1.TcpPort {
		return false
	}
	if this.UdpPort != that1.UdpPort {
		return false
	}
	if this.ServiceId != that1.ServiceId {
		return false
	}
	if len(this.ServiceTcp) != len(that1.ServiceTcp) {
		return false
	}
	for i := range this.ServiceTcp {
		if this.ServiceTcp[i] != that1.ServiceTcp[i] {
			return false
		}
	}
	if len(this.ServiceUdp) != len(that1.ServiceUdp) {
		return false
	}
	for i := range this.ServiceUdp {
		if this.ServiceUdp[i] != that1.ServiceUdp[i] {
			return false
		}
	}
	if this.Price != that1.Price {
		return false
	}
	if this.BeneficiaryAddr != that1.BeneficiaryAddr {
		return false
	}
	return true
}
func (this *StreamMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StreamMetadata)
	if !ok {
		that2, ok := that.(StreamMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceId != that1.ServiceId {
		return false
	}
	if this.PortId != that1.PortId {
		return false
	}
	if this.IsPayment != that1.IsPayment {
		return false
	}
	return true
}
func (this *ServiceMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&pb.ServiceMetadata{")
	s = append(s, "Ip: "+fmt.Sprintf("%#v", this.Ip)+",\n")
	s = append(s, "TcpPort: "+fmt.Sprintf("%#v", this.TcpPort)+",\n")
	s = append(s, "UdpPort: "+fmt.Sprintf("%#v", this.UdpPort)+",\n")
	s = append(s, "ServiceId: "+fmt.Sprintf("%#v", this.ServiceId)+",\n")
	s = append(s, "ServiceTcp: "+fmt.Sprintf("%#v", this.ServiceTcp)+",\n")
	s = append(s, "ServiceUdp: "+fmt.Sprintf("%#v", this.ServiceUdp)+",\n")
	s = append(s, "Price: "+fmt.Sprintf("%#v", this.Price)+",\n")
	s = append(s, "BeneficiaryAddr: "+fmt.Sprintf("%#v", this.BeneficiaryAddr)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StreamMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.StreamMetadata{")
	s = append(s, "ServiceId: "+fmt.Sprintf("%#v", this.ServiceId)+",\n")
	s = append(s, "PortId: "+fmt.Sprintf("%#v", this.PortId)+",\n")
	s = append(s, "IsPayment: "+fmt.Sprintf("%#v", this.IsPayment)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTuna(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ServiceMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ip) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTuna(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.TcpPort != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTuna(dAtA, i, uint64(m.TcpPort))
	}
	if m.UdpPort != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTuna(dAtA, i, uint64(m.UdpPort))
	}
	if m.ServiceId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTuna(dAtA, i, uint64(m.ServiceId))
	}
	if len(m.ServiceTcp) > 0 {
		dAtA2 := make([]byte, len(m.ServiceTcp)*10)
		var j1 int
		for _, num := range m.ServiceTcp {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTuna(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.ServiceUdp) > 0 {
		dAtA4 := make([]byte, len(m.ServiceUdp)*10)
		var j3 int
		for _, num := range m.ServiceUdp {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintTuna(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if len(m.Price) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintTuna(dAtA, i, uint64(len(m.Price)))
		i += copy(dAtA[i:], m.Price)
	}
	if len(m.BeneficiaryAddr) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintTuna(dAtA, i, uint64(len(m.BeneficiaryAddr)))
		i += copy(dAtA[i:], m.BeneficiaryAddr)
	}
	return i, nil
}

func (m *StreamMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ServiceId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTuna(dAtA, i, uint64(m.ServiceId))
	}
	if m.PortId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTuna(dAtA, i, uint64(m.PortId))
	}
	if m.IsPayment {
		dAtA[i] = 0x18
		i++
		if m.IsPayment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintTuna(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedServiceMetadata(r randyTuna, easy bool) *ServiceMetadata {
	this := &ServiceMetadata{}
	this.Ip = string(randStringTuna(r))
	this.TcpPort = uint32(r.Uint32())
	this.UdpPort = uint32(r.Uint32())
	this.ServiceId = uint32(r.Uint32())
	v1 := r.Intn(10)
	this.ServiceTcp = make([]uint32, v1)
	for i := 0; i < v1; i++ {
		this.ServiceTcp[i] = uint32(r.Uint32())
	}
	v2 := r.Intn(10)
	this.ServiceUdp = make([]uint32, v2)
	for i := 0; i < v2; i++ {
		this.ServiceUdp[i] = uint32(r.Uint32())
	}
	this.Price = string(randStringTuna(r))
	this.BeneficiaryAddr = string(randStringTuna(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStreamMetadata(r randyTuna, easy bool) *StreamMetadata {
	this := &StreamMetadata{}
	this.ServiceId = uint32(r.Uint32())
	this.PortId = uint32(r.Uint32())
	this.IsPayment = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTuna interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTuna(r randyTuna) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTuna(r randyTuna) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTuna(r)
	}
	return string(tmps)
}
func randUnrecognizedTuna(r randyTuna, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTuna(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTuna(dAtA []byte, r randyTuna, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTuna(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTuna(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ServiceMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovTuna(uint64(l))
	}
	if m.TcpPort != 0 {
		n += 1 + sovTuna(uint64(m.TcpPort))
	}
	if m.UdpPort != 0 {
		n += 1 + sovTuna(uint64(m.UdpPort))
	}
	if m.ServiceId != 0 {
		n += 1 + sovTuna(uint64(m.ServiceId))
	}
	if len(m.ServiceTcp) > 0 {
		l = 0
		for _, e := range m.ServiceTcp {
			l += sovTuna(uint64(e))
		}
		n += 1 + sovTuna(uint64(l)) + l
	}
	if len(m.ServiceUdp) > 0 {
		l = 0
		for _, e := range m.ServiceUdp {
			l += sovTuna(uint64(e))
		}
		n += 1 + sovTuna(uint64(l)) + l
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTuna(uint64(l))
	}
	l = len(m.BeneficiaryAddr)
	if l > 0 {
		n += 1 + l + sovTuna(uint64(l))
	}
	return n
}

func (m *StreamMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceId != 0 {
		n += 1 + sovTuna(uint64(m.ServiceId))
	}
	if m.PortId != 0 {
		n += 1 + sovTuna(uint64(m.PortId))
	}
	if m.IsPayment {
		n += 2
	}
	return n
}

func sovTuna(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTuna(x uint64) (n int) {
	return sovTuna(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ServiceMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceMetadata{`,
		`Ip:` + fmt.Sprintf("%v", this.Ip) + `,`,
		`TcpPort:` + fmt.Sprintf("%v", this.TcpPort) + `,`,
		`UdpPort:` + fmt.Sprintf("%v", this.UdpPort) + `,`,
		`ServiceId:` + fmt.Sprintf("%v", this.ServiceId) + `,`,
		`ServiceTcp:` + fmt.Sprintf("%v", this.ServiceTcp) + `,`,
		`ServiceUdp:` + fmt.Sprintf("%v", this.ServiceUdp) + `,`,
		`Price:` + fmt.Sprintf("%v", this.Price) + `,`,
		`BeneficiaryAddr:` + fmt.Sprintf("%v", this.BeneficiaryAddr) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StreamMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StreamMetadata{`,
		`ServiceId:` + fmt.Sprintf("%v", this.ServiceId) + `,`,
		`PortId:` + fmt.Sprintf("%v", this.PortId) + `,`,
		`IsPayment:` + fmt.Sprintf("%v", this.IsPayment) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTuna(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ServiceMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTuna
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTuna
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpPort", wireType)
			}
			m.TcpPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TcpPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UdpPort", wireType)
			}
			m.UdpPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UdpPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			m.ServiceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTuna
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ServiceTcp = append(m.ServiceTcp, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTuna
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTuna
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTuna
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ServiceTcp = append(m.ServiceTcp, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceTcp", wireType)
			}
		case 6:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTuna
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ServiceUdp = append(m.ServiceUdp, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTuna
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTuna
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTuna
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ServiceUdp = append(m.ServiceUdp, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceUdp", wireType)
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTuna
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeneficiaryAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTuna
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeneficiaryAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTuna(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTuna
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTuna
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			m.ServiceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			m.PortId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PortId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPayment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPayment = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTuna(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTuna
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTuna(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTuna
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTuna
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTuna
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTuna
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTuna(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTuna = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTuna   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/tuna.proto", fileDescriptor_tuna_607c72e3a0605a47) }

var fileDescriptor_tuna_607c72e3a0605a47 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3f, 0x6f, 0xe2, 0x30,
	0x18, 0xc6, 0xe3, 0x70, 0xfc, 0xf3, 0x09, 0x38, 0x45, 0x27, 0x5d, 0xee, 0xa4, 0x7b, 0x8b, 0x98,
	0xe8, 0x50, 0x18, 0xba, 0x76, 0x69, 0x37, 0x86, 0x4a, 0x28, 0xb4, 0x73, 0xe4, 0xd8, 0x26, 0xf5,
	0x00, 0xb1, 0x8c, 0x53, 0x89, 0xad, 0x1f, 0xa1, 0x1f, 0xa3, 0x1f, 0xa1, 0x1f, 0xa1, 0x23, 0x23,
	0x63, 0x63, 0x96, 0x8e, 0x8c, 0xed, 0x56, 0xc5, 0x86, 0x0a, 0xba, 0xe5, 0x79, 0x7f, 0xcf, 0x13,
	0x3d, 0xef, 0x6b, 0xdc, 0x92, 0xc9, 0x50, 0xe7, 0x73, 0x32, 0x90, 0x2a, 0xd3, 0x59, 0xe0, 0xcb,
	0xe4, 0xdf, 0x59, 0x2a, 0xf4, 0x5d, 0x9e, 0x0c, 0x68, 0x36, 0x1b, 0xa6, 0x59, 0x9a, 0x0d, 0x2d,
	0x4a, 0xf2, 0xa9, 0x55, 0x56, 0xd8, 0x2f, 0x17, 0xe9, 0x7d, 0x20, 0xdc, 0x99, 0x70, 0x75, 0x2f,
	0x28, 0xbf, 0xe6, 0x9a, 0x30, 0xa2, 0x49, 0xd0, 0xc6, 0xbe, 0x90, 0x21, 0xea, 0xa2, 0x7e, 0x33,
	0xf2, 0x85, 0x0c, 0xfe, 0xe2, 0x86, 0xa6, 0x32, 0x96, 0x99, 0xd2, 0xa1, 0xdf, 0x45, 0xfd, 0x56,
	0x54, 0xd7, 0x54, 0x8e, 0x33, 0xa5, 0x4b, 0x94, 0xb3, 0x1d, 0xaa, 0x38, 0x94, 0x33, 0x87, 0xfe,
	0x63, 0xbc, 0x70, 0x3f, 0x8e, 0x05, 0x0b, 0x7f, 0x58, 0xd8, 0xdc, 0x4d, 0x46, 0x2c, 0x38, 0xc1,
	0x3f, 0xf7, 0x58, 0x53, 0x19, 0x56, 0xbb, 0x95, 0x7e, 0x2b, 0xda, 0x27, 0x6e, 0xa8, 0x3c, 0x34,
	0xe4, 0x4c, 0x86, 0xb5, 0x23, 0xc3, 0x2d, 0x93, 0xc1, 0x6f, 0x5c, 0x95, 0x4a, 0x50, 0x1e, 0xd6,
	0x6d, 0x53, 0x27, 0x82, 0x53, 0xfc, 0x2b, 0xe1, 0x73, 0x3e, 0x15, 0x54, 0x10, 0xb5, 0x8c, 0x09,
	0x63, 0x2a, 0x6c, 0x58, 0x43, 0xe7, 0x60, 0x7e, 0xc9, 0x98, 0xea, 0xa5, 0xb8, 0x3d, 0xd1, 0x8a,
	0x93, 0xd9, 0xd7, 0xe6, 0xc7, 0x9d, 0xd1, 0xf7, 0xce, 0x7f, 0x70, 0xbd, 0xdc, 0xb4, 0x64, 0xee,
	0x0e, 0xb5, 0x52, 0x8e, 0x58, 0x99, 0x13, 0x8b, 0x58, 0x92, 0xe5, 0x8c, 0xcf, 0xdd, 0x21, 0x1a,
	0x51, 0x53, 0x2c, 0xc6, 0x6e, 0x70, 0x75, 0xb1, 0x2a, 0xc0, 0x5b, 0x17, 0xe0, 0x6d, 0x0b, 0x40,
	0xef, 0x05, 0xa0, 0x07, 0x03, 0xe8, 0xc9, 0x00, 0x7a, 0x36, 0x80, 0x5e, 0x0c, 0xa0, 0x95, 0x01,
	0xf4, 0x6a, 0x00, 0xbd, 0x19, 0xf0, 0xb6, 0x06, 0xd0, 0xe3, 0x06, 0xbc, 0xd5, 0x06, 0xbc, 0xf5,
	0x06, 0xbc, 0xa4, 0x66, 0x5f, 0xea, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xef, 0x55, 0xc0,
	0xed, 0x01, 0x00, 0x00,
}
