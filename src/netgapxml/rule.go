package netgapxml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"sync"
)

type SingleSrcAddrInfo struct {
	SrcAddrIndex      uint32 `xml:"SrcAddrIndex"`
	SrcDevName        string `xml:"SrcDevName"`
	SrcIPAddr         string `xml:"SrcIPAddr"`
	SrcNetMask        string `xml:"NetMask"`
	SrcMACAddr        string `xml:"SrcMACAddr"`
	SrcNATEnable      string `xml:"SrcNATEnable"`
	SrcIPAddrAfterNAT string `xml:"SrcIPAddrAfterNAT"`
	GwInfoTableIndex  string `xml:"GatewayTableInfoIndex"`
}

type AllSrcAddrInfo struct {
	SrcAddrInfo []SingleSrcAddrInfo `xml:"SrcAddrInfo"`
}

type SingleDstAddrInfo struct {
	DstAddrIndex      uint32 `xml:"DstAddrIndex"`
	DstDevName        string `xml:"DstDevName"`
	DstIPAddr         string `xml:"DstIPAddr"`
	DstNetMask        string `xml:"NetMask"`
	DstMACAddr        string `xml:"DstMACAddr"`
	DstNATEnable      string `xml:"DstNATEnable"`
	DstIPAddrAfterNAT string `xml:"DstIPAddrAfterNAT"`
	GwInfoTableIndex  string `xml:"GatewayTableInfoIndex"`
}

type AllDstAddrInfo struct {
	DstAddrInfo []SingleDstAddrInfo `xml:"DstAddrInfo"`
}

type SingleGwAddrInfo struct {
	GwAddrIndex      uint32 `xml:"GatewayAddrIndex"`
	GwDevName        string `xml:"GatewayDevName"`
	GwIPAddr         string `xml:"GatewayIPAddr"`
	GwNetMask        string `xml:"NetMask"`
	GwMACAddr        string `xml:"GatewayMACAddr"`
	GwNATEnable      string `xml:"GatewayNATEnable"`
	GwIPAddrAfterNAT string `xml:"GatewayIPAddrAfterNAT"`
}

type AllGwAddrInfo struct {
	GwAddrInfo []SingleGwAddrInfo `xml:"GatewayAddrInfo"`
}

type SingleMatchInfo struct {
	MatchIndex        uint32 `xml:"MatchRegularIndex"`
	SrcInfoTableIndex uint32 `xml:"SrcAddrTableInfoIndex"`
	DstInfoTableIndex uint32 `xml:"DstAddrTableInfoIndex"`
	ProtocolType      uint32 `xml:"ProtocolType"`
	SrcPortBegin      uint32 `xml:"SrcPortBegin"`
	SrcPortEnd        uint32 `xml:"SrcPortEnd"`
	DstPortBegin      uint32 `xml:"DstPortBegin"`
	DstPortEnd        uint32 `xml:"DstPortEnd"`
	TwoWayEnable      uint32 `xml:"TwoWayEnable"`
}

type AllMatchInfo struct {
	MatchInfo []SingleMatchInfo `xml:"MatchRegular"`
}

type AllRuleInfo struct {
	EthIndex      uint32         `xml:"EthIndex"`
	SrcAddrRule   AllSrcAddrInfo `xml:"SrcAddrTableInfo"`
	DstAddrRule   AllDstAddrInfo `xml:"DstAddrTableInfo"`
	GwAddrRule    AllGwAddrInfo  `xml:"GatewayTableInfo"`
	MatchRule     AllMatchInfo   `xml:"MatchRegularTableInfo"`
	RuleInfoMutex sync.Mutex
}

func ReadRuleInfoFromFile(pRuleInfo *AllRuleInfo, FilePath string) int {
	data, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Printf("error: %v", err)
		return -1
	}

	pRuleInfo.RuleInfoMutex.Lock()
	err = xml.Unmarshal(data, pRuleInfo)
	pRuleInfo.RuleInfoMutex.Unlock()
	if err != nil {
		log.Printf("error: %v", err)
		return -1
	}

	return 0
}
