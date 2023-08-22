package session

var AvailableTableIDs string

func CheckPreviousTableIDs(tableIDs string) (bool) {
    if (tableIDs == AvailableTableIDs) {
        return true
    }
    SetAvailableTableIDs(tableIDs)
    return false
}

func SetAvailableTableIDs(tableIDs string) {
    AvailableTableIDs = tableIDs
}

func ResetPreviousTableIDs() {
    AvailableTableIDs = ""
}