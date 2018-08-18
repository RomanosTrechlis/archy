package graph

const (
	EDGE = "edges"
	NODE = "nodes"
	NODE_CLASSES = "fn10273 fn6944 fn9471 fn6284 fn6956 fn6935 fn8147 fn6939 fn6936 fn6949 fn6629 fn7952 fn6680 fn6957 fn8786 fn6676 fn10713 fn7495 fn7500 fn9361 fn6279 fn6278 fn8569 fn7641 fn8568"
)

type LightElement struct {
	Data       Data     `json:"data"`
	// Common
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbed    bool     `json:"grabbed"`
	Grabbable  bool     `json:"grabbable"`
	Classes    string   `json:"classes"`
}

type Data struct {
	// Common
	Id    string  `json:"id"`

	// Node
	IdInt int     `json:"idInt"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Query bool    `json:"query"`
	Gene  bool    `json:"gene"`

	// Edge
	Source         string  `json:"source"`
	Target         string  `json:"target"`
	Weight         float64 `json:"weight"`
	Group          string  `json:"group"`
	NetworkId      int     `json:"networkId"`
	NetworkGroupId int     `json:"networkGroupId"`
	Intn           bool    `json:"intn"`
	RIntnId        int     `json:"rIntnId"`
}

func NewElement(t string) *LightElement {
	le := &LightElement{
		Data: Data{},
		Group: t,
		Removed: false,
		Selected: false,
		Selectable: true,
		Locked: false,
		Grabbed: false,
		Grabbable: true,
	}
	if t == NODE {
		le.Classes = NODE_CLASSES
	}
	return le
}

