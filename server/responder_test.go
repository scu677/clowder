package server

import (
	"fmt"
	"github.com/musec/clowder/dbase"
	"github.com/musec/clowder/pxedhcp"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var dhcpPackets = [][]byte{
	[]byte{
		0x01, 0x01, 0x06, 0x00, 0x46, 0x34, /* .>....F4 */
		0x6e, 0x00, 0x00, 0x20, 0x80, 0x00, 0x00, 0x00, /* n.. .... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0xa8, /* ......D. */
		0x42, 0x34, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00, /* B4n..... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, /* ......c. */
		0x53, 0x63, 0x35, 0x01, 0x01, 0x37, 0x18, 0x01, /* Sc5..7.. */
		0x02, 0x03, 0x05, 0x06, 0x0b, 0x0c, 0x0d, 0x0f, /* ........ */
		0x10, 0x11, 0x12, 0x2b, 0x36, 0x3c, 0x43, 0x80, /* ...+6<C. */
		0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x39, /* .......9 */
		0x02, 0x04, 0xec, 0x61, 0x11, 0x00, 0x44, 0x45, /* ...a..DE */
		0x4c, 0x4c, 0x36, 0x00, 0x10, 0x58, 0x80, 0x43, /* LL6..X.C */
		0xb5, 0xc0, 0x4f, 0x53, 0x35, 0x32, 0x5d, 0x02, /* ..OS52]. */
		0x00, 0x00, 0x5e, 0x03, 0x01, 0x02, 0x01, 0x3c, /* ..^....< */
		0x20, 0x50, 0x58, 0x45, 0x43, 0x6c, 0x69, 0x65, /*  PXEClie */
		0x6e, 0x74, 0x3a, 0x41, 0x72, 0x63, 0x68, 0x3a, /* nt:Arch: */
		0x30, 0x30, 0x30, 0x30, 0x30, 0x3a, 0x55, 0x4e, /* 00000:UN */
		0x44, 0x49, 0x3a, 0x30, 0x30, 0x32, 0x30, 0x30, /* DI:00200 */
		0x31, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* 1....... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ...... */
	},
	[]byte{
		0x01, 0x01, 0x06, 0x00, 0x00, 0x00, /* .[...... */
		0x25, 0x9b, 0x00, 0x00, 0x00, 0x00, 0xc0, 0xa8, /* %....... */
		0x01, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .d...... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc8, 0x00, /* ........ */
		0x84, 0x66, 0xd1, 0xc0, 0x00, 0x00, 0x00, 0x00, /* .f...... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, /* ......c. */
		0x53, 0x63, 0x35, 0x01, 0x03, 0x39, 0x02, 0x04, /* Sc5..9.. */
		0x80, 0x3d, 0x19, 0x00, 0x63, 0x69, 0x73, 0x63, /* .=..cisc */
		0x6f, 0x2d, 0x63, 0x38, 0x30, 0x30, 0x2e, 0x38, /* o-c800.8 */
		0x34, 0x36, 0x36, 0x2e, 0x64, 0x31, 0x63, 0x30, /* 466.d1c0 */
		0x2d, 0x56, 0x6c, 0x31, 0x33, 0x04, 0x00, 0x00, /* -Vl13... */
		0x02, 0x58, 0x37, 0x08, 0x01, 0x06, 0x0f, 0x2c, /* .X7...., */
		0x03, 0x21, 0x96, 0x2b, 0xff, /* .!.+. */
	},
	[]byte{
		0x01, 0x01, 0x06, 0x00, 0x46, 0x34, /* .R....F4 */
		0x6e, 0x00, 0x00, 0x20, 0x80, 0x00, 0x00, 0x00, /* n.. .... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0xa8, /* ......D. */
		0x42, 0x34, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00, /* B4n..... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, /* ......c. */
		0x53, 0x63, 0x35, 0x01, 0x03, 0x32, 0x04, 0xc0, /* Sc5..2.. */
		0xa8, 0x01, 0x15, 0x37, 0x18, 0x01, 0x02, 0x03, /* ...7.... */
		0x05, 0x06, 0x0b, 0x0c, 0x0d, 0x0f, 0x10, 0x11, /* ........ */
		0x12, 0x2b, 0x36, 0x3c, 0x43, 0x80, 0x81, 0x82, /* .+6<C... */
		0x83, 0x84, 0x85, 0x86, 0x87, 0x39, 0x02, 0x04, /* .....9.. */
		0xec, 0x36, 0x04, 0xc0, 0xa8, 0x01, 0x01, 0x61, /* .6.....a */
		0x11, 0x00, 0x44, 0x45, 0x4c, 0x4c, 0x36, 0x00, /* ..DELL6. */
		0x10, 0x58, 0x80, 0x43, 0xb5, 0xc0, 0x4f, 0x53, /* .X.C..OS */
		0x35, 0x32, 0x5d, 0x02, 0x00, 0x00, 0x5e, 0x03, /* 52]...^. */
		0x01, 0x02, 0x01, 0x3c, 0x20, 0x50, 0x58, 0x45, /* ...< PXE */
		0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x41, /* Client:A */
		0x72, 0x63, 0x68, 0x3a, 0x30, 0x30, 0x30, 0x30, /* rch:0000 */
		0x30, 0x3a, 0x55, 0x4e, 0x44, 0x49, 0x3a, 0x30, /* 0:UNDI:0 */
		0x30, 0x32, 0x30, 0x30, 0x31, 0xff, 0x00, 0x00, /* 02001... */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ...... */
	},
}

func TestResponder(t *testing.T) {

	//Create server
	serverIP := net.IP{192, 168, 1, 1}
	serverMask := net.IP{255, 255, 255, 0}
	duration := time.Minute * 10
	hostname, _ := os.Hostname()
	dns := net.IP{192, 168, 1, 1}
	router := net.IP{192, 168, 1, 1}
	domainName := "musec.engr.mun.ca"
	s := NewServer(serverIP, serverMask, 5000, duration, hostname, dns, router, domainName)
	s.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	s.Pxe = make(dbase.PxeTable, 10, 10)

	for i := range dhcpPackets {
		fmt.Println("******************************************\nTest case", i, "\n******************************************\n")
		s.MachineLeases = dbase.NewLeases(net.IP{192, 168, 1, 10}, 50)
		s.DeviceLeases = dbase.NewLeases(net.IP{192, 168, 1, 100}, 50)

		req := pxedhcp.Packet(dhcpPackets[i])
		options := req.ParseOptions()
		mac := req.GetHardwareAddr()
		uuid, ok := options[97]
		if ok {
			s.Pxe[0] = dbase.PxeRecord{uuid, "blackmarsh1", "pxeboot"}
			pool := s.MachineLeases
			ip := net.IP{192, 168, 1, 21}
			pool.SetIPStat(ip, dbase.RESERVED)
			//pool.SetMac(ip,mac)
		} else {
			pool := s.DeviceLeases
			ip := net.IP{192, 168, 1, 100}
			pool.SetIPStat(ip, dbase.RESERVED)
			pool.SetMac(ip, mac)
		}
		rep := s.DHCPResponder(req)

		if rep != nil {
			fmt.Println(rep)
			repOpt := rep.ParseOptions()
			for k := range repOpt {
				switch k {
				case 15, 17:
					fmt.Println(k, "\t", string(repOpt[k]))
				case 1, 3, 6, 54:
					fmt.Println(k, "\t", net.IP(repOpt[k]))
				default:
					fmt.Printf("%d\t%x \n", k, repOpt[k])
				}
			}
		}
	}
	fmt.Println(s.NewHardware)
}
