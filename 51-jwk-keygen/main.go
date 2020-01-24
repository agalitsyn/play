/*-
 * Copyright 2017 Square Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base32"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ed25519"
	"gopkg.in/alecthomas/kingpin.v2"

	"gopkg.in/square/go-jose.v2"
)

var (
	app = kingpin.New("jwk-keygen", "A command-line utility to generate public/pirvate keypairs in JWK format.")

	use = app.Flag("use", "Desrired key use").Required().Enum("enc", "sig")
	publicKey = app.Flag("public-key", "Path to key file (PEM or DER-encoded)").ExistingFile()
	privateKey = app.Flag("private-key", "Path to key file (PEM or DER-encoded)").ExistingFile()
	alg = app.Flag("alg", "Generate key to be used for ALG").Required().Enum(
		// `sig`
		string(jose.ES256), string(jose.ES384), string(jose.ES512), string(jose.EdDSA),
		string(jose.RS256), string(jose.RS384), string(jose.RS512), string(jose.PS256), string(jose.PS384), string(jose.PS512),
		// `enc`
		string(jose.RSA1_5), string(jose.RSA_OAEP), string(jose.RSA_OAEP_256),
		string(jose.ECDH_ES), string(jose.ECDH_ES_A128KW), string(jose.ECDH_ES_A192KW), string(jose.ECDH_ES_A256KW),
	)
	bits    = app.Flag("bits", "Key size in bits").Int()
	kid     = app.Flag("kid", "Key ID").String()
	kidRand = app.Flag("kid-rand", "Generate random Key ID").Bool()
)

// KeygenSig generates keypair for corresponding SignatureAlgorithm.
func KeygenSig(alg jose.SignatureAlgorithm, bits int) (crypto.PublicKey, crypto.PrivateKey, error) {
	switch alg {
	case jose.ES256, jose.ES384, jose.ES512, jose.EdDSA:
		keylen := map[jose.SignatureAlgorithm]int{
			jose.ES256: 256,
			jose.ES384: 384,
			jose.ES512: 521, // sic!
			jose.EdDSA: 256,
		}
		if bits != 0 && bits != keylen[alg] {
			return nil, nil, errors.New("this `alg` does not support arbitrary key length")
		}
	case jose.RS256, jose.RS384, jose.RS512, jose.PS256, jose.PS384, jose.PS512:
		if bits == 0 {
			bits = 2048
		}
		if bits < 2048 {
			return nil, nil, errors.New("too short key for RSA `alg`, 2048+ is required")
		}
	}
	switch alg {
	case jose.ES256:
		// The cryptographic operations are implemented using constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		return key.Public(), key, err
	case jose.ES384:
		// NB: The cryptographic operations do not use constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		return key.Public(), key, err
	case jose.ES512:
		// NB: The cryptographic operations do not use constant-time algorithms.
		key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		return key.Public(), key, err
	case jose.EdDSA:
		pub, key, err := ed25519.GenerateKey(rand.Reader)
		return pub, key, err
	case jose.RS256, jose.RS384, jose.RS512, jose.PS256, jose.PS384, jose.PS512:
		key, err := rsa.GenerateKey(rand.Reader, bits)
		//err = savePublicPEMKey(fmt.Sprintf("key_%s.pub", alg), key.PublicKey)
		//err = savePEMKey(fmt.Sprintf("key_%s", alg), key)
		return key.Public(), key, err
	default:
		return nil, nil, errors.New("unknown `alg` for `use` = `sig`")
	}
}

// KeygenEnc generates keypair for corresponding KeyAlgorithm.
func KeygenEnc(alg jose.KeyAlgorithm, bits int) (crypto.PublicKey, crypto.PrivateKey, error) {
	switch alg {
	case jose.RSA1_5, jose.RSA_OAEP, jose.RSA_OAEP_256:
		if bits == 0 {
			bits = 2048
		}
		if bits < 2048 {
			return nil, nil, errors.New("too short key for RSA `alg`, 2048+ is required")
		}
		key, err := rsa.GenerateKey(rand.Reader, bits)
		return key.Public(), key, err
	case jose.ECDH_ES, jose.ECDH_ES_A128KW, jose.ECDH_ES_A192KW, jose.ECDH_ES_A256KW:
		var crv elliptic.Curve
		switch bits {
		case 0, 256:
			crv = elliptic.P256()
		case 384:
			crv = elliptic.P384()
		case 521:
			crv = elliptic.P521()
		default:
			return nil, nil, errors.New("unknown elliptic curve bit length, use one of 256, 384, 521")
		}
		key, err := ecdsa.GenerateKey(crv, rand.Reader)
		return key.Public(), key, err
	default:
		return nil, nil, errors.New("unknown `alg` for `use` = `enc`")
	}
}

func main() {
	app.Version("v2")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *kidRand {
		if *kid == "" {
			b := make([]byte, 5)
			_, err := rand.Read(b)
			app.FatalIfError(err, "can't Read() crypto/rand")
			*kid = base32.StdEncoding.EncodeToString(b)
		} else {
			app.FatalUsage("can't combine --kid and --kid-rand")
		}
	}

	var privKey crypto.PublicKey
	var pubKey crypto.PrivateKey
	if *privateKey != "" && *publicKey != "" {
		key, err := rsaConfigSetup(*privateKey, "", *publicKey)
		app.FatalIfError(err, "")
		pubKey = key.Public()
		privKey = key
	} else {
		var err error
		switch *use {
		case "sig":
			pubKey, privKey, err = KeygenSig(jose.SignatureAlgorithm(*alg), *bits)
		case "enc":
			pubKey, privKey, err = KeygenEnc(jose.KeyAlgorithm(*alg), *bits)
		}
		app.FatalIfError(err, "unable to generate key")
	}

	priv := jose.JSONWebKey{Key: privKey, KeyID: *kid, Algorithm: *alg, Use: *use}
	pub := jose.JSONWebKey{Key: pubKey, KeyID: *kid, Algorithm: *alg, Use: *use}

	if priv.IsPublic() || !pub.IsPublic() || !priv.Valid() || !pub.Valid() {
		app.Fatalf("invalid keys were generated")
	}

	privJS, err := priv.MarshalJSON()
	app.FatalIfError(err, "can't Marshal private key to JSON")
	pubJS, err := pub.MarshalJSON()
	app.FatalIfError(err, "can't Marshal public key to JSON")

	if *kid == "" {
		fmt.Printf("==> jwk_%s.pub <==\n", *alg)
		fmt.Println(string(pubJS))
		fmt.Printf("==> jwk_%s <==\n", *alg)
		fmt.Println(string(privJS))
	} else {
		// JWK Thumbprint (RFC7638) is not used for key id because of
		// lack of canonical representation.
		fname := fmt.Sprintf("jwk_%s_%s_%s", *use, *alg, *kid)
		err = writeNewFile(fname+".pub", pubJS, 0444)
		app.FatalIfError(err, "can't write public key to file %s.pub", fname)
		fmt.Printf("Written public key to %s.pub\n", fname)
		err = writeNewFile(fname, privJS, 0400)
		app.FatalIfError(err, "cant' write private key to file %s", fname)
		fmt.Printf("Written private key to %s\n", fname)
	}
}

// writeNewFile is shameless copy-paste from ioutil.WriteFile with a bit
// different flags for OpenFile.
func writeNewFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}


func savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}
	return nil
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}
	return nil
}

func rsaConfigSetup(rsaPrivateKeyLocation, rsaPrivateKeyPassword, rsaPublicKeyLocation string) (*rsa.PrivateKey, error) {
	if rsaPrivateKeyLocation == "" {
		return nil, fmt.Errorf("no RSA Key given")
	}

	priv, err := ioutil.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		return nil, fmt.Errorf("no RSA private key found")
	}

	privPem, _ := pem.Decode(priv)
	var privPemBytes []byte
	//if privPem.Type != "RSA PRIVATE KEY" {
	//	return nil, fmt.Errorf("RSA private key is of the wrong type: %v", privPem.Type)
	//}

	if rsaPrivateKeyPassword != "" {
		privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(rsaPrivateKeyPassword))
	} else {
		privPemBytes = privPem.Bytes
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
			return nil, fmt.Errorf("unable to parse RSA private key, generating a temp one: %v", err)
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unable to parse RSA private key, generating a temp one: %v", err)
	}

	pub, err := ioutil.ReadFile(rsaPublicKeyLocation)
	if err != nil {
		return nil, fmt.Errorf("no RSA public key found")
	}
	pubPem, _ := pem.Decode(pub)
	//if pubPem == nil {
	//	return nil, fmt.Errorf("could not decode")
	//}
	//if pubPem.Type != "RSA PUBLIC KEY" {
	//	return nil, fmt.Errorf("RSA public key is of the wrong type", pubPem.Type)
	//}

	if parsedKey, err = x509.ParsePKIXPublicKey(pubPem.Bytes); err != nil {
		return nil, fmt.Errorf("could not parse RSA public key")
	}

	var pubKey *rsa.PublicKey
	if pubKey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("could not parse RSA public key")
	}

	privateKey.PublicKey = *pubKey

	return privateKey, nil
}
