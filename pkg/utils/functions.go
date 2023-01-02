package utils

func Contains(list []string, a string,) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
