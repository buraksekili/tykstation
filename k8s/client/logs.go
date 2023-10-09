package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"net/http"
	"sync"
)

func (c *Client) Logs(ctx context.Context, name, namespace string, plo v1.PodLogOptions) (io.ReadCloser, error) {
	plo = v1.PodLogOptions{}

	logStream, err := c.ClientSet.CoreV1().Pods(namespace).GetLogs(name, &plo).Stream(ctx)
	if err != nil {
		fmt.Printf("ERR: %v", err)
		return nil, err
	}

	return logStream, nil
}

func DefaultConsumeRequest(request rest.ResponseWrapper, out io.Writer) error {
	readCloser, err := request.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer readCloser.Close()

	r := bufio.NewReader(readCloser)
	for {
		bytes, err := r.ReadBytes('\n')
		if _, err := out.Write(bytes); err != nil {
			return err
		}

		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
}

func (c Client) ParallelConsumeRequest(w http.ResponseWriter, request rest.ResponseWrapper) error {
	reader, writer := io.Pipe()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(request rest.ResponseWrapper) {
		defer wg.Done()
		if err := DefaultConsumeRequest(request, writer); err != nil {
			fmt.Fprintf(writer, "error: %v\n", err)
		}

	}(request)

	go func() {
		wg.Wait()
		writer.Close()
	}()

	_, err := io.Copy(w, reader)
	return err
}
