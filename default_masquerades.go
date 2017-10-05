package fronted

var DefaultTrustedCAs = []*CA{
	&CA{
		CommonName: "VeriSign Class 3 Public Primary Certification Authority - G5",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIE0zCCA7ugAwIBAgIQGNrRniZ96LtKIVjNzGs7SjANBgkqhkiG9w0BAQUFADCB\nyjELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL\nExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJp\nU2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxW\nZXJpU2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0\naG9yaXR5IC0gRzUwHhcNMDYxMTA4MDAwMDAwWhcNMzYwNzE2MjM1OTU5WjCByjEL\nMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQLExZW\nZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJpU2ln\nbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxWZXJp\nU2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9y\naXR5IC0gRzUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCvJAgIKXo1\nnmAMqudLO07cfLw8RRy7K+D+KQL5VwijZIUVJ/XxrcgxiV0i6CqqpkKzj/i5Vbex\nt0uz/o9+B1fs70PbZmIVYc9gDaTY3vjgw2IIPVQT60nKWVSFJuUrjxuf6/WhkcIz\nSdhDY2pSS9KP6HBRTdGJaXvHcPaz3BJ023tdS1bTlr8Vd6Gw9KIl8q8ckmcY5fQG\nBO+QueQA5N06tRn/Arr0PO7gi+s3i+z016zy9vA9r911kTMZHRxAy3QkGSGT2RT+\nrCpSx4/VBEnkjWNHiDxpg8v+R70rfk/Fla4OndTRQ8Bnc+MUCH7lP59zuDMKz10/\nNIeWiu5T6CUVAgMBAAGjgbIwga8wDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8E\nBAMCAQYwbQYIKwYBBQUHAQwEYTBfoV2gWzBZMFcwVRYJaW1hZ2UvZ2lmMCEwHzAH\nBgUrDgMCGgQUj+XTGoasjY5rw8+AatRIGCx7GS4wJRYjaHR0cDovL2xvZ28udmVy\naXNpZ24uY29tL3ZzbG9nby5naWYwHQYDVR0OBBYEFH/TZafC3ey78DAJ80M5+gKv\nMzEzMA0GCSqGSIb3DQEBBQUAA4IBAQCTJEowX2LP2BqYLz3q3JktvXf2pXkiOOzE\np6B4Eq1iDkVwZMXnl2YtmAl+X6/WzChl8gGqCBpH3vn5fJJaCGkgDdk+bW48DW7Y\n5gaRQBi5+MHt39tBquCWIMnNZBU4gcmU7qKEKQsTb47bDN0lAtukixlE0kF6BWlK\nWE9gyn6CagsCqiUXObXbf+eEZSqVir2G3l6BFoMtEMze/aiCKm0oHw0LxOXnGiYZ\n4fQRbxC1lfznQgUy286dUV4otp6F01vvpX1FQHKOtw5rDgb7MzVIcbidJ4vEZV8N\nhnacRHr2lVz2XTIIM6RUthg/aFzyQkqFOFSDX9HoLPKsEdao7WNq\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "Starfield Services Root Certificate Authority - G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIID7zCCAtegAwIBAgIBADANBgkqhkiG9w0BAQsFADCBmDELMAkGA1UEBhMCVVMx\nEDAOBgNVBAgTB0FyaXpvbmExEzARBgNVBAcTClNjb3R0c2RhbGUxJTAjBgNVBAoT\nHFN0YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xOzA5BgNVBAMTMlN0YXJmaWVs\nZCBTZXJ2aWNlcyBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eSAtIEcyMB4XDTA5\nMDkwMTAwMDAwMFoXDTM3MTIzMTIzNTk1OVowgZgxCzAJBgNVBAYTAlVTMRAwDgYD\nVQQIEwdBcml6b25hMRMwEQYDVQQHEwpTY290dHNkYWxlMSUwIwYDVQQKExxTdGFy\nZmllbGQgVGVjaG5vbG9naWVzLCBJbmMuMTswOQYDVQQDEzJTdGFyZmllbGQgU2Vy\ndmljZXMgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgLSBHMjCCASIwDQYJKoZI\nhvcNAQEBBQADggEPADCCAQoCggEBANUMOsQq+U7i9b4Zl1+OiFOxHz/Lz58gE20p\nOsgPfTz3a3Y4Y9k2YKibXlwAgLIvWX/2h/klQ4bnaRtSmpDhcePYLQ1Ob/bISdm2\n8xpWriu2dBTrz/sm4xq6HZYuajtYlIlHVv8loJNwU4PahHQUw2eeBGg6345AWh1K\nTs9DkTvnVtYAcMtS7nt9rjrnvDH5RfbCYM8TWQIrgMw0R9+53pBlbQLPLJGmpufe\nhRhJfGZOozptqbXuNC66DQO4M99H67FrjSXZm86B0UVGMpZwh94CDklDhbZsc7tk\n6mFBrMnUVN+HL8cisibMn1lUaJ/8viovxFUcdUBgF4UCVTmLfwUCAwEAAaNCMEAw\nDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFJxfAN+q\nAdcwKziIorhtSpzyEZGDMA0GCSqGSIb3DQEBCwUAA4IBAQBLNqaEd2ndOxmfZyMI\nbw5hyf2E3F/YNoHN2BtBLZ9g3ccaaNnRbobhiCPPE95Dz+I0swSdHynVv/heyNXB\nve6SbzJ08pGCL72CQnqtKrcgfU28elUSwhXqvfdqlS5sdJ/PHLTyxQGjhdByPq1z\nqwubdQxtRbeOlKyWN7Wg0I8VRw7j6IPdj/3vQQF3zCepYoUz8jcI73HPdwbeyBkd\niEDPfUYd/x7H4c7/I9vG+o1VTqkC50cRRj70/b17KSa7qWFiNyi2LSr2EIZkyXCn\n0q23KXB56jzaYyWf/Wi3MOxw+3WKt21gZ7IeyLnp2KhvAotnDU0mV3HaIPzBSlCN\nsSi6\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "Go Daddy Root Certificate Authority - G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIBADANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMCVVMx\nEDAOBgNVBAgTB0FyaXpvbmExEzARBgNVBAcTClNjb3R0c2RhbGUxGjAYBgNVBAoT\nEUdvRGFkZHkuY29tLCBJbmMuMTEwLwYDVQQDEyhHbyBEYWRkeSBSb290IENlcnRp\nZmljYXRlIEF1dGhvcml0eSAtIEcyMB4XDTA5MDkwMTAwMDAwMFoXDTM3MTIzMTIz\nNTk1OVowgYMxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdBcml6b25hMRMwEQYDVQQH\nEwpTY290dHNkYWxlMRowGAYDVQQKExFHb0RhZGR5LmNvbSwgSW5jLjExMC8GA1UE\nAxMoR28gRGFkZHkgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgLSBHMjCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL9xYgjx+lk09xvJGKP3gElY6SKD\nE6bFIEMBO4Tx5oVJnyfq9oQbTqC023CYxzIBsQU+B07u9PpPL1kwIuerGVZr4oAH\n/PMWdYA5UXvl+TW2dE6pjYIT5LY/qQOD+qK+ihVqf94Lw7YZFAXK6sOoBJQ7Rnwy\nDfMAZiLIjWltNowRGLfTshxgtDj6AozO091GB94KPutdfMh8+7ArU6SSYmlRJQVh\nGkSBjCypQ5Yj36w6gZoOKcUcqeldHraenjAKOc7xiID7S13MMuyFYkMlNAJWJwGR\ntDtwKj9useiciAF9n9T521NtYJ2/LOdYq7hfRvzOxBsDPAnrSTFcaUaz4EcCAwEA\nAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYE\nFDqahQcQZyi27/a9BUFuIMGU2g/eMA0GCSqGSIb3DQEBCwUAA4IBAQCZ21151fmX\nWWcDYfF+OwYxdS2hII5PZYe096acvNjpL9DbWu7PdIxztDhC2gV7+AJ1uP2lsdeu\n9tfeE8tTEH6KRtGX+rcuKxGrkLAngPnon1rpN5+r5N9ss4UXnT3ZJE95kTXWXwTr\ngIOrmgIttRD02JDHBHNA7XIloKmf7J6raBKZV8aPEjoJpL1E/QYVN8Gb5DKj7Tjo\n2GTzLH4U/ALqn83/B2gX2yKQOC16jdFU8WnjXzPKej17CuPKf1855eJ1usV2GDPO\nLPAvTK33sefOT6jEm0pUBsV/fdUID+Ic/n4XuKxe9tQWskMJDE32p2u0mYRlynqI\n4uJEvlz36hz1\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "GeoTrust Global CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDVDCCAjygAwIBAgIDAjRWMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT\nMRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i\nYWwgQ0EwHhcNMDIwNTIxMDQwMDAwWhcNMjIwNTIxMDQwMDAwWjBCMQswCQYDVQQG\nEwJVUzEWMBQGA1UEChMNR2VvVHJ1c3QgSW5jLjEbMBkGA1UEAxMSR2VvVHJ1c3Qg\nR2xvYmFsIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2swYYzD9\n9BcjGlZ+W988bDjkcbd4kdS8odhM+KhDtgPpTSEHCIjaWC9mOSm9BXiLnTjoBbdq\nfnGk5sRgprDvgOSJKA+eJdbtg/OtppHHmMlCGDUUna2YRpIuT8rxh0PBFpVXLVDv\niS2Aelet8u5fa9IAjbkU+BQVNdnARqN7csiRv8lVK83Qlz6cJmTM386DGXHKTubU\n1XupGc1V3sjs0l44U+VcT4wt/lAjNvxm5suOpDkZALeVAjmRCw7+OC7RHQWa9k0+\nbw8HHa8sHo9gOeL6NlMTOdReJivbPagUvTLrGAMoUgRx5aszPeE4uwc2hGKceeoW\nMPRfwCvocWvk+QIDAQABo1MwUTAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTA\nephojYn7qwVkDBF9qn1luMrMTjAfBgNVHSMEGDAWgBTAephojYn7qwVkDBF9qn1l\nuMrMTjANBgkqhkiG9w0BAQUFAAOCAQEANeMpauUvXVSOKVCUn5kaFOSPeCpilKIn\nZ57QzxpeR+nBsqTP3UEaBU6bS+5Kb1VSsyShNwrrZHYqLizz/Tt1kL/6cdjHPTfS\ntQWVYrmm3ok9Nns4d0iXrKYgjy6myQzCsplFAMfOEVEiIuCl6rYVSAlk6l5PdPcF\nPseKUgzbFbS9bZvlxrFUaKnjaZC2mqUPuLk/IH2uSrW4nOQdtqvmlKXBx4Ot2/Un\nhw4EbNX/3aBd7YdStysVAq45pmp06drE57xNNB6pXE0zX5IJL4hmXXeXxx12E6nV\n5fEWCRE11azbJHFwLJhWC9kXtNHjUStedejV0NxPNO3CBWaAocvmMw==\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "DigiCert Global Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\nQTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB\nCSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97\nnh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt\n43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P\nT19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4\ngdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO\nBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR\nTLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw\nDQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr\nhMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg\n06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF\nPnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls\nYSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk\nCAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "COMODO RSA Certification Authority",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIF2DCCA8CgAwIBAgIQTKr5yttjb+Af907YWwOGnTANBgkqhkiG9w0BAQwFADCB\nhTELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G\nA1UEBxMHU2FsZm9yZDEaMBgGA1UEChMRQ09NT0RPIENBIExpbWl0ZWQxKzApBgNV\nBAMTIkNPTU9ETyBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTAwMTE5\nMDAwMDAwWhcNMzgwMTE4MjM1OTU5WjCBhTELMAkGA1UEBhMCR0IxGzAZBgNVBAgT\nEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4GA1UEBxMHU2FsZm9yZDEaMBgGA1UEChMR\nQ09NT0RPIENBIExpbWl0ZWQxKzApBgNVBAMTIkNPTU9ETyBSU0EgQ2VydGlmaWNh\ndGlvbiBBdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCR\n6FSS0gpWsawNJN3Fz0RndJkrN6N9I3AAcbxT38T6KhKPS38QVr2fcHK3YX/JSw8X\npz3jsARh7v8Rl8f0hj4K+j5c+ZPmNHrZFGvnnLOFoIJ6dq9xkNfs/Q36nGz637CC\n9BR++b7Epi9Pf5l/tfxnQ3K9DADWietrLNPtj5gcFKt+5eNu/Nio5JIk2kNrYrhV\n/erBvGy2i/MOjZrkm2xpmfh4SDBF1a3hDTxFYPwyllEnvGfDyi62a+pGx8cgoLEf\nZd5ICLqkTqnyg0Y3hOvozIFIQ2dOciqbXL1MGyiKXCJ7tKuY2e7gUYPDCUZObT6Z\n+pUX2nwzV0E8jVHtC7ZcryxjGt9XyD+86V3Em69FmeKjWiS0uqlWPc9vqv9JWL7w\nqP/0uK3pN/u6uPQLOvnoQ0IeidiEyxPx2bvhiWC4jChWrBQdnArncevPDt09qZah\nSL0896+1DSJMwBGB7FY79tOi4lu3sgQiUpWAk2nojkxl8ZEDLXB0AuqLZxUpaVIC\nu9ffUGpVRr+goyhhf3DQw6KqLCGqR84onAZFdr+CGCe01a60y1Dma/RMhnEw6abf\nFobg2P9A3fvQQoh/ozM6LlweQRGBY84YcWsr7KaKtzFcOmpH4MN5WdYgGq/yapiq\ncrxXStJLnbsQ/LBMQeXtHT1eKJ2czL+zUdqnR+WEUwIDAQABo0IwQDAdBgNVHQ4E\nFgQUu69+Aj36pvE8hI6t7jiY7NkyMtQwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB\n/wQFMAMBAf8wDQYJKoZIhvcNAQEMBQADggIBAArx1UaEt65Ru2yyTUEUAJNMnMvl\nwFTPoCWOAvn9sKIN9SCYPBMtrFaisNZ+EZLpLrqeLppysb0ZRGxhNaKatBYSaVqM\n4dc+pBroLwP0rmEdEBsqpIt6xf4FpuHA1sj+nq6PK7o9mfjYcwlYRm6mnPTXJ9OV\n2jeDchzTc+CiR5kDOF3VSXkAKRzH7JsgHAckaVd4sjn8OoSgtZx8jb8uk2Intzna\nFxiuvTwJaP+EmzzV1gsD41eeFPfR60/IvYcjt7ZJQ3mFXLrrkguhxuhoqEwWsRqZ\nCuhTLJK7oQkYdQxlqHvLI7cawiiFwxv/0Cti76R7CZGYZ4wUAc1oBmpjIXUDgIiK\nboHGhfKppC3n9KUkEEeDys30jXlYsQab5xoq2Z0B15R97QNKyvDb6KkBPvVWmcke\njkk9u+UJueBPSZI9FoJAzMxZxuY67RIuaTxslbH9qh17f4a+Hg4yRvv7E491f0yL\nS0Zj/gA0QHDBw7mh3aZw4gSzQbzpgJHqZJx64SIDqZxubw5lT2yHh17zbqD5daWb\nQOhTsiedSrnAdyGN/4fy3ryM7xfft0kL0fJuMAsaDk527RH89elWsn2/x20Kk4yl\n0MC2Hb46TpSi125sC8KKfPog88Tk5c0NqMuRkrF8hey1FGlmDoLnzc7ILaZRfyHB\nNVOFBkpdn627G190\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "DigiCert High Assurance EV Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIQAqxcJmoLQJuPC3nyrkYldzANBgkqhkiG9w0BAQUFADBs\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j\nZSBFViBSb290IENBMB4XDTA2MTExMDAwMDAwMFoXDTMxMTExMDAwMDAwMFowbDEL\nMAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3\nLmRpZ2ljZXJ0LmNvbTErMCkGA1UEAxMiRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug\nRVYgUm9vdCBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMbM5XPm\n+9S75S0tMqbf5YE/yc0lSbZxKsPVlDRnogocsF9ppkCxxLeyj9CYpKlBWTrT3JTW\nPNt0OKRKzE0lgvdKpVMSOO7zSW1xkX5jtqumX8OkhPhPYlG++MXs2ziS4wblCJEM\nxChBVfvLWokVfnHoNb9Ncgk9vjo4UFt3MRuNs8ckRZqnrG0AFFoEt7oT61EKmEFB\nIk5lYYeBQVCmeVyJ3hlKV9Uu5l0cUyx+mM0aBhakaHPQNAQTXKFx01p8VdteZOE3\nhzBWBOURtCmAEvF5OYiiAhF8J2a3iLd48soKqDirCmTCv2ZdlYTBoSUeh10aUAsg\nEsxBu24LUTi4S8sCAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQF\nMAMBAf8wHQYDVR0OBBYEFLE+w2kD+L9HAdSYJhoIAu9jZCvDMB8GA1UdIwQYMBaA\nFLE+w2kD+L9HAdSYJhoIAu9jZCvDMA0GCSqGSIb3DQEBBQUAA4IBAQAcGgaX3Nec\nnzyIZgYIVyHbIUf4KmeqvxgydkAQV8GK83rZEWWONfqe/EW1ntlMMUu4kehDLI6z\neM7b41N5cdblIZQB2lWHmiRk9opmzN6cN82oNLFpmyPInngiK3BD41VHMWEZ71jF\nhS9OMPagMRYjyOfiZRYzy78aG6A9+MpeizGLYAiJLQwGXFK3xPkKmNEVX58Svnw2\nYzi9RKR/5CYrCsSXaQ3pjOLAEFe4yHYSkVXySGnYvCoCWw9E1CAx2/S6cCZdkGCe\nvEsXCS+0yx5DaMkHJ8HSXPfqIbloEpw8nL+e/IBcm2PN7EeqJSdnoDfzAIJ9VNep\n+OkuE6N36B9K\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "GlobalSign Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG\nA1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv\nb3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw\nMDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i\nYWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT\naWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ\njc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp\nxy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp\n1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG\nsnUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ\nU26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8\n9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E\nBTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B\nAQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz\nyj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE\n38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP\nAbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad\nDKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME\nHMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==\n-----END CERTIFICATE-----\n",
	},
}

var DefaultCloudfrontMasquerades = []*Masquerade{
	&Masquerade{
		Domain:    "Images-na.ssl-images-amazon.com",
		IpAddress: "54.192.9.227",
	},
	&Masquerade{
		Domain:    "abcmouse.tw",
		IpAddress: "205.251.206.4",
	},
	&Masquerade{
		Domain:    "adbephotos-stage.com",
		IpAddress: "54.182.7.83",
	},
	&Masquerade{
		Domain:    "adtdp.com",
		IpAddress: "54.192.6.182",
	},
	&Masquerade{
		Domain:    "altium.com",
		IpAddress: "205.251.206.142",
	},
	&Masquerade{
		Domain:    "amazon.co.uk",
		IpAddress: "13.32.7.60",
	},
	&Masquerade{
		Domain:    "amazon.co.uk",
		IpAddress: "54.182.7.222",
	},
	&Masquerade{
		Domain:    "amazon.com",
		IpAddress: "54.230.2.31",
	},
	&Masquerade{
		Domain:    "amazon.com",
		IpAddress: "13.32.10.225",
	},
	&Masquerade{
		Domain:    "amazon.de",
		IpAddress: "54.230.2.127",
	},
	&Masquerade{
		Domain:    "amazon.es",
		IpAddress: "205.251.206.66",
	},
	&Masquerade{
		Domain:    "amazon.fr",
		IpAddress: "54.239.200.212",
	},
	&Masquerade{
		Domain:    "api.eab.com",
		IpAddress: "52.84.4.199",
	},
	&Masquerade{
		Domain:    "api.eab.com",
		IpAddress: "216.137.52.24",
	},
	&Masquerade{
		Domain:    "api.starmakerstudios.com",
		IpAddress: "54.192.0.54",
	},
	&Masquerade{
		Domain:    "api.starmakerstudios.com",
		IpAddress: "216.137.52.151",
	},
	&Masquerade{
		Domain:    "api.starmakerstudios.com",
		IpAddress: "205.251.206.107",
	},
	&Masquerade{
		Domain:    "api.starmakerstudios.com",
		IpAddress: "52.84.8.180",
	},
	&Masquerade{
		Domain:    "api.starmakerstudios.com",
		IpAddress: "54.192.4.182",
	},
	&Masquerade{
		Domain:    "apilivingsocial.co.uk",
		IpAddress: "13.32.14.198",
	},
	&Masquerade{
		Domain:    "apilivingsocial.co.uk",
		IpAddress: "13.32.2.198",
	},
	&Masquerade{
		Domain:    "appchoose.io",
		IpAddress: "54.182.6.103",
	},
	&Masquerade{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "13.32.7.152",
	},
	&Masquerade{
		Domain:    "appsdownload2.hkjc.com",
		IpAddress: "54.182.7.16",
	},
	&Masquerade{
		Domain:    "assets.tumblr.com",
		IpAddress: "52.84.8.71",
	},
	&Masquerade{
		Domain:    "batch.eu-west-1.amazonaws.com",
		IpAddress: "54.192.6.52",
	},
	&Masquerade{
		Domain:    "batch.eu-west-2.amazonaws.com",
		IpAddress: "13.32.7.120",
	},
	&Masquerade{
		Domain:    "batch.eu-west-2.amazonaws.com",
		IpAddress: "205.251.212.128",
	},
	&Masquerade{
		Domain:    "batch.eu-west-2.amazonaws.com",
		IpAddress: "216.137.36.225",
	},
	&Masquerade{
		Domain:    "berlin.buuteeq.com",
		IpAddress: "204.246.164.180",
	},
	&Masquerade{
		Domain:    "bigpanda.io",
		IpAddress: "54.182.0.51",
	},
	&Masquerade{
		Domain:    "buttonhub.com",
		IpAddress: "54.230.6.180",
	},
	&Masquerade{
		Domain:    "buzzfeed.com",
		IpAddress: "54.182.0.151",
	},
	&Masquerade{
		Domain:    "camp-fire.jp",
		IpAddress: "13.32.5.168",
	},
	&Masquerade{
		Domain:    "cbtalentnetwork.com",
		IpAddress: "216.137.52.171",
	},
	&Masquerade{
		Domain:    "cdn-test.worldpay.com",
		IpAddress: "216.137.52.87",
	},
	&Masquerade{
		Domain:    "cdn.admin.staging.checkmatenext.com",
		IpAddress: "54.239.130.137",
	},
	&Masquerade{
		Domain:    "cdn.concordnow.com",
		IpAddress: "54.192.15.160",
	},
	&Masquerade{
		Domain:    "cdn.fukuyamamasaharu.com",
		IpAddress: "52.84.13.193",
	},
	&Masquerade{
		Domain:    "cdn.getgo.com",
		IpAddress: "205.251.206.9",
	},
	&Masquerade{
		Domain:    "cdn.gotomeet.at",
		IpAddress: "13.32.9.205",
	},
	&Masquerade{
		Domain:    "cdn.medallia.com",
		IpAddress: "54.230.8.114",
	},
	&Masquerade{
		Domain:    "cdn.medallia.com",
		IpAddress: "13.32.5.183",
	},
	&Masquerade{
		Domain:    "cdn.mozilla.net",
		IpAddress: "54.230.4.28",
	},
	&Masquerade{
		Domain:    "cdn.mozilla.net",
		IpAddress: "54.192.9.39",
	},
	&Masquerade{
		Domain:    "cdn.mozilla.net",
		IpAddress: "13.32.12.197",
	},
	&Masquerade{
		Domain:    "cdn.shptrn.com",
		IpAddress: "216.137.52.105",
	},
	&Masquerade{
		Domain:    "cdn.wk-dev.wdesk.org",
		IpAddress: "52.84.16.19",
	},
	&Masquerade{
		Domain:    "cdnint.fca.telematics.net",
		IpAddress: "13.32.14.157",
	},
	&Masquerade{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "216.137.52.250",
	},
	&Masquerade{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "204.246.164.70",
	},
	&Masquerade{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "205.251.212.72",
	},
	&Masquerade{
		Domain:    "cdnsta.fca.telematics.net",
		IpAddress: "54.182.7.41",
	},
	&Masquerade{
		Domain:    "chaturbate.com",
		IpAddress: "13.32.12.96",
	},
	&Masquerade{
		Domain:    "chauffeur-prive.com",
		IpAddress: "204.246.164.71",
	},
	&Masquerade{
		Domain:    "chemistwarehouse.com",
		IpAddress: "54.182.5.77",
	},
	&Masquerade{
		Domain:    "chiwawa.one",
		IpAddress: "54.182.0.99",
	},
	&Masquerade{
		Domain:    "ciproductionportal.production.vf-leap.com",
		IpAddress: "216.137.45.196",
	},
	&Masquerade{
		Domain:    "ciproductionportal.production.vf-leap.com",
		IpAddress: "13.32.6.244",
	},
	&Masquerade{
		Domain:    "client.wc.ue1.app.chime.aws",
		IpAddress: "216.137.45.224",
	},
	&Masquerade{
		Domain:    "clients.chime.aws",
		IpAddress: "205.251.206.125",
	},
	&Masquerade{
		Domain:    "clocktree.com",
		IpAddress: "54.230.6.14",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.93",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.141",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.135",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.119",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.145",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.151",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.251",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.174",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.47",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.190",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.205",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.228",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.59",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.219",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.189",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.237",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.173",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.123",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.64",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.65",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.106",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.52",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.215",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.131",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.214",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.121",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.176",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.98",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.72",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.75",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.142",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.48",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.28",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.119",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.85",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.101",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.150",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.12",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.31",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.179",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.135",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.129",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.144",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.235",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.40",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.245",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.240",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.29",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.200",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.108",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.82",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.86",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.159",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.181",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.7",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.219",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.198",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.143",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.132",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.214",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.117",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.23",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.122",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.130",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.249",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.157",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.66",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.149",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.173",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.75",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.20",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.113",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.132",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.91",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.42",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.79",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.99",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.239",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.149",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.94",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.37",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.140",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.164",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.39",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.209",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.212",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.208",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.167",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.13",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.48",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.69",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.131",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.207",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.188",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.64",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.13",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.123",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.110",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.18",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.254",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.202",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.214",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.228",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.94",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.163",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.208",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.22",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.135",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.20",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.42",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.161",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.52",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.5",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.83",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.164",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.236",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.173",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.213",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.194",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.43",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.250",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.171",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.225",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.154",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.29",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.67",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.33",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.204",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.33",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.94",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.174",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.161",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.219",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.182",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.9",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.64",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.51",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.248",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.84",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.69",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.250",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.73",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.222",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.75",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.177",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.160",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.147",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.30",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.206",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.57",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.39",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.40",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.80",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.224",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.109",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.28",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.104",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.130",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.136",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.94",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.82",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.39",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.84",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.243",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.114",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.235",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.143",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.177",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.238",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.140",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.112",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.200",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.123",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.76",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.133",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.120",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.240",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.244",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.39",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.239",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.167",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.116",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.47",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.78",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.220",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.235",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.141",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.102",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.12",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.37",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.157",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.252",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.96",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.93",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.76",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.150",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.221",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.142",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.116",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.90",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.241",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.95",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.30",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.233",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.210",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.121",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.17",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.54",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.114",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.236",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.115",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.22",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.149",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.102",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.172",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.177",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.74",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.24",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.123",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.234",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.22",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.227",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.189",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.8",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.148",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.58",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.195",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.205",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.229",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.67",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.34",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.147",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.14",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.8",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.11",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.37",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.26",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.96",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.200",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.105",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.134",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.146",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.39",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.210",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.162",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.159",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.17",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.145",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.172",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.111",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.150",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.217",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.248",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.158",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.83",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.41",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.40",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.96",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.130",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.117",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.205",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.15",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.160",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.47",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.140",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.28",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.206",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.176",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.10",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.79",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.254",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.236",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.52",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.229",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.98",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.219",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.197",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.125",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.96",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.23",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.171",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.249",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.138",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.223",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.84",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.118",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.116",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.247",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.212",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.123",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.158",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.13",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.200",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.173",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.230",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.108",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.5",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.190",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.178",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.121",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.68",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.34",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.239",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.88",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.147",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.232",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.105",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.131",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.210",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.144",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.95",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.91",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.90",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.35",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.245",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.36",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.157",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.24",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.37",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.41",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.119",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.12",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.138",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.65",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.198",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.111",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.117",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.21",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.83",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.109",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.31",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.199",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.121",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.194",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.104",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.211",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.96",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.179",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.68",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.252",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.116",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.89",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.127",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.201",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.50",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.98",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.61",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.54",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.202",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.187",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.226",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.211",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.232",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.41",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.241",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.149",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.172",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.59",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.93",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.28",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.238",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.142",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.156",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.95",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.19",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.27",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.4",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.230",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.26",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.101",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.87",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.65",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.223",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "216.137.52.207",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.233",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.194",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.11",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.27",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.33",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.229",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.69",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.102",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.79",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.73",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.76",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.47",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.43",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.168",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.60",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.152",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.197",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.79",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.34",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.129",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.226",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.25",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.188",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.78",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.192",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.22",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.70",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.35",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.75",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.122",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.251",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.207",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.186",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.40",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.209",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.198",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.147",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.73",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.115",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.95",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.105",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.26",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.174",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.129",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.227",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.118",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.172",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.221",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.134",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.234",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.141",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.150",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.178",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.104",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.197",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.83",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.146",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.46",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.124",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.112",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.107",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.244",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.217",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.78",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.48",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.181",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.171",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.124",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.202",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.216",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.10",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.12",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.223",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.249",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.213",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.138",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.81",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.221",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.241",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.170",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.162",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.151",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.54",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.246",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.242",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.161",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.89",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.87",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.168",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.180",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.111",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.209",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.248",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.7",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.112",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.242",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.244",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.245",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.224",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.71",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.29",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.244",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.215",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.148",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.87",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.185",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.55",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.224",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.224",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.231",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.66",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.128",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.32",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.208",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.83",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.246",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.169",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.154",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.127",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.11",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.143",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.34",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.54",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.108",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.113",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.222",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.252",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.193",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.152",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.149.6",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.100",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.196",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.124",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.42",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.119",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.9",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.183",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.166",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.137",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.8",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.146",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.5",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.184",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.153",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.64",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.202",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.204",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.182",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.213",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.16",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.152",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.235",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.168",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.205",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.239.192.4",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.157",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.129",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.4.183",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.250",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.115",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.173",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.182",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.0.37",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.174",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.176",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.112",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.22.86",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.182.2.88",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.92",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.120",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.132.202",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.222",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.5.113",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.7.90",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.175",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.84.10.159",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.167",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.222",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.24",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.15.189",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.172",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.20.188",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "52.222.136.150",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.16.50",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "13.32.1.69",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.230.7.195",
	},
	&Masquerade{
		Domain:    "cloudfront.net",
		IpAddress: "54.192.20.127",
	},
	&Masquerade{
		Domain:    "cloudfront.rebirthdrive.jp",
		IpAddress: "54.182.0.95",
	},
	&Masquerade{
		Domain:    "cloudfront.rebirthdrive.jp",
		IpAddress: "205.251.212.14",
	},
	&Masquerade{
		Domain:    "cmcm.com",
		IpAddress: "13.32.10.125",
	},
	&Masquerade{
		Domain:    "cmix.com",
		IpAddress: "54.182.0.196",
	},
	&Masquerade{
		Domain:    "cms-beta.us-west-2.gg.iot.amazonaws.com",
		IpAddress: "52.84.19.124",
	},
	&Masquerade{
		Domain:    "cms.us-west-2.gg.iot.amazonaws.com",
		IpAddress: "52.84.20.141",
	},
	&Masquerade{
		Domain:    "consumerreportscdn.org",
		IpAddress: "205.251.212.68",
	},
	&Masquerade{
		Domain:    "consumerreportscdn.org",
		IpAddress: "54.182.5.90",
	},
	&Masquerade{
		Domain:    "contents.truefighting.bnsv.ca",
		IpAddress: "54.182.7.82",
	},
	&Masquerade{
		Domain:    "cop.intuitcdn.net",
		IpAddress: "205.251.212.186",
	},
	&Masquerade{
		Domain:    "cop.intuitcdn.net",
		IpAddress: "13.32.5.239",
	},
	&Masquerade{
		Domain:    "cpdcdn.officedepot.com",
		IpAddress: "54.182.5.98",
	},
	&Masquerade{
		Domain:    "craftsy.com",
		IpAddress: "216.137.52.137",
	},
	&Masquerade{
		Domain:    "csp.infoblox.com",
		IpAddress: "54.230.4.190",
	},
	&Masquerade{
		Domain:    "democrats.org",
		IpAddress: "54.192.8.184",
	},
	&Masquerade{
		Domain:    "democrats.org",
		IpAddress: "205.251.206.117",
	},
	&Masquerade{
		Domain:    "devcms-www.lifelock.com",
		IpAddress: "204.246.164.29",
	},
	&Masquerade{
		Domain:    "dispatch.me",
		IpAddress: "54.182.7.244",
	},
	&Masquerade{
		Domain:    "dl.ubnt.com",
		IpAddress: "54.182.5.17",
	},
	&Masquerade{
		Domain:    "domuso.com",
		IpAddress: "54.182.5.112",
	},
	&Masquerade{
		Domain:    "download.70mai.asia",
		IpAddress: "54.239.142.125",
	},
	&Masquerade{
		Domain:    "download.70mai.asia",
		IpAddress: "54.182.0.111",
	},
	&Masquerade{
		Domain:    "dwnld.filecatalogue.com",
		IpAddress: "54.182.0.179",
	},
	&Masquerade{
		Domain:    "elo7.com.br",
		IpAddress: "54.192.13.100",
	},
	&Masquerade{
		Domain:    "empowernetwork.com",
		IpAddress: "54.192.6.25",
	},
	&Masquerade{
		Domain:    "envs.nbcd.co",
		IpAddress: "216.137.52.172",
	},
	&Masquerade{
		Domain:    "ext.app-cloud.jp",
		IpAddress: "52.84.18.67",
	},
	&Masquerade{
		Domain:    "fandays.jp",
		IpAddress: "54.182.0.29",
	},
	&Masquerade{
		Domain:    "fanduel.com",
		IpAddress: "204.246.164.12",
	},
	&Masquerade{
		Domain:    "fareoffice.com",
		IpAddress: "52.84.8.9",
	},
	&Masquerade{
		Domain:    "fcr.freecreditreport.com",
		IpAddress: "52.84.13.225",
	},
	&Masquerade{
		Domain:    "fcs.freecreditscore.com",
		IpAddress: "54.230.6.132",
	},
	&Masquerade{
		Domain:    "flipagram.com",
		IpAddress: "54.182.6.151",
	},
	&Masquerade{
		Domain:    "flowaccount.com",
		IpAddress: "13.32.5.162",
	},
	&Masquerade{
		Domain:    "flowaccount.com",
		IpAddress: "54.182.0.133",
	},
	&Masquerade{
		Domain:    "foxsportsgo.com",
		IpAddress: "204.246.164.16",
	},
	&Masquerade{
		Domain:    "freshdesk.com",
		IpAddress: "13.32.13.254",
	},
	&Masquerade{
		Domain:    "gameiom.com",
		IpAddress: "54.182.5.72",
	},
	&Masquerade{
		Domain:    "gamma.us-west-2.iot.amazonaws.com",
		IpAddress: "54.192.4.33",
	},
	&Masquerade{
		Domain:    "get.com",
		IpAddress: "54.182.0.132",
	},
	&Masquerade{
		Domain:    "getstream.io",
		IpAddress: "216.137.52.55",
	},
	&Masquerade{
		Domain:    "gfycat.com",
		IpAddress: "216.137.52.148",
	},
	&Masquerade{
		Domain:    "glbl.adlegend.com",
		IpAddress: "54.192.1.100",
	},
	&Masquerade{
		Domain:    "gopro.com",
		IpAddress: "52.84.4.136",
	},
	&Masquerade{
		Domain:    "gumbuya.net",
		IpAddress: "54.182.6.232",
	},
	&Masquerade{
		Domain:    "hbonow.com",
		IpAddress: "54.182.6.35",
	},
	&Masquerade{
		Domain:    "healthgrades.com",
		IpAddress: "54.230.6.247",
	},
	&Masquerade{
		Domain:    "heartbeat-stage.v0.maxdome.cloud",
		IpAddress: "54.182.6.88",
	},
	&Masquerade{
		Domain:    "homes.co.jp",
		IpAddress: "54.192.6.73",
	},
	&Masquerade{
		Domain:    "honey.is",
		IpAddress: "216.137.52.80",
	},
	&Masquerade{
		Domain:    "iflix.com",
		IpAddress: "54.182.0.213",
	},
	&Masquerade{
		Domain:    "imeet.com",
		IpAddress: "52.84.1.136",
	},
	&Masquerade{
		Domain:    "imeet.net",
		IpAddress: "216.137.52.84",
	},
	&Masquerade{
		Domain:    "img.angelmaster.jp",
		IpAddress: "204.246.164.38",
	},
	&Masquerade{
		Domain:    "info.cookpad.com",
		IpAddress: "54.182.0.178",
	},
	&Masquerade{
		Domain:    "inform.com",
		IpAddress: "54.182.7.122",
	},
	&Masquerade{
		Domain:    "insighttimer.com",
		IpAddress: "52.84.0.152",
	},
	&Masquerade{
		Domain:    "interviewed-integrations.com",
		IpAddress: "54.192.3.220",
	},
	&Masquerade{
		Domain:    "interviewed-staging-api.com",
		IpAddress: "54.182.5.254",
	},
	&Masquerade{
		Domain:    "intwowcher.co.uk",
		IpAddress: "205.251.206.197",
	},
	&Masquerade{
		Domain:    "io-virtualvenue.com",
		IpAddress: "54.182.6.49",
	},
	&Masquerade{
		Domain:    "iot.eu-central-1.amazonaws.com",
		IpAddress: "205.251.212.9",
	},
	&Masquerade{
		Domain:    "ipredictive.com",
		IpAddress: "13.32.10.179",
	},
	&Masquerade{
		Domain:    "itravel2000.com",
		IpAddress: "13.32.11.54",
	},
	&Masquerade{
		Domain:    "izettle.com",
		IpAddress: "216.137.52.46",
	},
	&Masquerade{
		Domain:    "jagranjosh.com",
		IpAddress: "205.251.206.121",
	},
	&Masquerade{
		Domain:    "jwplayer.com",
		IpAddress: "54.182.6.192",
	},
	&Masquerade{
		Domain:    "kariusdx.com",
		IpAddress: "54.182.5.241",
	},
	&Masquerade{
		Domain:    "kik.com",
		IpAddress: "205.251.212.97",
	},
	&Masquerade{
		Domain:    "ksmobile.com",
		IpAddress: "52.84.19.51",
	},
	&Masquerade{
		Domain:    "lifelock.com",
		IpAddress: "216.137.52.57",
	},
	&Masquerade{
		Domain:    "logica.io",
		IpAddress: "205.251.212.113",
	},
	&Masquerade{
		Domain:    "lyft.com",
		IpAddress: "54.182.6.18",
	},
	&Masquerade{
		Domain:    "lyft.com",
		IpAddress: "54.192.8.48",
	},
	&Masquerade{
		Domain:    "lyft.com",
		IpAddress: "13.32.11.108",
	},
	&Masquerade{
		Domain:    "m.payacst.com",
		IpAddress: "54.182.7.113",
	},
	&Masquerade{
		Domain:    "massrelevance.com",
		IpAddress: "205.251.206.245",
	},
	&Masquerade{
		Domain:    "media.amazonwebservices.com",
		IpAddress: "205.251.212.59",
	},
	&Masquerade{
		Domain:    "media.specialized.com",
		IpAddress: "13.32.9.232",
	},
	&Masquerade{
		Domain:    "mercadolibre.com.mx",
		IpAddress: "13.32.11.230",
	},
	&Masquerade{
		Domain:    "mercadolibre.com.uy",
		IpAddress: "205.251.207.71",
	},
	&Masquerade{
		Domain:    "mercadolivre.com.br",
		IpAddress: "205.251.212.11",
	},
	&Masquerade{
		Domain:    "milb.com",
		IpAddress: "204.246.164.117",
	},
	&Masquerade{
		Domain:    "mlbstatic.com",
		IpAddress: "13.32.14.115",
	},
	&Masquerade{
		Domain:    "mobileposse.com",
		IpAddress: "54.192.3.124",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "216.137.45.185",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "54.182.0.155",
	},
	&Masquerade{
		Domain:    "mojang.com",
		IpAddress: "204.246.164.126",
	},
	&Masquerade{
		Domain:    "moneytree.jp",
		IpAddress: "216.137.52.158",
	},
	&Masquerade{
		Domain:    "mongodb.org",
		IpAddress: "204.246.164.78",
	},
	&Masquerade{
		Domain:    "mos.asia",
		IpAddress: "54.182.0.180",
	},
	&Masquerade{
		Domain:    "mos.asia",
		IpAddress: "205.251.206.55",
	},
	&Masquerade{
		Domain:    "mparticle.com",
		IpAddress: "54.192.4.76",
	},
	&Masquerade{
		Domain:    "mparticle.com",
		IpAddress: "13.32.10.154",
	},
	&Masquerade{
		Domain:    "munchery.com",
		IpAddress: "52.84.8.215",
	},
	&Masquerade{
		Domain:    "myfonts.net",
		IpAddress: "52.84.2.57",
	},
	&Masquerade{
		Domain:    "mymagazine.smt.docomo.ne.jp",
		IpAddress: "216.137.52.42",
	},
	&Masquerade{
		Domain:    "myportfolio.com",
		IpAddress: "54.230.22.180",
	},
	&Masquerade{
		Domain:    "myportfolio.com",
		IpAddress: "52.84.17.235",
	},
	&Masquerade{
		Domain:    "net.wwe.com",
		IpAddress: "204.246.164.88",
	},
	&Masquerade{
		Domain:    "nhlstatic.com",
		IpAddress: "205.251.212.196",
	},
	&Masquerade{
		Domain:    "nosto.com",
		IpAddress: "54.239.200.232",
	},
	&Masquerade{
		Domain:    "ooyala.com",
		IpAddress: "54.182.0.8",
	},
	&Masquerade{
		Domain:    "ooyala.com",
		IpAddress: "54.192.9.217",
	},
	&Masquerade{
		Domain:    "oprah.com",
		IpAddress: "54.192.5.233",
	},
	&Masquerade{
		Domain:    "orgsync.com",
		IpAddress: "54.192.0.217",
	},
	&Masquerade{
		Domain:    "password.amazonworkspaces.com",
		IpAddress: "54.192.6.241",
	},
	&Masquerade{
		Domain:    "password.amazonworkspaces.com",
		IpAddress: "13.32.7.127",
	},
	&Masquerade{
		Domain:    "payments.zynga.com",
		IpAddress: "54.182.6.109",
	},
	&Masquerade{
		Domain:    "payscale.com",
		IpAddress: "13.32.5.212",
	},
	&Masquerade{
		Domain:    "preflight.primevideo.com",
		IpAddress: "54.192.4.228",
	},
	&Masquerade{
		Domain:    "previews.envatousercontent.com",
		IpAddress: "54.182.7.40",
	},
	&Masquerade{
		Domain:    "product-downloads.atlassian.com",
		IpAddress: "52.84.6.111",
	},
	&Masquerade{
		Domain:    "qa.o.brightcove.com",
		IpAddress: "13.32.10.178",
	},
	&Masquerade{
		Domain:    "qpyou.cn",
		IpAddress: "54.182.0.184",
	},
	&Masquerade{
		Domain:    "rafflecopter.com",
		IpAddress: "54.239.200.29",
	},
	&Masquerade{
		Domain:    "rakuten.tv",
		IpAddress: "54.192.6.58",
	},
	&Masquerade{
		Domain:    "riffsy.com",
		IpAddress: "54.182.7.54",
	},
	&Masquerade{
		Domain:    "ring.com",
		IpAddress: "54.182.7.24",
	},
	&Masquerade{
		Domain:    "rlcdn.com",
		IpAddress: "216.137.52.8",
	},
	&Masquerade{
		Domain:    "s.salecycle.com",
		IpAddress: "54.192.0.99",
	},
	&Masquerade{
		Domain:    "s3-accelerate.amazonaws.com",
		IpAddress: "54.182.6.228",
	},
	&Masquerade{
		Domain:    "seal.beyondsecurity.com",
		IpAddress: "54.182.7.186",
	},
	&Masquerade{
		Domain:    "secondlife-staging.com",
		IpAddress: "54.182.5.228",
	},
	&Masquerade{
		Domain:    "secondlife-staging.com",
		IpAddress: "216.137.52.17",
	},
	&Masquerade{
		Domain:    "secretsales.com",
		IpAddress: "54.182.0.240",
	},
	&Masquerade{
		Domain:    "seriemundial.com",
		IpAddress: "13.32.18.214",
	},
	&Masquerade{
		Domain:    "shopch.jp",
		IpAddress: "54.192.6.169",
	},
	&Masquerade{
		Domain:    "shopch.jp",
		IpAddress: "52.84.16.31",
	},
	&Masquerade{
		Domain:    "siftscience.com",
		IpAddress: "216.137.52.66",
	},
	&Masquerade{
		Domain:    "silveregg.net",
		IpAddress: "54.230.4.217",
	},
	&Masquerade{
		Domain:    "sit.abun.do",
		IpAddress: "54.230.3.195",
	},
	&Masquerade{
		Domain:    "slack-files.com",
		IpAddress: "54.182.5.245",
	},
	&Masquerade{
		Domain:    "slack.com",
		IpAddress: "54.230.10.143",
	},
	&Masquerade{
		Domain:    "slack.com",
		IpAddress: "205.251.206.86",
	},
	&Masquerade{
		Domain:    "slack.com",
		IpAddress: "54.182.0.103",
	},
	&Masquerade{
		Domain:    "slack.com",
		IpAddress: "205.251.212.192",
	},
	&Masquerade{
		Domain:    "slack.com",
		IpAddress: "54.182.5.153",
	},
	&Masquerade{
		Domain:    "smallpdf.com",
		IpAddress: "205.251.251.185",
	},
	&Masquerade{
		Domain:    "sso.m-ft.co",
		IpAddress: "54.192.10.199",
	},
	&Masquerade{
		Domain:    "stage.experiancs.com",
		IpAddress: "54.230.8.142",
	},
	&Masquerade{
		Domain:    "static-mock.production.stitchfix.com",
		IpAddress: "13.32.9.32",
	},
	&Masquerade{
		Domain:    "static-preprod.turbo.intuit.com",
		IpAddress: "13.32.5.165",
	},
	&Masquerade{
		Domain:    "static.agent-search.rdc.moveaws.com",
		IpAddress: "54.230.6.75",
	},
	&Masquerade{
		Domain:    "static.amundi.com",
		IpAddress: "54.239.200.210",
	},
	&Masquerade{
		Domain:    "static.cld.navitime.jp",
		IpAddress: "13.32.13.231",
	},
	&Masquerade{
		Domain:    "static.counsyl.com",
		IpAddress: "13.32.6.25",
	},
	&Masquerade{
		Domain:    "static.counsyl.com",
		IpAddress: "216.137.52.18",
	},
	&Masquerade{
		Domain:    "static.id.fc2cn.com",
		IpAddress: "216.137.52.114",
	},
	&Masquerade{
		Domain:    "static.lendingclub.com",
		IpAddress: "54.230.6.63",
	},
	&Masquerade{
		Domain:    "static.neteller.com",
		IpAddress: "54.182.0.38",
	},
	&Masquerade{
		Domain:    "static02.global.mifile.cn",
		IpAddress: "216.137.45.218",
	},
	&Masquerade{
		Domain:    "sundaysky.com",
		IpAddress: "205.251.212.8",
	},
	&Masquerade{
		Domain:    "supercell.com",
		IpAddress: "54.182.6.26",
	},
	&Masquerade{
		Domain:    "support.atlassian.com",
		IpAddress: "204.246.164.9",
	},
	&Masquerade{
		Domain:    "support.atlassian.com",
		IpAddress: "54.230.3.186",
	},
	&Masquerade{
		Domain:    "t-x.io",
		IpAddress: "205.251.206.239",
	},
	&Masquerade{
		Domain:    "tapad.com",
		IpAddress: "54.182.5.141",
	},
	&Masquerade{
		Domain:    "tbhtime.com",
		IpAddress: "13.32.13.217",
	},
	&Masquerade{
		Domain:    "telemetry.mozilla.org",
		IpAddress: "54.192.1.46",
	},
	&Masquerade{
		Domain:    "telemetry.mozilla.org",
		IpAddress: "54.230.4.89",
	},
	&Masquerade{
		Domain:    "telemetry.mozilla.org",
		IpAddress: "205.251.212.127",
	},
	&Masquerade{
		Domain:    "telltalegames.com",
		IpAddress: "205.251.206.158",
	},
	&Masquerade{
		Domain:    "tf-cdn.net",
		IpAddress: "13.32.18.199",
	},
	&Masquerade{
		Domain:    "ticketfly.com",
		IpAddress: "52.84.13.249",
	},
	&Masquerade{
		Domain:    "toons.tv",
		IpAddress: "204.246.164.73",
	},
	&Masquerade{
		Domain:    "topspin.net",
		IpAddress: "54.182.0.73",
	},
	&Masquerade{
		Domain:    "traversedlp.com",
		IpAddress: "205.251.212.145",
	},
	&Masquerade{
		Domain:    "tresensa.com",
		IpAddress: "205.251.212.107",
	},
	&Masquerade{
		Domain:    "tvc-mall.com",
		IpAddress: "205.251.212.90",
	},
	&Masquerade{
		Domain:    "twitchcdn-shadow.net",
		IpAddress: "54.239.142.124",
	},
	&Masquerade{
		Domain:    "typekit.net",
		IpAddress: "54.192.0.249",
	},
	&Masquerade{
		Domain:    "uat.echelonccl.com",
		IpAddress: "54.192.8.118",
	},
	&Masquerade{
		Domain:    "uat.echelonccl.com",
		IpAddress: "54.182.0.25",
	},
	&Masquerade{
		Domain:    "uprinting.com",
		IpAddress: "204.246.164.80",
	},
	&Masquerade{
		Domain:    "usanaprjb.com",
		IpAddress: "54.192.3.19",
	},
	&Masquerade{
		Domain:    "video.theblaze.com",
		IpAddress: "13.32.5.177",
	},
	&Masquerade{
		Domain:    "vivoom.co",
		IpAddress: "54.182.5.193",
	},
	&Masquerade{
		Domain:    "wap.amazon.cn",
		IpAddress: "54.182.7.228",
	},
	&Masquerade{
		Domain:    "web.nhl.com",
		IpAddress: "13.32.7.57",
	},
	&Masquerade{
		Domain:    "web.sundaysky.com",
		IpAddress: "205.251.206.172",
	},
	&Masquerade{
		Domain:    "wowcher.co.uk",
		IpAddress: "13.32.13.44",
	},
	&Masquerade{
		Domain:    "wuaki.tv",
		IpAddress: "54.192.10.179",
	},
	&Masquerade{
		Domain:    "wuaki.tv",
		IpAddress: "54.182.6.162",
	},
	&Masquerade{
		Domain:    "wuaki.tv",
		IpAddress: "13.32.9.98",
	},
	&Masquerade{
		Domain:    "www.53.localytics.com",
		IpAddress: "13.32.7.117",
	},
	&Masquerade{
		Domain:    "www.a3cloud.net",
		IpAddress: "54.230.3.169",
	},
	&Masquerade{
		Domain:    "www.a3cloud.net",
		IpAddress: "54.182.0.64",
	},
	&Masquerade{
		Domain:    "www.actnx.com",
		IpAddress: "216.137.52.51",
	},
	&Masquerade{
		Domain:    "www.alpha.encomstudios.net",
		IpAddress: "216.137.52.78",
	},
	&Masquerade{
		Domain:    "www.amazon.co.in",
		IpAddress: "216.137.45.29",
	},
	&Masquerade{
		Domain:    "www.an1.srv.c-aas-test.com",
		IpAddress: "216.137.52.251",
	},
	&Masquerade{
		Domain:    "www.api.everforth.com",
		IpAddress: "54.230.6.6",
	},
	&Masquerade{
		Domain:    "www.assetscience.com",
		IpAddress: "205.251.207.211",
	},
	&Masquerade{
		Domain:    "www.assisted-tax-preprod.a.intuit.com",
		IpAddress: "54.192.5.237",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "205.251.212.117",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "54.192.8.100",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "13.32.11.196",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "54.182.0.143",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "52.84.16.222",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "216.137.43.112",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "54.230.6.130",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "54.192.6.141",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "13.32.9.189",
	},
	&Masquerade{
		Domain:    "www.awsstatic.com",
		IpAddress: "205.251.212.221",
	},
	&Masquerade{
		Domain:    "www.bounceexchange.com",
		IpAddress: "54.239.200.93",
	},
	&Masquerade{
		Domain:    "www.brawlstarlegends.com",
		IpAddress: "52.84.7.47",
	},
	&Masquerade{
		Domain:    "www.buuteeq.com",
		IpAddress: "52.84.13.209",
	},
	&Masquerade{
		Domain:    "www.ca.support.turbotax.com",
		IpAddress: "52.84.1.81",
	},
	&Masquerade{
		Domain:    "www.canaldapeca.com.br",
		IpAddress: "54.230.6.189",
	},
	&Masquerade{
		Domain:    "www.cdn.polar.com",
		IpAddress: "216.137.52.27",
	},
	&Masquerade{
		Domain:    "www.chinamoneynetwork.com",
		IpAddress: "54.182.5.186",
	},
	&Masquerade{
		Domain:    "www.clothesforever.com",
		IpAddress: "54.192.0.231",
	},
	&Masquerade{
		Domain:    "www.coalbirds.com",
		IpAddress: "216.137.52.33",
	},
	&Masquerade{
		Domain:    "www.connectwise.com",
		IpAddress: "54.192.13.224",
	},
	&Masquerade{
		Domain:    "www.connectwise.com",
		IpAddress: "205.251.206.60",
	},
	&Masquerade{
		Domain:    "www.dalyoung.pe.kr",
		IpAddress: "205.251.206.223",
	},
	&Masquerade{
		Domain:    "www.dlgdigitalapi.com",
		IpAddress: "54.182.6.29",
	},
	&Masquerade{
		Domain:    "www.dropcam.com",
		IpAddress: "205.251.206.44",
	},
	&Masquerade{
		Domain:    "www.elevenlife.com",
		IpAddress: "205.251.251.254",
	},
	&Masquerade{
		Domain:    "www.eu.auth0.com",
		IpAddress: "54.192.3.172",
	},
	&Masquerade{
		Domain:    "www.eu.auth0.com",
		IpAddress: "205.251.212.157",
	},
	&Masquerade{
		Domain:    "www.execute-api.ap-southeast-1.amazonaws.com",
		IpAddress: "54.192.9.200",
	},
	&Masquerade{
		Domain:    "www.execute-api.us-east-1.amazonaws.com",
		IpAddress: "54.182.5.199",
	},
	&Masquerade{
		Domain:    "www.fanduel.com",
		IpAddress: "205.251.212.217",
	},
	&Masquerade{
		Domain:    "www.findawayworld.com",
		IpAddress: "54.182.6.223",
	},
	&Masquerade{
		Domain:    "www.flarecloud.net",
		IpAddress: "216.137.52.117",
	},
	&Masquerade{
		Domain:    "www.fmmotorparts.com",
		IpAddress: "52.84.4.140",
	},
	&Masquerade{
		Domain:    "www.fuhu.com",
		IpAddress: "13.32.13.10",
	},
	&Masquerade{
		Domain:    "www.fuhu.com",
		IpAddress: "52.84.13.245",
	},
	&Masquerade{
		Domain:    "www.game1.rgsrvs.com",
		IpAddress: "54.182.7.7",
	},
	&Masquerade{
		Domain:    "www.game2.rgsrvs.com",
		IpAddress: "13.32.11.150",
	},
	&Masquerade{
		Domain:    "www.game34.klabgames.net",
		IpAddress: "13.32.5.65",
	},
	&Masquerade{
		Domain:    "www.games.dev.starmp.com",
		IpAddress: "13.32.5.104",
	},
	&Masquerade{
		Domain:    "www.games.dev.starmp.com",
		IpAddress: "205.251.212.132",
	},
	&Masquerade{
		Domain:    "www.gldvideo.com",
		IpAddress: "13.32.10.6",
	},
	&Masquerade{
		Domain:    "www.hbfiles.com",
		IpAddress: "205.251.206.214",
	},
	&Masquerade{
		Domain:    "www.hbfiles.com",
		IpAddress: "52.84.14.16",
	},
	&Masquerade{
		Domain:    "www.idm-test.cfadevelop.com",
		IpAddress: "216.137.52.220",
	},
	&Masquerade{
		Domain:    "www.katapad.net",
		IpAddress: "205.251.206.177",
	},
	&Masquerade{
		Domain:    "www.katapad.net",
		IpAddress: "13.32.7.190",
	},
	&Masquerade{
		Domain:    "www.katapad.net",
		IpAddress: "54.182.5.54",
	},
	&Masquerade{
		Domain:    "www.kdc.capitalone.com",
		IpAddress: "54.182.0.77",
	},
	&Masquerade{
		Domain:    "www.lebaraplay.com",
		IpAddress: "54.182.6.170",
	},
	&Masquerade{
		Domain:    "www.manheim.man-auto1.com",
		IpAddress: "54.239.142.108",
	},
	&Masquerade{
		Domain:    "www.mercadolibre.com.sv",
		IpAddress: "54.192.15.151",
	},
	&Masquerade{
		Domain:    "www.myconnectwise.net",
		IpAddress: "13.32.5.86",
	},
	&Masquerade{
		Domain:    "www.nabicloud.com",
		IpAddress: "216.137.52.191",
	},
	&Masquerade{
		Domain:    "www.novu.com",
		IpAddress: "54.230.2.17",
	},
	&Masquerade{
		Domain:    "www.nyc837.com",
		IpAddress: "13.32.17.131",
	},
	&Masquerade{
		Domain:    "www.ozstage.com",
		IpAddress: "54.239.142.101",
	},
	&Masquerade{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "54.230.4.75",
	},
	&Masquerade{
		Domain:    "www.qa.boltdns.net",
		IpAddress: "216.137.36.222",
	},
	&Masquerade{
		Domain:    "www.qa.dlgdigitalapi.com",
		IpAddress: "54.230.4.195",
	},
	&Masquerade{
		Domain:    "www.qa.newzag.com",
		IpAddress: "13.32.17.41",
	},
	&Masquerade{
		Domain:    "www.quick-cdn.com",
		IpAddress: "52.84.1.65",
	},
	&Masquerade{
		Domain:    "www.rapid7.com",
		IpAddress: "54.182.0.72",
	},
	&Masquerade{
		Domain:    "www.rapid7.com",
		IpAddress: "54.230.16.139",
	},
	&Masquerade{
		Domain:    "www.reach150.com",
		IpAddress: "205.251.212.29",
	},
	&Masquerade{
		Domain:    "www.res.netease.com",
		IpAddress: "204.246.164.222",
	},
	&Masquerade{
		Domain:    "www.res.netease.com",
		IpAddress: "13.32.14.242",
	},
	&Masquerade{
		Domain:    "www.scruff.com",
		IpAddress: "54.192.0.55",
	},
	&Masquerade{
		Domain:    "www.shopbop.com",
		IpAddress: "204.246.164.82",
	},
	&Masquerade{
		Domain:    "www.shopch.jp",
		IpAddress: "54.182.7.170",
	},
	&Masquerade{
		Domain:    "www.shopch.jp",
		IpAddress: "54.192.4.199",
	},
	&Masquerade{
		Domain:    "www.sourabh.club",
		IpAddress: "205.251.212.76",
	},
	&Masquerade{
		Domain:    "www.srv.ygles.com",
		IpAddress: "54.182.7.127",
	},
	&Masquerade{
		Domain:    "www.stage.boltdns.net",
		IpAddress: "13.32.7.145",
	},
	&Masquerade{
		Domain:    "www.staging.kicktag.com",
		IpAddress: "205.251.212.25",
	},
	&Masquerade{
		Domain:    "www.stg.classi.jp",
		IpAddress: "205.251.212.21",
	},
	&Masquerade{
		Domain:    "www.synology.com",
		IpAddress: "13.32.5.62",
	},
	&Masquerade{
		Domain:    "www.travelhook.com",
		IpAddress: "54.182.5.200",
	},
	&Masquerade{
		Domain:    "www.tripfactory.com",
		IpAddress: "205.251.207.98",
	},
	&Masquerade{
		Domain:    "www.twitchapp.net",
		IpAddress: "54.182.7.99",
	},
	&Masquerade{
		Domain:    "www.uat.newzag.com",
		IpAddress: "54.239.130.121",
	},
	&Masquerade{
		Domain:    "www.walker.souqcdn.com",
		IpAddress: "54.192.10.67",
	},
	&Masquerade{
		Domain:    "www.webdamdb.com",
		IpAddress: "205.251.212.169",
	},
	&Masquerade{
		Domain:    "www.weledaglobalgarden.com",
		IpAddress: "54.182.5.35",
	},
	&Masquerade{
		Domain:    "www.yumpu.com",
		IpAddress: "13.32.13.61",
	},
	&Masquerade{
		Domain:    "www.zag.com",
		IpAddress: "54.239.200.238",
	},
	&Masquerade{
		Domain:    "www.zenefits.com",
		IpAddress: "52.84.20.251",
	},
	&Masquerade{
		Domain:    "www2.maclog.info",
		IpAddress: "205.251.207.11",
	},
	&Masquerade{
		Domain:    "youview.tv",
		IpAddress: "54.192.6.69",
	},
	&Masquerade{
		Domain:    "z-fe.amazon-adsystem.com",
		IpAddress: "54.182.6.193",
	},
	&Masquerade{
		Domain:    "z-na.amazon-adsystem.com",
		IpAddress: "204.246.164.57",
	},
	&Masquerade{
		Domain:    "zeasn.tv",
		IpAddress: "54.192.4.38",
	},
	&Masquerade{
		Domain:    "zenput.com",
		IpAddress: "205.251.212.121",
	},
	&Masquerade{
		Domain:    "zigbang.com",
		IpAddress: "216.137.52.248",
	},
	&Masquerade{
		Domain:    "zigbang.com",
		IpAddress: "54.192.4.71",
	},
}
