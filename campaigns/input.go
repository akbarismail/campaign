package campaigns

import "campaign/user"

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	TotalAmount      int    `json:"total_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
