function FindProxyForURL(url, host) {
    if (isPlainHostName(host)) return "DIRECT";
    else if (shExpMatch(host, "music.163.com")) {
        return "PROXY proxy.uku.im:443"
    } else if (shExpMatch(host, "*.v.163.com")) {
        return "PROXY proxy.uku.im:443"
    } else if (shExpMatch(host, "*.music.163.com")) {
        return "PROXY proxy.uku.im:443"
    } else if (shExpMatch(host, "*.music.126.net")) {
        return "PROXY proxy.uku.im:443"
    } else return "DIRECT"
}