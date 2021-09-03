package polit

import (
	"fmt"
	"net"
	"sync"

	"github.com/liuwqiang/tykd/actor"
	"github.com/liuwqiang/tykd/config"
	"github.com/liuwqiang/tykd/db"
	pb "github.com/liuwqiang/tykd/protos/golang"

	"google.golang.org/grpc"
)

const (
	maxRequestSize = 1 * 1024 * 1024 // 1Mb
)

type T struct {
	pb.UnimplementedPolitServer
	actDesc  *actor.Descriptor
	listener net.Listener
	grpcSrv  *grpc.Server
	wg       sync.WaitGroup
	errorCh  chan error
	db       db.Repo
}

func New() (*T, error) {
	listener, err := net.Listen("tcp", config.Global().Server.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener %s", err)
	}
	grpcSrv := grpc.NewServer(grpc.MaxRecvMsgSize(maxRequestSize))
	s := T{
		actDesc:  actor.Root().NewChild(fmt.Sprintf("grpc://%s", config.Global().Server.Addr)),
		listener: listener,
		grpcSrv:  grpcSrv,
		errorCh:  make(chan error, 1),
	}
	pb.RegisterPolitServer(grpcSrv, &s)
	dbRepo, err := db.New()
	if err != nil {
		return nil, fmt.Errorf("new db err %s", err)
	}
	s.db = dbRepo
	return &s, nil
}

func (s *T) Start() {
	actor.Spawn(s.actDesc, &s.wg, func() {
		if err := s.grpcSrv.Serve(s.listener); err != nil {
			s.errorCh <- fmt.Errorf("gRPC API listener failed %s", err)
		}
	})
}

// ErrorCh returns an output channel that HTTP server running in another
// goroutine will use if it stops with error if one occurs. The channel will be
// closed when the server is fully stopped due to an error or otherwise..
func (s *T) ErrorCh() <-chan error {
	return s.errorCh
}

// Stop immediately stops gRPC server. So it is caller's responsibility to make
// sure that all pending requests are completed.
func (s *T) Stop() {
	s.grpcSrv.Stop()
	s.wg.Wait()
	close(s.errorCh)
}
