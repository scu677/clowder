package dbase

import (
	"fmt"
	"strings"
	"encoding/hex"
)

type UUID []byte

func (uuid UUID) String() string {
	time_low :=[]byte{uuid[3],uuid[2],uuid[1],uuid[0]}
	time_mid :=[]byte{uuid[5],uuid[4]}
	time_high_and_version :=[]byte{uuid[7],uuid[6]}
	clock_seq:=[]byte(uuid[8:10])
	node:=[]byte(uuid[10:])
	return fmt.Sprintf("%x-%x-%x-%x-%x",time_low, time_mid, time_high_and_version, clock_seq, node)
}

func ParseUUID(str string) (UUID, error) {
	s:=strings.Split(str,"-")
	if len(s)!=5 || len(s[0])!=4 || len(s[1])!=2 || len(s[2])!=2 || len(s[3])!=2 || len(s[4])!=6 {
		return nil,fmt.Errorf("Invalid UUID string.")
	}
	time_low,err0:=hex.DecodeString(s[0])
	time_mid,err1 :=hex.DecodeString(s[1])
	time_high_and_version,err2 := hex.DecodeString(s[2])
	clock_seq,err3:=hex.DecodeString(s[3])
	node,err4:=hex.DecodeString(s[4])
	if err0||err1||err2||err3||err4 { return nil,fmt.Errorf("Invalid UUID string.") }

	uuid:=[]byte{time_low[3],time_low[2],time_low[1],time_low[0],time_mid[1],time_mid[0],time_high_and_version[1],time_high_and_version[0]}
	uuid=append(uuid,clock_seq...)
	uuid=append(uuid,node...)
	return UUID(uuid),nil
}


type Hardwares map[string]UUID

func (hw Hardwares) String() string {
	s:=""
	for i,h:= range hw{
		s+=i+"\t"+h.String()+"\n"
	}
	if s!="" {
		s = s[:len(s)-1]
	}
	return s
}

