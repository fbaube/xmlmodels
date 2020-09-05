package xmlmodels

import (
	S "strings"
)

// AnalysisRecord is the results of content analysis. It is named
// "Record" because it is meant to be persisted to the database.
// It is embedded in db.ContentRecord
type AnalysisRecord struct {
	Contyping
	// RootTag        string // e.g. "html", enclosing both <head> and <body>
	// RootAtts       string // e.g. <html lang="en"> yields << lang="en" >>
	MarkdownFlavor string
	KeyElms
	ContentitySections
	XmlInfo
	DitaInfo
}

// IsXML is true for all XML, including all HTML.
func (p AnalysisRecord) IsXML() bool {
	s := p.FileType()
	return s == "XML" || s == "HTML"
}

func (p *AnalysisRecord) String() string {
	return ("AR!")
}

// FileType returns "XML", "MKDN", "HTML", or future stuff TBD.
func (p AnalysisRecord) FileType() string {
	// Exceptional case
	if S.HasPrefix(p.MType, "xml/html/") {
		return "HTML"
	}
	if S.HasPrefix(p.MimeType, "text/html") {
		return "HTML"
	}
	// Normal case
	// return S.ToUpper(MTypeSub(p.MType, 0))
	// Cut & Paste
	i := S.Index(p.MType, "/")
	return S.ToUpper(p.MType[:i])
}
