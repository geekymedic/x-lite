package neon

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"

	"github.com/geekymedic/neon/bff/types"
)

/*
_uid	用户ID	N	标示当前用户
_token	用户Token	N	当前用户令牌
_platform	接入平台	Y	IOS,Android,Web,WechatApp,Windows，Linux，...
_version	客户端版本	Y	1.0.0
_net	当前网络环境	N	2G,3G,4G,5G,wifi,unknown
_mobile	用户手机号	N	15644441111
_os	操作系统	N	系统获取的操作系统名称
_device	设备ID	N	客户端计算唯一的设备ID
_describe	设备描述	N	OPPO R33
_trace	调用链跟踪ID	Y	客户端生成唯一UUID
_sequence	调用序列	Y	客户端每次调用时加一
_time	调用时间戳	Y	调用时候客户端当前时间戳
_stack 服务各系统间的调用栈
_chain 服务调用链条
_storeId 店铺Id N 店铺的id
_path api URL
_struct_err 参数检查结果
*/

type Session struct {
	Uid         string
	Token       string
	Platform    string
	Version     string
	Net         string
	Mobile      string
	OS          string
	Device      string
	Describe    string
	Trace       string
	Sequence    string
	Time        string
	StoreId     string
	Path        string
	ClientIp    string
	StructError string
}

func (m *Session) SetRandomTrace() {
	m.Trace = fmt.Sprintf("_%s", uuid.Must(uuid.NewUUID()).String())
}

func (m *Session) Keys() map[string]*string {
	return map[string]*string{
		"_uid":      &m.Uid,
		"_token":    &m.Token,
		"_platform": &m.Platform,
		"_version":  &m.Version,
		"_net":      &m.Net,
		"_mobile":   &m.Mobile,
		"_os":       &m.OS,
		"_device":   &m.Device,
		"_describe": &m.Describe,
		"_trace":    &m.Trace,
		"_sequence": &m.Sequence,
		"_time":     &m.Time,
		"_storeId":  &m.StoreId,
		"_clientIp": &m.ClientIp,
	}
}

func (m *Session) KeysValues() map[string]string {

	return map[string]string{
		"_uid":      m.Uid,
		"_token":    m.Token,
		"_platform": m.Platform,
		"_version":  m.Version,
		"_net":      m.Net,
		"_mobile":   m.Mobile,
		"_os":       m.OS,
		"_device":   m.Device,
		"_describe": m.Describe,
		"_trace":    m.Trace,
		"_sequence": m.Sequence,
		"_time":     m.Time,
		"_storeId":  m.StoreId,
		"_clientIp": m.ClientIp,
	}
}

func (m *Session) ShortLog() []interface{} {
	return []interface{}{
		"_uid", m.Uid,
		"_token", m.Token,
		"_platform", m.Platform,
		"_version", m.Version,
		"_net", m.Net,
		"_mobile", m.Mobile,
		"_os", m.OS,
		"_device", m.Device,
		"_describe", m.Describe,
		"_trace", m.Trace,
		"_sequence", m.Sequence,
		"_time", m.Time,
		"_storeId", m.StoreId,
		"_path", m.Path,
		"_clientIp", m.ClientIp,
		"_struct_error", m.StructError,
	}
}

func CreateSessionFromGrpcIncomingContext(ctx context.Context) *Session {
	session := &Session{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return session
	}
	session = createSessionFromGrpcMetadata(md)
	return session
}

func CreateSessionFromGrpcOutgoingContext(ctx context.Context) *Session {
	session := &Session{}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return session
	}
	session = createSessionFromGrpcMetadata(md)
	return session
}

func createSessionFromGrpcMetadata(md metadata.MD) *Session {
	session := &Session{}
	fn := func(key string) string {
		if value := md.Get(key); len(value) > 0 {
			return value[0]
		} else {
			return ""
		}
	}

	session.Uid = fn("_uid")
	session.Trace = fn("_trace")
	session.Token = fn("_token")
	session.Describe = fn("_describe")
	session.Device = fn("_device")
	session.Mobile = fn("_mobile")
	session.Net = fn("_net")
	session.OS = fn("_os")
	session.Platform = fn("_platform")
	session.Sequence = fn("_sequence")
	session.Time = fn("_time")
	session.Version = fn("_version")
	session.StoreId = fn("_storeId")
	session.ClientIp = fn("_clientIp")
	return session
}

func NewSessionFromGinCtx(ctx *gin.Context) *Session {
	var (
		s = &Session{}
	)
	session, ok := ctx.Get(types.NeonSession)
	if ok {
		return session.(*Session)
	}

	for name, ref := range s.Keys() {
		*ref = ctx.Query(name)
	}
	if s.Trace == "" {
		s.Trace = uuid.Must(uuid.NewRandom()).String()
	}
	if s.ClientIp == "" {
		s.ClientIp = ctx.ClientIP()
	}
	if ctx.Request.URL.RawQuery != "" {
		s.Path = ctx.Request.URL.Path + "?" + ctx.Request.URL.RawQuery
	} else {
		s.Path = ctx.Request.URL.Path
	}
	return s
}

func SetSession(session *Session, ctx *gin.Context) {
	ctx.Set(types.NeonSession, session)
}

func getSystemName() string {
	return viper.GetString("Name")
}
