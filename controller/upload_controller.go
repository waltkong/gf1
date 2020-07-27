package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"os"
)

const FILE_PARAM_NAME = "file"

type UploadController struct {

}

func (c *UploadController) Upload(r *ghttp.Request)  {

	uploadPath := g.Cfg().Get("server.UploadPath").(string)
	files := r.GetUploadFiles(FILE_PARAM_NAME)
	if files == nil {
		r.Response.WriteHeader(403)
		r.Exit()
	}
	names, err := files.Save(uploadPath)
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}


// UploadShow shows uploading simgle file page.
func (c *UploadController) Uploadshow(r *ghttp.Request) {

	r.Response.Write(`
    <html>
    <head>
       <title>GF Upload File Demo</title>
    </head>
       <body>
           <form enctype="multipart/form-data" action="/upload/upload" method="post">
               <input type="file" name="file" />
               <input type="submit" value="upload" />
           </form>
       </body>
    </html>
    `)
}


func (c *UploadController) Download(r *ghttp.Request)  {
	filename := r.Get("filename").(string)
	file, err:= os.Open(filename)
	if err != nil {
		r.Response.WriteHeader(403)
		r.Response.Write("find no file to download")
		r.Exit()
	}
	defer file.Close()
	buff, _ := ioutil.ReadAll(file)
	r.Response.Write(buff)
}