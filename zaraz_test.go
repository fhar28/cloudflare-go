package cloudflare

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")

		fmt.Fprint(w, `{
			"errors": [],
			"messages": [],
			"success": true,
			"result": {
					"dataLayer": true,
					"debugKey": "cheese",
					"dlp": [],
					"historyChange": true,
					"settings": {
						"autoInjectScript": true
					},
					"tools": {
						"PBQr": {
							"blockingTriggers": [],
							"component": "html",
							"defaultFields": {},
							"enabled": true,
							"mode": {
							  "cloud": false,
							  "ignoreSPA": true,
							  "light": false,
							  "sample": false,
							  "segment": {
								"end": 100,
								"start": 0
							  },
							  "trigger": "pageload"
							},
							"name": "Custom HTML",
							"neoEvents": [],
							"permissions": [
							  "execute_unsafe_scripts"
							],
							"settings": {},
							"type": "component"
						  }
					},
					"triggers": {
						"Pageview": {
							"clientRules": [],
							"description": "All page loads",
							"excludeRules": [],
							"loadRules": [
							  {
								"match": "{{ client.__zarazTrack }}",
								"op": "EQUALS",
								"value": "Pageview"
							  }
							],
							"name": "Pageview",
							"system": "pageload"
						}
					},
					"variables": {},
					"zarazVersion": 44
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/config", handler)
	expected := ZarazConfigResponse{
		Result: ZarazConfig{
			"dataLayer":     true,
			"debugKey":      "cheese",
			"dlp":           []interface{}{},
			"historyChange": true,
			"settings": map[string]interface{}{
				"autoInjectScript": true,
			},
			"tools": map[string]interface{}{
				"PBQr": map[string]interface{}{
					"blockingTriggers": []interface{}{},
					"component":        "html",
					"defaultFields":    map[string]interface{}{},
					"enabled":          true,
					"mode": map[string]interface{}{
						"cloud":     false,
						"ignoreSPA": true,
						"light":     false,
						"sample":    false,
						"segment": map[string]interface{}{
							"end":   float64(100),
							"start": float64(0),
						},
						"trigger": "pageload",
					},
					"name":        "Custom HTML",
					"neoEvents":   []interface{}{},
					"permissions": []interface{}{"execute_unsafe_scripts"},
					"settings":    map[string]interface{}{},
					"type":        "component",
				},
			},
			"triggers": map[string]interface{}{
				"Pageview": map[string]interface{}{
					"clientRules": []interface{}{},
					"description": "All page loads",
					"loadRules": []interface{}{map[string]interface{}{
						"match": "{{ client.__zarazTrack }}",
						"op":    "EQUALS",
						"value": "Pageview",
					}},
					"excludeRules": []interface{}{},
					"name":         "Pageview",
					"system":       "pageload",
				},
			},
			"variables":    map[string]interface{}{},
			"zarazVersion": float64(44),
		},
		Response: Response{
			Success:  true,
			Messages: []ResponseInfo{},
			Errors:   []ResponseInfo{},
		},
	}

	actual, err := client.GetZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"errors": [],
			"messages": [],
			"success": true,
			"result": {
					"dataLayer": true,
					"debugKey": "cheese",
					"dlp": [],
					"historyChange": true,
					"settings": {
						"autoInjectScript": true
					},
					"tools": {
						"PBQr": {
							"blockingTriggers": [],
							"component": "html",
							"defaultFields": {},
							"enabled": true,
							"mode": {
							  "cloud": false,
							  "ignoreSPA": true,
							  "light": false,
							  "sample": false,
							  "segment": {
								"end": 100,
								"start": 0
							  },
							  "trigger": "pageload"
							},
							"name": "Custom HTML",
							"neoEvents": [],
							"permissions": [
							  "execute_unsafe_scripts"
							],
							"settings": {},
							"type": "component"
						  }
					},
					"triggers": {
						"Pageview": {
							"clientRules": [],
							"description": "All page loads",
							"excludeRules": [],
							"loadRules": [
							  {
								"match": "{{ client.__zarazTrack }}",
								"op": "EQUALS",
								"value": "Pageview"
							  }
							],
							"name": "Pageview",
							"system": "pageload"
						}
					},
					"variables": {},
					"zarazVersion": 44
			}
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/config", handler)
	payload := ZarazConfig{
		"dataLayer":     true,
		"debugKey":      "cluu08t8nne8gq9g3160",
		"dlp":           []interface{}{},
		"historyChange": true,
		"settings": map[string]interface{}{
			"autoInjectScript": true,
		},
		"tools": map[string]interface{}{
			"PBQr": map[string]interface{}{
				"blockingTriggers": []interface{}{},
				"component":        "html",
				"defaultFields":    map[string]interface{}{},
				"enabled":          true,
				"mode": map[string]interface{}{
					"cloud":     false,
					"ignoreSPA": true,
					"light":     false,
					"sample":    false,
					"segment": map[string]interface{}{
						"end":   float64(100),
						"start": float64(0),
					},
					"trigger": "pageload",
				},
				"name":        "Custom HTML",
				"neoEvents":   []interface{}{},
				"permissions": []interface{}{"execute_unsafe_scripts"},
				"settings":    map[string]interface{}{},
				"type":        "component",
			},
		},
		"triggers": map[string]interface{}{
			"Pageview": map[string]interface{}{
				"clientRules": []interface{}{},
				"description": "All page loads",
				"loadRules": []interface{}{map[string]interface{}{
					"match": "{{ client.__zarazTrack }}",
					"op":    "EQUALS",
					"value": "Pageview",
				}},
				"excludeRules": []interface{}{},
				"name":         "Pageview",
				"system":       "pageload",
			},
		},
		"variables":    map[string]interface{}{},
		"zarazVersion": float64(44),
	}
	expected := ZarazConfigResponse{
		Result: ZarazConfig{
			"dataLayer":     true,
			"debugKey":      "cheese",
			"dlp":           []interface{}{},
			"historyChange": true,
			"settings": map[string]interface{}{
				"autoInjectScript": true,
			},
			"tools": map[string]interface{}{
				"PBQr": map[string]interface{}{
					"blockingTriggers": []interface{}{},
					"component":        "html",
					"defaultFields":    map[string]interface{}{},
					"enabled":          true,
					"mode": map[string]interface{}{
						"cloud":     false,
						"ignoreSPA": true,
						"light":     false,
						"sample":    false,
						"segment": map[string]interface{}{
							"end":   float64(100),
							"start": float64(0),
						},
						"trigger": "pageload",
					},
					"name":        "Custom HTML",
					"neoEvents":   []interface{}{},
					"permissions": []interface{}{"execute_unsafe_scripts"},
					"settings":    map[string]interface{}{},
					"type":        "component",
				},
			},
			"triggers": map[string]interface{}{
				"Pageview": map[string]interface{}{
					"clientRules": []interface{}{},
					"description": "All page loads",
					"loadRules": []interface{}{map[string]interface{}{
						"match": "{{ client.__zarazTrack }}",
						"op":    "EQUALS",
						"value": "Pageview",
					}},
					"excludeRules": []interface{}{},
					"name":         "Pageview",
					"system":       "pageload",
				},
			},
			"variables":    map[string]interface{}{},
			"zarazVersion": float64(44),
		},
		Response: Response{
			Success:  true,
			Messages: []ResponseInfo{},
			Errors:   []ResponseInfo{},
		},
	}

	actual, err := client.UpdateZarazConfig(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetZarazWorkflow(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"errors": [],
			"messages": [],
			"success": true,
			"result": "realtime"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/workflow", handler)
	want := ZarazWorkflowResponse{
		Result: "realtime",
		Response: Response{
			Success:  true,
			Messages: []ResponseInfo{},
			Errors:   []ResponseInfo{},
		},
	}

	actual, err := client.GetZarazWorkflow(context.Background(), ZoneIdentifier(testZoneID))

	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestUpdateZarazWorkflow(t *testing.T) {
	setup()
	defer teardown()

	payload := "realtime"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		body, _ := io.ReadAll(r.Body)
		bodyString := string(body)
		assert.Equal(t, fmt.Sprintf("\"%s\"", payload), bodyString)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"errors": [],
			"messages": [],
			"success": true,
			"result": "realtime"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/workflow", handler)
	want := ZarazWorkflowResponse{
		Result: "realtime",
		Response: Response{
			Success:  true,
			Messages: []ResponseInfo{},
			Errors:   []ResponseInfo{},
		},
	}

	actual, err := client.UpdateZarazWorkflow(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestPublishZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	payload := "test description"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		body, _ := io.ReadAll(r.Body)
		bodyString := string(body)
		assert.Equal(t, fmt.Sprintf("\"%s\"", payload), bodyString)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"errors": [],
			"messages": [],
			"success": true,
			"result": "Config has been published successfully"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/publish", handler)
	want := ZarazPublishResponse{
		Result: "Config has been published successfully",
		Response: Response{
			Success:  true,
			Messages: []ResponseInfo{},
			Errors:   []ResponseInfo{},
		},
	}

	actual, err := client.PublishZarazConfig(context.Background(), ZoneIdentifier(testZoneID), payload)

	require.NoError(t, err)

	assert.Equal(t, want, actual)
}

func TestGetZarazConfigHistory(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"count": 1,
			"data": [
				{
					"createdAt": "2023-12-19T19:24:42.779683Z",
					"description": "Moving to Preview & Publish workflow",
					"id": 1005135,
					"updatedAt": "2023-12-19T19:24:42.779683Z",
					"userId": "9ceddf6f117afe04c64716c83468d3a4"
				}
			]
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/history", handler)

	_, _, err := client.GetZarazConfigHistory(context.Background(), ZoneIdentifier(testZoneID), GetZarazConfigHistoryParams{})
	require.NoError(t, err)
}

func TestGetDefaultZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "text/plain")
		fmt.Fprint(w, `{
			"someTestKeyThatRepsTheConfig": "test"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/default", handler)

	_, err := client.GetDefaultZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)
}

func TestExportZarazConfig(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "text/plain")
		fmt.Fprint(w, `{
			"someTestKeyThatRepsTheConfig": "test"
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/settings/zaraz/v2/export", handler)

	err := client.ExportZarazConfig(context.Background(), ZoneIdentifier(testZoneID))
	require.NoError(t, err)
}
