package models

type BackendError struct {
    Code int `json:"errcode"`
    Msg  string
}

func (b BackendError) Error() string {
    return b.Msg
}