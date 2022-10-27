//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Dashboard.
func (d Dashboard) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DashboardLens.
func (d DashboardLens) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metadata", d.Metadata)
	populate(objectMap, "order", d.Order)
	populate(objectMap, "parts", d.Parts)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DashboardPartMetadata.
func (d DashboardPartMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["type"] = d.Type
	if d.AdditionalProperties != nil {
		for key, val := range d.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DashboardPartMetadata.
func (d *DashboardPartMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		default:
			if d.AdditionalProperties == nil {
				d.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				d.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DashboardParts.
func (d DashboardParts) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metadata", d.Metadata)
	populate(objectMap, "position", d.Position)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DashboardParts.
func (d *DashboardParts) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "metadata":
			d.Metadata, err = unmarshalDashboardPartMetadataClassification(val)
			delete(rawMsg, key)
		case "position":
			err = unpopulate(val, "Position", &d.Position)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DashboardPartsPosition.
func (d DashboardPartsPosition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "colSpan", d.ColSpan)
	populate(objectMap, "metadata", d.Metadata)
	populate(objectMap, "rowSpan", d.RowSpan)
	populate(objectMap, "x", d.X)
	populate(objectMap, "y", d.Y)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DashboardProperties.
func (d DashboardProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "lenses", d.Lenses)
	populate(objectMap, "metadata", d.Metadata)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MarkdownPartMetadata.
func (m MarkdownPartMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inputs", m.Inputs)
	populate(objectMap, "settings", m.Settings)
	objectMap["type"] = "Extension/HubsExtension/PartType/MarkdownPart"
	if m.AdditionalProperties != nil {
		for key, val := range m.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarkdownPartMetadata.
func (m *MarkdownPartMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inputs":
			err = unpopulate(val, "Inputs", &m.Inputs)
			delete(rawMsg, key)
		case "settings":
			err = unpopulate(val, "Settings", &m.Settings)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		default:
			if m.AdditionalProperties == nil {
				m.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				m.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PatchableDashboard.
func (p PatchableDashboard) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}