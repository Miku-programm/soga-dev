package quicstream

import (
	"context"
	quic "github.com/apernet/quic-go"
	"time"
)

// 原 RTTI 16 字节单接口结构；Close 先取消读再关闭写侧。
type U6qURXBS2qu struct{ Fjf4om quic.Stream }

func (s *U6qURXBS2qu) StreamID() quic.StreamID               { return s.Fjf4om.StreamID() }                      //0x1c2bf40
func (s *U6qURXBS2qu) Read(b []byte) (int, error)            { return s.Fjf4om.Read(b) }                         //0x1c2bf80
func (s *U6qURXBS2qu) CancelRead(code quic.StreamErrorCode)  { s.Fjf4om.CancelRead(code) }                       //0x1c2bfe0
func (s *U6qURXBS2qu) SetReadDeadline(t time.Time) error     { return s.Fjf4om.SetReadDeadline(t) }              //0x1c2c020
func (s *U6qURXBS2qu) Write(b []byte) (int, error)           { return s.Fjf4om.Write(b) }                        //0x1c2c080
func (s *U6qURXBS2qu) Close() error                          { s.Fjf4om.CancelRead(0); return s.Fjf4om.Close() } //0x1c2c0e0
func (s *U6qURXBS2qu) CancelWrite(code quic.StreamErrorCode) { s.Fjf4om.CancelWrite(code) }                      //0x1c2c140
func (s *U6qURXBS2qu) Context() context.Context              { return s.Fjf4om.Context() }                       //0x1c2c180
func (s *U6qURXBS2qu) SetWriteDeadline(t time.Time) error    { return s.Fjf4om.SetWriteDeadline(t) }             //0x1c2c1c0
func (s *U6qURXBS2qu) SetDeadline(t time.Time) error         { return s.Fjf4om.SetDeadline(t) }                  //0x1c2c220
