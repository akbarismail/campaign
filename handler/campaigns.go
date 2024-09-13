package handler

import (
	"campaign/campaigns"
	"campaign/helper"
	"campaign/user"
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaigns.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMsg := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", campaigns.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaigns.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignById(input)
	if err != nil {
		response := helper.APIResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail campaign", http.StatusOK, "success", campaigns.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
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
