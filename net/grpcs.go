package net

import (
	"math/rand"
	"time"

	"github.com/hilqiqi0/golang-utils/tools/errs"
	"google.golang.org/grpc"
)

func NewGRPCConnect(address string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	errs.CheckCommonErr(err)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewGRPCConnects(addresses []string) (conns []*grpc.ClientConn, err error) {
	for _, address := range addresses {
		conn, err := NewGRPCConnect(address)
		errs.CheckCommonErr(err)
		if err != nil {
			return nil, err
		}
		conns = append(conns, conn)
	}
	return
}

func GetGRPCConnectByRand(conns []*grpc.ClientConn) *grpc.ClientConn {
	rand.Seed(time.Now().Unix())
	return conns[rand.Intn(len(conns))]
}
