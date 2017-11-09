// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api_rpc.proto

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	api_rpc.proto

It has these top-level messages:
	GetNebStateRequest
	GetNebStateResponse
	AccountsRequest
	AccountsResponse
	GetAccountStateRequest
	GetAccountStateResponse
	SendTransactionRequest
	CallRequest
	SendRawTransactionRequest
	SendTransactionResponse
	GetBlockByHashRequest
	GetTransactionByHashRequest
	BlockDumpRequest
	BlockDumpResponse
	TransactionReceiptResponse
*/
package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import corepb "github.com/nebulasio/go-nebulas/core/pb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message of GetNebState rpc.
type GetNebStateRequest struct {
}

func (m *GetNebStateRequest) Reset()                    { *m = GetNebStateRequest{} }
func (m *GetNebStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNebStateRequest) ProtoMessage()               {}
func (*GetNebStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Response message of GetNebState rpc.
type GetNebStateResponse struct {
	// Block chain id
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	// Current neb tail hash
	Tail string `protobuf:"bytes,2,opt,name=tail" json:"tail,omitempty"`
	// Neb coinbase
	Coinbase string `protobuf:"bytes,3,opt,name=coinbase" json:"coinbase,omitempty"`
	// Number of peers currenly connected
	PeerCount uint32 `protobuf:"varint,4,opt,name=peer_count,json=peerCount" json:"peer_count,omitempty"`
	// Neb mine status, minging is true ,otherwise false
	IsMining bool `protobuf:"varint,5,opt,name=is_mining,json=isMining" json:"is_mining,omitempty"`
	// The current neb protocol version.
	ProtocolVersion string `protobuf:"bytes,6,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	// The peer sync status.
	Synchronized bool `protobuf:"varint,7,opt,name=synchronized" json:"synchronized,omitempty"`
}

func (m *GetNebStateResponse) Reset()                    { *m = GetNebStateResponse{} }
func (m *GetNebStateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNebStateResponse) ProtoMessage()               {}
func (*GetNebStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetNebStateResponse) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *GetNebStateResponse) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *GetNebStateResponse) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *GetNebStateResponse) GetPeerCount() uint32 {
	if m != nil {
		return m.PeerCount
	}
	return 0
}

func (m *GetNebStateResponse) GetIsMining() bool {
	if m != nil {
		return m.IsMining
	}
	return false
}

func (m *GetNebStateResponse) GetProtocolVersion() string {
	if m != nil {
		return m.ProtocolVersion
	}
	return ""
}

func (m *GetNebStateResponse) GetSynchronized() bool {
	if m != nil {
		return m.Synchronized
	}
	return false
}

// Request message of Accounts rpc.
type AccountsRequest struct {
}

func (m *AccountsRequest) Reset()                    { *m = AccountsRequest{} }
func (m *AccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountsRequest) ProtoMessage()               {}
func (*AccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message of Accounts rpc.
type AccountsResponse struct {
	// Account list
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *AccountsResponse) Reset()                    { *m = AccountsResponse{} }
func (m *AccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountsResponse) ProtoMessage()               {}
func (*AccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountsResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// Request message of GetAccountState rpc.
type GetAccountStateRequest struct {
	// Hex string of the account addresss.
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Hex string block number, or one of "latest", "earliest" or "pending". If not specified, use "latest".
	Block string `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *GetAccountStateRequest) Reset()                    { *m = GetAccountStateRequest{} }
func (m *GetAccountStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountStateRequest) ProtoMessage()               {}
func (*GetAccountStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetAccountStateRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GetAccountStateRequest) GetBlock() string {
	if m != nil {
		return m.Block
	}
	return ""
}

// Response message of GetAccountState rpc.
type GetAccountStateResponse struct {
	// Current balance in unit of 1/(10^18) nas.
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	// Current transaction count.
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *GetAccountStateResponse) Reset()                    { *m = GetAccountStateResponse{} }
func (m *GetAccountStateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountStateResponse) ProtoMessage()               {}
func (*GetAccountStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetAccountStateResponse) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *GetAccountStateResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

// Request message of SendTransaction rpc.
type SendTransactionRequest struct {
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	// contract source code.
	Source string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	// the params of contract.
	Args string `protobuf:"bytes,6,opt,name=args" json:"args,omitempty"`
}

func (m *SendTransactionRequest) Reset()                    { *m = SendTransactionRequest{} }
func (m *SendTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTransactionRequest) ProtoMessage()               {}
func (*SendTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SendTransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendTransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendTransactionRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SendTransactionRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SendTransactionRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SendTransactionRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

type CallRequest struct {
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	// call contract function name
	Function string `protobuf:"bytes,4,opt,name=function" json:"function,omitempty"`
	// the params of contract.
	Args string `protobuf:"bytes,5,opt,name=args" json:"args,omitempty"`
}

func (m *CallRequest) Reset()                    { *m = CallRequest{} }
func (m *CallRequest) String() string            { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()               {}
func (*CallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CallRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *CallRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CallRequest) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CallRequest) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *CallRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

// Request message of SendRawTransactionRequest rpc.
type SendRawTransactionRequest struct {
	// Signed data of transaction
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SendRawTransactionRequest) Reset()                    { *m = SendRawTransactionRequest{} }
func (m *SendRawTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*SendRawTransactionRequest) ProtoMessage()               {}
func (*SendRawTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SendRawTransactionRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Response message of SendTransaction rpc.
type SendTransactionResponse struct {
	// Hex string of transaction hash.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *SendTransactionResponse) Reset()                    { *m = SendTransactionResponse{} }
func (m *SendTransactionResponse) String() string            { return proto.CompactTextString(m) }
func (*SendTransactionResponse) ProtoMessage()               {}
func (*SendTransactionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SendTransactionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// Request message of GetBlockByHash rpc.
type GetBlockByHashRequest struct {
	// Hex string of block hash.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *GetBlockByHashRequest) Reset()                    { *m = GetBlockByHashRequest{} }
func (m *GetBlockByHashRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockByHashRequest) ProtoMessage()               {}
func (*GetBlockByHashRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetBlockByHashRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// Request message of GetTransactionByHash rpc.
type GetTransactionByHashRequest struct {
	// Hex string of transaction hash.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *GetTransactionByHashRequest) Reset()                    { *m = GetTransactionByHashRequest{} }
func (m *GetTransactionByHashRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTransactionByHashRequest) ProtoMessage()               {}
func (*GetTransactionByHashRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetTransactionByHashRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// Request message of BlockDump.
type BlockDumpRequest struct {
	// the count of blocks to dump before current tail.
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *BlockDumpRequest) Reset()                    { *m = BlockDumpRequest{} }
func (m *BlockDumpRequest) String() string            { return proto.CompactTextString(m) }
func (*BlockDumpRequest) ProtoMessage()               {}
func (*BlockDumpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *BlockDumpRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Response message of BlockDump.
type BlockDumpResponse struct {
	// block dump info.
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *BlockDumpResponse) Reset()                    { *m = BlockDumpResponse{} }
func (m *BlockDumpResponse) String() string            { return proto.CompactTextString(m) }
func (*BlockDumpResponse) ProtoMessage()               {}
func (*BlockDumpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *BlockDumpResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Response message of TransactionReceipt.
type TransactionReceiptResponse struct {
	// Hex string of tx hash.
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	// Transaction nonce.
	Nonce           uint64 `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	Timestamp       int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Data            string `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	ChainId         uint32 `protobuf:"varint,7,opt,name=chainId" json:"chainId,omitempty"`
	ContractAddress string `protobuf:"bytes,8,opt,name=contract_address,json=contractAddress" json:"contract_address,omitempty"`
}

func (m *TransactionReceiptResponse) Reset()                    { *m = TransactionReceiptResponse{} }
func (m *TransactionReceiptResponse) String() string            { return proto.CompactTextString(m) }
func (*TransactionReceiptResponse) ProtoMessage()               {}
func (*TransactionReceiptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TransactionReceiptResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TransactionReceiptResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionReceiptResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransactionReceiptResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransactionReceiptResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TransactionReceiptResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TransactionReceiptResponse) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *TransactionReceiptResponse) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*GetNebStateRequest)(nil), "rpcpb.GetNebStateRequest")
	proto.RegisterType((*GetNebStateResponse)(nil), "rpcpb.GetNebStateResponse")
	proto.RegisterType((*AccountsRequest)(nil), "rpcpb.AccountsRequest")
	proto.RegisterType((*AccountsResponse)(nil), "rpcpb.AccountsResponse")
	proto.RegisterType((*GetAccountStateRequest)(nil), "rpcpb.GetAccountStateRequest")
	proto.RegisterType((*GetAccountStateResponse)(nil), "rpcpb.GetAccountStateResponse")
	proto.RegisterType((*SendTransactionRequest)(nil), "rpcpb.SendTransactionRequest")
	proto.RegisterType((*CallRequest)(nil), "rpcpb.CallRequest")
	proto.RegisterType((*SendRawTransactionRequest)(nil), "rpcpb.SendRawTransactionRequest")
	proto.RegisterType((*SendTransactionResponse)(nil), "rpcpb.SendTransactionResponse")
	proto.RegisterType((*GetBlockByHashRequest)(nil), "rpcpb.GetBlockByHashRequest")
	proto.RegisterType((*GetTransactionByHashRequest)(nil), "rpcpb.GetTransactionByHashRequest")
	proto.RegisterType((*BlockDumpRequest)(nil), "rpcpb.BlockDumpRequest")
	proto.RegisterType((*BlockDumpResponse)(nil), "rpcpb.BlockDumpResponse")
	proto.RegisterType((*TransactionReceiptResponse)(nil), "rpcpb.TransactionReceiptResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for APIService service

type APIServiceClient interface {
	// Return the state of the neb.
	GetNebState(ctx context.Context, in *GetNebStateRequest, opts ...grpc.CallOption) (*GetNebStateResponse, error)
	// Return the dump info of blockchain.
	BlockDump(ctx context.Context, in *BlockDumpRequest, opts ...grpc.CallOption) (*BlockDumpResponse, error)
	// Accounts return account list.
	Accounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error)
	// Return the state of the account.
	GetAccountState(ctx context.Context, in *GetAccountStateRequest, opts ...grpc.CallOption) (*GetAccountStateResponse, error)
	// Verify, sign, and send the transaction.
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
	// Call smart contract.
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
	// Submit the signed transaction.
	SendRawTransaction(ctx context.Context, in *SendRawTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
	// Get block header info by the block hash.
	GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*corepb.Block, error)
	// Get transactionReceipt info by tansaction hash.
	GetTransactionReceipt(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*TransactionReceiptResponse, error)
}

type aPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAPIServiceClient(cc *grpc.ClientConn) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) GetNebState(ctx context.Context, in *GetNebStateRequest, opts ...grpc.CallOption) (*GetNebStateResponse, error) {
	out := new(GetNebStateResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/GetNebState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) BlockDump(ctx context.Context, in *BlockDumpRequest, opts ...grpc.CallOption) (*BlockDumpResponse, error) {
	out := new(BlockDumpResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/BlockDump", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) Accounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error) {
	out := new(AccountsResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/Accounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetAccountState(ctx context.Context, in *GetAccountStateRequest, opts ...grpc.CallOption) (*GetAccountStateResponse, error) {
	out := new(GetAccountStateResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/GetAccountState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/SendTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) SendRawTransaction(ctx context.Context, in *SendRawTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/SendRawTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*corepb.Block, error) {
	out := new(corepb.Block)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/GetBlockByHash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetTransactionReceipt(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*TransactionReceiptResponse, error) {
	out := new(TransactionReceiptResponse)
	err := grpc.Invoke(ctx, "/rpcpb.APIService/GetTransactionReceipt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for APIService service

type APIServiceServer interface {
	// Return the state of the neb.
	GetNebState(context.Context, *GetNebStateRequest) (*GetNebStateResponse, error)
	// Return the dump info of blockchain.
	BlockDump(context.Context, *BlockDumpRequest) (*BlockDumpResponse, error)
	// Accounts return account list.
	Accounts(context.Context, *AccountsRequest) (*AccountsResponse, error)
	// Return the state of the account.
	GetAccountState(context.Context, *GetAccountStateRequest) (*GetAccountStateResponse, error)
	// Verify, sign, and send the transaction.
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error)
	// Call smart contract.
	Call(context.Context, *CallRequest) (*SendTransactionResponse, error)
	// Submit the signed transaction.
	SendRawTransaction(context.Context, *SendRawTransactionRequest) (*SendTransactionResponse, error)
	// Get block header info by the block hash.
	GetBlockByHash(context.Context, *GetBlockByHashRequest) (*corepb.Block, error)
	// Get transactionReceipt info by tansaction hash.
	GetTransactionReceipt(context.Context, *GetTransactionByHashRequest) (*TransactionReceiptResponse, error)
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_GetNebState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNebStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetNebState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/GetNebState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetNebState(ctx, req.(*GetNebStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_BlockDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockDumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).BlockDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/BlockDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).BlockDump(ctx, req.(*BlockDumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_Accounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).Accounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/Accounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).Accounts(ctx, req.(*AccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetAccountState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetAccountState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/GetAccountState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetAccountState(ctx, req.(*GetAccountStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_SendRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).SendRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/SendRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).SendRawTransaction(ctx, req.(*SendRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetBlockByHash(ctx, req.(*GetBlockByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetTransactionReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetTransactionReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.APIService/GetTransactionReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetTransactionReceipt(ctx, req.(*GetTransactionByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNebState",
			Handler:    _APIService_GetNebState_Handler,
		},
		{
			MethodName: "BlockDump",
			Handler:    _APIService_BlockDump_Handler,
		},
		{
			MethodName: "Accounts",
			Handler:    _APIService_Accounts_Handler,
		},
		{
			MethodName: "GetAccountState",
			Handler:    _APIService_GetAccountState_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _APIService_SendTransaction_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _APIService_Call_Handler,
		},
		{
			MethodName: "SendRawTransaction",
			Handler:    _APIService_SendRawTransaction_Handler,
		},
		{
			MethodName: "GetBlockByHash",
			Handler:    _APIService_GetBlockByHash_Handler,
		},
		{
			MethodName: "GetTransactionReceipt",
			Handler:    _APIService_GetTransactionReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_rpc.proto",
}

func init() { proto.RegisterFile("api_rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0x0d, 0xf5, 0x67, 0xf1, 0xc6, 0xb2, 0xec, 0xb1, 0x22, 0x53, 0x8c, 0x12, 0xe8, 0x1b, 0x7c,
	0x40, 0x95, 0x14, 0x11, 0x9b, 0x66, 0x97, 0x4d, 0xe1, 0xa4, 0x80, 0xe3, 0x45, 0xff, 0xe8, 0xa2,
	0x68, 0xd1, 0x85, 0x30, 0x1a, 0x4d, 0x24, 0xa2, 0xd4, 0x0c, 0xcb, 0x19, 0x39, 0x48, 0x97, 0x45,
	0xdf, 0xa0, 0x7d, 0xb3, 0x3e, 0x41, 0x81, 0x2e, 0x0a, 0xf4, 0x25, 0x8a, 0x19, 0x0e, 0x7f, 0x64,
	0xca, 0x35, 0xba, 0x9b, 0x7b, 0x79, 0x79, 0xce, 0xbd, 0xe7, 0x72, 0x0e, 0xa1, 0x47, 0x92, 0x68,
	0x9e, 0x26, 0x74, 0x96, 0xa4, 0x42, 0x09, 0xd4, 0x4e, 0x13, 0x9a, 0x2c, 0xfc, 0xf1, 0x4a, 0x88,
	0x55, 0xcc, 0x02, 0x92, 0x44, 0x01, 0xe1, 0x5c, 0x28, 0xa2, 0x22, 0xc1, 0x65, 0x56, 0xe4, 0xbf,
	0x58, 0x45, 0x6a, 0xbd, 0x5d, 0xcc, 0xa8, 0xd8, 0x04, 0x9c, 0x2d, 0xb6, 0x31, 0x91, 0x91, 0x08,
	0x56, 0xe2, 0x99, 0x0d, 0x02, 0x2a, 0x52, 0x16, 0x24, 0x8b, 0x60, 0x11, 0x0b, 0xfa, 0x43, 0xf6,
	0x12, 0x1e, 0x00, 0xba, 0x60, 0xea, 0x73, 0xb6, 0xb8, 0x52, 0x44, 0xb1, 0x90, 0xfd, 0xb8, 0x65,
	0x52, 0xe1, 0xbf, 0x1d, 0x38, 0xdd, 0x49, 0xcb, 0x44, 0x70, 0xc9, 0xd0, 0x08, 0xba, 0x74, 0x4d,
	0x22, 0x3e, 0x8f, 0x96, 0x9e, 0x33, 0x71, 0xa6, 0xbd, 0xf0, 0xc0, 0xc4, 0x97, 0x4b, 0x84, 0xa0,
	0xa5, 0x48, 0x14, 0x7b, 0x8d, 0x89, 0x33, 0x75, 0x43, 0x73, 0x46, 0x3e, 0x74, 0xa9, 0x88, 0xf8,
	0x82, 0x48, 0xe6, 0x35, 0x4d, 0xbe, 0x88, 0xd1, 0x23, 0x80, 0x84, 0xb1, 0x74, 0x4e, 0xc5, 0x96,
	0x2b, 0xaf, 0x65, 0xc0, 0x5c, 0x9d, 0x79, 0xad, 0x13, 0xe8, 0x21, 0xb8, 0x91, 0x9c, 0x6f, 0x22,
	0x1e, 0xf1, 0x95, 0xd7, 0x9e, 0x38, 0xd3, 0x6e, 0xd8, 0x8d, 0xe4, 0x67, 0x26, 0x46, 0x4f, 0xe0,
	0xd8, 0x74, 0x4f, 0x45, 0x3c, 0xbf, 0x66, 0xa9, 0x8c, 0x04, 0xf7, 0x3a, 0x06, 0xbf, 0x9f, 0xe7,
	0xbf, 0xc9, 0xd2, 0x08, 0xc3, 0xa1, 0x7c, 0xcf, 0xe9, 0x3a, 0x15, 0x3c, 0xfa, 0x89, 0x2d, 0xbd,
	0x03, 0x03, 0xb5, 0x93, 0xc3, 0x27, 0xd0, 0x3f, 0xa7, 0xa6, 0x0f, 0x99, 0x0b, 0xf0, 0x11, 0x1c,
	0x97, 0x29, 0x3b, 0xfc, 0x18, 0x5c, 0xb2, 0x5c, 0xa6, 0x4c, 0x4a, 0x26, 0x3d, 0x67, 0xd2, 0x9c,
	0xba, 0x61, 0x99, 0xc0, 0x6f, 0x60, 0x78, 0xc1, 0x94, 0x7d, 0xa9, 0x2a, 0x26, 0xf2, 0xe0, 0xc0,
	0x96, 0x19, 0xcd, 0xdc, 0x30, 0x0f, 0xd1, 0x00, 0xda, 0x66, 0x17, 0x56, 0xb4, 0x2c, 0xc0, 0x97,
	0x70, 0x56, 0x43, 0xb2, 0x2d, 0x78, 0x70, 0xb0, 0x20, 0x31, 0xe1, 0x94, 0x19, 0xa8, 0xc3, 0x30,
	0x0f, 0x35, 0x14, 0x17, 0x3a, 0xaf, 0xa1, 0x5a, 0x61, 0x16, 0xe0, 0xdf, 0x1c, 0x18, 0x5e, 0x31,
	0xbe, 0xfc, 0x3a, 0x25, 0x5c, 0x12, 0xaa, 0xbf, 0x96, 0xbc, 0x2b, 0x04, 0xad, 0xb7, 0xa9, 0xd8,
	0xd8, 0x96, 0xcc, 0x19, 0x1d, 0x41, 0x43, 0x09, 0xdb, 0x4c, 0x43, 0x09, 0x0d, 0x7a, 0x4d, 0xe2,
	0x6d, 0xb6, 0xbc, 0xc3, 0x30, 0x0b, 0x4a, 0xaa, 0x56, 0x85, 0x0a, 0x0d, 0xa1, 0x23, 0xc5, 0x36,
	0xa5, 0xcc, 0x6c, 0xcb, 0x0d, 0x6d, 0xa4, 0x79, 0x48, 0xba, 0x92, 0x76, 0x3f, 0xe6, 0x8c, 0xdf,
	0xc1, 0xfd, 0xd7, 0x24, 0x8e, 0xff, 0x63, 0x2b, 0x19, 0x69, 0xb3, 0x4a, 0xea, 0x43, 0xf7, 0xed,
	0x96, 0x9b, 0xb9, 0x4c, 0x37, 0x6e, 0x58, 0xc4, 0x05, 0x71, 0xbb, 0x42, 0x1c, 0xc0, 0x48, 0xcb,
	0x11, 0x92, 0x77, 0xfb, 0x15, 0x59, 0x12, 0x45, 0xac, 0xb2, 0xe6, 0x8c, 0x9f, 0xc1, 0x59, 0x4d,
	0x3f, 0xbb, 0x0b, 0x04, 0xad, 0x35, 0x91, 0xeb, 0xbc, 0x6b, 0x7d, 0xc6, 0x1f, 0xc2, 0x83, 0x0b,
	0xa6, 0x5e, 0xe9, 0x35, 0xbe, 0x7a, 0xff, 0x86, 0xc8, 0x75, 0x05, 0xbb, 0x56, 0xfc, 0x1c, 0x1e,
	0x5e, 0x30, 0x55, 0x81, 0xbe, 0xfb, 0x95, 0x29, 0x1c, 0x1b, 0xf0, 0x4f, 0xb7, 0x9b, 0x24, 0xaf,
	0x1b, 0x40, 0x3b, 0xbb, 0x43, 0xba, 0xb0, 0x1d, 0x66, 0x01, 0xfe, 0x00, 0x4e, 0x2a, 0x95, 0x65,
	0xcb, 0xc5, 0x84, 0xae, 0x9d, 0xf0, 0x0f, 0x07, 0xfc, 0x9d, 0xf1, 0x28, 0x8b, 0x12, 0xf5, 0x6f,
	0x53, 0x16, 0xfb, 0x6a, 0xd4, 0xf6, 0xd5, 0xac, 0xef, 0x6b, 0xe7, 0x23, 0x19, 0x83, 0xab, 0xa2,
	0x0d, 0x93, 0x8a, 0x6c, 0x12, 0xb3, 0x98, 0x66, 0x58, 0x26, 0x8a, 0xf6, 0x3a, 0x65, 0x7b, 0xfa,
	0x8b, 0xb7, 0x0e, 0x63, 0xae, 0x6e, 0xc5, 0x70, 0x9e, 0xc0, 0x31, 0x15, 0x5c, 0xa5, 0x84, 0xaa,
	0x79, 0x7e, 0xbf, 0xba, 0x99, 0x09, 0xe4, 0xf9, 0xf3, 0x2c, 0xfd, 0xf1, 0x5f, 0x1d, 0x80, 0xf3,
	0x2f, 0x2f, 0xaf, 0x58, 0x7a, 0x1d, 0x51, 0x86, 0xbe, 0x87, 0xfb, 0x15, 0x73, 0x43, 0xa3, 0x99,
	0x71, 0xd7, 0x59, 0xdd, 0x07, 0x7d, 0x7f, 0xdf, 0xa3, 0x4c, 0x19, 0xfc, 0xe0, 0xe7, 0xdf, 0xff,
	0xfc, 0xb5, 0xd1, 0x47, 0xbd, 0xe0, 0xfa, 0xb9, 0xf6, 0xdb, 0x40, 0x1a, 0xb4, 0xef, 0xc0, 0x2d,
	0x84, 0x47, 0x67, 0xf6, 0xfd, 0x9b, 0x4b, 0xf3, 0xbd, 0xfa, 0x03, 0x0b, 0x3b, 0x32, 0xb0, 0xa7,
	0xf8, 0x48, 0xc3, 0x1a, 0x43, 0x08, 0x96, 0xdb, 0x4d, 0xf2, 0xd2, 0x79, 0x8a, 0xbe, 0x82, 0x6e,
	0x6e, 0x4a, 0x68, 0x68, 0x01, 0x6e, 0x18, 0x97, 0x7f, 0x56, 0xcb, 0x5b, 0xdc, 0x81, 0xc1, 0x3d,
	0x42, 0x87, 0x1a, 0x97, 0xe4, 0x30, 0x1c, 0xfa, 0x37, 0xbc, 0x06, 0x3d, 0x2a, 0x67, 0xde, 0xe3,
	0x66, 0xfe, 0xe3, 0xdb, 0x1e, 0x5b, 0x9e, 0xb1, 0xe1, 0x19, 0xe2, 0x93, 0x0a, 0x4f, 0x26, 0x8d,
	0x1e, 0x21, 0x86, 0xfe, 0x8d, 0xfb, 0x54, 0xf0, 0xed, 0xf7, 0xa9, 0x82, 0xef, 0x96, 0x6b, 0x88,
	0x7d, 0xc3, 0x37, 0xc0, 0x7d, 0xcd, 0xa7, 0xca, 0x02, 0xcd, 0xf6, 0x05, 0xb4, 0xb4, 0xcf, 0x20,
	0x64, 0x31, 0x2a, 0xa6, 0x73, 0x27, 0xee, 0xa9, 0xc1, 0xed, 0xe1, 0xae, 0xc6, 0xa5, 0x24, 0x8e,
	0x35, 0xe0, 0xb7, 0x80, 0xea, 0xfe, 0x81, 0x26, 0x15, 0xa8, 0xbd, 0xd6, 0x72, 0x27, 0xd9, 0x3d,
	0xf4, 0x09, 0x1c, 0xed, 0x3a, 0x07, 0x1a, 0x97, 0x42, 0xd7, 0x0d, 0xc5, 0xef, 0xcd, 0xf4, 0xdf,
	0x3c, 0xff, 0x82, 0xf0, 0x3d, 0xf4, 0x8b, 0x63, 0xbc, 0xa7, 0x7e, 0x95, 0x11, 0x2e, 0x81, 0x6e,
	0x33, 0x1b, 0xff, 0x7f, 0xb6, 0xe6, 0x76, 0x27, 0xc0, 0xff, 0x37, 0x82, 0x3c, 0xc6, 0x23, 0x2d,
	0xc8, 0x6a, 0x1f, 0xd3, 0x4b, 0xe7, 0xe9, 0xa2, 0x63, 0x7e, 0xc0, 0x2f, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x99, 0x92, 0xa6, 0xc1, 0x08, 0x00, 0x00,
}
