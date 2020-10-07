// Written by Gon Yi
package areq

var mime = map[string]string{
	"3g2":    "video/3gpp2",
	"3gp":    "video/3gpp",
	"7z":     "application/x-7z-compressed",
	"aac":    "audio/aac",
	"abw":    "application/x-abiword",
	"arc":    "application/x-freearc",
	"avi":    "video/x-msvideo",
	"azw":    "application/vnd.amazon.ebook",
	"bin":    "application/octet-stream",
	"bmp":    "image/bmp",
	"bz":     "application/x-bzip",
	"bz2":    "application/x-bzip2",
	"csh":    "application/x-csh",
	"css":    "text/css",
	"csv":    "text/csv",
	"doc":    "application/msword",
	"docx":   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"eot":    "application/vnd.ms-fontobject",
	"epub":   " application/epub+zip",
	"gif":    "image/gif",
	"gz":     "application/gzip",
	"htm":    "text/html",
	"html":   "text/html",
	"ico":    "image/vnd.microsoft.icon",
	"ics":    "text/calendar",
	"jar":    "application/java-archive",
	"jpeg":   "image/jpeg",
	"jpg":    "image/jpeg",
	"js":     "text/javascript",
	"json":   "application/json",
	"jsonld": "application/ld+json",
	"mid":    "audio/midi",
	"midi":   "audio/midi",
	"mjs":    "text/javascript",
	"mp3":    "audio/mpeg",
	"mpeg":   "video/mpeg",
	"mpkg":   "application/vnd.apple.installer+xml",
	"odp":    "application/vnd.oasis.opendocument.presentation",
	"ods":    "application/vnd.oasis.opendocument.spreadsheet",
	"odt":    "application/vnd.oasis.opendocument.text",
	"oga":    "audio/ogg",
	"ogv":    "video/ogg",
	"ogx":    "application/ogg",
	"opus":   "audio/opus",
	"otf":    "font/otf",
	"pdf":    "application/pdf",
	"php":    "application/php",
	"png":    "image/png",
	"ppt":    "application/vnd.ms-powerpoint",
	"pptx":   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"rar":    "application/x-rar-compressed",
	"rtf":    "application/rtf",
	"sh":     "application/x-sh",
	"svg":    "image/svg+xml",
	"swf":    "application/x-shockwave-flash",
	"tar":    "application/x-tar",
	"tif":    "image/tiff",
	"tiff":   "image/tiff",
	"ts":     "video/mp2t",
	"ttf":    "font/ttf",
	"txt":    "text/plain",
	"vsd":    "application/vnd.visio",
	"wav":    "audio/wav",
	"weba":   "audio/webm",
	"webm":   "video/webm",
	"webp":   "image/webp",
	"woff":   "font/woff",
	"woff2":  "font/woff2",
	"xhtml":  "application/xhtml+xml",
	"xls":    "application/vnd.ms-excel",
	"xlsx":   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xml":    "text/xml",
	"xul":    "application/vnd.mozilla.xul+xml",
	"zip":    "archive application/zip",
}

func MimeByFileExt(fileExt, fallback string) string {
	// Search for the mime using file extension
	if v, ok := mime[fileExt]; ok {
		return v
	}
	// If fallback is an extension, use its mime
	if v, ok := mime[fallback]; ok {
		return v
	}
	// If not, return fileExt as is.
	return fileExt
}
