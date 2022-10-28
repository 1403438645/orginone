package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdministrativeDivisionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdministrativeDivisionLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAdministrativeDivisionLogic {
	return GetAdministrativeDivisionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdministrativeDivisionLogic) GetAdministrativeDivision(req types.GetAdministrativeDivisionReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}