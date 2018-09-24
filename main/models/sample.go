package models


type Sample struct {
	EncodedImage string `json:"encodedImage"`
	Label        int    `json:"label"`
}


// Factory of samples
func GenerateSampleList() []Sample {
	var samples []Sample
	for i := 0; i < 10; i++ {
		sample := Sample{Label: i, EncodedImage: "asdf"}
		samples = append(samples, sample)
	}
	return samples
}