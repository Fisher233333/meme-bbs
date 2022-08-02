package common

import (
	"bbs-go/model/constants"
	"bbs-go/pkg/config"
	"bbs-go/pkg/html"
	"bbs-go/pkg/markdown"
	"bbs-go/pkg/text"
	"net"
	"net/http"
	"strings"
)

// IsProd 是否是正式环境
func IsProd() bool {
	return config.Instance.Env == "prod"
}

func GetSummary(contentType string, content string) (summary string) {
	if contentType == constants.ContentTypeMarkdown {
		summary = markdown.GetSummary(content, constants.SummaryLen)
	} else if contentType == constants.ContentTypeHtml {
		summary = html.GetSummary(content, constants.SummaryLen)
	} else {
		summary = text.GetSummary(content, constants.SummaryLen)
	}
	return
}

// GetMarkdownSummary 截取markdown摘要
func GetMarkdownSummary(markdownStr string) string {
	return markdown.GetSummary(markdownStr, constants.SummaryLen)
}

// GetRequestIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func GetRequestIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func GetUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}
