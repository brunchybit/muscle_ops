package utils

import (
	"archive/tar"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
)

func buildImage(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dfReader, err := os.Open(dockerfile)
	if err != nil {
		return err
	}

	rf, err := ioutil.ReadAll(dfReader)
	if err != nil {
		return err
	}

	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(rf)),
	}

	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	_, err = tw.Write(rf)
	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	buildOpts := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	imageBuildResp, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOpts,
	)

	if err != nil {
		return err
	}

	defer imageBuildResp.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResp.Body)
	if err != nil {
		return err
	}

	return nil

}
