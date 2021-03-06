/*
 * @Copyright Reserved By Janusec (https://www.janusec.com/).
 * @Author: U2
 * @Date: 2018-07-14 16:31:19
 * @Last Modified: U2, 2018-07-14 16:31:19
 */

package data

import (
	"encoding/json"

	"janusec/models"
	"janusec/utils"
)

func RPCGetSettings() []*models.Setting {
	rpcRequest := &models.RPCRequest{
		Action: "get_settings", Object: nil}
	resp, err := GetRPCResponse(rpcRequest)
	utils.CheckError("RPCGetSettings", err)
	rpcSettings := &models.RPCSettings{}
	if err = json.Unmarshal(resp, rpcSettings); err != nil {
		utils.CheckError("RPCGetSettings Unmarshal", err)
	}
	return rpcSettings.Object
}

func RPCGetOAuthConfig() *models.OAuthConfig {
	rpcRequest := &models.RPCRequest{
		Action: "get_oauth_conf", Object: nil}
	resp, err := GetRPCResponse(rpcRequest)
	utils.CheckError("RPCGetOAuthConfig", err)
	rpcOAuthConf := &models.RPCOAuthConfig{}
	if err = json.Unmarshal(resp, rpcOAuthConf); err != nil {
		utils.CheckError("RPCGetOAuthConfig Unmarshal", err)
	}
	//fmt.Println("RPCGetOAuthConfig", rpcOAuthConf.Object)
	return rpcOAuthConf.Object
}
