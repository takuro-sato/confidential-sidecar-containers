package endpointcontroller

import (
	"encoding/base64"

	"github.com/Microsoft/confidential-sidecar-containers/pkg/attest"
	"github.com/Microsoft/confidential-sidecar-containers/pkg/common"
)

type EndpointController struct {
	CertState      attest.CertState
	Identity       common.Identity
	UvmInformation common.UvmInformation
}

// Status of response
const (
	StatusOK           = iota
	StatusInvalidInput = iota
	StatusForbidden    = iota
)

type MAAAttestInput struct {
	// MAA endpoint which authors the MAA token
	MAAEndpoint string `json:"maa_endpoint" binding:"required"`
	// Base64 encoded representation of runtime data to be encoded
	// as runtime claim in the MAA token
	RuntimeData string `json:"runtime_data" binding:"required"`
}

type MAAAttestOutput struct {
	Token string `json:"token" binding:"required"`
}

func (c *EndpointController) MAAAttest(input MAAAttestInput) (_ MAAAttestOutput, status uint, _ error) {

	// base64 decode the incoming runtime data
	runtimeDataBytes, err := base64.StdEncoding.DecodeString(input.RuntimeData)
	if err != nil {
		return MAAAttestOutput{}, StatusInvalidInput, err
	}

	maa := attest.MAA{
		Endpoint:   input.MAAEndpoint,
		TEEType:    "SevSnpVM",
		APIVersion: "api-version=2020-10-01",
	}

	maaToken, err := c.CertState.Attest(maa, runtimeDataBytes, c.UvmInformation)
	if err != nil {
		return MAAAttestOutput{}, StatusForbidden, err
	}

	return MAAAttestOutput{Token: maaToken}, StatusOK, nil
}
