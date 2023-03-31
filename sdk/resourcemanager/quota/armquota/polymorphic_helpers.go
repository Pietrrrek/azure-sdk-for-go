//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armquota

import "encoding/json"

func unmarshalLimitJSONObjectClassification(rawMsg json.RawMessage) (LimitJSONObjectClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b LimitJSONObjectClassification
	switch m["limitObjectType"] {
	case string(LimitTypeLimitValue):
		b = &LimitObject{}
	default:
		b = &LimitJSONObject{}
	}
	return b, json.Unmarshal(rawMsg, b)
}