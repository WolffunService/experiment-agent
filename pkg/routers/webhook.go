/****************************************************************************
 * Copyright 2019-2020, Optimizely, Inc. and contributors                   *
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

// Package routers //
package routers

import (
	"github.com/WolffunGame/experiment-agent/config"
	"github.com/WolffunGame/experiment-agent/pkg/handlers"

	"github.com/go-chi/chi"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/WolffunGame/experiment-agent/pkg/optimizely"
)

// NewWebhookRouter returns HTTP API router
func NewWebhookRouter(optlyCache optimizely.Cache, conf config.WebhookConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimw.AllowContentType("application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	webhookAPI := handlers.NewWebhookHandler(optlyCache, conf.Projects)

	r.Post("/webhooks/optimizely", webhookAPI.HandleWebhook)
	return r
}
