package main

import (
	"log"
	"netgapxml"
	"strconv"
	"time"
)

func main() {
	var ret, i int

	ret = netgapxml.ReadUserInfoFromFile(&netgapxml.NetGapAllUserInfo, "gap_user_info.xml")
	if ret >= 0 {
		log.Println(netgapxml.NetGapAllUserInfo)
	}

	for i = 0; i < netgapxml.NetGapEthNum; i++ {
		numstr := strconv.Itoa(i)
		str := "gap_config_info_eth" + numstr + ".xml"
		ret = netgapxml.ReadRuleInfoFromFile(&netgapxml.NetGapInnerAllEthInfo[i], str)
		if ret >= 0 {
			log.Println(netgapxml.NetGapInnerAllEthInfo[i])
		}
	}

	for {
		time.Sleep(time.Second)
	}

	return
}
