package main

import (
	gw "awesomeProject/userservice/pkg/grpc/pb"
	"awesomeProject/userservice/pkg/ui/swagger"
	jwt "awesomeProject/userservice/pkg/utils"
	"context"
	"flag"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"path"
	"strings"
)

var (
	endpoint = flag.String("endpoint", "localhost:8082", "gRPC listen address")
)

func main() {
	ctx := context.Background()
	Run(ctx)

}
func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", swaggerServer("pkg/grpc/pb"))
	mux.HandleFunc("/user/signFrom", signFrom())
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	//serveSwaggerUI(mux)
	gwm, _ := newgateway(ctx)
	mux.Handle("/", gwm)

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	//
	//mux :=runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	//err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	//if err != nil {
	//	log.Printf("启动gateway服务错误: %s", err)
	//	return err
	//}
	s := &http.Server{
		Addr:    ":9091",
		Handler: allowCORS(mux),
	}
	fmt.Print("gateway启动 9091")

	return s.ListenAndServe()
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		if r.URL.Path == "/user/login" || r.URL.Path == "/user/signFrom" || r.URL.Path == "/user/sign" || r.URL.Path == "/swagger/userservice.swagger.json" {
		} else {
			token := r.Header.Get("Authorization")
			if token == "" || !jwt.AuthCheck(token) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("token错误或非法"))
				w.Header().Set("Content-Type", "application/json")
				return
			}
		}
		h.ServeHTTP(w, r)

	})
}
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}
func swaggerServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}

func signFrom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body><h2>请输入您的学号</h2><form action='/user/sign' method='post'>学号：<input type='text' name='seq' /><br /><input type='submit' /></form></body></html>"))
	}
}
func newgateway(ctx context.Context) (http.Handler, error) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		log.Printf("启动gateway服务错误: %s", err)
		return nil, err
	}
	return mux, nil
}
