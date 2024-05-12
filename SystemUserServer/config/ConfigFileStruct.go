package config

type Yaml struct {
	Server   Server   `mapstructure:"Server"`
	DataBase DataBase `mapstructure:"database"`
}

type DataBase struct {
	Url                     string `mapstructure:"url"`                     // 数据库的URL
	Port                    int    `mapstructure:"port"`                    // 数据库服务的端口
	Username                string `mapstructure:"username"`                // 连接数据库的用户名
	Password                string `mapstructure:"password"`                // 连接数据库的密码
	Name                    string `mapstructure:"name"`                    // 数据库的名称
	UrlAdditionalParameters string `mapstructure:"urladditionalparameters"` // URL的额外参数
}

type Server struct {
	ApplicationName string  `mapstructure:"application-name"`
	Address         string  `mapstructure:"address"`
	Port            int     `mapstructure:"port"` // 服务的端口号
	Servlet         Servlet `mapstructure:"servlet"`
}

type Servlet struct {
	ContextPath string `mapstructure:"context-path"` // Servlet 的上下文路径
}
