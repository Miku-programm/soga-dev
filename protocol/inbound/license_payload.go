package inbound

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	jsonv1 "encoding/json"
	"encoding/json/v2"
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"recovered.local/soga/util/rtx"
	"recovered.local/soga/config"
	"recovered.local/soga/transport/realitykdf"
)

// gS8fWl recovers 0x1db1720. The parser at 0x72f1e0 is net/url.Parse;
// the original reads its Host field (offset 40), including any explicit port.
// Invalid webapi URLs print the normalized URL and exit with status zero.
func gS8fWl() string {
	_, api, address, host, port, _, _ := config.LicensePanelFields()
	if api != "webapi" {
		return host + ":" + strconv.Itoa(port)
	}
	if !strings.HasSuffix(address, "/") {
		address += "/"
	}
	parsed, err := url.Parse(address)
	if err != nil {
		log.Println("parse url failed:", address)
		os.Exit(0)
	}
	return parsed.Host
}

// a2Kpoy1 recovers 0x1db3a60. rss is Go MemStats.Sys, not process RSS.
// Failed load queries yield an empty string; failed memory queries yield zero.
func a2Kpoy1() (loads string, total uint64, rss uint64) {
	if averages, err := load.Avg(); err == nil {
		loads = fmt.Sprintf("%.2f %.2f %.2f", averages.Load1, averages.Load5, averages.Load15)
	}
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	memory, err := mem.VirtualMemory()
	rss = stats.Sys
	if err == nil {
		total = memory.Total
	}
	return
}

// licensePayload names the map-building portion of original bj4dWcdriU
// (0x1db5f00). Its argument-free production call retains original data sources.
func licensePayload() map[string]any {
	domain := gS8fWl()
	domainMD5 := md5.Sum([]byte(domain))
	key := config.FrFQQhzTR5fY()
	loads, total, rss := a2Kpoy1()
	panel, api, address, host, port, name, version := config.LicensePanelFields()
	if !strings.HasSuffix(address, "/") {
		address += "/"
	}
	apiInfo := address + ":" + config.H7eU0uGXU()
	dbInfo := host + ":" + strconv.Itoa(port) + ":" + name
	apiMD5 := md5.Sum([]byte(apiInfo))
	dbMD5 := md5.Sum([]byte(dbInfo))
	return map[string]any{
		"key": key, "domain_origin": domain, "domain": hex.EncodeToString(domainMD5[:]),
		"api_md5": base64.StdEncoding.EncodeToString(apiMD5[:]), "db_md5": base64.StdEncoding.EncodeToString(dbMD5[:]),
		"api_info": apiInfo, "db_info": dbInfo, "server_type": config.DnhHoT2aRvs(),
		"version": version, "user_num": userCountAt50C6DA0, "loads": loads, "rss_total": total, "rss": rss,
		"panel": panel, "api": api, "timestamp": uint32(rtx.C5Jff5vs0()), "is_docker": config.VSHlrRuK7_(),
	}
}

// bj4dWcdriU recovers 0x1db5f00, with original json/v1 compatibility options.
func bj4dWcdriU() (*realitykdf.KUD0DUCag, error) {
	data, err := json.Marshal(licensePayload(), jsonv1.DefaultOptionsV1())
	if err != nil {
		return nil, fmt.Errorf("marshal data: %w", err)
	}
	return realitykdf.LBab24XA(data)
}
