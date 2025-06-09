/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vwxutils

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// GetCertificateSerialNumber 获取证书序列号
func GetCertificateSerialNumber(cert *x509.Certificate) string {
	return utils.GetCertificateSerialNumber(*cert)
}

// EncryptRSA 使用RSA公钥加密
func EncryptRSA(plaintext string, publicKey *rsa.PublicKey) (string, error) {
	ciphertext, err := utils.EncryptOAEPWithPublicKey(plaintext, publicKey)
	if err != nil {
		return "", fmt.Errorf("encrypt error: %w", err)
	}

	return ciphertext, nil
}

// DecryptRSA 使用RSA私钥解密
func DecryptRSA(ciphertext string, privateKey *rsa.PrivateKey) (string, error) {
	plaintext, err := utils.DecryptOAEP(ciphertext, privateKey)
	if err != nil {
		return "", fmt.Errorf("decrypt error: %w", err)
	}

	return plaintext, nil
}

// EncryptAESGCM 使用AES-GCM加密
func EncryptAESGCM(plaintext, key string) (string, error) {
	keyBytes := []byte(key)
	plaintextBytes := []byte(plaintext)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("new cipher error: %w", err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generate nonce error: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("new GCM error: %w", err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintextBytes, nil)
	encoded := base64.StdEncoding.EncodeToString(append(nonce, ciphertext...))

	return encoded, nil
}

// DecryptAESGCM 使用AES-GCM解密
func DecryptAESGCM(ciphertext, key string) (string, error) {
	keyBytes := []byte(key)

	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode base64 error: %w", err)
	}

	if len(decoded) < 12 {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce := decoded[:12]
	ciphertextBytes := decoded[12:]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("new cipher error: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("new GCM error: %w", err)
	}

	plaintextBytes, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt error: %w", err)
	}

	return string(plaintextBytes), nil
}

// GenerateNonceStr 生成随机字符串
func GenerateNonceStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// GenerateTimestamp 生成时间戳
func GenerateTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// SHA256WithRSA 使用SHA256WithRSA算法签名
func SHA256WithRSA(message string, privateKey *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("sign error: %w", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySHA256WithRSA 验证SHA256WithRSA签名
func VerifySHA256WithRSA(message, signature string, publicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256([]byte(message))
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("decode signature error: %w", err)
	}

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signatureBytes)
}

// FileBase64 文件转base64
// bash:  base64 -w 0 -i file.txt
func FileBase64(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
