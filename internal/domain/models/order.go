package models

import "time"

type Order struct {
	id              string
	title           string
	isCanSendNotify bool
	dateStart       time.Time
	dateEnd         time.Time
}

func (o Order) GetId() string {
	return o.id
}

func (o Order) GetName() string {
	return o.title
}

func (o Order) GetIsCanSendNotify() bool {
	return o.isCanSendNotify
}

func (o Order) GetDateStart() time.Time {
	return o.dateStart
}

func (o Order) GetDateEnd() time.Time {
	return o.dateEnd
}

func (o Order) SetName(name string) {
	o.title = name
}

func (o Order) SetIsCanSendNotify(isCanSendNotify bool) {
	o.isCanSendNotify = isCanSendNotify
}

func (o Order) SetDateStart(dateStart time.Time) {
	o.dateStart = dateStart
}

func (o Order) SetDateEnd(dateEnd time.Time) {
	o.dateEnd = dateEnd
}
