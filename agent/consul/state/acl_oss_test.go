// +build !consulent

package state

import "github.com/hashicorp/consul/agent/structs"

func testIndexerTableACLPolicies() map[string]indexerTestCase {
	obj := &structs.ACLPolicy{
		ID:   "123e4567-e89b-12d3-a456-426614174abc",
		Name: "PoLiCyNaMe",
	}
	encodedID := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9b, 0x12, 0xd3, 0xa4, 0x56, 0x42, 0x66, 0x14, 0x17, 0x4a, 0xbc}
	return map[string]indexerTestCase{
		indexID: {
			read: indexValue{
				source:   obj.ID,
				expected: encodedID,
			},
			write: indexValue{
				source:   obj,
				expected: encodedID,
			},
		},
		indexName: {
			read: indexValue{
				source:   Query{Value: "PolicyName"},
				expected: []byte("policyname\x00"),
			},
			write: indexValue{
				source:   obj,
				expected: []byte("policyname\x00"),
			},
		},
	}
}

func testIndexerTableACLRoles() map[string]indexerTestCase {
	policyID1 := "123e4567-e89a-12d7-a456-426614174001"
	policyID2 := "123e4567-e89a-12d7-a456-426614174002"
	obj := &structs.ACLRole{
		ID:   "123e4567-e89a-12d7-a456-426614174abc",
		Name: "RoLe",
		Policies: []structs.ACLRolePolicyLink{
			{ID: policyID1}, {ID: policyID2},
		},
	}
	encodedID := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9a, 0x12, 0xd7, 0xa4, 0x56, 0x42, 0x66, 0x14, 0x17, 0x4a, 0xbc}
	encodedPID1 := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9a, 0x12, 0xd7, 0xa4, 0x56, 0x42, 0x66, 0x14, 0x17, 0x40, 0x01}
	encodedPID2 := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9a, 0x12, 0xd7, 0xa4, 0x56, 0x42, 0x66, 0x14, 0x17, 0x40, 0x02}
	return map[string]indexerTestCase{
		indexID: {
			read: indexValue{
				source:   obj.ID,
				expected: encodedID,
			},
			write: indexValue{
				source:   obj,
				expected: encodedID,
			},
		},
		indexName: {
			read: indexValue{
				source:   Query{Value: "RoLe"},
				expected: []byte("role\x00"),
			},
			write: indexValue{
				source:   obj,
				expected: []byte("role\x00"),
			},
		},
		indexPolicies: {
			read: indexValue{
				source:   Query{Value: policyID1},
				expected: encodedPID1,
			},
			writeMulti: indexValueMulti{
				source:   obj,
				expected: [][]byte{encodedPID1, encodedPID2},
			},
		},
	}
}
