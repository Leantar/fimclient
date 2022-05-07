package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/Leantar/fimproto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
	"path/filepath"
	"strconv"
)

type Config struct {
	Host        string `yaml:"host"`
	Port        int64  `yaml:"port"`
	CertFile    string `yaml:"cert_file"`
	CertKeyFile string `yaml:"cert_key_file"`
	CaFile      string `yaml:"ca_file"`
}

type Client struct {
	conn   *grpc.ClientConn
	client proto.FimClient
	conf   Config
}

func NewConnectedClient(config Config) (*Client, error) {
	c := &Client{
		conf: config,
	}

	err := c.connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return c, nil
}

func (c *Client) connect() error {
	creds, err := createGrpcCredentials(c.conf.CertFile, c.conf.CertKeyFile, c.conf.CaFile)
	if err != nil {
		return err
	}

	address := net.JoinHostPort(c.conf.Host, strconv.FormatInt(c.conf.Port, 10))

	c.conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}

	c.client = proto.NewFimClient(c.conn)

	return nil
}

func createGrpcCredentials(certPath, keyPath, caPath string) (credentials.TransportCredentials, error) {
	caFile, err := filepath.Abs(caPath)
	if err != nil {
		return nil, err
	}

	caBytes, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	certFile, err := filepath.Abs(certPath)
	if err != nil {
		return nil, err
	}

	keyFile, err := filepath.Abs(keyPath)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(caBytes)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s", caFile)
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{cert},
	}), nil
}
