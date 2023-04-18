package impl

//定义sql语句
//插入resource 表
const (
	InsertResourceSQL = `INSERT INTO resource (
		id,
		vendor,
		region,
		zone,
		create_at,
		expire_at,
		category,
		type,
		name,
		description,
		status,
		update_at,
		sync_at,
		sync_accout,
		public_ip,
		private_ip,
		pay_type,
		resource_hash,
		describe_hash 
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

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

	queryHostSQL = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
)
