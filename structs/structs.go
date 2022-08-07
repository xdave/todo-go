package structs

type TodoInsert struct {
	Text string `json:"text"`
	Checked bool `json:"checked"`
	Ip_Address string `json:"ip_address"`
}
type TodoAll struct {
	TodoInsert
	Id int64 `json:"id"`
}
type TodoQuery struct {
	Id int64 `json:"id"`
	Text string `json:"text"`
	Checked bool `json:"checked"`
}
type TodoCheck struct {
	Id int64 `json:"id"`
	Checked bool `json:"checked"`
}