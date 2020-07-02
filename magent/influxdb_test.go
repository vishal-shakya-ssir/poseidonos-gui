package main

import (
	"context"
	"fmt"
	client1 "github.com/influxdata/influxdb1-client"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"testing"
	"time"
)

//Should not raise an error if the data to be written to DB is valid
//This function calls the checkPointTypes function
//It verifies that there is no error produced if the data passed to the function is valid
func TestCheckPointTypes(t *testing.T) {
	pt := client1.Point{
		Measurement: "air",
		Tags:        map[string]string{},
		Fields: map[string]interface{}{
			"cpu": 2,
		},
		Time: time.Now()}
	err := checkPointTypes(pt)
	require.NoError(t, err)

}

//Should return error when data to be written to DB is invalid
//This function calls the checkPointTypes function
//It verifies that an error is produced if the data passed to the function is not valid
func TestCheckPointTypesError(t *testing.T) {
	p := client1.Point{}
	pt := client1.Point{
		Measurement: "air",
		Tags:        map[string]string{},
		Fields: map[string]interface{}{
			"cpu": p,
		},
		Time: time.Now()}
	err := checkPointTypes(pt)
	require.Error(t, err)
}

//Should write to DB using HTTP protocol
//This function calls the NewHTTPClient and then calls the InfluxDB.Write
//It verifies that there is no error produced if the datais send to DB using HTTP protocol
func TestWriteToDBHTTP(t *testing.T) {
	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	})
	ts := httptest.NewServer(httpHandler)
	defer ts.Close()
	u, err := url.Parse(fmt.Sprintf("http://%s", ts.Listener.Addr().String()))
	require.NoError(t, err)
	newPoint := client1.Point{
		Fields: map[string]interface{}{
			"cpu": 10,
		},
		Tags:        map[string]string{},
		Measurement: "cpu_magent",
		Time:        time.Now(),
	}
	bp := client1.BatchPoints{
		Precision:        "ms",
		WriteConsistency: "one",
		RetentionPolicy:  "autogen",
	}
	bp.Points = append(bp.Points, newPoint)
	ctx := context.Background()
	client, err := NewHTTPClient(config)
	require.NoError(t, err)
	influxdb := Init("http")
	config.URL = u

	//Should send an http request using http protocol with the points in the channel
	clientPoint := ClientPoint{
		Fields: map[string]interface{}{
			"cpu": 10,
		},
		Tags:            map[string]string{},
		Measurement:     "cpu_magent",
		RetentionPolicy: "autogen",
	}
	ctx, cancel := context.WithCancel(context.Background())
	dataChan := make(chan ClientPoint, 100)
	dataChan <- clientPoint
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	writeToInfluxDB(ctx, dataChan, client, config, 1, influxdb)

	//Should send an http request when the time interval for flushing is exceeded
	ctx, cancel = context.WithCancel(context.Background())
	dataChan <- clientPoint
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	writeToInfluxDB(ctx, dataChan, client, config, 10, influxdb)

	//Should raise an error when the request returns error
	config.URL, _ = url.Parse(fmt.Sprintf("http://%s/test", ts.Listener.Addr().String()))
	err = influxdb.Write(ctx, client, bp, config)
	require.Error(t, err)

	//Should not send request using  http protocol if data is invalid
	point := client1.Point{
		Fields: map[string]interface{}{
			"cpu": []int{10},
		},
		Tags:        map[string]string{},
		Measurement: "cpu_magent",
		Time:        time.Now(),
	}
	bp.Points = append(bp.Points, point)
	err = influxdb.Write(ctx, client, bp, config)
	require.Error(t, err)
}

//Should write to DB using Unix Protocol
//This function calls the NewHTTPClient and then calls the InfluxDB.Write
//It verifies that there is no error produced if the data is sent to DB using Unix socket
func TestWriteToDBUnix(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "telegraf-test")
	if err != nil {
		require.NoError(t, err)
	}
	defer os.RemoveAll(tmpdir)

	sock := path.Join(tmpdir, "test.sock")
	listener, err := net.Listen("unix", sock)
	require.NoError(t, err)

	ts := httptest.NewUnstartedServer(http.NotFoundHandler())
	ts.Listener = listener
	ts.Start()
	defer ts.Close()
	newPoint := client1.Point{
		Fields: map[string]interface{}{
			"cpu": 10,
		},
		Tags:        map[string]string{},
		Measurement: "cpu_magent",
		Time:        time.Now(),
	}
	bp := client1.BatchPoints{
		Precision:        "ms",
		WriteConsistency: "one",
		RetentionPolicy:  "autogen",
	}
	bp.Points = append(bp.Points, newPoint)
	ts.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/write":
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	})
	ctx := context.Background()
	influxdb := Init("unix")
	config.URL = &url.URL{Scheme: "unix", Path: sock}
	config.URLString = config.URL.String()
	client, err := NewHTTPClient(config)
	require.NoError(t, err)

	//Should send an http request using unix socket protocol with the points in the channel
	err = influxdb.Write(ctx, client, bp, config)
	require.NoError(t, err)

	//Should raise an error when the response status is not OK
	config.URL.Path = "/test"
	err = influxdb.Write(ctx, client, bp, config)
	require.Error(t, err)

	//Should not send request using  unix socket protocol if data is invalid
	point := client1.Point{
		Fields: map[string]interface{}{
			"cpu": []int{10},
		},
		Tags:        map[string]string{},
		Measurement: "cpu_magent",
		Time:        time.Now(),
	}
	bp.Points = append(bp.Points, point)
	err = influxdb.Write(ctx, client, bp, config)
	require.Error(t, err)
}
