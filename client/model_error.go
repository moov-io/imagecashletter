/*
 * X9 API
 *
 * Moov X9 () implements an HTTP API for creating, parsing and validating X9 (Check21) files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Error struct {
	// An error message describing the problem intended for humans.
	Error string `json:"error"`
}