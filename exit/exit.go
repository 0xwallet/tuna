package main

import (
	"encoding/hex"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	. "github.com/nknorg/nkn-sdk-go"
	"github.com/nknorg/nkn/vault"
	"github.com/nknorg/tuna"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/rdegges/go-ipify"
	"github.com/trueinsider/smux"
)

type Configuration struct {
	ListenTCP            int               `json:"ListenTCP"`
	ListenUDP            int               `json:"ListenUDP"`
	Reverse              bool              `json:"Reverse"`
	DialTimeout          uint16            `json:"DialTimeout"`
	UDPTimeout           uint16            `json:"UDPTimeout"`
	PrivateKey           string            `json:"PrivateKey"`
	SubscriptionPrefix   string            `json:"SubscriptionPrefix"`
	SubscriptionDuration uint32            `json:"SubscriptionDuration"`
	SubscriptionInterval uint32            `json:"SubscriptionInterval"`
	Services             map[string]string `json:"Services"`
}

type Service struct {
	Name string `json:"name"`
	TCP  []int  `json:"tcp"`
	UDP  []int  `json:"udp"`
}

type TunaExit struct {
	config      Configuration
	wallet      *WalletSDK
	services    []Service
	serviceConn *cache.Cache
	common      *tuna.Common
	clientConn  *net.UDPConn
}

func NewTunaExit(config Configuration, wallet *WalletSDK) *TunaExit {
	var services []Service
	tuna.ReadJson("services.json", &services)

	return &TunaExit{
		config:      config,
		wallet:      wallet,
		services:    services,
		serviceConn: cache.New(time.Duration(config.UDPTimeout)*time.Second, time.Second),
	}
}

func (te *TunaExit) getServiceId(serviceName string) (byte, error) {
	for i, service := range te.services {
		if service.Name == serviceName {
			return byte(i), nil
		}
	}

	return 0, errors.New("Service " + serviceName + " not found")
}

func (te *TunaExit) handleSession(conn net.Conn) {
	session, _ := smux.Server(conn, nil)

	for {
		stream, err := session.AcceptStream()
		if err != nil {
			log.Println("Couldn't accept stream:", err)
			break
		}

		metadata := stream.Metadata()
		serviceId := metadata[0]
		portId := int(metadata[1])

		service, err := te.getService(serviceId)
		if err != nil {
			log.Println(err)
			tuna.Close(stream)
			continue
		}
		tcpPortsCount := len(service.TCP)
		udpPortsCount := len(service.UDP)
		var protocol tuna.Protocol
		var port int
		if portId < tcpPortsCount {
			protocol = tuna.TCP
			port = service.TCP[portId]
		} else if portId-tcpPortsCount < udpPortsCount {
			protocol = tuna.UDP
			portId -= tcpPortsCount
			port = service.UDP[portId]
		} else {
			log.Println("Wrong portId received:", portId)
			tuna.Close(stream)
			continue
		}

		serviceIP := te.config.Services[service.Name]
		host := serviceIP + ":" + strconv.Itoa(port)

		conn, err := net.DialTimeout(string(protocol), host, time.Duration(te.config.DialTimeout)*time.Second)
		if err != nil {
			log.Println("Couldn't connect to host", host, "with error:", err)
			tuna.Close(stream)
			continue
		}

		go tuna.Pipe(conn, stream)
		go tuna.Pipe(stream, conn)
	}

	tuna.Close(session)
	tuna.Close(conn)
}

func (te *TunaExit) listenTCP(port int) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: port})
	if err != nil {
		log.Println("Couldn't bind listener:", err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Couldn't accept client connection:", err)
				tuna.Close(conn)
				continue
			}

			go te.handleSession(conn)
		}
	}()
}

func (te *TunaExit) getService(serviceId byte) (*Service, error) {
	if int(serviceId) >= len(te.services) {
		return nil, errors.New("Wrong serviceId: " + strconv.Itoa(int(serviceId)))
	}
	return &te.services[serviceId], nil
}

func (te *TunaExit) getServiceConn(addr *net.UDPAddr, connId []byte, serviceId byte, portId byte) (*net.UDPConn, error) {
	connKey := addr.String() + ":" + tuna.GetConnIdString(connId)
	var conn *net.UDPConn
	var x interface{}
	var ok bool
	if x, ok = te.serviceConn.Get(connKey); !ok {
		service, err := te.getService(serviceId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		port := service.UDP[portId]
		conn, err = net.DialUDP("udp", nil, &net.UDPAddr{Port: port})
		if err != nil {
			log.Println("Couldn't connect to local UDP port", port, "with error:", err)
			tuna.Close(conn)
			return conn, err
		}

		te.serviceConn.Set(connKey, conn, cache.DefaultExpiration)

		prefix := []byte{connId[0], connId[1], serviceId, portId}
		go func() {
			serviceBuffer := make([]byte, 2048)
			for {
				n, err := conn.Read(serviceBuffer)
				if err != nil {
					log.Println("Couldn't receive data from service:", err)
					tuna.Close(conn)
					break
				}
				_, err = te.clientConn.WriteToUDP(append(prefix, serviceBuffer[:n]...), addr)
				if err != nil {
					log.Println("Couldn't send data to client:", err)
					tuna.Close(conn)
					break
				}
			}
		}()
	} else {
		conn = x.(*net.UDPConn)
	}

	return conn, nil
}

func (te *TunaExit) listenUDP(port int) {
	var err error
	te.clientConn, err = net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		log.Println("Couldn't bind listener:", err)
	}
	te.readUDP()
}

func (te *TunaExit) readUDP() {
	go func() {
		clientBuffer := make([]byte, 2048)
		for {
			n, addr, err := te.clientConn.ReadFromUDP(clientBuffer)
			if err != nil {
				log.Println("Couldn't receive data from client:", err)
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				}
				continue
			}
			serviceConn, err := te.getServiceConn(addr, clientBuffer[0:2], clientBuffer[2], clientBuffer[3])
			if err != nil {
				continue
			}
			_, err = serviceConn.Write(clientBuffer[4:n])
			if err != nil {
				log.Println("Couldn't send data to service:", err)
			}
		}
	}()
}

func (te *TunaExit) updateAllMetadata(ip string, tcpPort int, udpPort int) {
	for serviceName := range te.config.Services {
		serviceId, err := te.getServiceId(serviceName)
		if err != nil {
			log.Panicln(err)
		}
		service, err := te.getService(serviceId)
		if err != nil {
			log.Panicln(err)
		}
		tuna.UpdateMetadata(
			serviceName,
			serviceId,
			service.TCP,
			service.UDP,
			ip,
			tcpPort,
			udpPort,
			te.config.SubscriptionPrefix,
			te.config.SubscriptionDuration,
			te.config.SubscriptionInterval,
			te.wallet,
		)
	}
}

func (te *TunaExit) Start() {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Panicln("Couldn't get IP:", err)
	}

	te.listenTCP(te.config.ListenTCP)
	te.listenUDP(te.config.ListenUDP)

	te.updateAllMetadata(ip, te.config.ListenTCP, te.config.ListenUDP)
}

func (te *TunaExit) StartReverse(serviceName string) {
	serviceId, err := te.getServiceId(serviceName)
	if err != nil {
		log.Panicln(err)
	}
	service, err := te.getService(serviceId)
	if err != nil {
		log.Panicln(err)
	}

	reverseMetadata := &tuna.Metadata{}
	reverseMetadata.ServiceTCP = service.TCP
	reverseMetadata.ServiceUDP = service.UDP

	te.common = &tuna.Common{
		ServiceName:        "reverse",
		Wallet:             te.wallet,
		DialTimeout:        te.config.DialTimeout,
		ReverseMetadata:    reverseMetadata,
		SubscriptionPrefix: te.config.SubscriptionPrefix,
	}

	go func() {
		var tcpConn net.Conn
		for {
			err := te.common.CreateServerConn(true)
			if err != nil {
				log.Println("Couldn't connect to reverse entry:", err)
				time.Sleep(1 * time.Second)
				continue
			}

			udpPort := -1
			udpConn, _ := te.common.GetServerUDPConn(false)
			if udpConn != nil {
				_, udpPortString, _ := net.SplitHostPort(udpConn.LocalAddr().String())
				udpPort, _ = strconv.Atoi(udpPortString)
			}

			serviceMetadata := tuna.CreateRawMetadata(
				serviceId,
				service.TCP,
				service.UDP,
				"",
				-1,
				udpPort,
			)

			tcpConn, _ = te.common.GetServerTCPConn(false)
			_, err = tcpConn.Write(serviceMetadata)
			if err != nil {
				log.Println("Couldn't send metadata to reverse entry:", err)
				time.Sleep(1 * time.Second)
				continue
			}

			te.handleSession(tcpConn)

			if udpConn != nil {
				te.clientConn = udpConn
				te.readUDP()
			}
		}
	}()
}

func main() {
	config := Configuration{SubscriptionPrefix: tuna.DefaultSubscriptionPrefix}
	tuna.ReadJson("config.json", &config)

	Init()

	privateKey, _ := hex.DecodeString(config.PrivateKey)
	account, err := vault.NewAccountWithPrivatekey(privateKey)
	if err != nil {
		log.Panicln("Couldn't load account:", err)
	}

	wallet := NewWalletSDK(account)

	if config.Reverse {
		for serviceName := range config.Services {
			NewTunaExit(config, wallet).StartReverse(serviceName)
		}
	} else {
		NewTunaExit(config, wallet).Start()
	}

	select {}
}