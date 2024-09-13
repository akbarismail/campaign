package campaigns

import "strings"

type CampaignDetailFormatter struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	Description      string                         `json:"description"`
	ImageUrl         string                         `json:"image_url"`
	TotalAmount      int                            `json:"total_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	UserId           int                            `json:"user_id"`
	Slug             string                         `json:"slug"`
	Perks            []string                       `json:"perks"`
	User             CampaignDetailUserFormatter    `json:"user"`
	Images           []CampaignDetailImageFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignDetailImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaigns) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.TotalAmount = campaign.TotalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserId = campaign.UserId
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	perks := []string{}
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	campaignDetailUserFormatter := CampaignDetailUserFormatter{}
	campaignDetailUserFormatter.Name = campaign.User.Name
	campaignDetailUserFormatter.ImageUrl = campaign.User.AvatarFileName

	campaignDetailFormatter.User = campaignDetailUserFormatter

	images := []CampaignDetailImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignDetailImageFormatter := CampaignDetailImageFormatter{}
		campaignDetailImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignDetailImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignDetailImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	TotalAmount      int    `json:"total_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaigns) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageUrl = ""
	campaignFormatter.TotalAmount = campaign.TotalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaigns) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
