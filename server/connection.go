package server

import (
	"github.com/spf13/viper"
	"net"
)

//
// A connection to a server (wraps an internal TCP connection).
//
type Connection struct {
	HasLogger

	connection net.Conn
}

// Connect to a server named in a Viper configuration.
func Connect(config *viper.Viper) (Connection, error) {
	var c Connection

	err := c.InitLog("")
	if err != nil {
		return c, err
	}

	host := config.GetString("server.host")
	port := config.GetString("server.controlPort")

	server := host + ":" + port
	c.connection, err = net.Dial("tcp", server)
	if err != nil {
		return c, err
	}

	return c, nil
}

// Connect to a server or terminate the application.
func ConnectOrDie(config *viper.Viper) Connection {
	c, err := Connect(config)

	if err != nil {
		c.FatalError(err)
	}

	return c
}

func (c Connection) SendCommand(cmd string) (string, error) {
	var buffer [2048]byte

	if _, err := c.connection.Write([]byte(cmd)); err != nil {
		return "", err
	}

	n, err := c.connection.Read(buffer[0:])
	if err != nil {
		return "", err
	}

	if cmd != "CLOSECONN" {
		if _, err := c.connection.Write([]byte("CLOSECONN")); err != nil {
			return "", err
		}
	}
	return string(buffer[:n]), nil
}
