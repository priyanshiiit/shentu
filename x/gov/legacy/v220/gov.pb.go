// Package v220 is taken from:
// https://github.com/certikfoundation/shentu/blob/v2.2.0/x/gov/types/gov.pb.go
// nolint

package v220

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// ProposalStatus enumerates the valid statuses of a proposal.
type ProposalStatus int32

const (
	// PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status.
	StatusNil ProposalStatus = 0
	// PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
	// period.
	StatusDepositPeriod ProposalStatus = 1
	// PROPOSAL_STATUS_VOTING_PERIOD defines a certifier voting period status.
	StatusCertifierVotingPeriod ProposalStatus = 2
	// PROPOSAL_STATUS_VOTING_PERIOD defines a validator voting period status.
	StatusValidatorVotingPeriod ProposalStatus = 3
	// PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
	// passed.
	StatusPassed ProposalStatus = 4
	// PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
	// been rejected.
	StatusRejected ProposalStatus = 5
	// PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
	// failed.
	StatusFailed ProposalStatus = 6
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_UNSPECIFIED",
	1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	2: "PROPOSAL_STATUS_CERTIFIER_VOTING_PERIOD",
	3: "PROPOSAL_STATUS_VALIDATOR_VOTING_PERIOD",
	4: "PROPOSAL_STATUS_PASSED",
	5: "PROPOSAL_STATUS_REJECTED",
	6: "PROPOSAL_STATUS_FAILED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_UNSPECIFIED":             0,
	"PROPOSAL_STATUS_DEPOSIT_PERIOD":          1,
	"PROPOSAL_STATUS_CERTIFIER_VOTING_PERIOD": 2,
	"PROPOSAL_STATUS_VALIDATOR_VOTING_PERIOD": 3,
	"PROPOSAL_STATUS_PASSED":                  4,
	"PROPOSAL_STATUS_REJECTED":                5,
	"PROPOSAL_STATUS_FAILED":                  6,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{0}
}

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty" yaml:"starting_proposal_id"`
	// deposits defines all the deposits present at genesis.
	Deposits Deposits `protobuf:"bytes,2,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	// votes defines all the votes present at genesis.
	Votes Votes `protobuf:"bytes,3,rep,name=votes,proto3,castrepeated=Votes" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals Proposals `protobuf:"bytes,4,rep,name=proposals,proto3,castrepeated=Proposals" json:"proposals"`
	// params defines all the parameters of related to deposit.
	DepositParams DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params" yaml:"deposit_params"`
	// params defines all the parameters of related to voting.
	VotingParams types.VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params" yaml:"voting_params"`
	// params defines all the parameters of related to tally.
	TallyParams TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params" yaml:"tally_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// Deposit defines an amount deposited by an account address to an active
// proposal.
type Deposit struct {
	*types.Deposit `protobuf:"bytes,1,opt,name=deposit,proto3,embedded=deposit" json:"deposit,omitempty"`
	TxHash         string `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"txhash"`
}

func (m *Deposit) Reset()      { *m = Deposit{} }
func (*Deposit) ProtoMessage() {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{1}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

// DepositParams defines the params for deposits on governance proposals.
type DepositParams struct {
	//  Minimum deposit for a proposal to enter voting period.
	MinInitialDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=min_initial_deposit,json=minInitialDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_initial_deposit,omitempty" yaml:"min_initial_deposit"`
	// Minimum deposit for a proposal to enter voting period.
	MinDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=min_deposit,json=minDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_deposit,omitempty" yaml:"min_deposit"`
	//  Maximum period for CTK holders to deposit on a proposal. Initial value: 2
	//  months.
	MaxDepositPeriod time.Duration `protobuf:"bytes,3,opt,name=max_deposit_period,json=maxDepositPeriod,proto3,stdduration" json:"max_deposit_period,omitempty" yaml:"max_deposit_period"`
}

func (m *DepositParams) Reset()      { *m = DepositParams{} }
func (*DepositParams) ProtoMessage() {}
func (*DepositParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{2}
}
func (m *DepositParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositParams.Merge(m, src)
}
func (m *DepositParams) XXX_Size() int {
	return m.Size()
}
func (m *DepositParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositParams.DiscardUnknown(m)
}

var xxx_messageInfo_DepositParams proto.InternalMessageInfo

// TallyParams defines the params for tallying votes on governance proposals.
type TallyParams struct {
	DefaultTally                     *types.TallyParams `protobuf:"bytes,1,opt,name=default_tally,json=defaultTally,proto3" json:"default_tally,omitempty"`
	CertifierUpdateSecurityVoteTally *types.TallyParams `protobuf:"bytes,2,opt,name=certifier_update_security_vote_tally,json=certifierUpdateSecurityVoteTally,proto3" json:"certifier_update_security_vote_tally,omitempty"`
	CertifierUpdateStakeVoteTally    *types.TallyParams `protobuf:"bytes,3,opt,name=certifier_update_stake_vote_tally,json=certifierUpdateStakeVoteTally,proto3" json:"certifier_update_stake_vote_tally,omitempty"`
}

func (m *TallyParams) Reset()      { *m = TallyParams{} }
func (*TallyParams) ProtoMessage() {}
func (*TallyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{3}
}
func (m *TallyParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TallyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TallyParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TallyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyParams.Merge(m, src)
}
func (m *TallyParams) XXX_Size() int {
	return m.Size()
}
func (m *TallyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyParams.DiscardUnknown(m)
}

var xxx_messageInfo_TallyParams proto.InternalMessageInfo

type Proposal struct {
	Content                 *types2.Any                              `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ProposalId              uint64                                   `protobuf:"varint,2,opt,name=proposal_id,json=proposalId,proto3" json:"id" yaml:"id"`
	Status                  ProposalStatus                           `protobuf:"varint,3,opt,name=status,proto3,enum=shentu.gov.v1alpha1.ProposalStatus" json:"status,omitempty" yaml:"proposal_status"`
	IsProposerCouncilMember bool                                     `protobuf:"varint,4,opt,name=is_proposer_council_member,json=isProposerCouncilMember,proto3" json:"is_proposer_council_member,omitempty" yaml:"is_proposer_council_member"`
	ProposerAddress         string                                   `protobuf:"bytes,5,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty" yaml:"proposer_address"`
	FinalTallyResult        types.TallyResult                        `protobuf:"bytes,6,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result" yaml:"final_tally_result"`
	SubmitTime              time.Time                                `protobuf:"bytes,7,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time" yaml:"submit_time"`
	DepositEndTime          time.Time                                `protobuf:"bytes,8,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time" yaml:"deposit_end_time"`
	TotalDeposit            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=total_deposit,json=totalDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_deposit" yaml:"total_deposit"`
	VotingStartTime         time.Time                                `protobuf:"bytes,10,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time" yaml:"voting_start_time"`
	VotingEndTime           time.Time                                `protobuf:"bytes,11,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time" yaml:"voting_end_time"`
}

func (m *Proposal) Reset()      { *m = Proposal{} }
func (*Proposal) ProtoMessage() {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{4}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

// Vote defines a vote on a governance proposal.
// A Vote consists of a proposal ID, the voter, and the vote option.
type Vote struct {
	*types.Vote `protobuf:"bytes,1,opt,name=deposit,proto3,embedded=deposit" json:"deposit,omitempty"`
	TxHash      string `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"txhash"`
}

func (m *Vote) Reset()      { *m = Vote{} }
func (*Vote) ProtoMessage() {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce5b0b452eb2673, []int{5}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func init() {
	// proto.RegisterEnum("shentu.gov.v1alpha1.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	// proto.RegisterType((*GenesisState)(nil), "shentu.gov.v1alpha1.GenesisState")
	// proto.RegisterType((*Deposit)(nil), "shentu.gov.v1alpha1.Deposit")
	// proto.RegisterType((*DepositParams)(nil), "shentu.gov.v1alpha1.DepositParams")
	// proto.RegisterType((*TallyParams)(nil), "shentu.gov.v1alpha1.TallyParams")
	// proto.RegisterType((*Proposal)(nil), "shentu.gov.v1alpha1.Proposal")
	// proto.RegisterType((*Vote)(nil), "shentu.gov.v1alpha1.Vote")
}

func init() { proto.RegisterFile("shentu/gov/v1alpha1/gov.proto", fileDescriptor_3ce5b0b452eb2673) }

var fileDescriptor_3ce5b0b452eb2673 = []byte{
	// 1472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xcd, 0x4f, 0x1b, 0x47,
	0x1b, 0xf7, 0x1a, 0xf3, 0x35, 0xc6, 0xc4, 0x0c, 0x24, 0x18, 0x03, 0x5e, 0x67, 0x93, 0x57, 0x2f,
	0x8a, 0x12, 0xfb, 0x85, 0xf7, 0x3d, 0xbc, 0xa2, 0x52, 0x25, 0x2f, 0x36, 0x89, 0x23, 0x0a, 0xee,
	0xda, 0xa1, 0x52, 0x7b, 0xd8, 0x8e, 0xbd, 0x03, 0x4c, 0xb3, 0x1f, 0x96, 0x67, 0x8c, 0x40, 0xea,
	0xa1, 0xc7, 0x88, 0x43, 0x95, 0x63, 0xa4, 0x0a, 0x29, 0x52, 0x6e, 0x91, 0xaa, 0x5e, 0x7a, 0xea,
	0x5f, 0x10, 0xf5, 0x94, 0x63, 0x4e, 0x4e, 0x43, 0x2e, 0x51, 0x8e, 0x56, 0x6f, 0xbd, 0x54, 0xbb,
	0x33, 0x6b, 0xef, 0x1a, 0x13, 0xd2, 0xf6, 0x04, 0xfb, 0x7c, 0xfc, 0x7e, 0xbf, 0x99, 0x79, 0xe6,
	0x79, 0xc6, 0x60, 0x99, 0x1e, 0x60, 0x9b, 0xb5, 0xf3, 0xfb, 0xce, 0x61, 0xfe, 0x70, 0x15, 0x99,
	0xcd, 0x03, 0xb4, 0xea, 0x7e, 0xe4, 0x9a, 0x2d, 0x87, 0x39, 0x70, 0x96, 0xbb, 0x73, 0xae, 0xc5,
	0x77, 0xa7, 0x33, 0x0d, 0x87, 0x5a, 0x0e, 0xcd, 0xd7, 0x11, 0xc5, 0xf9, 0xc3, 0xd5, 0x3a, 0x66,
	0x68, 0x35, 0xdf, 0x70, 0x88, 0xcd, 0x93, 0xd2, 0x4b, 0xc2, 0xcf, 0x31, 0xb9, 0xbb, 0x07, 0x99,
	0x5e, 0xe0, 0x5e, 0xdd, 0xfb, 0xca, 0xf3, 0x0f, 0xe1, 0x9a, 0xdb, 0x77, 0xf6, 0x1d, 0x6e, 0x77,
	0xff, 0x13, 0x56, 0x79, 0xdf, 0x71, 0xf6, 0x4d, 0x9c, 0xf7, 0xbe, 0xea, 0xed, 0xbd, 0x3c, 0x23,
	0x16, 0xa6, 0x0c, 0x59, 0x4d, 0x1f, 0x71, 0x30, 0x00, 0xd9, 0xc7, 0xc2, 0x95, 0x19, 0x74, 0x19,
	0xed, 0x16, 0x62, 0xc4, 0x11, 0x52, 0x95, 0xdf, 0x63, 0x60, 0xea, 0x2e, 0xb6, 0x31, 0x25, 0xb4,
	0xca, 0x10, 0xc3, 0xf0, 0x73, 0x30, 0x47, 0x19, 0x6a, 0x31, 0x62, 0xef, 0xbb, 0x0a, 0x9b, 0x0e,
	0x45, 0xa6, 0x4e, 0x8c, 0x94, 0x94, 0x95, 0x56, 0x62, 0xaa, 0xdc, 0xed, 0xc8, 0x8b, 0xc7, 0xc8,
	0x32, 0xd7, 0x95, 0x61, 0x51, 0x8a, 0x06, 0x7d, 0x73, 0x45, 0x58, 0xcb, 0x06, 0xbc, 0x0f, 0x26,
	0x0c, 0xdc, 0x74, 0x28, 0x61, 0x34, 0x15, 0xcd, 0x8e, 0xac, 0xc4, 0xd7, 0x96, 0x72, 0x43, 0xb6,
	0x35, 0x57, 0xe4, 0x41, 0x6a, 0xf2, 0x45, 0x47, 0x8e, 0x3c, 0x7f, 0x2d, 0x4f, 0x08, 0x03, 0xd5,
	0x7a, 0xf9, 0xf0, 0x53, 0x30, 0x7a, 0xe8, 0x30, 0x4c, 0x53, 0x23, 0x1e, 0xd0, 0xc2, 0x50, 0xa0,
	0x5d, 0x87, 0x61, 0x35, 0x21, 0x50, 0x46, 0xdd, 0x2f, 0xaa, 0xf1, 0x34, 0xb8, 0x0d, 0x26, 0x7d,
	0xbd, 0x34, 0x15, 0xf3, 0x30, 0x96, 0x87, 0x62, 0xf8, 0xfa, 0xd5, 0x19, 0x81, 0x33, 0xe9, 0x5b,
	0xa8, 0xd6, 0x87, 0x80, 0x07, 0x60, 0x5a, 0x68, 0xd3, 0x9b, 0xa8, 0x85, 0x2c, 0x9a, 0x1a, 0xcd,
	0x4a, 0x2b, 0xf1, 0x35, 0xe5, 0x43, 0x2b, 0xac, 0x78, 0x91, 0xea, 0xb2, 0x8b, 0xdc, 0xed, 0xc8,
	0x57, 0xf9, 0x86, 0x86, 0x71, 0x14, 0x2d, 0x61, 0x04, 0xa3, 0x61, 0x03, 0x24, 0x0e, 0x1d, 0xbe,
	0xe1, 0x9c, 0x68, 0xcc, 0x23, 0xca, 0xe6, 0x44, 0x05, 0x71, 0x22, 0xaf, 0xd8, 0xdc, 0x0d, 0x70,
	0x8f, 0x80, 0xd3, 0x2c, 0x09, 0x9a, 0x39, 0x4e, 0x13, 0x02, 0x51, 0xb4, 0xa9, 0xc3, 0x40, 0x2c,
	0xfc, 0x1a, 0x4c, 0x31, 0x64, 0x9a, 0xc7, 0x3e, 0xc7, 0xb8, 0xe0, 0x18, 0xb6, 0x98, 0x9a, 0x1b,
	0x28, 0x38, 0x16, 0x05, 0xc7, 0x2c, 0xe7, 0x08, 0x62, 0x28, 0x5a, 0x9c, 0xf5, 0x23, 0xd7, 0x63,
	0x4f, 0x9e, 0xca, 0x92, 0xf2, 0x2d, 0x18, 0x17, 0x7b, 0x01, 0x3f, 0x01, 0xe3, 0x62, 0xa1, 0x5e,
	0x8d, 0xc5, 0xd7, 0x16, 0x87, 0xad, 0xc8, 0xaf, 0x8d, 0xd8, 0xcb, 0x8e, 0x2c, 0x69, 0x7e, 0x06,
	0xbc, 0x05, 0xc6, 0xd9, 0x91, 0x7e, 0x80, 0xe8, 0x41, 0x2a, 0x9a, 0x95, 0x56, 0x26, 0xd5, 0x99,
	0x6e, 0x47, 0x4e, 0x08, 0x11, 0x47, 0xae, 0x5d, 0xd1, 0xc6, 0xd8, 0xd1, 0x3d, 0x44, 0x0f, 0xd6,
	0x27, 0x1e, 0x3d, 0x95, 0x23, 0xef, 0x9e, 0xca, 0x11, 0xe5, 0x8f, 0x11, 0x90, 0x08, 0x1d, 0x05,
	0xfc, 0x45, 0x02, 0xb3, 0x16, 0xb1, 0x75, 0x62, 0x13, 0x46, 0x90, 0xa9, 0xf7, 0x15, 0xf1, 0x2a,
	0x13, 0x8a, 0xdc, 0x0b, 0xdf, 0x93, 0xb4, 0xe1, 0x10, 0x5b, 0x75, 0xdc, 0x85, 0xbf, 0xef, 0xc8,
	0xcb, 0x43, 0xb2, 0x6f, 0x3b, 0x16, 0x61, 0xd8, 0x6a, 0xb2, 0xe3, 0x6e, 0x47, 0x4e, 0x73, 0x51,
	0x43, 0xc2, 0x94, 0xe7, 0xaf, 0xe5, 0x95, 0x7d, 0xc2, 0x0e, 0xda, 0xf5, 0x5c, 0xc3, 0xb1, 0x44,
	0x47, 0x10, 0x7f, 0xee, 0x50, 0xe3, 0x61, 0x9e, 0x1d, 0x37, 0x31, 0xf5, 0xf8, 0xa8, 0x36, 0x63,
	0x11, 0xbb, 0xcc, 0x01, 0xfc, 0x1d, 0xfc, 0x41, 0x02, 0x71, 0x17, 0xd7, 0x17, 0x1d, 0xbd, 0x4c,
	0xb4, 0x2e, 0x44, 0x5f, 0x0d, 0x64, 0x85, 0xc4, 0xc2, 0xbe, 0xd8, 0xbf, 0x25, 0x12, 0x58, 0xc4,
	0xf6, 0xd5, 0x7d, 0x2f, 0x01, 0x68, 0xa1, 0x23, 0xbd, 0x57, 0xde, 0xb8, 0x45, 0x1c, 0x23, 0x35,
	0xe2, 0x9d, 0xf5, 0x42, 0x8e, 0xf7, 0xa7, 0x9c, 0xdf, 0x9f, 0x72, 0x45, 0xd1, 0x9f, 0xd4, 0x92,
	0x10, 0xb9, 0x74, 0x3e, 0x39, 0xa4, 0x75, 0x41, 0x68, 0x3d, 0x17, 0xa5, 0x3c, 0x79, 0x2d, 0x4b,
	0x5a, 0xd2, 0x42, 0x47, 0xfe, 0x59, 0x73, 0xf3, 0x4f, 0x51, 0x10, 0x0f, 0xd4, 0x2e, 0x2c, 0x82,
	0x84, 0x81, 0xf7, 0x50, 0xdb, 0x64, 0xba, 0x57, 0xa8, 0xa2, 0x0c, 0xe5, 0x61, 0x65, 0x18, 0xc8,
	0xd3, 0xa6, 0x44, 0x96, 0x67, 0x83, 0x0e, 0xb8, 0xd9, 0xc0, 0x2d, 0x46, 0xf6, 0x08, 0x6e, 0xe9,
	0xed, 0xa6, 0x81, 0x18, 0xd6, 0x29, 0x6e, 0xb4, 0x5b, 0x84, 0x1d, 0xeb, 0x6e, 0xef, 0x11, 0xe0,
	0xd1, 0x8f, 0x03, 0xcf, 0xf6, 0xc0, 0x1e, 0x78, 0x58, 0x55, 0x01, 0xe5, 0x36, 0x33, 0x4e, 0x48,
	0xc0, 0xf5, 0xf3, 0x84, 0x0c, 0x3d, 0xc4, 0x41, 0xb6, 0x91, 0x8f, 0x63, 0x5b, 0x1e, 0x64, 0x73,
	0x71, 0x7a, 0x54, 0xca, 0xb3, 0x09, 0x30, 0xe1, 0x77, 0x3f, 0xf7, 0xbe, 0x36, 0x1c, 0x9b, 0x61,
	0xdb, 0xbf, 0xaf, 0x73, 0xe7, 0xce, 0xb0, 0x60, 0x1f, 0xab, 0xf1, 0x5f, 0x7f, 0xbe, 0x33, 0xbe,
	0xc1, 0x03, 0x35, 0x3f, 0x03, 0xfe, 0x0f, 0xc4, 0x83, 0x43, 0x25, 0xea, 0x0d, 0x95, 0xd9, 0xf7,
	0x1d, 0x39, 0x4a, 0x8c, 0x6e, 0x47, 0x9e, 0xe4, 0x67, 0xe9, 0x0e, 0x12, 0xd0, 0xec, 0x0f, 0x90,
	0x2f, 0xc0, 0x18, 0x65, 0x88, 0xb5, 0xa9, 0xb7, 0x9e, 0xe9, 0xb5, 0x1b, 0x1f, 0xec, 0xd8, 0x55,
	0x2f, 0x54, 0x4d, 0x77, 0x3b, 0xf2, 0x35, 0x8e, 0xd7, 0xa3, 0xe4, 0x28, 0x8a, 0x26, 0xe0, 0x60,
	0x1d, 0xa4, 0x09, 0x15, 0x03, 0x0c, 0xb7, 0xf4, 0x86, 0xd3, 0xb6, 0x1b, 0xc4, 0xd4, 0x2d, 0x6c,
	0xd5, 0x71, 0x2b, 0x15, 0xcb, 0x4a, 0x2b, 0x13, 0xea, 0xbf, 0xba, 0x1d, 0xf9, 0xba, 0xd0, 0x75,
	0x61, 0xac, 0xa2, 0xcd, 0x13, 0x5a, 0x11, 0xbe, 0x0d, 0xee, 0xfa, 0xcc, 0xf3, 0xc0, 0x4d, 0x90,
	0xec, 0x25, 0x21, 0xc3, 0x68, 0x61, 0xca, 0x67, 0xc4, 0xa4, 0xba, 0xd8, 0xed, 0xc8, 0xf3, 0x41,
	0x85, 0xfd, 0x08, 0x45, 0xbb, 0xe2, 0x9b, 0x0a, 0xdc, 0x02, 0x9b, 0x00, 0xee, 0x11, 0x1b, 0x99,
	0xfc, 0x64, 0xf5, 0x16, 0xa6, 0x6d, 0x93, 0x89, 0x21, 0x70, 0xf1, 0x01, 0x6b, 0x5e, 0x98, 0x7a,
	0x5d, 0xf4, 0x67, 0x71, 0x59, 0xce, 0x03, 0x29, 0x5a, 0xd2, 0x33, 0x06, 0x92, 0xe0, 0x57, 0x20,
	0x4e, 0xdb, 0x75, 0x8b, 0x30, 0xdd, 0x7d, 0x70, 0x88, 0x59, 0x90, 0x3e, 0x77, 0xda, 0x35, 0xff,
	0x35, 0xa2, 0x66, 0x04, 0x8b, 0x68, 0x1f, 0x81, 0x64, 0xe5, 0xb1, 0x7b, 0x17, 0x01, 0xb7, 0xb8,
	0x09, 0x90, 0x80, 0xa4, 0x7f, 0x5d, 0xb1, 0x6d, 0x70, 0x86, 0x89, 0x4b, 0x19, 0x6e, 0x08, 0x86,
	0xf9, 0xf0, 0xc8, 0xf4, 0x11, 0x38, 0x8d, 0x3f, 0x91, 0x4b, 0xb6, 0xe1, 0x51, 0x3d, 0x92, 0x40,
	0x82, 0x39, 0x2c, 0xd0, 0xd6, 0x27, 0x2f, 0xeb, 0x90, 0xf7, 0xc2, 0x33, 0x33, 0x94, 0xfd, 0xd7,
	0x5a, 0xe1, 0x94, 0x97, 0xeb, 0x37, 0x43, 0x13, 0xcc, 0x88, 0xf9, 0xeb, 0xbd, 0x93, 0xf8, 0xb2,
	0xc1, 0xa5, 0xcb, 0xbe, 0x29, 0xe4, 0xa4, 0x42, 0x23, 0xbc, 0x0f, 0xc1, 0xd7, 0x7d, 0x85, 0xdb,
	0xab, 0xae, 0xd9, 0x5b, 0xf8, 0x1e, 0x10, 0xa6, 0xfe, 0x16, 0xc7, 0x2f, 0xe5, 0x52, 0x04, 0xd7,
	0xb5, 0x10, 0x57, 0x78, 0x87, 0xc5, 0x4b, 0x44, 0x6c, 0xf0, 0x7a, 0xec, 0x9d, 0x3b, 0xd3, 0x0f,
	0x41, 0xcc, 0x6d, 0x19, 0xf0, 0xff, 0x83, 0x03, 0x3d, 0x75, 0xc1, 0x13, 0x05, 0xff, 0xa3, 0x69,
	0xfe, 0x44, 0x4c, 0xf3, 0x5b, 0x3f, 0x8e, 0x80, 0xe9, 0xf0, 0xdd, 0x87, 0x39, 0xb0, 0x58, 0xd1,
	0x76, 0x2a, 0x3b, 0xd5, 0xc2, 0x96, 0x5e, 0xad, 0x15, 0x6a, 0x0f, 0xaa, 0xfa, 0x83, 0xed, 0x6a,
	0xa5, 0xb4, 0x51, 0xde, 0x2c, 0x97, 0x8a, 0xc9, 0x48, 0x3a, 0x71, 0x72, 0x9a, 0x9d, 0xe4, 0xc1,
	0xdb, 0xc4, 0xed, 0x69, 0x99, 0xc1, 0xf8, 0x62, 0xa9, 0xb2, 0x53, 0x2d, 0xd7, 0xf4, 0x4a, 0x49,
	0x2b, 0xef, 0x14, 0x93, 0x52, 0x7a, 0xfe, 0xe4, 0x34, 0x3b, 0xcb, 0x53, 0x42, 0xf3, 0x04, 0x6e,
	0x81, 0x7f, 0x0f, 0x26, 0x6f, 0x94, 0xb4, 0x9a, 0x4b, 0xa5, 0xe9, 0xbb, 0x3b, 0xb5, 0xf2, 0xf6,
	0x5d, 0x1f, 0x25, 0x9a, 0x96, 0x4f, 0x4e, 0xb3, 0x8b, 0x1c, 0x65, 0xc3, 0xef, 0xb9, 0xe2, 0xb5,
	0x76, 0x21, 0xda, 0x6e, 0x61, 0xab, 0x5c, 0x2c, 0xd4, 0x76, 0x06, 0xd1, 0x46, 0x82, 0x68, 0xbb,
	0xc8, 0x24, 0x06, 0x62, 0x4e, 0x18, 0xed, 0x36, 0xb8, 0x36, 0x88, 0x56, 0x29, 0x54, 0xab, 0xa5,
	0x62, 0x32, 0x96, 0x4e, 0x9e, 0x9c, 0x66, 0xa7, 0x78, 0x72, 0x05, 0x51, 0x8a, 0x0d, 0xf8, 0x1f,
	0x90, 0x1a, 0x8c, 0xd6, 0x4a, 0xf7, 0x4b, 0x1b, 0xb5, 0x52, 0x31, 0x39, 0x9a, 0x86, 0x27, 0xa7,
	0xd9, 0x69, 0x1e, 0xaf, 0xe1, 0x6f, 0x70, 0x83, 0xe1, 0xa1, 0xf8, 0x9b, 0x85, 0xf2, 0x56, 0xa9,
	0x98, 0x1c, 0x0b, 0xe2, 0x6f, 0x22, 0x62, 0x62, 0x23, 0x1d, 0x7b, 0xf4, 0x2c, 0x13, 0x51, 0x6b,
	0x2f, 0xde, 0x64, 0x22, 0xaf, 0xde, 0x64, 0x22, 0xdf, 0x9d, 0x65, 0x22, 0x2f, 0xce, 0x32, 0xd2,
	0xcb, 0xb3, 0x8c, 0xf4, 0xdb, 0x59, 0x46, 0x7a, 0xfc, 0x36, 0x13, 0x79, 0xf9, 0x36, 0x13, 0x79,
	0xf5, 0x36, 0x13, 0xf9, 0x32, 0x17, 0xbc, 0x61, 0xee, 0x4e, 0x3d, 0xdc, 0x73, 0xda, 0xb6, 0xe1,
	0xbd, 0x0d, 0xf2, 0xe2, 0x37, 0xdb, 0x91, 0xf7, 0x0b, 0xcb, 0xbb, 0x6d, 0xf5, 0x31, 0xaf, 0x94,
	0xff, 0xfb, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x11, 0xf5, 0xaf, 0xd0, 0x0d, 0x00, 0x00,
}

func (this *Proposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Proposal)
	if !ok {
		that2, ok := that.(Proposal)
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
	if !this.Content.Equal(that1.Content) {
		return false
	}
	if this.ProposalId != that1.ProposalId {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.IsProposerCouncilMember != that1.IsProposerCouncilMember {
		return false
	}
	if this.ProposerAddress != that1.ProposerAddress {
		return false
	}
	if !this.FinalTallyResult.Equal(&that1.FinalTallyResult) {
		return false
	}
	if !this.SubmitTime.Equal(that1.SubmitTime) {
		return false
	}
	if !this.DepositEndTime.Equal(that1.DepositEndTime) {
		return false
	}
	if len(this.TotalDeposit) != len(that1.TotalDeposit) {
		return false
	}
	for i := range this.TotalDeposit {
		if !this.TotalDeposit[i].Equal(&that1.TotalDeposit[i]) {
			return false
		}
	}
	if !this.VotingStartTime.Equal(that1.VotingStartTime) {
		return false
	}
	if !this.VotingEndTime.Equal(that1.VotingEndTime) {
		return false
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGov(dAtA, i, m.StartingProposalId)
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintGov(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Deposit != nil {
		{
			size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DepositParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxDepositPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxDepositPeriod):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintGov(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x1a
	if len(m.MinDeposit) > 0 {
		for iNdEx := len(m.MinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MinInitialDeposit) > 0 {
		for iNdEx := len(m.MinInitialDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinInitialDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TallyParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TallyParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TallyParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertifierUpdateStakeVoteTally != nil {
		{
			size, err := m.CertifierUpdateStakeVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.CertifierUpdateSecurityVoteTally != nil {
		{
			size, err := m.CertifierUpdateSecurityVoteTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultTally != nil {
		{
			size, err := m.DefaultTally.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n9, err9 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotingEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime):])
	if err9 != nil {
		return 0, err9
	}
	i -= n9
	i = encodeVarintGov(dAtA, i, uint64(n9))
	i--
	dAtA[i] = 0x5a
	n10, err10 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotingStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingStartTime):])
	if err10 != nil {
		return 0, err10
	}
	i -= n10
	i = encodeVarintGov(dAtA, i, uint64(n10))
	i--
	dAtA[i] = 0x52
	if len(m.TotalDeposit) > 0 {
		for iNdEx := len(m.TotalDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	n11, err11 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DepositEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositEndTime):])
	if err11 != nil {
		return 0, err11
	}
	i -= n11
	i = encodeVarintGov(dAtA, i, uint64(n11))
	i--
	dAtA[i] = 0x42
	n12, err12 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SubmitTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime):])
	if err12 != nil {
		return 0, err12
	}
	i -= n12
	i = encodeVarintGov(dAtA, i, uint64(n12))
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.FinalTallyResult.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintGov(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.IsProposerCouncilMember {
		i--
		if m.IsProposerCouncilMember {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.ProposalId != 0 {
		i = encodeVarintGov(dAtA, i, m.ProposalId)
		i--
		dAtA[i] = 0x10
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintGov(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Vote != nil {
		{
			size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGov(m.StartingProposalId)
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = m.DepositParams.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.VotingParams.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Deposit != nil {
		l = m.Deposit.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *DepositParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinInitialDeposit) > 0 {
		for _, e := range m.MinInitialDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.MinDeposit) > 0 {
		for _, e := range m.MinDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxDepositPeriod)
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *TallyParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultTally != nil {
		l = m.DefaultTally.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.CertifierUpdateSecurityVoteTally != nil {
		l = m.CertifierUpdateSecurityVoteTally.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.CertifierUpdateStakeVoteTally != nil {
		l = m.CertifierUpdateStakeVoteTally.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.ProposalId != 0 {
		n += 1 + sovGov(m.ProposalId)
	}
	if m.Status != 0 {
		n += 1 + sovGov(uint64(m.Status))
	}
	if m.IsProposerCouncilMember {
		n += 2
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.FinalTallyResult.Size()
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime)
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositEndTime)
	n += 1 + l + sovGov(uint64(l))
	if len(m.TotalDeposit) > 0 {
		for _, e := range m.TotalDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingStartTime)
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime)
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vote != nil {
		l = m.Vote.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deposit == nil {
				m.Deposit = &types.Deposit{}
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *DepositParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: DepositParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinInitialDeposit = append(m.MinInitialDeposit, types1.Coin{})
			if err := m.MinInitialDeposit[len(m.MinInitialDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposit = append(m.MinDeposit, types1.Coin{})
			if err := m.MinDeposit[len(m.MinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDepositPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxDepositPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *TallyParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: TallyParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TallyParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultTally", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultTally == nil {
				m.DefaultTally = &types.TallyParams{}
			}
			if err := m.DefaultTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateSecurityVoteTally", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertifierUpdateSecurityVoteTally == nil {
				m.CertifierUpdateSecurityVoteTally = &types.TallyParams{}
			}
			if err := m.CertifierUpdateSecurityVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifierUpdateStakeVoteTally", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CertifierUpdateStakeVoteTally == nil {
				m.CertifierUpdateStakeVoteTally = &types.TallyParams{}
			}
			if err := m.CertifierUpdateStakeVoteTally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types2.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsProposerCouncilMember", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsProposerCouncilMember = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalTallyResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FinalTallyResult.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SubmitTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DepositEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDeposit = append(m.TotalDeposit, types1.Coin{})
			if err := m.TotalDeposit[len(m.TotalDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotingStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotingEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Vote == nil {
				m.Vote = &types.Vote{}
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
