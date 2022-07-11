package main

import(
  "context"
  "fmt"
  "net"
//  "strings"

  core "github.com/cilium/proxy/go/envoy/config/core/v3"
  auth "github.com/cilium/proxy/go/envoy/service/auth/v3"
  envoy_type "github.com/cilium/proxy/go/envoy/type/v3"

  "github.com/gogo/googleapis/google/rpc"
  status "google.golang.org/genproto/googleapis/rpc/status"

  "google.golang.org/grpc"

  "github.com/corazawaf/coraza/v2"
  "github.com/corazawaf/coraza/v2/seclang"
)

var waf *coraza.Waf

func checkRequest(ctx context.Context, req *auth.CheckRequest) bool {

    tx := waf.NewTransaction()
    defer func(){
        tx.ProcessLogging()
        tx.Clean()
    }()

    source := req.GetAttributes().GetSource().GetAddress().GetSocketAddress()
    destination := req.GetAttributes().GetDestination().GetAddress().GetSocketAddress()

    fmt.Printf("source=%v\n",source)
    fmt.Printf("destination=%v\n",destination)

    http := req.GetAttributes().GetRequest().GetHttp()
    host := http.GetHost()
//    scheme := http.GetScheme()
    protocol := http.GetProtocol()
    method := http.GetMethod()
    path := http.GetPath()
    headers := http.GetHeaders()

    tx.ProcessConnection(source.GetAddress(), int(source.GetPortValue()), destination.GetAddress(), int(destination.GetPortValue()))
    tx.ProcessURI(path, method, protocol)
    tx.AddRequestHeader("Host", host)
    
    for hn,hv := range headers {
        tx.AddRequestHeader(hn, hv)
    }

    if it := tx.ProcessRequestHeaders(); it != nil {
        fmt.Printf("Transaction was interrupted with status %d\n", it.Status)
        return false
    }
    
    fmt.Printf("Transaction was completed\n")
    return true
}


type authorizationServer struct{}

func (a *authorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {

  fmt.Printf("authorizationServer req=%v",req)

  ok := checkRequest(ctx, req)

  if ok {

      return &auth.CheckResponse{
        Status: &status.Status{
          Code: int32(rpc.OK),
        },
        HttpResponse: &auth.CheckResponse_OkResponse{
          OkResponse: &auth.OkHttpResponse{
            Headers: []*core.HeaderValueOption{
              {
                Header: &core.HeaderValue{
                  Key:   "my-credential-header",
                  Value: "permission6,permission9",
                },
              },
            },
          },
        },
      }, nil

  } else {

      return &auth.CheckResponse{
        Status: &status.Status{
          Code: int32(rpc.UNAUTHENTICATED),
        },
        HttpResponse: &auth.CheckResponse_DeniedResponse{
          DeniedResponse: &auth.DeniedHttpResponse{
            Status: &envoy_type.HttpStatus{Code: 401},
            Body:   "Not authorized",
          },
        },
      }, nil
  }

}

func main() {
  fmt.Println("Starting auth goraza")

  fmt.Println("Init goraza")
  waf = coraza.NewWaf()
  parser, _ := seclang.NewParser(waf)

  files := []string{ "/goraza/goraza.conf" }
  for _, f := range files {
    fmt.Printf("Parsing %v\n", f)
    if err := parser.FromFile(f); err != nil {
      panic(err)
    }
  }

  fmt.Println("Start listening 4041")
  lis, _ := net.Listen("tcp", ":4041")
  grpcServer := grpc.NewServer()
  authServer := &authorizationServer{}
  auth.RegisterAuthorizationServer(grpcServer, authServer)
  grpcServer.Serve(lis)
}
