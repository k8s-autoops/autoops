package autoops

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateServerCert(t *testing.T) {
	cert, certPEM, key, keyPEM, err := GenerateRootCA()
	require.NoError(t, err)
	t.Log(cert.Subject)
	t.Log(certPEM)
	t.Log(key)
	t.Log(keyPEM)

	cert, certPEM, key, keyPEM, err = GenerateServerCert("test autoops", certPEM, keyPEM)
	require.NoError(t, err)
	t.Log(cert.Subject)
	t.Log(certPEM)
	t.Log(key)
	t.Log(keyPEM)
}
