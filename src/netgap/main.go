package main

import (
	"log"
	"netgapxml"
)

var NetGapAllUserInfo netgapxml.AllUserInfo

var NetGapInnerAllRuleInfo netgapxml.AllRuleInfo

func main() {
	var ret int

	ret = netgapxml.ReadUserInfoFromFile(&NetGapAllUserInfo, "gap_user_info.xml")
	if ret < 0 {
		return
	}
	log.Println(NetGapAllUserInfo.UserInfo)

	ret = netgapxml.ReadRuleInfoFromFile(&NetGapInnerAllRuleInfo, "gap_config_info_eth0.xml")
	if ret < 0 {
		return
	}
	log.Println(NetGapInnerAllRuleInfo)

	return
}
