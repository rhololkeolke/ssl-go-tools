// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages_robocup_ssl_game_event.proto

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

type SSL_Referee_Game_Event_GameEventType int32

const (
	// not set
	SSL_Referee_Game_Event_UNKNOWN SSL_Referee_Game_Event_GameEventType = 0
	// an event that is not listed in this enum yet.
	// Give further details in the message below
	SSL_Referee_Game_Event_CUSTOM SSL_Referee_Game_Event_GameEventType = 1
	// Law 3: Number of players
	SSL_Referee_Game_Event_NUMBER_OF_PLAYERS SSL_Referee_Game_Event_GameEventType = 2
	// Law 9: Ball out of play
	SSL_Referee_Game_Event_BALL_LEFT_FIELD SSL_Referee_Game_Event_GameEventType = 3
	// Law 10: Team scored a goal
	SSL_Referee_Game_Event_GOAL SSL_Referee_Game_Event_GameEventType = 4
	// Law 9.3: lack of progress while bringing the ball into play
	SSL_Referee_Game_Event_KICK_TIMEOUT SSL_Referee_Game_Event_GameEventType = 5
	// Law ?: There is no progress in game for (10|15)? seconds
	SSL_Referee_Game_Event_NO_PROGRESS_IN_GAME SSL_Referee_Game_Event_GameEventType = 6
	// Law 12: Pushing / Substantial Contact
	SSL_Referee_Game_Event_BOT_COLLISION SSL_Referee_Game_Event_GameEventType = 7
	// Law 12.2: Defender is completely inside penalty area
	SSL_Referee_Game_Event_MULTIPLE_DEFENDER SSL_Referee_Game_Event_GameEventType = 8
	// Law 12: Defender is partially inside penalty area
	SSL_Referee_Game_Event_MULTIPLE_DEFENDER_PARTIALLY SSL_Referee_Game_Event_GameEventType = 9
	// Law 12.3: Attacker in defense area
	SSL_Referee_Game_Event_ATTACKER_IN_DEFENSE_AREA SSL_Referee_Game_Event_GameEventType = 10
	// Law 12: Icing (kicking over midline and opponent goal line)
	SSL_Referee_Game_Event_ICING SSL_Referee_Game_Event_GameEventType = 11
	// Law 12: Ball speed
	SSL_Referee_Game_Event_BALL_SPEED SSL_Referee_Game_Event_GameEventType = 12
	// Law 12: Robot speed during STOP
	SSL_Referee_Game_Event_ROBOT_STOP_SPEED SSL_Referee_Game_Event_GameEventType = 13
	// Law 12: Maximum dribbling distance
	SSL_Referee_Game_Event_BALL_DRIBBLING SSL_Referee_Game_Event_GameEventType = 14
	// Law 12: Touching the opponent goalkeeper
	SSL_Referee_Game_Event_ATTACKER_TOUCH_KEEPER SSL_Referee_Game_Event_GameEventType = 15
	// Law 12: Double touch
	SSL_Referee_Game_Event_DOUBLE_TOUCH SSL_Referee_Game_Event_GameEventType = 16
	// Law 13-17: Attacker not too close to the opponent's penalty area when ball enters play
	SSL_Referee_Game_Event_ATTACKER_TO_DEFENCE_AREA SSL_Referee_Game_Event_GameEventType = 17
	// Law 13-17: Keeping the correct distance to the ball during opponents freekicks
	SSL_Referee_Game_Event_DEFENDER_TO_KICK_POINT_DISTANCE SSL_Referee_Game_Event_GameEventType = 18
	// Law 12: A robot holds the ball deliberately
	SSL_Referee_Game_Event_BALL_HOLDING SSL_Referee_Game_Event_GameEventType = 19
	// Law 12: The ball entered the goal directly after an indirect kick was performed
	SSL_Referee_Game_Event_INDIRECT_GOAL SSL_Referee_Game_Event_GameEventType = 20
	// Law 9.2: Ball placement
	SSL_Referee_Game_Event_BALL_PLACEMENT_FAILED SSL_Referee_Game_Event_GameEventType = 21
	// Law 10: A goal is only scored if the ball has not exceeded a robot height (150mm) between the last
	// kick of an attacker and the time the ball crossed the goal line.
	SSL_Referee_Game_Event_CHIP_ON_GOAL SSL_Referee_Game_Event_GameEventType = 22
)

var SSL_Referee_Game_Event_GameEventType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CUSTOM",
	2:  "NUMBER_OF_PLAYERS",
	3:  "BALL_LEFT_FIELD",
	4:  "GOAL",
	5:  "KICK_TIMEOUT",
	6:  "NO_PROGRESS_IN_GAME",
	7:  "BOT_COLLISION",
	8:  "MULTIPLE_DEFENDER",
	9:  "MULTIPLE_DEFENDER_PARTIALLY",
	10: "ATTACKER_IN_DEFENSE_AREA",
	11: "ICING",
	12: "BALL_SPEED",
	13: "ROBOT_STOP_SPEED",
	14: "BALL_DRIBBLING",
	15: "ATTACKER_TOUCH_KEEPER",
	16: "DOUBLE_TOUCH",
	17: "ATTACKER_TO_DEFENCE_AREA",
	18: "DEFENDER_TO_KICK_POINT_DISTANCE",
	19: "BALL_HOLDING",
	20: "INDIRECT_GOAL",
	21: "BALL_PLACEMENT_FAILED",
	22: "CHIP_ON_GOAL",
}
var SSL_Referee_Game_Event_GameEventType_value = map[string]int32{
	"UNKNOWN":                         0,
	"CUSTOM":                          1,
	"NUMBER_OF_PLAYERS":               2,
	"BALL_LEFT_FIELD":                 3,
	"GOAL":                            4,
	"KICK_TIMEOUT":                    5,
	"NO_PROGRESS_IN_GAME":             6,
	"BOT_COLLISION":                   7,
	"MULTIPLE_DEFENDER":               8,
	"MULTIPLE_DEFENDER_PARTIALLY":     9,
	"ATTACKER_IN_DEFENSE_AREA":        10,
	"ICING":                           11,
	"BALL_SPEED":                      12,
	"ROBOT_STOP_SPEED":                13,
	"BALL_DRIBBLING":                  14,
	"ATTACKER_TOUCH_KEEPER":           15,
	"DOUBLE_TOUCH":                    16,
	"ATTACKER_TO_DEFENCE_AREA":        17,
	"DEFENDER_TO_KICK_POINT_DISTANCE": 18,
	"BALL_HOLDING":                    19,
	"INDIRECT_GOAL":                   20,
	"BALL_PLACEMENT_FAILED":           21,
	"CHIP_ON_GOAL":                    22,
}

func (x SSL_Referee_Game_Event_GameEventType) Enum() *SSL_Referee_Game_Event_GameEventType {
	p := new(SSL_Referee_Game_Event_GameEventType)
	*p = x
	return p
}
func (x SSL_Referee_Game_Event_GameEventType) String() string {
	return proto.EnumName(SSL_Referee_Game_Event_GameEventType_name, int32(x))
}
func (x *SSL_Referee_Game_Event_GameEventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Game_Event_GameEventType_value, data, "SSL_Referee_Game_Event_GameEventType")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Game_Event_GameEventType(value)
	return nil
}
func (SSL_Referee_Game_Event_GameEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a, []int{0, 0}
}

// a team
type SSL_Referee_Game_Event_Team int32

const (
	SSL_Referee_Game_Event_TEAM_UNKNOWN SSL_Referee_Game_Event_Team = 0
	SSL_Referee_Game_Event_TEAM_YELLOW  SSL_Referee_Game_Event_Team = 1
	SSL_Referee_Game_Event_TEAM_BLUE    SSL_Referee_Game_Event_Team = 2
)

var SSL_Referee_Game_Event_Team_name = map[int32]string{
	0: "TEAM_UNKNOWN",
	1: "TEAM_YELLOW",
	2: "TEAM_BLUE",
}
var SSL_Referee_Game_Event_Team_value = map[string]int32{
	"TEAM_UNKNOWN": 0,
	"TEAM_YELLOW":  1,
	"TEAM_BLUE":    2,
}

func (x SSL_Referee_Game_Event_Team) Enum() *SSL_Referee_Game_Event_Team {
	p := new(SSL_Referee_Game_Event_Team)
	*p = x
	return p
}
func (x SSL_Referee_Game_Event_Team) String() string {
	return proto.EnumName(SSL_Referee_Game_Event_Team_name, int32(x))
}
func (x *SSL_Referee_Game_Event_Team) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSL_Referee_Game_Event_Team_value, data, "SSL_Referee_Game_Event_Team")
	if err != nil {
		return err
	}
	*x = SSL_Referee_Game_Event_Team(value)
	return nil
}
func (SSL_Referee_Game_Event_Team) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a, []int{0, 1}
}

// a game event that caused a referee command
type SSL_Referee_Game_Event struct {
	// the game event type that happened
	GameEventType *SSL_Referee_Game_Event_GameEventType `protobuf:"varint,1,req,name=gameEventType,enum=SSL_Referee_Game_Event_GameEventType" json:"gameEventType,omitempty"`
	// the team and optionally a designated robot that is the originator of the game event
	Originator *SSL_Referee_Game_Event_Originator `protobuf:"bytes,2,opt,name=originator" json:"originator,omitempty"`
	// a message describing further details of this game event
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_Referee_Game_Event) Reset()         { *m = SSL_Referee_Game_Event{} }
func (m *SSL_Referee_Game_Event) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_Game_Event) ProtoMessage()    {}
func (*SSL_Referee_Game_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a, []int{0}
}
func (m *SSL_Referee_Game_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_Game_Event.Unmarshal(m, b)
}
func (m *SSL_Referee_Game_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_Game_Event.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_Game_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_Game_Event.Merge(dst, src)
}
func (m *SSL_Referee_Game_Event) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_Game_Event.Size(m)
}
func (m *SSL_Referee_Game_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_Game_Event.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_Game_Event proto.InternalMessageInfo

func (m *SSL_Referee_Game_Event) GetGameEventType() SSL_Referee_Game_Event_GameEventType {
	if m != nil && m.GameEventType != nil {
		return *m.GameEventType
	}
	return SSL_Referee_Game_Event_UNKNOWN
}

func (m *SSL_Referee_Game_Event) GetOriginator() *SSL_Referee_Game_Event_Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

func (m *SSL_Referee_Game_Event) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// information about an originator
type SSL_Referee_Game_Event_Originator struct {
	Team                 *SSL_Referee_Game_Event_Team `protobuf:"varint,1,req,name=team,enum=SSL_Referee_Game_Event_Team" json:"team,omitempty"`
	BotId                *uint32                      `protobuf:"varint,2,opt,name=botId" json:"botId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SSL_Referee_Game_Event_Originator) Reset()         { *m = SSL_Referee_Game_Event_Originator{} }
func (m *SSL_Referee_Game_Event_Originator) String() string { return proto.CompactTextString(m) }
func (*SSL_Referee_Game_Event_Originator) ProtoMessage()    {}
func (*SSL_Referee_Game_Event_Originator) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a, []int{0, 0}
}
func (m *SSL_Referee_Game_Event_Originator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_Referee_Game_Event_Originator.Unmarshal(m, b)
}
func (m *SSL_Referee_Game_Event_Originator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_Referee_Game_Event_Originator.Marshal(b, m, deterministic)
}
func (dst *SSL_Referee_Game_Event_Originator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_Referee_Game_Event_Originator.Merge(dst, src)
}
func (m *SSL_Referee_Game_Event_Originator) XXX_Size() int {
	return xxx_messageInfo_SSL_Referee_Game_Event_Originator.Size(m)
}
func (m *SSL_Referee_Game_Event_Originator) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_Referee_Game_Event_Originator.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_Referee_Game_Event_Originator proto.InternalMessageInfo

func (m *SSL_Referee_Game_Event_Originator) GetTeam() SSL_Referee_Game_Event_Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return SSL_Referee_Game_Event_TEAM_UNKNOWN
}

func (m *SSL_Referee_Game_Event_Originator) GetBotId() uint32 {
	if m != nil && m.BotId != nil {
		return *m.BotId
	}
	return 0
}

func init() {
	proto.RegisterType((*SSL_Referee_Game_Event)(nil), "SSL_Referee_Game_Event")
	proto.RegisterType((*SSL_Referee_Game_Event_Originator)(nil), "SSL_Referee_Game_Event.Originator")
	proto.RegisterEnum("SSL_Referee_Game_Event_GameEventType", SSL_Referee_Game_Event_GameEventType_name, SSL_Referee_Game_Event_GameEventType_value)
	proto.RegisterEnum("SSL_Referee_Game_Event_Team", SSL_Referee_Game_Event_Team_name, SSL_Referee_Game_Event_Team_value)
}

func init() {
	proto.RegisterFile("messages_robocup_ssl_game_event.proto", fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a)
}

var fileDescriptor_messages_robocup_ssl_game_event_078d80496a5b9d6a = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0xc7, 0x9f, 0x84, 0xf0, 0x92, 0x09, 0x09, 0xcb, 0xf0, 0xf2, 0xa4, 0x2d, 0x12, 0x28, 0x15,
	0x12, 0x27, 0x54, 0x71, 0xea, 0x75, 0x6d, 0x4f, 0xc2, 0x2a, 0xeb, 0x5d, 0x6b, 0xbd, 0x16, 0xe2,
	0x34, 0x0a, 0xad, 0x8b, 0x90, 0x1a, 0x8c, 0x92, 0xb4, 0x52, 0x3f, 0x65, 0x3f, 0x4f, 0x6f, 0xd5,
	0x3a, 0x29, 0x2f, 0x6a, 0x39, 0xee, 0xec, 0xef, 0xbf, 0xf3, 0x9b, 0xb1, 0xe1, 0x74, 0x5a, 0xce,
	0xe7, 0x93, 0xdb, 0x72, 0xce, 0xb3, 0xea, 0xa6, 0xfa, 0xf4, 0xed, 0x81, 0xe7, 0xf3, 0xaf, 0x7c,
	0x3b, 0x99, 0x96, 0x5c, 0x7e, 0x2f, 0xef, 0x17, 0xe7, 0x0f, 0xb3, 0x6a, 0x51, 0x0d, 0x7e, 0x6e,
	0xc0, 0x61, 0x9e, 0x6b, 0x76, 0xe5, 0x97, 0x72, 0x56, 0x96, 0x3c, 0x0a, 0x00, 0x05, 0x00, 0xc7,
	0xd0, 0x0d, 0x78, 0x7d, 0xf0, 0x3f, 0x1e, 0xca, 0x7e, 0xe3, 0xa4, 0x79, 0xd6, 0xbb, 0x38, 0x3d,
	0xff, 0x37, 0x7f, 0x3e, 0x7a, 0x0e, 0xbb, 0x97, 0x59, 0x8c, 0x00, 0xaa, 0xd9, 0xdd, 0xed, 0xdd,
	0xfd, 0x64, 0x51, 0xcd, 0xfa, 0xcd, 0x93, 0xc6, 0x59, 0xe7, 0x62, 0xf0, 0xda, 0x4b, 0xf6, 0x91,
	0x74, 0xcf, 0x52, 0xd8, 0x87, 0xcd, 0xd5, 0x50, 0xfd, 0xb5, 0x93, 0xc6, 0x59, 0xdb, 0xfd, 0x39,
	0xbe, 0xf5, 0x00, 0x4f, 0x19, 0xfc, 0x00, 0xad, 0x45, 0x39, 0x99, 0xae, 0x7c, 0x8f, 0x5e, 0xeb,
	0xe2, 0xcb, 0xc9, 0xd4, 0xd5, 0x24, 0xee, 0xc3, 0xfa, 0x4d, 0xb5, 0x50, 0x9f, 0x6b, 0xb1, 0xae,
	0x5b, 0x1e, 0x06, 0xbf, 0xd6, 0xa0, 0xfb, 0x62, 0x28, 0xec, 0xc0, 0x66, 0x61, 0xc6, 0xc6, 0x5e,
	0x19, 0xf1, 0x1f, 0x02, 0x6c, 0xc4, 0x45, 0xee, 0x6d, 0x2a, 0x1a, 0x78, 0x00, 0xbb, 0xa6, 0x48,
	0x23, 0x72, 0x6c, 0x87, 0x9c, 0x69, 0x79, 0x4d, 0x2e, 0x17, 0x4d, 0xdc, 0x83, 0x9d, 0x48, 0x6a,
	0xcd, 0x9a, 0x86, 0x9e, 0x87, 0x8a, 0x74, 0x22, 0xd6, 0x70, 0x0b, 0x5a, 0x23, 0x2b, 0xb5, 0x68,
	0xa1, 0x80, 0xed, 0xb1, 0x8a, 0xc7, 0xec, 0x55, 0x4a, 0xb6, 0xf0, 0x62, 0x1d, 0xff, 0x87, 0x3d,
	0x63, 0x39, 0x73, 0x76, 0xe4, 0x28, 0xcf, 0x59, 0x19, 0x1e, 0xc9, 0x94, 0xc4, 0x06, 0xee, 0x42,
	0x37, 0xb2, 0x9e, 0x63, 0xab, 0xb5, 0xca, 0x95, 0x35, 0x62, 0x33, 0xf4, 0x4c, 0x0b, 0xed, 0x55,
	0xa6, 0x89, 0x13, 0x1a, 0x92, 0x49, 0xc8, 0x89, 0x2d, 0x3c, 0x86, 0x77, 0x7f, 0x95, 0x39, 0x93,
	0xce, 0x2b, 0xa9, 0xf5, 0xb5, 0x68, 0xe3, 0x11, 0xf4, 0xa5, 0xf7, 0x32, 0x1e, 0x93, 0x0b, 0x0d,
	0x6a, 0x26, 0x27, 0x96, 0x8e, 0xa4, 0x00, 0x6c, 0xc3, 0xba, 0x8a, 0x95, 0x19, 0x89, 0x0e, 0xf6,
	0x00, 0x6a, 0xfb, 0x3c, 0x23, 0x4a, 0xc4, 0x36, 0xee, 0x83, 0x70, 0x36, 0x58, 0xe4, 0xde, 0x66,
	0xab, 0x6a, 0x17, 0x11, 0x7a, 0x35, 0x95, 0x38, 0x15, 0x45, 0x3a, 0x24, 0x7b, 0xf8, 0x06, 0x0e,
	0x1e, 0x5b, 0x78, 0x5b, 0xc4, 0x97, 0x3c, 0x26, 0xca, 0xc8, 0x89, 0x9d, 0x30, 0x73, 0x62, 0x8b,
	0x48, 0xd3, 0xf2, 0x42, 0x88, 0x17, 0x3e, 0xde, 0x2e, 0x7d, 0xe2, 0x95, 0xcf, 0x2e, 0xbe, 0x87,
	0xe3, 0xc7, 0x29, 0xbc, 0xe5, 0x7a, 0x5f, 0x99, 0x55, 0xc6, 0x73, 0xa2, 0x72, 0x2f, 0x4d, 0x4c,
	0x02, 0xc3, 0xa3, 0xb5, 0xc3, 0xa5, 0xd5, 0x49, 0x30, 0xd8, 0x0b, 0xfb, 0x52, 0x26, 0x51, 0x8e,
	0x62, 0xcf, 0xf5, 0xb6, 0xf7, 0x83, 0x54, 0x0d, 0x65, 0x5a, 0xc6, 0x94, 0x92, 0xf1, 0x3c, 0x94,
	0x4a, 0x53, 0x22, 0x0e, 0x42, 0x3e, 0xbe, 0x54, 0x19, 0x5b, 0xb3, 0x84, 0x0f, 0x07, 0x1f, 0xa1,
	0x15, 0xfe, 0x8f, 0x70, 0xe3, 0x49, 0xa6, 0xfc, 0xf4, 0xd9, 0x77, 0xa0, 0x53, 0x57, 0xae, 0x49,
	0x6b, 0x7b, 0x25, 0x1a, 0xd8, 0x85, 0x76, 0x5d, 0x88, 0x74, 0x41, 0xa2, 0xf9, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0xb1, 0x7c, 0xd6, 0x79, 0x03, 0x00, 0x00,
}
