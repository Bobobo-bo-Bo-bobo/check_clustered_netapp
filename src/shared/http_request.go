package shared

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPRequest make HTTP request
//
//   Parameters:
//     url         - URL
//     method      - HTTP method
//     header      - user defined HTTP headers
//     reader      - reader for payload to submit
//     userName    - authenticate as user
//     password    - authentication password
//     caFile      - file of CA certificate for verification of server SSL
//     insecureSSL - don't verify server certificate
//     timeout     - HTTP timeout
//   Return:
//     HTTPResult
//     error
func HTTPRequest(url string, method string, header *map[string]string, reader io.Reader, userName string, password string, caFile string, insecureSSL bool, timeout time.Duration) (HTTPResult, error) {
	var result HTTPResult
	var transp *http.Transport
	var cncl context.CancelFunc

	tlsCfg := tls.Config{}
	caPool := x509.NewCertPool()

	if insecureSSL {
		tlsCfg.InsecureSkipVerify = true

		transp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if caFile != "" {
		caData, err := ReadFileContent(caFile)
		if err != nil {
			return result, err
		}
		ok := caPool.AppendCertsFromPEM(caData)
		if !ok {
			return result, fmt.Errorf("Error: Can't append SSL data from CA file to CA pool")
		}
	}
	tlsCfg.RootCAs = caPool

	transp = &http.Transport{
		TLSClientConfig: &tlsCfg,
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: transp,
		// non-GET methods (like PATCH, POST, ...) may or may not work when encountering
		// HTTP redirect. Don't follow 301/302. The new location can be checked by looking
		// at the "Location" header.
		CheckRedirect: func(http_request *http.Request, http_via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	result.URL = url

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return result, err
	}

	cntx := context.Background()
	cntx, cncl = context.WithTimeout(cntx, timeout)
	defer cncl()
	request.WithContext(cntx)

	defer func() {
		if request.Body != nil {
			ioutil.ReadAll(request.Body)
			request.Body.Close()
		}
	}()

	request.SetBasicAuth(userName, password)

	// add required headers
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	// set User-Agent
	request.Header.Set("User-Agent", UserAgent)

	// close connection after response and prevent re-use of TCP connection because some implementations (e.g. HP iLO4)
	// don't like connection reuse and respond with EoF for the next connections
	request.Close = true

	// add supplied additional headers
	if header != nil {
		for key, value := range *header {
			request.Header.Add(key, value)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	defer func() {
		ioutil.ReadAll(response.Body)
		response.Body.Close()
	}()

	result.Status = response.Status
	result.StatusCode = response.StatusCode
	result.Header = response.Header
	result.Content, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}
