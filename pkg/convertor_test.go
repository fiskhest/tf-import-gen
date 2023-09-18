package tfimportgen

import (
	"github.com/kishaningithub/tf-import-gen/pkg/internal/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ComputeTerraformImportForResource(t *testing.T) {
	tests := []struct {
		name              string
		terraformResource parser.TerraformResource
		expected          TerraformImport
	}{
		{
			name: "For aws_iam_role_policy_attachment",
			terraformResource: parser.TerraformResource{
				Address: "aws_iam_role_policy_attachment.test",
				Type:    "aws_iam_role_policy_attachment",
				AttributeValues: map[string]any{
					"role":       "test-role",
					"policy_arn": "test-policy-arn",
				},
			},
			expected: TerraformImport{
				ResourceAddress: "aws_iam_role_policy_attachment.test",
				ResourceID:      "test-role/test-policy-arn",
			},
		},
		{
			name: "For aws_lambda_permission",
			terraformResource: parser.TerraformResource{
				Address: "aws_lambda_permission.test",
				Type:    "aws_lambda_permission",
				AttributeValues: map[string]any{
					"statement_id":  "test-statement-id",
					"function_name": "test-function-name",
				},
			},
			expected: TerraformImport{
				ResourceAddress: "aws_lambda_permission.test",
				ResourceID:      "test-function-name/test-statement-id",
			},
		},
		{
			name: "For aws_security_group_rule with source_security_group_id",
			terraformResource: parser.TerraformResource{
				Address: "aws_security_group_rule.test",
				Type:    "aws_security_group_rule",
				AttributeValues: map[string]any{
					"security_group_id":        "security-group-id",
					"type":                     "type",
					"protocol":                 "protocol",
					"from_port":                1234,
					"to_port":                  5678,
					"source_security_group_id": "source-security-group-id",
				},
			},
			expected: TerraformImport{
				ResourceAddress: "aws_security_group_rule.test",
				ResourceID:      "security-group-id_type_protocol_1234_5678_source-security-group-id",
			},
		},
		{
			name: "For aws_security_group_rule with cidr_blocks",
			terraformResource: parser.TerraformResource{
				Address: "aws_security_group_rule.test",
				Type:    "aws_security_group_rule",
				AttributeValues: map[string]any{
					"security_group_id": "security-group-id",
					"type":              "type",
					"protocol":          "protocol",
					"from_port":         1234,
					"to_port":           5678,
					"cidr_blocks":       []any{"cidr-block-1", "cidr-block-2"},
				},
			},
			expected: TerraformImport{
				ResourceAddress: "aws_security_group_rule.test",
				ResourceID:      "security-group-id_type_protocol_1234_5678_cidr-block-1_cidr-block-2",
			},
		},
		{
			name: "For everything else",
			terraformResource: parser.TerraformResource{
				Address: "example.address",
				Type:    "example_type",
				AttributeValues: map[string]any{
					"id": "test_id",
				},
			},
			expected: TerraformImport{
				ResourceAddress: "example.address",
				ResourceID:      "test_id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := computeTerraformImportForResource(tt.terraformResource)
			require.Equal(t, tt.expected, actual)
		})
	}
}