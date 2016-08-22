package secret

type Secret struct {
	ID         string `json:"Id,omitempty"`
	Name       string
	Data       []byte `json:",omitempty"`
	Mountpoint string `json:",omitempty"`
	Required   bool   `json:",omitempty"`
}
