package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"secondLife/imageRecog"
	"strings"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

var (
	graphModel   *tf.Graph
	sessionModel *tf.Session
	labels       []string
)

type ImageURLRequest struct {
	URL string `json:"url"`
}

func LoadModel() error {
	// Load inception model
	model, err := ioutil.ReadFile("/model/tensorflow_inception_graph.pb")
	if err != nil {
		return err
	}
	graphModel = tf.NewGraph()
	if err := graphModel.Import(model, ""); err != nil {
		return err
	}

	sessionModel, err = tf.NewSession(graphModel, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Load labels
	labelsFile, err := os.Open("/model/imagenet_comp_graph_label_strings.txt")
	if err != nil {
		return err
	}
	defer labelsFile.Close()
	scanner := bufio.NewScanner(labelsFile)
	// Labels are separated by newlines
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
func (h *APIServer) imageRecognize(w http.ResponseWriter, r *http.Request) error {
	var imageURLReq ImageURLRequest

	err := json.NewDecoder(r.Body).Decode(&imageURLReq)
	if err != nil {
		responseError(w, "Invalid JSON request body", http.StatusBadRequest)
		return err
	}

	resp, err := http.Get(imageURLReq.URL)
	if err != nil {
		responseError(w, "Failed to download image from the provided URL", http.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()

	// Copy image data to a buffer
	var imageBuffer bytes.Buffer
	_, err = io.Copy(&imageBuffer, resp.Body)
	if err != nil {
		responseError(w, "Failed to read image data", http.StatusInternalServerError)
		return err
	}

	// Extract filename and extension from URL
	imageName := strings.Split(imageURLReq.URL, "/")
	filename := imageName[len(imageName)-1]

	// Make tensor
	tensor, err := imageRecog.MakeTensorFromImage(&imageBuffer, filename)
	if err != nil {
		responseError(w, "Invalid image", http.StatusBadRequest)
		return err
	}

	// Run inference
	output, err := sessionModel.Run(
		map[tf.Output]*tf.Tensor{
			graphModel.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			graphModel.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		responseError(w, "Could not run inference", http.StatusInternalServerError)
		return err
	}

	// Return best labels
	responseJSON(w, ClassifyResult{
		Filename: filename,
		Labels:   findBestLabels(output[0].Value().([][]float32)[0]),
	})
	return nil
}
