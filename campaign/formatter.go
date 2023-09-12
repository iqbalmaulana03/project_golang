package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_dercription"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_dercription"`
	ImageUrl         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaignImagesFormatter `jsong:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	campaignFormatter := CampaignDetailFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignFormatter.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImagesFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignFormatter.Images = images

	return campaignFormatter
}
