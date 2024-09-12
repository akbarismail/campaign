package campaigns

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
