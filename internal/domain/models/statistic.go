package models

import "time"

type Statistic struct {
	UID          string
	UserUID      string
	AIServiceUID string
	Amount       float64
	CreatedAt    time.Time
}

type StatisticFilter struct {
	UserUID      *string
	AIServiceUID *string
	Order        *string
	PageNumber   *uint32
	PageSize     *uint32
}

type StatisticRead struct {
	UserUID      string
	AIServiceUID string
	Count        uint32
	FullAmount   float64
}

type StatisticWrite struct {
	UserUID      string
	AIServiceUID string
}
