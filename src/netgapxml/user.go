package netgapxml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"sync"
)

var NetGapAllUserInfo AllUserInfo

type SingleUserInfo struct {
	UserName       string `xml:"UserName"`
	UserPasswdHash string `xml:"PasswordHash"`
	UserPermission uint8  `xml:"Permission"`
	UserComment    string `xml:"Comments"`
}

type AllUserInfo struct {
	UserInfo      []SingleUserInfo `xml:"UserInfo"`
	UserInfoMutex sync.Mutex
}

func ReadUserInfoFromFile(pUserInfo *AllUserInfo, FilePath string) int {
	data, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Printf("error: %v", err)
		return -1
	}

	pUserInfo.UserInfoMutex.Lock()
	err = xml.Unmarshal(data, pUserInfo)
	pUserInfo.UserInfoMutex.Unlock()
	if err != nil {
		log.Printf("error: %v", err)
		return -1
	}

	return 0
}
