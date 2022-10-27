//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlocks

const (
	moduleName    = "armlocks"
	moduleVersion = "v1.0.0"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// LockLevel - The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized
// users are able to read and modify the resources, but not delete. ReadOnly means
// authorized users can only read from a resource, but they can't modify or delete it.
type LockLevel string

const (
	LockLevelCanNotDelete LockLevel = "CanNotDelete"
	LockLevelNotSpecified LockLevel = "NotSpecified"
	LockLevelReadOnly     LockLevel = "ReadOnly"
)

// PossibleLockLevelValues returns the possible values for the LockLevel const type.
func PossibleLockLevelValues() []LockLevel {
	return []LockLevel{
		LockLevelCanNotDelete,
		LockLevelNotSpecified,
		LockLevelReadOnly,
	}
}