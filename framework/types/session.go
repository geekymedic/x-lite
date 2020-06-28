package types

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

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
	session, ok := ctx.Get(XLiteSession)
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
	ctx.Set(XLiteSession, session)
}

func getSystemName() string {
	return viper.GetString("Name")
}
