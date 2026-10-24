package hysteriabrutal

import (
	"github.com/apernet/quic-go/congestion"
	common "recovered.local/soga/transport/hysteriaerr"
	"reflect"
	"testing"
	"time"
)

// 原样本 RTTI 硬编码证据，字段名、偏移、类型和 tag 均独立校验。
func TestOriginalRTTILayouts(t *testing.T) {
	type field struct {
		name     string
		offset   uintptr
		typ      reflect.Type
		tag      reflect.StructTag
		embedded bool
	}
	cases := []struct {
		typ    reflect.Type
		size   uintptr
		fields []field
	}{
		{reflect.TypeOf(iC6_vHy61ee{}), 40, []field{
			{"o00cUt", 0, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"rcyLtXR67hl", 1, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"fyuyH6zvX", 8, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"gvnxB1rdE", 16, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"pWaFRsxfpvC", 24, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"tcFhoDpK1YNk", 32, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(wzBRgOvr8BC{}), 128, []field{
			{"p0vSfsOlU", 0, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"jNhPqIq", 24, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"pLa7WgMaenyW", 32, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"v_nXPcs0HTe", 40, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"yXecWTdyLG", 64, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"ujKE8LV1OP", 88, reflect.TypeOf((*(iC6_vHy61ee))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(tWtTcf[wzBRgOvr8BC]{}), 136, []field{
			{"cPIy2wvomr", 0, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"ySYKCoJI", 8, reflect.TypeOf((*(wzBRgOvr8BC))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(YXV6CyZ5Ux[tWtTcf[wzBRgOvr8BC]]{}), 48, []field{
			{"ckRjtgECJ", 0, reflect.TypeOf((*([]tWtTcf[wzBRgOvr8BC]))(nil)).Elem(), "", false},
			{"rxa2MP2pFk", 24, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"gwZziNLjai8", 32, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"aY2XfPJecIWV", 40, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(h8SYlleQ[wzBRgOvr8BC]{}), 64, []field{
			{"fSrleCn8X1", 0, reflect.TypeOf((*(YXV6CyZ5Ux[tWtTcf[wzBRgOvr8BC]]))(nil)).Elem(), "", false},
			{"bDBgyW", 48, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"f8sQndcVreq", 56, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(kO8aHEaU9{}), 32, []field{
			{"m7dTK3NJaO", 0, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"joNWs9", 24, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(h4dvYPRUDTrt{}), 64, []field{
			{"yCiDnZ", 0, reflect.TypeOf((*([2]kO8aHEaU9))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(YXV6CyZ5Ux[kO8aHEaU9]{}), 48, []field{
			{"ckRjtgECJ", 0, reflect.TypeOf((*([]kO8aHEaU9))(nil)).Elem(), "", false},
			{"rxa2MP2pFk", 24, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"gwZziNLjai8", 32, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"aY2XfPJecIWV", 40, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(k8ofg5VH{}), 32, []field{
			{"sCcx3Ca6E", 0, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"yhpZvxZ", 8, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"tSOB2v", 16, reflect.TypeOf((*(time.Duration))(nil)).Elem(), "", false},
			{"qr1Cd4llNbRm", 24, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(f2ypPZu[k8ofg5VH, awuEkJg9p]{}), 40, []field{
			{"nvK8flg3g", 0, reflect.TypeOf((*(k8ofg5VH))(nil)).Elem(), "", false},
			{"jgaVJt", 32, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(Nqk7YFvzgcZg[k8ofg5VH, awuEkJg9p]{}), 40, []field{
			{"aaF6ngDQX", 0, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
			{"fPLAVE6X", 8, reflect.TypeOf((*([]f2ypPZu[k8ofg5VH, awuEkJg9p]))(nil)).Elem(), "", false},
			{"enH1PpcAyo", 32, reflect.TypeOf((*(func(k8ofg5VH, k8ofg5VH) int))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(a4a_ewpUFjg{}), 72, []field{
			{"jm_R6V", 0, reflect.TypeOf((*(*Nqk7YFvzgcZg[k8ofg5VH, awuEkJg9p]))(nil)).Elem(), "", false},
			{"r10Waeh", 8, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"cM5vf8", 32, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"pLwLGJtFf", 40, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"uxti2K1", 48, reflect.TypeOf((*(uint64))(nil)).Elem(), "", false},
			{"v9KunT", 56, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"tpXKRSNXF", 64, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"sI7jVWaahw", 65, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(oahYXcYd4{}), 272, []field{
			{"eoyKL70a", 0, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"cyAQ1b", 8, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"oZZJNF", 16, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"o8yuvw", 24, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"g4GkKaB7", 32, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"fhJyzuFBO", 40, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"ciarX4f", 64, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"n7S7Dan", 88, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"fhILy8i", 96, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"q7YCyTJjNN", 104, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"zISoP8ufhtI", 112, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"e7KrTFhqp", 120, reflect.TypeOf((*(*h8SYlleQ[wzBRgOvr8BC]))(nil)).Elem(), "", false},
			{"y_Z4NrytDBe4", 128, reflect.TypeOf((*(h4dvYPRUDTrt))(nil)).Elem(), "", false},
			{"vFlrYPDp4", 192, reflect.TypeOf((*(YXV6CyZ5Ux[kO8aHEaU9]))(nil)).Elem(), "", false},
			{"rDjgJc", 240, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"qgtQO9f7j", 248, reflect.TypeOf((*(*a4a_ewpUFjg))(nil)).Elem(), "", false},
			{"lWRr8Cm0w", 256, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"i4NwjwX59Gl", 264, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"exqrreSd1pd", 265, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]{}), 16, []field{
			{"nvK8flg3g", 0, reflect.TypeOf((*(JaadYOR_Vaf6))(nil)).Elem(), "", false},
			{"jgaVJt", 8, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]{}), 40, []field{
			{"aaF6ngDQX", 0, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
			{"fPLAVE6X", 8, reflect.TypeOf((*([]f2ypPZu[JaadYOR_Vaf6, awuEkJg9p]))(nil)).Elem(), "", false},
			{"enH1PpcAyo", 32, reflect.TypeOf((*(func(JaadYOR_Vaf6, JaadYOR_Vaf6) int))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(tGYu23kout{}), 416, []field{
			{"jzNuIRaj5L4f", 0, reflect.TypeOf((*(congestion.RTTStatsProvider))(nil)).Elem(), "", false},
			{"ajdznEwg3HP", 16, reflect.TypeOf((*(KYybyfkN))(nil)).Elem(), "", false},
			{"gMj1ag2rOX9T", 32, reflect.TypeOf((*(*common.OQpLDX))(nil)).Elem(), "", false},
			{"jOTShB", 40, reflect.TypeOf((*(lajiJA))(nil)).Elem(), "", false},
			{"eiff4Dk6J", 48, reflect.TypeOf((*(*oahYXcYd4))(nil)).Elem(), "", false},
			{"g4B8Dg", 56, reflect.TypeOf((*(awuEkJg9p))(nil)).Elem(), "", false},
			{"wUnF0r", 64, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"bkPDow", 72, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"pC_xa8", 80, reflect.TypeOf((*(uint64))(nil)).Elem(), "", false},
			{"sSPtxam", 88, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"clm34LfHOfH", 96, reflect.TypeOf((*(*Nqk7YFvzgcZg[JaadYOR_Vaf6, awuEkJg9p]))(nil)).Elem(), "", false},
			{"es2Rg_vDy", 104, reflect.TypeOf((*(time.Duration))(nil)).Elem(), "", false},
			{"p4pnrwiNi", 112, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"ndBPgZ", 136, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"kPUkZLy", 144, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"qkdjbFU0l", 152, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"irwaYXRFoja", 160, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"ytPxW1_vL", 168, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"fI9EUh", 176, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"s_Zx6iZ", 184, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"rbHB4J1ksIV", 192, reflect.TypeOf((*(JaadYOR_Vaf6))(nil)).Elem(), "", false},
			{"rDtgaAw", 200, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"qAZ13bY", 208, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"sZxgbLDX", 216, reflect.TypeOf((*(float64))(nil)).Elem(), "", false},
			{"xIZ4Qm6J81", 224, reflect.TypeOf((*(int64))(nil)).Elem(), "", false},
			{"sewKJeM4", 232, reflect.TypeOf((*(int))(nil)).Elem(), "", false},
			{"_62XJdnnNP", 240, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"dv_p1GSp", 264, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"ccZF6DrRpP", 272, reflect.TypeOf((*(int64))(nil)).Elem(), "", false},
			{"afwE420", 280, reflect.TypeOf((*(JaadYOR_Vaf6))(nil)).Elem(), "", false},
			{"bmSW0VMoFYsi", 288, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"hq2wUhQa", 296, reflect.TypeOf((*(time.Time))(nil)).Elem(), "", false},
			{"zFN_SDSij", 320, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"gZ285RGO1K", 321, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"hcFC7avZ", 322, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"daiTa6d", 328, reflect.TypeOf((*(hKvOmYf_pbF))(nil)).Elem(), "", false},
			{"zdCu53B_Jn", 336, reflect.TypeOf((*(congestion.PacketNumber))(nil)).Elem(), "", false},
			{"voFtPR2C", 344, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"xTTndm", 352, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"dxYh9Z", 353, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"i5pDu7Gmkf", 354, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"zQhgaatIp", 355, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"f4pwsMIQjjC", 356, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"btA4YPFSwnmT", 357, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"yL5qU8a", 358, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"wnghShWZg", 360, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"gX1RE3uW", 368, reflect.TypeOf((*(uint8))(nil)).Elem(), "", false},
			{"zaXoVmSmYwgs", 376, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"dzxASp", 384, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"hIPu5Tz9", 392, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"ehQ4ZmGJ4u", 400, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"d2Q76d8Uzv8K", 408, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
		}},
		{reflect.TypeOf(ri4dMyMP{}), 80, []field{
			{"w87e3ZbfBKJR", 0, reflect.TypeOf((*(JaadYOR_Vaf6))(nil)).Elem(), "", false},
			{"zkhqa7g1Vp3", 8, reflect.TypeOf((*(bool))(nil)).Elem(), "", false},
			{"e1LhJmPmT", 16, reflect.TypeOf((*(time.Duration))(nil)).Elem(), "", false},
			{"jSNh_B5Te", 24, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
			{"xDGjNq0R", 32, reflect.TypeOf((*(iC6_vHy61ee))(nil)).Elem(), "", false},
			{"kPsWl9", 72, reflect.TypeOf((*(congestion.ByteCount))(nil)).Elem(), "", false},
		}},
	}
	for _, c := range cases {
		t.Run(c.typ.String(), func(t *testing.T) {
			if c.typ.Size() != c.size || c.typ.NumField() != len(c.fields) {
				t.Fatalf("layout size=%d fields=%d", c.typ.Size(), c.typ.NumField())
			}
			for i, e := range c.fields {
				f := c.typ.Field(i)
				if f.Name != e.name || f.Offset != e.offset || f.Type != e.typ || f.Tag != e.tag || f.Anonymous != e.embedded {
					t.Fatalf("field %d: %+v expected %+v", i, f, e)
				}
			}
		})
	}
}
