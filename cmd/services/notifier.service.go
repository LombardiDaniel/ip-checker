package services

type NotifierService interface {
	SendNotification(msg string)				error
}
