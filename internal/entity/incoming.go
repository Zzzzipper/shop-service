package entity

import (
	"time"
)

type (
	Partner struct {
		ID         uint64    `json:"id" gorm:"id"`
		CreateTime time.Time `json:"create_time" binding:"required" gorm:"create_time"`
		FullName   string    `json:"full_name" binding:"required" gorm:"full_name"`
		Url        string    `json:"url" binding:"required" gorm:"url"`
		ApiToken   string    `json:"api_token" binding:"required" gorm:"api_token"`
		Role       int32     `json:"role" binding:"required" gorm:"role"`
	}

	Merchant struct {
		ID         uint64    `json:"id" gorm:"id"`
		CreateTime time.Time `json:"create_time" binding:"required" gorm:"create_time"`
		FullName   string    `json:"full_name" binding:"required" gorm:"full_name"`
		Url        string    `json:"url" binding:"required" gorm:"url"`
		PartnerID  uint64    `json:"partner_id" binding:"required" gorm:"partner_id"`
	}

	Shop struct {
		ID         uint64    `json:"id" gorm:"id"`
		CreateTime time.Time `json:"create_time" binding:"required" gorm:"create_time"`
		FullName   string    `json:"full_name" binding:"required" gorm:"full_name"`
		MerchantID uint64    `json:"merchant_id" binding:"required" gorm:"merchant_id"`
		Login      string    `json:"login" binding:"required" gorm:"login"`
		Password   string    `json:"password" binding:"required" gorm:"password"`
		Url        string    `json:"url" binding:"required" gorm:"url"`
		Settings   string    `json:"settings" binding:"required" gorm:"settings"`
	}

	ShopInfo struct {
		ShopId     uint64 `json:"shop_id" binding:"required" gorm:"shop_id"`
		MerchantId uint64 `json:"merchant_id" binding:"required" gorm:"merchant_id"`
		PartnerId  uint64 `json:"partner_id" binding:"required" gorm:"partner_id"`
		Settings   string `json:"settings" binding:"required" gorm:"settings"`
	}
)
