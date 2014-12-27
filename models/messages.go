package models

type Message struct {
    Message string `json:"msg" required`
}

func GetDefault500() Message {
    return Message {"System Error Occurred"}
}

func Get404(model string, id string) Message {
    return Message {"Unable to find " + model + " by id: " + id}
}

func GetDelMsg(model string, id string) Message {
    return Message {"Deleted " + model + " by id: " + id}
}
