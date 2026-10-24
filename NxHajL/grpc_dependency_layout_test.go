package NxHajL

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"recovered.local/soga/GYzZ5v"
	"reflect"
	"strings"
	"testing"
	"time"
)

// 这些期望值来自原 ELF RTTI，而非候选库的布局快照。依赖版本升级若改变
// 已核对的数据结构必须重新核查原机器码；布局一致本身不代表所有函数一致。
// 证据：work/ida-20261023/grpc-compare/type-map164.json。
type originalGRPCField struct {
	name         string
	offset, size uintptr
	kind         reflect.Kind
}
type originalGRPCType struct {
	size   uintptr
	fields []originalGRPCField
}

var originalGRPCTypes = map[string]originalGRPCType{
	"grpc.parser": {40, []originalGRPCField{
		{"r", 0, 16, reflect.Interface},
		{"header", 16, 5, reflect.Array},
		{"recvBufferPool", 24, 16, reflect.Interface},
	}},
	"grpc.clientStream": {272, []originalGRPCField{
		{"callHdr", 0, 8, reflect.Pointer},
		{"opts", 8, 24, reflect.Slice},
		{"callInfo", 32, 8, reflect.Pointer},
		{"cc", 40, 8, reflect.Pointer},
		{"desc", 48, 8, reflect.Pointer},
		{"codec", 56, 16, reflect.Interface},
		{"cp", 72, 16, reflect.Interface},
		{"comp", 88, 16, reflect.Interface},
		{"cancel", 104, 8, reflect.Func},
		{"sentLast", 112, 1, reflect.Bool},
		{"methodConfig", 120, 8, reflect.Pointer},
		{"ctx", 128, 16, reflect.Interface},
		{"retryThrottler", 144, 8, reflect.Pointer},
		{"binlogs", 152, 24, reflect.Slice},
		{"serverHeaderBinlogged", 176, 1, reflect.Bool},
		{"mu", 180, 8, reflect.Struct},
		{"firstAttempt", 188, 1, reflect.Bool},
		{"numRetries", 192, 8, reflect.Int},
		{"numRetriesSincePushback", 200, 8, reflect.Int},
		{"finished", 208, 1, reflect.Bool},
		{"attempt", 216, 8, reflect.Pointer},
		{"committed", 224, 1, reflect.Bool},
		{"onCommit", 232, 8, reflect.Func},
		{"buffer", 240, 24, reflect.Slice},
		{"bufferSize", 264, 8, reflect.Int},
	}},
	"grpc.csAttempt": {208, []originalGRPCField{
		{"ctx", 0, 16, reflect.Interface},
		{"cs", 16, 8, reflect.Pointer},
		{"t", 24, 16, reflect.Interface},
		{"s", 40, 8, reflect.Pointer},
		{"p", 48, 8, reflect.Pointer},
		{"pickResult", 56, 32, reflect.Struct},
		{"finished", 88, 1, reflect.Bool},
		{"dc", 96, 16, reflect.Interface},
		{"decomp", 112, 16, reflect.Interface},
		{"decompSet", 128, 1, reflect.Bool},
		{"mu", 132, 8, reflect.Struct},
		{"trInfo", 144, 8, reflect.Pointer},
		{"statsHandlers", 152, 24, reflect.Slice},
		{"beginTime", 176, 24, reflect.Struct},
		{"allowTransparentRetry", 200, 1, reflect.Bool},
		{"drop", 201, 1, reflect.Bool},
	}},
	"transport.Stream": {288, []originalGRPCField{
		{"id", 0, 4, reflect.Uint32},
		{"st", 8, 16, reflect.Interface},
		{"ct", 24, 8, reflect.Pointer},
		{"ctx", 32, 16, reflect.Interface},
		{"cancel", 48, 8, reflect.Func},
		{"done", 56, 8, reflect.Chan},
		{"doneFunc", 64, 8, reflect.Func},
		{"ctxDone", 72, 8, reflect.Chan},
		{"method", 80, 16, reflect.String},
		{"recvCompress", 96, 16, reflect.String},
		{"sendCompress", 112, 16, reflect.String},
		{"buf", 128, 8, reflect.Pointer},
		{"trReader", 136, 16, reflect.Interface},
		{"fc", 152, 8, reflect.Pointer},
		{"wq", 160, 8, reflect.Pointer},
		{"clientAdvertisedCompressors", 168, 16, reflect.String},
		{"requestRead", 184, 8, reflect.Func},
		{"headerChan", 192, 8, reflect.Chan},
		{"headerChanClosed", 200, 4, reflect.Uint32},
		{"headerValid", 204, 1, reflect.Bool},
		{"headerWireLength", 208, 8, reflect.Int},
		{"hdrMu", 216, 8, reflect.Struct},
		{"header", 224, 8, reflect.Map},
		{"trailer", 232, 8, reflect.Map},
		{"noHeaders", 240, 1, reflect.Bool},
		{"headerSent", 244, 4, reflect.Uint32},
		{"state", 248, 4, reflect.Uint32},
		{"status", 256, 8, reflect.Pointer},
		{"bytesReceived", 264, 4, reflect.Uint32},
		{"unprocessed", 268, 4, reflect.Uint32},
		{"contentSubtype", 272, 16, reflect.String},
	}},
	"grpc.ClientConn": {912, []originalGRPCField{
		{"ctx", 0, 16, reflect.Interface},
		{"cancel", 16, 8, reflect.Func},
		{"target", 24, 16, reflect.String},
		{"parsedTarget", 40, 144, reflect.Struct},
		{"authority", 184, 16, reflect.String},
		{"dopts", 200, 496, reflect.Struct},
		{"channelz", 696, 8, reflect.Pointer},
		{"resolverBuilder", 704, 16, reflect.Interface},
		{"idlenessMgr", 720, 8, reflect.Pointer},
		{"csMgr", 728, 8, reflect.Pointer},
		{"pickerWrapper", 736, 8, reflect.Pointer},
		{"safeConfigSelector", 744, 40, reflect.Struct},
		{"retryThrottler", 784, 16, reflect.Struct},
		{"mu", 800, 24, reflect.Struct},
		{"resolverWrapper", 824, 8, reflect.Pointer},
		{"balancerWrapper", 832, 8, reflect.Pointer},
		{"sc", 840, 8, reflect.Pointer},
		{"conns", 848, 8, reflect.Map},
		{"mkp", 856, 24, reflect.Struct},
		{"firstResolveEvent", 880, 8, reflect.Pointer},
		{"lceMu", 888, 8, reflect.Struct},
		{"lastConnectionError", 896, 16, reflect.Interface},
	}},
	"grpc.dialOptions": {496, []originalGRPCField{
		{"unaryInt", 0, 8, reflect.Func},
		{"streamInt", 8, 8, reflect.Func},
		{"chainUnaryInts", 16, 24, reflect.Slice},
		{"chainStreamInts", 40, 24, reflect.Slice},
		{"cp", 64, 16, reflect.Interface},
		{"dc", 80, 16, reflect.Interface},
		{"bs", 96, 16, reflect.Interface},
		{"block", 112, 1, reflect.Bool},
		{"returnLastError", 113, 1, reflect.Bool},
		{"timeout", 120, 8, reflect.Int64},
		{"authority", 128, 16, reflect.String},
		{"binaryLogger", 144, 16, reflect.Interface},
		{"copts", 160, 192, reflect.Struct},
		{"callOptions", 352, 24, reflect.Slice},
		{"channelzParent", 376, 16, reflect.Interface},
		{"disableServiceConfig", 392, 1, reflect.Bool},
		{"disableRetry", 393, 1, reflect.Bool},
		{"disableHealthCheck", 394, 1, reflect.Bool},
		{"healthCheckFunc", 400, 8, reflect.Func},
		{"minConnectTimeout", 408, 8, reflect.Func},
		{"defaultServiceConfig", 416, 8, reflect.Pointer},
		{"defaultServiceConfigRawJSON", 424, 8, reflect.Pointer},
		{"resolvers", 432, 24, reflect.Slice},
		{"idleTimeout", 456, 8, reflect.Int64},
		{"recvBufferPool", 464, 16, reflect.Interface},
		{"defaultScheme", 480, 16, reflect.String},
	}},
	"transport.ConnectOptions": {192, []originalGRPCField{
		{"UserAgent", 0, 16, reflect.String},
		{"Dialer", 16, 8, reflect.Func},
		{"FailOnNonTempDialError", 24, 1, reflect.Bool},
		{"PerRPCCredentials", 32, 24, reflect.Slice},
		{"TransportCredentials", 56, 16, reflect.Interface},
		{"CredsBundle", 72, 16, reflect.Interface},
		{"KeepaliveParams", 88, 24, reflect.Struct},
		{"StatsHandlers", 112, 24, reflect.Slice},
		{"InitialWindowSize", 136, 4, reflect.Int32},
		{"InitialConnWindowSize", 140, 4, reflect.Int32},
		{"WriteBufferSize", 144, 8, reflect.Int},
		{"ReadBufferSize", 152, 8, reflect.Int},
		{"SharedWriteBuffer", 160, 1, reflect.Bool},
		{"ChannelzParent", 168, 8, reflect.Pointer},
		{"MaxHeaderListSize", 176, 8, reflect.Pointer},
		{"UseProxy", 184, 1, reflect.Bool},
	}},
	"grpc.Server": {504, []originalGRPCField{
		{"opts", 0, 352, reflect.Struct},
		{"mu", 352, 8, reflect.Struct},
		{"lis", 360, 8, reflect.Map},
		{"conns", 368, 8, reflect.Map},
		{"serve", 376, 1, reflect.Bool},
		{"drain", 377, 1, reflect.Bool},
		{"cv", 384, 8, reflect.Pointer},
		{"services", 392, 8, reflect.Map},
		{"events", 400, 16, reflect.Interface},
		{"quit", 416, 8, reflect.Pointer},
		{"done", 424, 8, reflect.Pointer},
		{"channelzRemoveOnce", 432, 12, reflect.Struct},
		{"serveWG", 448, 16, reflect.Struct},
		{"handlersWG", 464, 16, reflect.Struct},
		{"channelz", 480, 8, reflect.Pointer},
		{"serverWorkerChannel", 488, 8, reflect.Chan},
		{"serverWorkerChannelClose", 496, 8, reflect.Func},
	}},
	"grpc.serverOptions": {352, []originalGRPCField{
		{"creds", 0, 16, reflect.Interface},
		{"codec", 16, 16, reflect.Interface},
		{"cp", 32, 16, reflect.Interface},
		{"dc", 48, 16, reflect.Interface},
		{"unaryInt", 64, 8, reflect.Func},
		{"streamInt", 72, 8, reflect.Func},
		{"chainUnaryInts", 80, 24, reflect.Slice},
		{"chainStreamInts", 104, 24, reflect.Slice},
		{"binaryLogger", 128, 16, reflect.Interface},
		{"inTapHandle", 144, 8, reflect.Func},
		{"statsHandlers", 152, 24, reflect.Slice},
		{"maxConcurrentStreams", 176, 4, reflect.Uint32},
		{"maxReceiveMessageSize", 184, 8, reflect.Int},
		{"maxSendMessageSize", 192, 8, reflect.Int},
		{"unknownStreamDesc", 200, 8, reflect.Pointer},
		{"keepaliveParams", 208, 40, reflect.Struct},
		{"keepalivePolicy", 248, 16, reflect.Struct},
		{"initialWindowSize", 264, 4, reflect.Int32},
		{"initialConnWindowSize", 268, 4, reflect.Int32},
		{"writeBufferSize", 272, 8, reflect.Int},
		{"readBufferSize", 280, 8, reflect.Int},
		{"sharedWriteBuffer", 288, 1, reflect.Bool},
		{"connectionTimeout", 296, 8, reflect.Int64},
		{"maxHeaderListSize", 304, 8, reflect.Pointer},
		{"headerTableSize", 312, 8, reflect.Pointer},
		{"numServerWorkers", 320, 4, reflect.Uint32},
		{"recvBufferPool", 328, 16, reflect.Interface},
		{"waitForHandlers", 344, 1, reflect.Bool},
	}},
	"grpc.serverStream": {232, []originalGRPCField{
		{"ctx", 0, 16, reflect.Interface}, {"t", 16, 16, reflect.Interface}, {"s", 32, 8, reflect.Pointer}, {"p", 40, 8, reflect.Pointer},
		{"codec", 48, 16, reflect.Interface}, {"cp", 64, 16, reflect.Interface}, {"dc", 80, 16, reflect.Interface}, {"comp", 96, 16, reflect.Interface}, {"decomp", 112, 16, reflect.Interface},
		{"sendCompressorName", 128, 16, reflect.String}, {"maxReceiveMessageSize", 144, 8, reflect.Int}, {"maxSendMessageSize", 152, 8, reflect.Int}, {"trInfo", 160, 8, reflect.Pointer},
		{"statsHandler", 168, 24, reflect.Slice}, {"binlogs", 192, 24, reflect.Slice}, {"serverHeaderBinlogged", 216, 1, reflect.Bool}, {"mu", 220, 8, reflect.Struct},
	}},
}

func checkOriginalGRPCType(typ reflect.Type) error {
	want, ok := originalGRPCTypes[typ.String()]
	if !ok {
		return nil
	}
	if typ.Size() != want.size || typ.NumField() != len(want.fields) {
		return fmt.Errorf("%v: size/field count = %d/%d, original %d/%d", typ, typ.Size(), typ.NumField(), want.size, len(want.fields))
	}
	for i, w := range want.fields {
		f := typ.Field(i)
		if f.Name != w.name || f.Offset != w.offset || f.Type.Size() != w.size || f.Type.Kind() != w.kind {
			return fmt.Errorf("%v field %d: got %s/%d/%d/%v, original %+v", typ, i, f.Name, f.Offset, f.Type.Size(), f.Type.Kind(), w)
		}
	}
	return nil
}
func TestOriginalGRPCDependencyRTTILayout(t *testing.T) {
	seen := map[reflect.Type]bool{}
	found := map[string]bool{}
	var visit func(reflect.Type)
	visit = func(typ reflect.Type) {
		if seen[typ] {
			return
		}
		seen[typ] = true
		if typ.PkgPath() != "" && !strings.HasPrefix(typ.PkgPath(), "google.golang.org/grpc") {
			return
		}
		if _, ok := originalGRPCTypes[typ.String()]; ok {
			found[typ.String()] = true
			if err := checkOriginalGRPCType(typ); err != nil {
				t.Error(err)
			}
		}
		switch typ.Kind() {
		case reflect.Struct:
			for i := 0; i < typ.NumField(); i++ {
				visit(typ.Field(i).Type)
			}
		case reflect.Pointer, reflect.Slice, reflect.Array, reflect.Chan:
			visit(typ.Elem())
		case reflect.Map:
			visit(typ.Key())
			visit(typ.Elem())
		case reflect.Func:
			for i := 0; i < typ.NumIn(); i++ {
				visit(typ.In(i))
			}
			for i := 0; i < typ.NumOut(); i++ {
				visit(typ.Out(i))
			}
		}
		for i := 0; i < typ.NumMethod(); i++ {
			visit(typ.Method(i).Type)
		}
	}
	visit(reflect.TypeOf(&grpc.Server{}))
	visit(reflect.TypeOf(&grpc.ClientConn{}))
	for name := range originalGRPCTypes {
		if name != "grpc.serverStream" && !found[name] {
			t.Errorf("original type %s not reached", name)
		}
	}
}

// 通过真实本地 RPC 取得库内部 serverStream 的实际类型，避免另造同布局类型。
func TestOriginalGRPCServerStreamRTTIAndWire(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	server := grpc.NewServer()
	defer server.Stop()
	defer listener.Close()
	checked := make(chan error, 1)
	server.RegisterService(&grpc.ServiceDesc{ServiceName: "evidence.Layout", Streams: []grpc.StreamDesc{{StreamName: "Exchange", ClientStreams: true, ServerStreams: true, Handler: func(_ any, s grpc.ServerStream) error {
		e := checkOriginalGRPCType(reflect.TypeOf(s).Elem())
		checked <- e
		if e != nil {
			return e
		}
		var m GYzZ5v.Z27tYEZ9hy1o
		if e = s.RecvMsg(&m); e != nil {
			return e
		}
		return s.SendMsg(&m)
	}}}}, nil)
	go server.Serve(listener)
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stream, err := conn.NewStream(ctx, &grpc.StreamDesc{ClientStreams: true, ServerStreams: true}, "/evidence.Layout/Exchange")
	if err != nil {
		t.Fatal(err)
	}
	if err = stream.SendMsg(&GYzZ5v.Z27tYEZ9hy1o{Q2bCNbJ12i: []byte("original grpc type evidence")}); err != nil {
		t.Fatal(err)
	}
	var got GYzZ5v.Z27tYEZ9hy1o
	if err = stream.RecvMsg(&got); err != nil {
		t.Fatal(err)
	}
	if string(got.Q2bCNbJ12i) != "original grpc type evidence" {
		t.Fatal("payload changed")
	}
	select {
	case err = <-checked:
		if err != nil {
			t.Fatal(err)
		}
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
}
