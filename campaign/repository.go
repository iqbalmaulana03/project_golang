package campaign

import "gorm.io/gorm"

type CampaignRepository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(id int) (Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images,is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images,is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindById(id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", id).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
