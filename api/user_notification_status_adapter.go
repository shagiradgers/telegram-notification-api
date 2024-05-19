package notification

func (s UserNotificationStatus) ToBool() bool {
	return s == UserNotificationStatus_ON
}
