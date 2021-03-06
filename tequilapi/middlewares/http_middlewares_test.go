/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestCacheControlHeadersAreAddedToResponse(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/not-important", nil)
	assert.NoError(t, err)
	respRecorder := httptest.NewRecorder()

	g := gin.Default()
	g.Use(ApplyCacheConfigMiddleware)
	g.ServeHTTP(respRecorder, req)

	assert.Equal(
		t,
		"no-cache, no-store, must-revalidate",
		respRecorder.Header().Get("Cache-Control"),
	)

}
