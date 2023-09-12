package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.CampaignService
}

func NewCampaignHandler(service campaign.CampaignService) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error to get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignById(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(input)
	if err != nil {
		response := helper.APIResponse("Error to get detail of Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.APIResponse("Error to get detail of Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail Of Campaigns", http.StatusOK, "success", campaign.FormatDetailCampaign(campaigns))
	c.JSON(http.StatusOK, response)
}
