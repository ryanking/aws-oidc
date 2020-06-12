package aws_config_client

import (
	"fmt"
	"testing"

	server "github.com/chanzuckerberg/aws-oidc/pkg/aws_config_server"
	"github.com/stretchr/testify/require"
)

// func TestLoop(t *testing.T) {
// 	r := require.New(t)

// 	// note how: "Account Name With Spaces" => "account-name-with-spaces"
// 	expected := `[profile account-name-with-spaces]
// output             = json
// credential_process = sh -c 'aws-oidc creds-process --issuer-url=issuer-url --client-id=foo_client_id --aws-role-arn=test1RoleName 2> /dev/tty'
// region             = us-west-2

// [profile my-second-new-profile]
// output             = json
// credential_process = sh -c 'aws-oidc creds-process --issuer-url=issuer-url --client-id=bar_client_id --aws-role-arn=test1RoleName 2> /dev/tty'
// region             = us-west-2

// [profile test1]
// output             = json
// credential_process = sh -c 'aws-oidc creds-process --issuer-url=issuer-url --client-id=bar_client_id --aws-role-arn=test2RoleName 2> /dev/tty'
// region             = us-west-2

// `

// 	out := ini.Empty()
// 	prompt := &MockPrompt{

// 		selectResponse: []int{
// 			0, 0, // select the first role in the first account
// 			1, 0, // select the first role in the second account
// 			1, 1, // select the second role in the second account
// 		},
// 		inputResponse: []string{
// 			"",                              // aws region
// 			"", "my-second-new-profile", "", // aws profile names
// 		},
// 		confirmResponse: []bool{true, true, false},
// 	}

// 	c := NewCompleter(prompt, generateDummyData())

// 	err := c.Loop(out)
// 	r.NoError(err)

// 	generatedConfig := bytes.NewBuffer(nil)
// 	_, err = out.WriteTo(generatedConfig)
// 	r.NoError(err)
// 	r.Equal(expected, generatedConfig.String())
// }

func TestAWSProfileNameValidator(t *testing.T) {
	type test struct {
		input interface{}
		err   error
	}
	r := require.New(t)

	tests := []test{
		{input: 1, err: fmt.Errorf("input not a string")},
		{input: "not valid", err: fmt.Errorf("Input (not valid) not a valid AWS profile name")},
		{input: "valid", err: nil},
	}

	c := NewCompleter(nil, generateDummyData())
	for _, test := range tests {
		err := c.awsProfileNameValidator(test.input)
		if test.err == nil {
			r.NoError(err)
		} else {
			r.Error(err)
			r.Equal(test.err.Error(), err.Error())
		}

	}
}

func generateDummyData() *server.AWSConfig {
	return &server.AWSConfig{
		Profiles: []server.AWSProfile{
			{
				ClientID: "bar_client_id",
				AWSAccount: server.AWSAccount{
					Name: "test1",
					ID:   "test_id_1",
				},
				RoleARN:   "test1RoleName",
				IssuerURL: "issuer-url",
			},
			{
				ClientID: "bar_client_id",
				AWSAccount: server.AWSAccount{
					Name: "test1",
					ID:   "test_id_1",
				},
				RoleARN:   "test2RoleName",
				IssuerURL: "issuer-url",
			},
			{
				ClientID: "foo_client_id",
				AWSAccount: server.AWSAccount{
					Name: "Account Name With Spaces",
					ID:   "account id 2",
				},
				RoleARN:   "test1RoleName",
				IssuerURL: "issuer-url",
			},
			{
				ClientID: "foo_client_id",
				AWSAccount: server.AWSAccount{
					Name: "Account Name With Spaces",
					ID:   "account id 2",
				},
				RoleARN:   "test1RoleName",
				IssuerURL: "issuer-url",
			},
		},
	}
}
