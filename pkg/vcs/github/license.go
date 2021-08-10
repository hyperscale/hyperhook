// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SPDXID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

/*
{
	"key": "mit",
	"name": "MIT License",
	"spdx_id": "MIT",
	"url": "https://api.github.com/licenses/mit",
	"node_id": "MDc6TGljZW5zZTEz"
}
*/
