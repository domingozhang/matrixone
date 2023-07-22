// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: status.proto

package status

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SessionField int32

const (
	SessionField_NODE_ID         SessionField = 0
	SessionField_CONN_ID         SessionField = 1
	SessionField_SESSION_ID      SessionField = 2
	SessionField_ACCOUNT         SessionField = 3
	SessionField_USER            SessionField = 4
	SessionField_HOST            SessionField = 5
	SessionField_DB              SessionField = 6
	SessionField_SESSION_START   SessionField = 7
	SessionField_COMMAND         SessionField = 8
	SessionField_INFO            SessionField = 9
	SessionField_TXN_ID          SessionField = 10
	SessionField_STATEMENT_ID    SessionField = 11
	SessionField_STATEMENT_TYPE  SessionField = 12
	SessionField_QUERY_TYPE      SessionField = 13
	SessionField_SQL_SOURCE_TYPE SessionField = 14
	SessionField_QUERY_START     SessionField = 15
)

var SessionField_name = map[int32]string{
	0:  "NODE_ID",
	1:  "CONN_ID",
	2:  "SESSION_ID",
	3:  "ACCOUNT",
	4:  "USER",
	5:  "HOST",
	6:  "DB",
	7:  "SESSION_START",
	8:  "COMMAND",
	9:  "INFO",
	10: "TXN_ID",
	11: "STATEMENT_ID",
	12: "STATEMENT_TYPE",
	13: "QUERY_TYPE",
	14: "SQL_SOURCE_TYPE",
	15: "QUERY_START",
}

var SessionField_value = map[string]int32{
	"NODE_ID":         0,
	"CONN_ID":         1,
	"SESSION_ID":      2,
	"ACCOUNT":         3,
	"USER":            4,
	"HOST":            5,
	"DB":              6,
	"SESSION_START":   7,
	"COMMAND":         8,
	"INFO":            9,
	"TXN_ID":          10,
	"STATEMENT_ID":    11,
	"STATEMENT_TYPE":  12,
	"QUERY_TYPE":      13,
	"SQL_SOURCE_TYPE": 14,
	"QUERY_START":     15,
}

func (x SessionField) String() string {
	return proto.EnumName(SessionField_name, int32(x))
}

func (SessionField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}

// Session is the information of a session.
type Session struct {
	// NodeID is the ID of CN node/service.
	NodeID string `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	// ConnID is the connection ID of backend server.
	ConnID uint32 `protobuf:"varint,2,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	// SessionID is the session ID.
	SessionID string `protobuf:"bytes,3,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	// Account is the tenant name.
	Account string `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
	// User is the username.
	User string `protobuf:"bytes,5,opt,name=User,proto3" json:"User,omitempty"`
	// Host is the host name/address and port.
	Host string `protobuf:"bytes,6,opt,name=Host,proto3" json:"Host,omitempty"`
	// DB is the database name.
	DB string `protobuf:"bytes,7,opt,name=DB,proto3" json:"DB,omitempty"`
	// SessionStart is the start time of this session.
	SessionStart time.Time `protobuf:"bytes,8,opt,name=SessionStart,proto3,stdtime" json:"SessionStart"`
	// Command is the command type.
	Command string `protobuf:"bytes,9,opt,name=Command,proto3" json:"Command,omitempty"`
	// Info is the current SQL statement.
	Info string `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
	// TxnID is the current transaction ID of the session.
	TxnID string `protobuf:"bytes,11,opt,name=TxnID,proto3" json:"TxnID,omitempty"`
	// StatementID is the last statement ID of the session.
	StatementID string `protobuf:"bytes,12,opt,name=StatementID,proto3" json:"StatementID,omitempty"`
	// StatementType is the type of the statement: Insert, Update, Delete, Execute, Select.
	StatementType string `protobuf:"bytes,13,opt,name=StatementType,proto3" json:"StatementType,omitempty"`
	// QueryType is the type of the query: DDL, DML, DQL ...
	QueryType string `protobuf:"bytes,14,opt,name=QueryType,proto3" json:"QueryType,omitempty"`
	// SQLSourceType is the SQL source type: internal_sql, cloud_nonuser_sql, external_sql, cloud_user_sql.
	SQLSourceType string `protobuf:"bytes,15,opt,name=SQLSourceType,proto3" json:"SQLSourceType,omitempty"`
	// QueryStart is the start time of query.
	QueryStart           time.Time `protobuf:"bytes,16,opt,name=QueryStart,proto3,stdtime" json:"QueryStart"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *Session) GetConnID() uint32 {
	if m != nil {
		return m.ConnID
	}
	return 0
}

func (m *Session) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *Session) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Session) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Session) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Session) GetDB() string {
	if m != nil {
		return m.DB
	}
	return ""
}

func (m *Session) GetSessionStart() time.Time {
	if m != nil {
		return m.SessionStart
	}
	return time.Time{}
}

func (m *Session) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Session) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Session) GetTxnID() string {
	if m != nil {
		return m.TxnID
	}
	return ""
}

func (m *Session) GetStatementID() string {
	if m != nil {
		return m.StatementID
	}
	return ""
}

func (m *Session) GetStatementType() string {
	if m != nil {
		return m.StatementType
	}
	return ""
}

func (m *Session) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *Session) GetSQLSourceType() string {
	if m != nil {
		return m.SQLSourceType
	}
	return ""
}

func (m *Session) GetQueryStart() time.Time {
	if m != nil {
		return m.QueryStart
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("status.SessionField", SessionField_name, SessionField_value)
	proto.RegisterType((*Session)(nil), "status.Session")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0x34, 0x75, 0x92, 0x71, 0xfe, 0x98, 0x05, 0xa1, 0x55, 0x85, 0xda, 0x08, 0x71,
	0xa8, 0x90, 0x88, 0x25, 0x38, 0x73, 0x48, 0x6c, 0x57, 0xb5, 0xd4, 0xda, 0xc4, 0xbb, 0x91, 0x28,
	0x97, 0xca, 0x49, 0xb7, 0xc6, 0xa2, 0xf6, 0x46, 0xf6, 0x5a, 0x6a, 0x1f, 0x81, 0x1b, 0x8f, 0xd5,
	0x23, 0x4f, 0x00, 0xa8, 0xaf, 0xc0, 0x0b, 0xa0, 0xdd, 0x75, 0xda, 0xf4, 0xc8, 0x6d, 0xbe, 0xdf,
	0xcc, 0x37, 0x9e, 0x19, 0x6b, 0xa1, 0x5f, 0x89, 0x44, 0xd4, 0xd5, 0x64, 0x5d, 0x72, 0xc1, 0x91,
	0xa9, 0xd5, 0xfe, 0xbb, 0x34, 0x13, 0x5f, 0xeb, 0xe5, 0x64, 0xc5, 0x73, 0x27, 0xe5, 0x29, 0x77,
	0x54, 0x7a, 0x59, 0x5f, 0x29, 0xa5, 0x84, 0x8a, 0xb4, 0x6d, 0xff, 0x30, 0xe5, 0x3c, 0xbd, 0x66,
	0x8f, 0x55, 0x22, 0xcb, 0x59, 0x25, 0x92, 0x7c, 0xad, 0x0b, 0x5e, 0x7f, 0x6f, 0x43, 0x87, 0xb0,
	0xaa, 0xca, 0x78, 0x81, 0x5e, 0x82, 0x19, 0xf2, 0x4b, 0x16, 0x78, 0xd8, 0x18, 0x1b, 0x47, 0xbd,
	0xb8, 0x51, 0x92, 0xbb, 0xbc, 0x28, 0x02, 0x0f, 0xb7, 0xc6, 0xc6, 0xd1, 0x20, 0x6e, 0x14, 0x7a,
	0x05, 0xbd, 0xc6, 0x1a, 0x78, 0x78, 0x57, 0x59, 0x1e, 0x01, 0xc2, 0xd0, 0x99, 0xae, 0x56, 0xbc,
	0x2e, 0x04, 0x6e, 0xab, 0xdc, 0x46, 0x22, 0x04, 0xed, 0x45, 0xc5, 0x4a, 0xbc, 0xa7, 0xb0, 0x8a,
	0x25, 0x3b, 0xe1, 0x95, 0xc0, 0xa6, 0x66, 0x32, 0x46, 0x43, 0x68, 0x79, 0x33, 0xdc, 0x51, 0xa4,
	0xe5, 0xcd, 0xd0, 0x09, 0xf4, 0x9b, 0xf6, 0x44, 0x24, 0xa5, 0xc0, 0xdd, 0xb1, 0x71, 0x64, 0xbd,
	0xdf, 0x9f, 0xe8, 0x1d, 0x27, 0x9b, 0x1d, 0x27, 0x74, 0xb3, 0xe3, 0xac, 0x7b, 0xf7, 0xeb, 0x70,
	0xe7, 0xc7, 0xef, 0x43, 0x23, 0x7e, 0xe2, 0x94, 0xb3, 0xb9, 0x3c, 0xcf, 0x93, 0xe2, 0x12, 0xf7,
	0xf4, 0x6c, 0x8d, 0x94, 0x73, 0x04, 0xc5, 0x15, 0xc7, 0xa0, 0xe7, 0x90, 0x31, 0x7a, 0x01, 0x7b,
	0xf4, 0x46, 0xee, 0x68, 0x29, 0xa8, 0x05, 0x1a, 0x83, 0x45, 0x44, 0x22, 0x58, 0xce, 0x0a, 0x11,
	0x78, 0xb8, 0xaf, 0x72, 0xdb, 0x08, 0xbd, 0x81, 0xc1, 0x83, 0xa4, 0xb7, 0x6b, 0x86, 0x07, 0xaa,
	0xe6, 0x29, 0x94, 0x57, 0x9c, 0xd7, 0xac, 0xbc, 0x55, 0x15, 0x43, 0x7d, 0xc5, 0x07, 0xa0, 0x7a,
	0xcc, 0x4f, 0x09, 0xaf, 0xcb, 0x15, 0x53, 0x15, 0xa3, 0xa6, 0xc7, 0x36, 0x44, 0x1e, 0x80, 0xb2,
	0xe8, 0xbb, 0xd8, 0xff, 0x71, 0x97, 0x2d, 0xdf, 0xdb, 0xbf, 0xc6, 0xc3, 0x81, 0x8f, 0x33, 0x76,
	0x7d, 0x89, 0x2c, 0xe8, 0x84, 0x91, 0xe7, 0x5f, 0x04, 0x9e, 0xbd, 0x23, 0x85, 0x1b, 0x85, 0xa1,
	0x14, 0x06, 0x1a, 0x02, 0x10, 0x9f, 0x90, 0x20, 0x52, 0xba, 0x25, 0x93, 0x53, 0xd7, 0x8d, 0x16,
	0x21, 0xb5, 0x77, 0x51, 0x17, 0xda, 0x0b, 0xe2, 0xc7, 0x76, 0x5b, 0x46, 0x27, 0x11, 0xa1, 0xf6,
	0x1e, 0x32, 0xe5, 0xbf, 0xb4, 0x4d, 0xf4, 0x0c, 0x06, 0x1b, 0x23, 0xa1, 0xd3, 0x98, 0xda, 0x1d,
	0xdd, 0xf8, 0xec, 0x6c, 0x1a, 0x7a, 0x76, 0x57, 0x3a, 0x82, 0xf0, 0x38, 0xb2, 0x7b, 0x08, 0xc0,
	0xa4, 0x9f, 0x55, 0x7b, 0x40, 0x36, 0xf4, 0x09, 0x9d, 0x52, 0xff, 0xcc, 0x0f, 0xa9, 0x24, 0x16,
	0x42, 0x30, 0x7c, 0x24, 0xf4, 0xfc, 0x93, 0x6f, 0xf7, 0xe5, 0x50, 0xf3, 0x85, 0x1f, 0x9f, 0x6b,
	0x3d, 0x40, 0xcf, 0x61, 0x44, 0xe6, 0xa7, 0x17, 0x24, 0x5a, 0xc4, 0xae, 0xaf, 0xe1, 0x10, 0x8d,
	0xc0, 0xd2, 0x45, 0xfa, 0xf3, 0xa3, 0xd9, 0xc7, 0xbb, 0xfb, 0x03, 0xe3, 0xe7, 0xfd, 0x81, 0xf1,
	0xe7, 0xfe, 0xc0, 0xf8, 0xe2, 0x6c, 0xbd, 0xaf, 0x3c, 0x11, 0x65, 0x76, 0xc3, 0xcb, 0x2c, 0xcd,
	0x8a, 0x8d, 0x28, 0x98, 0xb3, 0xfe, 0x96, 0x3a, 0xeb, 0xa5, 0xa3, 0x1f, 0xe4, 0xd2, 0x54, 0xe7,
	0xfd, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xd7, 0xab, 0x78, 0xaf, 0x03, 0x00, 0x00,
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.QueryStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.QueryStart):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStatus(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	if len(m.SQLSourceType) > 0 {
		i -= len(m.SQLSourceType)
		copy(dAtA[i:], m.SQLSourceType)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.SQLSourceType)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.StatementType) > 0 {
		i -= len(m.StatementType)
		copy(dAtA[i:], m.StatementType)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.StatementType)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.StatementID) > 0 {
		i -= len(m.StatementID)
		copy(dAtA[i:], m.StatementID)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.StatementID)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.TxnID) > 0 {
		i -= len(m.TxnID)
		copy(dAtA[i:], m.TxnID)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.TxnID)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0x4a
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SessionStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SessionStart):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintStatus(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	if len(m.DB) > 0 {
		i -= len(m.DB)
		copy(dAtA[i:], m.DB)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.DB)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ConnID != 0 {
		i = encodeVarintStatus(dAtA, i, uint64(m.ConnID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.NodeID) > 0 {
		i -= len(m.NodeID)
		copy(dAtA[i:], m.NodeID)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.NodeID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeID)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	if m.ConnID != 0 {
		n += 1 + sovStatus(uint64(m.ConnID))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.DB)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SessionStart)
	n += 1 + l + sovStatus(uint64(l))
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.TxnID)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.StatementID)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.StatementType)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.SQLSourceType)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.QueryStart)
	n += 2 + l + sovStatus(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnID", wireType)
			}
			m.ConnID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SessionStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxnID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatementID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatementID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatementType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatementType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SQLSourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SQLSourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.QueryStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStatus
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
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
				return 0, ErrInvalidLengthStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStatus = fmt.Errorf("proto: unexpected end of group")
)