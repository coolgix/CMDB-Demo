package impl

import (
	"context"
	"fmt"

	"github.com/CMDB-Demo/apps/host"
)

// 完成对象和sql之间的转换

//把host对象保存到数据库里面
//save的时候为了防止取消的情况需要使用ctx
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {

	var (
		err error
	)

	//检验完成后一次性要往两张表录入数据，我们需要2个操作 要么都成功，要么都失败，事务的逻辑
	//初始化一个事务，所有的操作都使用这个事务来提交
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx erro, %s", err)
	}

	//通过defer语句对事务的提交方式
	//1、无错误，则commit成功
	//2、有报错，则roll back事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback erro,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error.%s", err)
			}
		}
	}()

	//事务处理完了，进行事务的插入
	restmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer restmt.Close()
	_, err = restmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	//插入Describe数据
	dstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)

	return nil
}
