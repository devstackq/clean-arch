package main

import (
// _ "github.com/lib/pq"
)

func main() {
	fmt.Print()

	programmName := filepath.Base(os.Args[0])
	//system log
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programmName)

	if err != nil {
		log.Print(err)
		return
	}

	log.SetOutput(sysLog)
	//write in log journal
	log.Println("some data")

	// args := os.Args
	// log.Print(args)

	// io.WriteString(os.Stdout, args[1])

	// f := *os.NewFile(1, "file")
	// f := &os.File{}
	// f := os.Stdin //every, poka ne ctrl D
	// //os.Args - 1 time call

	// s := bufio.NewScanner(f) //reader - cli, os.Stdin

	// log.Print(string(s.Bytes()), s.Text(), 123, s.Scan())

	// for s.Scan() {
	// 	log.Print(">", s.Text())
	// }
	// log.Print(s)
	// w.Write()

	// if err := config.Init(); err != nil {
	// 	log.Println(err, "viper")
	// 	return
	// }
	// app := server.NewApp()

	// if err := app.Run(viper.GetString("port")); err != nil {
	// 	log.Print(err)
	// 	return
	// }
}

// ssl vpn
// 10 443
// 92.46.185.220
//io.WriteString(os.Stoud, "data") where, what
//Goo