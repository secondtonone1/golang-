package service

import (
	"golang-/seckill/secproxy/components"
	"sync"
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
