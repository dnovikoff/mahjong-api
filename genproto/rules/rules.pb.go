// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rules/rules.proto

package rules

import (
	fmt "fmt"
	base "github.com/dnovikoff/mahjong-api/genproto/base"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KanDoraOpen int32

const (
	KanDoraOpen_KANDORA_OPEN_UNSPECEFIED KanDoraOpen = 0
	// No kan dora open in JPML-A
	KanDoraOpen_DONT_OPEN KanDoraOpen = 1
	// Instant dora open in EMA
	KanDoraOpen_INSTANT KanDoraOpen = 2
	// Afer next kan call or drop (for opened kans)
	KanDoraOpen_AFTER_ACTION KanDoraOpen = 3
)

var KanDoraOpen_name = map[int32]string{
	0: "KANDORA_OPEN_UNSPECEFIED",
	1: "DONT_OPEN",
	2: "INSTANT",
	3: "AFTER_ACTION",
}

var KanDoraOpen_value = map[string]int32{
	"KANDORA_OPEN_UNSPECEFIED": 0,
	"DONT_OPEN":                1,
	"INSTANT":                  2,
	"AFTER_ACTION":             3,
}

func (x KanDoraOpen) String() string {
	return proto.EnumName(KanDoraOpen_name, int32(x))
}

func (KanDoraOpen) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{0}
}

type Atodzuke int32

const (
	Atodzuke_ATODZUKE_UNSPECIFIED Atodzuke = 0
	Atodzuke_ATODZUKE_ALLOWED     Atodzuke = 1
	// Allowed when all waits gives some yaku.
	Atodzuke_ATODZUKE_FORBIDDEN_SOFT Atodzuke = 2
	// Should already have yaku. Eg. syanpon for both yakuhai does not fit.
	Atodzuke_ATODZUKE_FORBIDDEN_STRICT Atodzuke = 3
)

var Atodzuke_name = map[int32]string{
	0: "ATODZUKE_UNSPECIFIED",
	1: "ATODZUKE_ALLOWED",
	2: "ATODZUKE_FORBIDDEN_SOFT",
	3: "ATODZUKE_FORBIDDEN_STRICT",
}

var Atodzuke_value = map[string]int32{
	"ATODZUKE_UNSPECIFIED":      0,
	"ATODZUKE_ALLOWED":          1,
	"ATODZUKE_FORBIDDEN_SOFT":   2,
	"ATODZUKE_FORBIDDEN_STRICT": 3,
}

func (x Atodzuke) String() string {
	return proto.EnumName(Atodzuke_name, int32(x))
}

func (Atodzuke) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{1}
}

type AgariYame int32

const (
	AgariYame_AGARI_YAME_UNSPECIFIED AgariYame = 0
	// Game ends.
	AgariYame_IMPLICIT_END AgariYame = 1
	// Game might be continued (player asked). (TODO: Implement)
	AgariYame_EXPLICIT_CONTINUE AgariYame = 2
	// Game must be continued.
	AgariYame_IMPLICIT_CONTINUE AgariYame = 3
)

var AgariYame_name = map[int32]string{
	0: "AGARI_YAME_UNSPECIFIED",
	1: "IMPLICIT_END",
	2: "EXPLICIT_CONTINUE",
	3: "IMPLICIT_CONTINUE",
}

var AgariYame_value = map[string]int32{
	"AGARI_YAME_UNSPECIFIED": 0,
	"IMPLICIT_END":           1,
	"EXPLICIT_CONTINUE":      2,
	"IMPLICIT_CONTINUE":      3,
}

func (x AgariYame) String() string {
	return proto.EnumName(AgariYame_name, int32(x))
}

func (AgariYame) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{2}
}

type Shifting int32

const (
	Shifting_SHIFTING_UNSPECIFIED Shifting = 0
	Shifting_SHIFTING_ALLOWED     Shifting = 1
	// Only shifting alloed. Ex. for 34+5, 2=allowed, 5=forbidden
	Shifting_SHIFTING_FORBIDDEN_SOFT Shifting = 2
	// for 24+5 both 2 and 5 forbidden
	Shifting_SHIFTING_FORBIDDEN_STRICT Shifting = 3
)

var Shifting_name = map[int32]string{
	0: "SHIFTING_UNSPECIFIED",
	1: "SHIFTING_ALLOWED",
	2: "SHIFTING_FORBIDDEN_SOFT",
	3: "SHIFTING_FORBIDDEN_STRICT",
}

var Shifting_value = map[string]int32{
	"SHIFTING_UNSPECIFIED":      0,
	"SHIFTING_ALLOWED":          1,
	"SHIFTING_FORBIDDEN_SOFT":   2,
	"SHIFTING_FORBIDDEN_STRICT": 3,
}

func (x Shifting) String() string {
	return proto.EnumName(Shifting_name, int32(x))
}

func (Shifting) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{3}
}

type Yaku struct {
	OpenTanyao bool `protobuf:"varint,1,opt,name=open_tanyao,json=openTanyao,proto3" json:"open_tanyao,omitempty"`
	// Does rinshan combine with haitei
	HaiteiFromLiveOnly bool `protobuf:"varint,2,opt,name=haitei_from_live_only,json=haiteiFromLiveOnly,proto3" json:"haitei_from_live_only,omitempty"`
	Ura                bool `protobuf:"varint,3,opt,name=ura,proto3" json:"ura,omitempty"`
	// No ipatsu in JPML-A
	Ipatsu  bool            `protobuf:"varint,4,opt,name=ipatsu,proto3" json:"ipatsu,omitempty"`
	AkaDora *base.Instances `protobuf:"bytes,5,opt,name=aka_dora,json=akaDora,proto3" json:"aka_dora,omitempty"`
	Renhou  base.Limit      `protobuf:"varint,6,opt,name=renhou,proto3,enum=base.Limit" json:"renhou,omitempty"`
	// No tsumo 2 fu for rinshan in jpmla
	RinshanFu bool `protobuf:"varint,7,opt,name=rinshan_fu,json=rinshanFu,proto3" json:"rinshan_fu,omitempty"`
	// Green required for ryuuiisou yakuman in jpmla
	GreenRequired        bool     `protobuf:"varint,8,opt,name=green_required,json=greenRequired,proto3" json:"green_required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Yaku) Reset()         { *m = Yaku{} }
func (m *Yaku) String() string { return proto.CompactTextString(m) }
func (*Yaku) ProtoMessage()    {}
func (*Yaku) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{0}
}

func (m *Yaku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Yaku.Unmarshal(m, b)
}
func (m *Yaku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Yaku.Marshal(b, m, deterministic)
}
func (m *Yaku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Yaku.Merge(m, src)
}
func (m *Yaku) XXX_Size() int {
	return xxx_messageInfo_Yaku.Size(m)
}
func (m *Yaku) XXX_DiscardUnknown() {
	xxx_messageInfo_Yaku.DiscardUnknown(m)
}

var xxx_messageInfo_Yaku proto.InternalMessageInfo

func (m *Yaku) GetOpenTanyao() bool {
	if m != nil {
		return m.OpenTanyao
	}
	return false
}

func (m *Yaku) GetHaiteiFromLiveOnly() bool {
	if m != nil {
		return m.HaiteiFromLiveOnly
	}
	return false
}

func (m *Yaku) GetUra() bool {
	if m != nil {
		return m.Ura
	}
	return false
}

func (m *Yaku) GetIpatsu() bool {
	if m != nil {
		return m.Ipatsu
	}
	return false
}

func (m *Yaku) GetAkaDora() *base.Instances {
	if m != nil {
		return m.AkaDora
	}
	return nil
}

func (m *Yaku) GetRenhou() base.Limit {
	if m != nil {
		return m.Renhou
	}
	return base.Limit_LIMIT_UNSPECIFIED
}

func (m *Yaku) GetRinshanFu() bool {
	if m != nil {
		return m.RinshanFu
	}
	return false
}

func (m *Yaku) GetGreenRequired() bool {
	if m != nil {
		return m.GreenRequired
	}
	return false
}

type Scoring struct {
	// 7700 rounds to 8000
	ManganRound  bool `protobuf:"varint,1,opt,name=mangan_round,json=manganRound,proto3" json:"mangan_round,omitempty"`
	KazoeYakuman bool `protobuf:"varint,2,opt,name=kazoe_yakuman,json=kazoeYakuman,proto3" json:"kazoe_yakuman,omitempty"`
	YakumanSum   bool `protobuf:"varint,3,opt,name=yakuman_sum,json=yakumanSum,proto3" json:"yakuman_sum,omitempty"`
	// The default should be 100, but could differ in some rules (eg. 500)
	HonbaMoney           int64          `protobuf:"varint,4,opt,name=honba_money,json=honbaMoney,proto3" json:"honba_money,omitempty"`
	DoubleYakumans       []base.Yakuman `protobuf:"varint,5,rep,packed,name=double_yakumans,json=doubleYakumans,proto3,enum=base.Yakuman" json:"double_yakumans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Scoring) Reset()         { *m = Scoring{} }
func (m *Scoring) String() string { return proto.CompactTextString(m) }
func (*Scoring) ProtoMessage()    {}
func (*Scoring) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{1}
}

func (m *Scoring) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scoring.Unmarshal(m, b)
}
func (m *Scoring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scoring.Marshal(b, m, deterministic)
}
func (m *Scoring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scoring.Merge(m, src)
}
func (m *Scoring) XXX_Size() int {
	return xxx_messageInfo_Scoring.Size(m)
}
func (m *Scoring) XXX_DiscardUnknown() {
	xxx_messageInfo_Scoring.DiscardUnknown(m)
}

var xxx_messageInfo_Scoring proto.InternalMessageInfo

func (m *Scoring) GetManganRound() bool {
	if m != nil {
		return m.ManganRound
	}
	return false
}

func (m *Scoring) GetKazoeYakuman() bool {
	if m != nil {
		return m.KazoeYakuman
	}
	return false
}

func (m *Scoring) GetYakumanSum() bool {
	if m != nil {
		return m.YakumanSum
	}
	return false
}

func (m *Scoring) GetHonbaMoney() int64 {
	if m != nil {
		return m.HonbaMoney
	}
	return 0
}

func (m *Scoring) GetDoubleYakumans() []base.Yakuman {
	if m != nil {
		return m.DoubleYakumans
	}
	return nil
}

type Draw struct {
	Winds                bool     `protobuf:"varint,1,opt,name=winds,proto3" json:"winds,omitempty"`
	Kokushi              bool     `protobuf:"varint,2,opt,name=kokushi,proto3" json:"kokushi,omitempty"`
	Kans                 bool     `protobuf:"varint,3,opt,name=kans,proto3" json:"kans,omitempty"`
	Riichi               bool     `protobuf:"varint,4,opt,name=riichi,proto3" json:"riichi,omitempty"`
	Ron3                 bool     `protobuf:"varint,5,opt,name=ron3,proto3" json:"ron3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Draw) Reset()         { *m = Draw{} }
func (m *Draw) String() string { return proto.CompactTextString(m) }
func (*Draw) ProtoMessage()    {}
func (*Draw) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{2}
}

func (m *Draw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Draw.Unmarshal(m, b)
}
func (m *Draw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Draw.Marshal(b, m, deterministic)
}
func (m *Draw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Draw.Merge(m, src)
}
func (m *Draw) XXX_Size() int {
	return xxx_messageInfo_Draw.Size(m)
}
func (m *Draw) XXX_DiscardUnknown() {
	xxx_messageInfo_Draw.DiscardUnknown(m)
}

var xxx_messageInfo_Draw proto.InternalMessageInfo

func (m *Draw) GetWinds() bool {
	if m != nil {
		return m.Winds
	}
	return false
}

func (m *Draw) GetKokushi() bool {
	if m != nil {
		return m.Kokushi
	}
	return false
}

func (m *Draw) GetKans() bool {
	if m != nil {
		return m.Kans
	}
	return false
}

func (m *Draw) GetRiichi() bool {
	if m != nil {
		return m.Riichi
	}
	return false
}

func (m *Draw) GetRon3() bool {
	if m != nil {
		return m.Ron3
	}
	return false
}

type Pao struct {
	Winds                bool     `protobuf:"varint,1,opt,name=winds,proto3" json:"winds,omitempty"`
	Dragons              bool     `protobuf:"varint,2,opt,name=dragons,proto3" json:"dragons,omitempty"`
	Kans                 bool     `protobuf:"varint,3,opt,name=kans,proto3" json:"kans,omitempty"`
	Rinshan              bool     `protobuf:"varint,4,opt,name=rinshan,proto3" json:"rinshan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pao) Reset()         { *m = Pao{} }
func (m *Pao) String() string { return proto.CompactTextString(m) }
func (*Pao) ProtoMessage()    {}
func (*Pao) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{3}
}

func (m *Pao) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pao.Unmarshal(m, b)
}
func (m *Pao) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pao.Marshal(b, m, deterministic)
}
func (m *Pao) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pao.Merge(m, src)
}
func (m *Pao) XXX_Size() int {
	return xxx_messageInfo_Pao.Size(m)
}
func (m *Pao) XXX_DiscardUnknown() {
	xxx_messageInfo_Pao.DiscardUnknown(m)
}

var xxx_messageInfo_Pao proto.InternalMessageInfo

func (m *Pao) GetWinds() bool {
	if m != nil {
		return m.Winds
	}
	return false
}

func (m *Pao) GetDragons() bool {
	if m != nil {
		return m.Dragons
	}
	return false
}

func (m *Pao) GetKans() bool {
	if m != nil {
		return m.Kans
	}
	return false
}

func (m *Pao) GetRinshan() bool {
	if m != nil {
		return m.Rinshan
	}
	return false
}

// Complex uma (binta) is used in JPML-A.
// Still classic uma could be described this way with all 3 field equal.
type ComplexUma struct {
	DefaultMoney         []int64  `protobuf:"zigzag64,1,rep,packed,name=default_money,json=defaultMoney,proto3" json:"default_money,omitempty"`
	Minus1Money          []int64  `protobuf:"zigzag64,2,rep,packed,name=minus1_money,json=minus1Money,proto3" json:"minus1_money,omitempty"`
	Plus1Money           []int64  `protobuf:"zigzag64,3,rep,packed,name=plus1_money,json=plus1Money,proto3" json:"plus1_money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplexUma) Reset()         { *m = ComplexUma{} }
func (m *ComplexUma) String() string { return proto.CompactTextString(m) }
func (*ComplexUma) ProtoMessage()    {}
func (*ComplexUma) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{4}
}

func (m *ComplexUma) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexUma.Unmarshal(m, b)
}
func (m *ComplexUma) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexUma.Marshal(b, m, deterministic)
}
func (m *ComplexUma) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexUma.Merge(m, src)
}
func (m *ComplexUma) XXX_Size() int {
	return xxx_messageInfo_ComplexUma.Size(m)
}
func (m *ComplexUma) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexUma.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexUma proto.InternalMessageInfo

func (m *ComplexUma) GetDefaultMoney() []int64 {
	if m != nil {
		return m.DefaultMoney
	}
	return nil
}

func (m *ComplexUma) GetMinus1Money() []int64 {
	if m != nil {
		return m.Minus1Money
	}
	return nil
}

func (m *ComplexUma) GetPlus1Money() []int64 {
	if m != nil {
		return m.Plus1Money
	}
	return nil
}

type Game struct {
	// Dealer could end the game if first in orasu.
	AgariYame              AgariYame `protobuf:"varint,1,opt,name=agari_yame,json=agariYame,proto3,enum=rules.AgariYame" json:"agari_yame,omitempty"`
	ChiShift               Shifting  `protobuf:"varint,2,opt,name=chi_shift,json=chiShift,proto3,enum=rules.Shifting" json:"chi_shift,omitempty"`
	EndByBancrocity        bool      `protobuf:"varint,3,opt,name=end_by_bancrocity,json=endByBancrocity,proto3" json:"end_by_bancrocity,omitempty"`
	Nagashi                bool      `protobuf:"varint,4,opt,name=nagashi,proto3" json:"nagashi,omitempty"`
	Atamahane              bool      `protobuf:"varint,5,opt,name=atamahane,proto3" json:"atamahane,omitempty"`
	HonbaPayedToAll        bool      `protobuf:"varint,6,opt,name=honba_payed_to_all,json=honbaPayedToAll,proto3" json:"honba_payed_to_all,omitempty"`
	RiichiReturnOnMultiron bool      `protobuf:"varint,7,opt,name=riichi_return_on_multiron,json=riichiReturnOnMultiron,proto3" json:"riichi_return_on_multiron,omitempty"`
	// Chi priority if clicked first.
	SpeedChi bool `protobuf:"varint,8,opt,name=speed_chi,json=speedChi,proto3" json:"speed_chi,omitempty"`
	// Calls displayed, event if waiting for a higher priority choises.
	SayOnClick               bool     `protobuf:"varint,9,opt,name=say_on_click,json=sayOnClick,proto3" json:"say_on_click,omitempty"`
	ShouldHaveMoneyForRiichi bool     `protobuf:"varint,10,opt,name=should_have_money_for_riichi,json=shouldHaveMoneyForRiichi,proto3" json:"should_have_money_for_riichi,omitempty"`
	Atodzuke                 Atodzuke `protobuf:"varint,11,opt,name=atodzuke,proto3,enum=rules.Atodzuke" json:"atodzuke,omitempty"`
	StartMoney               int64    `protobuf:"varint,12,opt,name=start_money,json=startMoney,proto3" json:"start_money,omitempty"`
	// At the end of the game, money reduced by start points
	EndReduceMoney int64 `protobuf:"varint,13,opt,name=end_reduce_money,json=endReduceMoney,proto3" json:"end_reduce_money,omitempty"`
	OkaMoney       int64 `protobuf:"varint,14,opt,name=oka_money,json=okaMoney,proto3" json:"oka_money,omitempty"`
	// Ex. At least 30000 to end the game.
	MinWinMoney int64 `protobuf:"varint,15,opt,name=min_win_money,json=minWinMoney,proto3" json:"min_win_money,omitempty"`
	// Ex. West for hanchans, until min_win_money condition fits.
	MaxLastWind base.Wind `protobuf:"varint,16,opt,name=max_last_wind,json=maxLastWind,proto3,enum=base.Wind" json:"max_last_wind,omitempty"`
	// Ex. East for hanchans. Could end here if conditions meets.
	LastWind    base.Wind   `protobuf:"varint,17,opt,name=last_wind,json=lastWind,proto3,enum=base.Wind" json:"last_wind,omitempty"`
	Uma         *ComplexUma `protobuf:"bytes,18,opt,name=uma,proto3" json:"uma,omitempty"`
	Draw        *Draw       `protobuf:"bytes,19,opt,name=draw,proto3" json:"draw,omitempty"`
	KanDoraOpen KanDoraOpen `protobuf:"varint,20,opt,name=kan_dora_open,json=kanDoraOpen,proto3,enum=rules.KanDoraOpen" json:"kan_dora_open,omitempty"`
	// In case of same points, players share uma.
	UmaShare bool `protobuf:"varint,21,opt,name=uma_share,json=umaShare,proto3" json:"uma_share,omitempty"`
	// In case there are riichi sticks
	RiichiSticksGoesToFirst bool `protobuf:"varint,22,opt,name=riichi_sticks_goes_to_first,json=riichiSticksGoesToFirst,proto3" json:"riichi_sticks_goes_to_first,omitempty"`
	Pao                     *Pao `protobuf:"bytes,23,opt,name=pao,proto3" json:"pao,omitempty"`
	// Possible values: 2,3,4
	NumberOfPlayers int64 `protobuf:"varint,24,opt,name=number_of_players,json=numberOfPlayers,proto3" json:"number_of_players,omitempty"`
	// Not allowed in sanma
	ChiAllowed bool `protobuf:"varint,25,opt,name=chi_allowed,json=chiAllowed,proto3" json:"chi_allowed,omitempty"`
	// User asked if he wants to show noten (not for riichi)
	SuggestNotean        bool     `protobuf:"varint,26,opt,name=suggest_notean,json=suggestNotean,proto3" json:"suggest_notean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{5}
}

func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetAgariYame() AgariYame {
	if m != nil {
		return m.AgariYame
	}
	return AgariYame_AGARI_YAME_UNSPECIFIED
}

func (m *Game) GetChiShift() Shifting {
	if m != nil {
		return m.ChiShift
	}
	return Shifting_SHIFTING_UNSPECIFIED
}

func (m *Game) GetEndByBancrocity() bool {
	if m != nil {
		return m.EndByBancrocity
	}
	return false
}

func (m *Game) GetNagashi() bool {
	if m != nil {
		return m.Nagashi
	}
	return false
}

func (m *Game) GetAtamahane() bool {
	if m != nil {
		return m.Atamahane
	}
	return false
}

func (m *Game) GetHonbaPayedToAll() bool {
	if m != nil {
		return m.HonbaPayedToAll
	}
	return false
}

func (m *Game) GetRiichiReturnOnMultiron() bool {
	if m != nil {
		return m.RiichiReturnOnMultiron
	}
	return false
}

func (m *Game) GetSpeedChi() bool {
	if m != nil {
		return m.SpeedChi
	}
	return false
}

func (m *Game) GetSayOnClick() bool {
	if m != nil {
		return m.SayOnClick
	}
	return false
}

func (m *Game) GetShouldHaveMoneyForRiichi() bool {
	if m != nil {
		return m.ShouldHaveMoneyForRiichi
	}
	return false
}

func (m *Game) GetAtodzuke() Atodzuke {
	if m != nil {
		return m.Atodzuke
	}
	return Atodzuke_ATODZUKE_UNSPECIFIED
}

func (m *Game) GetStartMoney() int64 {
	if m != nil {
		return m.StartMoney
	}
	return 0
}

func (m *Game) GetEndReduceMoney() int64 {
	if m != nil {
		return m.EndReduceMoney
	}
	return 0
}

func (m *Game) GetOkaMoney() int64 {
	if m != nil {
		return m.OkaMoney
	}
	return 0
}

func (m *Game) GetMinWinMoney() int64 {
	if m != nil {
		return m.MinWinMoney
	}
	return 0
}

func (m *Game) GetMaxLastWind() base.Wind {
	if m != nil {
		return m.MaxLastWind
	}
	return base.Wind_WIND_UNSPECIFIED
}

func (m *Game) GetLastWind() base.Wind {
	if m != nil {
		return m.LastWind
	}
	return base.Wind_WIND_UNSPECIFIED
}

func (m *Game) GetUma() *ComplexUma {
	if m != nil {
		return m.Uma
	}
	return nil
}

func (m *Game) GetDraw() *Draw {
	if m != nil {
		return m.Draw
	}
	return nil
}

func (m *Game) GetKanDoraOpen() KanDoraOpen {
	if m != nil {
		return m.KanDoraOpen
	}
	return KanDoraOpen_KANDORA_OPEN_UNSPECEFIED
}

func (m *Game) GetUmaShare() bool {
	if m != nil {
		return m.UmaShare
	}
	return false
}

func (m *Game) GetRiichiSticksGoesToFirst() bool {
	if m != nil {
		return m.RiichiSticksGoesToFirst
	}
	return false
}

func (m *Game) GetPao() *Pao {
	if m != nil {
		return m.Pao
	}
	return nil
}

func (m *Game) GetNumberOfPlayers() int64 {
	if m != nil {
		return m.NumberOfPlayers
	}
	return 0
}

func (m *Game) GetChiAllowed() bool {
	if m != nil {
		return m.ChiAllowed
	}
	return false
}

func (m *Game) GetSuggestNotean() bool {
	if m != nil {
		return m.SuggestNotean
	}
	return false
}

type Ruleset struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Scoring              *Scoring `protobuf:"bytes,3,opt,name=scoring,proto3" json:"scoring,omitempty"`
	Yaku                 *Yaku    `protobuf:"bytes,4,opt,name=yaku,proto3" json:"yaku,omitempty"`
	Game                 *Game    `protobuf:"bytes,5,opt,name=game,proto3" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ruleset) Reset()         { *m = Ruleset{} }
func (m *Ruleset) String() string { return proto.CompactTextString(m) }
func (*Ruleset) ProtoMessage()    {}
func (*Ruleset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{6}
}

func (m *Ruleset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ruleset.Unmarshal(m, b)
}
func (m *Ruleset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ruleset.Marshal(b, m, deterministic)
}
func (m *Ruleset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ruleset.Merge(m, src)
}
func (m *Ruleset) XXX_Size() int {
	return xxx_messageInfo_Ruleset.Size(m)
}
func (m *Ruleset) XXX_DiscardUnknown() {
	xxx_messageInfo_Ruleset.DiscardUnknown(m)
}

var xxx_messageInfo_Ruleset proto.InternalMessageInfo

func (m *Ruleset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ruleset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Ruleset) GetScoring() *Scoring {
	if m != nil {
		return m.Scoring
	}
	return nil
}

func (m *Ruleset) GetYaku() *Yaku {
	if m != nil {
		return m.Yaku
	}
	return nil
}

func (m *Ruleset) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type Timeouts struct {
	// Base time to think on server suggest. Ex. 10 seconds.
	// Max value: 30 seconds.
	// Min value: 5 seconds.
	Base *duration.Duration `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// Additional time to think in this round (if base time passed). Ex. 5 seconds.
	// Max value: 30 seconds.
	// Min value: 0. (or nil)
	Extra *duration.Duration `protobuf:"bytes,2,opt,name=extra,proto3" json:"extra,omitempty"`
	// If player makes a decition in base time, than recover time added to extra.
	// Limited by extra value. Ex. 1 second.
	// Value: >=0 (or nil)
	Recover *duration.Duration `protobuf:"bytes,3,opt,name=recover,proto3" json:"recover,omitempty"`
	// The game will move very fast in case there are no suggest.
	// This is a delay after each player action.
	// Max value: 1 second.
	// Could be 0 (nil) for robot games.
	Delay                *duration.Duration `protobuf:"bytes,4,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Timeouts) Reset()         { *m = Timeouts{} }
func (m *Timeouts) String() string { return proto.CompactTextString(m) }
func (*Timeouts) ProtoMessage()    {}
func (*Timeouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_0addb16c7cb9da8e, []int{7}
}

func (m *Timeouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timeouts.Unmarshal(m, b)
}
func (m *Timeouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timeouts.Marshal(b, m, deterministic)
}
func (m *Timeouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timeouts.Merge(m, src)
}
func (m *Timeouts) XXX_Size() int {
	return xxx_messageInfo_Timeouts.Size(m)
}
func (m *Timeouts) XXX_DiscardUnknown() {
	xxx_messageInfo_Timeouts.DiscardUnknown(m)
}

var xxx_messageInfo_Timeouts proto.InternalMessageInfo

func (m *Timeouts) GetBase() *duration.Duration {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *Timeouts) GetExtra() *duration.Duration {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *Timeouts) GetRecover() *duration.Duration {
	if m != nil {
		return m.Recover
	}
	return nil
}

func (m *Timeouts) GetDelay() *duration.Duration {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterEnum("rules.KanDoraOpen", KanDoraOpen_name, KanDoraOpen_value)
	proto.RegisterEnum("rules.Atodzuke", Atodzuke_name, Atodzuke_value)
	proto.RegisterEnum("rules.AgariYame", AgariYame_name, AgariYame_value)
	proto.RegisterEnum("rules.Shifting", Shifting_name, Shifting_value)
	proto.RegisterType((*Yaku)(nil), "rules.Yaku")
	proto.RegisterType((*Scoring)(nil), "rules.Scoring")
	proto.RegisterType((*Draw)(nil), "rules.Draw")
	proto.RegisterType((*Pao)(nil), "rules.Pao")
	proto.RegisterType((*ComplexUma)(nil), "rules.ComplexUma")
	proto.RegisterType((*Game)(nil), "rules.Game")
	proto.RegisterType((*Ruleset)(nil), "rules.Ruleset")
	proto.RegisterType((*Timeouts)(nil), "rules.Timeouts")
}

func init() { proto.RegisterFile("rules/rules.proto", fileDescriptor_0addb16c7cb9da8e) }

var fileDescriptor_0addb16c7cb9da8e = []byte{
	// 1507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x6e, 0xdb, 0xc8,
	0x15, 0xae, 0x2c, 0x3b, 0x92, 0x8e, 0x6c, 0x59, 0x9e, 0xe6, 0x87, 0xf9, 0xd9, 0xc6, 0x55, 0x50,
	0xd4, 0xf0, 0x76, 0x2d, 0xc4, 0x01, 0x16, 0x28, 0x50, 0x14, 0x50, 0x2c, 0x29, 0x2b, 0xc4, 0x91,
	0x0c, 0x4a, 0x41, 0xea, 0xbd, 0x19, 0x1c, 0x91, 0x23, 0x72, 0x2a, 0x72, 0x46, 0x9d, 0x21, 0xed,
	0x68, 0x5f, 0xa5, 0x0f, 0xd0, 0xab, 0x3e, 0x49, 0x6f, 0xfa, 0x48, 0xc5, 0xfc, 0x48, 0xde, 0x74,
	0xbd, 0xb9, 0x11, 0x78, 0xbe, 0xef, 0x13, 0xcf, 0x9c, 0xdf, 0x21, 0x1c, 0xa9, 0x32, 0x63, 0xba,
	0x6b, 0x7f, 0xcf, 0x56, 0x4a, 0x16, 0x92, 0xec, 0x59, 0xe3, 0xd9, 0xef, 0x12, 0x29, 0x93, 0x8c,
	0x75, 0x2d, 0x38, 0x2f, 0x17, 0xdd, 0xb8, 0x54, 0x58, 0x70, 0x29, 0x9c, 0xec, 0xd9, 0xe1, 0x1c,
	0x35, 0xeb, 0x9a, 0x9f, 0x2f, 0x80, 0x35, 0x2e, 0x4b, 0x07, 0x74, 0xfe, 0xb9, 0x03, 0xbb, 0xd7,
	0xb8, 0x2c, 0xc9, 0x4b, 0x68, 0xca, 0x15, 0x13, 0xb4, 0x40, 0xb1, 0x46, 0x19, 0x54, 0x8e, 0x2b,
	0x27, 0xf5, 0x10, 0x0c, 0x34, 0xb3, 0x08, 0x79, 0x0d, 0x8f, 0x52, 0xe4, 0x05, 0xe3, 0x74, 0xa1,
	0x64, 0x4e, 0x33, 0x7e, 0xc3, 0xa8, 0x14, 0xd9, 0x3a, 0xd8, 0xb1, 0x52, 0xe2, 0xc8, 0xa1, 0x92,
	0xf9, 0x25, 0xbf, 0x61, 0x13, 0x91, 0xad, 0x49, 0x1b, 0xaa, 0xa5, 0xc2, 0xa0, 0x6a, 0x05, 0xe6,
	0x91, 0x3c, 0x86, 0x07, 0x7c, 0x85, 0x85, 0x2e, 0x83, 0x5d, 0x0b, 0x7a, 0x8b, 0x9c, 0x42, 0x1d,
	0x97, 0x48, 0x63, 0xa9, 0x30, 0xd8, 0x3b, 0xae, 0x9c, 0x34, 0xcf, 0x0f, 0xcf, 0xec, 0xb1, 0x47,
	0x42, 0x17, 0x28, 0x22, 0xa6, 0xc3, 0x1a, 0x2e, 0xb1, 0x2f, 0x15, 0x92, 0x57, 0xf0, 0x40, 0x31,
	0x91, 0xca, 0x32, 0x78, 0x70, 0x5c, 0x39, 0x69, 0x9d, 0x37, 0x9d, 0xf2, 0x92, 0xe7, 0xbc, 0x08,
	0x3d, 0x45, 0xbe, 0x01, 0x50, 0x5c, 0xe8, 0x14, 0x05, 0x5d, 0x94, 0x41, 0xcd, 0x3a, 0x6b, 0x78,
	0x64, 0x58, 0x92, 0x3f, 0x40, 0x2b, 0x51, 0x8c, 0x09, 0xaa, 0xd8, 0x3f, 0x4a, 0xae, 0x58, 0x1c,
	0xd4, 0xad, 0xe4, 0xc0, 0xa2, 0xa1, 0x07, 0x3b, 0xff, 0xa9, 0x40, 0x6d, 0x1a, 0x49, 0xc5, 0x45,
	0x42, 0x7e, 0x0f, 0xfb, 0x39, 0x8a, 0x04, 0x05, 0x55, 0xb2, 0x14, 0xb1, 0xcf, 0x50, 0xd3, 0x61,
	0xa1, 0x81, 0xc8, 0x2b, 0x38, 0x58, 0xe2, 0x4f, 0x92, 0x51, 0x93, 0xe0, 0x1c, 0x85, 0x4f, 0xcd,
	0xbe, 0x05, 0xaf, 0x1d, 0x66, 0x12, 0xed, 0x69, 0xaa, 0xcb, 0xdc, 0x27, 0x07, 0x3c, 0x34, 0x2d,
	0x73, 0x23, 0x48, 0xa5, 0x98, 0x23, 0xcd, 0xa5, 0x60, 0x6b, 0x9b, 0xa8, 0x6a, 0x08, 0x16, 0xfa,
	0x60, 0x10, 0xf2, 0x3d, 0x1c, 0xc6, 0xb2, 0x9c, 0x67, 0x5b, 0x3f, 0x3a, 0xd8, 0x3b, 0xae, 0x9e,
	0xb4, 0xce, 0x0f, 0x5c, 0x26, 0xbc, 0xa7, 0xb0, 0xe5, 0x54, 0xde, 0xd4, 0x9d, 0x1b, 0xd8, 0xed,
	0x2b, 0xbc, 0x25, 0x0f, 0x61, 0xef, 0x96, 0x8b, 0x58, 0xfb, 0x10, 0x9c, 0x41, 0x02, 0xa8, 0x2d,
	0xe5, 0xb2, 0xd4, 0x29, 0xf7, 0xc7, 0xde, 0x98, 0x84, 0xc0, 0xee, 0xd2, 0x38, 0x71, 0x47, 0xb5,
	0xcf, 0xa6, 0x90, 0x8a, 0xf3, 0x28, 0xe5, 0x9b, 0x42, 0x3a, 0xcb, 0x68, 0x95, 0x14, 0x6f, 0x6c,
	0x11, 0xeb, 0xa1, 0x7d, 0xee, 0x44, 0x50, 0xbd, 0x42, 0xf9, 0xeb, 0x6e, 0x63, 0x85, 0x89, 0x14,
	0x7a, 0xe3, 0xd6, 0x9b, 0xf7, 0xba, 0x0d, 0xa0, 0xe6, 0x8b, 0xe8, 0xfd, 0x6e, 0xcc, 0x4e, 0x09,
	0x70, 0x21, 0xf3, 0x55, 0xc6, 0x3e, 0x7f, 0xcc, 0x4d, 0x8f, 0x1c, 0xc4, 0x6c, 0x81, 0x65, 0x56,
	0xf8, 0x2c, 0x56, 0x8e, 0xab, 0x27, 0x24, 0xdc, 0xf7, 0xa0, 0xcb, 0xa3, 0xa9, 0x28, 0x17, 0xa5,
	0x7e, 0xed, 0x35, 0x3b, 0x56, 0xd3, 0x74, 0x98, 0x93, 0xbc, 0x84, 0xe6, 0x2a, 0xbb, 0x53, 0x54,
	0xad, 0x02, 0x2c, 0x64, 0x05, 0x9d, 0x7f, 0xd7, 0x61, 0xf7, 0x1d, 0xe6, 0x8c, 0x74, 0x01, 0x30,
	0x41, 0xc5, 0xe9, 0x1a, 0x73, 0x66, 0x43, 0x6c, 0x9d, 0xb7, 0xcf, 0xdc, 0xcc, 0xf6, 0x0c, 0x71,
	0x8d, 0x39, 0x0b, 0x1b, 0xb8, 0x79, 0x24, 0x7f, 0x82, 0x46, 0x94, 0x72, 0xaa, 0x53, 0xbe, 0x28,
	0x6c, 0xe8, 0xad, 0xf3, 0x43, 0xaf, 0x9f, 0x1a, 0x8c, 0x8b, 0x24, 0xac, 0x47, 0x29, 0xb7, 0x06,
	0x39, 0x85, 0x23, 0x26, 0x62, 0x3a, 0x5f, 0xd3, 0x39, 0x8a, 0x48, 0xc9, 0x88, 0x17, 0x6b, 0x9f,
	0x99, 0x43, 0x26, 0xe2, 0xb7, 0xeb, 0xb7, 0x5b, 0xd8, 0x24, 0x49, 0x60, 0x82, 0x7a, 0x5b, 0x9c,
	0x8d, 0x49, 0x5e, 0x40, 0x03, 0x0b, 0xcc, 0x31, 0x45, 0xc1, 0x7c, 0x89, 0xee, 0x00, 0xf2, 0x2d,
	0x10, 0xd7, 0x78, 0x2b, 0x5c, 0xb3, 0x98, 0x16, 0x92, 0x62, 0x96, 0xd9, 0x21, 0xab, 0x87, 0x87,
	0x96, 0xb9, 0x32, 0xc4, 0x4c, 0xf6, 0xb2, 0x8c, 0xfc, 0x19, 0x9e, 0xba, 0x92, 0x53, 0xc5, 0x8a,
	0x52, 0x09, 0x2a, 0x05, 0xcd, 0xcb, 0xac, 0xe0, 0x4a, 0x0a, 0x3f, 0x6f, 0x8f, 0x9d, 0x20, 0xb4,
	0xfc, 0x44, 0x7c, 0xf0, 0x2c, 0x79, 0x0e, 0x0d, 0xbd, 0x62, 0x2c, 0xa6, 0xa6, 0x7d, 0xdc, 0xdc,
	0xd5, 0x2d, 0x70, 0x91, 0x72, 0x72, 0x0c, 0xfb, 0x1a, 0xd7, 0xe6, 0x6d, 0x51, 0xc6, 0xa3, 0x65,
	0xd0, 0x70, 0xf3, 0xa1, 0x71, 0x3d, 0x11, 0x17, 0x06, 0x21, 0x7f, 0x85, 0x17, 0x3a, 0x95, 0x65,
	0x16, 0xd3, 0x14, 0x6f, 0x98, 0xab, 0x0c, 0x5d, 0x48, 0x45, 0x7d, 0x43, 0x82, 0xfd, 0x47, 0xe0,
	0x34, 0x3f, 0xe0, 0x0d, 0xb3, 0x95, 0x1a, 0x4a, 0x15, 0xba, 0x16, 0xfd, 0x16, 0xea, 0x58, 0xc8,
	0xf8, 0xa7, 0x72, 0xc9, 0x82, 0xe6, 0x17, 0x79, 0xef, 0x79, 0x38, 0xdc, 0x0a, 0x4c, 0x03, 0xe8,
	0x02, 0xd5, 0xa6, 0x8d, 0xf6, 0xdd, 0x30, 0x5a, 0xc8, 0x75, 0xc8, 0x09, 0xb4, 0x4d, 0x61, 0x14,
	0x8b, 0xcb, 0xc8, 0x1f, 0x26, 0x38, 0xb0, 0xaa, 0x16, 0x13, 0x71, 0x68, 0x61, 0xa7, 0x7c, 0x0e,
	0x0d, 0xb9, 0xdc, 0x4c, 0x75, 0xcb, 0x4a, 0xea, 0x72, 0xe9, 0x67, 0xba, 0x03, 0x07, 0x39, 0x17,
	0xf4, 0x96, 0x0b, 0x2f, 0x38, 0xb4, 0x02, 0xd3, 0x8c, 0x9f, 0xb8, 0x70, 0x9a, 0x33, 0x38, 0xc8,
	0xf1, 0x33, 0xcd, 0x50, 0x17, 0x46, 0x18, 0x07, 0x6d, 0x7b, 0x7a, 0x70, 0x53, 0xff, 0x89, 0x8b,
	0xd8, 0xac, 0xa3, 0xcf, 0x97, 0xa8, 0x0b, 0x63, 0x90, 0x3f, 0x42, 0xe3, 0x4e, 0x7b, 0xf4, 0x0b,
	0x6d, 0x3d, 0xdb, 0x08, 0x5f, 0x41, 0xb5, 0xcc, 0x31, 0x20, 0x76, 0xf1, 0x1e, 0xf9, 0x64, 0xdc,
	0x4d, 0x53, 0x68, 0x58, 0xf2, 0x12, 0x76, 0x63, 0x85, 0xb7, 0xc1, 0x6f, 0xad, 0xaa, 0xe9, 0x55,
	0x66, 0xa1, 0x84, 0x96, 0x20, 0xdf, 0x9b, 0xed, 0x27, 0xec, 0x0e, 0xa7, 0xe6, 0xde, 0x08, 0x1e,
	0x5a, 0x97, 0xc4, 0x2b, 0xdf, 0xa3, 0x30, 0xeb, 0x7b, 0xb2, 0x62, 0x22, 0x6c, 0x2e, 0xef, 0x0c,
	0x93, 0x97, 0x32, 0x47, 0xaa, 0x53, 0x54, 0x2c, 0x78, 0xe4, 0xda, 0xa1, 0xcc, 0x71, 0x6a, 0x6c,
	0xf2, 0x17, 0x78, 0xee, 0xdb, 0x4c, 0x17, 0x3c, 0x5a, 0x6a, 0x9a, 0x48, 0xa6, 0x4d, 0x6b, 0x2e,
	0xb8, 0xd2, 0x45, 0xf0, 0xd8, 0xca, 0x9f, 0x38, 0xc9, 0xd4, 0x2a, 0xde, 0x49, 0xa6, 0x67, 0x72,
	0x68, 0x68, 0xf2, 0x02, 0xaa, 0x2b, 0x94, 0xc1, 0x13, 0x7b, 0x64, 0xf0, 0x07, 0xb9, 0x42, 0x19,
	0x1a, 0xd8, 0xcc, 0x94, 0x28, 0xf3, 0x39, 0x53, 0x54, 0x2e, 0xe8, 0x2a, 0xc3, 0x35, 0x53, 0x3a,
	0x08, 0x6c, 0xde, 0x0f, 0x1d, 0x31, 0x59, 0x5c, 0x39, 0xd8, 0xf4, 0x81, 0x39, 0x04, 0x66, 0x99,
	0xbc, 0x65, 0x71, 0xf0, 0xd4, 0x75, 0x65, 0x94, 0xf2, 0x9e, 0x43, 0xcc, 0x8d, 0xa2, 0xcb, 0x24,
	0x61, 0xba, 0xa0, 0x42, 0x16, 0x0c, 0x45, 0xf0, 0xcc, 0xdd, 0x28, 0x1e, 0x1d, 0x5b, 0xb0, 0xf3,
	0xaf, 0x0a, 0xd4, 0x42, 0x73, 0x0c, 0x56, 0x90, 0x16, 0xec, 0x70, 0x77, 0x8f, 0x34, 0xc2, 0x1d,
	0x1e, 0x93, 0x63, 0x68, 0xc6, 0x4c, 0x47, 0x8a, 0xaf, 0xcc, 0x15, 0x6e, 0x77, 0x42, 0x23, 0xfc,
	0x39, 0x44, 0x4e, 0xa0, 0xa6, 0xdd, 0x75, 0x64, 0x67, 0xbf, 0x79, 0xde, 0xda, 0x6c, 0x0c, 0x87,
	0x86, 0x1b, 0xda, 0x54, 0xcb, 0x5c, 0x0e, 0x76, 0x01, 0xdc, 0x55, 0xcb, 0x5c, 0x05, 0xa1, 0x25,
	0x8c, 0x20, 0x31, 0x9b, 0x6a, 0xef, 0x0b, 0x81, 0x59, 0x65, 0xa1, 0x25, 0x3a, 0xff, 0xad, 0x40,
	0x7d, 0xc6, 0x73, 0x26, 0xcb, 0x42, 0x93, 0xef, 0x60, 0xd7, 0x34, 0x8e, 0x3d, 0x6c, 0xf3, 0xfc,
	0xe9, 0x99, 0xfb, 0xee, 0x38, 0xdb, 0x7c, 0x77, 0x9c, 0xf5, 0xfd, 0x77, 0x47, 0x68, 0x65, 0xa4,
	0x0b, 0x7b, 0xec, 0x73, 0xa1, 0xd0, 0xc6, 0xf0, 0x55, 0xbd, 0xd3, 0x91, 0x37, 0x50, 0x53, 0x2c,
	0x92, 0x37, 0x4c, 0xf9, 0xc0, 0xbe, 0xf2, 0x97, 0x8d, 0xd2, 0x78, 0x89, 0x59, 0x86, 0x6b, 0x1f,
	0xe4, 0xd7, 0xbc, 0x58, 0xdd, 0xe9, 0x35, 0x34, 0x7f, 0xd6, 0x85, 0xe4, 0x05, 0x04, 0xef, 0x7b,
	0xe3, 0xfe, 0x24, 0xec, 0xd1, 0xc9, 0xd5, 0x60, 0x4c, 0x3f, 0x8e, 0xa7, 0x57, 0x83, 0x8b, 0xc1,
	0x70, 0x34, 0xe8, 0xb7, 0x7f, 0x43, 0x0e, 0xa0, 0xd1, 0x9f, 0x8c, 0x67, 0x96, 0x6a, 0x57, 0x48,
	0x13, 0x6a, 0xa3, 0xf1, 0x74, 0xd6, 0x1b, 0xcf, 0xda, 0x3b, 0xa4, 0x0d, 0xfb, 0xbd, 0xe1, 0x6c,
	0x10, 0xd2, 0xde, 0xc5, 0x6c, 0x34, 0x19, 0xb7, 0xab, 0xa7, 0x37, 0x50, 0xdf, 0x6c, 0x0f, 0x12,
	0xc0, 0xc3, 0xde, 0x6c, 0xd2, 0xff, 0xf1, 0xe3, 0xfb, 0x81, 0x7f, 0xe7, 0xc8, 0xbf, 0xf3, 0x21,
	0xb4, 0xb7, 0x4c, 0xef, 0xf2, 0x72, 0xf2, 0x69, 0xd0, 0x6f, 0x57, 0xc8, 0x73, 0x78, 0xb2, 0x45,
	0x87, 0x93, 0xf0, 0xed, 0xa8, 0xdf, 0x1f, 0x8c, 0xe9, 0x74, 0x32, 0x34, 0xae, 0xbe, 0x81, 0xa7,
	0xf7, 0x91, 0xb3, 0x70, 0x74, 0x31, 0x6b, 0x57, 0x4f, 0x13, 0x68, 0x6c, 0x6f, 0x17, 0xf2, 0x0c,
	0x1e, 0xf7, 0xde, 0xf5, 0xc2, 0x11, 0xbd, 0xee, 0x7d, 0xf8, 0x7f, 0xd7, 0x6d, 0xd8, 0x1f, 0x7d,
	0xb8, 0xba, 0x1c, 0x5d, 0x8c, 0x66, 0x74, 0x30, 0x36, 0x6e, 0x1f, 0xc1, 0xd1, 0xe0, 0x6f, 0x1e,
	0xb9, 0x98, 0x8c, 0x67, 0xa3, 0xf1, 0xc7, 0x41, 0x7b, 0xc7, 0xc0, 0x5b, 0xe1, 0x16, 0xb6, 0x01,
	0x6e, 0xae, 0x25, 0x13, 0xe0, 0xf4, 0x87, 0xd1, 0x70, 0x36, 0x1a, 0xbf, 0xfb, 0x65, 0x80, 0x5b,
	0xe6, 0x8b, 0x00, 0xb7, 0xe8, 0x7d, 0x01, 0xde, 0x47, 0xfa, 0x00, 0xdf, 0xbe, 0xfe, 0xb1, 0x9b,
	0xf0, 0x22, 0x2d, 0xe7, 0x67, 0x91, 0xcc, 0xbb, 0xb1, 0x90, 0x37, 0x7c, 0x29, 0x17, 0x8b, 0x6e,
	0x8e, 0xe9, 0xdf, 0xa5, 0x48, 0xbe, 0xc3, 0x15, 0xef, 0x26, 0x4c, 0xd8, 0x9a, 0xbb, 0x4f, 0xe4,
	0xf9, 0x03, 0x6b, 0xbc, 0xf9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x30, 0x26, 0x2c, 0x38,
	0x0b, 0x00, 0x00,
}