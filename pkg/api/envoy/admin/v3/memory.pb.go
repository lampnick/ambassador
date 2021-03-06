// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v3/memory.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Proto representation of the internal memory consumption of an Envoy instance. These represent
// values extracted from an internal TCMalloc instance. For more information, see the section of the
// docs entitled ["Generic Tcmalloc Status"](https://gperftools.github.io/gperftools/tcmalloc.html).
// [#next-free-field: 6]
type Memory struct {
	// The number of bytes allocated by the heap for Envoy. This is an alias for
	// `generic.current_allocated_bytes`.
	Allocated uint64 `protobuf:"varint,1,opt,name=allocated,proto3" json:"allocated,omitempty"`
	// The number of bytes reserved by the heap but not necessarily allocated. This is an alias for
	// `generic.heap_size`.
	HeapSize uint64 `protobuf:"varint,2,opt,name=heap_size,json=heapSize,proto3" json:"heap_size,omitempty"`
	// The number of bytes in free, unmapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and depending on the OS, typically do not count towards physical memory
	// usage. This is an alias for `tcmalloc.pageheap_unmapped_bytes`.
	PageheapUnmapped uint64 `protobuf:"varint,3,opt,name=pageheap_unmapped,json=pageheapUnmapped,proto3" json:"pageheap_unmapped,omitempty"`
	// The number of bytes in free, mapped pages in the page heap. These bytes always count towards
	// virtual memory usage, and unless the underlying memory is swapped out by the OS, they also
	// count towards physical memory usage. This is an alias for `tcmalloc.pageheap_free_bytes`.
	PageheapFree uint64 `protobuf:"varint,4,opt,name=pageheap_free,json=pageheapFree,proto3" json:"pageheap_free,omitempty"`
	// The amount of memory used by the TCMalloc thread caches (for small objects). This is an alias
	// for `tcmalloc.current_total_thread_cache_bytes`.
	TotalThreadCache     uint64   `protobuf:"varint,5,opt,name=total_thread_cache,json=totalThreadCache,proto3" json:"total_thread_cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ea6f9dbe7a0a918, []int{0}
}
func (m *Memory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return m.Size()
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetAllocated() uint64 {
	if m != nil {
		return m.Allocated
	}
	return 0
}

func (m *Memory) GetHeapSize() uint64 {
	if m != nil {
		return m.HeapSize
	}
	return 0
}

func (m *Memory) GetPageheapUnmapped() uint64 {
	if m != nil {
		return m.PageheapUnmapped
	}
	return 0
}

func (m *Memory) GetPageheapFree() uint64 {
	if m != nil {
		return m.PageheapFree
	}
	return 0
}

func (m *Memory) GetTotalThreadCache() uint64 {
	if m != nil {
		return m.TotalThreadCache
	}
	return 0
}

func init() {
	proto.RegisterType((*Memory)(nil), "envoy.admin.v3.Memory")
}

func init() { proto.RegisterFile("envoy/admin/v3/memory.proto", fileDescriptor_1ea6f9dbe7a0a918) }

var fileDescriptor_1ea6f9dbe7a0a918 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x59, 0xad, 0xc5, 0xae, 0x3f, 0xe8, 0x9e, 0x42, 0x5b, 0x82, 0xd5, 0x8b, 0xa0, 0x6c,
	0xc0, 0x9e, 0xf4, 0x58, 0xc1, 0x9b, 0x50, 0xfc, 0x39, 0x97, 0x31, 0x19, 0x9b, 0x85, 0x64, 0x77,
	0xd9, 0x6c, 0x83, 0xe9, 0x13, 0xf8, 0x0c, 0xbe, 0x8f, 0xe0, 0xd1, 0x47, 0x90, 0x1c, 0x7c, 0x0e,
	0xc9, 0xc4, 0x2a, 0xbd, 0x7e, 0xdf, 0x37, 0x03, 0x33, 0x7c, 0x80, 0xba, 0x34, 0x55, 0x04, 0x49,
	0xae, 0x74, 0x54, 0x8e, 0xa3, 0x1c, 0x73, 0xe3, 0x2a, 0x69, 0x9d, 0xf1, 0x46, 0xec, 0x93, 0x94,
	0x24, 0x65, 0x39, 0xee, 0x8f, 0x16, 0x89, 0x85, 0x08, 0xb4, 0x36, 0x1e, 0xbc, 0x32, 0xba, 0x88,
	0x4a, 0x74, 0x85, 0x32, 0x5a, 0xe9, 0x79, 0x3b, 0x72, 0xfc, 0xcd, 0x78, 0xf7, 0x96, 0x76, 0x88,
	0x21, 0xef, 0x41, 0x96, 0x99, 0x18, 0x3c, 0x26, 0x01, 0x3b, 0x62, 0xa7, 0x9d, 0xbb, 0x7f, 0x20,
	0x06, 0xbc, 0x97, 0x22, 0xd8, 0x59, 0xa1, 0x96, 0x18, 0x6c, 0x90, 0xdd, 0x6e, 0xc0, 0xbd, 0x5a,
	0xa2, 0x38, 0xe3, 0x87, 0x16, 0xe6, 0x48, 0xc1, 0x42, 0xe7, 0x60, 0x2d, 0x26, 0xc1, 0x26, 0x45,
	0x07, 0x2b, 0xf1, 0xf8, 0xcb, 0xc5, 0x09, 0xdf, 0xfb, 0x8b, 0x9f, 0x1d, 0x62, 0xd0, 0xa1, 0x70,
	0x77, 0x05, 0x6f, 0x1c, 0xa2, 0x38, 0xe7, 0xc2, 0x1b, 0x0f, 0xd9, 0xcc, 0xa7, 0x0e, 0x21, 0x99,
	0xc5, 0x10, 0xa7, 0x18, 0x6c, 0xb5, 0x2b, 0xc9, 0x3c, 0x90, 0xb8, 0x6e, 0xf8, 0xd5, 0xe8, 0xed,
	0xfd, 0x35, 0x1c, 0xf2, 0xfe, 0xda, 0xfd, 0x17, 0x90, 0xd9, 0x14, 0x64, 0x7b, 0xdd, 0xe4, 0xf2,
	0xa3, 0x0e, 0xd9, 0x67, 0x1d, 0xb2, 0xaf, 0x3a, 0x64, 0x7c, 0xa8, 0x8c, 0xa4, 0xd8, 0x3a, 0xf3,
	0x52, 0xc9, 0xf5, 0xbf, 0x4d, 0x76, 0xda, 0x99, 0x69, 0xf3, 0xa1, 0x29, 0x7b, 0xea, 0xd2, 0xab,
	0xc6, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x76, 0x10, 0x7b, 0x7c, 0x01, 0x00, 0x00,
}

func (m *Memory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Memory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Memory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TotalThreadCache != 0 {
		i = encodeVarintMemory(dAtA, i, uint64(m.TotalThreadCache))
		i--
		dAtA[i] = 0x28
	}
	if m.PageheapFree != 0 {
		i = encodeVarintMemory(dAtA, i, uint64(m.PageheapFree))
		i--
		dAtA[i] = 0x20
	}
	if m.PageheapUnmapped != 0 {
		i = encodeVarintMemory(dAtA, i, uint64(m.PageheapUnmapped))
		i--
		dAtA[i] = 0x18
	}
	if m.HeapSize != 0 {
		i = encodeVarintMemory(dAtA, i, uint64(m.HeapSize))
		i--
		dAtA[i] = 0x10
	}
	if m.Allocated != 0 {
		i = encodeVarintMemory(dAtA, i, uint64(m.Allocated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMemory(dAtA []byte, offset int, v uint64) int {
	offset -= sovMemory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Memory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allocated != 0 {
		n += 1 + sovMemory(uint64(m.Allocated))
	}
	if m.HeapSize != 0 {
		n += 1 + sovMemory(uint64(m.HeapSize))
	}
	if m.PageheapUnmapped != 0 {
		n += 1 + sovMemory(uint64(m.PageheapUnmapped))
	}
	if m.PageheapFree != 0 {
		n += 1 + sovMemory(uint64(m.PageheapFree))
	}
	if m.TotalThreadCache != 0 {
		n += 1 + sovMemory(uint64(m.TotalThreadCache))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMemory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMemory(x uint64) (n int) {
	return sovMemory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Memory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMemory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Memory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Memory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocated", wireType)
			}
			m.Allocated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Allocated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeapSize", wireType)
			}
			m.HeapSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeapSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageheapUnmapped", wireType)
			}
			m.PageheapUnmapped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageheapUnmapped |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageheapFree", wireType)
			}
			m.PageheapFree = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageheapFree |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalThreadCache", wireType)
			}
			m.TotalThreadCache = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalThreadCache |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMemory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMemory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMemory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMemory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMemory
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
					return 0, ErrIntOverflowMemory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMemory
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
			if length < 0 {
				return 0, ErrInvalidLengthMemory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMemory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMemory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMemory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMemory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMemory = fmt.Errorf("proto: unexpected end of group")
)
