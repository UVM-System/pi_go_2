package photo

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type ImageBytes struct {
	FieldName string
	FileName string
	Content []byte
	ContentType string
}
type PostBytes struct {
	FileMap []ImageBytes
	FieldMap map[string]string
}

func createFormDataFromBytes(formdata PostBytes) (string,*bytes.Buffer,error)  {
	bodyBuf := &bytes.Buffer{}
	bodyWriter :=multipart.NewWriter(bodyBuf)
	//写入文件
	for _,imageBytes :=range formdata.FileMap{
		bufferWriter,_ :=bodyWriter.CreatePart(mimeHeader(imageBytes.FieldName,imageBytes.FileName,imageBytes.ContentType))
		bufferWriter.Write(imageBytes.Content)
	}
	for key, val := range formdata.FieldMap {
		_ = bodyWriter.WriteField(key, val)
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	return contentType,bodyBuf,nil
}

//封装文件内容

func mimeHeader(fieldname, filename,contenttype string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", contenttype)
	return h
}
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
