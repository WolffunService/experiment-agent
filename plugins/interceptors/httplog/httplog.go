/****************************************************************************
 * Copyright 2020, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

// Package httplog //
package httplog

import (
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog/log"

	"github.com/WolffunGame/experiment-agent/plugins/interceptors"
)

type httpLog struct{}

func (h *httpLog) Handler() func(http.Handler) http.Handler {
	return httplog.Handler(log.Logger)
}

func init() {
	interceptors.Add("httplog", func() interceptors.Interceptor {
		return &httpLog{}
	})
}
