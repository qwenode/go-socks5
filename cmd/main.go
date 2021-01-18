package main

import (
	"github.com/qwenode/gogo/nic"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"net"
	"socks5"
	"time"
)

var address = nic.GetAllIPV4BindAddress()
var addressLen = len(address)

func main() {
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					"vvpppsssk": "lOI1UerioJ2Ieuralaai"}}},
		//Credentials: socks5.StaticCredentials{"vvpppsssk":"lOI#UerioJ#Ieuralaai"},
		//Resolver:    nil,
		//Rules:       nil,
		//Rewriter:    nil,
		//BindIP:      nil,
		//Logger:      nil,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			ip := address[rand.Intn(addressLen)]
			localAddr, err := net.ResolveIPAddr("ip", ip.String())
			log.Println(localAddr, err)
			if err != nil {
				log.Println("xxx", err)
				panic(err)
			}
			localTCPAddr := &net.TCPAddr{
				IP: localAddr.IP,
			}
			localAddrDialier := &net.Dialer{LocalAddr: localTCPAddr,
				Timeout: time.Second * 15,
			}
			return localAddrDialier.Dial(network, addr)
		},
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:65501"); err != nil {
		panic(err)
	}

}
