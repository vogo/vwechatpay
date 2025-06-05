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

package vwechatpay

import (
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vos"
)

func TestLoadCert(t *testing.T) {
	MerchantID := vos.EnvString("MERCHANT_ID")
	MerchantAPIv3Key := vos.EnvString("MERCHANT_API_V3_KEY")
	MerchantCertSerialNumber := vos.EnvString("MERCHANT_CERT_SERIAL_NUMBER")

	if MerchantID == "" || MerchantAPIv3Key == "" || MerchantCertSerialNumber == "" {
		t.Skip("skip test, env not found")
	}

	cli := NewManagerFromEnv()
	cert := LoadPlatformCert(cli)
	assert.Nil(t, cert)
}

func TestInitConn(t *testing.T) {
	platCertString := `BagBjoI0TAF1PMdfBYNBNgpqh5OnHQrbJ9yVMPiqvGeSMyR0kJO6GawAIhRDgFXemSnuN2QoRGnCh3Zdz54Kv6K+OdjKbwCUZMPSV1/1r5iKZtPrn/Aay/+ZJlqWrIwhr3Jc39ma/hxj2gmB/4gQdWaI/RebSvA/am0aaOuboVG8MhuxMsBdBf53BLYSeHUjEQPwzQOKlW9CFGvf3MbKxbzYN2Jkzj1XhAvUCbIh8bt+LsB4nvGxYdc/mn4tWTVnQ3X88JSTEj126Pck4UkZHCJwHtc6JD16vyK+KD/ribM2MKY58YoAaKcr7CpLwjYgpJbjnetvnIcV28xPZY1aLiJpETCK4nqgSg/HaIDWJJBKFEvgkgAJnwXvBl7BHzh/+W7phT6iGyCs4Jh7ozV0lEPexK9Ytl5b0wUQ5i73its2B/MJEKhXkZBtlv8J22sVasHaYxTA15Z//SjDCi+YKHySbSqif2mjS8LEPHtUwpJ9WyQIx45XxamXcvRjz958O809VzFrGdybaDiM5SeWp0MzI3a9LAZLoJn0XZkjWypBYXvlVsdv44iUTY3qAPmYJtKoxi6LaFkj1uirEPBpfY7Q3pldJvPDfK2OYzgY2YjbjK9epOjBx1F639fuvdEwX0+o3Xmsq0JOBWRGHIOereLCNeJthkdDH5sSQ3BkzKa+Q6X7QJ3WSr+dnKEULC2enzVZ1HCp41fmZF1+gDBl/G53274TWIDzKERWD0hLOEJZlIwQVQSLU/3SmncUPkk57ozZUdXqkRcAJPW7wttsh0Cjl4B22g1OSNUMxUrh9Wz9ejmiNQJFRS0iks8cuajYFF3Ug7uRjiKB+EKms2JR9GaTo+iW5pQ0U7ged+uOLk1x4pN6YnOIVzNH7HE7EMSmHfOD7yZkdMeNYJTDoRVMzHnER9FIX6e607XAXA9Ksi/UcUIMTJWNe+8YzIXY0CSIVtPR3Xh42/NO+gqma95ziOqZaHpncSd0W7E9vNhSi5V1G6EfxFDhXIWdfsUOCsMxWf/gDx8XHW4kSW7iAsbDFdB7gIoJZuxqFvlpZ54k+hBKB5pZ9K63tQvt59aw8WfHFBlcjw6ZIsaHwugVC79mAD86qVoafPA+hCd7zdpyNuJNlB4DyX3zUT3Jw+duK5EJuZTDouLxWTZ4PFmFu69mNc1jymRgE3vcPLQz2Kisu6d+OJxzpk7tqiLk+gYnAA0cgrx80lmcS2nXsaJoErOVURFFlSwW/N8oR4nf6Wrvp3t8bqr/zmcYDJEbzy+isE7gf2Pt004qm2APWHzUXgy6mY8cBY9rztYaXA0YJW/58jNEyfB97v6GO8RJScmGFY1ZPnR79GgzHFlhLVee63k96KEml3m5Bbg265T69sPFpYEC88wxUwufZ6cnc2VtD1ca/iBgNUtJxlax8PuQoE5GTejN224bm+11+abAc1kKIGnvfJrVyjLeqVWLLfQuMzfOlOaogPVzZlx+03M74YjHgncSR/8pjo2n+KHFa/LNFsTH24Qdazv7OBMkk3pT5MpnJP79jZMnPrcRH6GQ+eoXotKIyQsK2UcCbCbX8hj5Fq8J6PFEzv4EufTIy+jJ/4xpoRq4rc/RlmuzsByZv5Lr1czVV2lXla01pLjZPpEgO2yTi348H8DTThftrM6Zzjwy5yFHViOA8QG+7MmLpqh0eoei+pOQPcqMro/jSjDviWBInju2AFouGZCE2/LnyRJhCS/80Xj9JhITSQmn3NxgtIxYVyxnEKPIaZlVXMqAAf8z8SiK+NG4ieV+sUXs+KIGr3M8rPGKdl3D7mU8n80DpzyEIlynKM7DjrCyeIWFeRzNTkdswwVz7jaMT+apEW5QDFFnIumRXB1vVq0SSF2nft3ciCvapGKL/Cvz86cPvxxXeUALpcH67V2n7ipRZ+q44lejY9eKUeTQtetlHvC+shAjyA699pKH0WQwLNRyxIjU6WHLbYurCH55jTpW43M62XoC`
	//bytes, _ := base64.StdEncoding.DecodeString(platCertString)
	//privateKey, _ := utils.LoadPrivateKeyWithPath("/Users/tiltwind/cert/1681252795_20240816_cert/apiclient_key.pem")
	cert, err := utils.LoadCertificate(string(platCertString))
	if err != nil {
		t.Error(err)
	}
	t.Log(cert)
}
