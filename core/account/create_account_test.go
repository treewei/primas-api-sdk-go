/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package account

import (
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestCreateAccount(t *testing.T) {
	address := "0xd75407ad8cabeeebfed78c4f3794208b3339fbf4"
	name := "kevin"
	abstract := "kevin test"
	avatar := "06354a3ee794cf210e0db0d6729710ee047227d679668de697431fdbd1232ffc"
	account_id := ""
	sub_account_id := ""
	created := time.Now().Unix()
	extra_hash := "test_value"

	signature, preObj, err := CreateAccount_SignatureStr(address, name, abstract, avatar, account_id, sub_account_id, int(created), extra_hash)
	if err != nil {
		t.Errorf("CreateAccount_SignatureStr error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("CreateAccount_SignatureStr preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("CreateAccount_SignatureStr signature value is empty")
		return
	}

	// mock Sign
	privateKey := tool.GetClientPrivateKey()
	signValue, err := tool.Sign([]byte(signature), privateKey)
	if err != nil {
		t.Errorf("Sign error %v:", err.Error())
		return
	}
	//

	createAccount, err := CreateAccount(signValue, preObj)
	if err != nil {
		t.Errorf("CreateAccount error %v:", err.Error())
		return
	}

	if createAccount != nil {
		t.Logf("CreateAccount response value:%v", createAccount)
		if createAccount.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("CreateAccount response error:%v", createAccount.ResultMsg)
			return
		}
		if createAccount.Data != nil {
			t.Logf("CreateAccount response value:%v", createAccount.Data)
		} else {
			t.Logf("CreateAccount response value don't find ")
		}
	}
}

func TestCreateAccount_sub(t *testing.T) {
	address := "0xd75407ad8cabeeebfed78c4f3794208b3339fbf4"
	name := "yoyou"
	abstract := "yoyou test"
	avatar := "06354a3ee794cf210e0db0d6729710ee047227d679668de697431fdbd1232ffc"
	account_id := "2fe750b56cc7f0949803fdd94075337cfb64a8f48c44b542ac4daaa52b799886"
	sub_account_id := "a_0006"
	created := time.Now().Unix()
	extra_hash := "test_value"

	signature, preObj, err := CreateAccount_SignatureStr(address, name, abstract, avatar, account_id, sub_account_id, int(created), extra_hash)
	if err != nil {
		t.Errorf("TestCreateAccount_sub error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestCreateAccount_sub preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestCreateAccount_sub signature value is empty")
		return
	}

	// mock Sign
	privateKey := tool.GetClientPrivateKey()
	signValue, err := tool.Sign([]byte(signature), privateKey)
	if err != nil {
		t.Errorf("Sign error %v:", err.Error())
		return
	}
	//

	createAccount, err := CreateAccount(signValue, preObj)
	if err != nil {
		t.Errorf("TestCreateAccount_sub error %v:", err.Error())
		return
	}

	if createAccount != nil {
		t.Logf("TestCreateAccount_sub response value:%v", createAccount)
		if createAccount.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestCreateAccount_sub response error:%v", createAccount.ResultMsg)
			return
		}
		if createAccount.Data != nil {
			t.Logf("TestCreateAccount_sub response value:%v", createAccount.Data)
		} else {
			t.Logf("TestCreateAccount_sub response value don't find ")
		}
	}
}
