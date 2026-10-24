package CbO_7wp3B

import (
	A "recovered.local/soga/a3PdV_bmpRUK"
	C "recovered.local/soga/cNT24rG4F"
	M "recovered.local/soga/muc279Mpa6S4"
	Q "recovered.local/soga/q5bqCgTdxaTM"
)

// gJwSzy 对应 0x1f9a140，依次尝试支持 GetNode 的面板提供的证书。
func gJwSzy(panels []Q.VyFk5uT) bool {
	for _, panel := range panels {
		nodePanel, ok := panel.(Q.Chd8czf6)
		if !ok {
			continue
		}
		nodeID := panel.GetAPIConfig().Yf1KqSqJpNI7
		node, err := nodePanel.GetNode()
		if err != nil {
			cert, key := A.J6c4NB_R(nodeID)
			if A.RUKs8pSwE4(cert, key) == nil {
				M.V8GN5Stk(4, "get node %v failed (%v), use cached panel cert: %v", nodeID, err, cert)
				C.SetCertificateFiles(cert, key)
				return true
			}
			M.V8GN5Stk(4, "get node %v failed, skip panel cert (if this node uses the panel content cert mode, restart soga after the panel is reachable): %v", nodeID, err)
			continue
		}
		if node.I3k4JnNp7nB == "" || node.XATiVun == "" {
			continue
		}
		cert, key, err := A.Qr4n7sWM5di(node.EMlXEc, node.I3k4JnNp7nB, node.XATiVun)
		if err != nil {
			M.TRTbrquD(4, "write panel cert failed:", err)
		} else {
			C.SetCertificateFiles(cert, key)
		}
		// 原逻辑在保存失败后仍报告已使用面板证书并返回 true。
		M.V8GN5Stk(0, "use panel cert for node %v", node.EMlXEc)
		return true
	}
	return false
}
