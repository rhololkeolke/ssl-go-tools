// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages_robocup_ssl_referee.proto

package sslproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are the "coarse" stages of the game.
type SSL_Referee_Stage int32

const (
	// The first half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_NORMAL_FIRST_HALF_PRE SSL_Referee_Stage = 0
	// The first half of the normal game, before half time.
	SSL_Referee_NORMAL_FIRST_HALF SSL_Referee_Stage = 1
	// Half time between first and second halves.
	SSL_Referee_NORMAL_HALF_TIME SSL_Referee_Stage = 2
	// The second half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_NORMAL_SECOND_HALF_PRE SSL_Referee_Stage = 3
	// The second half of the normal game, after half time.
	SSL_Referee_NORMAL_SECOND_HALF SSL_Referee_Stage = 4
	// The break before extra time.
	SSL_Referee_EXTRA_TIME_BREAK SSL_Referee_Stage = 5
	// The first half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_EXTRA_FIRST_HALF_PRE SSL_Referee_Stage = 6
	// The first half of extra time.
	SSL_Referee_EXTRA_FIRST_HALF SSL_Referee_Stage = 7
	// Half time between first and second extra halves.
	SSL_Referee_EXTRA_HALF_TIME SSL_Referee_Stage = 8
	// The second half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	SSL_Referee_EXTRA_SECOND_HALF_PRE SSL_Referee_Stage = 9
	// The second half of extra time.
	SSL_Referee_EXTRA_SECOND_HALF SSL_Referee_Stage = 10
	// The break before penalty shootout.
	SSL_Referee_PENALTY_SHOOTOUT_BREAK SSL_Referee_Stage = 11
	// The penalty shootout.
	SSL_Referee_PENALTY_SHOOTOUT SSL_Referee_Stage = 12
	// The game is over.
	SSL_Referee_POST_GAME SSL_Referee_Stage = 13
)

var SSL_Referee_Stage_name = map[int32]string{
	0:  "NORMAL_FIRST_HALF_PRE",
	1:  "NORMAL_FIRST_HALF",
	2:  "NORMAL_HALF_TIME",
	3:  "NORMAL_SECOND_HALF_PRE",
	4:  "NORMAL_SECOND_HALF",
	5:  "EXTRA_TIME_BREAK",
	6:  "EXTRA_FIRST_HALF_PRE",
	7:  "EXTRA_FIRST_HALF",
	8:  "EXTRA_HALF_TIME",
	9:  "EXTRA_SECOND_HALF_PRE",
	10: "EXTRA_SECOND_HALF",
	11: "PENALTY_SHOOTOUT_BREAK",
	12: "PENALTY_SHOOTOUT",
	13: "POST_GAME",
}
var SSL_Referee_Stage_value = map[string]int32{
	"NORMAL_FIRST_HALF_PRE":  0,
	"NORMAL_FIRST_HALF":      1,
	"NORMAL_HALF_TIME":       2,
	"NORMAL_SECOND_HALF_PRE": 3,
	"NORMAL_SECOND_HALF":     4,
	"EXTRA_TIME_BREAK":       5,
	"EXTRA_FIRST_HALF_PRE":   6,
	"EXTRA_FIRST_HALF":       7,
	"EXTRA_HALF_TIME":        8,
	"EXTRA_SECOND_HALF_PRE":  9,
	"EXTRA_SECOND_HALF":      10,
	"PENALTY_SHOOTOUT_BREAK": 11,
	"PENALTY_SHOOTOUT":       12,
	"POST_GAME":              13,
}

func (x SSL_Referee_Stage) Enum() *SSL_Referee_Stage {
	p := new(SSL_Referee_Stage)
	*p = x
	return p
}
func (x SSL_Referee_Stage) String() string {
	return proto.EnumName(SSL_Referee_Stage_name, int32(x))
}
func (x *SSL_Referee_Stage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Stage_value, data, "SSL_Referee_Stage")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Stage(value)
	return nil
}
func (SSL_Referee_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e, []int{0, 0}
}

// These are the "fine" states of play on the field.
type SSL_Referee_Command int32

const (
	// All robots should completely stop moving.
	SSL_Referee_HALT SSL_Referee_Command = 0
	// Robots must keep 50 cm from the ball.
	SSL_Referee_STOP SSL_Referee_Command = 1
	// A prepared kickoff or penalty may now be taken.
	SSL_Referee_NORMAL_START SSL_Referee_Command = 2
	// The ball is dropped and free for either team.
	SSL_Referee_FORCE_START SSL_Referee_Command = 3
	// The yellow team may move into kickoff position.
	SSL_Referee_PREPARE_KICKOFF_YELLOW SSL_Referee_Command = 4
	// The blue team may move into kickoff position.
	SSL_Referee_PREPARE_KICKOFF_BLUE SSL_Referee_Command = 5
	// The yellow team may move into penalty position.
	SSL_Referee_PREPARE_PENALTY_YELLOW SSL_Referee_Command = 6
	// The blue team may move into penalty position.
	SSL_Referee_PREPARE_PENALTY_BLUE SSL_Referee_Command = 7
	// The yellow team may take a direct free kick.
	SSL_Referee_DIRECT_FREE_YELLOW SSL_Referee_Command = 8
	// The blue team may take a direct free kick.
	SSL_Referee_DIRECT_FREE_BLUE SSL_Referee_Command = 9
	// The yellow team may take an indirect free kick.
	SSL_Referee_INDIRECT_FREE_YELLOW SSL_Referee_Command = 10
	// The blue team may take an indirect free kick.
	SSL_Referee_INDIRECT_FREE_BLUE SSL_Referee_Command = 11
	// The yellow team is currently in a timeout.
	SSL_Referee_TIMEOUT_YELLOW SSL_Referee_Command = 12
	// The blue team is currently in a timeout.
	SSL_Referee_TIMEOUT_BLUE SSL_Referee_Command = 13
	// The yellow team just scored a goal.
	// For information only.
	// For rules compliance, teams must treat as STOP.
	SSL_Referee_GOAL_YELLOW SSL_Referee_Command = 14
	// The blue team just scored a goal.
	SSL_Referee_GOAL_BLUE SSL_Referee_Command = 15
	// Equivalent to STOP, but the yellow team must pick up the ball and
	// drop it in the Designated Position.
	SSL_Referee_BALL_PLACEMENT_YELLOW SSL_Referee_Command = 16
	// Equivalent to STOP, but the blue team must pick up the ball and drop
	// it in the Designated Position.
	SSL_Referee_BALL_PLACEMENT_BLUE SSL_Referee_Command = 17
)

var SSL_Referee_Command_name = map[int32]string{
	0:  "HALT",
	1:  "STOP",
	2:  "NORMAL_START",
	3:  "FORCE_START",
	4:  "PREPARE_KICKOFF_YELLOW",
	5:  "PREPARE_KICKOFF_BLUE",
	6:  "PREPARE_PENALTY_YELLOW",
	7:  "PREPARE_PENALTY_BLUE",
	8:  "DIRECT_FREE_YELLOW",
	9:  "DIRECT_FREE_BLUE",
	10: "INDIRECT_FREE_YELLOW",
	11: "INDIRECT_FREE_BLUE",
	12: "TIMEOUT_YELLOW",
	13: "TIMEOUT_BLUE",
	14: "GOAL_YELLOW",
	15: "GOAL_BLUE",
	16: "BALL_PLACEMENT_YELLOW",
	17: "BALL_PLACEMENT_BLUE",
}
var SSL_Referee_Command_value = map[string]int32{
	"HALT":                   0,
	"STOP":                   1,
	"NORMAL_START":           2,
	"FORCE_START":            3,
	"PREPARE_KICKOFF_YELLOW": 4,
	"PREPARE_KICKOFF_BLUE":   5,
	"PREPARE_PENALTY_YELLOW": 6,
	"PREPARE_PENALTY_BLUE":   7,
	"DIRECT_FREE_YELLOW":     8,
	"DIRECT_FREE_BLUE":       9,
	"INDIRECT_FREE_YELLOW":   10,
	"INDIRECT_FREE_BLUE":     11,
	"TIMEOUT_YELLOW":         12,
	"TIMEOUT_BLUE":           13,
	"GOAL_YELLOW":            14,
	"GOAL_BLUE":              15,
	"BALL_PLACEMENT_YELLOW":  16,
	"BALL_PLACEMENT_BLUE":    17,
}

func (x SSL_Referee_Command) Enum() *SSL_Referee_Command {
	p := new(SSL_Referee_Command)
	*p = x
	return p
}
func (x SSL_Referee_Command) String() string {
	return proto.EnumName(SSL_Referee_Command_name, int32(x))
}
func (x *SSL_Referee_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Command_value, data, "SSL_Referee_Command")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Command(value)
	return nil
}
func (SSL_Referee_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e, []int{0, 1}
}

// Each UDP packet contains one of these messages.
type SSL_Referee struct {
	// The UNIX timestamp when the packet was sent, in microseconds.
	// Divide by 1,000,000 to get a time_t.
	PacketTimestamp *uint64            `protobuf:"varint,1,req,name=packet_timestamp,json=packetTimestamp" json:"packet_timestamp,omitempty"`
	Stage           *SSL_Referee_Stage `protobuf:"varint,2,req,name=stage,enum=SSL_Referee_Stage" json:"stage,omitempty"`
	// The number of microseconds left in the stage.
	// The following stages have this value; the rest do not:
	// NORMAL_FIRST_HALF
	// NORMAL_HALF_TIME
	// NORMAL_SECOND_HALF
	// EXTRA_TIME_BREAK
	// EXTRA_FIRST_HALF
	// EXTRA_HALF_TIME
	// EXTRA_SECOND_HALF
	// PENALTY_SHOOTOUT_BREAK
	//
	// If the stage runs over its specified time, this value
	// becomes negative.
	StageTimeLeft *int32               `protobuf:"zigzag32,3,opt,name=stage_time_left,json=stageTimeLeft" json:"stage_time_left,omitempty"`
	Command       *SSL_Referee_Command `protobuf:"varint,4,req,name=command,enum=SSL_Referee_Command" json:"command,omitempty"`
	// The number of commands issued since startup (mod 2^32).
	CommandCounter *uint32 `protobuf:"varint,5,req,name=command_counter,json=commandCounter" json:"command_counter,omitempty"`
	// The UNIX timestamp when the command was issued, in microseconds.
	// This value changes only when a new command is issued, not on each packet.
	CommandTimestamp *uint64 `protobuf:"varint,6,req,name=command_timestamp,json=commandTimestamp" json:"command_timestamp,omitempty"`
	// Information about the two teams.
	Yellow             *SSL_Referee_TeamInfo `protobuf:"bytes,7,req,name=yellow" json:"yellow,omitempty"`
	Blue               *SSL_Referee_TeamInfo `protobuf:"bytes,8,req,name=blue" json:"blue,omitempty"`
	DesignatedPosition *SSL_Referee_Point    `protobuf:"bytes,9,opt,name=designated_position,json=designatedPosition" json:"designated_position,omitempty"`
	// Information about the direction of play.
	// True, if the blue team will have it's goal on the positive x-axis of the ssl-vision coordinate system
	// Obviously, the yellow team will play on the opposide half
	// For compatibility, this field is optional
	BlueTeamOnPositiveHalf *bool `protobuf:"varint,10,opt,name=blueTeamOnPositiveHalf" json:"blueTeamOnPositiveHalf,omitempty"`
	// The game event that caused the referee command
	GameEvent            *SSL_Referee_Game_Event `protobuf:"bytes,11,opt,name=gameEvent" json:"gameEvent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SSL_Referee) Reset()         { *m = SSL_Referee{} }
func (m *SSL_Referee) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee) ProtoMessage()    {}
func (*SSL_Referee) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e, []int{0}
}
func (m *SSL_Referee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee.Unmarshal(m, b)
}
func (m *SSL_Referee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee.Merge(dst, src)
}
func (m *SSL_Referee) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee.Size(m)
}
func (m *SSL_Referee) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee proto.InternalMessageInfo

func (m *SSL_Referee) GetPacketTimestamp() uint64 {
	if m != nil && m.PacketTimestamp != nil {
		return *m.PacketTimestamp
	}
	return 0
}

func (m *SSL_Referee) GetStage() SSL_Referee_Stage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return SSL_Referee_NORMAL_FIRST_HALF_PRE
}

func (m *SSL_Referee) GetStageTimeLeft() int32 {
	if m != nil && m.StageTimeLeft != nil {
		return *m.StageTimeLeft
	}
	return 0
}

func (m *SSL_Referee) GetCommand() SSL_Referee_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return SSL_Referee_HALT
}

func (m *SSL_Referee) GetCommandCounter() uint32 {
	if m != nil && m.CommandCounter != nil {
		return *m.CommandCounter
	}
	return 0
}

func (m *SSL_Referee) GetCommandTimestamp() uint64 {
	if m != nil && m.CommandTimestamp != nil {
		return *m.CommandTimestamp
	}
	return 0
}

func (m *SSL_Referee) GetYellow() *SSL_Referee_TeamInfo {
	if m != nil {
		return m.Yellow
	}
	return nil
}

func (m *SSL_Referee) GetBlue() *SSL_Referee_TeamInfo {
	if m != nil {
		return m.Blue
	}
	return nil
}

func (m *SSL_Referee) GetDesignatedPosition() *SSL_Referee_Point {
	if m != nil {
		return m.DesignatedPosition
	}
	return nil
}

func (m *SSL_Referee) GetBlueTeamOnPositiveHalf() bool {
	if m != nil && m.BlueTeamOnPositiveHalf != nil {
		return *m.BlueTeamOnPositiveHalf
	}
	return false
}

func (m *SSL_Referee) GetGameEvent() *SSL_Referee_Game_Event {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

// Information about a single team.
type SSL_Referee_TeamInfo struct {
	// The team's name (empty string if operator has not typed anything).
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The number of goals scored by the team during normal play and overtime.
	Score *uint32 `protobuf:"varint,2,req,name=score" json:"score,omitempty"`
	// The number of red cards issued to the team since the beginning of the game.
	RedCards *uint32 `protobuf:"varint,3,req,name=red_cards,json=redCards" json:"red_cards,omitempty"`
	// The amount of time (in microseconds) left on each yellow card issued to the team.
	// If no yellow cards are issued, this array has no elements.
	// Otherwise, times are ordered from smallest to largest.
	YellowCardTimes []uint32 `protobuf:"varint,4,rep,packed,name=yellow_card_times,json=yellowCardTimes" json:"yellow_card_times,omitempty"`
	// The total number of yellow cards ever issued to the team.
	YellowCards *uint32 `protobuf:"varint,5,req,name=yellow_cards,json=yellowCards" json:"yellow_cards,omitempty"`
	// The number of timeouts this team can still call.
	// If in a timeout right now, that timeout is excluded.
	Timeouts *uint32 `protobuf:"varint,6,req,name=timeouts" json:"timeouts,omitempty"`
	// The number of microseconds of timeout this team can use.
	TimeoutTime *uint32 `protobuf:"varint,7,req,name=timeout_time,json=timeoutTime" json:"timeout_time,omitempty"`
	// The pattern number of this team's goalie.
	Goalie               *uint32  `protobuf:"varint,8,req,name=goalie" json:"goalie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_Referee_TeamInfo) Reset()         { *m = SSL_Referee_TeamInfo{} }
func (m *SSL_Referee_TeamInfo) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_TeamInfo) ProtoMessage()    {}
func (*SSL_Referee_TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e, []int{0, 0}
}
func (m *SSL_Referee_TeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Unmarshal(m, b)
}
func (m *SSL_Referee_TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_TeamInfo.Merge(dst, src)
}
func (m *SSL_Referee_TeamInfo) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_TeamInfo.Size(m)
}
func (m *SSL_Referee_TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_TeamInfo proto.InternalMessageInfo

func (m *SSL_Referee_TeamInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SSL_Referee_TeamInfo) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetRedCards() uint32 {
	if m != nil && m.RedCards != nil {
		return *m.RedCards
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetYellowCardTimes() []uint32 {
	if m != nil {
		return m.YellowCardTimes
	}
	return nil
}

func (m *SSL_Referee_TeamInfo) GetYellowCards() uint32 {
	if m != nil && m.YellowCards != nil {
		return *m.YellowCards
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetTimeouts() uint32 {
	if m != nil && m.Timeouts != nil {
		return *m.Timeouts
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetTimeoutTime() uint32 {
	if m != nil && m.TimeoutTime != nil {
		return *m.TimeoutTime
	}
	return 0
}

func (m *SSL_Referee_TeamInfo) GetGoalie() uint32 {
	if m != nil && m.Goalie != nil {
		return *m.Goalie
	}
	return 0
}

// The coordinates of the Designated Position. These are measured in
// millimetres and correspond to SSL-Vision coordinates. These fields are
// always either both present (in the case of a ball placement command) or
// both absent (in the case of any other command).
type SSL_Referee_Point struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_Referee_Point) Reset()         { *m = SSL_Referee_Point{} }
func (m *SSL_Referee_Point) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_Point) ProtoMessage()    {}
func (*SSL_Referee_Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e, []int{0, 1}
}
func (m *SSL_Referee_Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_Point.Unmarshal(m, b)
}
func (m *SSL_Referee_Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_Point.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_Point.Merge(dst, src)
}
func (m *SSL_Referee_Point) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_Point.Size(m)
}
func (m *SSL_Referee_Point) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_Point.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_Point proto.InternalMessageInfo

func (m *SSL_Referee_Point) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *SSL_Referee_Point) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*SSL_Referee)(nil), "SSL_Referee")
	proto.RegisterType((*SSL_Referee_TeamInfo)(nil), "SSL_Referee.TeamInfo")
	proto.RegisterType((*SSL_Referee_Point)(nil), "SSL_Referee.Point")
	proto.RegisterEnum("SSL_Referee_Stage", SSL_Referee_Stage_name, SSL_Referee_Stage_value)
	proto.RegisterEnum("SSL_Referee_Command", SSL_Referee_Command_name, SSL_Referee_Command_value)
}

func init() {
	proto.RegisterFile("messages_robocup_ssl_referee.proto", fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e)
}

var fileDescriptor_messages_robocup_ssl_referee_a0009bf4f26b1c0e = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xe2, 0x46,
	0x14, 0x5d, 0x9b, 0xef, 0x6b, 0x1c, 0x86, 0x9b, 0x8f, 0x75, 0xe9, 0x0b, 0x4d, 0xd5, 0xd6, 0xab,
	0xaa, 0x3c, 0xac, 0xd4, 0xbe, 0x3b, 0xde, 0x21, 0x41, 0x71, 0x30, 0x1a, 0xbc, 0x6a, 0xf7, 0x69,
	0xe4, 0x85, 0x01, 0xa1, 0x82, 0x8d, 0x6c, 0xb3, 0xdd, 0xfc, 0x81, 0xfe, 0x8c, 0xfe, 0xa7, 0xfe,
	0x9e, 0xbe, 0x54, 0x33, 0xb6, 0xd7, 0x84, 0xa4, 0x7d, 0x9b, 0x39, 0xe7, 0xdc, 0x33, 0xf7, 0xde,
	0x63, 0xc3, 0xf5, 0x4e, 0xa4, 0x69, 0xb8, 0x16, 0x29, 0x4f, 0xe2, 0x8f, 0xf1, 0xe2, 0xb0, 0xe7,
	0x69, 0xba, 0xe5, 0x89, 0x58, 0x89, 0x44, 0x88, 0xd1, 0x3e, 0x89, 0xb3, 0x78, 0xf0, 0xdd, 0x8b,
	0x9a, 0x75, 0xb8, 0x13, 0x5c, 0x7c, 0x12, 0x51, 0x96, 0xcb, 0xae, 0xff, 0xea, 0x82, 0x31, 0x9f,
	0x7b, 0x9c, 0xe5, 0xc5, 0xf8, 0x06, 0xc8, 0x3e, 0x5c, 0xfc, 0x2e, 0x32, 0x9e, 0x6d, 0x76, 0x22,
	0xcd, 0xc2, 0xdd, 0xde, 0xd2, 0x86, 0xba, 0x5d, 0x67, 0xbd, 0x1c, 0x0f, 0x4a, 0x18, 0x6d, 0x68,
	0xa4, 0x59, 0xb8, 0x16, 0x96, 0x3e, 0xd4, 0xed, 0xb3, 0xb7, 0x38, 0x3a, 0xf2, 0x19, 0xcd, 0x25,
	0xc3, 0x72, 0x01, 0x7e, 0x0f, 0x3d, 0x75, 0x50, 0x9e, 0x7c, 0x2b, 0x56, 0x99, 0x55, 0x1b, 0x6a,
	0x76, 0x9f, 0x99, 0x0a, 0x96, 0x96, 0x9e, 0x58, 0x65, 0x38, 0x82, 0xd6, 0x22, 0xde, 0xed, 0xc2,
	0x68, 0x69, 0xd5, 0x95, 0xe7, 0xc5, 0x13, 0x4f, 0x37, 0xe7, 0x58, 0x29, 0xc2, 0x1f, 0xa0, 0x57,
	0x1c, 0xf9, 0x22, 0x3e, 0x44, 0x99, 0x48, 0xac, 0xc6, 0x50, 0xb7, 0x4d, 0x76, 0x56, 0xc0, 0x6e,
	0x8e, 0xe2, 0x8f, 0xd0, 0x2f, 0x85, 0xd5, 0x58, 0x4d, 0x35, 0x16, 0x29, 0x88, 0x6a, 0xae, 0x9f,
	0xa0, 0xf9, 0x28, 0xb6, 0xdb, 0xf8, 0x0f, 0xab, 0x35, 0xd4, 0x6d, 0xe3, 0xed, 0xe5, 0x93, 0x26,
	0x02, 0x11, 0xee, 0x26, 0xd1, 0x2a, 0x66, 0x85, 0x08, 0xdf, 0x40, 0xfd, 0xe3, 0xf6, 0x20, 0xac,
	0xf6, 0xff, 0x89, 0x95, 0x04, 0x5d, 0x38, 0x5f, 0x8a, 0x74, 0xb3, 0x8e, 0xc2, 0x4c, 0x2c, 0xf9,
	0x3e, 0x4e, 0x37, 0xd9, 0x26, 0x8e, 0xac, 0xce, 0x50, 0xb3, 0x8d, 0x93, 0xfd, 0xcd, 0xe2, 0x4d,
	0x94, 0x31, 0xac, 0xe4, 0xb3, 0x42, 0x8d, 0xbf, 0xc0, 0x95, 0x34, 0x93, 0xd6, 0x7e, 0x94, 0xa3,
	0x9f, 0xc4, 0x5d, 0xb8, 0x5d, 0x59, 0x30, 0xd4, 0xec, 0x36, 0xfb, 0x0f, 0x16, 0x7f, 0x86, 0x8e,
	0x4c, 0x9f, 0xca, 0xf0, 0x2d, 0x43, 0x3d, 0xf9, 0xfa, 0xf8, 0x49, 0x7e, 0x2b, 0xbf, 0x0d, 0x45,
	0xb3, 0x4a, 0x39, 0xf8, 0x47, 0x83, 0x76, 0x39, 0x06, 0x22, 0xd4, 0xa3, 0x70, 0x27, 0xd4, 0x17,
	0xd1, 0x61, 0xea, 0x8c, 0x17, 0xd0, 0x48, 0x17, 0x71, 0x92, 0x7f, 0x06, 0x26, 0xcb, 0x2f, 0xf8,
	0x35, 0x74, 0x12, 0xb1, 0xe4, 0x8b, 0x30, 0x59, 0xa6, 0x56, 0x4d, 0x31, 0xed, 0x44, 0x2c, 0x5d,
	0x79, 0xc7, 0x11, 0xf4, 0xf3, 0xe5, 0x29, 0x3e, 0x8f, 0xc4, 0xaa, 0x0f, 0x6b, 0xb6, 0x79, 0xa3,
	0x13, 0x8d, 0xf5, 0x72, 0x52, 0x6a, 0x55, 0x2a, 0xf8, 0x0d, 0x74, 0x8f, 0xf4, 0x69, 0x11, 0xb2,
	0x51, 0xc9, 0x52, 0x1c, 0x40, 0x5b, 0xda, 0xc4, 0x87, 0x2c, 0x55, 0xc1, 0x9a, 0xec, 0xcb, 0x5d,
	0x96, 0x17, 0x67, 0xf5, 0x94, 0x8a, 0xd5, 0x64, 0x46, 0x81, 0xc9, 0x27, 0xf0, 0x0a, 0x9a, 0xeb,
	0x38, 0xdc, 0x6e, 0xf2, 0x18, 0x4d, 0x56, 0xdc, 0x06, 0xdf, 0x42, 0x43, 0x25, 0x81, 0x5d, 0xd0,
	0x3e, 0xab, 0xb1, 0x75, 0xa6, 0x7d, 0x96, 0xb7, 0x47, 0x35, 0xaf, 0xce, 0xb4, 0xc7, 0xeb, 0xbf,
	0x75, 0x68, 0xa8, 0xef, 0x1d, 0xbf, 0x82, 0xcb, 0xa9, 0xcf, 0x1e, 0x1c, 0x8f, 0x8f, 0x27, 0x6c,
	0x1e, 0xf0, 0x3b, 0xc7, 0x1b, 0xf3, 0x19, 0xa3, 0xe4, 0x15, 0x5e, 0x42, 0xff, 0x19, 0x45, 0x34,
	0xbc, 0x00, 0x52, 0xc0, 0x4a, 0x1b, 0x4c, 0x1e, 0x28, 0xd1, 0x71, 0x00, 0x57, 0x05, 0x3a, 0xa7,
	0xae, 0x3f, 0x7d, 0x57, 0x19, 0xd5, 0xf0, 0x0a, 0xf0, 0x39, 0x47, 0xea, 0xd2, 0x89, 0xfe, 0x16,
	0x30, 0x47, 0x79, 0xf0, 0x1b, 0x46, 0x9d, 0x7b, 0xd2, 0x40, 0x0b, 0x2e, 0x72, 0xf4, 0xa4, 0xa1,
	0x66, 0xa5, 0x3f, 0xea, 0xa7, 0x85, 0xe7, 0xd0, 0xcb, 0xd1, 0xaa, 0x9d, 0xb6, 0x1c, 0x2b, 0x07,
	0x4f, 0xbb, 0xe9, 0xc8, 0xb1, 0x9e, 0x51, 0x04, 0xe4, 0x00, 0x33, 0x3a, 0x75, 0xbc, 0xe0, 0x03,
	0x9f, 0xdf, 0xf9, 0x7e, 0xe0, 0xbf, 0x0f, 0x8a, 0x96, 0x0c, 0xf9, 0xf0, 0x29, 0x47, 0xba, 0x68,
	0x42, 0x67, 0xe6, 0xcf, 0x03, 0x7e, 0xeb, 0x3c, 0x50, 0x62, 0x5e, 0xff, 0x59, 0x83, 0x56, 0xf1,
	0xbf, 0x63, 0x1b, 0xea, 0x77, 0x8e, 0x17, 0x90, 0x57, 0xf2, 0x34, 0x0f, 0xfc, 0x19, 0xd1, 0x90,
	0x40, 0xb7, 0xdc, 0x42, 0xe0, 0xb0, 0x80, 0xe8, 0xd8, 0x03, 0x63, 0xec, 0x33, 0x97, 0x16, 0x40,
	0x4d, 0xf5, 0xc0, 0xe8, 0xcc, 0x61, 0x94, 0xdf, 0x4f, 0xdc, 0x7b, 0x7f, 0x3c, 0xe6, 0x1f, 0xa8,
	0xe7, 0xf9, 0xbf, 0x92, 0xba, 0x5c, 0xcb, 0x29, 0x77, 0xe3, 0xbd, 0xa7, 0xa4, 0x71, 0x5c, 0x55,
	0x76, 0x59, 0x54, 0x35, 0x8f, 0xab, 0x4a, 0x4e, 0x55, 0xb5, 0x64, 0x28, 0xef, 0x26, 0x8c, 0xba,
	0x01, 0x1f, 0x33, 0x4a, 0xcb, 0x8a, 0xb6, 0x9c, 0xf5, 0x18, 0x57, 0xea, 0x8e, 0xf4, 0x99, 0x4c,
	0x5f, 0xd0, 0x83, 0xf4, 0x79, 0xca, 0xa8, 0x0a, 0x03, 0x11, 0xce, 0x64, 0x16, 0x72, 0x8d, 0x85,
	0xb6, 0x2b, 0x57, 0x50, 0x62, 0x4a, 0x65, 0xca, 0x15, 0xdc, 0xfa, 0x8e, 0x57, 0x4a, 0xce, 0xe4,
	0x52, 0x15, 0xa0, 0xf8, 0x9e, 0xcc, 0xf1, 0xc6, 0xf1, 0x3c, 0x3e, 0xf3, 0x1c, 0x97, 0x3e, 0xd0,
	0xe9, 0x17, 0x33, 0x82, 0xaf, 0xe1, 0xfc, 0x84, 0x52, 0x35, 0xfd, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x18, 0xdc, 0xbe, 0x6c, 0x06, 0x00, 0x00,
}
