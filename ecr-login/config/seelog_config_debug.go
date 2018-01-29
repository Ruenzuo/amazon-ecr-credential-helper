// +build debug

package config

func configFile(logfile string) string {
  configfile := `
  <seelog type="asyncloop" minlevel="debug">
    <outputs formatid="main">
      <rollingfile filename="` + logfile + `" type="date"
       datepattern="2006-01-02-15" archivetype="none" maxrolls="2" />
      <console />
    </outputs>
    <formats>
      <format id="main" format="%UTCDate(2006-01-02T15:04:05Z07:00) [%LEVEL] %Msg%n" />
    </formats>
  </seelog>
`
  return configfile
}
