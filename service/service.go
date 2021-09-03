package service

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/liuwqiang/tykd/actor"
	"github.com/liuwqiang/tykd/polit"
	"github.com/liuwqiang/tykd/server"
)

type T struct {
	actDesc *actor.Descriptor
	servers []server.T
	stopCh  chan struct{}
	wg      sync.WaitGroup
}

func Spawn() (*T, error) {
	s := &T{
		actDesc: actor.Root().NewChild("service"),
		stopCh:  make(chan struct{}),
	}

	grpcSrv, err := polit.New()
	if err != nil {
		return nil, fmt.Errorf("failed to start gRPC server %s", err)
	}
	s.servers = append(s.servers, grpcSrv)

	if len(s.servers) == 0 {
		return nil, fmt.Errorf("at least one API server should be configured")
	}

	actor.Spawn(s.actDesc, &s.wg, s.run)
	return s, nil
}

func (s *T) Stop() {
	close(s.stopCh)
	s.wg.Wait()
}

func (s *T) run() {
	selectCases := make([]reflect.SelectCase, len(s.servers)+1)
	for i, srv := range s.servers {
		srv.Start()
		selectCases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(srv.ErrorCh()),
		}
	}
	selectCases[len(s.servers)] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(s.stopCh),
	}

	// Block until either an server error is reported or a Stop is called.
	chosen, val, ok := reflect.Select(selectCases)
	if chosen < len(s.servers) && ok {
		serverErr := val.Interface().(error)
		s.actDesc.Log().WithError(serverErr).Error("API server crashed")
	}

	s.actDesc.Log().Info("Shutting down")

	// Stop all API servers.
	var wg sync.WaitGroup
	for _, srv := range s.servers {
		actor.Spawn(s.actDesc.NewChild("srv_stop"), &wg, srv.Stop)
	}
	wg.Wait()
	s.actDesc.Log().Info("All API servers shutdown")
}
