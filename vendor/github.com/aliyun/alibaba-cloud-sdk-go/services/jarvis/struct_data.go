package jarvis

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in jarvis response
type Data struct {
	GmtExpire        string                 `json:"GmtExpire" xml:"GmtExpire"`
	TacticsName      string                 `json:"TacticsName" xml:"TacticsName"`
	SrcPort          int                    `json:"SrcPort" xml:"SrcPort"`
	RiskType         string                 `json:"RiskType" xml:"RiskType"`
	FeedBack         int                    `json:"FeedBack" xml:"FeedBack"`
	Reason           string                 `json:"Reason" xml:"Reason"`
	GmtCreate        string                 `json:"GmtCreate" xml:"GmtCreate"`
	RegionId         string                 `json:"RegionId" xml:"RegionId"`
	RiskInstance     string                 `json:"RiskInstance" xml:"RiskInstance"`
	IgnoreTime       string                 `json:"IgnoreTime" xml:"IgnoreTime"`
	DstIP            string                 `json:"DstIP" xml:"DstIP"`
	Product          string                 `json:"Product" xml:"Product"`
	Protocol         string                 `json:"Protocol" xml:"Protocol"`
	PunishCount      int                    `json:"PunishCount" xml:"PunishCount"`
	UpdateTime       string                 `json:"UpdateTime" xml:"UpdateTime"`
	SrcIP            string                 `json:"SrcIP" xml:"SrcIP"`
	DstPort          int                    `json:"DstPort" xml:"DstPort"`
	AutoConfig       int                    `json:"AutoConfig" xml:"AutoConfig"`
	PunishType       string                 `json:"PunishType" xml:"PunishType"`
	GroupId          int                    `json:"GroupId" xml:"GroupId"`
	PunishResult     string                 `json:"PunishResult" xml:"PunishResult"`
	RiskDescribe     string                 `json:"RiskDescribe" xml:"RiskDescribe"`
	RiskId           int                    `json:"RiskId" xml:"RiskId"`
	Status           string                 `json:"Status" xml:"Status"`
	GmtRealExpire    string                 `json:"GmtRealExpire" xml:"GmtRealExpire"`
	SrcUid           string                 `json:"SrcUid" xml:"SrcUid"`
	InstanceList     []string               `json:"InstanceList" xml:"InstanceList"`
	Items            []Item                 `json:"Items" xml:"Items"`
	RdsWhitelistRisk []RdsWhitelistRiskItem `json:"RdsWhitelistRisk" xml:"RdsWhitelistRisk"`
	EcsSecGroupRisk  []EcsSecGroupRiskItem  `json:"EcsSecGroupRisk" xml:"EcsSecGroupRisk"`
}
