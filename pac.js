function FindProxyForURL(url, host) {
    if (isPlainHostName(host)) return "DIRECT";
    else if (shExpMatch(host, "music.163.com")) {
        return "PROXY pac.onenorthtech.com:7999"
    } else if (shExpMatch(host, "*.v.163.com")) {
        return "PROXY pac.onenorthtech.com:7999"
    } else if (shExpMatch(host, "*.music.163.com")) {
        return "PROXY pac.onenorthtech.com:7999"
    } else if (shExpMatch(host, "*.music.126.net")) {
        return "PROXY pac.onenorthtech.com:7999"
    } else return "DIRECT"
}