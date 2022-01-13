package lib

import "github.com/onflow/cadence"

// define metadata struct
type Metadata struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	ExternalURL  string `json:"external_url"`
	Image        string `json:"image"`
	Attributes   string `json:"attributes"`
	AnimationURL string `json:"animation_url"`
}

// v1 format extend metadata struct
type MetadataV1 struct {
	Metadata
	Collection string `json:"collection"`
}

// create a new metadataV1 struct
func NewMetadataV1(name string, version string, description string, externalURL string, image string, attributes string, animationURL string, collection string) *MetadataV1 {
	return &MetadataV1{
		Metadata: Metadata{
			Name:         name,
			Version:      version,
			Description:  description,
			ExternalURL:  externalURL,
			Image:        image,
			Attributes:   attributes,
			AnimationURL: animationURL,
		},
		Collection: collection,
	}
}

func CadenceString(s string) cadence.String {
	ret, err := cadence.NewString(s)
	if err != nil {
		panic(err)
	}
	return ret
}

// ToCadenceDictionary
func (m *MetadataV1) ToCadenceDictionary() cadence.Dictionary {
	data := []cadence.KeyValuePair{
		{Key: CadenceString("name"), Value: CadenceString(m.Name)},
		{Key: CadenceString("version"), Value: CadenceString(m.Version)},
		{Key: CadenceString("description"), Value: CadenceString(m.Description)},
		{Key: CadenceString("external_url"), Value: CadenceString(m.ExternalURL)},
		{Key: CadenceString("image"), Value: CadenceString(m.Image)},
		{Key: CadenceString("attributes"), Value: CadenceString(m.Attributes)},
		{Key: CadenceString("animation_url"), Value: CadenceString(m.AnimationURL)},
		{Key: CadenceString("collection"), Value: CadenceString(m.Collection)},
	}
	return cadence.NewDictionary(data)
}
