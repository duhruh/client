// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/common.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type Time int64
type StringKVPair struct {
	Key   string `codec:"key" json:"key"`
	Value string `codec:"value" json:"value"`
}

type Status struct {
	Code   int            `codec:"code" json:"code"`
	Name   string         `codec:"name" json:"name"`
	Desc   string         `codec:"desc" json:"desc"`
	Fields []StringKVPair `codec:"fields" json:"fields"`
}

type UID string
type DeviceID string
type SigID string
type KID string
type BinaryKID []byte
type TLFID string
type Bytes32 [32]byte
type Text struct {
	Data   string `codec:"data" json:"data"`
	Markup bool   `codec:"markup" json:"markup"`
}

type PGPIdentity struct {
	Username string `codec:"username" json:"username"`
	Comment  string `codec:"comment" json:"comment"`
	Email    string `codec:"email" json:"email"`
}

type PublicKey struct {
	KID               KID           `codec:"KID" json:"KID"`
	PGPFingerprint    string        `codec:"PGPFingerprint" json:"PGPFingerprint"`
	PGPIdentities     []PGPIdentity `codec:"PGPIdentities" json:"PGPIdentities"`
	IsSibkey          bool          `codec:"isSibkey" json:"isSibkey"`
	IsEldest          bool          `codec:"isEldest" json:"isEldest"`
	ParentID          string        `codec:"parentID" json:"parentID"`
	DeviceID          DeviceID      `codec:"deviceID" json:"deviceID"`
	DeviceDescription string        `codec:"deviceDescription" json:"deviceDescription"`
	DeviceType        string        `codec:"deviceType" json:"deviceType"`
	CTime             Time          `codec:"cTime" json:"cTime"`
	ETime             Time          `codec:"eTime" json:"eTime"`
}

type KeybaseTime struct {
	Unix  Time `codec:"unix" json:"unix"`
	Chain int  `codec:"chain" json:"chain"`
}

type RevokedKey struct {
	Key  PublicKey   `codec:"key" json:"key"`
	Time KeybaseTime `codec:"time" json:"time"`
	By   KID         `codec:"by" json:"by"`
}

type User struct {
	Uid      UID    `codec:"uid" json:"uid"`
	Username string `codec:"username" json:"username"`
}

type Device struct {
	Type         string   `codec:"type" json:"type"`
	Name         string   `codec:"name" json:"name"`
	DeviceID     DeviceID `codec:"deviceID" json:"deviceID"`
	CTime        Time     `codec:"cTime" json:"cTime"`
	MTime        Time     `codec:"mTime" json:"mTime"`
	LastUsedTime Time     `codec:"lastUsedTime" json:"lastUsedTime"`
	EncryptKey   KID      `codec:"encryptKey" json:"encryptKey"`
	VerifyKey    KID      `codec:"verifyKey" json:"verifyKey"`
	Status       int      `codec:"status" json:"status"`
}

type DeviceType int

const (
	DeviceType_DESKTOP DeviceType = 0
	DeviceType_MOBILE  DeviceType = 1
)

var DeviceTypeMap = map[string]DeviceType{
	"DESKTOP": 0,
	"MOBILE":  1,
}

var DeviceTypeRevMap = map[DeviceType]string{
	0: "DESKTOP",
	1: "MOBILE",
}

type Stream struct {
	Fd int `codec:"fd" json:"fd"`
}

type LogLevel int

const (
	LogLevel_NONE     LogLevel = 0
	LogLevel_DEBUG    LogLevel = 1
	LogLevel_INFO     LogLevel = 2
	LogLevel_NOTICE   LogLevel = 3
	LogLevel_WARN     LogLevel = 4
	LogLevel_ERROR    LogLevel = 5
	LogLevel_CRITICAL LogLevel = 6
	LogLevel_FATAL    LogLevel = 7
)

var LogLevelMap = map[string]LogLevel{
	"NONE":     0,
	"DEBUG":    1,
	"INFO":     2,
	"NOTICE":   3,
	"WARN":     4,
	"ERROR":    5,
	"CRITICAL": 6,
	"FATAL":    7,
}

var LogLevelRevMap = map[LogLevel]string{
	0: "NONE",
	1: "DEBUG",
	2: "INFO",
	3: "NOTICE",
	4: "WARN",
	5: "ERROR",
	6: "CRITICAL",
	7: "FATAL",
}

type ClientType int

const (
	ClientType_NONE ClientType = 0
	ClientType_CLI  ClientType = 1
	ClientType_GUI  ClientType = 2
	ClientType_KBFS ClientType = 3
)

var ClientTypeMap = map[string]ClientType{
	"NONE": 0,
	"CLI":  1,
	"GUI":  2,
	"KBFS": 3,
}

var ClientTypeRevMap = map[ClientType]string{
	0: "NONE",
	1: "CLI",
	2: "GUI",
	3: "KBFS",
}

type UserVersionVector struct {
	Id               int64 `codec:"id" json:"id"`
	SigHints         int   `codec:"sigHints" json:"sigHints"`
	SigChain         int64 `codec:"sigChain" json:"sigChain"`
	CachedAt         Time  `codec:"cachedAt" json:"cachedAt"`
	LastIdentifiedAt Time  `codec:"lastIdentifiedAt" json:"lastIdentifiedAt"`
}

type UserPlusKeys struct {
	Uid               UID               `codec:"uid" json:"uid"`
	Username          string            `codec:"username" json:"username"`
	DeviceKeys        []PublicKey       `codec:"deviceKeys" json:"deviceKeys"`
	RevokedDeviceKeys []RevokedKey      `codec:"revokedDeviceKeys" json:"revokedDeviceKeys"`
	PGPKeyCount       int               `codec:"pgpKeyCount" json:"pgpKeyCount"`
	Uvv               UserVersionVector `codec:"uvv" json:"uvv"`
}

type UserPlusAllKeys struct {
	Base    UserPlusKeys `codec:"base" json:"base"`
	PGPKeys []PublicKey  `codec:"pgpKeys" json:"pgpKeys"`
}

type MerkleTreeID int

const (
	MerkleTreeID_MASTER       MerkleTreeID = 0
	MerkleTreeID_KBFS_PUBLIC  MerkleTreeID = 1
	MerkleTreeID_KBFS_PRIVATE MerkleTreeID = 2
)

var MerkleTreeIDMap = map[string]MerkleTreeID{
	"MASTER":       0,
	"KBFS_PUBLIC":  1,
	"KBFS_PRIVATE": 2,
}

var MerkleTreeIDRevMap = map[MerkleTreeID]string{
	0: "MASTER",
	1: "KBFS_PUBLIC",
	2: "KBFS_PRIVATE",
}

// SocialAssertionService is a service that can be used to assert proofs for a
// user.
type SocialAssertionService string

// SocialAssertion contains a service and username for that service, that
// together form an assertion about a user. Resolving an assertion requires
// that the user posts a Keybase proof on the asserted service as the asserted
// user.
type SocialAssertion struct {
	User    string                 `codec:"user" json:"user"`
	Service SocialAssertionService `codec:"service" json:"service"`
}

// UserResolution maps how an unresolved user assertion has been resolved.
type UserResolution struct {
	Assertion SocialAssertion `codec:"assertion" json:"assertion"`
	UserID    UID             `codec:"userID" json:"userID"`
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.Common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
