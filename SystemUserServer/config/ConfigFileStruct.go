package config

type Yaml struct {
	Server   Server   `yaml:"Server"`
	DataBase DataBase `yaml:"database"`
}

type DataBase struct {
	Url                     string `yaml:"url"`                     // 数据库的URL
	Port                    int    `yaml:"port"`                    // 数据库服务的端口
	Username                string `yaml:"username"`                // 连接数据库的用户名
	Password                string `yaml:"password"`                // 连接数据库的密码
	Name                    string `yaml:"name"`                    // 数据库的名称
	UrlAdditionalParameters string `yaml:"urladditionalparameters"` // URL的额外参数
}

type Server struct {
	Address string  `yaml:"address"`
	Port    int     `yaml:"port"` // 服务的端口号
	Servlet Servlet `yaml:"servlet"`
}

type Servlet struct {
	ContextPath string `yaml:"contextpath"` // Servlet 的上下文路径
}
