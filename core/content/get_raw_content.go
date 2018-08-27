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

package content

import (
	"errors"
	"net/http"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func GetRawContent(content_id string) (*http.Response, error) {
	if content_id == "" {
		return nil, errors.New("content_id is empty")
	}

	url := config.Gogal_Server + `/content/` + content_id + `/raw`
	queryParams := make(map[string]interface{}, 0)

	response, err := tool.Http_Get_Direct(url, queryParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
