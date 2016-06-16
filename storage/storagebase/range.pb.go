// Code generated by protoc-gen-gogo.
// source: cockroach/storage/storagebase/range.proto
// DO NOT EDIT!

/*
	Package storagebase is a generated protocol buffer package.

	It is generated from these files:
		cockroach/storage/storagebase/range.proto

	It has these top-level messages:
		RangeState
		RangeInfo
*/
package storagebase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_storage_engine_enginepb "github.com/cockroachdb/cockroach/storage/engine/enginepb"
import cockroach_roachpb2 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/util/hlc"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// RangeState is the part of the Range Raft state machine which is cached in
// memory and which is manipulated exclusively through consensus.
type RangeState struct {
	// The highest (and last) index applied to the state machine.
	RaftAppliedIndex uint64 `protobuf:"varint,4,opt,name=raftAppliedIndex,proto3" json:"raftAppliedIndex,omitempty"`
	// The highest (and last) lease index applied to the state machine.
	LeaseAppliedIndex uint64 `protobuf:"varint,5,opt,name=leaseAppliedIndex,proto3" json:"leaseAppliedIndex,omitempty"`
	// The Range descriptor.
	// The pointer may change, but the referenced RangeDescriptor struct itself
	// must be treated as immutable; it is leaked out of the lock.
	//
	// Changes of the descriptor should normally go through one of the
	// (*Replica).setDesc* methods.
	Desc *cockroach_roachpb.RangeDescriptor `protobuf:"bytes,6,opt,name=desc" json:"desc,omitempty"`
	// The latest lease holder, if any.
	Lease *cockroach_roachpb1.Lease `protobuf:"bytes,7,opt,name=lease" json:"lease,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *cockroach_roachpb2.RaftTruncatedState `protobuf:"bytes,8,opt,name=truncatedState" json:"truncatedState,omitempty"`
	// gcThreshold is the GC threshold of the Range, typically updated when keys
	// are garbage collected. Reads and writes at timestamps <= this time will
	// not be served.
	GCThreshold cockroach_util_hlc.Timestamp                `protobuf:"bytes,9,opt,name=gcThreshold" json:"gcThreshold"`
	Stats       cockroach_storage_engine_enginepb.MVCCStats `protobuf:"bytes,10,opt,name=stats" json:"stats"`
	Frozen      bool                                        `protobuf:"varint,11,opt,name=frozen,proto3" json:"frozen,omitempty"`
}

func (m *RangeState) Reset()                    { *m = RangeState{} }
func (m *RangeState) String() string            { return proto.CompactTextString(m) }
func (*RangeState) ProtoMessage()               {}
func (*RangeState) Descriptor() ([]byte, []int) { return fileDescriptorRange, []int{0} }

type RangeInfo struct {
	RangeState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
}

func (m *RangeInfo) Reset()                    { *m = RangeInfo{} }
func (m *RangeInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeInfo) ProtoMessage()               {}
func (*RangeInfo) Descriptor() ([]byte, []int) { return fileDescriptorRange, []int{1} }

func init() {
	proto.RegisterType((*RangeState)(nil), "cockroach.storage.storagebase.RangeState")
	proto.RegisterType((*RangeInfo)(nil), "cockroach.storage.storagebase.RangeInfo")
}
func (m *RangeState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RangeState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintRange(data, i, uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintRange(data, i, uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRange(data, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lease != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintRange(data, i, uint64(m.Lease.Size()))
		n2, err := m.Lease.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.TruncatedState != nil {
		data[i] = 0x42
		i++
		i = encodeVarintRange(data, i, uint64(m.TruncatedState.Size()))
		n3, err := m.TruncatedState.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	data[i] = 0x4a
	i++
	i = encodeVarintRange(data, i, uint64(m.GCThreshold.Size()))
	n4, err := m.GCThreshold.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	data[i] = 0x52
	i++
	i = encodeVarintRange(data, i, uint64(m.Stats.Size()))
	n5, err := m.Stats.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.Frozen {
		data[i] = 0x58
		i++
		if m.Frozen {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RangeInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RangeInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintRange(data, i, uint64(m.RangeState.Size()))
	n6, err := m.RangeState.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.LastIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintRange(data, i, uint64(m.LastIndex))
	}
	return i, nil
}

func encodeFixed64Range(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Range(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRange(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RangeState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovRange(uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		n += 1 + sovRange(uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovRange(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovRange(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovRange(uint64(l))
	}
	l = m.GCThreshold.Size()
	n += 1 + l + sovRange(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovRange(uint64(l))
	if m.Frozen {
		n += 2
	}
	return n
}

func (m *RangeInfo) Size() (n int) {
	var l int
	_ = l
	l = m.RangeState.Size()
	n += 1 + l + sovRange(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovRange(uint64(m.LastIndex))
	}
	return n
}

func sovRange(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRange(x uint64) (n int) {
	return sovRange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RangeState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RangeState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseAppliedIndex", wireType)
			}
			m.LeaseAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LeaseAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &cockroach_roachpb1.Lease{}
			}
			if err := m.Lease.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &cockroach_roachpb2.RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GCThreshold.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Frozen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRange(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRange
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
func (m *RangeInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRange
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RangeState.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRange(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRange
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
func skipRange(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRange
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRange
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRange
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRange(data[start:])
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
	ErrInvalidLengthRange = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRange   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRange = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x5b, 0x6d, 0x6b, 0x9b, 0x01, 0xd1, 0x28, 0x12, 0x96, 0xdd, 0x76, 0x29, 0x14, 0x5c,
	0x29, 0x19, 0x50, 0xf0, 0x6e, 0x2b, 0x68, 0xc1, 0xbd, 0x8c, 0x45, 0xc4, 0x8b, 0xa4, 0x99, 0x74,
	0x3a, 0x38, 0x4d, 0x86, 0xcc, 0x5b, 0x11, 0x3f, 0x85, 0x1f, 0xab, 0xc7, 0xbd, 0x08, 0x9e, 0x8a,
	0xae, 0x5f, 0xc4, 0x4c, 0x32, 0xb3, 0x33, 0xeb, 0x94, 0x3d, 0xa4, 0x93, 0xe6, 0xfd, 0xde, 0xff,
	0xe5, 0xfd, 0xf3, 0xd0, 0x19, 0x57, 0xfc, 0x8b, 0x56, 0x8c, 0x6f, 0xfc, 0x0c, 0x94, 0x66, 0x91,
	0x28, 0xbf, 0x2b, 0x96, 0x09, 0x5f, 0x33, 0x19, 0x09, 0x9a, 0x6a, 0x05, 0x0a, 0x9f, 0x5c, 0xa3,
	0xb4, 0x40, 0x68, 0x0d, 0x3d, 0x9a, 0x36, 0x95, 0x84, 0x8c, 0x62, 0x59, 0x7e, 0xd2, 0x95, 0xbf,
	0xfd, 0xca, 0xb9, 0x13, 0x3b, 0x9a, 0x54, 0xb4, 0xfd, 0x35, 0xd1, 0x58, 0x82, 0xd0, 0x92, 0x25,
	0x9f, 0x35, 0x5b, 0x43, 0x81, 0x9d, 0x36, 0xb1, 0xad, 0x00, 0x16, 0x32, 0x60, 0x05, 0x71, 0xdc,
	0x24, 0x6a, 0xd1, 0x71, 0x15, 0xbd, 0x80, 0x38, 0xf1, 0x37, 0x09, 0xf7, 0x21, 0xde, 0x8a, 0x0c,
	0xd8, 0x36, 0x2d, 0x98, 0xc7, 0x91, 0x8a, 0x94, 0xdd, 0xfa, 0xf9, 0xce, 0x9d, 0x8e, 0x7f, 0xde,
	0x45, 0x28, 0xc8, 0xbb, 0x7f, 0x0f, 0x0c, 0x04, 0x7e, 0x86, 0x1e, 0xe4, 0xd7, 0x7a, 0x95, 0xa6,
	0x49, 0x2c, 0xc2, 0x85, 0x0c, 0xc5, 0x37, 0xd2, 0x39, 0x6d, 0x3f, 0xed, 0x04, 0x8d, 0x73, 0x3c,
	0x45, 0x0f, 0x13, 0x61, 0x2c, 0xb9, 0x01, 0x77, 0x2d, 0xdc, 0x0c, 0xe0, 0x97, 0xa8, 0x13, 0x8a,
	0x8c, 0x93, 0x9e, 0x01, 0xbc, 0xe7, 0x63, 0x5a, 0xb9, 0x5c, 0xf4, 0x43, 0xed, 0x35, 0x5e, 0x1b,
	0x46, 0xc7, 0xa9, 0x71, 0x36, 0xb0, 0x3c, 0xa6, 0xa8, 0x6b, 0xc5, 0xc8, 0x3d, 0x9b, 0x48, 0x0e,
	0x24, 0xbe, 0xcb, 0xe3, 0x81, 0xc3, 0xf0, 0x39, 0xba, 0x0f, 0xfa, 0x42, 0x72, 0xd3, 0x4d, 0x68,
	0x7b, 0x22, 0x7d, 0x9b, 0x38, 0x39, 0x58, 0x71, 0x0d, 0xcb, 0x1b, 0x70, 0xf0, 0x5f, 0x32, 0x5e,
	0x22, 0x2f, 0xe2, 0xcb, 0x8d, 0x16, 0xd9, 0x46, 0x25, 0x21, 0x19, 0x58, 0xad, 0x93, 0x9a, 0x56,
	0xee, 0x37, 0x35, 0x7e, 0xd3, 0x65, 0xe9, 0xf7, 0xec, 0xd1, 0x6e, 0x3f, 0x6a, 0x5d, 0xed, 0x47,
	0xde, 0x9b, 0xf9, 0x75, 0x66, 0x50, 0x97, 0xc1, 0x6f, 0x51, 0xd7, 0xa0, 0x90, 0x11, 0x64, 0xf5,
	0xa6, 0xb4, 0x39, 0x73, 0x6e, 0x9a, 0x68, 0x39, 0x54, 0xf4, 0xfc, 0xc3, 0x7c, 0x9e, 0x5f, 0x29,
	0x9b, 0x75, 0x72, 0xf9, 0xc0, 0x09, 0xe0, 0x27, 0xa8, 0xb7, 0xd6, 0xea, 0xbb, 0x90, 0xc4, 0x33,
	0x52, 0xfd, 0xa0, 0xf8, 0x37, 0x06, 0x34, 0xb0, 0x7e, 0x2e, 0xe4, 0x5a, 0xe1, 0x85, 0x2b, 0x27,
	0x48, 0xdb, 0x96, 0x3b, 0xa3, 0xb7, 0x8e, 0x38, 0xad, 0xe6, 0x61, 0xd6, 0xcf, 0x6b, 0x5d, 0xee,
	0x47, 0x6d, 0x57, 0x4f, 0xe0, 0x63, 0x34, 0x48, 0x58, 0x06, 0xee, 0xb1, 0xef, 0xd8, 0xc7, 0xae,
	0x0e, 0x66, 0x93, 0xdd, 0x9f, 0x61, 0x6b, 0x77, 0x35, 0x6c, 0x5f, 0x9a, 0xf5, 0xcb, 0xac, 0xdf,
	0x66, 0xfd, 0xf8, 0x3b, 0x6c, 0x7d, 0xf2, 0x6a, 0x05, 0x3e, 0x76, 0x57, 0x3d, 0x3b, 0x7d, 0x2f,
	0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x24, 0x59, 0x44, 0x32, 0x98, 0x03, 0x00, 0x00,
}