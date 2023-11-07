/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

import (
	"encoding/json"
)

// AudioMessage
// AudioMessage
// https://developers.line.biz/en/reference/messaging-api/#audio-message
type AudioMessage struct {
	Message

	/**
	 * Get QuickReply
	 */
	QuickReply *QuickReply `json:"quickReply,omitempty"`

	/**
	 * Get Sender
	 */
	Sender *Sender `json:"sender,omitempty"`

	/**
	 * Get OriginalContentUrl
	 */
	OriginalContentUrl string `json:"originalContentUrl"`

	/**
	 * Get Duration
	 */
	Duration int64 `json:"duration"`
}

// MarshalJSON customizes the JSON serialization of the AudioMessage struct.
func (r *AudioMessage) MarshalJSON() ([]byte, error) {

	type Alias AudioMessage
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "audio",
	})
}