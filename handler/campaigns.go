package handler

import (
	"campaign/campaigns"
	"campaign/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaigns.Service
}

func NewCampaignHandler(campaignService campaigns.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	allCampaigns, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns.FormatCampaigns(allCampaigns))
	c.JSON(http.StatusOK, response)
}
