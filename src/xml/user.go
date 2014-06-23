
package netgapxml

import (
	"sync"
)


type SingleUserInfo struct {
	m_UserName string
	m_UserPasswdHash string
	m_UserPermission uint8
	m_UserComment string
}

type AllUserInfo struct {
	m_UserInfo []SSingleUserInfo
	m_Sem sync.Mutex
}

func ReadUserInfoFromFile(pUserInfo *SAllUserInfo) int {

	return 0;
}