package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang-/seckill/config"
	"golang-/seckill/secproxy/components"
	"sync"

	"github.com/astaxie/beego/logs"
)

var (
	FrequencyMgrInst *FrequencyMgr
)

func init() {
	FrequencyMgrInst = &FrequencyMgr{
		UserFrequency: make(map[int]*FrequencyLimit, components.INIT_INFO_SIZE),
		IPFrequency:   make(map[string]*FrequencyLimit, components.INIT_INFO_SIZE),
	}
}

type FrequencyLimit struct {
	Frequency int
	TimeStamp int64
}

type FrequencyMgr struct {
	UserFrequency map[int]*FrequencyLimit
	IPFrequency   map[string]*FrequencyLimit
	FrequencyLock sync.RWMutex
}

func (fm *FrequencyMgr) CalFrequency(userId int, visitTime int64) (frequency int) {
	fm.FrequencyLock.Lock()
	defer fm.FrequencyLock.Unlock()
	frequencylm, ok := fm.UserFrequency[userId]
	if !ok {
		fm.UserFrequency[userId] = &FrequencyLimit{
			Frequency: 1,
			TimeStamp: visitTime,
		}
		return 1
	}

	if frequencylm.TimeStamp != visitTime {
		frequencylm.Frequency = 1
		frequencylm.TimeStamp = visitTime
		return 1
	}

	frequencylm.Frequency++
	return frequencylm.Frequency

}

func (fm *FrequencyMgr) CalIPFrequency(ip string, visitTime int64) (frequency int) {
	fm.FrequencyLock.Lock()
	defer fm.FrequencyLock.Unlock()
	frequencylm, ok := fm.IPFrequency[ip]

	if !ok {
		fm.IPFrequency[ip] = &FrequencyLimit{
			Frequency: 1,
			TimeStamp: visitTime,
		}
		return 1
	}

	if frequencylm.TimeStamp != visitTime {
		frequencylm.Frequency = 1
		frequencylm.TimeStamp = visitTime
		return 1
	}

	frequencylm.Frequency++
	return frequencylm.Frequency
}

func isInblacklist(userid int, ip string) bool {
	//判断是否在黑名单中
	components.SKConfData.BlacklistRWLock.RLock()
	defer components.SKConfData.BlacklistRWLock.RUnlock()
	_, bok := components.SKConfData.IDBlacklist[userid]
	if bok {
		logs.Debug("user id[%v] in black list", userid)
		return true
	}

	_, bok = components.SKConfData.IPBlacklist[ip]
	if bok {
		logs.Debug("ip[%v] in black list", ip)
		return true
	}
	return false
}

func userCheck(req *config.SecRequest) (err error) {
	//检测跳转
	/*
		found := false
		for _, refer := range components.SKConfData.ReferWhitelist {
			if refer == req.ReferAddr {
				found = true
			}
		}

		if found == false {
			return errors.New("refer addr is invalid ")
		}
	*/
	authData := fmt.Sprintf("%d:%s", req.UserId, components.SKConfData.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		return errors.New("user auth sign dosen't match ")
	}

	return nil
}
