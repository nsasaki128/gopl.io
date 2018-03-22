package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type ftp struct {
	conn net.Conn
	cwd  string
	addr string
}

func newFtp(conn net.Conn) *ftp {
	return &ftp{
		conn: conn,
		cwd:  os.Getenv("HOME"),
		addr: conn.RemoteAddr().String(),
	}
}

func (ftp *ftp) reply(code int, text string) {
	if text == "" {
		io.WriteString(ftp.conn, fmt.Sprintf("%d\r\n", code))
	} else {
		io.WriteString(ftp.conn, fmt.Sprintf("%d %s\r\n", code, text))
	}
}
func (ftp *ftp) replyCode(code int) {
	ftp.reply(code, "")
}

func (ftp *ftp) run() {
	defer ftp.conn.Close()
	ftp.reply(ServiceReadyForNewUser, "Service ready for new user.")

	scanner := bufio.NewScanner(ftp.conn)
	for scanner.Scan() {
		command := scanner.Text()
		log.Println(command)
		if !ftp.handleConn(command) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error:", err)
	}
}

func (ftp *ftp) handleConn(c string) bool {
	command := strings.Split(c, " ")
	fmt.Println(command[0])

	switch command[0] {
	case "USER":
		ftp.handleUserCommand(c)
	case "TYPE":
		ftp.handleTypeCommand(c)
	case "PASV":
		ftp.handlePasvCommand(c)
	case "EPRT":
		//ftp.handlePasvCommand(c)
		ftp.handleEprtCommand(command[1])
	case "CWD":
		ftp.handleCwdCommand(c)
	case "LIST":
		ftp.handleListCommand(command)
	case "STOR":
		ftp.handleStorCommand(c)
	case "SYST":
		ftp.handleSystCommand(c)
	case "PWD":
		ftp.handlePwdCommand(c)
	case "QUIT":
		ftp.handleQuitCommand(c)
		return false
	default:
		ftp.reply(CommandNotImplemented, "Command not implemented.")
	}
	return true
}

func (ftp *ftp) handleUserCommand(c string) {
	ftp.reply(UserLoggedOn, "User Logged on.")
}

func (ftp *ftp) handleTypeCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}

	typeCode := command[1]
	if typeCode != "A" && typeCode != "I" {
		ftp.reply(CommandNotImplementedForParameter, "Command not implemented for that parameter.")
		return
	}
	ftp.reply(CommandOk, "Command okay.")
}
func (ftp *ftp) handlePasvCommand(c string) {
	listenIP := ftp.addr
	lastIdx := strings.LastIndex(listenIP, ":")
	socket, _ := strconv.Atoi(listenIP[lastIdx+1:])
	quads := strings.Split(listenIP[:lastIdx], ".")

	p1 := socket / 256
	p2 := socket - (p1 * 256)
	target := fmt.Sprintf("(%s,%s,%s,%s,%d,%d)", quads[0], quads[1], quads[2], quads[3], p1, p2)
	msg := "Entering Passive Mode " + target
	fmt.Println(msg)
	ftp.reply(Entering, msg)
}

func (ftp *ftp) handleCwdCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}
	cm := command[1]
	fmt.Println(cm)
	if !strings.HasPrefix(cm, "/") {
		cm = path.Join(ftp.cwd, cm)
	}

	info, err := os.Stat(cm)
	if err != nil || !info.IsDir() {
		ftp.reply(FileUnavailableBusy, "Directory not found.")
		return
	}
	fmt.Println(info.Name())
	ftp.cwd = cm

	ftp.reply(RequestedFileActionOkey, "Requested file action okay, completed.")
}
func (ftp *ftp) handleListCommand(command []string) {
	ftp.replyCode(DataConnectionAlreadyOpen)
	if len(command) == 1 {
		command = nil
	} else {
		command = command[1:]
	}
	cmd := exec.Command("/bin/ls", command...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		ftp.conn.Write([]byte(err.Error()))
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		ftp.conn.Write([]byte(err.Error()))
		return
	}

	if ftp.conn == nil {
		panic("Data connection has not been established")
	}

	a := ascii{ftp.conn, nil}
	go io.Copy(&a, stdout)
	go io.Copy(&a, stderr)
	if err := cmd.Start(); err != nil {
		ftp.conn.Write([]byte(err.Error()))
		return
	}

	cmd.Wait()

	ftp.replyCode(ClosingDataConnection)
}

func (ftp *ftp) handleStorCommand(c string) {
	command := strings.Split(c, " ")
	if len(command) < 2 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
		return
	}
	p := path.Join(ftp.cwd, command[1])

	ftp.reply(FileStatusOkay, "File status okay; about to open data connection.")

	conn, err := net.Dial("tcp", ftp.addr)
	if err != nil {
		log.Println("error: ", err)
		ftp.reply(CantOpenDataConnection, "Can't open data connection.")
		return
	}
	defer conn.Close()

	f, err := os.Create(p)
	if err != nil {
		log.Println("error: ", err)
		ftp.reply(FileUnavailable, "Requested file action not taken.")
		return
	}
	defer f.Close()

	if _, err = io.Copy(f, conn); err != nil {
		log.Println("error: ", err)
		ftp.reply(FileUnavailable, "Requested file action not taken.")
		return
	}

	ftp.reply(RequestedFileActionOkey, "File transfer completed.")
}
func (ftp *ftp) handleSystCommand(c string) {
	ftp.reply(CommandOk, "Simple FTP")
}

func (ftp *ftp) handlePwdCommand(c string) {
	fmt.Println(ftp.cwd)
	ftp.reply(PathNameCreated, fmt.Sprintf(`"%s" is the current directory`, ftp.cwd))
}
func (ftp *ftp) handleQuitCommand(c string) {
	ftp.reply(ServiceClosingTELNETConnection, "Close connection.")
}

func (ftp *ftp) handleEprtCommand(c string) {
	ipInfo := strings.Split(c[1:len(c)-1], c[0:1])
	fmt.Println(ipInfo)
	// protocol type, network address, port
	if len(ipInfo) != 3 {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
	}
	var address string
	switch ipInfo[0] {
	case "1": // IPv4
		address = ipInfo[1] + ":" + ipInfo[2]
	case "2": // IPv6
		address = fmt.Sprintf("[%s]:%s", ipInfo[1], ipInfo[2])
	default:

	}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		ftp.reply(CommandNotImplementedForParameter, "Invalid Parameter.")
	}
	defer conn.Close()

	log.Print("Data Connect established\n")

	ftp.replyCode(CommandOk)
}
