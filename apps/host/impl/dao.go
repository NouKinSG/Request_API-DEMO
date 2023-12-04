package impl

import (
	"context"
	"fmt"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
)

// 完成对象核数据库直接转化
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {

	var err error

	// 3、把数据入库到resource表和host表
	// 一次需要往两张表插入数据，需要使用2个操作  要么成功，要么失败，事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 4、通过defer语句来处理事务的提交和回滚
	// 无错误，则提交  Commit事务
	// 有错误，则回滚  Rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error %s", err)
			}
		}
	}()

	// 5、插入Resource数据
	rstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	//6、Describe 数据
	dstmt, err := tx.Prepare(ins.Description)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)

	if err != nil {
		return err
	}

	return nil
}
