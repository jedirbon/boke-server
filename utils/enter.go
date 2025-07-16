package utils

import "net"

// IsPrivateIP 判断IP地址是否为内网地址
// 参数: ip - 要检查的IP地址字符串(如 "192.168.1.1")
// 返回值: bool - true表示是内网地址，false表示是公网地址
func IsPrivateIP(ip string) bool {
	// 解析IP地址
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false // 无效IP按公网处理
	}

	// IPv4私有地址范围
	if parsedIP.To4() != nil {
		// 检查IPv4私有地址范围
		return parsedIP.IsPrivate() // Go 1.17+ 内置方法

		// 如果使用Go 1.16或更早版本，可以使用以下代码替代:
		/*
			ip4 := parsedIP.To4()
			return ip4[0] == 10 || // 10.0.0.0/8
				(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
				(ip4[0] == 192 && ip4[1] == 168) || // 192.168.0.0/16
				ip4[0] == 127 || // 127.0.0.0/8 (环回地址)
				ip4[0] == 0 || // 0.0.0.0/8 (本地网络)
				(ip4[0] == 169 && ip4[1] == 254) // 169.254.0.0/16 (链路本地)
		*/
	}

	// IPv6私有地址范围
	if parsedIP.To16() != nil {
		// 检查IPv6私有地址范围
		return parsedIP.IsPrivate() // Go 1.17+ 内置方法

		// 如果使用Go 1.16或更早版本，可以使用以下代码替代:
		/*
			return parsedIP.IsLoopback() || // ::1/128
				parsedIP.IsLinkLocalUnicast() || // fe80::/10
				parsedIP.IsLinkLocalMulticast() || // ffx2::/16
				parsedIP.IsUnspecified() || // ::/128
				parsedIP.Equal(net.IPv6zero) || // ::/128
				parsedIP.IsInterfaceLocalMulticast() // ffx1::/16
		*/
	}

	return false
}

func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
