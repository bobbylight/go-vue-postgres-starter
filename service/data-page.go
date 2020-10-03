package main

type DataPage struct {
    Offset uint32 `json:"offset"`
    Data []Task `json:"data"`
    Total uint32 `json:"count"`
}
