package models

import (
	"time"
)

type Stats struct {
	ID           uint      `gorm:"primaryKey"`
	Timestamp    time.Time `gorm:"column:timestamp;type:timestamp with time zone;default:current_timestamp" json:"timestamp"`
	OrgName      string    `gorm:"column:org_name" json:"orgname"`
	SrcIP        string    `gorm:"column:src_ip" json:"scrip"`
	SrcPort      int       `gorm:"column:src_port" json:"srcport"`
	DstIP        string    `gorm:"column:dst_ip" json:"dstip"`
	DstPort      int       `gorm:"column:dst_port" json:"dstport"`
	PacketsCount int       `gorm:"column:packets_count" json:"packetscount"`
	BytesCount   int       `gorm:"column:bytes_count" json:"bytescount"`
}
type GeneratedStats struct {
	UserID uint
	Bucket string
	Name   string
}
