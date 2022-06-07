package oci

// Credentials defines OCI credentials
type Credentials struct {
	TenancyID              string
	UserID                 string
	Region                 string
	CredentialsKey         string
	CredentialsFingerprint string
	CredentialsPassphrase  string
}