// @BeeOverwrite YES
// @BeeGenerateTime 20200911_202853
package controllers

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
