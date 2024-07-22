package fronted

var DefaultTrustedCAs = []*CA{
	{
		CommonName: "Amazon Root CA 1",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDQTCCAimgAwIBAgITBmyfz5m/jAo54vB4ikPmljZbyjANBgkqhkiG9w0BAQsF\nADA5MQswCQYDVQQGEwJVUzEPMA0GA1UEChMGQW1hem9uMRkwFwYDVQQDExBBbWF6\nb24gUm9vdCBDQSAxMB4XDTE1MDUyNjAwMDAwMFoXDTM4MDExNzAwMDAwMFowOTEL\nMAkGA1UEBhMCVVMxDzANBgNVBAoTBkFtYXpvbjEZMBcGA1UEAxMQQW1hem9uIFJv\nb3QgQ0EgMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALJ4gHHKeNXj\nca9HgFB0fW7Y14h29Jlo91ghYPl0hAEvrAIthtOgQ3pOsqTQNroBvo3bSMgHFzZM\n9O6II8c+6zf1tRn4SWiw3te5djgdYZ6k/oI2peVKVuRF4fn9tBb6dNqcmzU5L/qw\nIFAGbHrQgLKm+a/sRxmPUDgH3KKHOVj4utWp+UhnMJbulHheb4mjUcAwhmahRWa6\nVOujw5H5SNz/0egwLX0tdHA114gk957EWW67c4cX8jJGKLhD+rcdqsq08p8kDi1L\n93FcXmn/6pUCyziKrlA4b9v7LWIbxcceVOF34GfID5yHI9Y/QCB/IIDEgEw+OyQm\njgSubJrIqg0CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMC\nAYYwHQYDVR0OBBYEFIQYzIU07LwMlJQuCFmcx7IQTgoIMA0GCSqGSIb3DQEBCwUA\nA4IBAQCY8jdaQZChGsV2USggNiMOruYou6r4lK5IpDB/G/wkjUu0yKGX9rbxenDI\nU5PMCCjjmCXPI6T53iHTfIUJrU6adTrCC2qJeHZERxhlbI1Bjjt/msv0tadQ1wUs\nN+gDS63pYaACbvXy8MWy7Vu33PqUXHeeE6V/Uq2V8viTO96LXFvKWlJbYK8U90vv\no/ufQJVtMVT8QtPHRh8jrdkPSHCa2XV4cdFyQzR1bldZwgJcJmApzyMZFo6IQ6XU\n5MsI+yMRQ+hDKXJioaldXgjUkK642M4UwtBV8ob2xJNDd2ZhwLnoQdeXeGADbkpy\nrqXRfboQnoZsG4q5WTP468SQvvG5\n-----END CERTIFICATE-----\n",
	},
	{
		CommonName: "DigiCert Global Root G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH\nMjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI\n2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx\n1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ\nq2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz\ntCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ\nvIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP\nBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV\n5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY\n1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4\nNeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG\nFdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91\n8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe\npLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl\nMrY=\n-----END CERTIFICATE-----\n",
	},
	{
		CommonName: "DigiCert Global Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\nQTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB\nCSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97\nnh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt\n43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P\nT19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4\ngdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO\nBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR\nTLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw\nDQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr\nhMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg\n06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF\nPnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls\nYSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk\nCAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=\n-----END CERTIFICATE-----\n",
	},
	{
		CommonName: "Go Daddy Root Certificate Authority - G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIBADANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMCVVMx\nEDAOBgNVBAgTB0FyaXpvbmExEzARBgNVBAcTClNjb3R0c2RhbGUxGjAYBgNVBAoT\nEUdvRGFkZHkuY29tLCBJbmMuMTEwLwYDVQQDEyhHbyBEYWRkeSBSb290IENlcnRp\nZmljYXRlIEF1dGhvcml0eSAtIEcyMB4XDTA5MDkwMTAwMDAwMFoXDTM3MTIzMTIz\nNTk1OVowgYMxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdBcml6b25hMRMwEQYDVQQH\nEwpTY290dHNkYWxlMRowGAYDVQQKExFHb0RhZGR5LmNvbSwgSW5jLjExMC8GA1UE\nAxMoR28gRGFkZHkgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgLSBHMjCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL9xYgjx+lk09xvJGKP3gElY6SKD\nE6bFIEMBO4Tx5oVJnyfq9oQbTqC023CYxzIBsQU+B07u9PpPL1kwIuerGVZr4oAH\n/PMWdYA5UXvl+TW2dE6pjYIT5LY/qQOD+qK+ihVqf94Lw7YZFAXK6sOoBJQ7Rnwy\nDfMAZiLIjWltNowRGLfTshxgtDj6AozO091GB94KPutdfMh8+7ArU6SSYmlRJQVh\nGkSBjCypQ5Yj36w6gZoOKcUcqeldHraenjAKOc7xiID7S13MMuyFYkMlNAJWJwGR\ntDtwKj9useiciAF9n9T521NtYJ2/LOdYq7hfRvzOxBsDPAnrSTFcaUaz4EcCAwEA\nAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYE\nFDqahQcQZyi27/a9BUFuIMGU2g/eMA0GCSqGSIb3DQEBCwUAA4IBAQCZ21151fmX\nWWcDYfF+OwYxdS2hII5PZYe096acvNjpL9DbWu7PdIxztDhC2gV7+AJ1uP2lsdeu\n9tfeE8tTEH6KRtGX+rcuKxGrkLAngPnon1rpN5+r5N9ss4UXnT3ZJE95kTXWXwTr\ngIOrmgIttRD02JDHBHNA7XIloKmf7J6raBKZV8aPEjoJpL1E/QYVN8Gb5DKj7Tjo\n2GTzLH4U/ALqn83/B2gX2yKQOC16jdFU8WnjXzPKej17CuPKf1855eJ1usV2GDPO\nLPAvTK33sefOT6jEm0pUBsV/fdUID+Ic/n4XuKxe9tQWskMJDE32p2u0mYRlynqI\n4uJEvlz36hz1\n-----END CERTIFICATE-----\n",
	},
	{
		CommonName: "USERTrust RSA Certification Authority",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIF3jCCA8agAwIBAgIQAf1tMPyjylGoG7xkDjUDLTANBgkqhkiG9w0BAQwFADCB\niDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0pl\ncnNleSBDaXR5MR4wHAYDVQQKExVUaGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNV\nBAMTJVVTRVJUcnVzdCBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTAw\nMjAxMDAwMDAwWhcNMzgwMTE4MjM1OTU5WjCBiDELMAkGA1UEBhMCVVMxEzARBgNV\nBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0plcnNleSBDaXR5MR4wHAYDVQQKExVU\naGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNVBAMTJVVTRVJUcnVzdCBSU0EgQ2Vy\ndGlmaWNhdGlvbiBBdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK\nAoICAQCAEmUXNg7D2wiz0KxXDXbtzSfTTK1Qg2HiqiBNCS1kCdzOiZ/MPans9s/B\n3PHTsdZ7NygRK0faOca8Ohm0X6a9fZ2jY0K2dvKpOyuR+OJv0OwWIJAJPuLodMkY\ntJHUYmTbf6MG8YgYapAiPLz+E/CHFHv25B+O1ORRxhFnRghRy4YUVD+8M/5+bJz/\nFp0YvVGONaanZshyZ9shZrHUm3gDwFA66Mzw3LyeTP6vBZY1H1dat//O+T23LLb2\nVN3I5xI6Ta5MirdcmrS3ID3KfyI0rn47aGYBROcBTkZTmzNg95S+UzeQc0PzMsNT\n79uq/nROacdrjGCT3sTHDN/hMq7MkztReJVni+49Vv4M0GkPGw/zJSZrM233bkf6\nc0Plfg6lZrEpfDKEY1WJxA3Bk1QwGROs0303p+tdOmw1XNtB1xLaqUkL39iAigmT\nYo61Zs8liM2EuLE/pDkP2QKe6xJMlXzzawWpXhaDzLhn4ugTncxbgtNMs+1b/97l\nc6wjOy0AvzVVdAlJ2ElYGn+SNuZRkg7zJn0cTRe8yexDJtC/QV9AqURE9JnnV4ee\nUB9XVKg+/XRjL7FQZQnmWEIuQxpMtPAlR1n6BB6T1CZGSlCBst6+eLf8ZxXhyVeE\nHg9j1uliutZfVS7qXMYoCAQlObgOK6nyTJccBz8NUvXt7y+CDwIDAQABo0IwQDAd\nBgNVHQ4EFgQUU3m/WqorSs9UgOHYm8Cd8rIDZsswDgYDVR0PAQH/BAQDAgEGMA8G\nA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQEMBQADggIBAFzUfA3P9wF9QZllDHPF\nUp/L+M+ZBn8b2kMVn54CVVeWFPFSPCeHlCjtHzoBN6J2/FNQwISbxmtOuowhT6KO\nVWKR82kV2LyI48SqC/3vqOlLVSoGIG1VeCkZ7l8wXEskEVX/JJpuXior7gtNn3/3\nATiUFJVDBwn7YKnuHKsSjKCaXqeYalltiz8I+8jRRa8YFWSQEg9zKC7F4iRO/Fjs\n8PRF/iKz6y+O0tlFYQXBl2+odnKPi4w2r78NBc5xjeambx9spnFixdjQg3IM8WcR\niQycE0xyNN+81XHfqnHd4blsjDwSXWXavVcStkNr/+XeTWYRUc+ZruwXtuhxkYze\nSf7dNXGiFSeUHM9h4ya7b6NnJSFd5t0dCy5oGzuCr+yDZ4XUmFF0sbmZgIn/f3gZ\nXHlKYC6SQK5MNyosycdiyA5d9zZbyuAlJQG03RoHnHcAP9Dc1ew91Pq7P8yF1m9/\nqS3fuQL39ZeatTXaw2ewh0qpKJ4jjv9cJ2vhsE/zB+4ALtRZh8tSQZXq9EfX7mRB\nVXyNWQKV3WKdwrnuWih0hKWbt5DHDAff9Yk2dDLWKMGwsAvgnEzDHNb842m1R0aB\nL6KCq9NjRHDEjf8tM7qtj3u1cIiuPhnPQCjY/MiQu12ZIvVS5ljFH4gxQ+6IHdfG\njjxDah2nGN59PRbxYvnKkKj9\n-----END CERTIFICATE-----\n",
	},
}

var DefaultCloudfrontMasquerades = []*Masquerade{
	{
		Domain:    "www.amazon.ae",
		IpAddress: "3.164.6.125",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.27",
	},
	{
		Domain:    "assetserv.com",
		IpAddress: "13.224.6.236",
	},
	{
		Domain:    "datad0g.com",
		IpAddress: "204.246.177.89",
	},
	{
		Domain:    "t.mail.optimumemail1.com",
		IpAddress: "13.224.7.64",
	},
	{
		Domain:    "realisticgames.co.uk",
		IpAddress: "13.224.5.3",
	},
	{
		Domain:    "geocomply.com",
		IpAddress: "204.246.177.39",
	},
	{
		Domain:    "dev.twitch.tv",
		IpAddress: "143.204.1.60",
	},
	{
		Domain:    "dolphin-fe.amazon.com",
		IpAddress: "204.246.178.22",
	},
	{
		Domain:    "brightcove.com",
		IpAddress: "143.204.1.30",
	},
	{
		Domain:    "assets.bwbx.io",
		IpAddress: "204.246.178.184",
	},
	{
		Domain:    "www.dcm-icwweb-dev.com",
		IpAddress: "204.246.178.19",
	},
	{
		Domain:    "www.adbephotos-stage.com",
		IpAddress: "99.86.2.9",
	},
	{
		Domain:    "www.connectwisedev.com",
		IpAddress: "99.86.0.153",
	},
	{
		Domain:    "www.uat.catchplay.com",
		IpAddress: "99.86.0.178",
	},
	{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "205.251.212.47",
	},
	{
		Domain:    "cdn.admin.staging.checkmatenext.com",
		IpAddress: "99.86.3.179",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.227",
	},
	{
		Domain:    "www.amazon.ae",
		IpAddress: "99.84.2.180",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.168",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "99.84.2.162",
	},
	{
		Domain:    "rheemcert.com",
		IpAddress: "99.86.4.107",
	},
	{
		Domain:    "samsungqbe.com",
		IpAddress: "99.84.0.159",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.32",
	},
	{
		Domain:    "api.smartpass.auone.jp",
		IpAddress: "54.239.130.234",
	},
	{
		Domain:    "undercovertourist.com",
		IpAddress: "13.249.7.66",
	},
	{
		Domain:    "emergency.wa.gov.au",
		IpAddress: "13.249.6.72",
	},
	{
		Domain:    "cdn.hands.net",
		IpAddress: "99.86.6.135",
	},
	{
		Domain:    "sellercentral.amazon.com",
		IpAddress: "143.204.2.163",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.36",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.25",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.26",
	},
	{
		Domain:    "www.samsungsmartcam.com",
		IpAddress: "13.249.6.29",
	},
	{
		Domain:    "phdvasia.com",
		IpAddress: "52.222.129.149",
	},
	{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "52.222.134.184",
	},
	{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "52.222.128.161",
	},
	{
		Domain:    "soccerladuma.co.za",
		IpAddress: "52.222.129.150",
	},
	{
		Domain:    "shopch.jp",
		IpAddress: "13.35.3.101",
	},
	{
		Domain:    "www.animelo.jp",
		IpAddress: "143.204.1.59",
	},
	{
		Domain:    "www.netdespatch.com",
		IpAddress: "13.35.4.228",
	},
	{
		Domain:    "trusteerqa.com",
		IpAddress: "13.35.4.21",
	},
	{
		Domain:    "gluon-cv.mxnet.io",
		IpAddress: "13.35.1.114",
	},
	{
		Domain:    "www.smentertainment.com",
		IpAddress: "204.246.177.25",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.166",
	},
	{
		Domain:    "www.api.brightcove.com",
		IpAddress: "204.246.164.127",
	},
	{
		Domain:    "bd1.awsstatic.com",
		IpAddress: "204.246.164.185",
	},
	{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "99.86.0.222",
	},
	{
		Domain:    "as0.awsstatic.com",
		IpAddress: "204.246.178.33",
	},
	{
		Domain:    "isao.net",
		IpAddress: "54.239.192.166",
	},
	{
		Domain:    "rca-upload-cloudstation-us-east-2.qa.hydra.sophos.com",
		IpAddress: "143.204.7.42",
	},
	{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "52.222.132.11",
	},
	{
		Domain:    "unrealengine.com",
		IpAddress: "99.86.4.156",
	},
	{
		Domain:    "pimg.jp",
		IpAddress: "99.86.2.27",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.178",
	},
	{
		Domain:    "webspectator.com",
		IpAddress: "204.246.169.132",
	},
	{
		Domain:    "kaltura.com",
		IpAddress: "99.86.5.174",
	},
	{
		Domain:    "jwo.amazon.com",
		IpAddress: "143.204.1.26",
	},
	{
		Domain:    "z-na.amazon-adsystem.com",
		IpAddress: "54.182.6.41",
	},
	{
		Domain:    "kucoin.com",
		IpAddress: "99.84.7.32",
	},
	{
		Domain:    "supplychainconnect.amazon.com",
		IpAddress: "204.246.169.15",
	},
	{
		Domain:    "product-downloads.atlassian.com",
		IpAddress: "54.182.6.185",
	},
	{
		Domain:    "smtown.com",
		IpAddress: "13.224.5.137",
	},
	{
		Domain:    "www.vistarmedia.com",
		IpAddress: "204.246.169.96",
	},
	{
		Domain:    "www.wowma.jp",
		IpAddress: "99.86.1.78",
	},
	{
		Domain:    "wpcp.shiseido.co.jp",
		IpAddress: "216.137.39.7",
	},
	{
		Domain:    "www.fastretailing.com",
		IpAddress: "99.86.2.108",
	},
	{
		Domain:    "altium.com",
		IpAddress: "99.86.0.186",
	},
	{
		Domain:    "eprocurement.marketplace.us-east-1.amazonaws.com",
		IpAddress: "99.86.2.105",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.132",
	},
	{
		Domain:    "www7.amazon.com",
		IpAddress: "13.249.2.67",
	},
	{
		Domain:    "i-parcel.com",
		IpAddress: "13.224.6.188",
	},
	{
		Domain:    "searchandexplore.com",
		IpAddress: "13.224.6.168",
	},
	{
		Domain:    "oih-gamma-eu.aka.amazon.com",
		IpAddress: "13.224.6.108",
	},
	{
		Domain:    "slackfrontiers.com",
		IpAddress: "99.84.6.196",
	},
	{
		Domain:    "panda.chtbl.com",
		IpAddress: "52.222.135.66",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-west-1.qa.hydra.sophos.com",
		IpAddress: "52.222.134.162",
	},
	{
		Domain:    "we-stats.com",
		IpAddress: "54.182.6.113",
	},
	{
		Domain:    "www.c.ooyala.com",
		IpAddress: "13.249.5.199",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.224.6.97",
	},
	{
		Domain:    "mobizen.com",
		IpAddress: "99.86.2.50",
	},
	{
		Domain:    "www.gdl.imtxwy.com",
		IpAddress: "99.86.5.155",
	},
	{
		Domain:    "d-hrp.com",
		IpAddress: "13.249.2.238",
	},
	{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "13.35.6.59",
	},
	{
		Domain:    "images-cn.ssl-images-amazon.com",
		IpAddress: "52.222.132.50",
	},
	{
		Domain:    "cf.test.frontier.a2z.com",
		IpAddress: "52.222.132.117",
	},
	{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "54.182.2.161",
	},
	{
		Domain:    "mheducation.com",
		IpAddress: "99.84.2.133",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.197",
	},
	{
		Domain:    "www.awscfdns.com",
		IpAddress: "13.35.2.242",
	},
	{
		Domain:    "www.diageohorizon.com",
		IpAddress: "99.84.2.203",
	},
	{
		Domain:    "fujifilmimagine.com",
		IpAddress: "99.86.0.19",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.94",
	},
	{
		Domain:    "www.amazon.com",
		IpAddress: "23.199.14.80",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.107",
	},
	{
		Domain:    "www.53.localytics.com",
		IpAddress: "13.249.2.62",
	},
	{
		Domain:    "load-test6.eu-west-2.cf-embed.net",
		IpAddress: "54.239.130.82",
	},
	{
		Domain:    "www.nosto.com",
		IpAddress: "13.249.6.10",
	},
	{
		Domain:    "oasiscdn.com",
		IpAddress: "204.246.177.64",
	},
	{
		Domain:    "gaijinent.com",
		IpAddress: "99.86.0.103",
	},
	{
		Domain:    "media.preziusercontent.com",
		IpAddress: "13.224.6.207",
	},
	{
		Domain:    "buildinglink.com",
		IpAddress: "143.204.5.63",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.219",
	},
	{
		Domain:    "www.observian.com",
		IpAddress: "52.222.134.2",
	},
	{
		Domain:    "www.amazon.sa",
		IpAddress: "54.182.2.233",
	},
	{
		Domain:    "ebookstore.sony.jp",
		IpAddress: "216.137.39.53",
	},
	{
		Domain:    "amazon.com.au",
		IpAddress: "99.84.2.84",
	},
	{
		Domain:    "alexa-comms-mobile-service.amazon.com",
		IpAddress: "108.139.184.238",
	},
	{
		Domain:    "hkcp08.com",
		IpAddress: "99.86.1.88",
	},
	{
		Domain:    "api.area-hinan-test.au.com",
		IpAddress: "204.246.169.51",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-central-1.qa.hydra.sophos.com",
		IpAddress: "13.35.6.175",
	},
	{
		Domain:    "mapbox.cn",
		IpAddress: "13.35.4.151",
	},
	{
		Domain:    "www.indigoag.tech",
		IpAddress: "204.246.164.179",
	},
	{
		Domain:    "www.production.scrabble.withbuddies.com",
		IpAddress: "99.86.3.120",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.23",
	},
	{
		Domain:    "www.bcovlive.io",
		IpAddress: "99.86.1.216",
	},
	{
		Domain:    "media.aircorsica.com",
		IpAddress: "13.35.1.153",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.129",
	},
	{
		Domain:    "www.dev.aws.casualty.cccis.com",
		IpAddress: "13.35.5.145",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.132",
	},
	{
		Domain:    "www.tigocloud.net",
		IpAddress: "52.222.130.214",
	},
	{
		Domain:    "tenki.auone.jp",
		IpAddress: "13.35.5.68",
	},
	{
		Domain:    "www.thinknearhub.com",
		IpAddress: "13.35.6.112",
	},
	{
		Domain:    "unrulymedia.com",
		IpAddress: "54.182.0.208",
	},
	{
		Domain:    "craftsy.com",
		IpAddress: "54.182.2.209",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.213",
	},
	{
		Domain:    "collectivehealth.com",
		IpAddress: "52.222.131.128",
	},
	{
		Domain:    "www.execute-api.us-east-1.amazonaws.com",
		IpAddress: "99.86.6.34",
	},
	{
		Domain:    "ext-test.app-cloud.jp",
		IpAddress: "143.204.6.27",
	},
	{
		Domain:    "aloseguro.com",
		IpAddress: "52.222.129.160",
	},
	{
		Domain:    "mpago.la",
		IpAddress: "143.204.2.90",
	},
	{
		Domain:    "api.stg.smartpass.auone.jp",
		IpAddress: "99.86.1.57",
	},
	{
		Domain:    "thetvdb.com",
		IpAddress: "54.182.3.230",
	},
	{
		Domain:    "esd.sentinelcloud.com",
		IpAddress: "204.246.178.124",
	},
	{
		Domain:    "www.project-a.videoprojects.net",
		IpAddress: "54.182.0.149",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.5.67",
	},
	{
		Domain:    "deploygate.com",
		IpAddress: "13.35.6.226",
	},
	{
		Domain:    "www.freshdesk.com",
		IpAddress: "99.86.5.24",
	},
	{
		Domain:    "ewrzfr.com",
		IpAddress: "143.204.0.165",
	},
	{
		Domain:    "www.stg.forecast.elyza.ai",
		IpAddress: "143.204.6.5",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.102",
	},
	{
		Domain:    "www.audible.com.au",
		IpAddress: "99.86.0.213",
	},
	{
		Domain:    "tripkit-test4.jeppesen.com",
		IpAddress: "13.249.6.140",
	},
	{
		Domain:    "cdn-legacy.contentful.com",
		IpAddress: "54.182.3.24",
	},
	{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "99.84.0.114",
	},
	{
		Domain:    "poptropica.com",
		IpAddress: "54.182.2.44",
	},
	{
		Domain:    "as0.awsstatic.com",
		IpAddress: "143.204.5.23",
	},
	{
		Domain:    "api1.platformdxc-d2.com",
		IpAddress: "13.35.4.200",
	},
	{
		Domain:    "smallpdf.com",
		IpAddress: "54.182.4.169",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.229",
	},
	{
		Domain:    "geocomply.net",
		IpAddress: "143.204.2.82",
	},
	{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "13.249.5.29",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.42",
	},
	{
		Domain:    "blim.com",
		IpAddress: "13.35.2.206",
	},
	{
		Domain:    "www.realizedev-test.com",
		IpAddress: "52.222.132.130",
	},
	{
		Domain:    "amazonsmile.com",
		IpAddress: "204.246.177.77",
	},
	{
		Domain:    "www.audible.com.au",
		IpAddress: "99.84.2.38",
	},
	{
		Domain:    "lucidhq.com",
		IpAddress: "99.86.6.52",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.42",
	},
	{
		Domain:    "www.53.localytics.com",
		IpAddress: "204.246.164.60",
	},
	{
		Domain:    "ekdgd.com",
		IpAddress: "143.204.5.126",
	},
	{
		Domain:    "twitchsvc.net",
		IpAddress: "204.246.169.158",
	},
	{
		Domain:    "mheducation.com",
		IpAddress: "52.222.134.79",
	},
	{
		Domain:    "zeasn.tv",
		IpAddress: "204.246.169.108",
	},
	{
		Domain:    "www.dev.pos.paylabo.com",
		IpAddress: "99.86.1.54",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.168",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.33",
	},
	{
		Domain:    "int3.machieco.nestle.jp",
		IpAddress: "52.222.135.49",
	},
	{
		Domain:    "download.epicgames.com",
		IpAddress: "54.239.192.66",
	},
	{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "99.84.6.108",
	},
	{
		Domain:    "a1v.starfall.com",
		IpAddress: "204.246.177.46",
	},
	{
		Domain:    "democrats.org",
		IpAddress: "204.246.178.41",
	},
	{
		Domain:    "www7.amazon.com",
		IpAddress: "54.239.130.171",
	},
	{
		Domain:    "www.tipico.com",
		IpAddress: "204.246.177.191",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.22",
	},
	{
		Domain:    "datadoghq.com",
		IpAddress: "65.8.214.61",
	},
	{
		Domain:    "demandbase.com",
		IpAddress: "143.204.5.108",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.7",
	},
	{
		Domain:    "www.staging.truecardev.com",
		IpAddress: "13.35.0.176",
	},
	{
		Domain:    "cf.test.frontier.a2z.com",
		IpAddress: "143.204.2.34",
	},
	{
		Domain:    "sftelemetry.sophos.com",
		IpAddress: "143.204.1.86",
	},
	{
		Domain:    "brain-market.com",
		IpAddress: "13.35.3.22",
	},
	{
		Domain:    "oneblood.org",
		IpAddress: "13.249.2.139",
	},
	{
		Domain:    "www.predix.io",
		IpAddress: "143.204.2.217",
	},
	{
		Domain:    "club-beta2.pokemon.com",
		IpAddress: "204.246.164.166",
	},
	{
		Domain:    "inspector-agent.amazonaws.com",
		IpAddress: "13.35.1.152",
	},
	{
		Domain:    "www.studysapuri.jp",
		IpAddress: "204.246.164.59",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.173",
	},
	{
		Domain:    "bc-citi.providersml.com",
		IpAddress: "13.35.2.155",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.193",
	},
	{
		Domain:    "mobile.mercadopago.com",
		IpAddress: "108.158.166.197",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "99.86.5.57",
	},
	{
		Domain:    "test.samsunghealth.com",
		IpAddress: "54.239.192.172",
	},
	{
		Domain:    "www.realizedev-test.com",
		IpAddress: "13.224.6.183",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.111",
	},
	{
		Domain:    "jwpsrv.com",
		IpAddress: "52.222.132.182",
	},
	{
		Domain:    "www.cequintsptecid.com",
		IpAddress: "54.182.4.216",
	},
	{
		Domain:    "mark1.dev",
		IpAddress: "204.246.178.88",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "52.222.131.118",
	},
	{
		Domain:    "code.org",
		IpAddress: "99.86.6.227",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.16",
	},
	{
		Domain:    "knowledgevision.com",
		IpAddress: "13.249.6.171",
	},
	{
		Domain:    "www.dwell.com",
		IpAddress: "13.35.5.126",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.124",
	},
	{
		Domain:    "fe.dazn-stage.com",
		IpAddress: "54.239.130.147",
	},
	{
		Domain:    "www.allianz-connect.com",
		IpAddress: "204.246.169.145",
	},
	{
		Domain:    "www.cp.misumi.jp",
		IpAddress: "54.182.3.154",
	},
	{
		Domain:    "www.iglobalstores.com",
		IpAddress: "13.35.4.110",
	},
	{
		Domain:    "www.tfly-aws.com",
		IpAddress: "52.222.134.36",
	},
	{
		Domain:    "payments.zynga.com",
		IpAddress: "99.86.2.135",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.129",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "99.84.2.194",
	},
	{
		Domain:    "cascade.madmimi.com",
		IpAddress: "13.35.3.201",
	},
	{
		Domain:    "api.mapbox.com",
		IpAddress: "13.35.1.183",
	},
	{
		Domain:    "www.test.iot.irobotapi.com",
		IpAddress: "99.84.2.101",
	},
	{
		Domain:    "offerup.com",
		IpAddress: "99.84.6.126",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "143.204.5.153",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.148",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.27",
	},
	{
		Domain:    "virmanig.myinstance.com",
		IpAddress: "54.182.7.61",
	},
	{
		Domain:    "coincheck.com",
		IpAddress: "13.35.2.34",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.106",
	},
	{
		Domain:    "prcp.pass.auone.jp",
		IpAddress: "13.224.5.108",
	},
	{
		Domain:    "kddi-fs.com",
		IpAddress: "143.204.2.103",
	},
	{
		Domain:    "stag.dazn.com",
		IpAddress: "204.246.177.149",
	},
	{
		Domain:    "www.sigalert.com",
		IpAddress: "99.84.0.6",
	},
	{
		Domain:    "nowforce.com",
		IpAddress: "52.222.135.12",
	},
	{
		Domain:    "truste.com",
		IpAddress: "13.35.2.195",
	},
	{
		Domain:    "login.schibsted.com",
		IpAddress: "13.35.6.148",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.31",
	},
	{
		Domain:    "www.linebc.jp",
		IpAddress: "54.182.4.177",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.109",
	},
	{
		Domain:    "werally.com",
		IpAddress: "13.249.6.176",
	},
	{
		Domain:    "www.toukei-kentei.jp",
		IpAddress: "13.35.5.50",
	},
	{
		Domain:    "www.nmrodam.com",
		IpAddress: "143.204.6.50",
	},
	{
		Domain:    "ccpsx.com",
		IpAddress: "13.35.4.35",
	},
	{
		Domain:    "www.vistarmedia.com",
		IpAddress: "143.204.2.86",
	},
	{
		Domain:    "www.connectwisedev.com",
		IpAddress: "99.84.6.153",
	},
	{
		Domain:    "tigocloud.net",
		IpAddress: "99.86.6.197",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.225",
	},
	{
		Domain:    "liftoff.io",
		IpAddress: "99.86.0.215",
	},
	{
		Domain:    "cdn.discounttire.com",
		IpAddress: "54.239.130.112",
	},
	{
		Domain:    "api.sandbox.repayonline.com",
		IpAddress: "143.204.5.225",
	},
	{
		Domain:    "mercadopago.com",
		IpAddress: "13.227.126.107",
	},
	{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "52.192.248.133",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.59",
	},
	{
		Domain:    "ecnavi.jp",
		IpAddress: "13.224.6.10",
	},
	{
		Domain:    "www.amazon.sa",
		IpAddress: "18.67.145.124",
	},
	{
		Domain:    "workflow-stage.licenses.adobe.com",
		IpAddress: "204.246.164.210",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.35.1.190",
	},
	{
		Domain:    "omsdocs.magento.com",
		IpAddress: "54.182.6.200",
	},
	{
		Domain:    "rca-upload-cloudstation-us-west-2.prod.hydra.sophos.com",
		IpAddress: "13.35.6.156",
	},
	{
		Domain:    "www.cafewell.com",
		IpAddress: "99.84.6.202",
	},
	{
		Domain:    "cdn.mozilla.net",
		IpAddress: "13.224.5.58",
	},
	{
		Domain:    "test.dazn.com",
		IpAddress: "99.86.1.48",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "99.84.0.212",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.41",
	},
	{
		Domain:    "abcmouse.com",
		IpAddress: "99.86.4.33",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.4.12",
	},
	{
		Domain:    "www.sf-cdn.net",
		IpAddress: "99.86.1.142",
	},
	{
		Domain:    "sni.to",
		IpAddress: "99.86.5.20",
	},
	{
		Domain:    "www.api.brightcove.com",
		IpAddress: "143.204.6.126",
	},
	{
		Domain:    "www.connectwise.com",
		IpAddress: "99.86.3.187",
	},
	{
		Domain:    "liftoff.io",
		IpAddress: "54.182.3.69",
	},
	{
		Domain:    "mojang.com",
		IpAddress: "143.204.2.134",
	},
	{
		Domain:    "eprocurement.marketplace.us-east-1.amazonaws.com",
		IpAddress: "99.84.7.42",
	},
	{
		Domain:    "tonglueyun.com",
		IpAddress: "13.35.2.65",
	},
	{
		Domain:    "highwebmedia.com",
		IpAddress: "52.222.131.158",
	},
	{
		Domain:    "www.desmos.com",
		IpAddress: "13.224.5.90",
	},
	{
		Domain:    "flipagram.com",
		IpAddress: "13.35.1.103",
	},
	{
		Domain:    "oqs.amb.cybird.ne.jp",
		IpAddress: "204.246.164.101",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.98",
	},
	{
		Domain:    "z-eu.associates-amazon.com",
		IpAddress: "13.35.1.27",
	},
	{
		Domain:    "alexa.amazon.com.mx",
		IpAddress: "52.222.134.131",
	},
	{
		Domain:    "forgesvc.net",
		IpAddress: "54.182.6.144",
	},
	{
		Domain:    "aa0.awsstatic.com",
		IpAddress: "13.224.5.28",
	},
	{
		Domain:    "signal.is",
		IpAddress: "99.84.2.196",
	},
	{
		Domain:    "api.mistore.jp",
		IpAddress: "143.204.1.71",
	},
	{
		Domain:    "tripkit-test5.jeppesen.com",
		IpAddress: "99.84.2.116",
	},
	{
		Domain:    "company-target.com",
		IpAddress: "13.224.6.192",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.51",
	},
	{
		Domain:    "bittorrent.com",
		IpAddress: "13.35.0.167",
	},
	{
		Domain:    "www.lottedfs.com",
		IpAddress: "99.86.2.19",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.76",
	},
	{
		Domain:    "zuora.identity.fcl-02.prep.fcagcv.com",
		IpAddress: "205.251.212.182",
	},
	{
		Domain:    "amazon.ca",
		IpAddress: "204.246.169.232",
	},
	{
		Domain:    "www.accordiagolf.com",
		IpAddress: "143.204.6.150",
	},
	{
		Domain:    "custom-api.bigpanda.io",
		IpAddress: "54.239.130.107",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.184",
	},
	{
		Domain:    "identitynow.com",
		IpAddress: "99.84.6.47",
	},
	{
		Domain:    "simple-workflow.licenses.adobe.com",
		IpAddress: "99.86.5.16",
	},
	{
		Domain:    "wework.com",
		IpAddress: "52.222.129.203",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.184",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.4",
	},
	{
		Domain:    "stage-spectrum.net",
		IpAddress: "143.204.5.9",
	},
	{
		Domain:    "ix-cdn.brightedge.com",
		IpAddress: "143.204.0.139",
	},
	{
		Domain:    "sunsky-online.com",
		IpAddress: "99.84.0.14",
	},
	{
		Domain:    "edge.disstg.commercecloud.salesforce.com",
		IpAddress: "99.86.4.165",
	},
	{
		Domain:    "toysrus.co.jp",
		IpAddress: "13.35.2.29",
	},
	{
		Domain:    "iot.ap-southeast-2.amazonaws.com",
		IpAddress: "99.84.2.155",
	},
	{
		Domain:    "mix.tokyo",
		IpAddress: "99.84.0.120",
	},
	{
		Domain:    "img-en.fs.com",
		IpAddress: "13.249.2.110",
	},
	{
		Domain:    "gomlab.com",
		IpAddress: "54.182.4.103",
	},
	{
		Domain:    "update.hicloud.com",
		IpAddress: "13.35.1.45",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.78",
	},
	{
		Domain:    "www.cafewell.com",
		IpAddress: "54.182.4.149",
	},
	{
		Domain:    "myfitnesspal.com.tw",
		IpAddress: "143.204.1.69",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.81",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.131",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.88",
	},
	{
		Domain:    "seesaw.me",
		IpAddress: "54.239.130.206",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.147",
	},
	{
		Domain:    "segment.com",
		IpAddress: "99.86.0.85",
	},
	{
		Domain:    "sings-download.twitch.tv",
		IpAddress: "99.86.4.124",
	},
	{
		Domain:    "dsdfpay.com",
		IpAddress: "13.35.2.20",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "99.86.4.208",
	},
	{
		Domain:    "www.static.lottedfs.com",
		IpAddress: "143.204.6.76",
	},
	{
		Domain:    "www.linebc.jp",
		IpAddress: "52.222.134.178",
	},
	{
		Domain:    "preprod.apac.amway.net",
		IpAddress: "143.204.1.166",
	},
	{
		Domain:    "cdn.burlingtonenglish.com",
		IpAddress: "52.222.131.236",
	},
	{
		Domain:    "www.chatbar.me",
		IpAddress: "99.86.4.90",
	},
	{
		Domain:    "us.whispir.com",
		IpAddress: "54.239.130.204",
	},
	{
		Domain:    "pod-point.com",
		IpAddress: "13.249.5.54",
	},
	{
		Domain:    "t-x.io",
		IpAddress: "143.204.6.217",
	},
	{
		Domain:    "ads-interfaces.sc-cdn.net",
		IpAddress: "99.86.3.166",
	},
	{
		Domain:    "www.quipper.net",
		IpAddress: "54.239.192.48",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.73",
	},
	{
		Domain:    "cardgames.io",
		IpAddress: "143.204.7.40",
	},
	{
		Domain:    "primer.typekit.net",
		IpAddress: "13.249.5.2",
	},
	{
		Domain:    "www.accuplacer.org",
		IpAddress: "99.84.0.231",
	},
	{
		Domain:    "www.epop.cf.eu.aiv-cdn.net",
		IpAddress: "204.246.178.165",
	},
	{
		Domain:    "www.fp.ps.easebar.com",
		IpAddress: "52.222.129.108",
	},
	{
		Domain:    "club-beta2.pokemon.com",
		IpAddress: "54.182.2.19",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.31",
	},
	{
		Domain:    "smartrecruiters.com",
		IpAddress: "13.35.5.77",
	},
	{
		Domain:    "product-downloads.atlassian.com",
		IpAddress: "143.204.7.33",
	},
	{
		Domain:    "dev.faceid.paylabo.com",
		IpAddress: "54.182.3.224",
	},
	{
		Domain:    "mymathacademy.com",
		IpAddress: "204.246.164.80",
	},
	{
		Domain:    "amazon.co.jp",
		IpAddress: "13.35.1.16",
	},
	{
		Domain:    "samsungknowledge.com",
		IpAddress: "143.204.2.218",
	},
	{
		Domain:    "s3-turbo.amazonaws.com",
		IpAddress: "13.35.2.190",
	},
	{
		Domain:    "device-firmware.gp-static.com",
		IpAddress: "54.182.0.179",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.77",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "205.251.213.10",
	},
	{
		Domain:    "bd0.awsstatic.com",
		IpAddress: "99.86.2.33",
	},
	{
		Domain:    "static.pontoslivelo.com.br",
		IpAddress: "204.246.177.114",
	},
	{
		Domain:    "enigmasoftware.com",
		IpAddress: "216.137.39.62",
	},
	{
		Domain:    "adtpulseaws.net",
		IpAddress: "143.204.5.12",
	},
	{
		Domain:    "amazon.co.uk",
		IpAddress: "54.239.195.221",
	},
	{
		Domain:    "workflow-stage.licenses.adobe.com",
		IpAddress: "13.35.2.213",
	},
	{
		Domain:    "www.dev.pos.paylabo.com",
		IpAddress: "13.35.2.117",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.185",
	},
	{
		Domain:    "ads.chtbl.com",
		IpAddress: "99.86.2.117",
	},
	{
		Domain:    "keyuca.com",
		IpAddress: "143.204.2.164",
	},
	{
		Domain:    "tvc-mall.com",
		IpAddress: "13.35.2.78",
	},
	{
		Domain:    "static.datadoghq.com",
		IpAddress: "13.224.5.68",
	},
	{
		Domain:    "mymathacademy.com",
		IpAddress: "143.204.5.189",
	},
	{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "13.35.3.15",
	},
	{
		Domain:    "www.playwithsea.com",
		IpAddress: "99.84.0.81",
	},
	{
		Domain:    "cdn.sw.altova.com",
		IpAddress: "143.204.2.113",
	},
	{
		Domain:    "www.siksine.com",
		IpAddress: "54.239.192.51",
	},
	{
		Domain:    "static.emarsys.com",
		IpAddress: "99.84.6.190",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.120",
	},
	{
		Domain:    "www.samsungsmartcam.com",
		IpAddress: "54.182.7.12",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.21",
	},
	{
		Domain:    "oihxray-fe.aka.amazon.com",
		IpAddress: "99.84.7.43",
	},
	{
		Domain:    "static.adobelogin.com",
		IpAddress: "143.204.3.70",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.177",
	},
	{
		Domain:    "www.cafewellstage.com",
		IpAddress: "52.222.130.237",
	},
	{
		Domain:    "static.lendingclub.com",
		IpAddress: "13.35.3.33",
	},
	{
		Domain:    "arkoselabs.com",
		IpAddress: "13.224.6.49",
	},
	{
		Domain:    "ekdgd.com",
		IpAddress: "204.246.178.96",
	},
	{
		Domain:    "cdn.shptrn.com",
		IpAddress: "54.182.3.46",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.8",
	},
	{
		Domain:    "beta.awsapps.com",
		IpAddress: "54.239.130.36",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.118",
	},
	{
		Domain:    "forgesvc.net",
		IpAddress: "99.86.6.66",
	},
	{
		Domain:    "ccpsx.com",
		IpAddress: "99.84.6.111",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.86",
	},
	{
		Domain:    "boleto.pagseguro.com.br",
		IpAddress: "52.222.128.225",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.0.191",
	},
	{
		Domain:    "api1.platformdxc-d2.com",
		IpAddress: "13.249.6.218",
	},
	{
		Domain:    "specialized.com",
		IpAddress: "99.84.7.34",
	},
	{
		Domain:    "test.api.seek.co.nz",
		IpAddress: "99.84.0.86",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.148",
	},
	{
		Domain:    "twitchcdn.net",
		IpAddress: "54.182.2.159",
	},
	{
		Domain:    "kuvo.com",
		IpAddress: "204.246.164.230",
	},
	{
		Domain:    "api.shopbop.com",
		IpAddress: "99.84.2.45",
	},
	{
		Domain:    "www.srv.ygles-test.com",
		IpAddress: "99.86.0.162",
	},
	{
		Domain:    "www.apteligent.com",
		IpAddress: "204.246.177.4",
	},
	{
		Domain:    "www.accordiagolf.com",
		IpAddress: "13.35.1.150",
	},
	{
		Domain:    "static.uber-adsystem.com",
		IpAddress: "13.35.2.42",
	},
	{
		Domain:    "versal.com",
		IpAddress: "99.86.2.35",
	},
	{
		Domain:    "cptuat.net",
		IpAddress: "52.222.132.38",
	},
	{
		Domain:    "dsdfpay.com",
		IpAddress: "52.222.132.125",
	},
	{
		Domain:    "parsely.com",
		IpAddress: "99.84.2.17",
	},
	{
		Domain:    "www.neuweb.biz",
		IpAddress: "99.86.5.182",
	},
	{
		Domain:    "predix.io",
		IpAddress: "13.35.6.163",
	},
	{
		Domain:    "www.neuweb.biz",
		IpAddress: "13.249.5.184",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.145",
	},
	{
		Domain:    "www.project-a.videoprojects.net",
		IpAddress: "54.182.5.149",
	},
	{
		Domain:    "www.airchip.com",
		IpAddress: "99.86.3.197",
	},
	{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "13.35.4.154",
	},
	{
		Domain:    "origin-help.imdb.com",
		IpAddress: "52.222.131.73",
	},
	{
		Domain:    "cdn.hands.net",
		IpAddress: "52.222.132.13",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.14",
	},
	{
		Domain:    "opmsec.sophos.com",
		IpAddress: "13.224.5.69",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-west-1.prod.hydra.sophos.com",
		IpAddress: "99.84.6.222",
	},
	{
		Domain:    "fifaconnect.org",
		IpAddress: "99.84.2.216",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.15",
	},
	{
		Domain:    "www.mytaxi.com",
		IpAddress: "99.86.6.223",
	},
	{
		Domain:    "olt-content.sans.org",
		IpAddress: "143.204.5.18",
	},
	{
		Domain:    "orgsync.com",
		IpAddress: "99.86.5.44",
	},
	{
		Domain:    "www.listrakbi.com",
		IpAddress: "99.84.6.9",
	},
	{
		Domain:    "abcmouse.com",
		IpAddress: "54.239.192.214",
	},
	{
		Domain:    "www.withbuddies.com",
		IpAddress: "205.251.212.194",
	},
	{
		Domain:    "panda.chtbl.com",
		IpAddress: "54.239.192.235",
	},
	{
		Domain:    "www.tosconfig.com",
		IpAddress: "99.84.0.7",
	},
	{
		Domain:    "nba-cdn.2ksports.com",
		IpAddress: "54.239.192.183",
	},
	{
		Domain:    "www.culqi.com",
		IpAddress: "204.246.169.162",
	},
	{
		Domain:    "www.adison.co",
		IpAddress: "52.222.134.82",
	},
	{
		Domain:    "angular.mrowl.com",
		IpAddress: "99.86.5.228",
	},
	{
		Domain:    "js.pusher.com",
		IpAddress: "204.246.164.33",
	},
	{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "13.249.2.17",
	},
	{
		Domain:    "iot.us-west-2.amazonaws.com",
		IpAddress: "143.204.1.41",
	},
	{
		Domain:    "forgecdn.net",
		IpAddress: "99.86.4.184",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.68",
	},
	{
		Domain:    "mheducation.com",
		IpAddress: "13.35.2.71",
	},
	{
		Domain:    "media.aircorsica.com",
		IpAddress: "13.249.2.122",
	},
	{
		Domain:    "www.api.brightcove.com",
		IpAddress: "204.246.177.6",
	},
	{
		Domain:    "ad1.awsstatic.com",
		IpAddress: "99.86.1.95",
	},
	{
		Domain:    "www.brinkpos.net",
		IpAddress: "54.239.130.65",
	},
	{
		Domain:    "www.twitch.tv",
		IpAddress: "99.86.0.72",
	},
	{
		Domain:    "behance.net",
		IpAddress: "54.239.192.115",
	},
	{
		Domain:    "tripkit-test2.jeppesen.com",
		IpAddress: "99.86.2.75",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.38",
	},
	{
		Domain:    "qa.o.brightcove.com",
		IpAddress: "204.246.169.37",
	},
	{
		Domain:    "payment.global.rakuten.com",
		IpAddress: "13.35.3.96",
	},
	{
		Domain:    "origin-gql.beta.api.imdb.a2z.com",
		IpAddress: "143.204.5.147",
	},
	{
		Domain:    "img-en.fs.com",
		IpAddress: "13.35.5.90",
	},
	{
		Domain:    "cofanet.coface.com",
		IpAddress: "54.182.4.137",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.154",
	},
	{
		Domain:    "www.o9.de",
		IpAddress: "54.182.6.126",
	},
	{
		Domain:    "virmanig.myinstance.com",
		IpAddress: "52.222.134.61",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.59",
	},
	{
		Domain:    "www.dst.vpsvc.com",
		IpAddress: "143.204.5.89",
	},
	{
		Domain:    "s3-accelerate.amazonaws.com",
		IpAddress: "143.204.2.199",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.92",
	},
	{
		Domain:    "static.datadoghq.com",
		IpAddress: "99.86.1.203",
	},
	{
		Domain:    "www.swipesense.com",
		IpAddress: "143.204.1.81",
	},
	{
		Domain:    "ecnavi.jp",
		IpAddress: "54.239.192.18",
	},
	{
		Domain:    "www.cafewellstage.com",
		IpAddress: "143.204.2.15",
	},
	{
		Domain:    "multisandbox.api.fluentretail.com",
		IpAddress: "204.246.177.110",
	},
	{
		Domain:    "mheducation.com",
		IpAddress: "54.182.2.11",
	},
	{
		Domain:    "mfi-tc02.fnopf.jp",
		IpAddress: "99.86.1.187",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.17",
	},
	{
		Domain:    "sftelemetry-test.sophos.com",
		IpAddress: "13.35.4.10",
	},
	{
		Domain:    "qtest.abcmouse.com",
		IpAddress: "13.35.4.11",
	},
	{
		Domain:    "clients.amazonworkspaces.com",
		IpAddress: "143.204.2.114",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.25",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.25",
	},
	{
		Domain:    "amazon.co.jp",
		IpAddress: "99.86.1.112",
	},
	{
		Domain:    "mojang.com",
		IpAddress: "13.35.5.134",
	},
	{
		Domain:    "gluon-cv.mxnet.io",
		IpAddress: "52.222.131.187",
	},
	{
		Domain:    "www.bijiaqi.xyz",
		IpAddress: "13.249.6.6",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.2.129",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.7",
	},
	{
		Domain:    "assets.cameloteurope.com",
		IpAddress: "99.86.6.109",
	},
	{
		Domain:    "twitchsvc-shadow.net",
		IpAddress: "143.204.6.23",
	},
	{
		Domain:    "z-na.associates-amazon.com",
		IpAddress: "143.204.6.100",
	},
	{
		Domain:    "cf.pumlo.awsps.myinstance.com",
		IpAddress: "204.246.178.16",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "54.182.6.226",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.117",
	},
	{
		Domain:    "api.imdbws.com",
		IpAddress: "204.246.164.5",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.52",
	},
	{
		Domain:    "www.stg.ui.com",
		IpAddress: "52.222.131.109",
	},
	{
		Domain:    "qpyou.cn",
		IpAddress: "13.224.7.29",
	},
	{
		Domain:    "gbf.game-a.mbga.jp",
		IpAddress: "54.182.6.23",
	},
	{
		Domain:    "www.aya.quipper.net",
		IpAddress: "54.182.6.173",
	},
	{
		Domain:    "forestry.trimble.com",
		IpAddress: "52.222.131.238",
	},
	{
		Domain:    "ap1.whispir.com",
		IpAddress: "205.251.212.140",
	},
	{
		Domain:    "openfin.co",
		IpAddress: "143.204.1.117",
	},
	{
		Domain:    "wordsearchbible.com",
		IpAddress: "13.35.0.231",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.228",
	},
	{
		Domain:    "www.mytaxi.com",
		IpAddress: "204.246.178.147",
	},
	{
		Domain:    "www.sodexomyway.com",
		IpAddress: "13.35.4.49",
	},
	{
		Domain:    "globalwip.cms.pearson.com",
		IpAddress: "13.35.1.233",
	},
	{
		Domain:    "pv.media-amazon.com",
		IpAddress: "52.222.131.206",
	},
	{
		Domain:    "www.execute-api.us-west-2.amazonaws.com",
		IpAddress: "54.182.3.135",
	},
	{
		Domain:    "prod1.superobscuredomains.com",
		IpAddress: "143.204.2.243",
	},
	{
		Domain:    "www.dst.vpsvc.com",
		IpAddress: "204.246.164.89",
	},
	{
		Domain:    "playwith.com.tw",
		IpAddress: "99.86.0.148",
	},
	{
		Domain:    "www.indigoag.build",
		IpAddress: "52.222.129.232",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.127",
	},
	{
		Domain:    "envysion.com",
		IpAddress: "13.224.5.125",
	},
	{
		Domain:    "i.fyu.se",
		IpAddress: "52.222.134.137",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.81",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.36",
	},
	{
		Domain:    "smile.amazon.de",
		IpAddress: "52.222.134.71",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.192",
	},
	{
		Domain:    "www.iglobalstores.com",
		IpAddress: "99.86.3.137",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.136",
	},
	{
		Domain:    "www.brinkpos.net",
		IpAddress: "54.182.3.65",
	},
	{
		Domain:    "kaercher.com",
		IpAddress: "99.84.0.105",
	},
	{
		Domain:    "dfoneople.com",
		IpAddress: "52.222.128.198",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.134.242",
	},
	{
		Domain:    "www.awscfdns.com",
		IpAddress: "99.84.0.107",
	},
	{
		Domain:    "www.adison.co",
		IpAddress: "13.224.5.162",
	},
	{
		Domain:    "saiercdn.imtxwy.com",
		IpAddress: "143.204.6.220",
	},
	{
		Domain:    "www.c.misumi-ec.com",
		IpAddress: "13.35.4.98",
	},
	{
		Domain:    "www.cloud.tenable.com",
		IpAddress: "99.84.5.238",
	},
	{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "54.182.2.238",
	},
	{
		Domain:    "smtown.com",
		IpAddress: "99.86.4.11",
	},
	{
		Domain:    "origin-api.amazonalexa.com",
		IpAddress: "54.182.6.235",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.45",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.78",
	},
	{
		Domain:    "paradoxplaza.com",
		IpAddress: "13.224.7.16",
	},
	{
		Domain:    "www.ladymay.net",
		IpAddress: "143.204.1.170",
	},
	{
		Domain:    "mcoc-cdn.net",
		IpAddress: "52.222.130.146",
	},
	{
		Domain:    "gimmegimme.it",
		IpAddress: "99.84.6.59",
	},
	{
		Domain:    "www.webapp.easebar.com",
		IpAddress: "13.35.3.61",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.21",
	},
	{
		Domain:    "www.goldspotmedia.com",
		IpAddress: "99.86.4.52",
	},
	{
		Domain:    "smtown.com",
		IpAddress: "54.182.4.202",
	},
	{
		Domain:    "syapp.jp",
		IpAddress: "13.249.5.36",
	},
	{
		Domain:    "www.uat.catchplay.com",
		IpAddress: "54.182.3.118",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.17",
	},
	{
		Domain:    "angels.camp-fire.jp",
		IpAddress: "143.204.6.45",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.33",
	},
	{
		Domain:    "www.midasplayer.com",
		IpAddress: "99.84.0.20",
	},
	{
		Domain:    "sellercentral.amazon.com",
		IpAddress: "99.86.6.59",
	},
	{
		Domain:    "gcsp.jnj.com",
		IpAddress: "54.239.192.139",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.189",
	},
	{
		Domain:    "www.culqi.com",
		IpAddress: "99.84.2.239",
	},
	{
		Domain:    "verti.iptiq.de",
		IpAddress: "54.239.130.181",
	},
	{
		Domain:    "kucoin.com",
		IpAddress: "54.182.3.54",
	},
	{
		Domain:    "versal.com",
		IpAddress: "205.251.212.103",
	},
	{
		Domain:    "rest.immobilienscout24.de",
		IpAddress: "13.35.1.151",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.179",
	},
	{
		Domain:    "www.brickworksoftware.com",
		IpAddress: "143.204.6.216",
	},
	{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "204.246.178.146",
	},
	{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "54.182.3.37",
	},
	{
		Domain:    "api.mapbox.com",
		IpAddress: "99.86.1.140",
	},
	{
		Domain:    "qpyou.cn",
		IpAddress: "204.246.164.92",
	},
	{
		Domain:    "thescore.com",
		IpAddress: "13.35.5.152",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.19",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.152",
	},
	{
		Domain:    "smartica.jp",
		IpAddress: "99.84.0.164",
	},
	{
		Domain:    "www.tfly-aws.com",
		IpAddress: "99.86.3.147",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.69",
	},
	{
		Domain:    "fe.dazn-stage.com",
		IpAddress: "13.249.6.77",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.31",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.35.2.108",
	},
	{
		Domain:    "clients.a.chime.aws",
		IpAddress: "13.224.5.171",
	},
	{
		Domain:    "passporthealthglobal.com",
		IpAddress: "99.86.1.100",
	},
	{
		Domain:    "dl.amazon.co.uk",
		IpAddress: "99.86.2.34",
	},
	{
		Domain:    "samsunghealth.com",
		IpAddress: "204.246.164.214",
	},
	{
		Domain:    "www.xp-assets.aiv-cdn.net",
		IpAddress: "99.86.4.189",
	},
	{
		Domain:    "dev.awsapps.com",
		IpAddress: "13.35.6.36",
	},
	{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "204.246.177.85",
	},
	{
		Domain:    "ubnt.com",
		IpAddress: "13.35.5.112",
	},
	{
		Domain:    "dev.twitch.tv",
		IpAddress: "13.224.5.153",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "54.239.130.207",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.18",
	},
	{
		Domain:    "origin-www.amazon.com.tr",
		IpAddress: "52.222.134.218",
	},
	{
		Domain:    "cdn.realtimeprocess.net",
		IpAddress: "52.222.128.177",
	},
	{
		Domain:    "cdn.burlingtonenglish.com",
		IpAddress: "13.249.2.98",
	},
	{
		Domain:    "enetscores.com",
		IpAddress: "99.86.2.84",
	},
	{
		Domain:    "s3-turbo.amazonaws.com",
		IpAddress: "52.222.131.175",
	},
	{
		Domain:    "seesaw.me",
		IpAddress: "54.239.195.207",
	},
	{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "13.35.6.155",
	},
	{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "143.204.1.12",
	},
	{
		Domain:    "classic.dm.amplience.net",
		IpAddress: "13.224.6.86",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.103",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.133",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.45",
	},
	{
		Domain:    "www.recoru.in",
		IpAddress: "52.222.129.155",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.82",
	},
	{
		Domain:    "coupang.net",
		IpAddress: "54.239.192.206",
	},
	{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "204.246.177.33",
	},
	{
		Domain:    "www.ooyala.com",
		IpAddress: "99.84.2.225",
	},
	{
		Domain:    "www.pearsonperspective.com",
		IpAddress: "143.204.2.70",
	},
	{
		Domain:    "cdn.fdp.foreflight.com",
		IpAddress: "13.35.6.28",
	},
	{
		Domain:    "camp-fire.jp",
		IpAddress: "205.251.212.38",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.189",
	},
	{
		Domain:    "apps.bahrain.bh",
		IpAddress: "99.86.6.96",
	},
	{
		Domain:    "www.quipper.com",
		IpAddress: "99.86.6.173",
	},
	{
		Domain:    "dev.sotappm.auone.jp",
		IpAddress: "99.84.6.151",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.169",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.27",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.140",
	},
	{
		Domain:    "adtpulseaws.net",
		IpAddress: "13.35.3.181",
	},
	{
		Domain:    "cont-test.mydaiz.jp",
		IpAddress: "204.246.178.51",
	},
	{
		Domain:    "payment.global.rakuten.com",
		IpAddress: "52.222.131.101",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.221",
	},
	{
		Domain:    "www.freshdesk.com",
		IpAddress: "54.182.6.230",
	},
	{
		Domain:    "www.test.iot.irobotapi.com",
		IpAddress: "13.224.6.150",
	},
	{
		Domain:    "mark1.dev",
		IpAddress: "13.224.7.33",
	},
	{
		Domain:    "auth.nightowlx.com",
		IpAddress: "13.35.1.86",
	},
	{
		Domain:    "media.edgenuity.com",
		IpAddress: "13.224.5.206",
	},
	{
		Domain:    "iot.eu-west-1.amazonaws.com",
		IpAddress: "13.35.6.97",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.119",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.77",
	},
	{
		Domain:    "custom-api.bigpanda.io",
		IpAddress: "13.35.6.130",
	},
	{
		Domain:    "souqcdn.com",
		IpAddress: "13.35.3.69",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "54.182.4.112",
	},
	{
		Domain:    "www.bcovlive.io",
		IpAddress: "52.222.132.176",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.15",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.143",
	},
	{
		Domain:    "www.milkvr.rocks",
		IpAddress: "204.246.164.119",
	},
	{
		Domain:    "dsdfpay.com",
		IpAddress: "13.249.2.216",
	},
	{
		Domain:    "amoad.com",
		IpAddress: "54.182.4.231",
	},
	{
		Domain:    "www.me2zengame.com",
		IpAddress: "13.35.3.117",
	},
	{
		Domain:    "esd.sentinelcloud.com",
		IpAddress: "99.84.6.78",
	},
	{
		Domain:    "login.schibsted.com",
		IpAddress: "204.246.164.147",
	},
	{
		Domain:    "wpcp.shiseido.co.jp",
		IpAddress: "205.251.212.233",
	},
	{
		Domain:    "one.accedo.tv",
		IpAddress: "204.246.169.10",
	},
	{
		Domain:    "www.binance.vision",
		IpAddress: "52.222.132.68",
	},
	{
		Domain:    "www.apkimage.io",
		IpAddress: "54.182.3.140",
	},
	{
		Domain:    "www.mytaxi.com",
		IpAddress: "54.182.5.163",
	},
	{
		Domain:    "www.uniqlo.com",
		IpAddress: "99.86.4.113",
	},
	{
		Domain:    "www.travelhook.com",
		IpAddress: "143.204.2.10",
	},
	{
		Domain:    "static.counsyl.com",
		IpAddress: "54.182.7.68",
	},
	{
		Domain:    "forestry.trimble.com",
		IpAddress: "143.204.5.122",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.6.163",
	},
	{
		Domain:    "www.clearlinkdata.com",
		IpAddress: "143.204.1.141",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.6.122",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.90",
	},
	{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "54.182.3.87",
	},
	{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "52.222.130.161",
	},
	{
		Domain:    "saucelabs.com",
		IpAddress: "13.35.1.160",
	},
	{
		Domain:    "snapfinance.com",
		IpAddress: "99.86.0.29",
	},
	{
		Domain:    "jtvnw-30eb2e4e018997e11b2884b1f80a025c.twitchcdn.net",
		IpAddress: "52.222.129.59",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.217",
	},
	{
		Domain:    "datad0g.com",
		IpAddress: "99.84.2.72",
	},
	{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "204.246.169.87",
	},
	{
		Domain:    "sparxcdn.net",
		IpAddress: "99.86.4.105",
	},
	{
		Domain:    "rubiconproject.com",
		IpAddress: "54.182.4.117",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.149",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-central-1.inf.hydra.sophos.com",
		IpAddress: "99.86.4.85",
	},
	{
		Domain:    "pimg.jp",
		IpAddress: "13.249.5.27",
	},
	{
		Domain:    "evident.io",
		IpAddress: "13.249.7.29",
	},
	{
		Domain:    "api.beta.tab.com.au",
		IpAddress: "99.84.6.53",
	},
	{
		Domain:    "dolphin-fe.amazon.com",
		IpAddress: "54.182.5.228",
	},
	{
		Domain:    "sywm-kr.gdl.imtxwy.com",
		IpAddress: "54.182.3.124",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.91",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.225",
	},
	{
		Domain:    "gallery.mailchimp.com",
		IpAddress: "13.249.5.78",
	},
	{
		Domain:    "www.binancechain.io",
		IpAddress: "54.239.130.144",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.11",
	},
	{
		Domain:    "iot.us-east-2.amazonaws.com",
		IpAddress: "13.249.6.198",
	},
	{
		Domain:    "z-na.associates-amazon.com",
		IpAddress: "54.182.6.51",
	},
	{
		Domain:    "static-cdn.jtvnw.net",
		IpAddress: "204.246.178.131",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.172",
	},
	{
		Domain:    "docomo-ntsupport.jp",
		IpAddress: "13.249.2.140",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.55",
	},
	{
		Domain:    "ubnt.com",
		IpAddress: "143.204.2.125",
	},
	{
		Domain:    "yieldoptimizer.com",
		IpAddress: "13.35.2.196",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.11",
	},
	{
		Domain:    "panda.chtbl.com",
		IpAddress: "204.246.169.18",
	},
	{
		Domain:    "ekdgd.com",
		IpAddress: "99.86.4.2",
	},
	{
		Domain:    "dev.sotappm.auone.jp",
		IpAddress: "13.35.2.4",
	},
	{
		Domain:    "webarchive.nationalarchives.gov.uk",
		IpAddress: "52.222.135.38",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "13.35.5.162",
	},
	{
		Domain:    "cpe.wtf",
		IpAddress: "13.35.3.225",
	},
	{
		Domain:    "oneblood.org",
		IpAddress: "99.84.2.137",
	},
	{
		Domain:    "thetvdb.com",
		IpAddress: "54.182.6.98",
	},
	{
		Domain:    "offerup.com",
		IpAddress: "52.222.135.30",
	},
	{
		Domain:    "www.channel4.com",
		IpAddress: "143.204.5.164",
	},
	{
		Domain:    "amazon.ca",
		IpAddress: "143.204.2.43",
	},
	{
		Domain:    "arevea.tv",
		IpAddress: "99.86.0.219",
	},
	{
		Domain:    "imbd-pro.net",
		IpAddress: "13.224.0.239",
	},
	{
		Domain:    "bd1.awsstatic.com",
		IpAddress: "143.204.1.113",
	},
	{
		Domain:    "club.ubisoft.com",
		IpAddress: "99.84.0.142",
	},
	{
		Domain:    "samsungqbe.com",
		IpAddress: "13.224.6.206",
	},
	{
		Domain:    "www.ladymay.net",
		IpAddress: "13.35.1.13",
	},
	{
		Domain:    "www.collegescheduler.com",
		IpAddress: "54.239.195.136",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "54.239.130.100",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.159",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.126",
	},
	{
		Domain:    "www.readingiq.com",
		IpAddress: "13.35.6.8",
	},
	{
		Domain:    "ring.com",
		IpAddress: "13.249.6.62",
	},
	{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "54.182.5.160",
	},
	{
		Domain:    "dsdfpay.com",
		IpAddress: "99.86.6.199",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.29",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.10",
	},
	{
		Domain:    "geocomply.net",
		IpAddress: "143.204.1.175",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.93",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.34",
	},
	{
		Domain:    "d.nanairo.coop",
		IpAddress: "54.182.4.82",
	},
	{
		Domain:    "scoring.pearsonassessments.com",
		IpAddress: "54.182.5.199",
	},
	{
		Domain:    "www.tipico.com",
		IpAddress: "143.204.6.52",
	},
	{
		Domain:    "www.adbephotos.com",
		IpAddress: "99.86.4.16",
	},
	{
		Domain:    "update.hicloud.com",
		IpAddress: "54.182.7.16",
	},
	{
		Domain:    "myfonts.net",
		IpAddress: "13.249.2.36",
	},
	{
		Domain:    "assets.cameloteurope.com",
		IpAddress: "143.204.1.112",
	},
	{
		Domain:    "widencdn.net",
		IpAddress: "52.222.128.185",
	},
	{
		Domain:    "pactsafe.io",
		IpAddress: "204.246.169.122",
	},
	{
		Domain:    "www.ebookstore.sony.jp",
		IpAddress: "13.249.7.32",
	},
	{
		Domain:    "lottedfs.com",
		IpAddress: "54.239.130.146",
	},
	{
		Domain:    "file.samsungcloud.com",
		IpAddress: "52.222.130.194",
	},
	{
		Domain:    "hicloud.com",
		IpAddress: "99.84.2.169",
	},
	{
		Domain:    "site.skychnl.net",
		IpAddress: "99.84.0.236",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "13.35.2.23",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-west-1.dev.hydra.sophos.com",
		IpAddress: "52.222.132.17",
	},
	{
		Domain:    "aa0.awsstatic.com",
		IpAddress: "143.204.7.30",
	},
	{
		Domain:    "d.nanairo.coop",
		IpAddress: "99.84.0.207",
	},
	{
		Domain:    "rlmcdn.net",
		IpAddress: "143.204.0.186",
	},
	{
		Domain:    "rebrandly.com",
		IpAddress: "54.239.192.92",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.78",
	},
	{
		Domain:    "livethumb.huluim.com",
		IpAddress: "13.35.6.136",
	},
	{
		Domain:    "www.sigalert.com",
		IpAddress: "13.35.3.6",
	},
	{
		Domain:    "us.whispir.com",
		IpAddress: "13.35.6.166",
	},
	{
		Domain:    "unrealengine.com",
		IpAddress: "54.182.0.209",
	},
	{
		Domain:    "yq.gph.imtxwy.com",
		IpAddress: "99.86.4.47",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.8",
	},
	{
		Domain:    "spd.samsungdm.com",
		IpAddress: "204.246.178.48",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.71",
	},
	{
		Domain:    "developercentral.amazon.com",
		IpAddress: "99.86.0.115",
	},
	{
		Domain:    "giv-dev.nmgcloud.io",
		IpAddress: "52.222.134.156",
	},
	{
		Domain:    "www.myharmony.com",
		IpAddress: "54.239.192.14",
	},
	{
		Domain:    "rlmcdn.net",
		IpAddress: "99.84.6.140",
	},
	{
		Domain:    "tapad.com",
		IpAddress: "99.86.5.2",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.22",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.218",
	},
	{
		Domain:    "rview.com",
		IpAddress: "143.204.2.185",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.4.23",
	},
	{
		Domain:    "binance.sg",
		IpAddress: "13.35.1.73",
	},
	{
		Domain:    "m.betaex.com",
		IpAddress: "54.239.192.223",
	},
	{
		Domain:    "www.nyc837-dev.gin-dev.com",
		IpAddress: "13.249.2.60",
	},
	{
		Domain:    "www.g.mkey.163.com",
		IpAddress: "99.86.5.51",
	},
	{
		Domain:    "www.findawayworld.com",
		IpAddress: "143.204.0.137",
	},
	{
		Domain:    "clients.chime.aws",
		IpAddress: "54.239.192.84",
	},
	{
		Domain:    "pubcerts-stage.licenses.adobe.com",
		IpAddress: "99.84.2.20",
	},
	{
		Domain:    "ring.com",
		IpAddress: "54.182.4.106",
	},
	{
		Domain:    "searchandexplore.com",
		IpAddress: "13.249.5.127",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.18",
	},
	{
		Domain:    "assets1.uswitch.com",
		IpAddress: "54.239.130.50",
	},
	{
		Domain:    "smile.amazon.de",
		IpAddress: "54.182.5.238",
	},
	{
		Domain:    "api.stage.context.cloud.sap",
		IpAddress: "13.224.6.156",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.99",
	},
	{
		Domain:    "static.ddog-gov.com",
		IpAddress: "204.246.178.46",
	},
	{
		Domain:    "club.ubisoft.com",
		IpAddress: "52.222.132.95",
	},
	{
		Domain:    "vdownload.cyberoam.com",
		IpAddress: "13.35.1.197",
	},
	{
		Domain:    "whopper.com",
		IpAddress: "99.86.6.236",
	},
	{
		Domain:    "vlive-simulcast.sans.org",
		IpAddress: "204.246.177.102",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.86.0.127",
	},
	{
		Domain:    "www.ladymay.net",
		IpAddress: "54.182.6.199",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.4",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.13",
	},
	{
		Domain:    "phdvasia.com",
		IpAddress: "13.35.6.26",
	},
	{
		Domain:    "twitchsvc.tech",
		IpAddress: "52.222.131.154",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.66",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.79",
	},
	{
		Domain:    "select.au.com",
		IpAddress: "13.249.2.51",
	},
	{
		Domain:    "avatax.avalara.net",
		IpAddress: "204.246.178.68",
	},
	{
		Domain:    "cdn.discounttire.com",
		IpAddress: "13.35.5.114",
	},
	{
		Domain:    "secb2b.com",
		IpAddress: "54.182.2.2",
	},
	{
		Domain:    "www.appservers.net",
		IpAddress: "52.222.132.232",
	},
	{
		Domain:    "datadoghq.com",
		IpAddress: "99.86.2.170",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "54.182.4.213",
	},
	{
		Domain:    "forhims.com",
		IpAddress: "54.182.6.101",
	},
	{
		Domain:    "jfrog.io",
		IpAddress: "54.182.4.87",
	},
	{
		Domain:    "zuora.identity.fcl-01.fcagcv.com",
		IpAddress: "52.222.131.216",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.8",
	},
	{
		Domain:    "www.c.ooyala.com",
		IpAddress: "13.224.0.227",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.193",
	},
	{
		Domain:    "dev1-www.lifelockunlocked.com",
		IpAddress: "99.84.2.124",
	},
	{
		Domain:    "file.samsungcloud.com",
		IpAddress: "143.204.6.146",
	},
	{
		Domain:    "gallery.mailchimp.com",
		IpAddress: "99.86.0.77",
	},
	{
		Domain:    "nexon.com",
		IpAddress: "54.239.130.51",
	},
	{
		Domain:    "perseus.de",
		IpAddress: "13.35.3.175",
	},
	{
		Domain:    "kaltura.com",
		IpAddress: "13.35.2.216",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.178",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.226",
	},
	{
		Domain:    "mfi-tc02.fnopf.jp",
		IpAddress: "99.86.2.187",
	},
	{
		Domain:    "bolindadigital.com",
		IpAddress: "54.182.6.177",
	},
	{
		Domain:    "es-navi.com",
		IpAddress: "143.204.2.32",
	},
	{
		Domain:    "sha-cf.v.uname.link",
		IpAddress: "99.84.6.86",
	},
	{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "99.86.2.36",
	},
	{
		Domain:    "wa.aws.amazon.com",
		IpAddress: "99.86.3.62",
	},
	{
		Domain:    "ad1.awsstatic.com",
		IpAddress: "204.246.169.161",
	},
	{
		Domain:    "aws.amazon.com",
		IpAddress: "99.84.4.72",
	},
	{
		Domain:    "www.brinkpos.net",
		IpAddress: "204.246.169.65",
	},
	{
		Domain:    "amazon.ca",
		IpAddress: "54.239.130.195",
	},
	{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "13.224.6.113",
	},
	{
		Domain:    "www.dazndn.com",
		IpAddress: "143.204.1.85",
	},
	{
		Domain:    "www.accordiagolf.com",
		IpAddress: "13.249.6.157",
	},
	{
		Domain:    "www.cphostaccess.com",
		IpAddress: "13.35.6.4",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.147",
	},
	{
		Domain:    "www.execute-api.us-east-1.amazonaws.com",
		IpAddress: "143.204.5.97",
	},
	{
		Domain:    "mpago.la",
		IpAddress: "13.35.6.206",
	},
	{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "13.35.5.196",
	},
	{
		Domain:    "www.playwithsea.com",
		IpAddress: "52.222.129.192",
	},
	{
		Domain:    "www.toukei-kentei.jp",
		IpAddress: "13.249.6.163",
	},
	{
		Domain:    "www.production.scrabble.withbuddies.com",
		IpAddress: "54.182.0.158",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.89",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.164",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.53",
	},
	{
		Domain:    "marketpulse.com",
		IpAddress: "143.204.1.165",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.194",
	},
	{
		Domain:    "sup-gcsp.jnj.com",
		IpAddress: "13.35.5.237",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.12",
	},
	{
		Domain:    "iot.us-east-1.amazonaws.com",
		IpAddress: "13.249.6.92",
	},
	{
		Domain:    "nba-cdn.2ksports.com",
		IpAddress: "52.222.131.180",
	},
	{
		Domain:    "carevisor.com",
		IpAddress: "54.182.6.225",
	},
	{
		Domain:    "enetscores.com",
		IpAddress: "204.246.178.44",
	},
	{
		Domain:    "www.bl.booklive.jp",
		IpAddress: "99.84.6.166",
	},
	{
		Domain:    "cdn-cloudfront.krxd.net",
		IpAddress: "99.84.2.49",
	},
	{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "54.182.0.160",
	},
	{
		Domain:    "smartica.jp",
		IpAddress: "13.35.6.129",
	},
	{
		Domain:    "api.digitalstudios.discovery.com",
		IpAddress: "204.246.164.165",
	},
	{
		Domain:    "as0.awsstatic.com",
		IpAddress: "99.86.2.23",
	},
	{
		Domain:    "zurple.com",
		IpAddress: "99.84.0.42",
	},
	{
		Domain:    "apps.bahrain.bh",
		IpAddress: "99.84.0.130",
	},
	{
		Domain:    "update.hicloud.com",
		IpAddress: "204.246.178.15",
	},
	{
		Domain:    "dl.amazon.com",
		IpAddress: "54.182.5.235",
	},
	{
		Domain:    "www.gph.imtxwy.com",
		IpAddress: "13.35.5.229",
	},
	{
		Domain:    "preprod.apac.amway.net",
		IpAddress: "143.204.6.208",
	},
	{
		Domain:    "origin-gql.beta.api.imdb.a2z.com",
		IpAddress: "99.86.0.142",
	},
	{
		Domain:    "specialized.com",
		IpAddress: "13.35.2.115",
	},
	{
		Domain:    "contestimg.wish.com",
		IpAddress: "204.246.164.57",
	},
	{
		Domain:    "siedev.net",
		IpAddress: "204.246.177.70",
	},
	{
		Domain:    "amazon.nl",
		IpAddress: "204.246.177.2",
	},
	{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "52.222.134.129",
	},
	{
		Domain:    "www.srv.ygles.com",
		IpAddress: "99.86.4.49",
	},
	{
		Domain:    "www.cafewell.com",
		IpAddress: "13.224.6.193",
	},
	{
		Domain:    "api.area-hinan-test.au.com",
		IpAddress: "54.239.192.54",
	},
	{
		Domain:    "www.shufu-job.jp",
		IpAddress: "13.224.0.226",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.20",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.17",
	},
	{
		Domain:    "static.amundi.com",
		IpAddress: "99.84.6.174",
	},
	{
		Domain:    "cf.pumlo.awsps.myinstance.com",
		IpAddress: "52.222.135.18",
	},
	{
		Domain:    "bks.cybird.ne.jp",
		IpAddress: "54.182.0.182",
	},
	{
		Domain:    "tvc-mall.com",
		IpAddress: "204.246.178.128",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.6.60",
	},
	{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "204.246.164.94",
	},
	{
		Domain:    "www.bounceexchange.com",
		IpAddress: "205.251.212.30",
	},
	{
		Domain:    "my.ellotte.com",
		IpAddress: "205.251.212.144",
	},
	{
		Domain:    "www.sodexomyway.com",
		IpAddress: "204.246.164.49",
	},
	{
		Domain:    "origin-client.legacy-app.games.a2z.com",
		IpAddress: "204.246.178.66",
	},
	{
		Domain:    "enetscores.com",
		IpAddress: "54.182.5.146",
	},
	{
		Domain:    "www.p7s1.io",
		IpAddress: "54.239.192.89",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.9",
	},
	{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "143.204.1.44",
	},
	{
		Domain:    "geocomply.com",
		IpAddress: "99.86.0.26",
	},
	{
		Domain:    "dev.ctrf.api.eden.mediba.jp",
		IpAddress: "13.35.4.53",
	},
	{
		Domain:    "www.chartboost.com",
		IpAddress: "54.182.2.144",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.2",
	},
	{
		Domain:    "dl.amazon.com",
		IpAddress: "54.182.0.235",
	},
	{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "52.222.129.84",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.19",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.196",
	},
	{
		Domain:    "www.swipesense.com",
		IpAddress: "13.249.5.128",
	},
	{
		Domain:    "www.indigoag.tech",
		IpAddress: "13.35.6.107",
	},
	{
		Domain:    "aiag.i-memo.jp",
		IpAddress: "99.84.6.68",
	},
	{
		Domain:    "www.linebc.jp",
		IpAddress: "54.239.130.178",
	},
	{
		Domain:    "avatax.avalara.net",
		IpAddress: "99.84.6.54",
	},
	{
		Domain:    "undercovertourist.com",
		IpAddress: "13.224.5.62",
	},
	{
		Domain:    "giv-dev.nmgcloud.io",
		IpAddress: "143.204.1.139",
	},
	{
		Domain:    "adventureacademy.com",
		IpAddress: "204.246.177.99",
	},
	{
		Domain:    "amazon.de",
		IpAddress: "13.35.5.149",
	},
	{
		Domain:    "kindle-guru.amazon.com",
		IpAddress: "204.246.164.124",
	},
	{
		Domain:    "www.innov8.space",
		IpAddress: "54.182.2.184",
	},
	{
		Domain:    "www.quick-cdn.com",
		IpAddress: "13.249.6.199",
	},
	{
		Domain:    "iot.ap-southeast-2.amazonaws.com",
		IpAddress: "54.182.3.218",
	},
	{
		Domain:    "amazon.co.uk",
		IpAddress: "54.239.192.101",
	},
	{
		Domain:    "video.counsyl.com",
		IpAddress: "205.251.212.161",
	},
	{
		Domain:    "origin-beta.client.legacy-app.games.a2z.com",
		IpAddress: "13.35.3.193",
	},
	{
		Domain:    "www.janrain.com",
		IpAddress: "99.86.5.158",
	},
	{
		Domain:    "www.i-ready.com",
		IpAddress: "52.222.130.187",
	},
	{
		Domain:    "api.msg.ue1.b.app.chime.aws",
		IpAddress: "13.249.6.161",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.5",
	},
	{
		Domain:    "livemeat.jp",
		IpAddress: "13.35.5.109",
	},
	{
		Domain:    "www.gph.imtxwy.com",
		IpAddress: "99.86.0.111",
	},
	{
		Domain:    "mpago.la",
		IpAddress: "13.249.2.25",
	},
	{
		Domain:    "mheducation.com",
		IpAddress: "54.182.2.33",
	},
	{
		Domain:    "musixmatch.com",
		IpAddress: "13.35.6.2",
	},
	{
		Domain:    "www.dreambox.com",
		IpAddress: "13.224.5.156",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.4",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.199",
	},
	{
		Domain:    "www.patient-create.orthofi-dev.com",
		IpAddress: "143.204.6.54",
	},
	{
		Domain:    "static.datadoghq.com",
		IpAddress: "54.182.6.111",
	},
	{
		Domain:    "www.update.easebar.com",
		IpAddress: "52.222.129.162",
	},
	{
		Domain:    "www.taggstar.com",
		IpAddress: "13.35.5.132",
	},
	{
		Domain:    "www.twitch.tv",
		IpAddress: "52.222.129.130",
	},
	{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "13.35.6.80",
	},
	{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "13.35.5.173",
	},
	{
		Domain:    "www.infomedia.com.au",
		IpAddress: "13.35.3.167",
	},
	{
		Domain:    "altium.com",
		IpAddress: "13.35.5.195",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.97",
	},
	{
		Domain:    "plaync.com",
		IpAddress: "54.239.192.140",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.25",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.44",
	},
	{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "52.222.128.156",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.73",
	},
	{
		Domain:    "nowforce.com",
		IpAddress: "54.182.4.186",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.100",
	},
	{
		Domain:    "www.placelocal.com",
		IpAddress: "13.35.0.178",
	},
	{
		Domain:    "zeasn.tv",
		IpAddress: "54.239.192.109",
	},
	{
		Domain:    "gateway.prod.compass.pioneer.com",
		IpAddress: "204.246.178.185",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.32",
	},
	{
		Domain:    "www.nyc837-dev.gin-dev.com",
		IpAddress: "143.204.1.92",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.99",
	},
	{
		Domain:    "forgecdn.net",
		IpAddress: "52.222.134.14",
	},
	{
		Domain:    "prod2.superobscuredomains.com",
		IpAddress: "52.222.133.252",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.205",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.3",
	},
	{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "52.222.130.150",
	},
	{
		Domain:    "www.dev.dgame.dmkt-sp.jp",
		IpAddress: "99.84.6.131",
	},
	{
		Domain:    "www.nmrodam.com",
		IpAddress: "13.35.2.209",
	},
	{
		Domain:    "assets.bwbx.io",
		IpAddress: "54.239.192.207",
	},
	{
		Domain:    "www.freshdesk.com",
		IpAddress: "99.84.2.42",
	},
	{
		Domain:    "macmillanyounglearners.com",
		IpAddress: "13.224.7.18",
	},
	{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "13.224.0.167",
	},
	{
		Domain:    "movergames.com",
		IpAddress: "204.246.177.156",
	},
	{
		Domain:    "chime.aws",
		IpAddress: "99.86.6.11",
	},
	{
		Domain:    "sings-download.twitch.tv",
		IpAddress: "143.204.2.9",
	},
	{
		Domain:    "club.ubisoft.com",
		IpAddress: "13.35.3.170",
	},
	{
		Domain:    "adn.wyzant.com",
		IpAddress: "13.35.5.184",
	},
	{
		Domain:    "www.ashcream.xyz",
		IpAddress: "13.224.5.212",
	},
	{
		Domain:    "www.awsapps.com",
		IpAddress: "52.222.132.94",
	},
	{
		Domain:    "api.shopbop.com",
		IpAddress: "52.222.135.11",
	},
	{
		Domain:    "saucelabs.com",
		IpAddress: "99.84.6.210",
	},
	{
		Domain:    "mkw.melbourne.vic.gov.au",
		IpAddress: "143.204.2.225",
	},
	{
		Domain:    "bethesda.net",
		IpAddress: "143.204.2.219",
	},
	{
		Domain:    "api.stg.smartpass.auone.jp",
		IpAddress: "99.84.2.222",
	},
	{
		Domain:    "lovewall-missdior.dior.com",
		IpAddress: "54.182.2.148",
	},
	{
		Domain:    "s-onetag.com",
		IpAddress: "52.222.134.188",
	},
	{
		Domain:    "ap1.whispir.com",
		IpAddress: "54.239.192.142",
	},
	{
		Domain:    "mojang.com",
		IpAddress: "54.182.2.20",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.187",
	},
	{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "205.251.212.184",
	},
	{
		Domain:    "appgallery.huawei.com",
		IpAddress: "52.222.134.107",
	},
	{
		Domain:    "twitchcdn-shadow.net",
		IpAddress: "13.249.6.151",
	},
	{
		Domain:    "www.period-calendar.com",
		IpAddress: "205.251.212.239",
	},
	{
		Domain:    "www.hungama.com",
		IpAddress: "52.222.129.70",
	},
	{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "54.239.130.229",
	},
	{
		Domain:    "api.sandbox.repayonline.com",
		IpAddress: "99.86.2.87",
	},
	{
		Domain:    "siftscience.com",
		IpAddress: "13.249.2.64",
	},
	{
		Domain:    "ssi.servicestream.com.au",
		IpAddress: "143.204.6.199",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.24",
	},
	{
		Domain:    "smile.amazon.de",
		IpAddress: "13.224.5.77",
	},
	{
		Domain:    "internal-weedmaps.com",
		IpAddress: "54.182.2.85",
	},
	{
		Domain:    "cdn.discounttire.com",
		IpAddress: "99.86.0.211",
	},
	{
		Domain:    "siftscience.com",
		IpAddress: "99.86.1.124",
	},
	{
		Domain:    "cdn-legacy.contentful.com",
		IpAddress: "13.35.5.14",
	},
	{
		Domain:    "bamsec.com",
		IpAddress: "13.249.5.44",
	},
	{
		Domain:    "pay.2go.com",
		IpAddress: "205.251.212.16",
	},
	{
		Domain:    "www.connectwise.com",
		IpAddress: "52.222.132.120",
	},
	{
		Domain:    "s.salecycle.com",
		IpAddress: "52.222.135.8",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.89",
	},
	{
		Domain:    "www.gr-assets.com",
		IpAddress: "99.86.2.8",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.28",
	},
	{
		Domain:    "js-assets.aiv-cdn.net",
		IpAddress: "204.246.164.44",
	},
	{
		Domain:    "www.misumi.jp",
		IpAddress: "54.182.6.120",
	},
	{
		Domain:    "xgcpaa.com",
		IpAddress: "13.249.5.22",
	},
	{
		Domain:    "iproc.originenergy.com.au",
		IpAddress: "99.86.6.189",
	},
	{
		Domain:    "rca-upload-cloudstation-eu-central-1.dev.hydra.sophos.com",
		IpAddress: "204.246.178.167",
	},
	{
		Domain:    "public-rca-cloudstation-us-east-2.qa.hydra.sophos.com",
		IpAddress: "54.182.3.60",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.157",
	},
	{
		Domain:    "www.creditloan.com",
		IpAddress: "204.246.178.6",
	},
	{
		Domain:    "www.fp.ps.easebar.com",
		IpAddress: "143.204.6.213",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.192",
	},
	{
		Domain:    "oqs.amb.cybird.ne.jp",
		IpAddress: "52.222.131.144",
	},
	{
		Domain:    "static.amundi.com",
		IpAddress: "99.86.6.33",
	},
	{
		Domain:    "achievers.com",
		IpAddress: "13.35.6.128",
	},
	{
		Domain:    "www.sigalert.com",
		IpAddress: "13.224.7.5",
	},
	{
		Domain:    "static.yub-cdn.com",
		IpAddress: "13.249.2.12",
	},
	{
		Domain:    "boleto.sandbox.pagseguro.com.br",
		IpAddress: "54.182.2.149",
	},
	{
		Domain:    "gaijinent.com",
		IpAddress: "13.249.5.21",
	},
	{
		Domain:    "edwardsdoc.com",
		IpAddress: "13.35.4.166",
	},
	{
		Domain:    "www.goldspotmedia.com",
		IpAddress: "13.35.3.52",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.206",
	},
	{
		Domain:    "www.cookpad.com",
		IpAddress: "204.246.177.11",
	},
	{
		Domain:    "iot.eu-west-2.amazonaws.com",
		IpAddress: "99.86.2.91",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.105",
	},
	{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "13.224.5.27",
	},
	{
		Domain:    "www.canadamats.ca",
		IpAddress: "204.246.164.8",
	},
	{
		Domain:    "rca-upload-cloudstation-us-west-2.dev3.hydra.sophos.com",
		IpAddress: "204.246.177.29",
	},
	{
		Domain:    "polaris.lhinside.com",
		IpAddress: "143.204.6.108",
	},
	{
		Domain:    "behance.net",
		IpAddress: "13.35.6.71",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.201",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.2",
	},
	{
		Domain:    "uploads.skyhighnetworks.com",
		IpAddress: "13.224.5.118",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.61",
	},
	{
		Domain:    "www.milkvr.rocks",
		IpAddress: "99.86.4.34",
	},
	{
		Domain:    "freight.amazon.com",
		IpAddress: "99.86.2.12",
	},
	{
		Domain:    "binance.com",
		IpAddress: "13.35.3.114",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.8",
	},
	{
		Domain:    "www.tosconfig.com",
		IpAddress: "99.86.0.192",
	},
	{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.103",
	},
	{
		Domain:    "iot.ap-northeast-1.amazonaws.com",
		IpAddress: "13.224.5.194",
	},
	{
		Domain:    "mobizen.com",
		IpAddress: "143.204.1.15",
	},
	{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "143.204.2.210",
	},
	{
		Domain:    "www.iot.irobot.cn",
		IpAddress: "99.86.2.14",
	},
	{
		Domain:    "bglen.net",
		IpAddress: "13.224.5.60",
	},
	{
		Domain:    "read.amazon.com",
		IpAddress: "99.86.2.62",
	},
	{
		Domain:    "www.nmrodam.com",
		IpAddress: "54.239.192.179",
	},
	{
		Domain:    "netmarble.net",
		IpAddress: "54.182.0.164",
	},
}
