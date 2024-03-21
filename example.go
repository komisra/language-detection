package main
package client

func main() {

	var texts = []string{"labas rytas", "good morning"}
	results, err = client.DetectBatch(texts)

	if err != nil {
    	fmt.Fprintln(os.Stderr, "error detecting language:", err)
    	os.Exit(1)
    	return
	}

	fmt.Fprintln(os.Stdout, "First text language:", detections[0][0].Language)
	fmt.Fprintln(os.Stdout, "Second text language:", detections[1][0].Language)
}