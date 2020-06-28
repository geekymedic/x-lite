package bff

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc/metadata"

	"git.gmtshenzhen.com/geeky-medic/x-lite/framework/bff/codes"
	"git.gmtshenzhen.com/geeky-medic/x-lite/framework/types"
	"git.gmtshenzhen.com/geeky-medic/x-lite/logger"
	"git.gmtshenzhen.com/geeky-medic/x-lite/logger/extend"
	errors "git.gmtshenzhen.com/geeky-medic/x-lite/xerrors"
)

var (
	empty = struct{}{}
)

func newSessionCtx(ctx context.Context, session *types.Session) context.Context {
	return metadata.NewOutgoingContext(
		ctx,
		metadata.New(session.KeysValues()),
	)
}

func NewState(ctx *gin.Context) *State {
	x, exists := ctx.Get(types.StateName)
	if exists {
		return x.(*State)
	}

	var (
		context = context.Background()
		session = types.NewSessionFromGinCtx(ctx)
		state   = &State{
			Gin:     ctx,
			Session: session,
			Logger:  extend.NewSessionLog(session),
			ctx:     newSessionCtx(context, session),
		}
	)

	ctx.Set(types.StateName, state)

	return state
}

type State struct {
	*types.Session
	logger.Logger
	Gin *gin.Context
	ctx context.Context
}

func (m *State) Error(code int, err error) {
	m.Gin.Set(types.ResponseStatusCode, code)
	if err != nil {
		m.Gin.Set(types.ResponseErr, err)
	}
	m.httpJson(code, empty)
}

func (m *State) Json(code int, v interface{}, err ...error) {
	m.Gin.Set(types.ResponseStatusCode, code)
	if len(err) != 0 {
		m.Gin.Set(types.ResponseErr, err[0])
	}
	m.httpJson(code, v)
}

func (m *State) ErrorMessage(code int, txt string) {
	m.Gin.Set(types.ResponseStatusCode, code)
	if txt != "" {
		m.Gin.Set(types.ResponseErr, fmt.Errorf("%s", txt))
	}
	m.httpJsonMessage(code, txt, empty)
}

func (m *State) Success(v interface{}) {
	m.Gin.Set(types.ResponseStatusCode, codes.CodeSuccess)
	m.httpJson(codes.CodeSuccess, v)
}

func (m *State) Context() context.Context {
	return m.ctx
}

func (m *State) GrpcClientCtx() context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(m.Session.KeysValues()))
}

func (m *State) ShouldBindJSON(v interface{}) error {
	if m.Gin.Writer.Written() {
		return nil
	}
	if err := m.Gin.ShouldBindBodyWith(v, binding.JSON); err != nil {
		m.Session.StructError = err.Error()
		m.Gin.Set(types.XLiteSession, m.Session)
		return errors.By(err)
	}
	return nil
}

func (m *State) httpJson(code int, v interface{}) {
	if m.Gin.Writer.Written() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(map[string]interface{}{
		"Code":    code,
		"Message": codes.GetMessage(code),
		"Data":    v,
	})
	if err != nil {
		logger.With("err", err).Error("fail to encode response")
	}
	m.Gin.Writer.WriteHeader(http.StatusOK)
	_, err = m.Gin.Writer.Write(buf.Bytes())
	if err != nil {
		m.Gin.Set(types.ResponseErr, err)
	}
	m.Gin.Set(types.ResponseBody, buf.String())
}

func (m *State) httpJsonMessage(code int, message string, v interface{}) {
	if m.Gin.Writer.Written() {
		return
	}
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(map[string]interface{}{
		"Code":    code,
		"Message": message,
		"Data":    v,
	})
	if err != nil {
		logger.With("err", err).Error("fail to encode response")
	}
	m.Gin.Writer.WriteHeader(http.StatusOK)
	_, err = m.Gin.Writer.Write(buf.Bytes())
	if err != nil {
		m.Gin.Set(types.ResponseErr, err)
	}
	m.Gin.Set(types.ResponseBody, buf.String())
}
