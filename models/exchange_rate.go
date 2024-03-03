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

package models

type (
	ExchangeRate       = UnipayExchangeRate
	UnipayExchangeRate struct {
		Id   uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                  // id
		Rate string `gorm:"column:rate;type:varchar(20);not null;default:'';comment:美元对人民币汇率 CNY=汇率*金额" json:"rate"` // 美元对人民币汇率 CNY=汇率*金额
	}
)