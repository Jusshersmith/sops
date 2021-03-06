// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package credentials

import (
	credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockIamCredentialsServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	credentialspb.IAMCredentialsServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockIamCredentialsServer) GenerateAccessToken(ctx context.Context, req *credentialspb.GenerateAccessTokenRequest) (*credentialspb.GenerateAccessTokenResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*credentialspb.GenerateAccessTokenResponse), nil
}

func (s *mockIamCredentialsServer) GenerateIdToken(ctx context.Context, req *credentialspb.GenerateIdTokenRequest) (*credentialspb.GenerateIdTokenResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*credentialspb.GenerateIdTokenResponse), nil
}

func (s *mockIamCredentialsServer) SignBlob(ctx context.Context, req *credentialspb.SignBlobRequest) (*credentialspb.SignBlobResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*credentialspb.SignBlobResponse), nil
}

func (s *mockIamCredentialsServer) SignJwt(ctx context.Context, req *credentialspb.SignJwtRequest) (*credentialspb.SignJwtResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*credentialspb.SignJwtResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockIamCredentials mockIamCredentialsServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	credentialspb.RegisterIAMCredentialsServer(serv, &mockIamCredentials)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestIamCredentialsGenerateAccessToken(t *testing.T) {
	var accessToken string = "accessToken-1938933922"
	var expectedResponse = &credentialspb.GenerateAccessTokenResponse{
		AccessToken: accessToken,
	}

	mockIamCredentials.err = nil
	mockIamCredentials.reqs = nil

	mockIamCredentials.resps = append(mockIamCredentials.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var scope []string = nil
	var request = &credentialspb.GenerateAccessTokenRequest{
		Name:  formattedName,
		Scope: scope,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GenerateAccessToken(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockIamCredentials.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestIamCredentialsGenerateAccessTokenError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockIamCredentials.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var scope []string = nil
	var request = &credentialspb.GenerateAccessTokenRequest{
		Name:  formattedName,
		Scope: scope,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GenerateAccessToken(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestIamCredentialsGenerateIdToken(t *testing.T) {
	var token string = "token110541305"
	var expectedResponse = &credentialspb.GenerateIdTokenResponse{
		Token: token,
	}

	mockIamCredentials.err = nil
	mockIamCredentials.reqs = nil

	mockIamCredentials.resps = append(mockIamCredentials.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var audience string = "audience975628804"
	var request = &credentialspb.GenerateIdTokenRequest{
		Name:     formattedName,
		Audience: audience,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GenerateIdToken(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockIamCredentials.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestIamCredentialsGenerateIdTokenError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockIamCredentials.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var audience string = "audience975628804"
	var request = &credentialspb.GenerateIdTokenRequest{
		Name:     formattedName,
		Audience: audience,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GenerateIdToken(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestIamCredentialsSignBlob(t *testing.T) {
	var keyId string = "keyId-1134673157"
	var signedBlob []byte = []byte("-32")
	var expectedResponse = &credentialspb.SignBlobResponse{
		KeyId:      keyId,
		SignedBlob: signedBlob,
	}

	mockIamCredentials.err = nil
	mockIamCredentials.reqs = nil

	mockIamCredentials.resps = append(mockIamCredentials.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var payload []byte = []byte("-114")
	var request = &credentialspb.SignBlobRequest{
		Name:    formattedName,
		Payload: payload,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SignBlob(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockIamCredentials.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestIamCredentialsSignBlobError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockIamCredentials.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var payload []byte = []byte("-114")
	var request = &credentialspb.SignBlobRequest{
		Name:    formattedName,
		Payload: payload,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SignBlob(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestIamCredentialsSignJwt(t *testing.T) {
	var keyId string = "keyId-1134673157"
	var signedJwt string = "signedJwt-979546844"
	var expectedResponse = &credentialspb.SignJwtResponse{
		KeyId:     keyId,
		SignedJwt: signedJwt,
	}

	mockIamCredentials.err = nil
	mockIamCredentials.reqs = nil

	mockIamCredentials.resps = append(mockIamCredentials.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var payload string = "-114"
	var request = &credentialspb.SignJwtRequest{
		Name:    formattedName,
		Payload: payload,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SignJwt(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockIamCredentials.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestIamCredentialsSignJwtError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockIamCredentials.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/serviceAccounts/%s", "[PROJECT]", "[SERVICE_ACCOUNT]")
	var payload string = "-114"
	var request = &credentialspb.SignJwtRequest{
		Name:    formattedName,
		Payload: payload,
	}

	c, err := NewIamCredentialsClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SignJwt(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
