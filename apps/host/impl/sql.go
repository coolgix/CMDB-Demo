package impl

//定义sql语句
//插入resource 表
const (
	InsertResourceSQL = `INSERT INTO resource (
		id,
		vendor,
		region,
		create_at,
		expire_at,
		type,
		name,
		description,
		status,
		update_at,
		sync_at,
		accout,
		public_ip,
		private_ip,
		pay_type
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	// 使用占位符 是为了prepare语句
	//插入host表
	InsertDescribeSQL = `INSERT INTO host (
		resource_id,
		cpu,
		memory,
		gpu_amount,
		gpu_spec,
		os_type,
		os_name,
		serial_number
	)
	VALUES
		(?,?,?,?,?,?,?,?);`
)
