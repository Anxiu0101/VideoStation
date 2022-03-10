- MapTo 函数执行映射时需要检查设置结构体字段名与配置字段名相符

  ```go
  type Server struct {
  	RunMode      string
  	HttpPort     int
  	ReadTimeout  time.Duration
  	WriteTimeout time.Duration
  }
  
  var ServerSetting = &Server{}
  
  func () {
      cfg, err = ini.Load("conf/app.ini")
  	if err != nil {
  		log.Fatalf("setting.Setup, Fail to parse 'conf/app.ini': %v", err)
  	}
      mapTo("server", ServerSetting)
  }
  
  // mapTo map section
  func mapTo(section string, v interface{}) {
  	err := cfg.Section(section).MapTo(v)
  	if err != nil {
  		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
  	}
  }
  ```

  ```ini
  [server]
  RunMode = debug
  HttpPort = 8000
  ReadTimeout = 60
  WriteTimeout = 60
  ```

  