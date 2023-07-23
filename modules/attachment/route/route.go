package route

import (
	"newsfeed/modules/attachment/controller"
	"newsfeed/modules/attachment/file_uploader"
	"newsfeed/modules/attachment/pb"
	"newsfeed/modules/attachment/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func AttachmentSetup(api *gin.RouterGroup) {
	FileInterface := file_uploader.NewFileUploaderFactory()
	conn, err := grpc.Dial(viper.GetString("ATTACHMENTURL"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	gRPCCLient := pb.NewAttachmentServiceClient(conn)
	svc := service.NewAttachmentService(gRPCCLient, FileInterface)
	ctlr := controller.NewAttachmentController(svc)

	file := api.Group("file")

	file.GET("", ctlr.GetSingleFile)
	file.DELETE("", ctlr.Delete)
	file.POST("/upload", ctlr.UploadAttachments)
	file.GET("/single/:path", ctlr.GetSingleAttachmentFile)

}
