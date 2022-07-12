package gce

import (
  "fmt"
  "net/http"
  "io"
  "bytes"
)

type Web struct{
    Writer http.ResponseWriter
    Request *http.Request
}

func Init_Web(res http.ResponseWriter, req *http.Request)Web{
    return Web{Writer:res,Request:req}
}
func Start_Web(port string){
    http.ListenAndServe(":"+port, nil)
}
func Start_Web_TLS(port string,certpem string,keypem string){
    http.ListenAndServeTLS(":"+port,certpem,keypem, nil)
}
func (Conf *Web)GET(id string)string{
    query := Conf.Request.URL.Query()
    ids := query.Get(id)
    return ids
}
func (Conf *Web)Echo(text string){
    fmt.Fprintf(Conf.Writer,text)
}
func (Conf *Web)WebPath()string{
    return Conf.Request.URL.Path
}
func (Conf *Web)FILE(id string)map[string]interface{}{
    var Buf bytes.Buffer
    formFile, header, _ := Conf.Request.FormFile(id)
    io.Copy(&Buf, formFile)
    contents := Buf.String()
    formFile.Close()
    resmap:=make(map[string]interface{})
    resmap["Filename"]=header.Filename
    resmap["Size"]=header.Size
    resmap["Type"]=header.Header["Content-Type"]
    resmap["FileFrom"]=contents
    return resmap
}
func (Conf *Web)POST(id string)string{
    return Conf.Request.PostFormValue(id)
}
