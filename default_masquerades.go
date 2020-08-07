package fronted

var DefaultTrustedCAs = []*CA{
	&CA{
		CommonName: "Amazon Root CA 1",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDQTCCAimgAwIBAgITBmyfz5m/jAo54vB4ikPmljZbyjANBgkqhkiG9w0BAQsF\nADA5MQswCQYDVQQGEwJVUzEPMA0GA1UEChMGQW1hem9uMRkwFwYDVQQDExBBbWF6\nb24gUm9vdCBDQSAxMB4XDTE1MDUyNjAwMDAwMFoXDTM4MDExNzAwMDAwMFowOTEL\nMAkGA1UEBhMCVVMxDzANBgNVBAoTBkFtYXpvbjEZMBcGA1UEAxMQQW1hem9uIFJv\nb3QgQ0EgMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALJ4gHHKeNXj\nca9HgFB0fW7Y14h29Jlo91ghYPl0hAEvrAIthtOgQ3pOsqTQNroBvo3bSMgHFzZM\n9O6II8c+6zf1tRn4SWiw3te5djgdYZ6k/oI2peVKVuRF4fn9tBb6dNqcmzU5L/qw\nIFAGbHrQgLKm+a/sRxmPUDgH3KKHOVj4utWp+UhnMJbulHheb4mjUcAwhmahRWa6\nVOujw5H5SNz/0egwLX0tdHA114gk957EWW67c4cX8jJGKLhD+rcdqsq08p8kDi1L\n93FcXmn/6pUCyziKrlA4b9v7LWIbxcceVOF34GfID5yHI9Y/QCB/IIDEgEw+OyQm\njgSubJrIqg0CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMC\nAYYwHQYDVR0OBBYEFIQYzIU07LwMlJQuCFmcx7IQTgoIMA0GCSqGSIb3DQEBCwUA\nA4IBAQCY8jdaQZChGsV2USggNiMOruYou6r4lK5IpDB/G/wkjUu0yKGX9rbxenDI\nU5PMCCjjmCXPI6T53iHTfIUJrU6adTrCC2qJeHZERxhlbI1Bjjt/msv0tadQ1wUs\nN+gDS63pYaACbvXy8MWy7Vu33PqUXHeeE6V/Uq2V8viTO96LXFvKWlJbYK8U90vv\no/ufQJVtMVT8QtPHRh8jrdkPSHCa2XV4cdFyQzR1bldZwgJcJmApzyMZFo6IQ6XU\n5MsI+yMRQ+hDKXJioaldXgjUkK642M4UwtBV8ob2xJNDd2ZhwLnoQdeXeGADbkpy\nrqXRfboQnoZsG4q5WTP468SQvvG5\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "DigiCert Global Root G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH\nMjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI\n2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx\n1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ\nq2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz\ntCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ\nvIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP\nBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV\n5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY\n1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4\nNeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG\nFdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91\n8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe\npLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl\nMrY=\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "DigiCert Global Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\nQTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB\nCSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97\nnh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt\n43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P\nT19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4\ngdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO\nBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR\nTLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw\nDQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr\nhMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg\n06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF\nPnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls\nYSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk\nCAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "Go Daddy Root Certificate Authority - G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIBADANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMCVVMx\nEDAOBgNVBAgTB0FyaXpvbmExEzARBgNVBAcTClNjb3R0c2RhbGUxGjAYBgNVBAoT\nEUdvRGFkZHkuY29tLCBJbmMuMTEwLwYDVQQDEyhHbyBEYWRkeSBSb290IENlcnRp\nZmljYXRlIEF1dGhvcml0eSAtIEcyMB4XDTA5MDkwMTAwMDAwMFoXDTM3MTIzMTIz\nNTk1OVowgYMxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdBcml6b25hMRMwEQYDVQQH\nEwpTY290dHNkYWxlMRowGAYDVQQKExFHb0RhZGR5LmNvbSwgSW5jLjExMC8GA1UE\nAxMoR28gRGFkZHkgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgLSBHMjCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL9xYgjx+lk09xvJGKP3gElY6SKD\nE6bFIEMBO4Tx5oVJnyfq9oQbTqC023CYxzIBsQU+B07u9PpPL1kwIuerGVZr4oAH\n/PMWdYA5UXvl+TW2dE6pjYIT5LY/qQOD+qK+ihVqf94Lw7YZFAXK6sOoBJQ7Rnwy\nDfMAZiLIjWltNowRGLfTshxgtDj6AozO091GB94KPutdfMh8+7ArU6SSYmlRJQVh\nGkSBjCypQ5Yj36w6gZoOKcUcqeldHraenjAKOc7xiID7S13MMuyFYkMlNAJWJwGR\ntDtwKj9useiciAF9n9T521NtYJ2/LOdYq7hfRvzOxBsDPAnrSTFcaUaz4EcCAwEA\nAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYE\nFDqahQcQZyi27/a9BUFuIMGU2g/eMA0GCSqGSIb3DQEBCwUAA4IBAQCZ21151fmX\nWWcDYfF+OwYxdS2hII5PZYe096acvNjpL9DbWu7PdIxztDhC2gV7+AJ1uP2lsdeu\n9tfeE8tTEH6KRtGX+rcuKxGrkLAngPnon1rpN5+r5N9ss4UXnT3ZJE95kTXWXwTr\ngIOrmgIttRD02JDHBHNA7XIloKmf7J6raBKZV8aPEjoJpL1E/QYVN8Gb5DKj7Tjo\n2GTzLH4U/ALqn83/B2gX2yKQOC16jdFU8WnjXzPKej17CuPKf1855eJ1usV2GDPO\nLPAvTK33sefOT6jEm0pUBsV/fdUID+Ic/n4XuKxe9tQWskMJDE32p2u0mYRlynqI\n4uJEvlz36hz1\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "USERTrust RSA Certification Authority",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIF3jCCA8agAwIBAgIQAf1tMPyjylGoG7xkDjUDLTANBgkqhkiG9w0BAQwFADCB\niDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0pl\ncnNleSBDaXR5MR4wHAYDVQQKExVUaGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNV\nBAMTJVVTRVJUcnVzdCBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTAw\nMjAxMDAwMDAwWhcNMzgwMTE4MjM1OTU5WjCBiDELMAkGA1UEBhMCVVMxEzARBgNV\nBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0plcnNleSBDaXR5MR4wHAYDVQQKExVU\naGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNVBAMTJVVTRVJUcnVzdCBSU0EgQ2Vy\ndGlmaWNhdGlvbiBBdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK\nAoICAQCAEmUXNg7D2wiz0KxXDXbtzSfTTK1Qg2HiqiBNCS1kCdzOiZ/MPans9s/B\n3PHTsdZ7NygRK0faOca8Ohm0X6a9fZ2jY0K2dvKpOyuR+OJv0OwWIJAJPuLodMkY\ntJHUYmTbf6MG8YgYapAiPLz+E/CHFHv25B+O1ORRxhFnRghRy4YUVD+8M/5+bJz/\nFp0YvVGONaanZshyZ9shZrHUm3gDwFA66Mzw3LyeTP6vBZY1H1dat//O+T23LLb2\nVN3I5xI6Ta5MirdcmrS3ID3KfyI0rn47aGYBROcBTkZTmzNg95S+UzeQc0PzMsNT\n79uq/nROacdrjGCT3sTHDN/hMq7MkztReJVni+49Vv4M0GkPGw/zJSZrM233bkf6\nc0Plfg6lZrEpfDKEY1WJxA3Bk1QwGROs0303p+tdOmw1XNtB1xLaqUkL39iAigmT\nYo61Zs8liM2EuLE/pDkP2QKe6xJMlXzzawWpXhaDzLhn4ugTncxbgtNMs+1b/97l\nc6wjOy0AvzVVdAlJ2ElYGn+SNuZRkg7zJn0cTRe8yexDJtC/QV9AqURE9JnnV4ee\nUB9XVKg+/XRjL7FQZQnmWEIuQxpMtPAlR1n6BB6T1CZGSlCBst6+eLf8ZxXhyVeE\nHg9j1uliutZfVS7qXMYoCAQlObgOK6nyTJccBz8NUvXt7y+CDwIDAQABo0IwQDAd\nBgNVHQ4EFgQUU3m/WqorSs9UgOHYm8Cd8rIDZsswDgYDVR0PAQH/BAQDAgEGMA8G\nA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQEMBQADggIBAFzUfA3P9wF9QZllDHPF\nUp/L+M+ZBn8b2kMVn54CVVeWFPFSPCeHlCjtHzoBN6J2/FNQwISbxmtOuowhT6KO\nVWKR82kV2LyI48SqC/3vqOlLVSoGIG1VeCkZ7l8wXEskEVX/JJpuXior7gtNn3/3\nATiUFJVDBwn7YKnuHKsSjKCaXqeYalltiz8I+8jRRa8YFWSQEg9zKC7F4iRO/Fjs\n8PRF/iKz6y+O0tlFYQXBl2+odnKPi4w2r78NBc5xjeambx9spnFixdjQg3IM8WcR\niQycE0xyNN+81XHfqnHd4blsjDwSXWXavVcStkNr/+XeTWYRUc+ZruwXtuhxkYze\nSf7dNXGiFSeUHM9h4ya7b6NnJSFd5t0dCy5oGzuCr+yDZ4XUmFF0sbmZgIn/f3gZ\nXHlKYC6SQK5MNyosycdiyA5d9zZbyuAlJQG03RoHnHcAP9Dc1ew91Pq7P8yF1m9/\nqS3fuQL39ZeatTXaw2ewh0qpKJ4jjv9cJ2vhsE/zB+4ALtRZh8tSQZXq9EfX7mRB\nVXyNWQKV3WKdwrnuWih0hKWbt5DHDAff9Yk2dDLWKMGwsAvgnEzDHNb842m1R0aB\nL6KCq9NjRHDEjf8tM7qtj3u1cIiuPhnPQCjY/MiQu12ZIvVS5ljFH4gxQ+6IHdfG\njjxDah2nGN59PRbxYvnKkKj9\n-----END CERTIFICATE-----\n",
	},
}

var DefaultCloudfrontMasquerades = []*Masquerade{
	&Masquerade{
		Domain:    "www.amazon.ae",
		IpAddress: "13.224.6.43",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.27",
	},
	&Masquerade{
		Domain:    "assetserv.com",
		IpAddress: "13.224.6.236",
	},
	&Masquerade{
		Domain:    "datad0g.com",
		IpAddress: "204.246.177.89",
	},
	&Masquerade{
		Domain:    "t.mail.optimumemail1.com",
		IpAddress: "13.224.7.64",
	},
	&Masquerade{
		Domain:    "realisticgames.co.uk",
		IpAddress: "13.224.5.3",
	},
	&Masquerade{
		Domain:    "geocomply.com",
		IpAddress: "204.246.177.39",
	},
	&Masquerade{
		Domain:    "dev.twitch.tv",
		IpAddress: "143.204.1.60",
	},
	&Masquerade{
		Domain:    "dolphin-fe.amazon.com",
		IpAddress: "204.246.178.22",
	},
	&Masquerade{
		Domain:    "brightcove.com",
		IpAddress: "143.204.1.30",
	},
	&Masquerade{
		Domain:    "assets.bwbx.io",
		IpAddress: "204.246.178.184",
	},
	&Masquerade{
		Domain:    "www.dcm-icwweb-dev.com",
		IpAddress: "204.246.178.19",
	},
	&Masquerade{
		Domain:    "www.adbephotos-stage.com",
		IpAddress: "99.86.2.9",
	},
	&Masquerade{
		Domain:    "www.connectwisedev.com",
		IpAddress: "99.86.0.153",
	},
	&Masquerade{
		Domain:    "www.uat.catchplay.com",
		IpAddress: "99.86.0.178",
	},
	&Masquerade{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "205.251.212.47",
	},
	&Masquerade{
		Domain:    "cdn.admin.staging.checkmatenext.com",
		IpAddress: "99.86.3.179",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.227",
	},
	&Masquerade{
		Domain:    "www.amazon.ae",
		IpAddress: "99.84.2.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.168",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "99.84.2.162",
	},
	&Masquerade{
		Domain:    "rheemcert.com",
		IpAddress: "99.86.4.107",
	},
	&Masquerade{
		Domain:    "samsungqbe.com",
		IpAddress: "99.84.0.159",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.32",
	},
	&Masquerade{
		Domain:    "api.smartpass.auone.jp",
		IpAddress: "54.239.130.234",
	},
	&Masquerade{
		Domain:    "undercovertourist.com",
		IpAddress: "13.249.7.66",
	},
	&Masquerade{
		Domain:    "emergency.wa.gov.au",
		IpAddress: "13.249.6.72",
	},
	&Masquerade{
		Domain:    "cdn.hands.net",
		IpAddress: "99.86.6.135",
	},
	&Masquerade{
		Domain:    "sellercentral.amazon.com",
		IpAddress: "143.204.2.163",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.36",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.26",
	},
	&Masquerade{
		Domain:    "www.samsungsmartcam.com",
		IpAddress: "13.249.6.29",
	},
	&Masquerade{
		Domain:    "phdvasia.com",
		IpAddress: "52.222.129.149",
	},
	&Masquerade{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "52.222.134.184",
	},
	&Masquerade{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "52.222.128.161",
	},
	&Masquerade{
		Domain:    "soccerladuma.co.za",
		IpAddress: "52.222.129.150",
	},
	&Masquerade{
		Domain:    "shopch.jp",
		IpAddress: "13.35.3.101",
	},
	&Masquerade{
		Domain:    "www.animelo.jp",
		IpAddress: "143.204.1.59",
	},
	&Masquerade{
		Domain:    "www.netdespatch.com",
		IpAddress: "13.35.4.228",
	},
	&Masquerade{
		Domain:    "trusteerqa.com",
		IpAddress: "13.35.4.21",
	},
	&Masquerade{
		Domain:    "gluon-cv.mxnet.io",
		IpAddress: "13.35.1.114",
	},
	&Masquerade{
		Domain:    "www.smentertainment.com",
		IpAddress: "204.246.177.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.166",
	},
	&Masquerade{
		Domain:    "www.api.brightcove.com",
		IpAddress: "204.246.164.127",
	},
	&Masquerade{
		Domain:    "bd1.awsstatic.com",
		IpAddress: "204.246.164.185",
	},
	&Masquerade{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "99.86.0.222",
	},
	&Masquerade{
		Domain:    "as0.awsstatic.com",
		IpAddress: "204.246.178.33",
	},
	&Masquerade{
		Domain:    "isao.net",
		IpAddress: "54.239.192.166",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-us-east-2.qa.hydra.sophos.com",
		IpAddress: "143.204.7.42",
	},
	&Masquerade{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "52.222.132.11",
	},
	&Masquerade{
		Domain:    "unrealengine.com",
		IpAddress: "99.86.4.156",
	},
	&Masquerade{
		Domain:    "pimg.jp",
		IpAddress: "99.86.2.27",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.178",
	},
	&Masquerade{
		Domain:    "webspectator.com",
		IpAddress: "204.246.169.132",
	},
	&Masquerade{
		Domain:    "kaltura.com",
		IpAddress: "99.86.5.174",
	},
	&Masquerade{
		Domain:    "jwo.amazon.com",
		IpAddress: "143.204.1.26",
	},
	&Masquerade{
		Domain:    "z-na.amazon-adsystem.com",
		IpAddress: "54.182.6.41",
	},
	&Masquerade{
		Domain:    "kucoin.com",
		IpAddress: "99.84.7.32",
	},
	&Masquerade{
		Domain:    "supplychainconnect.amazon.com",
		IpAddress: "204.246.169.15",
	},
	&Masquerade{
		Domain:    "product-downloads.atlassian.com",
		IpAddress: "54.182.6.185",
	},
	&Masquerade{
		Domain:    "smtown.com",
		IpAddress: "13.224.5.137",
	},
	&Masquerade{
		Domain:    "www.vistarmedia.com",
		IpAddress: "204.246.169.96",
	},
	&Masquerade{
		Domain:    "www.wowma.jp",
		IpAddress: "99.86.1.78",
	},
	&Masquerade{
		Domain:    "wpcp.shiseido.co.jp",
		IpAddress: "216.137.39.7",
	},
	&Masquerade{
		Domain:    "www.fastretailing.com",
		IpAddress: "99.86.2.108",
	},
	&Masquerade{
		Domain:    "altium.com",
		IpAddress: "99.86.0.186",
	},
	&Masquerade{
		Domain:    "eprocurement.marketplace.us-east-1.amazonaws.com",
		IpAddress: "99.86.2.105",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.132",
	},
	&Masquerade{
		Domain:    "www7.amazon.com",
		IpAddress: "13.249.2.67",
	},
	&Masquerade{
		Domain:    "i-parcel.com",
		IpAddress: "13.224.6.188",
	},
	&Masquerade{
		Domain:    "searchandexplore.com",
		IpAddress: "13.224.6.168",
	},
	&Masquerade{
		Domain:    "oih-gamma-eu.aka.amazon.com",
		IpAddress: "13.224.6.108",
	},
	&Masquerade{
		Domain:    "slackfrontiers.com",
		IpAddress: "99.84.6.196",
	},
	&Masquerade{
		Domain:    "panda.chtbl.com",
		IpAddress: "52.222.135.66",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-west-1.qa.hydra.sophos.com",
		IpAddress: "52.222.134.162",
	},
	&Masquerade{
		Domain:    "we-stats.com",
		IpAddress: "54.182.6.113",
	},
	&Masquerade{
		Domain:    "www.c.ooyala.com",
		IpAddress: "13.249.5.199",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.224.6.97",
	},
	&Masquerade{
		Domain:    "mobizen.com",
		IpAddress: "99.86.2.50",
	},
	&Masquerade{
		Domain:    "www.gdl.imtxwy.com",
		IpAddress: "99.86.5.155",
	},
	&Masquerade{
		Domain:    "d-hrp.com",
		IpAddress: "13.249.2.238",
	},
	&Masquerade{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "13.35.6.59",
	},
	&Masquerade{
		Domain:    "images-cn.ssl-images-amazon.com",
		IpAddress: "52.222.132.50",
	},
	&Masquerade{
		Domain:    "cf.test.frontier.a2z.com",
		IpAddress: "52.222.132.117",
	},
	&Masquerade{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "54.182.2.161",
	},
	&Masquerade{
		Domain:    "mheducation.com",
		IpAddress: "99.84.2.133",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.197",
	},
	&Masquerade{
		Domain:    "www.awscfdns.com",
		IpAddress: "13.35.2.242",
	},
	&Masquerade{
		Domain:    "www.diageohorizon.com",
		IpAddress: "99.84.2.203",
	},
	&Masquerade{
		Domain:    "fujifilmimagine.com",
		IpAddress: "99.86.0.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.94",
	},
	&Masquerade{
		Domain:    "gbf.game-a.mbga.jp",
		IpAddress: "13.224.0.132",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.107",
	},
	&Masquerade{
		Domain:    "www.53.localytics.com",
		IpAddress: "13.249.2.62",
	},
	&Masquerade{
		Domain:    "load-test6.eu-west-2.cf-embed.net",
		IpAddress: "54.239.130.82",
	},
	&Masquerade{
		Domain:    "www.nosto.com",
		IpAddress: "13.249.6.10",
	},
	&Masquerade{
		Domain:    "oasiscdn.com",
		IpAddress: "204.246.177.64",
	},
	&Masquerade{
		Domain:    "gaijinent.com",
		IpAddress: "99.86.0.103",
	},
	&Masquerade{
		Domain:    "media.preziusercontent.com",
		IpAddress: "13.224.6.207",
	},
	&Masquerade{
		Domain:    "buildinglink.com",
		IpAddress: "143.204.5.63",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.219",
	},
	&Masquerade{
		Domain:    "www.observian.com",
		IpAddress: "52.222.134.2",
	},
	&Masquerade{
		Domain:    "www.amazon.sa",
		IpAddress: "54.182.2.233",
	},
	&Masquerade{
		Domain:    "ebookstore.sony.jp",
		IpAddress: "216.137.39.53",
	},
	&Masquerade{
		Domain:    "amazon.com.au",
		IpAddress: "99.84.2.84",
	},
	&Masquerade{
		Domain:    "alexa-comms-mobile-service.amazon.com",
		IpAddress: "13.224.0.182",
	},
	&Masquerade{
		Domain:    "hkcp08.com",
		IpAddress: "99.86.1.88",
	},
	&Masquerade{
		Domain:    "api.area-hinan-test.au.com",
		IpAddress: "204.246.169.51",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-central-1.qa.hydra.sophos.com",
		IpAddress: "13.35.6.175",
	},
	&Masquerade{
		Domain:    "mapbox.cn",
		IpAddress: "13.35.4.151",
	},
	&Masquerade{
		Domain:    "www.indigoag.tech",
		IpAddress: "204.246.164.179",
	},
	&Masquerade{
		Domain:    "www.production.scrabble.withbuddies.com",
		IpAddress: "99.86.3.120",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.23",
	},
	&Masquerade{
		Domain:    "www.bcovlive.io",
		IpAddress: "99.86.1.216",
	},
	&Masquerade{
		Domain:    "media.aircorsica.com",
		IpAddress: "13.35.1.153",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.129",
	},
	&Masquerade{
		Domain:    "www.dev.aws.casualty.cccis.com",
		IpAddress: "13.35.5.145",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.132",
	},
	&Masquerade{
		Domain:    "www.tigocloud.net",
		IpAddress: "52.222.130.214",
	},
	&Masquerade{
		Domain:    "tenki.auone.jp",
		IpAddress: "13.35.5.68",
	},
	&Masquerade{
		Domain:    "www.thinknearhub.com",
		IpAddress: "13.35.6.112",
	},
	&Masquerade{
		Domain:    "unrulymedia.com",
		IpAddress: "54.182.0.208",
	},
	&Masquerade{
		Domain:    "craftsy.com",
		IpAddress: "54.182.2.209",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.213",
	},
	&Masquerade{
		Domain:    "collectivehealth.com",
		IpAddress: "52.222.131.128",
	},
	&Masquerade{
		Domain:    "www.execute-api.us-east-1.amazonaws.com",
		IpAddress: "99.86.6.34",
	},
	&Masquerade{
		Domain:    "ext-test.app-cloud.jp",
		IpAddress: "143.204.6.27",
	},
	&Masquerade{
		Domain:    "aloseguro.com",
		IpAddress: "52.222.129.160",
	},
	&Masquerade{
		Domain:    "mpago.la",
		IpAddress: "143.204.2.90",
	},
	&Masquerade{
		Domain:    "api.stg.smartpass.auone.jp",
		IpAddress: "99.86.1.57",
	},
	&Masquerade{
		Domain:    "thetvdb.com",
		IpAddress: "54.182.3.230",
	},
	&Masquerade{
		Domain:    "esd.sentinelcloud.com",
		IpAddress: "204.246.178.124",
	},
	&Masquerade{
		Domain:    "www.project-a.videoprojects.net",
		IpAddress: "54.182.0.149",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.5.67",
	},
	&Masquerade{
		Domain:    "deploygate.com",
		IpAddress: "13.35.6.226",
	},
	&Masquerade{
		Domain:    "www.freshdesk.com",
		IpAddress: "99.86.5.24",
	},
	&Masquerade{
		Domain:    "ewrzfr.com",
		IpAddress: "143.204.0.165",
	},
	&Masquerade{
		Domain:    "www.stg.forecast.elyza.ai",
		IpAddress: "143.204.6.5",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.102",
	},
	&Masquerade{
		Domain:    "www.audible.com.au",
		IpAddress: "99.86.0.213",
	},
	&Masquerade{
		Domain:    "tripkit-test4.jeppesen.com",
		IpAddress: "13.249.6.140",
	},
	&Masquerade{
		Domain:    "cdn-legacy.contentful.com",
		IpAddress: "54.182.3.24",
	},
	&Masquerade{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "99.84.0.114",
	},
	&Masquerade{
		Domain:    "poptropica.com",
		IpAddress: "54.182.2.44",
	},
	&Masquerade{
		Domain:    "as0.awsstatic.com",
		IpAddress: "143.204.5.23",
	},
	&Masquerade{
		Domain:    "api1.platformdxc-d2.com",
		IpAddress: "13.35.4.200",
	},
	&Masquerade{
		Domain:    "smallpdf.com",
		IpAddress: "54.182.4.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.229",
	},
	&Masquerade{
		Domain:    "geocomply.net",
		IpAddress: "143.204.2.82",
	},
	&Masquerade{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "13.249.5.29",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.42",
	},
	&Masquerade{
		Domain:    "blim.com",
		IpAddress: "13.35.2.206",
	},
	&Masquerade{
		Domain:    "www.realizedev-test.com",
		IpAddress: "52.222.132.130",
	},
	&Masquerade{
		Domain:    "amazonsmile.com",
		IpAddress: "204.246.177.77",
	},
	&Masquerade{
		Domain:    "www.audible.com.au",
		IpAddress: "99.84.2.38",
	},
	&Masquerade{
		Domain:    "lucidhq.com",
		IpAddress: "99.86.6.52",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.42",
	},
	&Masquerade{
		Domain:    "www.53.localytics.com",
		IpAddress: "204.246.164.60",
	},
	&Masquerade{
		Domain:    "ekdgd.com",
		IpAddress: "143.204.5.126",
	},
	&Masquerade{
		Domain:    "twitchsvc.net",
		IpAddress: "204.246.169.158",
	},
	&Masquerade{
		Domain:    "mheducation.com",
		IpAddress: "52.222.134.79",
	},
	&Masquerade{
		Domain:    "zeasn.tv",
		IpAddress: "204.246.169.108",
	},
	&Masquerade{
		Domain:    "www.dev.pos.paylabo.com",
		IpAddress: "99.86.1.54",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.168",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.33",
	},
	&Masquerade{
		Domain:    "int3.machieco.nestle.jp",
		IpAddress: "52.222.135.49",
	},
	&Masquerade{
		Domain:    "download.epicgames.com",
		IpAddress: "54.239.192.66",
	},
	&Masquerade{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "99.84.6.108",
	},
	&Masquerade{
		Domain:    "a1v.starfall.com",
		IpAddress: "204.246.177.46",
	},
	&Masquerade{
		Domain:    "democrats.org",
		IpAddress: "204.246.178.41",
	},
	&Masquerade{
		Domain:    "www7.amazon.com",
		IpAddress: "54.239.130.171",
	},
	&Masquerade{
		Domain:    "www.tipico.com",
		IpAddress: "204.246.177.191",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.22",
	},
	&Masquerade{
		Domain:    "datadoghq.com",
		IpAddress: "13.249.5.87",
	},
	&Masquerade{
		Domain:    "demandbase.com",
		IpAddress: "143.204.5.108",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.7",
	},
	&Masquerade{
		Domain:    "www.staging.truecardev.com",
		IpAddress: "13.35.0.176",
	},
	&Masquerade{
		Domain:    "cf.test.frontier.a2z.com",
		IpAddress: "143.204.2.34",
	},
	&Masquerade{
		Domain:    "sftelemetry.sophos.com",
		IpAddress: "143.204.1.86",
	},
	&Masquerade{
		Domain:    "brain-market.com",
		IpAddress: "13.35.3.22",
	},
	&Masquerade{
		Domain:    "oneblood.org",
		IpAddress: "13.249.2.139",
	},
	&Masquerade{
		Domain:    "www.predix.io",
		IpAddress: "143.204.2.217",
	},
	&Masquerade{
		Domain:    "club-beta2.pokemon.com",
		IpAddress: "204.246.164.166",
	},
	&Masquerade{
		Domain:    "inspector-agent.amazonaws.com",
		IpAddress: "13.35.1.152",
	},
	&Masquerade{
		Domain:    "www.studysapuri.jp",
		IpAddress: "204.246.164.59",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.173",
	},
	&Masquerade{
		Domain:    "bc-citi.providersml.com",
		IpAddress: "13.35.2.155",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.193",
	},
	&Masquerade{
		Domain:    "mobile.mercadopago.com",
		IpAddress: "99.86.1.210",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "99.86.5.57",
	},
	&Masquerade{
		Domain:    "test.samsunghealth.com",
		IpAddress: "54.239.192.172",
	},
	&Masquerade{
		Domain:    "www.realizedev-test.com",
		IpAddress: "13.224.6.183",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.111",
	},
	&Masquerade{
		Domain:    "jwpsrv.com",
		IpAddress: "52.222.132.182",
	},
	&Masquerade{
		Domain:    "www.cequintsptecid.com",
		IpAddress: "54.182.4.216",
	},
	&Masquerade{
		Domain:    "mark1.dev",
		IpAddress: "204.246.178.88",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "52.222.131.118",
	},
	&Masquerade{
		Domain:    "code.org",
		IpAddress: "99.86.6.227",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.16",
	},
	&Masquerade{
		Domain:    "knowledgevision.com",
		IpAddress: "13.249.6.171",
	},
	&Masquerade{
		Domain:    "www.dwell.com",
		IpAddress: "13.35.5.126",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.124",
	},
	&Masquerade{
		Domain:    "fe.dazn-stage.com",
		IpAddress: "54.239.130.147",
	},
	&Masquerade{
		Domain:    "www.allianz-connect.com",
		IpAddress: "204.246.169.145",
	},
	&Masquerade{
		Domain:    "www.cp.misumi.jp",
		IpAddress: "54.182.3.154",
	},
	&Masquerade{
		Domain:    "www.iglobalstores.com",
		IpAddress: "13.35.4.110",
	},
	&Masquerade{
		Domain:    "www.tfly-aws.com",
		IpAddress: "52.222.134.36",
	},
	&Masquerade{
		Domain:    "payments.zynga.com",
		IpAddress: "99.86.2.135",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.129",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "99.84.2.194",
	},
	&Masquerade{
		Domain:    "cascade.madmimi.com",
		IpAddress: "13.35.3.201",
	},
	&Masquerade{
		Domain:    "api.mapbox.com",
		IpAddress: "13.35.1.183",
	},
	&Masquerade{
		Domain:    "www.test.iot.irobotapi.com",
		IpAddress: "99.84.2.101",
	},
	&Masquerade{
		Domain:    "offerup.com",
		IpAddress: "99.84.6.126",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "143.204.5.153",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.148",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.27",
	},
	&Masquerade{
		Domain:    "virmanig.myinstance.com",
		IpAddress: "54.182.7.61",
	},
	&Masquerade{
		Domain:    "coincheck.com",
		IpAddress: "13.35.2.34",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.106",
	},
	&Masquerade{
		Domain:    "prcp.pass.auone.jp",
		IpAddress: "13.224.5.108",
	},
	&Masquerade{
		Domain:    "kddi-fs.com",
		IpAddress: "143.204.2.103",
	},
	&Masquerade{
		Domain:    "stag.dazn.com",
		IpAddress: "204.246.177.149",
	},
	&Masquerade{
		Domain:    "www.sigalert.com",
		IpAddress: "99.84.0.6",
	},
	&Masquerade{
		Domain:    "nowforce.com",
		IpAddress: "52.222.135.12",
	},
	&Masquerade{
		Domain:    "truste.com",
		IpAddress: "13.35.2.195",
	},
	&Masquerade{
		Domain:    "login.schibsted.com",
		IpAddress: "13.35.6.148",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.31",
	},
	&Masquerade{
		Domain:    "customerfi.com",
		IpAddress: "13.224.0.137",
	},
	&Masquerade{
		Domain:    "www.linebc.jp",
		IpAddress: "54.182.4.177",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.109",
	},
	&Masquerade{
		Domain:    "werally.com",
		IpAddress: "13.249.6.176",
	},
	&Masquerade{
		Domain:    "www.toukei-kentei.jp",
		IpAddress: "13.35.5.50",
	},
	&Masquerade{
		Domain:    "www.nmrodam.com",
		IpAddress: "143.204.6.50",
	},
	&Masquerade{
		Domain:    "ccpsx.com",
		IpAddress: "13.35.4.35",
	},
	&Masquerade{
		Domain:    "www.vistarmedia.com",
		IpAddress: "143.204.2.86",
	},
	&Masquerade{
		Domain:    "www.connectwisedev.com",
		IpAddress: "99.84.6.153",
	},
	&Masquerade{
		Domain:    "tigocloud.net",
		IpAddress: "99.86.6.197",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.225",
	},
	&Masquerade{
		Domain:    "liftoff.io",
		IpAddress: "99.86.0.215",
	},
	&Masquerade{
		Domain:    "cdn.discounttire.com",
		IpAddress: "54.239.130.112",
	},
	&Masquerade{
		Domain:    "api.sandbox.repayonline.com",
		IpAddress: "143.204.5.225",
	},
	&Masquerade{
		Domain:    "mercadopago.com",
		IpAddress: "13.249.6.109",
	},
	&Masquerade{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "99.86.2.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.59",
	},
	&Masquerade{
		Domain:    "ecnavi.jp",
		IpAddress: "13.224.6.10",
	},
	&Masquerade{
		Domain:    "www.amazon.sa",
		IpAddress: "54.239.130.180",
	},
	&Masquerade{
		Domain:    "workflow-stage.licenses.adobe.com",
		IpAddress: "204.246.164.210",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.35.1.190",
	},
	&Masquerade{
		Domain:    "omsdocs.magento.com",
		IpAddress: "54.182.6.200",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-us-west-2.prod.hydra.sophos.com",
		IpAddress: "13.35.6.156",
	},
	&Masquerade{
		Domain:    "www.cafewell.com",
		IpAddress: "99.84.6.202",
	},
	&Masquerade{
		Domain:    "cdn.mozilla.net",
		IpAddress: "13.224.5.58",
	},
	&Masquerade{
		Domain:    "test.dazn.com",
		IpAddress: "99.86.1.48",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "99.84.0.212",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.41",
	},
	&Masquerade{
		Domain:    "abcmouse.com",
		IpAddress: "99.86.4.33",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.4.12",
	},
	&Masquerade{
		Domain:    "www.sf-cdn.net",
		IpAddress: "99.86.1.142",
	},
	&Masquerade{
		Domain:    "sni.to",
		IpAddress: "99.86.5.20",
	},
	&Masquerade{
		Domain:    "www.api.brightcove.com",
		IpAddress: "143.204.6.126",
	},
	&Masquerade{
		Domain:    "www.connectwise.com",
		IpAddress: "99.86.3.187",
	},
	&Masquerade{
		Domain:    "liftoff.io",
		IpAddress: "54.182.3.69",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "143.204.2.134",
	},
	&Masquerade{
		Domain:    "eprocurement.marketplace.us-east-1.amazonaws.com",
		IpAddress: "99.84.7.42",
	},
	&Masquerade{
		Domain:    "tonglueyun.com",
		IpAddress: "13.35.2.65",
	},
	&Masquerade{
		Domain:    "highwebmedia.com",
		IpAddress: "52.222.131.158",
	},
	&Masquerade{
		Domain:    "www.desmos.com",
		IpAddress: "13.224.5.90",
	},
	&Masquerade{
		Domain:    "flipagram.com",
		IpAddress: "13.35.1.103",
	},
	&Masquerade{
		Domain:    "oqs.amb.cybird.ne.jp",
		IpAddress: "204.246.164.101",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.98",
	},
	&Masquerade{
		Domain:    "z-eu.associates-amazon.com",
		IpAddress: "13.35.1.27",
	},
	&Masquerade{
		Domain:    "alexa.amazon.com.mx",
		IpAddress: "52.222.134.131",
	},
	&Masquerade{
		Domain:    "forgesvc.net",
		IpAddress: "54.182.6.144",
	},
	&Masquerade{
		Domain:    "aa0.awsstatic.com",
		IpAddress: "13.224.5.28",
	},
	&Masquerade{
		Domain:    "signal.is",
		IpAddress: "99.84.2.196",
	},
	&Masquerade{
		Domain:    "api.mistore.jp",
		IpAddress: "143.204.1.71",
	},
	&Masquerade{
		Domain:    "tripkit-test5.jeppesen.com",
		IpAddress: "99.84.2.116",
	},
	&Masquerade{
		Domain:    "company-target.com",
		IpAddress: "13.224.6.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.51",
	},
	&Masquerade{
		Domain:    "bittorrent.com",
		IpAddress: "13.35.0.167",
	},
	&Masquerade{
		Domain:    "www.lottedfs.com",
		IpAddress: "99.86.2.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.76",
	},
	&Masquerade{
		Domain:    "zuora.identity.fcl-02.prep.fcagcv.com",
		IpAddress: "205.251.212.182",
	},
	&Masquerade{
		Domain:    "amazon.ca",
		IpAddress: "204.246.169.232",
	},
	&Masquerade{
		Domain:    "www.accordiagolf.com",
		IpAddress: "143.204.6.150",
	},
	&Masquerade{
		Domain:    "custom-api.bigpanda.io",
		IpAddress: "54.239.130.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.184",
	},
	&Masquerade{
		Domain:    "identitynow.com",
		IpAddress: "99.84.6.47",
	},
	&Masquerade{
		Domain:    "simple-workflow.licenses.adobe.com",
		IpAddress: "99.86.5.16",
	},
	&Masquerade{
		Domain:    "wework.com",
		IpAddress: "52.222.129.203",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.4",
	},
	&Masquerade{
		Domain:    "stage-spectrum.net",
		IpAddress: "143.204.5.9",
	},
	&Masquerade{
		Domain:    "ix-cdn.brightedge.com",
		IpAddress: "143.204.0.139",
	},
	&Masquerade{
		Domain:    "sunsky-online.com",
		IpAddress: "99.84.0.14",
	},
	&Masquerade{
		Domain:    "edge.disstg.commercecloud.salesforce.com",
		IpAddress: "99.86.4.165",
	},
	&Masquerade{
		Domain:    "toysrus.co.jp",
		IpAddress: "13.35.2.29",
	},
	&Masquerade{
		Domain:    "iot.ap-southeast-2.amazonaws.com",
		IpAddress: "99.84.2.155",
	},
	&Masquerade{
		Domain:    "mix.tokyo",
		IpAddress: "99.84.0.120",
	},
	&Masquerade{
		Domain:    "img-en.fs.com",
		IpAddress: "13.249.2.110",
	},
	&Masquerade{
		Domain:    "gomlab.com",
		IpAddress: "54.182.4.103",
	},
	&Masquerade{
		Domain:    "update.hicloud.com",
		IpAddress: "13.35.1.45",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.78",
	},
	&Masquerade{
		Domain:    "www.cafewell.com",
		IpAddress: "54.182.4.149",
	},
	&Masquerade{
		Domain:    "myfitnesspal.com.tw",
		IpAddress: "143.204.1.69",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.131",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.88",
	},
	&Masquerade{
		Domain:    "seesaw.me",
		IpAddress: "54.239.130.206",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.147",
	},
	&Masquerade{
		Domain:    "segment.com",
		IpAddress: "99.86.0.85",
	},
	&Masquerade{
		Domain:    "sings-download.twitch.tv",
		IpAddress: "99.86.4.124",
	},
	&Masquerade{
		Domain:    "dsdfpay.com",
		IpAddress: "13.35.2.20",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "99.86.4.208",
	},
	&Masquerade{
		Domain:    "www.static.lottedfs.com",
		IpAddress: "143.204.6.76",
	},
	&Masquerade{
		Domain:    "www.linebc.jp",
		IpAddress: "52.222.134.178",
	},
	&Masquerade{
		Domain:    "preprod.apac.amway.net",
		IpAddress: "143.204.1.166",
	},
	&Masquerade{
		Domain:    "cdn.burlingtonenglish.com",
		IpAddress: "52.222.131.236",
	},
	&Masquerade{
		Domain:    "www.chatbar.me",
		IpAddress: "99.86.4.90",
	},
	&Masquerade{
		Domain:    "us.whispir.com",
		IpAddress: "54.239.130.204",
	},
	&Masquerade{
		Domain:    "pod-point.com",
		IpAddress: "13.249.5.54",
	},
	&Masquerade{
		Domain:    "t-x.io",
		IpAddress: "143.204.6.217",
	},
	&Masquerade{
		Domain:    "ads-interfaces.sc-cdn.net",
		IpAddress: "99.86.3.166",
	},
	&Masquerade{
		Domain:    "www.quipper.net",
		IpAddress: "54.239.192.48",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.73",
	},
	&Masquerade{
		Domain:    "cardgames.io",
		IpAddress: "143.204.7.40",
	},
	&Masquerade{
		Domain:    "primer.typekit.net",
		IpAddress: "13.249.5.2",
	},
	&Masquerade{
		Domain:    "www.accuplacer.org",
		IpAddress: "99.84.0.231",
	},
	&Masquerade{
		Domain:    "www.epop.cf.eu.aiv-cdn.net",
		IpAddress: "204.246.178.165",
	},
	&Masquerade{
		Domain:    "www.fp.ps.easebar.com",
		IpAddress: "52.222.129.108",
	},
	&Masquerade{
		Domain:    "club-beta2.pokemon.com",
		IpAddress: "54.182.2.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.31",
	},
	&Masquerade{
		Domain:    "smartrecruiters.com",
		IpAddress: "13.35.5.77",
	},
	&Masquerade{
		Domain:    "product-downloads.atlassian.com",
		IpAddress: "143.204.7.33",
	},
	&Masquerade{
		Domain:    "dev.faceid.paylabo.com",
		IpAddress: "54.182.3.224",
	},
	&Masquerade{
		Domain:    "mymathacademy.com",
		IpAddress: "204.246.164.80",
	},
	&Masquerade{
		Domain:    "amazon.co.jp",
		IpAddress: "13.35.1.16",
	},
	&Masquerade{
		Domain:    "samsungknowledge.com",
		IpAddress: "143.204.2.218",
	},
	&Masquerade{
		Domain:    "s3-turbo.amazonaws.com",
		IpAddress: "13.35.2.190",
	},
	&Masquerade{
		Domain:    "device-firmware.gp-static.com",
		IpAddress: "54.182.0.179",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.77",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "205.251.213.10",
	},
	&Masquerade{
		Domain:    "bd0.awsstatic.com",
		IpAddress: "99.86.2.33",
	},
	&Masquerade{
		Domain:    "static.pontoslivelo.com.br",
		IpAddress: "204.246.177.114",
	},
	&Masquerade{
		Domain:    "enigmasoftware.com",
		IpAddress: "216.137.39.62",
	},
	&Masquerade{
		Domain:    "adtpulseaws.net",
		IpAddress: "143.204.5.12",
	},
	&Masquerade{
		Domain:    "amazon.co.uk",
		IpAddress: "54.239.195.221",
	},
	&Masquerade{
		Domain:    "workflow-stage.licenses.adobe.com",
		IpAddress: "13.35.2.213",
	},
	&Masquerade{
		Domain:    "www.dev.pos.paylabo.com",
		IpAddress: "13.35.2.117",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.185",
	},
	&Masquerade{
		Domain:    "ads.chtbl.com",
		IpAddress: "99.86.2.117",
	},
	&Masquerade{
		Domain:    "keyuca.com",
		IpAddress: "143.204.2.164",
	},
	&Masquerade{
		Domain:    "tvc-mall.com",
		IpAddress: "13.35.2.78",
	},
	&Masquerade{
		Domain:    "static.datadoghq.com",
		IpAddress: "13.224.5.68",
	},
	&Masquerade{
		Domain:    "mymathacademy.com",
		IpAddress: "143.204.5.189",
	},
	&Masquerade{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "13.35.3.15",
	},
	&Masquerade{
		Domain:    "www.playwithsea.com",
		IpAddress: "99.84.0.81",
	},
	&Masquerade{
		Domain:    "cdn.sw.altova.com",
		IpAddress: "143.204.2.113",
	},
	&Masquerade{
		Domain:    "www.siksine.com",
		IpAddress: "54.239.192.51",
	},
	&Masquerade{
		Domain:    "static.emarsys.com",
		IpAddress: "99.84.6.190",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.120",
	},
	&Masquerade{
		Domain:    "www.samsungsmartcam.com",
		IpAddress: "54.182.7.12",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.21",
	},
	&Masquerade{
		Domain:    "oihxray-fe.aka.amazon.com",
		IpAddress: "99.84.7.43",
	},
	&Masquerade{
		Domain:    "static.adobelogin.com",
		IpAddress: "143.204.3.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.177",
	},
	&Masquerade{
		Domain:    "www.cafewellstage.com",
		IpAddress: "52.222.130.237",
	},
	&Masquerade{
		Domain:    "static.lendingclub.com",
		IpAddress: "13.35.3.33",
	},
	&Masquerade{
		Domain:    "arkoselabs.com",
		IpAddress: "13.224.6.49",
	},
	&Masquerade{
		Domain:    "ekdgd.com",
		IpAddress: "204.246.178.96",
	},
	&Masquerade{
		Domain:    "cdn.shptrn.com",
		IpAddress: "54.182.3.46",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.8",
	},
	&Masquerade{
		Domain:    "beta.awsapps.com",
		IpAddress: "54.239.130.36",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.118",
	},
	&Masquerade{
		Domain:    "forgesvc.net",
		IpAddress: "99.86.6.66",
	},
	&Masquerade{
		Domain:    "ccpsx.com",
		IpAddress: "99.84.6.111",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.86",
	},
	&Masquerade{
		Domain:    "boleto.pagseguro.com.br",
		IpAddress: "52.222.128.225",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.0.191",
	},
	&Masquerade{
		Domain:    "api1.platformdxc-d2.com",
		IpAddress: "13.249.6.218",
	},
	&Masquerade{
		Domain:    "specialized.com",
		IpAddress: "99.84.7.34",
	},
	&Masquerade{
		Domain:    "test.api.seek.co.nz",
		IpAddress: "99.84.0.86",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.148",
	},
	&Masquerade{
		Domain:    "twitchcdn.net",
		IpAddress: "54.182.2.159",
	},
	&Masquerade{
		Domain:    "kuvo.com",
		IpAddress: "204.246.164.230",
	},
	&Masquerade{
		Domain:    "api.shopbop.com",
		IpAddress: "99.84.2.45",
	},
	&Masquerade{
		Domain:    "www.srv.ygles-test.com",
		IpAddress: "99.86.0.162",
	},
	&Masquerade{
		Domain:    "www.apteligent.com",
		IpAddress: "204.246.177.4",
	},
	&Masquerade{
		Domain:    "www.accordiagolf.com",
		IpAddress: "13.35.1.150",
	},
	&Masquerade{
		Domain:    "static.uber-adsystem.com",
		IpAddress: "13.35.2.42",
	},
	&Masquerade{
		Domain:    "versal.com",
		IpAddress: "99.86.2.35",
	},
	&Masquerade{
		Domain:    "cptuat.net",
		IpAddress: "52.222.132.38",
	},
	&Masquerade{
		Domain:    "dsdfpay.com",
		IpAddress: "52.222.132.125",
	},
	&Masquerade{
		Domain:    "parsely.com",
		IpAddress: "99.84.2.17",
	},
	&Masquerade{
		Domain:    "www.neuweb.biz",
		IpAddress: "99.86.5.182",
	},
	&Masquerade{
		Domain:    "predix.io",
		IpAddress: "13.35.6.163",
	},
	&Masquerade{
		Domain:    "www.neuweb.biz",
		IpAddress: "13.249.5.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.145",
	},
	&Masquerade{
		Domain:    "www.project-a.videoprojects.net",
		IpAddress: "54.182.5.149",
	},
	&Masquerade{
		Domain:    "www.airchip.com",
		IpAddress: "99.86.3.197",
	},
	&Masquerade{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "13.35.4.154",
	},
	&Masquerade{
		Domain:    "origin-help.imdb.com",
		IpAddress: "52.222.131.73",
	},
	&Masquerade{
		Domain:    "cdn.hands.net",
		IpAddress: "52.222.132.13",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.14",
	},
	&Masquerade{
		Domain:    "opmsec.sophos.com",
		IpAddress: "13.224.5.69",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-west-1.prod.hydra.sophos.com",
		IpAddress: "99.84.6.222",
	},
	&Masquerade{
		Domain:    "fifaconnect.org",
		IpAddress: "99.84.2.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.15",
	},
	&Masquerade{
		Domain:    "www.mytaxi.com",
		IpAddress: "99.86.6.223",
	},
	&Masquerade{
		Domain:    "olt-content.sans.org",
		IpAddress: "143.204.5.18",
	},
	&Masquerade{
		Domain:    "orgsync.com",
		IpAddress: "99.86.5.44",
	},
	&Masquerade{
		Domain:    "www.listrakbi.com",
		IpAddress: "99.84.6.9",
	},
	&Masquerade{
		Domain:    "abcmouse.com",
		IpAddress: "54.239.192.214",
	},
	&Masquerade{
		Domain:    "www.withbuddies.com",
		IpAddress: "205.251.212.194",
	},
	&Masquerade{
		Domain:    "panda.chtbl.com",
		IpAddress: "54.239.192.235",
	},
	&Masquerade{
		Domain:    "www.tosconfig.com",
		IpAddress: "99.84.0.7",
	},
	&Masquerade{
		Domain:    "nba-cdn.2ksports.com",
		IpAddress: "54.239.192.183",
	},
	&Masquerade{
		Domain:    "www.culqi.com",
		IpAddress: "204.246.169.162",
	},
	&Masquerade{
		Domain:    "www.adison.co",
		IpAddress: "52.222.134.82",
	},
	&Masquerade{
		Domain:    "angular.mrowl.com",
		IpAddress: "99.86.5.228",
	},
	&Masquerade{
		Domain:    "js.pusher.com",
		IpAddress: "204.246.164.33",
	},
	&Masquerade{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "13.249.2.17",
	},
	&Masquerade{
		Domain:    "iot.us-west-2.amazonaws.com",
		IpAddress: "143.204.1.41",
	},
	&Masquerade{
		Domain:    "forgecdn.net",
		IpAddress: "99.86.4.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.68",
	},
	&Masquerade{
		Domain:    "mheducation.com",
		IpAddress: "13.35.2.71",
	},
	&Masquerade{
		Domain:    "media.aircorsica.com",
		IpAddress: "13.249.2.122",
	},
	&Masquerade{
		Domain:    "www.api.brightcove.com",
		IpAddress: "204.246.177.6",
	},
	&Masquerade{
		Domain:    "ad1.awsstatic.com",
		IpAddress: "99.86.1.95",
	},
	&Masquerade{
		Domain:    "www.brinkpos.net",
		IpAddress: "54.239.130.65",
	},
	&Masquerade{
		Domain:    "www.twitch.tv",
		IpAddress: "99.86.0.72",
	},
	&Masquerade{
		Domain:    "behance.net",
		IpAddress: "54.239.192.115",
	},
	&Masquerade{
		Domain:    "tripkit-test2.jeppesen.com",
		IpAddress: "99.86.2.75",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.38",
	},
	&Masquerade{
		Domain:    "qa.o.brightcove.com",
		IpAddress: "204.246.169.37",
	},
	&Masquerade{
		Domain:    "payment.global.rakuten.com",
		IpAddress: "13.35.3.96",
	},
	&Masquerade{
		Domain:    "origin-gql.beta.api.imdb.a2z.com",
		IpAddress: "143.204.5.147",
	},
	&Masquerade{
		Domain:    "img-en.fs.com",
		IpAddress: "13.35.5.90",
	},
	&Masquerade{
		Domain:    "cofanet.coface.com",
		IpAddress: "54.182.4.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.154",
	},
	&Masquerade{
		Domain:    "www.o9.de",
		IpAddress: "54.182.6.126",
	},
	&Masquerade{
		Domain:    "virmanig.myinstance.com",
		IpAddress: "52.222.134.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.59",
	},
	&Masquerade{
		Domain:    "www.dst.vpsvc.com",
		IpAddress: "143.204.5.89",
	},
	&Masquerade{
		Domain:    "s3-accelerate.amazonaws.com",
		IpAddress: "143.204.2.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.92",
	},
	&Masquerade{
		Domain:    "static.datadoghq.com",
		IpAddress: "99.86.1.203",
	},
	&Masquerade{
		Domain:    "www.swipesense.com",
		IpAddress: "143.204.1.81",
	},
	&Masquerade{
		Domain:    "ecnavi.jp",
		IpAddress: "54.239.192.18",
	},
	&Masquerade{
		Domain:    "www.cafewellstage.com",
		IpAddress: "143.204.2.15",
	},
	&Masquerade{
		Domain:    "multisandbox.api.fluentretail.com",
		IpAddress: "204.246.177.110",
	},
	&Masquerade{
		Domain:    "mheducation.com",
		IpAddress: "54.182.2.11",
	},
	&Masquerade{
		Domain:    "mfi-tc02.fnopf.jp",
		IpAddress: "99.86.1.187",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.17",
	},
	&Masquerade{
		Domain:    "sftelemetry-test.sophos.com",
		IpAddress: "13.35.4.10",
	},
	&Masquerade{
		Domain:    "qtest.abcmouse.com",
		IpAddress: "13.35.4.11",
	},
	&Masquerade{
		Domain:    "clients.amazonworkspaces.com",
		IpAddress: "143.204.2.114",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.25",
	},
	&Masquerade{
		Domain:    "amazon.co.jp",
		IpAddress: "99.86.1.112",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "13.35.5.134",
	},
	&Masquerade{
		Domain:    "gluon-cv.mxnet.io",
		IpAddress: "52.222.131.187",
	},
	&Masquerade{
		Domain:    "www.bijiaqi.xyz",
		IpAddress: "13.249.6.6",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.2.129",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.7",
	},
	&Masquerade{
		Domain:    "assets.cameloteurope.com",
		IpAddress: "99.86.6.109",
	},
	&Masquerade{
		Domain:    "twitchsvc-shadow.net",
		IpAddress: "143.204.6.23",
	},
	&Masquerade{
		Domain:    "z-na.associates-amazon.com",
		IpAddress: "143.204.6.100",
	},
	&Masquerade{
		Domain:    "cf.pumlo.awsps.myinstance.com",
		IpAddress: "204.246.178.16",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "54.182.6.226",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.117",
	},
	&Masquerade{
		Domain:    "api.imdbws.com",
		IpAddress: "204.246.164.5",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.52",
	},
	&Masquerade{
		Domain:    "www.stg.ui.com",
		IpAddress: "52.222.131.109",
	},
	&Masquerade{
		Domain:    "qpyou.cn",
		IpAddress: "13.224.7.29",
	},
	&Masquerade{
		Domain:    "gbf.game-a.mbga.jp",
		IpAddress: "54.182.6.23",
	},
	&Masquerade{
		Domain:    "www.aya.quipper.net",
		IpAddress: "54.182.6.173",
	},
	&Masquerade{
		Domain:    "forestry.trimble.com",
		IpAddress: "52.222.131.238",
	},
	&Masquerade{
		Domain:    "ap1.whispir.com",
		IpAddress: "205.251.212.140",
	},
	&Masquerade{
		Domain:    "openfin.co",
		IpAddress: "143.204.1.117",
	},
	&Masquerade{
		Domain:    "wordsearchbible.com",
		IpAddress: "13.35.0.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.228",
	},
	&Masquerade{
		Domain:    "www.mytaxi.com",
		IpAddress: "204.246.178.147",
	},
	&Masquerade{
		Domain:    "www.sodexomyway.com",
		IpAddress: "13.35.4.49",
	},
	&Masquerade{
		Domain:    "globalwip.cms.pearson.com",
		IpAddress: "13.35.1.233",
	},
	&Masquerade{
		Domain:    "pv.media-amazon.com",
		IpAddress: "52.222.131.206",
	},
	&Masquerade{
		Domain:    "www.execute-api.us-west-2.amazonaws.com",
		IpAddress: "54.182.3.135",
	},
	&Masquerade{
		Domain:    "prod1.superobscuredomains.com",
		IpAddress: "143.204.2.243",
	},
	&Masquerade{
		Domain:    "www.dst.vpsvc.com",
		IpAddress: "204.246.164.89",
	},
	&Masquerade{
		Domain:    "playwith.com.tw",
		IpAddress: "99.86.0.148",
	},
	&Masquerade{
		Domain:    "www.indigoag.build",
		IpAddress: "52.222.129.232",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.127",
	},
	&Masquerade{
		Domain:    "envysion.com",
		IpAddress: "13.224.5.125",
	},
	&Masquerade{
		Domain:    "i.fyu.se",
		IpAddress: "52.222.134.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.36",
	},
	&Masquerade{
		Domain:    "smile.amazon.de",
		IpAddress: "52.222.134.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.192",
	},
	&Masquerade{
		Domain:    "www.iglobalstores.com",
		IpAddress: "99.86.3.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.136",
	},
	&Masquerade{
		Domain:    "www.brinkpos.net",
		IpAddress: "54.182.3.65",
	},
	&Masquerade{
		Domain:    "kaercher.com",
		IpAddress: "99.84.0.105",
	},
	&Masquerade{
		Domain:    "dfoneople.com",
		IpAddress: "52.222.128.198",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.134.242",
	},
	&Masquerade{
		Domain:    "www.awscfdns.com",
		IpAddress: "99.84.0.107",
	},
	&Masquerade{
		Domain:    "www.adison.co",
		IpAddress: "13.224.5.162",
	},
	&Masquerade{
		Domain:    "saiercdn.imtxwy.com",
		IpAddress: "143.204.6.220",
	},
	&Masquerade{
		Domain:    "www.c.misumi-ec.com",
		IpAddress: "13.35.4.98",
	},
	&Masquerade{
		Domain:    "www.cloud.tenable.com",
		IpAddress: "99.84.5.238",
	},
	&Masquerade{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "54.182.2.238",
	},
	&Masquerade{
		Domain:    "smtown.com",
		IpAddress: "99.86.4.11",
	},
	&Masquerade{
		Domain:    "origin-api.amazonalexa.com",
		IpAddress: "54.182.6.235",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.45",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.78",
	},
	&Masquerade{
		Domain:    "paradoxplaza.com",
		IpAddress: "13.224.7.16",
	},
	&Masquerade{
		Domain:    "www.ladymay.net",
		IpAddress: "143.204.1.170",
	},
	&Masquerade{
		Domain:    "mcoc-cdn.net",
		IpAddress: "52.222.130.146",
	},
	&Masquerade{
		Domain:    "gimmegimme.it",
		IpAddress: "99.84.6.59",
	},
	&Masquerade{
		Domain:    "www.webapp.easebar.com",
		IpAddress: "13.35.3.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.21",
	},
	&Masquerade{
		Domain:    "www.goldspotmedia.com",
		IpAddress: "99.86.4.52",
	},
	&Masquerade{
		Domain:    "smtown.com",
		IpAddress: "54.182.4.202",
	},
	&Masquerade{
		Domain:    "syapp.jp",
		IpAddress: "13.249.5.36",
	},
	&Masquerade{
		Domain:    "www.uat.catchplay.com",
		IpAddress: "54.182.3.118",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.17",
	},
	&Masquerade{
		Domain:    "angels.camp-fire.jp",
		IpAddress: "143.204.6.45",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.33",
	},
	&Masquerade{
		Domain:    "www.midasplayer.com",
		IpAddress: "99.84.0.20",
	},
	&Masquerade{
		Domain:    "sellercentral.amazon.com",
		IpAddress: "99.86.6.59",
	},
	&Masquerade{
		Domain:    "gcsp.jnj.com",
		IpAddress: "54.239.192.139",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.189",
	},
	&Masquerade{
		Domain:    "www.culqi.com",
		IpAddress: "99.84.2.239",
	},
	&Masquerade{
		Domain:    "verti.iptiq.de",
		IpAddress: "54.239.130.181",
	},
	&Masquerade{
		Domain:    "kucoin.com",
		IpAddress: "54.182.3.54",
	},
	&Masquerade{
		Domain:    "versal.com",
		IpAddress: "205.251.212.103",
	},
	&Masquerade{
		Domain:    "rest.immobilienscout24.de",
		IpAddress: "13.35.1.151",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.179",
	},
	&Masquerade{
		Domain:    "www.brickworksoftware.com",
		IpAddress: "143.204.6.216",
	},
	&Masquerade{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "204.246.178.146",
	},
	&Masquerade{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "54.182.3.37",
	},
	&Masquerade{
		Domain:    "api.mapbox.com",
		IpAddress: "99.86.1.140",
	},
	&Masquerade{
		Domain:    "qpyou.cn",
		IpAddress: "204.246.164.92",
	},
	&Masquerade{
		Domain:    "thescore.com",
		IpAddress: "13.35.5.152",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.152",
	},
	&Masquerade{
		Domain:    "smartica.jp",
		IpAddress: "99.84.0.164",
	},
	&Masquerade{
		Domain:    "www.tfly-aws.com",
		IpAddress: "99.86.3.147",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.69",
	},
	&Masquerade{
		Domain:    "fe.dazn-stage.com",
		IpAddress: "13.249.6.77",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.31",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "13.35.2.108",
	},
	&Masquerade{
		Domain:    "clients.a.chime.aws",
		IpAddress: "13.224.5.171",
	},
	&Masquerade{
		Domain:    "passporthealthglobal.com",
		IpAddress: "99.86.1.100",
	},
	&Masquerade{
		Domain:    "dl.amazon.co.uk",
		IpAddress: "99.86.2.34",
	},
	&Masquerade{
		Domain:    "samsunghealth.com",
		IpAddress: "204.246.164.214",
	},
	&Masquerade{
		Domain:    "www.xp-assets.aiv-cdn.net",
		IpAddress: "99.86.4.189",
	},
	&Masquerade{
		Domain:    "dev.awsapps.com",
		IpAddress: "13.35.6.36",
	},
	&Masquerade{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "204.246.177.85",
	},
	&Masquerade{
		Domain:    "ubnt.com",
		IpAddress: "13.35.5.112",
	},
	&Masquerade{
		Domain:    "dev.twitch.tv",
		IpAddress: "13.224.5.153",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "54.239.130.207",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.18",
	},
	&Masquerade{
		Domain:    "origin-www.amazon.com.tr",
		IpAddress: "52.222.134.218",
	},
	&Masquerade{
		Domain:    "cdn.realtimeprocess.net",
		IpAddress: "52.222.128.177",
	},
	&Masquerade{
		Domain:    "cdn.burlingtonenglish.com",
		IpAddress: "13.249.2.98",
	},
	&Masquerade{
		Domain:    "enetscores.com",
		IpAddress: "99.86.2.84",
	},
	&Masquerade{
		Domain:    "s3-turbo.amazonaws.com",
		IpAddress: "52.222.131.175",
	},
	&Masquerade{
		Domain:    "seesaw.me",
		IpAddress: "54.239.195.207",
	},
	&Masquerade{
		Domain:    "www.enjoy.point.auone.jp",
		IpAddress: "13.35.6.155",
	},
	&Masquerade{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "143.204.1.12",
	},
	&Masquerade{
		Domain:    "classic.dm.amplience.net",
		IpAddress: "13.224.6.86",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.103",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.133",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.45",
	},
	&Masquerade{
		Domain:    "www.recoru.in",
		IpAddress: "52.222.129.155",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.82",
	},
	&Masquerade{
		Domain:    "coupang.net",
		IpAddress: "54.239.192.206",
	},
	&Masquerade{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "204.246.177.33",
	},
	&Masquerade{
		Domain:    "www.ooyala.com",
		IpAddress: "99.84.2.225",
	},
	&Masquerade{
		Domain:    "www.pearsonperspective.com",
		IpAddress: "143.204.2.70",
	},
	&Masquerade{
		Domain:    "cdn.fdp.foreflight.com",
		IpAddress: "13.35.6.28",
	},
	&Masquerade{
		Domain:    "camp-fire.jp",
		IpAddress: "205.251.212.38",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.189",
	},
	&Masquerade{
		Domain:    "apps.bahrain.bh",
		IpAddress: "99.86.6.96",
	},
	&Masquerade{
		Domain:    "www.quipper.com",
		IpAddress: "99.86.6.173",
	},
	&Masquerade{
		Domain:    "dev.sotappm.auone.jp",
		IpAddress: "99.84.6.151",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.27",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.140",
	},
	&Masquerade{
		Domain:    "adtpulseaws.net",
		IpAddress: "13.35.3.181",
	},
	&Masquerade{
		Domain:    "cont-test.mydaiz.jp",
		IpAddress: "204.246.178.51",
	},
	&Masquerade{
		Domain:    "payment.global.rakuten.com",
		IpAddress: "52.222.131.101",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.221",
	},
	&Masquerade{
		Domain:    "www.freshdesk.com",
		IpAddress: "54.182.6.230",
	},
	&Masquerade{
		Domain:    "www.test.iot.irobotapi.com",
		IpAddress: "13.224.6.150",
	},
	&Masquerade{
		Domain:    "mark1.dev",
		IpAddress: "13.224.7.33",
	},
	&Masquerade{
		Domain:    "auth.nightowlx.com",
		IpAddress: "13.35.1.86",
	},
	&Masquerade{
		Domain:    "media.edgenuity.com",
		IpAddress: "13.224.5.206",
	},
	&Masquerade{
		Domain:    "iot.eu-west-1.amazonaws.com",
		IpAddress: "13.35.6.97",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.119",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.77",
	},
	&Masquerade{
		Domain:    "custom-api.bigpanda.io",
		IpAddress: "13.35.6.130",
	},
	&Masquerade{
		Domain:    "souqcdn.com",
		IpAddress: "13.35.3.69",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "54.182.4.112",
	},
	&Masquerade{
		Domain:    "www.bcovlive.io",
		IpAddress: "52.222.132.176",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.15",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.143",
	},
	&Masquerade{
		Domain:    "www.milkvr.rocks",
		IpAddress: "204.246.164.119",
	},
	&Masquerade{
		Domain:    "dsdfpay.com",
		IpAddress: "13.249.2.216",
	},
	&Masquerade{
		Domain:    "amoad.com",
		IpAddress: "54.182.4.231",
	},
	&Masquerade{
		Domain:    "www.me2zengame.com",
		IpAddress: "13.35.3.117",
	},
	&Masquerade{
		Domain:    "esd.sentinelcloud.com",
		IpAddress: "99.84.6.78",
	},
	&Masquerade{
		Domain:    "login.schibsted.com",
		IpAddress: "204.246.164.147",
	},
	&Masquerade{
		Domain:    "wpcp.shiseido.co.jp",
		IpAddress: "205.251.212.233",
	},
	&Masquerade{
		Domain:    "one.accedo.tv",
		IpAddress: "204.246.169.10",
	},
	&Masquerade{
		Domain:    "www.binance.vision",
		IpAddress: "52.222.132.68",
	},
	&Masquerade{
		Domain:    "www.apkimage.io",
		IpAddress: "54.182.3.140",
	},
	&Masquerade{
		Domain:    "www.mytaxi.com",
		IpAddress: "54.182.5.163",
	},
	&Masquerade{
		Domain:    "www.uniqlo.com",
		IpAddress: "99.86.4.113",
	},
	&Masquerade{
		Domain:    "www.travelhook.com",
		IpAddress: "143.204.2.10",
	},
	&Masquerade{
		Domain:    "static.counsyl.com",
		IpAddress: "54.182.7.68",
	},
	&Masquerade{
		Domain:    "forestry.trimble.com",
		IpAddress: "143.204.5.122",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.6.163",
	},
	&Masquerade{
		Domain:    "www.clearlinkdata.com",
		IpAddress: "143.204.1.141",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.6.122",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.90",
	},
	&Masquerade{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "54.182.3.87",
	},
	&Masquerade{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "52.222.130.161",
	},
	&Masquerade{
		Domain:    "saucelabs.com",
		IpAddress: "13.35.1.160",
	},
	&Masquerade{
		Domain:    "snapfinance.com",
		IpAddress: "99.86.0.29",
	},
	&Masquerade{
		Domain:    "jtvnw-30eb2e4e018997e11b2884b1f80a025c.twitchcdn.net",
		IpAddress: "52.222.129.59",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.217",
	},
	&Masquerade{
		Domain:    "datad0g.com",
		IpAddress: "99.84.2.72",
	},
	&Masquerade{
		Domain:    "cdn.venividivicci.de",
		IpAddress: "204.246.169.87",
	},
	&Masquerade{
		Domain:    "sparxcdn.net",
		IpAddress: "99.86.4.105",
	},
	&Masquerade{
		Domain:    "rubiconproject.com",
		IpAddress: "54.182.4.117",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.149",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-central-1.inf.hydra.sophos.com",
		IpAddress: "99.86.4.85",
	},
	&Masquerade{
		Domain:    "pimg.jp",
		IpAddress: "13.249.5.27",
	},
	&Masquerade{
		Domain:    "evident.io",
		IpAddress: "13.249.7.29",
	},
	&Masquerade{
		Domain:    "api.beta.tab.com.au",
		IpAddress: "99.84.6.53",
	},
	&Masquerade{
		Domain:    "dolphin-fe.amazon.com",
		IpAddress: "54.182.5.228",
	},
	&Masquerade{
		Domain:    "sywm-kr.gdl.imtxwy.com",
		IpAddress: "54.182.3.124",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.91",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.225",
	},
	&Masquerade{
		Domain:    "gallery.mailchimp.com",
		IpAddress: "13.249.5.78",
	},
	&Masquerade{
		Domain:    "www.binancechain.io",
		IpAddress: "54.239.130.144",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.11",
	},
	&Masquerade{
		Domain:    "iot.us-east-2.amazonaws.com",
		IpAddress: "13.249.6.198",
	},
	&Masquerade{
		Domain:    "z-na.associates-amazon.com",
		IpAddress: "54.182.6.51",
	},
	&Masquerade{
		Domain:    "static-cdn.jtvnw.net",
		IpAddress: "204.246.178.131",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.172",
	},
	&Masquerade{
		Domain:    "docomo-ntsupport.jp",
		IpAddress: "13.249.2.140",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.55",
	},
	&Masquerade{
		Domain:    "ubnt.com",
		IpAddress: "143.204.2.125",
	},
	&Masquerade{
		Domain:    "yieldoptimizer.com",
		IpAddress: "13.35.2.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.11",
	},
	&Masquerade{
		Domain:    "panda.chtbl.com",
		IpAddress: "204.246.169.18",
	},
	&Masquerade{
		Domain:    "ekdgd.com",
		IpAddress: "99.86.4.2",
	},
	&Masquerade{
		Domain:    "dev.sotappm.auone.jp",
		IpAddress: "13.35.2.4",
	},
	&Masquerade{
		Domain:    "webarchive.nationalarchives.gov.uk",
		IpAddress: "52.222.135.38",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "13.35.5.162",
	},
	&Masquerade{
		Domain:    "cpe.wtf",
		IpAddress: "13.35.3.225",
	},
	&Masquerade{
		Domain:    "oneblood.org",
		IpAddress: "99.84.2.137",
	},
	&Masquerade{
		Domain:    "thetvdb.com",
		IpAddress: "54.182.6.98",
	},
	&Masquerade{
		Domain:    "offerup.com",
		IpAddress: "52.222.135.30",
	},
	&Masquerade{
		Domain:    "www.channel4.com",
		IpAddress: "143.204.5.164",
	},
	&Masquerade{
		Domain:    "amazon.ca",
		IpAddress: "143.204.2.43",
	},
	&Masquerade{
		Domain:    "arevea.tv",
		IpAddress: "99.86.0.219",
	},
	&Masquerade{
		Domain:    "imbd-pro.net",
		IpAddress: "13.224.0.239",
	},
	&Masquerade{
		Domain:    "bd1.awsstatic.com",
		IpAddress: "143.204.1.113",
	},
	&Masquerade{
		Domain:    "club.ubisoft.com",
		IpAddress: "99.84.0.142",
	},
	&Masquerade{
		Domain:    "samsungqbe.com",
		IpAddress: "13.224.6.206",
	},
	&Masquerade{
		Domain:    "www.ladymay.net",
		IpAddress: "13.35.1.13",
	},
	&Masquerade{
		Domain:    "www.collegescheduler.com",
		IpAddress: "54.239.195.136",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "54.239.130.100",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.159",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.126",
	},
	&Masquerade{
		Domain:    "www.readingiq.com",
		IpAddress: "13.35.6.8",
	},
	&Masquerade{
		Domain:    "ring.com",
		IpAddress: "13.249.6.62",
	},
	&Masquerade{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "54.182.5.160",
	},
	&Masquerade{
		Domain:    "dsdfpay.com",
		IpAddress: "99.86.6.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.29",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.10",
	},
	&Masquerade{
		Domain:    "geocomply.net",
		IpAddress: "143.204.1.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.93",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.34",
	},
	&Masquerade{
		Domain:    "d.nanairo.coop",
		IpAddress: "54.182.4.82",
	},
	&Masquerade{
		Domain:    "scoring.pearsonassessments.com",
		IpAddress: "54.182.5.199",
	},
	&Masquerade{
		Domain:    "www.tipico.com",
		IpAddress: "143.204.6.52",
	},
	&Masquerade{
		Domain:    "www.adbephotos.com",
		IpAddress: "99.86.4.16",
	},
	&Masquerade{
		Domain:    "update.hicloud.com",
		IpAddress: "54.182.7.16",
	},
	&Masquerade{
		Domain:    "myfonts.net",
		IpAddress: "13.249.2.36",
	},
	&Masquerade{
		Domain:    "assets.cameloteurope.com",
		IpAddress: "143.204.1.112",
	},
	&Masquerade{
		Domain:    "widencdn.net",
		IpAddress: "52.222.128.185",
	},
	&Masquerade{
		Domain:    "pactsafe.io",
		IpAddress: "204.246.169.122",
	},
	&Masquerade{
		Domain:    "www.ebookstore.sony.jp",
		IpAddress: "13.249.7.32",
	},
	&Masquerade{
		Domain:    "lottedfs.com",
		IpAddress: "54.239.130.146",
	},
	&Masquerade{
		Domain:    "file.samsungcloud.com",
		IpAddress: "52.222.130.194",
	},
	&Masquerade{
		Domain:    "hicloud.com",
		IpAddress: "99.84.2.169",
	},
	&Masquerade{
		Domain:    "site.skychnl.net",
		IpAddress: "99.84.0.236",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "13.35.2.23",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-west-1.dev.hydra.sophos.com",
		IpAddress: "52.222.132.17",
	},
	&Masquerade{
		Domain:    "aa0.awsstatic.com",
		IpAddress: "143.204.7.30",
	},
	&Masquerade{
		Domain:    "d.nanairo.coop",
		IpAddress: "99.84.0.207",
	},
	&Masquerade{
		Domain:    "rlmcdn.net",
		IpAddress: "143.204.0.186",
	},
	&Masquerade{
		Domain:    "rebrandly.com",
		IpAddress: "54.239.192.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.78",
	},
	&Masquerade{
		Domain:    "livethumb.huluim.com",
		IpAddress: "13.35.6.136",
	},
	&Masquerade{
		Domain:    "www.sigalert.com",
		IpAddress: "13.35.3.6",
	},
	&Masquerade{
		Domain:    "us.whispir.com",
		IpAddress: "13.35.6.166",
	},
	&Masquerade{
		Domain:    "unrealengine.com",
		IpAddress: "54.182.0.209",
	},
	&Masquerade{
		Domain:    "yq.gph.imtxwy.com",
		IpAddress: "99.86.4.47",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.8",
	},
	&Masquerade{
		Domain:    "spd.samsungdm.com",
		IpAddress: "204.246.178.48",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.71",
	},
	&Masquerade{
		Domain:    "developercentral.amazon.com",
		IpAddress: "99.86.0.115",
	},
	&Masquerade{
		Domain:    "giv-dev.nmgcloud.io",
		IpAddress: "52.222.134.156",
	},
	&Masquerade{
		Domain:    "www.myharmony.com",
		IpAddress: "54.239.192.14",
	},
	&Masquerade{
		Domain:    "rlmcdn.net",
		IpAddress: "99.84.6.140",
	},
	&Masquerade{
		Domain:    "tapad.com",
		IpAddress: "99.86.5.2",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.22",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.218",
	},
	&Masquerade{
		Domain:    "rview.com",
		IpAddress: "143.204.2.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.4.23",
	},
	&Masquerade{
		Domain:    "binance.sg",
		IpAddress: "13.35.1.73",
	},
	&Masquerade{
		Domain:    "m.betaex.com",
		IpAddress: "54.239.192.223",
	},
	&Masquerade{
		Domain:    "www.nyc837-dev.gin-dev.com",
		IpAddress: "13.249.2.60",
	},
	&Masquerade{
		Domain:    "www.g.mkey.163.com",
		IpAddress: "99.86.5.51",
	},
	&Masquerade{
		Domain:    "www.findawayworld.com",
		IpAddress: "143.204.0.137",
	},
	&Masquerade{
		Domain:    "clients.chime.aws",
		IpAddress: "54.239.192.84",
	},
	&Masquerade{
		Domain:    "pubcerts-stage.licenses.adobe.com",
		IpAddress: "99.84.2.20",
	},
	&Masquerade{
		Domain:    "ring.com",
		IpAddress: "54.182.4.106",
	},
	&Masquerade{
		Domain:    "searchandexplore.com",
		IpAddress: "13.249.5.127",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.18",
	},
	&Masquerade{
		Domain:    "assets1.uswitch.com",
		IpAddress: "54.239.130.50",
	},
	&Masquerade{
		Domain:    "smile.amazon.de",
		IpAddress: "54.182.5.238",
	},
	&Masquerade{
		Domain:    "api.stage.context.cloud.sap",
		IpAddress: "13.224.6.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.99",
	},
	&Masquerade{
		Domain:    "static.ddog-gov.com",
		IpAddress: "204.246.178.46",
	},
	&Masquerade{
		Domain:    "club.ubisoft.com",
		IpAddress: "52.222.132.95",
	},
	&Masquerade{
		Domain:    "vdownload.cyberoam.com",
		IpAddress: "13.35.1.197",
	},
	&Masquerade{
		Domain:    "whopper.com",
		IpAddress: "99.86.6.236",
	},
	&Masquerade{
		Domain:    "vlive-simulcast.sans.org",
		IpAddress: "204.246.177.102",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.86.0.127",
	},
	&Masquerade{
		Domain:    "www.ladymay.net",
		IpAddress: "54.182.6.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.4",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.13",
	},
	&Masquerade{
		Domain:    "phdvasia.com",
		IpAddress: "13.35.6.26",
	},
	&Masquerade{
		Domain:    "twitchsvc.tech",
		IpAddress: "52.222.131.154",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.66",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.79",
	},
	&Masquerade{
		Domain:    "select.au.com",
		IpAddress: "13.249.2.51",
	},
	&Masquerade{
		Domain:    "avatax.avalara.net",
		IpAddress: "204.246.178.68",
	},
	&Masquerade{
		Domain:    "cdn.discounttire.com",
		IpAddress: "13.35.5.114",
	},
	&Masquerade{
		Domain:    "secb2b.com",
		IpAddress: "54.182.2.2",
	},
	&Masquerade{
		Domain:    "www.appservers.net",
		IpAddress: "52.222.132.232",
	},
	&Masquerade{
		Domain:    "datadoghq.com",
		IpAddress: "99.86.2.170",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "54.182.4.213",
	},
	&Masquerade{
		Domain:    "forhims.com",
		IpAddress: "54.182.6.101",
	},
	&Masquerade{
		Domain:    "jfrog.io",
		IpAddress: "54.182.4.87",
	},
	&Masquerade{
		Domain:    "zuora.identity.fcl-01.fcagcv.com",
		IpAddress: "52.222.131.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.8",
	},
	&Masquerade{
		Domain:    "www.c.ooyala.com",
		IpAddress: "13.224.0.227",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.193",
	},
	&Masquerade{
		Domain:    "dev1-www.lifelockunlocked.com",
		IpAddress: "99.84.2.124",
	},
	&Masquerade{
		Domain:    "file.samsungcloud.com",
		IpAddress: "143.204.6.146",
	},
	&Masquerade{
		Domain:    "gallery.mailchimp.com",
		IpAddress: "99.86.0.77",
	},
	&Masquerade{
		Domain:    "nexon.com",
		IpAddress: "54.239.130.51",
	},
	&Masquerade{
		Domain:    "perseus.de",
		IpAddress: "13.35.3.175",
	},
	&Masquerade{
		Domain:    "kaltura.com",
		IpAddress: "13.35.2.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.178",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.226",
	},
	&Masquerade{
		Domain:    "mfi-tc02.fnopf.jp",
		IpAddress: "99.86.2.187",
	},
	&Masquerade{
		Domain:    "bolindadigital.com",
		IpAddress: "54.182.6.177",
	},
	&Masquerade{
		Domain:    "es-navi.com",
		IpAddress: "143.204.2.32",
	},
	&Masquerade{
		Domain:    "sha-cf.v.uname.link",
		IpAddress: "99.84.6.86",
	},
	&Masquerade{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "99.86.2.36",
	},
	&Masquerade{
		Domain:    "wa.aws.amazon.com",
		IpAddress: "99.86.3.62",
	},
	&Masquerade{
		Domain:    "ad1.awsstatic.com",
		IpAddress: "204.246.169.161",
	},
	&Masquerade{
		Domain:    "aws.amazon.com",
		IpAddress: "99.84.4.72",
	},
	&Masquerade{
		Domain:    "www.brinkpos.net",
		IpAddress: "204.246.169.65",
	},
	&Masquerade{
		Domain:    "amazon.ca",
		IpAddress: "54.239.130.195",
	},
	&Masquerade{
		Domain:    "www.dn.nexoncdn.co.kr",
		IpAddress: "13.224.6.113",
	},
	&Masquerade{
		Domain:    "www.dazndn.com",
		IpAddress: "143.204.1.85",
	},
	&Masquerade{
		Domain:    "www.accordiagolf.com",
		IpAddress: "13.249.6.157",
	},
	&Masquerade{
		Domain:    "www.cphostaccess.com",
		IpAddress: "13.35.6.4",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.147",
	},
	&Masquerade{
		Domain:    "www.execute-api.us-east-1.amazonaws.com",
		IpAddress: "143.204.5.97",
	},
	&Masquerade{
		Domain:    "mpago.la",
		IpAddress: "13.35.6.206",
	},
	&Masquerade{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "13.35.5.196",
	},
	&Masquerade{
		Domain:    "www.playwithsea.com",
		IpAddress: "52.222.129.192",
	},
	&Masquerade{
		Domain:    "www.toukei-kentei.jp",
		IpAddress: "13.249.6.163",
	},
	&Masquerade{
		Domain:    "www.production.scrabble.withbuddies.com",
		IpAddress: "54.182.0.158",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.89",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.164",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.53",
	},
	&Masquerade{
		Domain:    "marketpulse.com",
		IpAddress: "143.204.1.165",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.194",
	},
	&Masquerade{
		Domain:    "sup-gcsp.jnj.com",
		IpAddress: "13.35.5.237",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.12",
	},
	&Masquerade{
		Domain:    "iot.us-east-1.amazonaws.com",
		IpAddress: "13.249.6.92",
	},
	&Masquerade{
		Domain:    "nba-cdn.2ksports.com",
		IpAddress: "52.222.131.180",
	},
	&Masquerade{
		Domain:    "carevisor.com",
		IpAddress: "54.182.6.225",
	},
	&Masquerade{
		Domain:    "enetscores.com",
		IpAddress: "204.246.178.44",
	},
	&Masquerade{
		Domain:    "www.bl.booklive.jp",
		IpAddress: "99.84.6.166",
	},
	&Masquerade{
		Domain:    "cdn-cloudfront.krxd.net",
		IpAddress: "99.84.2.49",
	},
	&Masquerade{
		Domain:    "ba0.awsstatic.com",
		IpAddress: "54.182.0.160",
	},
	&Masquerade{
		Domain:    "smartica.jp",
		IpAddress: "13.35.6.129",
	},
	&Masquerade{
		Domain:    "api.digitalstudios.discovery.com",
		IpAddress: "204.246.164.165",
	},
	&Masquerade{
		Domain:    "as0.awsstatic.com",
		IpAddress: "99.86.2.23",
	},
	&Masquerade{
		Domain:    "zurple.com",
		IpAddress: "99.84.0.42",
	},
	&Masquerade{
		Domain:    "apps.bahrain.bh",
		IpAddress: "99.84.0.130",
	},
	&Masquerade{
		Domain:    "update.hicloud.com",
		IpAddress: "204.246.178.15",
	},
	&Masquerade{
		Domain:    "dl.amazon.com",
		IpAddress: "54.182.5.235",
	},
	&Masquerade{
		Domain:    "www.gph.imtxwy.com",
		IpAddress: "13.35.5.229",
	},
	&Masquerade{
		Domain:    "preprod.apac.amway.net",
		IpAddress: "143.204.6.208",
	},
	&Masquerade{
		Domain:    "origin-gql.beta.api.imdb.a2z.com",
		IpAddress: "99.86.0.142",
	},
	&Masquerade{
		Domain:    "specialized.com",
		IpAddress: "13.35.2.115",
	},
	&Masquerade{
		Domain:    "contestimg.wish.com",
		IpAddress: "204.246.164.57",
	},
	&Masquerade{
		Domain:    "siedev.net",
		IpAddress: "204.246.177.70",
	},
	&Masquerade{
		Domain:    "amazon.nl",
		IpAddress: "204.246.177.2",
	},
	&Masquerade{
		Domain:    "www.thinkthroughmath.com",
		IpAddress: "52.222.134.129",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "99.86.4.49",
	},
	&Masquerade{
		Domain:    "www.cafewell.com",
		IpAddress: "13.224.6.193",
	},
	&Masquerade{
		Domain:    "api.area-hinan-test.au.com",
		IpAddress: "54.239.192.54",
	},
	&Masquerade{
		Domain:    "www.shufu-job.jp",
		IpAddress: "13.224.0.226",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "143.204.3.20",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.17",
	},
	&Masquerade{
		Domain:    "static.amundi.com",
		IpAddress: "99.84.6.174",
	},
	&Masquerade{
		Domain:    "cf.pumlo.awsps.myinstance.com",
		IpAddress: "52.222.135.18",
	},
	&Masquerade{
		Domain:    "bks.cybird.ne.jp",
		IpAddress: "54.182.0.182",
	},
	&Masquerade{
		Domain:    "tvc-mall.com",
		IpAddress: "204.246.178.128",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "13.224.6.60",
	},
	&Masquerade{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "204.246.164.94",
	},
	&Masquerade{
		Domain:    "www.bounceexchange.com",
		IpAddress: "205.251.212.30",
	},
	&Masquerade{
		Domain:    "my.ellotte.com",
		IpAddress: "205.251.212.144",
	},
	&Masquerade{
		Domain:    "www.sodexomyway.com",
		IpAddress: "204.246.164.49",
	},
	&Masquerade{
		Domain:    "origin-client.legacy-app.games.a2z.com",
		IpAddress: "204.246.178.66",
	},
	&Masquerade{
		Domain:    "enetscores.com",
		IpAddress: "54.182.5.146",
	},
	&Masquerade{
		Domain:    "www.p7s1.io",
		IpAddress: "54.239.192.89",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.9",
	},
	&Masquerade{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "143.204.1.44",
	},
	&Masquerade{
		Domain:    "geocomply.com",
		IpAddress: "99.86.0.26",
	},
	&Masquerade{
		Domain:    "dev.ctrf.api.eden.mediba.jp",
		IpAddress: "13.35.4.53",
	},
	&Masquerade{
		Domain:    "www.chartboost.com",
		IpAddress: "54.182.2.144",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.2",
	},
	&Masquerade{
		Domain:    "dl.amazon.com",
		IpAddress: "54.182.0.235",
	},
	&Masquerade{
		Domain:    "www.gamma.awsapps.com",
		IpAddress: "52.222.129.84",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.3.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.196",
	},
	&Masquerade{
		Domain:    "www.swipesense.com",
		IpAddress: "13.249.5.128",
	},
	&Masquerade{
		Domain:    "www.indigoag.tech",
		IpAddress: "13.35.6.107",
	},
	&Masquerade{
		Domain:    "aiag.i-memo.jp",
		IpAddress: "99.84.6.68",
	},
	&Masquerade{
		Domain:    "www.linebc.jp",
		IpAddress: "54.239.130.178",
	},
	&Masquerade{
		Domain:    "avatax.avalara.net",
		IpAddress: "99.84.6.54",
	},
	&Masquerade{
		Domain:    "undercovertourist.com",
		IpAddress: "13.224.5.62",
	},
	&Masquerade{
		Domain:    "giv-dev.nmgcloud.io",
		IpAddress: "143.204.1.139",
	},
	&Masquerade{
		Domain:    "adventureacademy.com",
		IpAddress: "204.246.177.99",
	},
	&Masquerade{
		Domain:    "amazon.de",
		IpAddress: "13.35.5.149",
	},
	&Masquerade{
		Domain:    "kindle-guru.amazon.com",
		IpAddress: "204.246.164.124",
	},
	&Masquerade{
		Domain:    "www.innov8.space",
		IpAddress: "54.182.2.184",
	},
	&Masquerade{
		Domain:    "www.quick-cdn.com",
		IpAddress: "13.249.6.199",
	},
	&Masquerade{
		Domain:    "iot.ap-southeast-2.amazonaws.com",
		IpAddress: "54.182.3.218",
	},
	&Masquerade{
		Domain:    "amazon.co.uk",
		IpAddress: "54.239.192.101",
	},
	&Masquerade{
		Domain:    "video.counsyl.com",
		IpAddress: "205.251.212.161",
	},
	&Masquerade{
		Domain:    "origin-beta.client.legacy-app.games.a2z.com",
		IpAddress: "13.35.3.193",
	},
	&Masquerade{
		Domain:    "www.janrain.com",
		IpAddress: "99.86.5.158",
	},
	&Masquerade{
		Domain:    "www.i-ready.com",
		IpAddress: "52.222.130.187",
	},
	&Masquerade{
		Domain:    "api.msg.ue1.b.app.chime.aws",
		IpAddress: "13.249.6.161",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.5",
	},
	&Masquerade{
		Domain:    "livemeat.jp",
		IpAddress: "13.35.5.109",
	},
	&Masquerade{
		Domain:    "www.gph.imtxwy.com",
		IpAddress: "99.86.0.111",
	},
	&Masquerade{
		Domain:    "mpago.la",
		IpAddress: "13.249.2.25",
	},
	&Masquerade{
		Domain:    "mheducation.com",
		IpAddress: "54.182.2.33",
	},
	&Masquerade{
		Domain:    "musixmatch.com",
		IpAddress: "13.35.6.2",
	},
	&Masquerade{
		Domain:    "www.dreambox.com",
		IpAddress: "13.224.5.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.4",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.199",
	},
	&Masquerade{
		Domain:    "www.patient-create.orthofi-dev.com",
		IpAddress: "143.204.6.54",
	},
	&Masquerade{
		Domain:    "static.datadoghq.com",
		IpAddress: "54.182.6.111",
	},
	&Masquerade{
		Domain:    "www.update.easebar.com",
		IpAddress: "52.222.129.162",
	},
	&Masquerade{
		Domain:    "www.taggstar.com",
		IpAddress: "13.35.5.132",
	},
	&Masquerade{
		Domain:    "www.twitch.tv",
		IpAddress: "52.222.129.130",
	},
	&Masquerade{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "13.35.6.80",
	},
	&Masquerade{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "13.35.5.173",
	},
	&Masquerade{
		Domain:    "www.infomedia.com.au",
		IpAddress: "13.35.3.167",
	},
	&Masquerade{
		Domain:    "altium.com",
		IpAddress: "13.35.5.195",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.97",
	},
	&Masquerade{
		Domain:    "plaync.com",
		IpAddress: "54.239.192.140",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.44",
	},
	&Masquerade{
		Domain:    "guipitan.amazon.co.jp",
		IpAddress: "52.222.128.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.73",
	},
	&Masquerade{
		Domain:    "nowforce.com",
		IpAddress: "54.182.4.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.100",
	},
	&Masquerade{
		Domain:    "www.placelocal.com",
		IpAddress: "13.35.0.178",
	},
	&Masquerade{
		Domain:    "zeasn.tv",
		IpAddress: "54.239.192.109",
	},
	&Masquerade{
		Domain:    "gateway.prod.compass.pioneer.com",
		IpAddress: "204.246.178.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.32",
	},
	&Masquerade{
		Domain:    "www.nyc837-dev.gin-dev.com",
		IpAddress: "143.204.1.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.99",
	},
	&Masquerade{
		Domain:    "forgecdn.net",
		IpAddress: "52.222.134.14",
	},
	&Masquerade{
		Domain:    "prod2.superobscuredomains.com",
		IpAddress: "52.222.133.252",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.205",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.3",
	},
	&Masquerade{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "52.222.130.150",
	},
	&Masquerade{
		Domain:    "www.dev.dgame.dmkt-sp.jp",
		IpAddress: "99.84.6.131",
	},
	&Masquerade{
		Domain:    "www.nmrodam.com",
		IpAddress: "13.35.2.209",
	},
	&Masquerade{
		Domain:    "assets.bwbx.io",
		IpAddress: "54.239.192.207",
	},
	&Masquerade{
		Domain:    "www.freshdesk.com",
		IpAddress: "99.84.2.42",
	},
	&Masquerade{
		Domain:    "macmillanyounglearners.com",
		IpAddress: "13.224.7.18",
	},
	&Masquerade{
		Domain:    "www.stg.misumi-ec.com",
		IpAddress: "13.224.0.167",
	},
	&Masquerade{
		Domain:    "movergames.com",
		IpAddress: "204.246.177.156",
	},
	&Masquerade{
		Domain:    "chime.aws",
		IpAddress: "99.86.6.11",
	},
	&Masquerade{
		Domain:    "sings-download.twitch.tv",
		IpAddress: "143.204.2.9",
	},
	&Masquerade{
		Domain:    "club.ubisoft.com",
		IpAddress: "13.35.3.170",
	},
	&Masquerade{
		Domain:    "adn.wyzant.com",
		IpAddress: "13.35.5.184",
	},
	&Masquerade{
		Domain:    "www.ashcream.xyz",
		IpAddress: "13.224.5.212",
	},
	&Masquerade{
		Domain:    "www.awsapps.com",
		IpAddress: "52.222.132.94",
	},
	&Masquerade{
		Domain:    "api.shopbop.com",
		IpAddress: "52.222.135.11",
	},
	&Masquerade{
		Domain:    "saucelabs.com",
		IpAddress: "99.84.6.210",
	},
	&Masquerade{
		Domain:    "mkw.melbourne.vic.gov.au",
		IpAddress: "143.204.2.225",
	},
	&Masquerade{
		Domain:    "bethesda.net",
		IpAddress: "143.204.2.219",
	},
	&Masquerade{
		Domain:    "api.stg.smartpass.auone.jp",
		IpAddress: "99.84.2.222",
	},
	&Masquerade{
		Domain:    "lovewall-missdior.dior.com",
		IpAddress: "54.182.2.148",
	},
	&Masquerade{
		Domain:    "s-onetag.com",
		IpAddress: "52.222.134.188",
	},
	&Masquerade{
		Domain:    "ap1.whispir.com",
		IpAddress: "54.239.192.142",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "54.182.2.20",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.187",
	},
	&Masquerade{
		Domain:    "resources.licenses.adobe.com",
		IpAddress: "205.251.212.184",
	},
	&Masquerade{
		Domain:    "appgallery.huawei.com",
		IpAddress: "52.222.134.107",
	},
	&Masquerade{
		Domain:    "twitchcdn-shadow.net",
		IpAddress: "13.249.6.151",
	},
	&Masquerade{
		Domain:    "www.period-calendar.com",
		IpAddress: "205.251.212.239",
	},
	&Masquerade{
		Domain:    "www.hungama.com",
		IpAddress: "52.222.129.70",
	},
	&Masquerade{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "54.239.130.229",
	},
	&Masquerade{
		Domain:    "api.sandbox.repayonline.com",
		IpAddress: "99.86.2.87",
	},
	&Masquerade{
		Domain:    "siftscience.com",
		IpAddress: "13.249.2.64",
	},
	&Masquerade{
		Domain:    "ssi.servicestream.com.au",
		IpAddress: "143.204.6.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.131.24",
	},
	&Masquerade{
		Domain:    "smile.amazon.de",
		IpAddress: "13.224.5.77",
	},
	&Masquerade{
		Domain:    "internal-weedmaps.com",
		IpAddress: "54.182.2.85",
	},
	&Masquerade{
		Domain:    "cdn.discounttire.com",
		IpAddress: "99.86.0.211",
	},
	&Masquerade{
		Domain:    "siftscience.com",
		IpAddress: "99.86.1.124",
	},
	&Masquerade{
		Domain:    "cdn-legacy.contentful.com",
		IpAddress: "13.35.5.14",
	},
	&Masquerade{
		Domain:    "bamsec.com",
		IpAddress: "13.249.5.44",
	},
	&Masquerade{
		Domain:    "pay.2go.com",
		IpAddress: "205.251.212.16",
	},
	&Masquerade{
		Domain:    "www.connectwise.com",
		IpAddress: "52.222.132.120",
	},
	&Masquerade{
		Domain:    "s.salecycle.com",
		IpAddress: "52.222.135.8",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.89",
	},
	&Masquerade{
		Domain:    "www.gr-assets.com",
		IpAddress: "99.86.2.8",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.28",
	},
	&Masquerade{
		Domain:    "js-assets.aiv-cdn.net",
		IpAddress: "204.246.164.44",
	},
	&Masquerade{
		Domain:    "www.misumi.jp",
		IpAddress: "54.182.6.120",
	},
	&Masquerade{
		Domain:    "xgcpaa.com",
		IpAddress: "13.249.5.22",
	},
	&Masquerade{
		Domain:    "iproc.originenergy.com.au",
		IpAddress: "99.86.6.189",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-eu-central-1.dev.hydra.sophos.com",
		IpAddress: "204.246.178.167",
	},
	&Masquerade{
		Domain:    "public-rca-cloudstation-us-east-2.qa.hydra.sophos.com",
		IpAddress: "54.182.3.60",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.157",
	},
	&Masquerade{
		Domain:    "www.creditloan.com",
		IpAddress: "204.246.178.6",
	},
	&Masquerade{
		Domain:    "www.fp.ps.easebar.com",
		IpAddress: "143.204.6.213",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.192",
	},
	&Masquerade{
		Domain:    "oqs.amb.cybird.ne.jp",
		IpAddress: "52.222.131.144",
	},
	&Masquerade{
		Domain:    "static.amundi.com",
		IpAddress: "99.86.6.33",
	},
	&Masquerade{
		Domain:    "achievers.com",
		IpAddress: "13.35.6.128",
	},
	&Masquerade{
		Domain:    "www.sigalert.com",
		IpAddress: "13.224.7.5",
	},
	&Masquerade{
		Domain:    "static.yub-cdn.com",
		IpAddress: "13.249.2.12",
	},
	&Masquerade{
		Domain:    "boleto.sandbox.pagseguro.com.br",
		IpAddress: "54.182.2.149",
	},
	&Masquerade{
		Domain:    "gaijinent.com",
		IpAddress: "13.249.5.21",
	},
	&Masquerade{
		Domain:    "edwardsdoc.com",
		IpAddress: "13.35.4.166",
	},
	&Masquerade{
		Domain:    "www.goldspotmedia.com",
		IpAddress: "13.35.3.52",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.5.206",
	},
	&Masquerade{
		Domain:    "www.cookpad.com",
		IpAddress: "204.246.177.11",
	},
	&Masquerade{
		Domain:    "iot.eu-west-2.amazonaws.com",
		IpAddress: "99.86.2.91",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.133.105",
	},
	&Masquerade{
		Domain:    "www.lps.lottedfs.com",
		IpAddress: "13.224.5.27",
	},
	&Masquerade{
		Domain:    "www.canadamats.ca",
		IpAddress: "204.246.164.8",
	},
	&Masquerade{
		Domain:    "rca-upload-cloudstation-us-west-2.dev3.hydra.sophos.com",
		IpAddress: "204.246.177.29",
	},
	&Masquerade{
		Domain:    "polaris.lhinside.com",
		IpAddress: "143.204.6.108",
	},
	&Masquerade{
		Domain:    "behance.net",
		IpAddress: "13.35.6.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "99.84.4.2",
	},
	&Masquerade{
		Domain:    "uploads.skyhighnetworks.com",
		IpAddress: "13.224.5.118",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.224.2.61",
	},
	&Masquerade{
		Domain:    "www.milkvr.rocks",
		IpAddress: "99.86.4.34",
	},
	&Masquerade{
		Domain:    "freight.amazon.com",
		IpAddress: "99.86.2.12",
	},
	&Masquerade{
		Domain:    "binance.com",
		IpAddress: "13.35.3.114",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.249.4.8",
	},
	&Masquerade{
		Domain:    "www.tosconfig.com",
		IpAddress: "99.86.0.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.1.103",
	},
	&Masquerade{
		Domain:    "iot.ap-northeast-1.amazonaws.com",
		IpAddress: "13.224.5.194",
	},
	&Masquerade{
		Domain:    "mobizen.com",
		IpAddress: "143.204.1.15",
	},
	&Masquerade{
		Domain:    "www.suezwatertechnologies.com",
		IpAddress: "143.204.2.210",
	},
	&Masquerade{
		Domain:    "www.iot.irobot.cn",
		IpAddress: "99.86.2.14",
	},
	&Masquerade{
		Domain:    "bglen.net",
		IpAddress: "13.224.5.60",
	},
	&Masquerade{
		Domain:    "read.amazon.com",
		IpAddress: "99.86.2.62",
	},
	&Masquerade{
		Domain:    "www.nmrodam.com",
		IpAddress: "54.239.192.179",
	},
	&Masquerade{
		Domain:    "netmarble.net",
		IpAddress: "54.182.0.164",
	},
}
