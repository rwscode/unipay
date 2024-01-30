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

package base

import (
	"errors"
	"fmt"
	"github.com/rwscode/unipay/deps/db"
	"github.com/rwscode/unipay/deps/models"
)

func CheckChannelExist(channelId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.Channel)).Where("id=?", channelId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("支付通道[%d]不存在", channelId))
	}
	return
}

func CheckChannelParamExist(channelParamId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.ChannelParam)).Where("id=?", channelParamId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("支付通道参数[%s]不存在", channelParamId))
	}
	return
}

func CheckOrderExist(transactionId string) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.ChannelParam)).Where("id=?", transactionId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("支付订单[%s]不存在", transactionId))
	}
	return
}