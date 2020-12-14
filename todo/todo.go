package todo

type Item struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type List struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type ListWithItems struct {
	List
	Items []Item `json:"items"`
}
