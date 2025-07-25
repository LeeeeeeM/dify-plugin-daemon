// Code generated by controller generator. DO NOT EDIT.

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/langgenius/dify-plugin-daemon/internal/service"
	"github.com/langgenius/dify-plugin-daemon/internal/types/app"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/plugin_entities"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/requests"
)

func GetAuthorizationURL(config *app.Config) gin.HandlerFunc {
	type request = plugin_entities.InvokePluginRequest[requests.RequestOAuthGetAuthorizationURL]

	return func(c *gin.Context) {
		BindPluginDispatchRequest(
			c,
			func(itr request) {
				service.GetAuthorizationURL(&itr, c, config.PluginMaxExecutionTimeout)
			},
		)
	}
}

func GetCredentials(config *app.Config) gin.HandlerFunc {
	type request = plugin_entities.InvokePluginRequest[requests.RequestOAuthGetCredentials]

	return func(c *gin.Context) {
		BindPluginDispatchRequest(
			c,
			func(itr request) {
				service.GetCredentials(&itr, c, config.PluginMaxExecutionTimeout)
			},
		)
	}
}

func RefreshCredentials(config *app.Config) gin.HandlerFunc {
	type request = plugin_entities.InvokePluginRequest[requests.RequestOAuthRefreshCredentials]

	return func(c *gin.Context) {
		BindPluginDispatchRequest(
			c,
			func(itr request) {
				service.RefreshCredentials(&itr, c, config.PluginMaxExecutionTimeout)
			},
		)
	}
}
