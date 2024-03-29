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

package pay

type Resp struct {
	OrderId    string `json:"order_id"`     // 订单id
	PayPageUrl string `json:"pay_page_url"` // 页面url
	PayQrUrl   string `json:"pay_qr_url"`   // 二维码url
	Message    string `json:"message"`      // 支付信息
}
