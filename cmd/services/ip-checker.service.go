package services

type IpCheckerService interface {
	GetCurrIp() (string, error)
	StoreIp(currIp string) error
	ReadOldIp() (string, error)
}
