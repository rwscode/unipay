// Copyright 2024 unipay Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e20svc

type (
	OrderPayHtmlReq struct {
		OrderId            string `form:"order_id"`              // 订单id
		CheckOrderStateUrl string `form:"check_order_state_url"` // 检查订单状态Url
		RedirectUrl        string `form:"redirect_url"`          // 支付成功跳转Url
		Platform           string `form:"platform"`              // 平台 1Android 2iOSWeb
	}
	E20HtmlReq struct {
		OrderId            string // 订单id
		Protocol           string // 协议 erc20/trc20
		Amount             string // 支付金额
		Address            string // 钱包地址
		ExpirationTime     string // 订单失效时间毫秒时间戳
		CheckOrderStateUrl string // 检查订单状态Url
		RedirectUrl        string // 支付成功跳转Url
		Platform           string // 平台 1Android 2iOSWeb
	}
)
